// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns information about the requested human task user interface (worker task
// template).
func (c *Client) DescribeHumanTaskUi(ctx context.Context, params *DescribeHumanTaskUiInput, optFns ...func(*Options)) (*DescribeHumanTaskUiOutput, error) {
	if params == nil {
		params = &DescribeHumanTaskUiInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeHumanTaskUi", params, optFns, c.addOperationDescribeHumanTaskUiMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeHumanTaskUiOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeHumanTaskUiInput struct {

	// The name of the human task user interface (worker task template) you want
	// information about.
	//
	// This member is required.
	HumanTaskUiName *string

	noSmithyDocumentSerde
}

type DescribeHumanTaskUiOutput struct {

	// The timestamp when the human task user interface was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The Amazon Resource Name (ARN) of the human task user interface (worker task
	// template).
	//
	// This member is required.
	HumanTaskUiArn *string

	// The name of the human task user interface (worker task template).
	//
	// This member is required.
	HumanTaskUiName *string

	// Container for user interface template information.
	//
	// This member is required.
	UiTemplate *types.UiTemplateInfo

	// The status of the human task user interface (worker task template). Valid values
	// are listed below.
	HumanTaskUiStatus types.HumanTaskUiStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeHumanTaskUiMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeHumanTaskUi{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeHumanTaskUi{}, middleware.After)
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
	if err = addOpDescribeHumanTaskUiValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeHumanTaskUi(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeHumanTaskUi(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DescribeHumanTaskUi",
	}
}
