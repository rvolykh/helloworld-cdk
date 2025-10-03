package resources

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	gocdk "github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type LambdaParams struct {
	NamePrefix string
}

func HelloWorldNodeJSLambda(scope constructs.Construct, params LambdaParams) awslambda.Function {
	name := params.NamePrefix + "nodejs"

	return awslambda.NewFunction(scope, jsii.String(name), &awslambda.FunctionProps{
		Runtime:      awslambda.Runtime_NODEJS_22_X(),
		Handler:      jsii.String("handler.handler"),
		Code:         awslambda.Code_FromAsset(jsii.String("resources/src/nodejs/"), nil),
		FunctionName: jsii.String(name),
	})
}

func HelloWorldPythonLambda(scope constructs.Construct, params LambdaParams) awslambda.Function {
	name := params.NamePrefix + "python"

	return awslambda.NewFunction(scope, jsii.String(name), &awslambda.FunctionProps{
		Runtime:      awslambda.Runtime_PYTHON_3_13(),
		Handler:      jsii.String("handler.handler"),
		Code:         awslambda.Code_FromAsset(jsii.String("resources/src/python/"), nil),
		FunctionName: jsii.String(name),
	})
}

func HelloWorldGoLambda(scope constructs.Construct, params LambdaParams) awslambda.Function {
	name := params.NamePrefix + "golang"

	return gocdk.NewGoFunction(scope, jsii.String(name), &gocdk.GoFunctionProps{
		Entry:        jsii.String("resources/src/golang/handler.go"),
		FunctionName: jsii.String(name),
		Runtime:      awslambda.Runtime_PROVIDED_AL2023(),
	})
}
