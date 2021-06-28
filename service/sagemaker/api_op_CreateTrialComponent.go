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

// Creates a trial component, which is a stage of a machine learning trial. A trial
// is composed of one or more trial components. A trial component can be used in
// multiple trials. Trial components include pre-processing jobs, training jobs,
// and batch transform jobs. When you use SageMaker Studio or the SageMaker Python
// SDK, all experiments, trials, and trial components are automatically tracked,
// logged, and indexed. When you use the Amazon Web Services SDK for Python (Boto),
// you must use the logging APIs provided by the SDK. You can add tags to a trial
// component and then use the Search API to search for the tags.
func (c *Client) CreateTrialComponent(ctx context.Context, params *CreateTrialComponentInput, optFns ...func(*Options)) (*CreateTrialComponentOutput, error) {
	if params == nil {
		params = &CreateTrialComponentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateTrialComponent", params, optFns, c.addOperationCreateTrialComponentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateTrialComponentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTrialComponentInput struct {

	// The name of the component. The name must be unique in your Amazon Web Services
	// account and is not case-sensitive.
	//
	// This member is required.
	TrialComponentName *string

	// The name of the component as displayed. The name doesn't need to be unique. If
	// DisplayName isn't specified, TrialComponentName is displayed.
	DisplayName *string

	// When the component ended.
	EndTime *time.Time

	// The input artifacts for the component. Examples of input artifacts are datasets,
	// algorithms, hyperparameters, source code, and instance types.
	InputArtifacts map[string]types.TrialComponentArtifact

	// Metadata properties of the tracking entity, trial, or trial component.
	MetadataProperties *types.MetadataProperties

	// The output artifacts for the component. Examples of output artifacts are
	// metrics, snapshots, logs, and images.
	OutputArtifacts map[string]types.TrialComponentArtifact

	// The hyperparameters for the component.
	Parameters map[string]types.TrialComponentParameterValue

	// When the component started.
	StartTime *time.Time

	// The status of the component. States include:
	//
	// * InProgress
	//
	// * Completed
	//
	// *
	// Failed
	Status *types.TrialComponentStatus

	// A list of tags to associate with the component. You can use Search API to search
	// on the tags.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateTrialComponentOutput struct {

	// The Amazon Resource Name (ARN) of the trial component.
	TrialComponentArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateTrialComponentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateTrialComponent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateTrialComponent{}, middleware.After)
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
	if err = addOpCreateTrialComponentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTrialComponent(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateTrialComponent(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "CreateTrialComponent",
	}
}
