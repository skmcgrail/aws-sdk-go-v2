// Code generated by smithy-go-codegen DO NOT EDIT.

package route53domains

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This operation disables automatic renewal of domain registration for the
// specified domain.
func (c *Client) DisableDomainAutoRenew(ctx context.Context, params *DisableDomainAutoRenewInput, optFns ...func(*Options)) (*DisableDomainAutoRenewOutput, error) {
	stack := middleware.NewStack("DisableDomainAutoRenew", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDisableDomainAutoRenewMiddlewares(stack)
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
	addOpDisableDomainAutoRenewValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisableDomainAutoRenew(options.Region), middleware.Before)
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
			OperationName: "DisableDomainAutoRenew",
			Err:           err,
		}
	}
	out := result.(*DisableDomainAutoRenewOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisableDomainAutoRenewInput struct {

	// The name of the domain that you want to disable automatic renewal for.
	//
	// This member is required.
	DomainName *string
}

type DisableDomainAutoRenewOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDisableDomainAutoRenewMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDisableDomainAutoRenew{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisableDomainAutoRenew{}, middleware.After)
}

func newServiceMetadataMiddleware_opDisableDomainAutoRenew(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53domains",
		OperationName: "DisableDomainAutoRenew",
	}
}
