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

	stacks.NewHelloWorldStack(app, "HelloworldCdkStack", &stacks.HelloWorldStackProps{
		StackProps: awscdk.StackProps{
			Env: config.Env(),
		},
	})

	app.Synth(nil)
}
