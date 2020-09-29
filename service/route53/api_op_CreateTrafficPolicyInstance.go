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

// Creates resource record sets in a specified hosted zone based on the settings in
// a specified traffic policy version. In addition, CreateTrafficPolicyInstance
// associates the resource record sets with a specified domain name (such as
// example.com) or subdomain name (such as www.example.com). Amazon Route 53
// responds to DNS queries for the domain or subdomain name by using the resource
// record sets that CreateTrafficPolicyInstance created.
func (c *Client) CreateTrafficPolicyInstance(ctx context.Context, params *CreateTrafficPolicyInstanceInput, optFns ...func(*Options)) (*CreateTrafficPolicyInstanceOutput, error) {
	stack := middleware.NewStack("CreateTrafficPolicyInstance", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpCreateTrafficPolicyInstanceMiddlewares(stack)
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
	addOpCreateTrafficPolicyInstanceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTrafficPolicyInstance(options.Region), middleware.Before)
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
			OperationName: "CreateTrafficPolicyInstance",
			Err:           err,
		}
	}
	out := result.(*CreateTrafficPolicyInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A complex type that contains information about the resource record sets that you
// want to create based on a specified traffic policy.
type CreateTrafficPolicyInstanceInput struct {
	// The domain name (such as example.com) or subdomain name (such as
	// www.example.com) for which Amazon Route 53 responds to DNS queries by using the
	// resource record sets that Route 53 creates for this traffic policy instance.
	Name *string
	// (Optional) The TTL that you want Amazon Route 53 to assign to all of the
	// resource record sets that it creates in the specified hosted zone.
	TTL *int64
	// The ID of the traffic policy that you want to use to create resource record sets
	// in the specified hosted zone.
	TrafficPolicyId *string
	// The version of the traffic policy that you want to use to create resource record
	// sets in the specified hosted zone.
	TrafficPolicyVersion *int32
	// The ID of the hosted zone that you want Amazon Route 53 to create resource
	// record sets in by using the configuration in a traffic policy.
	HostedZoneId *string
}

// A complex type that contains the response information for the
// CreateTrafficPolicyInstance request.
type CreateTrafficPolicyInstanceOutput struct {
	// A complex type that contains settings for the new traffic policy instance.
	TrafficPolicyInstance *types.TrafficPolicyInstance
	// A unique URL that represents a new traffic policy instance.
	Location *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpCreateTrafficPolicyInstanceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpCreateTrafficPolicyInstance{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpCreateTrafficPolicyInstance{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateTrafficPolicyInstance(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "CreateTrafficPolicyInstance",
	}
}