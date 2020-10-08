// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a flow definition.
func (c *Client) CreateFlowDefinition(ctx context.Context, params *CreateFlowDefinitionInput, optFns ...func(*Options)) (*CreateFlowDefinitionOutput, error) {
	stack := middleware.NewStack("CreateFlowDefinition", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateFlowDefinitionMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpCreateFlowDefinitionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFlowDefinition(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "CreateFlowDefinition",
			Err:           err,
		}
	}
	out := result.(*CreateFlowDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateFlowDefinitionInput struct {

	// The name of your flow definition.
	//
	// This member is required.
	FlowDefinitionName *string

	// An object containing information about the tasks the human reviewers will
	// perform.
	//
	// This member is required.
	HumanLoopConfig *types.HumanLoopConfig

	// An object containing information about where the human review results will be
	// uploaded.
	//
	// This member is required.
	OutputConfig *types.FlowDefinitionOutputConfig

	// The Amazon Resource Name (ARN) of the role needed to call other services on your
	// behalf. For example,
	// arn:aws:iam::1234567890:role/service-role/AmazonSageMaker-ExecutionRole-20180111T151298.
	//
	// This member is required.
	RoleArn *string

	// An object containing information about the events that trigger a human workflow.
	HumanLoopActivationConfig *types.HumanLoopActivationConfig

	// Container for configuring the source of human task requests. Use to specify if
	// Amazon Rekognition or Amazon Textract is used as an integration source.
	HumanLoopRequestSource *types.HumanLoopRequestSource

	// An array of key-value pairs that contain metadata to help you categorize and
	// organize a flow definition. Each tag consists of a key and a value, both of
	// which you define.
	Tags []*types.Tag
}

type CreateFlowDefinitionOutput struct {

	// The Amazon Resource Name (ARN) of the flow definition you create.
	//
	// This member is required.
	FlowDefinitionArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateFlowDefinitionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateFlowDefinition{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateFlowDefinition{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateFlowDefinition(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "CreateFlowDefinition",
	}
}
