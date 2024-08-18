package main

import (
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"

	// "github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type AuthGoStackProps struct {
	awscdk.StackProps
}

func NewAuthGoStack(scope constructs.Construct, id string, props *AuthGoStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// The code that defines your stack goes here

	// Define a new lambda function
	lambda := awslambda.NewFunction(stack, jsii.String("HandlerRequest"), &awslambda.FunctionProps{
		Runtime: awslambda.Runtime_PROVIDED_AL2023(),
		Code: awslambda.Code_FromAsset(jsii.String("handler/main.zip"), nil),
		Handler: jsii.String("main"), // this is the name of function in the handler/main.go file
	})

	// create DynamoDB table
	table := awsdynamodb.NewTable(stack, jsii.String("UsersTable"), &awsdynamodb.TableProps{
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("username"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		TableName: jsii.String("Users"),
	})

	// grant the lambda function read/write permissions to the table
	table.GrantReadWriteData(lambda)


	
	return stack
}

func main() {
	defer jsii.Close() // this will close the JSII runtime when the program exits. defer is used to ensure this line is executed after everything is done

	app := awscdk.NewApp(nil)  // this will create a new CDK app with no context

	NewAuthGoStack(app, "AuthGoStack", &AuthGoStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	// return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String("123456789012"),
	//  Region:  jsii.String("us-east-1"),
	// }

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	return &awscdk.Environment{
	 Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	 Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	}
}