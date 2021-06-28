// Code generated by smithy-go-codegen DO NOT EDIT.

package workmail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Mark a user, group, or resource as no longer used in Amazon WorkMail. This
// action disassociates the mailbox and schedules it for clean-up. WorkMail keeps
// mailboxes for 30 days before they are permanently removed. The functionality in
// the console is Disable.
func (c *Client) DeregisterFromWorkMail(ctx context.Context, params *DeregisterFromWorkMailInput, optFns ...func(*Options)) (*DeregisterFromWorkMailOutput, error) {
	if params == nil {
		params = &DeregisterFromWorkMailInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeregisterFromWorkMail", params, optFns, c.addOperationDeregisterFromWorkMailMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeregisterFromWorkMailOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeregisterFromWorkMailInput struct {

	// The identifier for the member (user or group) to be updated.
	//
	// This member is required.
	EntityId *string

	// The identifier for the organization under which the Amazon WorkMail entity
	// exists.
	//
	// This member is required.
	OrganizationId *string

	noSmithyDocumentSerde
}

type DeregisterFromWorkMailOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeregisterFromWorkMailMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeregisterFromWorkMail{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeregisterFromWorkMail{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDeregisterFromWorkMailValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeregisterFromWorkMail(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDeregisterFromWorkMail(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workmail",
		OperationName: "DeregisterFromWorkMail",
	}
}
