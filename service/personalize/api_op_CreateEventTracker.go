// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an event tracker that you use when adding event data to a specified
// dataset group using the PutEvents
// (https://docs.aws.amazon.com/personalize/latest/dg/API_UBS_PutEvents.html) API.
// Only one event tracker can be associated with a dataset group. You will get an
// error if you call CreateEventTracker using the same dataset group as an existing
// event tracker. When you create an event tracker, the response includes a
// tracking ID, which you pass as a parameter when you use the PutEvents
// (https://docs.aws.amazon.com/personalize/latest/dg/API_UBS_PutEvents.html)
// operation. Amazon Personalize then appends the event data to the Interactions
// dataset of the dataset group you specify in your event tracker. The event
// tracker can be in one of the following states:
//
// * CREATE PENDING > CREATE
// IN_PROGRESS > ACTIVE -or- CREATE FAILED
//
// * DELETE PENDING > DELETE
// IN_PROGRESS
//
// To get the status of the event tracker, call DescribeEventTracker.
// The event tracker must be in the ACTIVE state before using the tracking ID.
// Related APIs
//
// * ListEventTrackers
//
// * DescribeEventTracker
//
// * DeleteEventTracker
func (c *Client) CreateEventTracker(ctx context.Context, params *CreateEventTrackerInput, optFns ...func(*Options)) (*CreateEventTrackerOutput, error) {
	if params == nil {
		params = &CreateEventTrackerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateEventTracker", params, optFns, c.addOperationCreateEventTrackerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateEventTrackerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateEventTrackerInput struct {

	// The Amazon Resource Name (ARN) of the dataset group that receives the event
	// data.
	//
	// This member is required.
	DatasetGroupArn *string

	// The name for the event tracker.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type CreateEventTrackerOutput struct {

	// The ARN of the event tracker.
	EventTrackerArn *string

	// The ID of the event tracker. Include this ID in requests to the PutEvents
	// (https://docs.aws.amazon.com/personalize/latest/dg/API_UBS_PutEvents.html) API.
	TrackingId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateEventTrackerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateEventTracker{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateEventTracker{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateEventTrackerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateEventTracker(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateEventTracker(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "CreateEventTracker",
	}
}
