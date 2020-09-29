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

// Gets information about a specified hosted zone including the four name servers
// assigned to the hosted zone.
func (c *Client) GetHostedZone(ctx context.Context, params *GetHostedZoneInput, optFns ...func(*Options)) (*GetHostedZoneOutput, error) {
	stack := middleware.NewStack("GetHostedZone", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpGetHostedZoneMiddlewares(stack)
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
	addOpGetHostedZoneValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetHostedZone(options.Region), middleware.Before)
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
			OperationName: "GetHostedZone",
			Err:           err,
		}
	}
	out := result.(*GetHostedZoneOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to get information about a specified hosted zone.
type GetHostedZoneInput struct {
	// The ID of the hosted zone that you want to get information about.
	Id *string
}

// A complex type that contain the response to a GetHostedZone request.
type GetHostedZoneOutput struct {
	// A complex type that lists the Amazon Route 53 name servers for the specified
	// hosted zone.
	DelegationSet *types.DelegationSet
	// A complex type that contains general information about the specified hosted
	// zone.
	HostedZone *types.HostedZone
	// A complex type that contains information about the VPCs that are associated with
	// the specified hosted zone.
	VPCs []*types.VPC

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpGetHostedZoneMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpGetHostedZone{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpGetHostedZone{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetHostedZone(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "GetHostedZone",
	}
}