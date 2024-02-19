package env

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

type EnvVar struct {
	PORT               int
	AppEnv             string
	GithubKey          string
	GithubSecret       string
	SessionKey         string
	BuilderName        string
	TaskDefination     string
	Cluster            string
	AwsRegion          string
	BucketName         string
	DatabaseUrl        string
	RefreshTokenSecret string
	AccessTokenSecret  string
}

var Env *EnvVar

func LoadEnv(filename string) error {
	if err := godotenv.Load(filename); err != nil {
		return err
	}
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return err
	}
	appEnv := os.Getenv("AppEnv")
	githubKey := os.Getenv("GITHUB_KEY")
	githubSecret := os.Getenv("GITHUB_SECRET")
	sessionKey := os.Getenv("SESSION_KEY")
	builderName := os.Getenv("BUILDER_IMAGE")
	taskDefination := os.Getenv("TASK_DEFINATION")
	cluster := os.Getenv("CLUSTER")
	awsRegion := os.Getenv("AWS_REGION")
	bucketName := os.Getenv("BUCKET_NAME")
	databaseUrl := os.Getenv("DB_URL")
	accessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")
	refreshTokenSecret := os.Getenv("REFRESH_TOKEN_SECRET")

	Env = &EnvVar{
		PORT:               port,
		AppEnv:             appEnv,
		GithubKey:          githubKey,
		GithubSecret:       githubSecret,
		SessionKey:         sessionKey,
		BuilderName:        builderName,
		TaskDefination:     taskDefination,
		Cluster:            cluster,
		AwsRegion:          awsRegion,
		BucketName:         bucketName,
		DatabaseUrl:        databaseUrl,
		RefreshTokenSecret: refreshTokenSecret,
		AccessTokenSecret:  accessTokenSecret,
	}
	return nil
}
