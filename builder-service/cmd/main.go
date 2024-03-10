package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"sync"

	"github.com/joho/godotenv"
	"github.com/subrotokumar/builder-service/internal/cloud"
)

func main() {
	godotenv.Load()
	bucketName := os.Getenv("BUCKET_NAME")
	if bucketName == "" {
		log.Panic("Empty Bucket name")
	}

	projectId := os.Getenv("PROJECT_ID")
	if projectId == "" {
		log.Panic("PROJECT_ID can't be empty")
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Panic(err.Error())
	}

	outPutDir := path.Join(cwd, "output")
	distFolderPath := path.Join(outPutDir, "dist")

	err = os.Chdir(outPutDir)
	if err != nil {
		log.Panic(err)
	}

	npmInstallCmd := exec.Command("npm", "install")
	output, err := npmInstallCmd.CombinedOutput()
	if err != nil {
		log.Panic(err.Error())
	}
	println(string(output))

	npmBuildCmd := exec.Command("npm", "run", "build")
	output, err = npmBuildCmd.CombinedOutput()
	if err != nil {
		log.Panic(err.Error())
	}
	println(string(output))

	awsClient, err := cloud.GetAwsClient()
	if err != nil {
		log.Panic(err.Error())
	}

	var wg = &sync.WaitGroup{}

	handleFile := func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		// wg.Add(1)
		// go func(wg *sync.WaitGroup) {
		// 	defer wg.Done()
		print("Uploading: ", path)

		relativePath := strings.ReplaceAll(path, distFolderPath, "")
		objectKey := fmt.Sprintf("__outputs/%s%s", projectId, relativePath)
		if err := awsClient.UploadFile(bucketName, objectKey, path); err != nil {
			println(" ⚠️")
		} else {
			println(" ✅")
		}
		// }(wg)
		return nil
	}
	wg.Wait()

	if err := filepath.Walk(distFolderPath, handleFile); err != nil {
		println("Error", err)
	}
	print("Done 🔥")
}
