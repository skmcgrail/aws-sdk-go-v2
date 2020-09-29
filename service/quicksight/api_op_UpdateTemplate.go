// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates a template from an existing Amazon QuickSight analysis or another
// template.
func (c *Client) UpdateTemplate(ctx context.Context, params *UpdateTemplateInput, optFns ...func(*Options)) (*UpdateTemplateOutput, error) {
	stack := middleware.NewStack("UpdateTemplate", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateTemplateMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpUpdateTemplateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateTemplate(options.Region), middleware.Before)
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
			OperationName: "UpdateTemplate",
			Err:           err,
		}
	}
	out := result.(*UpdateTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateTemplateInput struct {
	// The entity that you are using as a source when you update the template. In
	// SourceEntity, you specify the type of object you're using as source:
	// SourceTemplate for a template or SourceAnalysis for an analysis. Both of these
	// require an Amazon Resource Name (ARN). For SourceTemplate, specify the ARN of
	// the source template. For SourceAnalysis, specify the ARN of the source analysis.
	// The SourceTemplate ARN can contain any AWS Account and any QuickSight-supported
	// AWS Region. Use the DataSetReferences entity within SourceTemplate or
	// SourceAnalysis to list the replacement datasets for the placeholders listed in
	// the original. The schema in each dataset must match its placeholder.
	SourceEntity *types.TemplateSourceEntity
	// A description of the current template version that is being updated. Every time
	// you call UpdateTemplate, you create a new version of the template. Each version
	// of the template maintains a description of the version in the VersionDescription
	// field.
	VersionDescription *string
	// The ID for the template.
	TemplateId *string
	// The ID of the AWS account that contains the template that you're updating.
	AwsAccountId *string
	// The name for the template.
	Name *string
}

type UpdateTemplateOutput struct {
	// The creation status of the template.
	CreationStatus types.ResourceStatus
	// The Amazon Resource Name (ARN) for the template.
	Arn *string
	// The ARN for the template, including the version information of the first
	// version.
	VersionArn *string
	// The AWS request ID for this operation.
	RequestId *string
	// The ID for the template.
	TemplateId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateTemplateMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateTemplate{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateTemplate{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateTemplate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "UpdateTemplate",
	}
}