// Code generated by smithy-go-codegen DO NOT EDIT.

package datapipeline

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datapipeline/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds tasks, schedules, and preconditions to the specified pipeline. You can use
// PutPipelineDefinition to populate a new pipeline. PutPipelineDefinition also
// validates the configuration as it adds it to the pipeline. Changes to the
// pipeline are saved unless one of the following three validation errors exists in
// the pipeline.
//
// * An object is missing a name or identifier field.
//
// * A string or
// reference field is empty.
//
// * The number of objects in the pipeline exceeds the
// maximum allowed objects.
//
// * The pipeline is in a FINISHED state.
//
// Pipeline
// object definitions are passed to the PutPipelineDefinition action and returned
// by the GetPipelineDefinition action. Example 1 This example sets an valid
// pipeline configuration and returns success. POST / HTTP/1.1 Content-Type:
// application/x-amz-json-1.1 X-Amz-Target: DataPipeline.PutPipelineDefinition
// Content-Length: 914 Host: datapipeline.us-east-1.amazonaws.com X-Amz-Date: Mon,
// 12 Nov 2012 17:49:52 GMT Authorization: AuthParams {"pipelineId":
// "df-0937003356ZJEXAMPLE", "pipelineObjects": [ {"id": "Default", "name":
// "Default", "fields": [ {"key": "workerGroup", "stringValue": "workerGroup"} ] },
// {"id": "Schedule", "name": "Schedule", "fields": [ {"key": "startDateTime",
// "stringValue": "2012-12-12T00:00:00"}, {"key": "type", "stringValue":
// "Schedule"}, {"key": "period", "stringValue": "1 hour"}, {"key": "endDateTime",
// "stringValue": "2012-12-21T18:00:00"} ] }, {"id": "SayHello", "name":
// "SayHello", "fields": [ {"key": "type", "stringValue": "ShellCommandActivity"},
// {"key": "command", "stringValue": "echo hello"}, {"key": "parent", "refValue":
// "Default"}, {"key": "schedule", "refValue": "Schedule"} ] } ] } HTTP/1.1 200
// x-amzn-RequestId: f74afc14-0754-11e2-af6f-6bc7a6be60d9 Content-Type:
// application/x-amz-json-1.1 Content-Length: 18 Date: Mon, 12 Nov 2012 17:50:53
// GMT {"errored": false} Example 2 This example sets an invalid pipeline
// configuration (the value for workerGroup is an empty string) and returns an
// error message. POST / HTTP/1.1 Content-Type: application/x-amz-json-1.1
// X-Amz-Target: DataPipeline.PutPipelineDefinition Content-Length: 903 Host:
// datapipeline.us-east-1.amazonaws.com X-Amz-Date: Mon, 12 Nov 2012 17:49:52 GMT
// Authorization: AuthParams {"pipelineId": "df-06372391ZG65EXAMPLE",
// "pipelineObjects": [ {"id": "Default", "name": "Default", "fields": [ {"key":
// "workerGroup", "stringValue": ""} ] }, {"id": "Schedule", "name": "Schedule",
// "fields": [ {"key": "startDateTime", "stringValue": "2012-09-25T17:00:00"},
// {"key": "type", "stringValue": "Schedule"}, {"key": "period", "stringValue": "1
// hour"}, {"key": "endDateTime", "stringValue": "2012-09-25T18:00:00"} ] }, {"id":
// "SayHello", "name": "SayHello", "fields": [ {"key": "type", "stringValue":
// "ShellCommandActivity"}, {"key": "command", "stringValue": "echo hello"},
// {"key": "parent", "refValue": "Default"}, {"key": "schedule", "refValue":
// "Schedule"} ] } ] } HTTP/1.1 200 x-amzn-RequestId:
// f74afc14-0754-11e2-af6f-6bc7a6be60d9 Content-Type: application/x-amz-json-1.1
// Content-Length: 18 Date: Mon, 12 Nov 2012 17:50:53 GMT {"__type":
// "com.amazon.setl.webservice#InvalidRequestException", "message": "Pipeline
// definition has errors: Could not save the pipeline definition due to FATAL
// errors: [com.amazon.setl.webservice.ValidationError@108d7ea9] Please call
// Validate to validate your pipeline"}
func (c *Client) PutPipelineDefinition(ctx context.Context, params *PutPipelineDefinitionInput, optFns ...func(*Options)) (*PutPipelineDefinitionOutput, error) {
	if params == nil {
		params = &PutPipelineDefinitionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutPipelineDefinition", params, optFns, c.addOperationPutPipelineDefinitionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutPipelineDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for PutPipelineDefinition.
type PutPipelineDefinitionInput struct {

	// The ID of the pipeline.
	//
	// This member is required.
	PipelineId *string

	// The objects that define the pipeline. These objects overwrite the existing
	// pipeline definition.
	//
	// This member is required.
	PipelineObjects []types.PipelineObject

	// The parameter objects used with the pipeline.
	ParameterObjects []types.ParameterObject

	// The parameter values used with the pipeline.
	ParameterValues []types.ParameterValue

	noSmithyDocumentSerde
}

// Contains the output of PutPipelineDefinition.
type PutPipelineDefinitionOutput struct {

	// Indicates whether there were validation errors, and the pipeline definition is
	// stored but cannot be activated until you correct the pipeline and call
	// PutPipelineDefinition to commit the corrected pipeline.
	//
	// This member is required.
	Errored bool

	// The validation errors that are associated with the objects defined in
	// pipelineObjects.
	ValidationErrors []types.ValidationError

	// The validation warnings that are associated with the objects defined in
	// pipelineObjects.
	ValidationWarnings []types.ValidationWarning

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutPipelineDefinitionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutPipelineDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutPipelineDefinition{}, middleware.After)
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
	if err = addOpPutPipelineDefinitionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutPipelineDefinition(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutPipelineDefinition(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datapipeline",
		OperationName: "PutPipelineDefinition",
	}
}
