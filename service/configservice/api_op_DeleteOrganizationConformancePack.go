// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specified organization conformance pack and all of the config rules
// and remediation actions from all member accounts in that organization. Only a
// master account or a delegated administrator account can delete an organization
// conformance pack. When calling this API with a delegated administrator, you must
// ensure AWS Organizations ListDelegatedAdministrator permissions are added. AWS
// Config sets the state of a conformance pack to DELETE_IN_PROGRESS until the
// deletion is complete. You cannot update a conformance pack while it is in this
// state.
func (c *Client) DeleteOrganizationConformancePack(ctx context.Context, params *DeleteOrganizationConformancePackInput, optFns ...func(*Options)) (*DeleteOrganizationConformancePackOutput, error) {
	stack := middleware.NewStack("DeleteOrganizationConformancePack", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteOrganizationConformancePackMiddlewares(stack)
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
	addOpDeleteOrganizationConformancePackValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteOrganizationConformancePack(options.Region), middleware.Before)
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
			OperationName: "DeleteOrganizationConformancePack",
			Err:           err,
		}
	}
	out := result.(*DeleteOrganizationConformancePackOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteOrganizationConformancePackInput struct {

	// The name of organization conformance pack that you want to delete.
	//
	// This member is required.
	OrganizationConformancePackName *string
}

type DeleteOrganizationConformancePackOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteOrganizationConformancePackMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteOrganizationConformancePack{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteOrganizationConformancePack{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteOrganizationConformancePack(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "DeleteOrganizationConformancePack",
	}
}
