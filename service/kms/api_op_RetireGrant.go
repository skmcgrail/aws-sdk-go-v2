// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retires a grant. To clean up, you can retire a grant when you're done using it.
// You should revoke a grant when you intend to actively deny operations that
// depend on it. The following are permitted to call this API:
//
//     * The AWS
// account (root user) under which the grant was created
//
//     * The
// RetiringPrincipal, if present in the grant
//
//     * The GranteePrincipal, if
// RetireGrant is an operation specified in the grant
//
// You must identify the grant
// to retire by its grant token or by a combination of the grant ID and the Amazon
// Resource Name (ARN) of the customer master key (CMK). A grant token is a unique
// variable-length base64-encoded string. A grant ID is a 64 character unique
// identifier of a grant. The CreateGrant () operation returns both.
func (c *Client) RetireGrant(ctx context.Context, params *RetireGrantInput, optFns ...func(*Options)) (*RetireGrantOutput, error) {
	stack := middleware.NewStack("RetireGrant", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpRetireGrantMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opRetireGrant(options.Region), middleware.Before)

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
			OperationName: "RetireGrant",
			Err:           err,
		}
	}
	out := result.(*RetireGrantOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RetireGrantInput struct {
	// Token that identifies the grant to be retired.
	GrantToken *string
	// The Amazon Resource Name (ARN) of the CMK associated with the grant. For
	// example:
	// arn:aws:kms:us-east-2:444455556666:key/1234abcd-12ab-34cd-56ef-1234567890ab
	KeyId *string
	// Unique identifier of the grant to retire. The grant ID is returned in the
	// response to a CreateGrant operation.
	//
	//     * Grant ID Example -
	// 0123456789012345678901234567890123456789012345678901234567890123
	GrantId *string
}

type RetireGrantOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpRetireGrantMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpRetireGrant{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpRetireGrant{}, middleware.After)
}

func newServiceMetadataMiddleware_opRetireGrant(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "RetireGrant",
	}
}