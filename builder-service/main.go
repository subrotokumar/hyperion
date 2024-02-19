package main

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"sync"
)

func main() {
	bucketName := os.Getenv("BUCKET_NAME")
	if bucketName == "" {
		panic("Empty Bucket name")
	}

	projectId := os.Getenv("PROJECT_ID")
	if projectId == "" {
		panic("PROJECT_ID can't be empty")
	}

	cwd, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}

	outPutDir := path.Join(cwd, "output")
	distFolderPath := path.Join(outPutDir, "dist")

	err = os.Chdir(outPutDir)
	if err != nil {
		panic(err)
	}

	npmInstallCmd := exec.Command("npm", "install")
	output, err := npmInstallCmd.CombinedOutput()
	if err != nil {
		panic(err.Error())
	}
	println(string(output))

	npmBuildCmd := exec.Command("npm", "run", "build")
	output, err = npmBuildCmd.CombinedOutput()
	if err != nil {
		panic(err.Error())
	}
	println(string(output))

	if err := initAWS(); err != nil {
		panic(err.Error())
	}

	var wg = &sync.WaitGroup{}

	handleFile := func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			print("Uploading: ", path)

			relativePath := strings.ReplaceAll(path, distFolderPath, "")
			objectKey := fmt.Sprintf("__outputs/%s%s", projectId, relativePath)
			if err := awsClient.UploadFile(bucketName, objectKey, path); err != nil {
				println(" ⚠️")
				return
			}
			println(" ✅")
		}(wg)
		return nil
	}
	wg.Wait()

	if err := filepath.Walk(distFolderPath, handleFile); err != nil {
		println("Error", err)
	}
	print("Done 🔥")
}
