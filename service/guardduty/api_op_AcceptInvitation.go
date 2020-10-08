// Code generated by smithy-go-codegen DO NOT EDIT.

package guardduty

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Accepts the invitation to be monitored by a master GuardDuty account.
func (c *Client) AcceptInvitation(ctx context.Context, params *AcceptInvitationInput, optFns ...func(*Options)) (*AcceptInvitationOutput, error) {
	stack := middleware.NewStack("AcceptInvitation", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpAcceptInvitationMiddlewares(stack)
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
	addOpAcceptInvitationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAcceptInvitation(options.Region), middleware.Before)
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
			OperationName: "AcceptInvitation",
			Err:           err,
		}
	}
	out := result.(*AcceptInvitationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AcceptInvitationInput struct {

	// The unique ID of the detector of the GuardDuty member account.
	//
	// This member is required.
	DetectorId *string

	// The value that is used to validate the master account to the member account.
	//
	// This member is required.
	InvitationId *string

	// The account ID of the master GuardDuty account whose invitation you're
	// accepting.
	//
	// This member is required.
	MasterId *string
}

type AcceptInvitationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpAcceptInvitationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpAcceptInvitation{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpAcceptInvitation{}, middleware.After)
}

func newServiceMetadataMiddleware_opAcceptInvitation(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "guardduty",
		OperationName: "AcceptInvitation",
	}
}
