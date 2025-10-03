package config

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
)

func Tags(tm awscdk.TagManager) awscdk.TagManager {
	tm.SetTag(jsii.String("ManagedBy"), jsii.String("CDK"), jsii.Number(1), jsii.Bool(true))
	return tm
}
