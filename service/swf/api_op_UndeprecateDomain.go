// Code generated by smithy-go-codegen DO NOT EDIT.

package swf

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Undeprecates a previously deprecated domain. After a domain has been
// undeprecated it can be used to create new workflow executions or register new
// types. This operation is eventually consistent. The results are best effort and
// may not exactly reflect recent updates and changes. Access Control You can use
// IAM policies to control this action's access to Amazon SWF resources as
// follows:
//
//     * Use a Resource element with the domain name to limit the action
// to only specified domains.
//
//     * Use an Action element to allow or deny
// permission to call this action.
//
//     * You cannot use an IAM policy to constrain
// this action's parameters.
//
// If the caller doesn't have sufficient permissions to
// invoke the action, or the parameter values fall outside the specified
// constraints, the action fails. The associated event attribute's cause parameter
// is set to OPERATION_NOT_PERMITTED. For details and example IAM policies, see
// Using IAM to Manage Access to Amazon SWF Workflows
// (https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
// in the Amazon SWF Developer Guide.
func (c *Client) UndeprecateDomain(ctx context.Context, params *UndeprecateDomainInput, optFns ...func(*Options)) (*UndeprecateDomainOutput, error) {
	stack := middleware.NewStack("UndeprecateDomain", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson10_serdeOpUndeprecateDomainMiddlewares(stack)
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
	addOpUndeprecateDomainValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUndeprecateDomain(options.Region), middleware.Before)

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
			OperationName: "UndeprecateDomain",
			Err:           err,
		}
	}
	out := result.(*UndeprecateDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UndeprecateDomainInput struct {
	// The name of the domain of the deprecated workflow type.
	Name *string
}

type UndeprecateDomainOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson10_serdeOpUndeprecateDomainMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson10_serializeOpUndeprecateDomain{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson10_deserializeOpUndeprecateDomain{}, middleware.After)
}

func newServiceMetadataMiddleware_opUndeprecateDomain(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "swf",
		OperationName: "UndeprecateDomain",
	}
}