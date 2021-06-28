// Code generated by smithy-go-codegen DO NOT EDIT.

package appflow

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appflow/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an existing flow.
func (c *Client) UpdateFlow(ctx context.Context, params *UpdateFlowInput, optFns ...func(*Options)) (*UpdateFlowOutput, error) {
	if params == nil {
		params = &UpdateFlowInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateFlow", params, optFns, c.addOperationUpdateFlowMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateFlowOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateFlowInput struct {

	// The configuration that controls how Amazon AppFlow transfers data to the
	// destination connector.
	//
	// This member is required.
	DestinationFlowConfigList []types.DestinationFlowConfig

	// The specified name of the flow. Spaces are not allowed. Use underscores (_) or
	// hyphens (-) only.
	//
	// This member is required.
	FlowName *string

	// A list of tasks that Amazon AppFlow performs while transferring the data in the
	// flow run.
	//
	// This member is required.
	Tasks []types.Task

	// The trigger settings that determine how and when the flow runs.
	//
	// This member is required.
	TriggerConfig *types.TriggerConfig

	// A description of the flow.
	Description *string

	// Contains information about the configuration of the source connector used in the
	// flow.
	SourceFlowConfig *types.SourceFlowConfig

	noSmithyDocumentSerde
}

type UpdateFlowOutput struct {

	// Indicates the current status of the flow.
	FlowStatus types.FlowStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateFlowMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateFlow{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateFlow{}, middleware.After)
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
	if err = addOpUpdateFlowValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateFlow(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateFlow(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appflow",
		OperationName: "UpdateFlow",
	}
}
