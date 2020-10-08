// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Sets the key state of a customer master key (CMK) to enabled. This allows you to
// use the CMK for cryptographic operations
// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations).
// You cannot perform this operation on a CMK in a different AWS account. The CMK
// that you use for this operation must be in a compatible key state. For details,
// see How Key State Affects Use of a Customer Master Key
// (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html) in the
// AWS Key Management Service Developer Guide.
func (c *Client) EnableKey(ctx context.Context, params *EnableKeyInput, optFns ...func(*Options)) (*EnableKeyOutput, error) {
	stack := middleware.NewStack("EnableKey", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpEnableKeyMiddlewares(stack)
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
	addOpEnableKeyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opEnableKey(options.Region), middleware.Before)
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
			OperationName: "EnableKey",
			Err:           err,
		}
	}
	out := result.(*EnableKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EnableKeyInput struct {

	// A unique identifier for the customer master key (CMK). Specify the key ID or the
	// Amazon Resource Name (ARN) of the CMK. For example:
	//
	//     * Key ID:
	// 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//     * Key ARN:
	// arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// To
	// get the key ID and key ARN for a CMK, use ListKeys () or DescribeKey ().
	//
	// This member is required.
	KeyId *string
}

type EnableKeyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpEnableKeyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpEnableKey{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpEnableKey{}, middleware.After)
}

func newServiceMetadataMiddleware_opEnableKey(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "EnableKey",
	}
}
