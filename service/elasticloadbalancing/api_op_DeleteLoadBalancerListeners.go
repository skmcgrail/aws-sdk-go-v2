// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancing

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specified listeners from the specified load balancer.
func (c *Client) DeleteLoadBalancerListeners(ctx context.Context, params *DeleteLoadBalancerListenersInput, optFns ...func(*Options)) (*DeleteLoadBalancerListenersOutput, error) {
	stack := middleware.NewStack("DeleteLoadBalancerListeners", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDeleteLoadBalancerListenersMiddlewares(stack)
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
	addOpDeleteLoadBalancerListenersValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteLoadBalancerListeners(options.Region), middleware.Before)
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
			OperationName: "DeleteLoadBalancerListeners",
			Err:           err,
		}
	}
	out := result.(*DeleteLoadBalancerListenersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DeleteLoadBalancerListeners.
type DeleteLoadBalancerListenersInput struct {

	// The name of the load balancer.
	//
	// This member is required.
	LoadBalancerName *string

	// The client port numbers of the listeners.
	//
	// This member is required.
	LoadBalancerPorts []*int32
}

// Contains the output of DeleteLoadBalancerListeners.
type DeleteLoadBalancerListenersOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDeleteLoadBalancerListenersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDeleteLoadBalancerListeners{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteLoadBalancerListeners{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteLoadBalancerListeners(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticloadbalancing",
		OperationName: "DeleteLoadBalancerListeners",
	}
}
