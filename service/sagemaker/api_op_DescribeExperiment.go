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
	"time"
)

// Provides a list of an experiment's properties.
func (c *Client) DescribeExperiment(ctx context.Context, params *DescribeExperimentInput, optFns ...func(*Options)) (*DescribeExperimentOutput, error) {
	stack := middleware.NewStack("DescribeExperiment", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeExperimentMiddlewares(stack)
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
	addOpDescribeExperimentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeExperiment(options.Region), middleware.Before)
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
			OperationName: "DescribeExperiment",
			Err:           err,
		}
	}
	out := result.(*DescribeExperimentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeExperimentInput struct {

	// The name of the experiment to describe.
	//
	// This member is required.
	ExperimentName *string
}

type DescribeExperimentOutput struct {

	// Who created the experiment.
	CreatedBy *types.UserContext

	// When the experiment was created.
	CreationTime *time.Time

	// The description of the experiment.
	Description *string

	// The name of the experiment as displayed. If DisplayName isn't specified,
	// ExperimentName is displayed.
	DisplayName *string

	// The Amazon Resource Name (ARN) of the experiment.
	ExperimentArn *string

	// The name of the experiment.
	ExperimentName *string

	// Who last modified the experiment.
	LastModifiedBy *types.UserContext

	// When the experiment was last modified.
	LastModifiedTime *time.Time

	// The ARN of the source and, optionally, the type.
	Source *types.ExperimentSource

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeExperimentMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeExperiment{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeExperiment{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeExperiment(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DescribeExperiment",
	}
}
