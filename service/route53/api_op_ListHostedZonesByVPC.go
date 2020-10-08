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

// Lists all the private hosted zones that a specified VPC is associated with,
// regardless of which AWS account or AWS service owns the hosted zones. The
// HostedZoneOwner structure in the response contains one of the following
// values:
//
//     * An OwningAccount element, which contains the account number of
// either the current AWS account or another AWS account. Some services, such as
// AWS Cloud Map, create hosted zones using the current account.
//
//     * An
// OwningService element, which identifies the AWS service that created and owns
// the hosted zone. For example, if a hosted zone was created by Amazon Elastic
// File System (Amazon EFS), the value of Owner is efs.amazonaws.com.
func (c *Client) ListHostedZonesByVPC(ctx context.Context, params *ListHostedZonesByVPCInput, optFns ...func(*Options)) (*ListHostedZonesByVPCOutput, error) {
	stack := middleware.NewStack("ListHostedZonesByVPC", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpListHostedZonesByVPCMiddlewares(stack)
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
	addOpListHostedZonesByVPCValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListHostedZonesByVPC(options.Region), middleware.Before)
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
			OperationName: "ListHostedZonesByVPC",
			Err:           err,
		}
	}
	out := result.(*ListHostedZonesByVPCOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Lists all the private hosted zones that a specified VPC is associated with,
// regardless of which AWS account created the hosted zones.
type ListHostedZonesByVPCInput struct {

	// The ID of the Amazon VPC that you want to list hosted zones for.
	//
	// This member is required.
	VPCId *string

	// For the Amazon VPC that you specified for VPCId, the AWS Region that you created
	// the VPC in.
	//
	// This member is required.
	VPCRegion types.VPCRegion

	// (Optional) The maximum number of hosted zones that you want Amazon Route 53 to
	// return. If the specified VPC is associated with more than MaxItems hosted zones,
	// the response includes a NextToken element. NextToken contains the hosted zone ID
	// of the first hosted zone that Route 53 will return if you submit another
	// request.
	MaxItems *string

	// If the previous response included a NextToken element, the specified VPC is
	// associated with more hosted zones. To get more hosted zones, submit another
	// ListHostedZonesByVPC request. For the value of NextToken, specify the value of
	// NextToken from the previous response. If the previous response didn't include a
	// NextToken element, there are no more hosted zones to get.
	NextToken *string
}

type ListHostedZonesByVPCOutput struct {

	// A list that contains one HostedZoneSummary element for each hosted zone that the
	// specified Amazon VPC is associated with. Each HostedZoneSummary element contains
	// the hosted zone name and ID, and information about who owns the hosted zone.
	//
	// This member is required.
	HostedZoneSummaries []*types.HostedZoneSummary

	// The value that you specified for MaxItems in the most recent
	// ListHostedZonesByVPC request.
	//
	// This member is required.
	MaxItems *string

	// The value that you specified for NextToken in the most recent
	// ListHostedZonesByVPC request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpListHostedZonesByVPCMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpListHostedZonesByVPC{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpListHostedZonesByVPC{}, middleware.After)
}

func newServiceMetadataMiddleware_opListHostedZonesByVPC(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "ListHostedZonesByVPC",
	}
}
