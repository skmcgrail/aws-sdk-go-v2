// Code generated by smithy-go-codegen DO NOT EDIT.

package route53domains

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53domains/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This operation updates the contact information for a particular domain. You must
// specify information for at least one contact: registrant, administrator, or
// technical. If the update is successful, this method returns an operation ID that
// you can use to track the progress and completion of the action. If the request
// is not completed successfully, the domain registrant will be notified by email.
func (c *Client) UpdateDomainContact(ctx context.Context, params *UpdateDomainContactInput, optFns ...func(*Options)) (*UpdateDomainContactOutput, error) {
	stack := middleware.NewStack("UpdateDomainContact", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateDomainContactMiddlewares(stack)
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
	addOpUpdateDomainContactValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDomainContact(options.Region), middleware.Before)

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
			OperationName: "UpdateDomainContact",
			Err:           err,
		}
	}
	out := result.(*UpdateDomainContactOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The UpdateDomainContact request includes the following elements.
type UpdateDomainContactInput struct {
	// Provides detailed contact information.
	AdminContact *types.ContactDetail
	// Provides detailed contact information.
	RegistrantContact *types.ContactDetail
	// The name of the domain that you want to update contact information for.
	DomainName *string
	// Provides detailed contact information.
	TechContact *types.ContactDetail
}

// The UpdateDomainContact response includes the following element.
type UpdateDomainContactOutput struct {
	// Identifier for tracking the progress of the request. To query the operation
	// status, use GetOperationDetail
	// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetOperationDetail.html).
	OperationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateDomainContactMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateDomainContact{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateDomainContact{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateDomainContact(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53domains",
		OperationName: "UpdateDomainContact",
	}
}