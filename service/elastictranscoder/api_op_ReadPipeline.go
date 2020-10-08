// Code generated by smithy-go-codegen DO NOT EDIT.

package elastictranscoder

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The ReadPipeline operation gets detailed information about a pipeline.
func (c *Client) ReadPipeline(ctx context.Context, params *ReadPipelineInput, optFns ...func(*Options)) (*ReadPipelineOutput, error) {
	stack := middleware.NewStack("ReadPipeline", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpReadPipelineMiddlewares(stack)
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
	addOpReadPipelineValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opReadPipeline(options.Region), middleware.Before)
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
			OperationName: "ReadPipeline",
			Err:           err,
		}
	}
	out := result.(*ReadPipelineOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The ReadPipelineRequest structure.
type ReadPipelineInput struct {

	// The identifier of the pipeline to read.
	//
	// This member is required.
	Id *string
}

// The ReadPipelineResponse structure.
type ReadPipelineOutput struct {

	// A section of the response body that provides information about the pipeline.
	Pipeline *types.Pipeline

	// Elastic Transcoder returns a warning if the resources used by your pipeline are
	// not in the same region as the pipeline. Using resources in the same region, such
	// as your Amazon S3 buckets, Amazon SNS notification topics, and AWS KMS key,
	// reduces processing time and prevents cross-regional charges.
	Warnings []*types.Warning

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpReadPipelineMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpReadPipeline{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpReadPipeline{}, middleware.After)
}

func newServiceMetadataMiddleware_opReadPipeline(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elastictranscoder",
		OperationName: "ReadPipeline",
	}
}
