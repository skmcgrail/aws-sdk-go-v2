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
	"time"
)

// Creates a Spot Instance request. For more information, see Spot Instance
// requests
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-requests.html) in the
// Amazon EC2 User Guide for Linux Instances.
func (c *Client) RequestSpotInstances(ctx context.Context, params *RequestSpotInstancesInput, optFns ...func(*Options)) (*RequestSpotInstancesOutput, error) {
	stack := middleware.NewStack("RequestSpotInstances", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpRequestSpotInstancesMiddlewares(stack)
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
	addOpRequestSpotInstancesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRequestSpotInstances(options.Region), middleware.Before)
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
			OperationName: "RequestSpotInstances",
			Err:           err,
		}
	}
	out := result.(*RequestSpotInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for RequestSpotInstances.
type RequestSpotInstancesInput struct {

	// The user-specified name for a logical grouping of requests. When you specify an
	// Availability Zone group in a Spot Instance request, all Spot Instances in the
	// request are launched in the same Availability Zone. Instance proximity is
	// maintained with this parameter, but the choice of Availability Zone is not. The
	// group applies only to requests for Spot Instances of the same instance type. Any
	// additional Spot Instance requests that are specified with the same Availability
	// Zone group name are launched in that same Availability Zone, as long as at least
	// one instance from the group is still active. If there is no active instance
	// running in the Availability Zone group that you specify for a new Spot Instance
	// request (all instances are terminated, the request is expired, or the maximum
	// price you specified falls below current Spot price), then Amazon EC2 launches
	// the instance in any Availability Zone where the constraint can be met.
	// Consequently, the subsequent set of Spot Instances could be placed in a
	// different zone from the original request, even if you specified the same
	// Availability Zone group. Default: Instances are launched in any available
	// Availability Zone.
	AvailabilityZoneGroup *string

	// The required duration for the Spot Instances (also known as Spot blocks), in
	// minutes. This value must be a multiple of 60 (60, 120, 180, 240, 300, or 360).
	// The duration period starts as soon as your Spot Instance receives its instance
	// ID. At the end of the duration period, Amazon EC2 marks the Spot Instance for
	// termination and provides a Spot Instance termination notice, which gives the
	// instance a two-minute warning before it terminates. You can't specify an
	// Availability Zone group or a launch group if you specify a duration.
	BlockDurationMinutes *int32

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. For more information, see How to Ensure Idempotency
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Run_Instance_Idempotency.html)
	// in the Amazon EC2 User Guide for Linux Instances.
	ClientToken *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The maximum number of Spot Instances to launch. Default: 1
	InstanceCount *int32

	// The behavior when a Spot Instance is interrupted. The default is terminate.
	InstanceInterruptionBehavior types.InstanceInterruptionBehavior

	// The instance launch group. Launch groups are Spot Instances that launch together
	// and terminate together. Default: Instances are launched and terminated
	// individually
	LaunchGroup *string

	// The launch specification.
	LaunchSpecification *types.RequestSpotLaunchSpecification

	// The maximum price per hour that you are willing to pay for a Spot Instance. The
	// default is the On-Demand price.
	SpotPrice *string

	// The key-value pair for tagging the Spot Instance request on creation. The value
	// for ResourceType must be spot-instances-request, otherwise the Spot Instance
	// request fails. To tag the Spot Instance request after it has been created, see
	// CreateTags
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateTags.html).
	TagSpecifications []*types.TagSpecification

	// The Spot Instance request type. Default: one-time
	Type types.SpotInstanceType

	// The start date of the request. If this is a one-time request, the request
	// becomes active at this date and time and remains active until all instances
	// launch, the request expires, or the request is canceled. If the request is
	// persistent, the request becomes active at this date and time and remains active
	// until it expires or is canceled. The specified start date and time cannot be
	// equal to the current date and time. You must specify a start date and time that
	// occurs after the current date and time.
	ValidFrom *time.Time

	// The end date of the request. If this is a one-time request, the request remains
	// active until all instances launch, the request is canceled, or this date is
	// reached. If the request is persistent, it remains active until it is canceled or
	// this date is reached. The default end date is 7 days from the current date.
	ValidUntil *time.Time
}

// Contains the output of RequestSpotInstances.
type RequestSpotInstancesOutput struct {

	// One or more Spot Instance requests.
	SpotInstanceRequests []*types.SpotInstanceRequest

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpRequestSpotInstancesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpRequestSpotInstances{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpRequestSpotInstances{}, middleware.After)
}

func newServiceMetadataMiddleware_opRequestSpotInstances(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "RequestSpotInstances",
	}
}
