// Code generated by smithy-go-codegen DO NOT EDIT.

package machinelearning

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Assigns the DELETED status to a BatchPrediction, rendering it unusable. After
// using the DeleteBatchPrediction operation, you can use the GetBatchPrediction
// operation to verify that the status of the BatchPrediction changed to DELETED.
// Caution: The result of the DeleteBatchPrediction operation is irreversible.
func (c *Client) DeleteBatchPrediction(ctx context.Context, params *DeleteBatchPredictionInput, optFns ...func(*Options)) (*DeleteBatchPredictionOutput, error) {
	if params == nil {
		params = &DeleteBatchPredictionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteBatchPrediction", params, optFns, c.addOperationDeleteBatchPredictionMiddlewares)
	if err != nil {
		return nil, err
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

	noSmithyDocumentSerde
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

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteBatchPredictionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteBatchPrediction{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteBatchPrediction{}, middleware.After)
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
	if err = addOpDeleteBatchPredictionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteBatchPrediction(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteBatchPrediction(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "machinelearning",
		OperationName: "DeleteBatchPrediction",
	}
}
