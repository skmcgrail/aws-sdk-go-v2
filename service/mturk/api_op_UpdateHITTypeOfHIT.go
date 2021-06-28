// Code generated by smithy-go-codegen DO NOT EDIT.

package mturk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// The UpdateHITTypeOfHIT operation allows you to change the HITType properties of
// a HIT. This operation disassociates the HIT from its old HITType properties and
// associates it with the new HITType properties. The HIT takes on the properties
// of the new HITType in place of the old ones.
func (c *Client) UpdateHITTypeOfHIT(ctx context.Context, params *UpdateHITTypeOfHITInput, optFns ...func(*Options)) (*UpdateHITTypeOfHITOutput, error) {
	if params == nil {
		params = &UpdateHITTypeOfHITInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateHITTypeOfHIT", params, optFns, c.addOperationUpdateHITTypeOfHITMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateHITTypeOfHITOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateHITTypeOfHITInput struct {

	// The HIT to update.
	//
	// This member is required.
	HITId *string

	// The ID of the new HIT type.
	//
	// This member is required.
	HITTypeId *string

	noSmithyDocumentSerde
}

type UpdateHITTypeOfHITOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateHITTypeOfHITMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateHITTypeOfHIT{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateHITTypeOfHIT{}, middleware.After)
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
	if err = addOpUpdateHITTypeOfHITValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateHITTypeOfHIT(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateHITTypeOfHIT(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mturk-requester",
		OperationName: "UpdateHITTypeOfHIT",
	}
}
