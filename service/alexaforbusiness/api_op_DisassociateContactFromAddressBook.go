// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disassociates a contact from a given address book.
func (c *Client) DisassociateContactFromAddressBook(ctx context.Context, params *DisassociateContactFromAddressBookInput, optFns ...func(*Options)) (*DisassociateContactFromAddressBookOutput, error) {
	stack := middleware.NewStack("DisassociateContactFromAddressBook", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDisassociateContactFromAddressBookMiddlewares(stack)
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
	addOpDisassociateContactFromAddressBookValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateContactFromAddressBook(options.Region), middleware.Before)
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
			OperationName: "DisassociateContactFromAddressBook",
			Err:           err,
		}
	}
	out := result.(*DisassociateContactFromAddressBookOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateContactFromAddressBookInput struct {

	// The ARN of the address from which to disassociate the contact.
	//
	// This member is required.
	AddressBookArn *string

	// The ARN of the contact to disassociate from an address book.
	//
	// This member is required.
	ContactArn *string
}

type DisassociateContactFromAddressBookOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDisassociateContactFromAddressBookMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDisassociateContactFromAddressBook{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisassociateContactFromAddressBook{}, middleware.After)
}

func newServiceMetadataMiddleware_opDisassociateContactFromAddressBook(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "DisassociateContactFromAddressBook",
	}
}
