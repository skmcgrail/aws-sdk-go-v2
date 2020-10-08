// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deprecated. Use the ListIdentities operation to list the email addresses and
// domains associated with your account.
func (c *Client) ListVerifiedEmailAddresses(ctx context.Context, params *ListVerifiedEmailAddressesInput, optFns ...func(*Options)) (*ListVerifiedEmailAddressesOutput, error) {
	stack := middleware.NewStack("ListVerifiedEmailAddresses", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpListVerifiedEmailAddressesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListVerifiedEmailAddresses(options.Region), middleware.Before)
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
			OperationName: "ListVerifiedEmailAddresses",
			Err:           err,
		}
	}
	out := result.(*ListVerifiedEmailAddressesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListVerifiedEmailAddressesInput struct {
}

// A list of email addresses that you have verified with Amazon SES under your AWS
// account.
type ListVerifiedEmailAddressesOutput struct {

	// A list of email addresses that have been verified.
	VerifiedEmailAddresses []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpListVerifiedEmailAddressesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpListVerifiedEmailAddresses{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpListVerifiedEmailAddresses{}, middleware.After)
}

func newServiceMetadataMiddleware_opListVerifiedEmailAddresses(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "ListVerifiedEmailAddresses",
	}
}
