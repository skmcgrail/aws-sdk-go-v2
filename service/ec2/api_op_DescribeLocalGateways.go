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

// Describes one or more local gateways. By default, all local gateways are
// described. Alternatively, you can filter the results.
func (c *Client) DescribeLocalGateways(ctx context.Context, params *DescribeLocalGatewaysInput, optFns ...func(*Options)) (*DescribeLocalGatewaysOutput, error) {
	stack := middleware.NewStack("DescribeLocalGateways", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDescribeLocalGatewaysMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeLocalGateways(options.Region), middleware.Before)
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
			OperationName: "DescribeLocalGateways",
			Err:           err,
		}
	}
	out := result.(*DescribeLocalGatewaysOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeLocalGatewaysInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// One or more filters.
	Filters []*types.Filter

	// One or more filters.
	//
	//     * local-gateway-id - The ID of a local gateway.
	//
	//     *
	// local-gateway-route-table-id - The ID of the local gateway route table.
	//
	//     *
	// local-gateway-route-table-virtual-interface-group-association-id - The ID of the
	// association.
	//
	//     * local-gateway-route-table-virtual-interface-group-id - The
	// ID of the virtual interface group.
	//
	//     * outpost-arn - The Amazon Resource Name
	// (ARN) of the Outpost.
	//
	//     * state - The state of the association.
	LocalGatewayIds []*string

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string
}

type DescribeLocalGatewaysOutput struct {

	// Information about the local gateways.
	LocalGateways []*types.LocalGateway

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDescribeLocalGatewaysMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDescribeLocalGateways{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeLocalGateways{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeLocalGateways(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeLocalGateways",
	}
}
