// Code generated by smithy-go-codegen DO NOT EDIT.

package mturk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The AcceptQualificationRequest operation approves a Worker's request for a
// Qualification. Only the owner of the Qualification type can grant a
// Qualification request for that type. A successful request for the
// AcceptQualificationRequest operation returns with no errors and an empty body.
func (c *Client) AcceptQualificationRequest(ctx context.Context, params *AcceptQualificationRequestInput, optFns ...func(*Options)) (*AcceptQualificationRequestOutput, error) {
	stack := middleware.NewStack("AcceptQualificationRequest", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAcceptQualificationRequestMiddlewares(stack)
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
	addOpAcceptQualificationRequestValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAcceptQualificationRequest(options.Region), middleware.Before)
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
			OperationName: "AcceptQualificationRequest",
			Err:           err,
		}
	}
	out := result.(*AcceptQualificationRequestOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AcceptQualificationRequestInput struct {

	// The ID of the Qualification request, as returned by the GetQualificationRequests
	// operation.
	//
	// This member is required.
	QualificationRequestId *string

	// The value of the Qualification. You can omit this value if you are using the
	// presence or absence of the Qualification as the basis for a HIT requirement.
	IntegerValue *int32
}

type AcceptQualificationRequestOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAcceptQualificationRequestMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAcceptQualificationRequest{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAcceptQualificationRequest{}, middleware.After)
}

func newServiceMetadataMiddleware_opAcceptQualificationRequest(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mturk-requester",
		OperationName: "AcceptQualificationRequest",
	}
}
