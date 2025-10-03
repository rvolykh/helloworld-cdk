package stacks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"

	"helloworld-cdk/config"
	"helloworld-cdk/resources"
)

type HelloWorldStackProps struct {
	awscdk.StackProps
}

func NewHelloWorldStack(scope constructs.Construct, id string, props *HelloWorldStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	config.Tags(stack.Tags())
	cfg := config.Config()

	resources.ApiGateway(stack, resources.ApiGatewayParams{
		NamePrefix: cfg.NamePrefix,
		Routes: []resources.Route{
			{
				Path: "python", Method: "GET",
				LambdaFunction: resources.HelloWorldPythonLambda(
					stack, resources.LambdaParams{NamePrefix: cfg.NamePrefix},
				),
			},
			{
				Path: "golang", Method: "GET",
				LambdaFunction: resources.HelloWorldGoLambda(
					stack, resources.LambdaParams{NamePrefix: cfg.NamePrefix},
				),
			},
			{
				Path: "nodejs", Method: "GET",
				LambdaFunction: resources.HelloWorldNodeJSLambda(
					stack, resources.LambdaParams{NamePrefix: cfg.NamePrefix},
				),
			},
		},
	})

	return stack
}
