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

// Modifies a Traffic Mirror session.
func (c *Client) ModifyTrafficMirrorSession(ctx context.Context, params *ModifyTrafficMirrorSessionInput, optFns ...func(*Options)) (*ModifyTrafficMirrorSessionOutput, error) {
	stack := middleware.NewStack("ModifyTrafficMirrorSession", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpModifyTrafficMirrorSessionMiddlewares(stack)
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
	addOpModifyTrafficMirrorSessionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opModifyTrafficMirrorSession(options.Region), middleware.Before)
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
			OperationName: "ModifyTrafficMirrorSession",
			Err:           err,
		}
	}
	out := result.(*ModifyTrafficMirrorSessionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyTrafficMirrorSessionInput struct {
	// The description to assign to the Traffic Mirror session.
	Description *string
	// The properties that you want to remove from the Traffic Mirror session. When you
	// remove a property from a Traffic Mirror session, the property is set to the
	// default.
	RemoveFields []types.TrafficMirrorSessionField
	// The Traffic Mirror target. The target must be in the same VPC as the source, or
	// have a VPC peering connection with the source.
	TrafficMirrorTargetId *string
	// The ID of the Traffic Mirror filter.
	TrafficMirrorFilterId *string
	// The ID of the Traffic Mirror session.
	TrafficMirrorSessionId *string
	// The number of bytes in each packet to mirror. These are bytes after the VXLAN
	// header. To mirror a subset, set this to the length (in bytes) to mirror. For
	// example, if you set this value to 100, then the first 100 bytes that meet the
	// filter criteria are copied to the target. Do not specify this parameter when you
	// want to mirror the entire packet.
	PacketLength *int32
	// The session number determines the order in which sessions are evaluated when an
	// interface is used by multiple sessions. The first session with a matching filter
	// is the one that mirrors the packets. Valid values are 1-32766.
	SessionNumber *int32
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
	// The virtual network ID of the Traffic Mirror session.
	VirtualNetworkId *int32
}

type ModifyTrafficMirrorSessionOutput struct {
	// Information about the Traffic Mirror session.
	TrafficMirrorSession *types.TrafficMirrorSession

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpModifyTrafficMirrorSessionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpModifyTrafficMirrorSession{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpModifyTrafficMirrorSession{}, middleware.After)
}

func newServiceMetadataMiddleware_opModifyTrafficMirrorSession(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ModifyTrafficMirrorSession",
	}
}