package stacks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
)

type AnotherCdkStackProps struct {
	awscdk.StackProps
}

func NewAnotherCdkStack(scope constructs.Construct, id string, props *AnotherCdkStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// The code that defines your stack goes here

	// example resource
	queue := awssqs.NewQueue(stack, jsii.String("AnotherCdkQueue.fifo"), &awssqs.QueueProps{
		VisibilityTimeout:         awscdk.Duration_Seconds(jsii.Number(300)),
		Fifo:                      jsii.Bool(true),
		ContentBasedDeduplication: jsii.Bool(true),
	})
	_ = queue

	return stack
}
