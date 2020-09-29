// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about the traffic policy instances that you created by using a
// specify traffic policy version. After you submit a CreateTrafficPolicyInstance
// or an UpdateTrafficPolicyInstance request, there's a brief delay while Amazon
// Route 53 creates the resource record sets that are specified in the traffic
// policy definition. For more information, see the State response element. Route
// 53 returns a maximum of 100 items in each response. If you have a lot of traffic
// policy instances, you can use the MaxItems parameter to list them in groups of
// up to 100.
func (c *Client) ListTrafficPolicyInstancesByPolicy(ctx context.Context, params *ListTrafficPolicyInstancesByPolicyInput, optFns ...func(*Options)) (*ListTrafficPolicyInstancesByPolicyOutput, error) {
	stack := middleware.NewStack("ListTrafficPolicyInstancesByPolicy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpListTrafficPolicyInstancesByPolicyMiddlewares(stack)
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
	addOpListTrafficPolicyInstancesByPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListTrafficPolicyInstancesByPolicy(options.Region), middleware.Before)
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
			OperationName: "ListTrafficPolicyInstancesByPolicy",
			Err:           err,
		}
	}
	out := result.(*ListTrafficPolicyInstancesByPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A complex type that contains the information about the request to list your
// traffic policy instances.
type ListTrafficPolicyInstancesByPolicyInput struct {
	// If the value of IsTruncated in the previous response was true, you have more
	// traffic policy instances. To get more traffic policy instances, submit another
	// ListTrafficPolicyInstancesByPolicy request. For the value of
	// trafficpolicyinstancetype, specify the value of TrafficPolicyInstanceTypeMarker
	// from the previous response, which is the name of the first traffic policy
	// instance that Amazon Route 53 will return if you submit another request. If the
	// value of IsTruncated in the previous response was false, there are no more
	// traffic policy instances to get.
	TrafficPolicyInstanceTypeMarker types.RRType
	// If the value of IsTruncated in the previous response was true, you have more
	// traffic policy instances. To get more traffic policy instances, submit another
	// ListTrafficPolicyInstancesByPolicy request. For the value of hostedzoneid,
	// specify the value of HostedZoneIdMarker from the previous response, which is the
	// hosted zone ID of the first traffic policy instance that Amazon Route 53 will
	// return if you submit another request. If the value of IsTruncated in the
	// previous response was false, there are no more traffic policy instances to get.
	HostedZoneIdMarker *string
	// The ID of the traffic policy for which you want to list traffic policy
	// instances.
	TrafficPolicyId *string
	// If the value of IsTruncated in the previous response was true, you have more
	// traffic policy instances. To get more traffic policy instances, submit another
	// ListTrafficPolicyInstancesByPolicy request. For the value of
	// trafficpolicyinstancename, specify the value of TrafficPolicyInstanceNameMarker
	// from the previous response, which is the name of the first traffic policy
	// instance that Amazon Route 53 will return if you submit another request. If the
	// value of IsTruncated in the previous response was false, there are no more
	// traffic policy instances to get.
	TrafficPolicyInstanceNameMarker *string
	// The maximum number of traffic policy instances to be included in the response
	// body for this request. If you have more than MaxItems traffic policy instances,
	// the value of the IsTruncated element in the response is true, and the values of
	// HostedZoneIdMarker, TrafficPolicyInstanceNameMarker, and
	// TrafficPolicyInstanceTypeMarker represent the first traffic policy instance that
	// Amazon Route 53 will return if you submit another request.
	MaxItems *string
	// The version of the traffic policy for which you want to list traffic policy
	// instances. The version must be associated with the traffic policy that is
	// specified by TrafficPolicyId.
	TrafficPolicyVersion *int32
}

// A complex type that contains the response information for the request.
type ListTrafficPolicyInstancesByPolicyOutput struct {
	// A list that contains one TrafficPolicyInstance element for each traffic policy
	// instance that matches the elements in the request.
	TrafficPolicyInstances []*types.TrafficPolicyInstance
	// If IsTruncated is true, TrafficPolicyInstanceNameMarker is the name of the first
	// traffic policy instance in the next group of MaxItems traffic policy instances.
	TrafficPolicyInstanceNameMarker *string
	// A flag that indicates whether there are more traffic policy instances to be
	// listed. If the response was truncated, you can get the next group of traffic
	// policy instances by calling ListTrafficPolicyInstancesByPolicy again and
	// specifying the values of the HostedZoneIdMarker,
	// TrafficPolicyInstanceNameMarker, and TrafficPolicyInstanceTypeMarker elements in
	// the corresponding request parameters.
	IsTruncated *bool
	// The value that you specified for the MaxItems parameter in the call to
	// ListTrafficPolicyInstancesByPolicy that produced the current response.
	MaxItems *string
	// If IsTruncated is true, TrafficPolicyInstanceTypeMarker is the DNS type of the
	// resource record sets that are associated with the first traffic policy instance
	// in the next group of MaxItems traffic policy instances.
	TrafficPolicyInstanceTypeMarker types.RRType
	// If IsTruncated is true, HostedZoneIdMarker is the ID of the hosted zone of the
	// first traffic policy instance in the next group of traffic policy instances.
	HostedZoneIdMarker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpListTrafficPolicyInstancesByPolicyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpListTrafficPolicyInstancesByPolicy{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpListTrafficPolicyInstancesByPolicy{}, middleware.After)
}

func newServiceMetadataMiddleware_opListTrafficPolicyInstancesByPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "ListTrafficPolicyInstancesByPolicy",
	}
}