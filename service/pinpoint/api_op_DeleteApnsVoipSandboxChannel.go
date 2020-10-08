// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpoint

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disables the APNs VoIP sandbox channel for an application and deletes any
// existing settings for the channel.
func (c *Client) DeleteApnsVoipSandboxChannel(ctx context.Context, params *DeleteApnsVoipSandboxChannelInput, optFns ...func(*Options)) (*DeleteApnsVoipSandboxChannelOutput, error) {
	stack := middleware.NewStack("DeleteApnsVoipSandboxChannel", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteApnsVoipSandboxChannelMiddlewares(stack)
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
	addOpDeleteApnsVoipSandboxChannelValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteApnsVoipSandboxChannel(options.Region), middleware.Before)
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
			OperationName: "DeleteApnsVoipSandboxChannel",
			Err:           err,
		}
	}
	out := result.(*DeleteApnsVoipSandboxChannelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteApnsVoipSandboxChannelInput struct {

	// The unique identifier for the application. This identifier is displayed as the
	// Project ID on the Amazon Pinpoint console.
	//
	// This member is required.
	ApplicationId *string
}

type DeleteApnsVoipSandboxChannelOutput struct {

	// Provides information about the status and settings of the APNs (Apple Push
	// Notification service) VoIP sandbox channel for an application.
	//
	// This member is required.
	APNSVoipSandboxChannelResponse *types.APNSVoipSandboxChannelResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteApnsVoipSandboxChannelMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteApnsVoipSandboxChannel{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteApnsVoipSandboxChannel{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteApnsVoipSandboxChannel(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mobiletargeting",
		OperationName: "DeleteApnsVoipSandboxChannel",
	}
}
