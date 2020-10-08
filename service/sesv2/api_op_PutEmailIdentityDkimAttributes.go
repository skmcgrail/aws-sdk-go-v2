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

// Used to enable or disable DKIM authentication for an email identity.
func (c *Client) PutEmailIdentityDkimAttributes(ctx context.Context, params *PutEmailIdentityDkimAttributesInput, optFns ...func(*Options)) (*PutEmailIdentityDkimAttributesOutput, error) {
	stack := middleware.NewStack("PutEmailIdentityDkimAttributes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpPutEmailIdentityDkimAttributesMiddlewares(stack)
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
	addOpPutEmailIdentityDkimAttributesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutEmailIdentityDkimAttributes(options.Region), middleware.Before)
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
			OperationName: "PutEmailIdentityDkimAttributes",
			Err:           err,
		}
	}
	out := result.(*PutEmailIdentityDkimAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to enable or disable DKIM signing of email that you send from an email
// identity.
type PutEmailIdentityDkimAttributesInput struct {

	// The email identity that you want to change the DKIM settings for.
	//
	// This member is required.
	EmailIdentity *string

	// Sets the DKIM signing configuration for the identity. When you set this value
	// true, then the messages that are sent from the identity are signed using DKIM.
	// If you set this value to false, your messages are sent without DKIM signing.
	SigningEnabled *bool
}

// An HTTP 200 response if the request succeeds, or an error message if the request
// fails.
type PutEmailIdentityDkimAttributesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpPutEmailIdentityDkimAttributesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpPutEmailIdentityDkimAttributes{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpPutEmailIdentityDkimAttributes{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutEmailIdentityDkimAttributes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "PutEmailIdentityDkimAttributes",
	}
}
