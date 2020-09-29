// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Searches for routes in the specified transit gateway route table.
func (c *Client) SearchTransitGatewayRoutes(ctx context.Context, params *SearchTransitGatewayRoutesInput, optFns ...func(*Options)) (*SearchTransitGatewayRoutesOutput, error) {
	stack := middleware.NewStack("SearchTransitGatewayRoutes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpSearchTransitGatewayRoutesMiddlewares(stack)
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
	addOpSearchTransitGatewayRoutesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSearchTransitGatewayRoutes(options.Region), middleware.Before)
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
			OperationName: "SearchTransitGatewayRoutes",
			Err:           err,
		}
	}
	out := result.(*SearchTransitGatewayRoutesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchTransitGatewayRoutesInput struct {
	// One or more filters. The possible values are:
	//
	//     *
	// attachment.transit-gateway-attachment-id- The id of the transit gateway
	// attachment.
	//
	//     * attachment.resource-id - The resource id of the transit
	// gateway attachment.
	//
	//     * attachment.resource-type - The attachment resource
	// type (vpc | vpn).
	//
	//     * route-search.exact-match - The exact match of the
	// specified filter.
	//
	//     * route-search.longest-prefix-match - The longest prefix
	// that matches the route.
	//
	//     * route-search.subnet-of-match - The routes with a
	// subnet that match the specified CIDR filter.
	//
	//     *
	// route-search.supernet-of-match - The routes with a CIDR that encompass the CIDR
	// filter. For example, if you have 10.0.1.0/29 and 10.0.1.0/31 routes in your
	// route table and you specify supernet-of-match as 10.0.1.0/30, then the result
	// returns 10.0.1.0/29.
	//
	//     * state - The state of the route (active |
	// blackhole).
	//
	//     * type - The type of route (propagated | static).
	Filters []*types.Filter
	// The ID of the transit gateway route table.
	TransitGatewayRouteTableId *string
	// The maximum number of routes to return.
	MaxResults *int32
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type SearchTransitGatewayRoutesOutput struct {
	// Information about the routes.
	Routes []*types.TransitGatewayRoute
	// Indicates whether there are additional routes available.
	AdditionalRoutesAvailable *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpSearchTransitGatewayRoutesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpSearchTransitGatewayRoutes{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpSearchTransitGatewayRoutes{}, middleware.After)
}

func newServiceMetadataMiddleware_opSearchTransitGatewayRoutes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "SearchTransitGatewayRoutes",
	}
}