// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the estimated monthly cost of a template. The return value is an AWS
// Simple Monthly Calculator URL with a query string that describes the resources
// required to run the template.
func (c *Client) EstimateTemplateCost(ctx context.Context, params *EstimateTemplateCostInput, optFns ...func(*Options)) (*EstimateTemplateCostOutput, error) {
	stack := middleware.NewStack("EstimateTemplateCost", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpEstimateTemplateCostMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opEstimateTemplateCost(options.Region), middleware.Before)
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
			OperationName: "EstimateTemplateCost",
			Err:           err,
		}
	}
	out := result.(*EstimateTemplateCostOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for an EstimateTemplateCost () action.
type EstimateTemplateCostInput struct {

	// A list of Parameter structures that specify input parameters.
	Parameters []*types.Parameter

	// Structure containing the template body with a minimum length of 1 byte and a
	// maximum length of 51,200 bytes. (For more information, go to Template Anatomy
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-anatomy.html)
	// in the AWS CloudFormation User Guide.) Conditional: You must pass TemplateBody
	// or TemplateURL. If both are passed, only TemplateBody is used.
	TemplateBody *string

	// Location of file containing the template body. The URL must point to a template
	// that is located in an Amazon S3 bucket. For more information, go to Template
	// Anatomy
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-anatomy.html)
	// in the AWS CloudFormation User Guide. Conditional: You must pass TemplateURL or
	// TemplateBody. If both are passed, only TemplateBody is used.
	TemplateURL *string
}

// The output for a EstimateTemplateCost () action.
type EstimateTemplateCostOutput struct {

	// An AWS Simple Monthly Calculator URL with a query string that describes the
	// resources required to run the template.
	Url *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpEstimateTemplateCostMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpEstimateTemplateCost{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpEstimateTemplateCost{}, middleware.After)
}

func newServiceMetadataMiddleware_opEstimateTemplateCost(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "EstimateTemplateCost",
	}
}
