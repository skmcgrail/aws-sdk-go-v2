// Code generated by smithy-go-codegen DO NOT EDIT.

package mediaconnect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Revokes an entitlement from a flow. Once an entitlement is revoked, the content
// becomes unavailable to the subscriber and the associated output is removed.
func (c *Client) RevokeFlowEntitlement(ctx context.Context, params *RevokeFlowEntitlementInput, optFns ...func(*Options)) (*RevokeFlowEntitlementOutput, error) {
	if params == nil {
		params = &RevokeFlowEntitlementInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RevokeFlowEntitlement", params, optFns, c.addOperationRevokeFlowEntitlementMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RevokeFlowEntitlementOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RevokeFlowEntitlementInput struct {

	// The ARN of the entitlement that you want to revoke.
	//
	// This member is required.
	EntitlementArn *string

	// The flow that you want to revoke an entitlement from.
	//
	// This member is required.
	FlowArn *string

	noSmithyDocumentSerde
}

type RevokeFlowEntitlementOutput struct {

	// The ARN of the entitlement that was revoked.
	EntitlementArn *string

	// The ARN of the flow that the entitlement was revoked from.
	FlowArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRevokeFlowEntitlementMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpRevokeFlowEntitlement{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpRevokeFlowEntitlement{}, middleware.After)
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
	if err = addOpRevokeFlowEntitlementValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRevokeFlowEntitlement(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRevokeFlowEntitlement(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediaconnect",
		OperationName: "RevokeFlowEntitlement",
	}
}
