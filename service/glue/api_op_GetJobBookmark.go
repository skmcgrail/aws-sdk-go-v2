// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information on a job bookmark entry.
func (c *Client) GetJobBookmark(ctx context.Context, params *GetJobBookmarkInput, optFns ...func(*Options)) (*GetJobBookmarkOutput, error) {
	stack := middleware.NewStack("GetJobBookmark", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetJobBookmarkMiddlewares(stack)
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
	addOpGetJobBookmarkValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetJobBookmark(options.Region), middleware.Before)
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
			OperationName: "GetJobBookmark",
			Err:           err,
		}
	}
	out := result.(*GetJobBookmarkOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetJobBookmarkInput struct {

	// The name of the job in question.
	//
	// This member is required.
	JobName *string

	// The unique run identifier associated with this job run.
	RunId *string
}

type GetJobBookmarkOutput struct {

	// A structure that defines a point that a job can resume processing.
	JobBookmarkEntry *types.JobBookmarkEntry

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetJobBookmarkMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetJobBookmark{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetJobBookmark{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetJobBookmark(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "GetJobBookmark",
	}
}
