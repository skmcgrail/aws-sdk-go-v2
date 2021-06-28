// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates capacity settings for a fleet. For fleets with multiple locations, use
// this operation to manage capacity settings in each location individually. Fleet
// capacity determines the number of game sessions and players that can be hosted
// based on the fleet configuration. Use this operation to set the following fleet
// capacity properties:
//
// * Minimum/maximum size: Set hard limits on fleet capacity.
// GameLift cannot set the fleet's capacity to a value outside of this range,
// whether the capacity is changed manually or through automatic scaling.
//
// *
// Desired capacity: Manually set the number of EC2 instances to be maintained in a
// fleet location. Before changing a fleet's desired capacity, you may want to call
// DescribeEC2InstanceLimits to get the maximum capacity of the fleet's EC2
// instance type. Alternatively, consider using automatic scaling to adjust
// capacity based on player demand.
//
// This operation can be used in the following
// ways:
//
// * To update capacity for a fleet's home Region, or if the fleet has no
// remote locations, omit the Location parameter. The fleet must be in ACTIVE
// status.
//
// * To update capacity for a fleet's remote location, include the
// Location parameter set to the location to be updated. The location must be in
// ACTIVE status.
//
// If successful, capacity settings are updated immediately. In
// response a change in desired capacity, GameLift initiates steps to start new
// instances or terminate existing instances in the requested fleet location. This
// continues until the location's active instance count matches the new desired
// instance count. You can track a fleet's current capacity by calling
// DescribeFleetCapacity or DescribeFleetLocationCapacity. If the requested desired
// instance count is higher than the instance type's limit, the LimitExceeded
// exception occurs. Learn more Scaling fleet capacity
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-manage-capacity.html)
// Related actions CreateFleetLocations | UpdateFleetAttributes |
// UpdateFleetCapacity | UpdateFleetPortSettings | UpdateRuntimeConfiguration |
// StopFleetActions | StartFleetActions | PutScalingPolicy | DeleteFleet |
// DeleteFleetLocations | DeleteScalingPolicy | All APIs by task
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html#reference-awssdk-resources-fleets)
func (c *Client) UpdateFleetCapacity(ctx context.Context, params *UpdateFleetCapacityInput, optFns ...func(*Options)) (*UpdateFleetCapacityOutput, error) {
	if params == nil {
		params = &UpdateFleetCapacityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateFleetCapacity", params, optFns, c.addOperationUpdateFleetCapacityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateFleetCapacityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request operation.
type UpdateFleetCapacityInput struct {

	// A unique identifier for the fleet to update capacity settings for. You can use
	// either the fleet ID or ARN value.
	//
	// This member is required.
	FleetId *string

	// The number of EC2 instances you want to maintain in the specified fleet
	// location. This value must fall between the minimum and maximum size limits.
	DesiredInstances *int32

	// The name of a remote location to update fleet capacity settings for, in the form
	// of an AWS Region code such as us-west-2.
	Location *string

	// The maximum number of instances that are allowed in the specified fleet
	// location. If this parameter is not set, the default is 1.
	MaxSize *int32

	// The minimum number of instances that are allowed in the specified fleet
	// location. If this parameter is not set, the default is 0.
	MinSize *int32

	noSmithyDocumentSerde
}

// Represents the returned data in response to a request operation.
type UpdateFleetCapacityOutput struct {

	// The Amazon Resource Name (ARN
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-arn-format.html)) that is
	// assigned to a GameLift fleet resource and uniquely identifies it. ARNs are
	// unique across all Regions. Format is
	// arn:aws:gamelift:::fleet/fleet-a1234567-b8c9-0d1e-2fa3-b45c6d7e8912.
	FleetArn *string

	// A unique identifier for the fleet that was updated.
	FleetId *string

	// The remote location being updated, expressed as an AWS Region code, such as
	// us-west-2.
	Location *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateFleetCapacityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateFleetCapacity{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateFleetCapacity{}, middleware.After)
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
	if err = addOpUpdateFleetCapacityValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateFleetCapacity(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateFleetCapacity(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "UpdateFleetCapacity",
	}
}
