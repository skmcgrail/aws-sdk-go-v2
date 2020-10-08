// Code generated by smithy-go-codegen DO NOT EDIT.

package machinelearning

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Assigns the DELETED status to a BatchPrediction, rendering it unusable. After
// using the DeleteBatchPrediction operation, you can use the GetBatchPrediction ()
// operation to verify that the status of the BatchPrediction changed to DELETED.
// <p> <b>Caution:</b> The result of the <code>DeleteBatchPrediction</code>
// operation is irreversible.</p>
func (c *Client) DeleteBatchPrediction(ctx context.Context, params *DeleteBatchPredictionInput, optFns ...func(*Options)) (*DeleteBatchPredictionOutput, error) {
	stack := middleware.NewStack("DeleteBatchPrediction", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteBatchPredictionMiddlewares(stack)
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
	addOpDeleteBatchPredictionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteBatchPrediction(options.Region), middleware.Before)
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
			OperationName: "DeleteBatchPrediction",
			Err:           err,
		}
	}
	out := result.(*DeleteBatchPredictionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteBatchPredictionInput struct {

	// A user-supplied ID that uniquely identifies the BatchPrediction.
	//
	// This member is required.
	BatchPredictionId *string
}

// Represents the output of a DeleteBatchPrediction operation. You can use the
// GetBatchPrediction operation and check the value of the Status parameter to see
// whether a BatchPrediction is marked as DELETED.
type DeleteBatchPredictionOutput struct {

	// A user-supplied ID that uniquely identifies the BatchPrediction. This value
	// should be identical to the value of the BatchPredictionID in the request.
	BatchPredictionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteBatchPredictionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteBatchPrediction{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteBatchPrediction{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteBatchPrediction(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "machinelearning",
		OperationName: "DeleteBatchPrediction",
	}
}
