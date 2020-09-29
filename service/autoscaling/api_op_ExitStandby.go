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

// Moves the specified instances out of the standby state. After you put the
// instances back in service, the desired capacity is incremented. For more
// information, see Temporarily Removing Instances from Your Auto Scaling Group
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-enter-exit-standby.html)
// in the Amazon EC2 Auto Scaling User Guide.
func (c *Client) ExitStandby(ctx context.Context, params *ExitStandbyInput, optFns ...func(*Options)) (*ExitStandbyOutput, error) {
	stack := middleware.NewStack("ExitStandby", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpExitStandbyMiddlewares(stack)
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
	addOpExitStandbyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opExitStandby(options.Region), middleware.Before)
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
			OperationName: "ExitStandby",
			Err:           err,
		}
	}
	out := result.(*ExitStandbyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ExitStandbyInput struct {
	// The name of the Auto Scaling group.
	AutoScalingGroupName *string
	// The IDs of the instances. You can specify up to 20 instances.
	InstanceIds []*string
}

type ExitStandbyOutput struct {
	// The activities related to moving instances out of Standby mode.
	Activities []*types.Activity

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpExitStandbyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpExitStandby{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpExitStandby{}, middleware.After)
}

func newServiceMetadataMiddleware_opExitStandby(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "ExitStandby",
	}
}