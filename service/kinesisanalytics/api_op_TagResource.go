// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisanalytics

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kinesisanalytics/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds one or more key-value tags to a Kinesis Analytics application. Note that
// the maximum number of application tags includes system tags. The maximum number
// of user-defined application tags is 50. For more information, see Using Tagging
// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-tagging.html).
func (c *Client) TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) {
	stack := middleware.NewStack("TagResource", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpTagResourceMiddlewares(stack)
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
	addOpTagResourceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opTagResource(options.Region), middleware.Before)
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
			OperationName: "TagResource",
			Err:           err,
		}
	}
	out := result.(*TagResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TagResourceInput struct {

	// The ARN of the application to assign the tags.
	//
	// This member is required.
	ResourceARN *string

	// The key-value tags to assign to the application.
	//
	// This member is required.
	Tags []*types.Tag
}

type TagResourceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpTagResourceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpTagResource{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpTagResource{}, middleware.After)
}

func newServiceMetadataMiddleware_opTagResource(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisanalytics",
		OperationName: "TagResource",
	}
}
