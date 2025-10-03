package stacks

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/assertions"
	"github.com/aws/jsii-runtime-go"
)

func TestMain(m *testing.M) {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Join(filepath.Dir(filename), "..")
	err := os.Chdir(dir)
	if err != nil {
		log.Fatalf("could not change directory: %v", err)
	}

	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestHelloworldCdkStack(t *testing.T) {
	// GIVEN
	app := awscdk.NewApp(nil)

	// WHEN
	stack := NewHelloWorldStack(app, "UnitTestHelloWorld", nil)

	// THEN
	template := assertions.Template_FromStack(stack, nil)

	t.Run("verify api gateway", func(t *testing.T) {
		template.ResourceCountIs(jsii.String("AWS::ApiGateway::RestApi"), jsii.Number(1))

		template.HasResourceProperties(jsii.String("AWS::ApiGateway::RestApi"), map[string]interface{}{
			"Name": "hello-world",
		})
	})

	t.Run("verify lambda functions", func(t *testing.T) {
		template.ResourceCountIs(jsii.String("AWS::Lambda::Function"), jsii.Number(3))

		template.HasResourceProperties(jsii.String("AWS::Lambda::Function"), map[string]interface{}{
			"FunctionName": "golang",
			"Runtime":      "provided.al2023",
		})
		template.HasResourceProperties(jsii.String("AWS::Lambda::Function"), map[string]interface{}{
			"FunctionName": "nodejs",
			"Runtime":      "nodejs22.x",
		})
		template.HasResourceProperties(jsii.String("AWS::Lambda::Function"), map[string]interface{}{
			"FunctionName": "python",
			"Runtime":      "python3.13",
		})
	})
}
