// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Provides information about a stream processor created by CreateStreamProcessor
// (). You can get information about the input and output streams, the input
// parameters for the face recognition being performed, and the current status of
// the stream processor.
func (c *Client) DescribeStreamProcessor(ctx context.Context, params *DescribeStreamProcessorInput, optFns ...func(*Options)) (*DescribeStreamProcessorOutput, error) {
	stack := middleware.NewStack("DescribeStreamProcessor", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeStreamProcessorMiddlewares(stack)
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
	addOpDescribeStreamProcessorValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeStreamProcessor(options.Region), middleware.Before)
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
			OperationName: "DescribeStreamProcessor",
			Err:           err,
		}
	}
	out := result.(*DescribeStreamProcessorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeStreamProcessorInput struct {

	// Name of the stream processor for which you want information.
	//
	// This member is required.
	Name *string
}

type DescribeStreamProcessorOutput struct {

	// Date and time the stream processor was created
	CreationTimestamp *time.Time

	// Kinesis video stream that provides the source streaming video.
	Input *types.StreamProcessorInput

	// The time, in Unix format, the stream processor was last updated. For example,
	// when the stream processor moves from a running state to a failed state, or when
	// the user starts or stops the stream processor.
	LastUpdateTimestamp *time.Time

	// Name of the stream processor.
	Name *string

	// Kinesis data stream to which Amazon Rekognition Video puts the analysis results.
	Output *types.StreamProcessorOutput

	// ARN of the IAM role that allows access to the stream processor.
	RoleArn *string

	// Face recognition input parameters that are being used by the stream processor.
	// Includes the collection to use for face recognition and the face attributes to
	// detect.
	Settings *types.StreamProcessorSettings

	// Current status of the stream processor.
	Status types.StreamProcessorStatus

	// Detailed status message about the stream processor.
	StatusMessage *string

	// ARN of the stream processor.
	StreamProcessorArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeStreamProcessorMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeStreamProcessor{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeStreamProcessor{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeStreamProcessor(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "DescribeStreamProcessor",
	}
}
