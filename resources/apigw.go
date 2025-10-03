package resources

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"

	"helloworld-cdk/tools"
)

type ApiGatewayParams struct {
	NamePrefix string
	Routes     []Route
}

type Route struct {
	Path           string
	Method         string
	LambdaFunction awslambda.Function
}

func ApiGateway(scope constructs.Construct, params ApiGatewayParams) awsapigateway.RestApi {
	name := params.NamePrefix + "hello-world"

	// API GW definition
	api := awsapigateway.NewRestApi(scope, jsii.String(name), &awsapigateway.RestApiProps{
		RestApiName: jsii.String(name),
		Description: jsii.String("Hello World REST API"),
		DefaultCorsPreflightOptions: &awsapigateway.CorsOptions{
			AllowOrigins: awsapigateway.Cors_ALL_ORIGINS(),
			AllowMethods: awsapigateway.Cors_ALL_METHODS(),
			AllowHeaders: jsii.Strings("Content-Type", "Authorization"),
		},
		Policy: awsiam.NewPolicyDocument(&awsiam.PolicyDocumentProps{
			Statements: &[]awsiam.PolicyStatement{
				awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
					Actions:   jsii.Strings("execute-api:Invoke"),
					Resources: jsii.Strings("*"),
					Principals: &[]awsiam.IPrincipal{
						awsiam.NewAnyPrincipal(),
					},
					Conditions: &map[string]interface{}{
						"IpAddress": map[string]interface{}{
							"aws:SourceIp": tools.MustGetExternalIP(nil) + "/32",
						},
					},
				}),
			},
		}),
	})

	// Lambda integrations
	for _, route := range params.Routes {
		lambdaIntegration := awsapigateway.NewLambdaIntegration(route.LambdaFunction, &awsapigateway.LambdaIntegrationOptions{
			RequestTemplates: &map[string]*string{
				"application/json": jsii.String(`{"statusCode": "200"}`),
			},
		})
		api.Root().AddResource(jsii.String(route.Path), nil).AddMethod(jsii.String(route.Method), lambdaIntegration, nil)
	}

	return api
}
