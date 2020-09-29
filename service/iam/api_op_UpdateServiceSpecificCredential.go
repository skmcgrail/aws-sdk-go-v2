// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Sets the status of a service-specific credential to Active or Inactive.
// Service-specific credentials that are inactive cannot be used for authentication
// to the service. This operation can be used to disable a user's service-specific
// credential as part of a credential rotation work flow.
func (c *Client) UpdateServiceSpecificCredential(ctx context.Context, params *UpdateServiceSpecificCredentialInput, optFns ...func(*Options)) (*UpdateServiceSpecificCredentialOutput, error) {
	stack := middleware.NewStack("UpdateServiceSpecificCredential", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpUpdateServiceSpecificCredentialMiddlewares(stack)
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
	addOpUpdateServiceSpecificCredentialValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateServiceSpecificCredential(options.Region), middleware.Before)
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
			OperationName: "UpdateServiceSpecificCredential",
			Err:           err,
		}
	}
	out := result.(*UpdateServiceSpecificCredentialOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateServiceSpecificCredentialInput struct {
	// The status to be assigned to the service-specific credential.
	Status types.StatusType
	// The name of the IAM user associated with the service-specific credential. If you
	// do not specify this value, then the operation assumes the user whose credentials
	// are used to call the operation. This parameter allows (through its regex pattern
	// (http://wikipedia.org/wiki/regex)) a string of characters consisting of upper
	// and lowercase alphanumeric characters with no spaces. You can also include any
	// of the following characters: _+=,.@-
	UserName *string
	// The unique identifier of the service-specific credential. This parameter allows
	// (through its regex pattern (http://wikipedia.org/wiki/regex)) a string of
	// characters that can consist of any upper or lowercased letter or digit.
	ServiceSpecificCredentialId *string
}

type UpdateServiceSpecificCredentialOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpUpdateServiceSpecificCredentialMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpUpdateServiceSpecificCredential{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpUpdateServiceSpecificCredential{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateServiceSpecificCredential(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "UpdateServiceSpecificCredential",
	}
}