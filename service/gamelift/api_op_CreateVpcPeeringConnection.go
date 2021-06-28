// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Establishes a VPC peering connection between a virtual private cloud (VPC) in an
// AWS account with the VPC for your Amazon GameLift fleet. VPC peering enables the
// game servers on your fleet to communicate directly with other AWS resources. You
// can peer with VPCs in any AWS account that you have access to, including the
// account that you use to manage your Amazon GameLift fleets. You cannot peer with
// VPCs that are in different Regions. For more information, see VPC Peering with
// Amazon GameLift Fleets
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/vpc-peering.html).
// Before calling this operation to establish the peering connection, you first
// need to call CreateVpcPeeringAuthorization and identify the VPC you want to peer
// with. Once the authorization for the specified VPC is issued, you have 24 hours
// to establish the connection. These two operations handle all tasks necessary to
// peer the two VPCs, including acceptance, updating routing tables, etc. To
// establish the connection, call this operation from the AWS account that is used
// to manage the Amazon GameLift fleets. Identify the following values: (1) The ID
// of the fleet you want to be enable a VPC peering connection for; (2) The AWS
// account with the VPC that you want to peer with; and (3) The ID of the VPC you
// want to peer with. This operation is asynchronous. If successful, a
// VpcPeeringConnection request is created. You can use continuous polling to track
// the request's status using DescribeVpcPeeringConnections, or by monitoring fleet
// events for success or failure using DescribeFleetEvents. Related actions
// CreateVpcPeeringAuthorization | DescribeVpcPeeringAuthorizations |
// DeleteVpcPeeringAuthorization | CreateVpcPeeringConnection |
// DescribeVpcPeeringConnections | DeleteVpcPeeringConnection | All APIs by task
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html#reference-awssdk-resources-fleets)
func (c *Client) CreateVpcPeeringConnection(ctx context.Context, params *CreateVpcPeeringConnectionInput, optFns ...func(*Options)) (*CreateVpcPeeringConnectionOutput, error) {
	if params == nil {
		params = &CreateVpcPeeringConnectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateVpcPeeringConnection", params, optFns, c.addOperationCreateVpcPeeringConnectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateVpcPeeringConnectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request operation.
type CreateVpcPeeringConnectionInput struct {

	// A unique identifier for the fleet. You can use either the fleet ID or ARN value.
	// This tells Amazon GameLift which GameLift VPC to peer with.
	//
	// This member is required.
	FleetId *string

	// A unique identifier for the AWS account with the VPC that you want to peer your
	// Amazon GameLift fleet with. You can find your Account ID in the AWS Management
	// Console under account settings.
	//
	// This member is required.
	PeerVpcAwsAccountId *string

	// A unique identifier for a VPC with resources to be accessed by your GameLift
	// fleet. The VPC must be in the same Region as your fleet. To look up a VPC ID,
	// use the VPC Dashboard (https://console.aws.amazon.com/vpc/) in the AWS
	// Management Console. Learn more about VPC peering in VPC Peering with GameLift
	// Fleets
	// (https://docs.aws.amazon.com/gamelift/latest/developerguide/vpc-peering.html).
	//
	// This member is required.
	PeerVpcId *string

	noSmithyDocumentSerde
}

type CreateVpcPeeringConnectionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateVpcPeeringConnectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateVpcPeeringConnection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateVpcPeeringConnection{}, middleware.After)
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
	if err = addOpCreateVpcPeeringConnectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateVpcPeeringConnection(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateVpcPeeringConnection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "CreateVpcPeeringConnection",
	}
}
