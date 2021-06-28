// Code generated by smithy-go-codegen DO NOT EDIT.

package codeguruprofiler

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codeguruprofiler/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sends feedback to CodeGuru Profiler about whether the anomaly detected by the
// analysis is useful or not.
func (c *Client) SubmitFeedback(ctx context.Context, params *SubmitFeedbackInput, optFns ...func(*Options)) (*SubmitFeedbackOutput, error) {
	if params == nil {
		params = &SubmitFeedbackInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SubmitFeedback", params, optFns, c.addOperationSubmitFeedbackMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SubmitFeedbackOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The structure representing the SubmitFeedbackRequest.
type SubmitFeedbackInput struct {

	// The universally unique identifier (UUID) of the AnomalyInstance
	// (https://docs.aws.amazon.com/codeguru/latest/profiler-api/API_AnomalyInstance.html)
	// object that is included in the analysis data.
	//
	// This member is required.
	AnomalyInstanceId *string

	// The name of the profiling group that is associated with the analysis data.
	//
	// This member is required.
	ProfilingGroupName *string

	// The feedback tpye. Thee are two valid values, Positive and Negative.
	//
	// This member is required.
	Type types.FeedbackType

	// Optional feedback about this anomaly.
	Comment *string

	noSmithyDocumentSerde
}

// The structure representing the SubmitFeedbackResponse.
type SubmitFeedbackOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSubmitFeedbackMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSubmitFeedback{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSubmitFeedback{}, middleware.After)
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
	if err = addOpSubmitFeedbackValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSubmitFeedback(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSubmitFeedback(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeguru-profiler",
		OperationName: "SubmitFeedback",
	}
}
