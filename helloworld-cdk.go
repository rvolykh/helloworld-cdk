package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"

	"helloworld-cdk/config"
	"helloworld-cdk/stacks"
)

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	stacks.NewHelloworldCdkStack(app, "HelloworldCdkStack", &stacks.HelloworldCdkStackProps{
		StackProps: awscdk.StackProps{
			Env: config.Env(),
		},
	})

	stacks.NewAnotherCdkStack(app, "AnotherCdkStack", &stacks.AnotherCdkStackProps{
		StackProps: awscdk.StackProps{
			Env: config.Env(),
		},
	})

	app.Synth(nil)
}
