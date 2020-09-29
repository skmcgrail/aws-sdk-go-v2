// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes one or more instances from the specified Auto Scaling group. After the
// instances are detached, you can manage them independent of the Auto Scaling
// group. If you do not specify the option to decrement the desired capacity,
// Amazon EC2 Auto Scaling launches instances to replace the ones that are
// detached. If there is a Classic Load Balancer attached to the Auto Scaling
// group, the instances are deregistered from the load balancer. If there are
// target groups attached to the Auto Scaling group, the instances are deregistered
// from the target groups. For more information, see Detach EC2 Instances from Your
// Auto Scaling Group
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/detach-instance-asg.html)
// in the Amazon EC2 Auto Scaling User Guide.
func (c *Client) DetachInstances(ctx context.Context, params *DetachInstancesInput, optFns ...func(*Options)) (*DetachInstancesOutput, error) {
	stack := middleware.NewStack("DetachInstances", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDetachInstancesMiddlewares(stack)
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
	addOpDetachInstancesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDetachInstances(options.Region), middleware.Before)
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
			OperationName: "DetachInstances",
			Err:           err,
		}
	}
	out := result.(*DetachInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DetachInstancesInput struct {
	// The IDs of the instances. You can specify up to 20 instances.
	InstanceIds []*string
	// The name of the Auto Scaling group.
	AutoScalingGroupName *string
	// Indicates whether the Auto Scaling group decrements the desired capacity value
	// by the number of instances detached.
	ShouldDecrementDesiredCapacity *bool
}

type DetachInstancesOutput struct {
	// The activities related to detaching the instances from the Auto Scaling group.
	Activities []*types.Activity

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDetachInstancesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDetachInstances{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDetachInstances{}, middleware.After)
}

func newServiceMetadataMiddleware_opDetachInstances(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "DetachInstances",
	}
}