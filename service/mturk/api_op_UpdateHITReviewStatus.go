// Code generated by smithy-go-codegen DO NOT EDIT.

package mturk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The UpdateHITReviewStatus operation updates the status of a HIT. If the status
// is Reviewable, this operation can update the status to Reviewing, or it can
// revert a Reviewing HIT back to the Reviewable status.
func (c *Client) UpdateHITReviewStatus(ctx context.Context, params *UpdateHITReviewStatusInput, optFns ...func(*Options)) (*UpdateHITReviewStatusOutput, error) {
	stack := middleware.NewStack("UpdateHITReviewStatus", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateHITReviewStatusMiddlewares(stack)
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
	addOpUpdateHITReviewStatusValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateHITReviewStatus(options.Region), middleware.Before)
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
			OperationName: "UpdateHITReviewStatus",
			Err:           err,
		}
	}
	out := result.(*UpdateHITReviewStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateHITReviewStatusInput struct {
	// Specifies how to update the HIT status. Default is False.
	//
	//     * Setting this to
	// false will only transition a HIT from Reviewable to Reviewing
	//
	//     * Setting
	// this to true will only transition a HIT from Reviewing to Reviewable
	Revert *bool
	// The ID of the HIT to update.
	HITId *string
}

type UpdateHITReviewStatusOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateHITReviewStatusMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateHITReviewStatus{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateHITReviewStatus{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateHITReviewStatus(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mturk-requester",
		OperationName: "UpdateHITReviewStatus",
	}
}