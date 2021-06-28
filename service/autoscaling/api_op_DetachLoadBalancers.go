// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Detaches one or more Classic Load Balancers from the specified Auto Scaling
// group. This operation detaches only Classic Load Balancers. If you have
// Application Load Balancers, Network Load Balancers, or Gateway Load Balancers,
// use the DetachLoadBalancerTargetGroups API instead. When you detach a load
// balancer, it enters the Removing state while deregistering the instances in the
// group. When all instances are deregistered, then you can no longer describe the
// load balancer using the DescribeLoadBalancers API call. The instances remain
// running.
func (c *Client) DetachLoadBalancers(ctx context.Context, params *DetachLoadBalancersInput, optFns ...func(*Options)) (*DetachLoadBalancersOutput, error) {
	if params == nil {
		params = &DetachLoadBalancersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DetachLoadBalancers", params, optFns, c.addOperationDetachLoadBalancersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DetachLoadBalancersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DetachLoadBalancersInput struct {

	// The name of the Auto Scaling group.
	//
	// This member is required.
	AutoScalingGroupName *string

	// The names of the load balancers. You can specify up to 10 load balancers.
	//
	// This member is required.
	LoadBalancerNames []string

	noSmithyDocumentSerde
}

type DetachLoadBalancersOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDetachLoadBalancersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDetachLoadBalancers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDetachLoadBalancers{}, middleware.After)
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
	if err = addOpDetachLoadBalancersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDetachLoadBalancers(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDetachLoadBalancers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "DetachLoadBalancers",
	}
}
