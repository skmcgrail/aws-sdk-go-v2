// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Associates an Amazon VPC with a private hosted zone.  <note> <p>To perform the
// association, the VPC and the private hosted zone must already exist. Also, you
// can't convert a public hosted zone into a private hosted zone.</p> </note> <p>If
// you want to associate a VPC that was created by one AWS account with a private
// hosted zone that was created by a different account, do one of the
// following:</p> <ul> <li> <p>Use the AWS account that created the private hosted
// zone to submit a <a
// href="https://docs.aws.amazon.com/Route53/latest/APIReference/API_CreateVPCAssociationAuthorization.html">CreateVPCAssociationAuthorization</a>
// request. Then use the account that created the VPC to submit an
// <code>AssociateVPCWithHostedZone</code> request.</p> </li> <li> <p>If a subnet
// in the VPC was shared with another account, you can use the account that the
// subnet was shared with to submit an <code>AssociateVPCWithHostedZone</code>
// request. For more information about sharing subnets, see <a
// href="https://docs.aws.amazon.com/vpc/latest/userguide/vpc-sharing.html">Working
// with Shared VPCs</a>.</p> </li> </ul>
func (c *Client) AssociateVPCWithHostedZone(ctx context.Context, params *AssociateVPCWithHostedZoneInput, optFns ...func(*Options)) (*AssociateVPCWithHostedZoneOutput, error) {
	stack := middleware.NewStack("AssociateVPCWithHostedZone", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpAssociateVPCWithHostedZoneMiddlewares(stack)
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
	addOpAssociateVPCWithHostedZoneValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateVPCWithHostedZone(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addSanitizeURLMiddleware(stack)

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
			OperationName: "AssociateVPCWithHostedZone",
			Err:           err,
		}
	}
	out := result.(*AssociateVPCWithHostedZoneOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A complex type that contains information about the request to associate a VPC
// with a private hosted zone.
type AssociateVPCWithHostedZoneInput struct {

	// The ID of the private hosted zone that you want to associate an Amazon VPC with.
	// Note that you can't associate a VPC with a hosted zone that doesn't have an
	// existing VPC association.
	//
	// This member is required.
	HostedZoneId *string

	// A complex type that contains information about the VPC that you want to
	// associate with a private hosted zone.
	//
	// This member is required.
	VPC *types.VPC

	// Optional: A comment about the association request.
	Comment *string
}

// A complex type that contains the response information for the
// AssociateVPCWithHostedZone request.
type AssociateVPCWithHostedZoneOutput struct {

	// A complex type that describes the changes made to your hosted zone.
	//
	// This member is required.
	ChangeInfo *types.ChangeInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpAssociateVPCWithHostedZoneMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpAssociateVPCWithHostedZone{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpAssociateVPCWithHostedZone{}, middleware.After)
}

func newServiceMetadataMiddleware_opAssociateVPCWithHostedZone(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "AssociateVPCWithHostedZone",
	}
}
