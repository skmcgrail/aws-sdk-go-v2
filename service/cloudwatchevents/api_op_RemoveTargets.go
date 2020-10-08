// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchevents

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes the specified targets from the specified rule. When the rule is
// triggered, those targets are no longer be invoked.  <p>When you remove a target,
// when the associated rule triggers, removed targets might continue to be invoked.
// Allow a short period of time for changes to take effect.</p> <p>This action can
// partially fail if too many requests are made at the same time. If that happens,
// <code>FailedEntryCount</code> is non-zero in the response and each entry in
// <code>FailedEntries</code> provides the ID of the failed target and the error
// code.</p>
func (c *Client) RemoveTargets(ctx context.Context, params *RemoveTargetsInput, optFns ...func(*Options)) (*RemoveTargetsOutput, error) {
	stack := middleware.NewStack("RemoveTargets", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpRemoveTargetsMiddlewares(stack)
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
	addOpRemoveTargetsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRemoveTargets(options.Region), middleware.Before)
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
			OperationName: "RemoveTargets",
			Err:           err,
		}
	}
	out := result.(*RemoveTargetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RemoveTargetsInput struct {

	// The IDs of the targets to remove from the rule.
	//
	// This member is required.
	Ids []*string

	// The name of the rule.
	//
	// This member is required.
	Rule *string

	// The name of the event bus associated with the rule.
	EventBusName *string

	// If this is a managed rule, created by an AWS service on your behalf, you must
	// specify Force as True to remove targets. This parameter is ignored for rules
	// that are not managed rules. You can check whether a rule is a managed rule by
	// using DescribeRule or ListRules and checking the ManagedBy field of the
	// response.
	Force *bool
}

type RemoveTargetsOutput struct {

	// The failed target entries.
	FailedEntries []*types.RemoveTargetsResultEntry

	// The number of failed entries.
	FailedEntryCount *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpRemoveTargetsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpRemoveTargets{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpRemoveTargets{}, middleware.After)
}

func newServiceMetadataMiddleware_opRemoveTargets(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "events",
		OperationName: "RemoveTargets",
	}
}
