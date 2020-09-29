// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes the specified tags from the role. For more information about tagging,
// see Tagging IAM Identities
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html) in the IAM User
// Guide.
func (c *Client) UntagRole(ctx context.Context, params *UntagRoleInput, optFns ...func(*Options)) (*UntagRoleOutput, error) {
	stack := middleware.NewStack("UntagRole", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpUntagRoleMiddlewares(stack)
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
	addOpUntagRoleValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUntagRole(options.Region), middleware.Before)
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
			OperationName: "UntagRole",
			Err:           err,
		}
	}
	out := result.(*UntagRoleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UntagRoleInput struct {
	// The name of the IAM role from which you want to remove tags. This parameter
	// accepts (through its regex pattern (http://wikipedia.org/wiki/regex)) a string
	// of characters that consist of upper and lowercase alphanumeric characters with
	// no spaces. You can also include any of the following characters: _+=,.@-
	RoleName *string
	// A list of key names as a simple array of strings. The tags with matching keys
	// are removed from the specified role.
	TagKeys []*string
}

type UntagRoleOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpUntagRoleMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpUntagRole{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpUntagRole{}, middleware.After)
}

func newServiceMetadataMiddleware_opUntagRole(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "UntagRole",
	}
}