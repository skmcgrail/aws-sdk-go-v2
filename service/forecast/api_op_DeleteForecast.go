// Code generated by smithy-go-codegen DO NOT EDIT.

package forecast

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a forecast created using the CreateForecast () operation. You can delete
// only forecasts that have a status of ACTIVE or CREATE_FAILED. To get the status,
// use the DescribeForecast () operation. You can't delete a forecast while it is
// being exported. After a forecast is deleted, you can no longer query the
// forecast.
func (c *Client) DeleteForecast(ctx context.Context, params *DeleteForecastInput, optFns ...func(*Options)) (*DeleteForecastOutput, error) {
	stack := middleware.NewStack("DeleteForecast", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteForecastMiddlewares(stack)
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
	addOpDeleteForecastValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteForecast(options.Region), middleware.Before)
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
			OperationName: "DeleteForecast",
			Err:           err,
		}
	}
	out := result.(*DeleteForecastOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteForecastInput struct {

	// The Amazon Resource Name (ARN) of the forecast to delete.
	//
	// This member is required.
	ForecastArn *string
}

type DeleteForecastOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteForecastMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteForecast{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteForecast{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteForecast(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "forecast",
		OperationName: "DeleteForecast",
	}
}
