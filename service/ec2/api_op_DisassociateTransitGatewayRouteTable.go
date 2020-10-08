// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disassociates a resource attachment from a transit gateway route table.
func (c *Client) DisassociateTransitGatewayRouteTable(ctx context.Context, params *DisassociateTransitGatewayRouteTableInput, optFns ...func(*Options)) (*DisassociateTransitGatewayRouteTableOutput, error) {
	stack := middleware.NewStack("DisassociateTransitGatewayRouteTable", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDisassociateTransitGatewayRouteTableMiddlewares(stack)
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
	addOpDisassociateTransitGatewayRouteTableValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateTransitGatewayRouteTable(options.Region), middleware.Before)
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
			OperationName: "DisassociateTransitGatewayRouteTable",
			Err:           err,
		}
	}
	out := result.(*DisassociateTransitGatewayRouteTableOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateTransitGatewayRouteTableInput struct {

	// The ID of the attachment.
	//
	// This member is required.
	TransitGatewayAttachmentId *string

	// The ID of the transit gateway route table.
	//
	// This member is required.
	TransitGatewayRouteTableId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type DisassociateTransitGatewayRouteTableOutput struct {

	// Information about the association.
	Association *types.TransitGatewayAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDisassociateTransitGatewayRouteTableMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDisassociateTransitGatewayRouteTable{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDisassociateTransitGatewayRouteTable{}, middleware.After)
}

func newServiceMetadataMiddleware_opDisassociateTransitGatewayRouteTable(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DisassociateTransitGatewayRouteTable",
	}
}
