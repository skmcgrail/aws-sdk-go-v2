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

// Removes authorization to submit an AssociateVPCWithHostedZone request to
// associate a specified VPC with a hosted zone that was created by a different
// account. You must use the account that created the hosted zone to submit a
// DeleteVPCAssociationAuthorization request. Sending this request only prevents
// the AWS account that created the VPC from associating the VPC with the Amazon
// Route 53 hosted zone in the future. If the VPC is already associated with the
// hosted zone, DeleteVPCAssociationAuthorization won't disassociate the VPC from
// the hosted zone. If you want to delete an existing association, use
// DisassociateVPCFromHostedZone.
func (c *Client) DeleteVPCAssociationAuthorization(ctx context.Context, params *DeleteVPCAssociationAuthorizationInput, optFns ...func(*Options)) (*DeleteVPCAssociationAuthorizationOutput, error) {
	stack := middleware.NewStack("DeleteVPCAssociationAuthorization", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpDeleteVPCAssociationAuthorizationMiddlewares(stack)
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
	addOpDeleteVPCAssociationAuthorizationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteVPCAssociationAuthorization(options.Region), middleware.Before)
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
			OperationName: "DeleteVPCAssociationAuthorization",
			Err:           err,
		}
	}
	out := result.(*DeleteVPCAssociationAuthorizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A complex type that contains information about the request to remove
// authorization to associate a VPC that was created by one AWS account with a
// hosted zone that was created with a different AWS account.
type DeleteVPCAssociationAuthorizationInput struct {

	// When removing authorization to associate a VPC that was created by one AWS
	// account with a hosted zone that was created with a different AWS account, the ID
	// of the hosted zone.
	//
	// This member is required.
	HostedZoneId *string

	// When removing authorization to associate a VPC that was created by one AWS
	// account with a hosted zone that was created with a different AWS account, a
	// complex type that includes the ID and region of the VPC.
	//
	// This member is required.
	VPC *types.VPC
}

// Empty response for the request.
type DeleteVPCAssociationAuthorizationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpDeleteVPCAssociationAuthorizationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpDeleteVPCAssociationAuthorization{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteVPCAssociationAuthorization{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteVPCAssociationAuthorization(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "DeleteVPCAssociationAuthorization",
	}
}
