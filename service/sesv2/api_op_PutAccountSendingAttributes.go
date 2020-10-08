// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Enable or disable the ability of your account to send email.
func (c *Client) PutAccountSendingAttributes(ctx context.Context, params *PutAccountSendingAttributesInput, optFns ...func(*Options)) (*PutAccountSendingAttributesOutput, error) {
	stack := middleware.NewStack("PutAccountSendingAttributes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpPutAccountSendingAttributesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutAccountSendingAttributes(options.Region), middleware.Before)
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
			OperationName: "PutAccountSendingAttributes",
			Err:           err,
		}
	}
	out := result.(*PutAccountSendingAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to change the ability of your account to send email.
type PutAccountSendingAttributesInput struct {

	// Enables or disables your account's ability to send email. Set to true to enable
	// email sending, or set to false to disable email sending. If AWS paused your
	// account's ability to send email, you can't use this operation to resume your
	// account's ability to send email.
	SendingEnabled *bool
}

// An HTTP 200 response if the request succeeds, or an error message if the request
// fails.
type PutAccountSendingAttributesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpPutAccountSendingAttributesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpPutAccountSendingAttributes{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpPutAccountSendingAttributes{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutAccountSendingAttributes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "PutAccountSendingAttributes",
	}
}
