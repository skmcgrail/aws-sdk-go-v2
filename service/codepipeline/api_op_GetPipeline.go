// Code generated by smithy-go-codegen DO NOT EDIT.

package codepipeline

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the metadata, structure, stages, and actions of a pipeline. Can be used
// to return the entire structure of a pipeline in JSON format, which can then be
// modified and used to update the pipeline structure with UpdatePipeline ().
func (c *Client) GetPipeline(ctx context.Context, params *GetPipelineInput, optFns ...func(*Options)) (*GetPipelineOutput, error) {
	stack := middleware.NewStack("GetPipeline", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetPipelineMiddlewares(stack)
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
	addOpGetPipelineValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetPipeline(options.Region), middleware.Before)

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
			OperationName: "GetPipeline",
			Err:           err,
		}
	}
	out := result.(*GetPipelineOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a GetPipeline action.
type GetPipelineInput struct {
	// The version number of the pipeline. If you do not specify a version, defaults to
	// the current version.
	Version *int32
	// The name of the pipeline for which you want to get information. Pipeline names
	// must be unique under an AWS user account.
	Name *string
}

// Represents the output of a GetPipeline action.
type GetPipelineOutput struct {
	// Represents the pipeline metadata information returned as part of the output of a
	// GetPipeline action.
	Metadata *types.PipelineMetadata
	// Represents the structure of actions and stages to be performed in the pipeline.
	Pipeline *types.PipelineDeclaration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetPipelineMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetPipeline{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetPipeline{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetPipeline(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codepipeline",
		OperationName: "GetPipeline",
	}
}