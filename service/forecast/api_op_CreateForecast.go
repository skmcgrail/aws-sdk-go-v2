// Code generated by smithy-go-codegen DO NOT EDIT.

package forecast

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/forecast/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a forecast for each item in the TARGET_TIME_SERIES dataset that was used
// to train the predictor. This is known as inference. To retrieve the forecast for
// a single item at low latency, use the operation. To export the complete forecast
// into your Amazon Simple Storage Service (Amazon S3) bucket, use the
// CreateForecastExportJob operation. The range of the forecast is determined by
// the ForecastHorizon value, which you specify in the CreatePredictor request.
// When you query a forecast, you can request a specific date range within the
// forecast. To get a list of all your forecasts, use the ListForecasts operation.
// The forecasts generated by Amazon Forecast are in the same time zone as the
// dataset that was used to create the predictor. For more information, see
// howitworks-forecast. The Status of the forecast must be ACTIVE before you can
// query or export the forecast. Use the DescribeForecast operation to get the
// status.
func (c *Client) CreateForecast(ctx context.Context, params *CreateForecastInput, optFns ...func(*Options)) (*CreateForecastOutput, error) {
	if params == nil {
		params = &CreateForecastInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateForecast", params, optFns, c.addOperationCreateForecastMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateForecastOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateForecastInput struct {

	// A name for the forecast.
	//
	// This member is required.
	ForecastName *string

	// The Amazon Resource Name (ARN) of the predictor to use to generate the forecast.
	//
	// This member is required.
	PredictorArn *string

	// The quantiles at which probabilistic forecasts are generated. You can currently
	// specify up to 5 quantiles per forecast. Accepted values include 0.01 to 0.99
	// (increments of .01 only) and mean. The mean forecast is different from the
	// median (0.50) when the distribution is not symmetric (for example, Beta and
	// Negative Binomial). The default value is ["0.1", "0.5", "0.9"].
	ForecastTypes []string

	// The optional metadata that you apply to the forecast to help you categorize and
	// organize them. Each tag consists of a key and an optional value, both of which
	// you define. The following basic restrictions apply to tags:
	//
	// * Maximum number of
	// tags per resource - 50.
	//
	// * For each resource, each tag key must be unique, and
	// each tag key can have only one value.
	//
	// * Maximum key length - 128 Unicode
	// characters in UTF-8.
	//
	// * Maximum value length - 256 Unicode characters in
	// UTF-8.
	//
	// * If your tagging schema is used across multiple services and resources,
	// remember that other services may have restrictions on allowed characters.
	// Generally allowed characters are: letters, numbers, and spaces representable in
	// UTF-8, and the following characters: + - = . _ : / @.
	//
	// * Tag keys and values are
	// case sensitive.
	//
	// * Do not use aws:, AWS:, or any upper or lowercase combination
	// of such as a prefix for keys as it is reserved for AWS use. You cannot edit or
	// delete tag keys with this prefix. Values can have this prefix. If a tag value
	// has aws as its prefix but the key does not, then Forecast considers it to be a
	// user tag and will count against the limit of 50 tags. Tags with only the key
	// prefix of aws do not count against your tags per resource limit.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateForecastOutput struct {

	// The Amazon Resource Name (ARN) of the forecast.
	ForecastArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateForecastMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateForecast{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateForecast{}, middleware.After)
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
	if err = addOpCreateForecastValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateForecast(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateForecast(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "forecast",
		OperationName: "CreateForecast",
	}
}
