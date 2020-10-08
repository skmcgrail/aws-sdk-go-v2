// Code generated by smithy-go-codegen DO NOT EDIT.

package codeartifact

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the resource policy attached to the specified domain. The policy is a
// resource-based policy, not an identity-based policy. For more information, see
// Identity-based policies and resource-based policies
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_identity-vs-resource.html)
// in the AWS Identity and Access Management User Guide.
func (c *Client) GetDomainPermissionsPolicy(ctx context.Context, params *GetDomainPermissionsPolicyInput, optFns ...func(*Options)) (*GetDomainPermissionsPolicyOutput, error) {
	stack := middleware.NewStack("GetDomainPermissionsPolicy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetDomainPermissionsPolicyMiddlewares(stack)
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
	addOpGetDomainPermissionsPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetDomainPermissionsPolicy(options.Region), middleware.Before)
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
			OperationName: "GetDomainPermissionsPolicy",
			Err:           err,
		}
	}
	out := result.(*GetDomainPermissionsPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDomainPermissionsPolicyInput struct {

	// The name of the domain to which the resource policy is attached.
	//
	// This member is required.
	Domain *string

	// The 12-digit account number of the AWS account that owns the domain. It does not
	// include dashes or spaces.
	DomainOwner *string
}

type GetDomainPermissionsPolicyOutput struct {

	// The returned resource policy.
	Policy *types.ResourcePolicy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetDomainPermissionsPolicyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetDomainPermissionsPolicy{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDomainPermissionsPolicy{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetDomainPermissionsPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeartifact",
		OperationName: "GetDomainPermissionsPolicy",
	}
}
