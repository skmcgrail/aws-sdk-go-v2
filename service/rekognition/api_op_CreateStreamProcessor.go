// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an Amazon Rekognition stream processor that you can use to detect and
// recognize faces in a streaming video. Amazon Rekognition Video is a consumer of
// live video from Amazon Kinesis Video Streams. Amazon Rekognition Video sends
// analysis results to Amazon Kinesis Data Streams. You provide as input a Kinesis
// video stream (Input) and a Kinesis data stream (Output) stream. You also specify
// the face recognition criteria in Settings. For example, the collection
// containing faces that you want to recognize. Use Name to assign an identifier
// for the stream processor. You use Name to manage the stream processor. For
// example, you can start processing the source video by calling
// StartStreamProcessor () with the Name field. After you have finished analyzing a
// streaming video, use StopStreamProcessor () to stop processing. You can delete
// the stream processor by calling DeleteStreamProcessor ().
func (c *Client) CreateStreamProcessor(ctx context.Context, params *CreateStreamProcessorInput, optFns ...func(*Options)) (*CreateStreamProcessorOutput, error) {
	stack := middleware.NewStack("CreateStreamProcessor", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateStreamProcessorMiddlewares(stack)
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
	addOpCreateStreamProcessorValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateStreamProcessor(options.Region), middleware.Before)

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
			OperationName: "CreateStreamProcessor",
			Err:           err,
		}
	}
	out := result.(*CreateStreamProcessorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateStreamProcessorInput struct {
	// Face recognition input parameters to be used by the stream processor. Includes
	// the collection to use for face recognition and the face attributes to detect.
	Settings *types.StreamProcessorSettings
	// An identifier you assign to the stream processor. You can use Name to manage the
	// stream processor. For example, you can get the current status of the stream
	// processor by calling DescribeStreamProcessor (). Name is idempotent.
	Name *string
	// Kinesis data stream stream to which Amazon Rekognition Video puts the analysis
	// results. If you are using the AWS CLI, the parameter name is
	// StreamProcessorOutput.
	Output *types.StreamProcessorOutput
	// Kinesis video stream stream that provides the source streaming video. If you are
	// using the AWS CLI, the parameter name is StreamProcessorInput.
	Input *types.StreamProcessorInput
	// ARN of the IAM role that allows access to the stream processor.
	RoleArn *string
}

type CreateStreamProcessorOutput struct {
	// ARN for the newly create stream processor.
	StreamProcessorArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateStreamProcessorMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateStreamProcessor{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateStreamProcessor{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateStreamProcessor(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "CreateStreamProcessor",
	}
}