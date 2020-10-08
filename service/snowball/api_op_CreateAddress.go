// Code generated by smithy-go-codegen DO NOT EDIT.

package snowball

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/snowball/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an address for a Snowball to be shipped to. In most regions, addresses
// are validated at the time of creation. The address you provide must be located
// within the serviceable area of your region. If the address is invalid or
// unsupported, then an exception is thrown.
func (c *Client) CreateAddress(ctx context.Context, params *CreateAddressInput, optFns ...func(*Options)) (*CreateAddressOutput, error) {
	stack := middleware.NewStack("CreateAddress", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateAddressMiddlewares(stack)
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
	addOpCreateAddressValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAddress(options.Region), middleware.Before)
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
			OperationName: "CreateAddress",
			Err:           err,
		}
	}
	out := result.(*CreateAddressOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAddressInput struct {

	// The address that you want the Snowball shipped to.
	//
	// This member is required.
	Address *types.Address
}

type CreateAddressOutput struct {

	// The automatically generated ID for a specific address. You'll use this ID when
	// you create a job to specify which address you want the Snowball for that job
	// shipped to.
	AddressId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateAddressMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateAddress{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateAddress{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateAddress(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "snowball",
		OperationName: "CreateAddress",
	}
}
