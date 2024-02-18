package aws

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
)

type AwsClient struct {
	Ecs *ecs.Client
}

var awsClient *AwsClient

func initAwsService() error {
	sdkConfig, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		return fmt.Errorf("Couldn't load default configuration. Have you set up your AWS account?")
	}
	ecsClient := ecs.NewFromConfig(sdkConfig)
	awsClient = &AwsClient{
		Ecs: ecsClient,
	}
	return nil

}

func (awsClient *AwsClient) RunTask(ctx context.Context, githubRepoUrl, projectId string) (*ecs.RunTaskOutput, error) {
	builderImage := os.Getenv("BUILDER_IMAGE")
	taskDefination := os.Getenv("TASK_DEFINATION")
	cluster := os.Getenv("CLUSTER")
	return awsClient.Ecs.RunTask(
		ctx,
		&ecs.RunTaskInput{
			TaskDefinition: aws.String(taskDefination),
			Cluster:        aws.String(cluster),
			LaunchType:     types.LaunchTypeFargate,
			Count:          aws.Int32(1),
			NetworkConfiguration: &types.NetworkConfiguration{
				AwsvpcConfiguration: &types.AwsVpcConfiguration{
					Subnets:        []string{},
					SecurityGroups: []string{},
					AssignPublicIp: types.AssignPublicIpEnabled,
				},
			},
			Overrides: &types.TaskOverride{
				ContainerOverrides: []types.ContainerOverride{
					{
						Name: aws.String(builderImage),
						Environment: []types.KeyValuePair{
							{
								Name:  aws.String("GITHUB_REPO_URL"),
								Value: aws.String(githubRepoUrl),
							},
							{
								Name:  aws.String("PROJECT_ID"),
								Value: aws.String(projectId),
							},
						},
					},
				},
			},
		},
	)
}
