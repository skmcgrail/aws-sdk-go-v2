// Code generated by smithy-go-codegen DO NOT EDIT.

package apigatewayv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets a VPC link.
func (c *Client) GetVpcLink(ctx context.Context, params *GetVpcLinkInput, optFns ...func(*Options)) (*GetVpcLinkOutput, error) {
	if params == nil {
		params = &GetVpcLinkInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetVpcLink", params, optFns, c.addOperationGetVpcLinkMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetVpcLinkOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetVpcLinkInput struct {

	// The ID of the VPC link.
	//
	// This member is required.
	VpcLinkId *string

	noSmithyDocumentSerde
}

type GetVpcLinkOutput struct {

	// The timestamp when the VPC link was created.
	CreatedDate *time.Time

	// The name of the VPC link.
	Name *string

	// A list of security group IDs for the VPC link.
	SecurityGroupIds []string

	// A list of subnet IDs to include in the VPC link.
	SubnetIds []string

	// Tags for the VPC link.
	Tags map[string]string

	// The ID of the VPC link.
	VpcLinkId *string

	// The status of the VPC link.
	VpcLinkStatus types.VpcLinkStatus

	// A message summarizing the cause of the status of the VPC link.
	VpcLinkStatusMessage *string

	// The version of the VPC link.
	VpcLinkVersion types.VpcLinkVersion

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetVpcLinkMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetVpcLink{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetVpcLink{}, middleware.After)
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
	if err = addOpGetVpcLinkValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetVpcLink(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetVpcLink(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "GetVpcLink",
	}
}
