// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a definition for a job that monitors model quality and drift. For
// information about model monitor, see Amazon SageMaker Model Monitor
// (https://docs.aws.amazon.com/sagemaker/latest/dg/model-monitor.html).
func (c *Client) CreateModelQualityJobDefinition(ctx context.Context, params *CreateModelQualityJobDefinitionInput, optFns ...func(*Options)) (*CreateModelQualityJobDefinitionOutput, error) {
	if params == nil {
		params = &CreateModelQualityJobDefinitionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateModelQualityJobDefinition", params, optFns, c.addOperationCreateModelQualityJobDefinitionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateModelQualityJobDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateModelQualityJobDefinitionInput struct {

	// The name of the monitoring job definition.
	//
	// This member is required.
	JobDefinitionName *string

	// Identifies the resources to deploy for a monitoring job.
	//
	// This member is required.
	JobResources *types.MonitoringResources

	// The container that runs the monitoring job.
	//
	// This member is required.
	ModelQualityAppSpecification *types.ModelQualityAppSpecification

	// A list of the inputs that are monitored. Currently endpoints are supported.
	//
	// This member is required.
	ModelQualityJobInput *types.ModelQualityJobInput

	// The output configuration for monitoring jobs.
	//
	// This member is required.
	ModelQualityJobOutputConfig *types.MonitoringOutputConfig

	// The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume
	// to perform tasks on your behalf.
	//
	// This member is required.
	RoleArn *string

	// Specifies the constraints and baselines for the monitoring job.
	ModelQualityBaselineConfig *types.ModelQualityBaselineConfig

	// Specifies the network configuration for the monitoring job.
	NetworkConfig *types.MonitoringNetworkConfig

	// A time limit for how long the monitoring job is allowed to run before stopping.
	StoppingCondition *types.MonitoringStoppingCondition

	// (Optional) An array of key-value pairs. For more information, see Using Cost
	// Allocation Tags
	// (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-whatURL)
	// in the Amazon Web Services Billing and Cost Management User Guide.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateModelQualityJobDefinitionOutput struct {

	// The Amazon Resource Name (ARN) of the model quality monitoring job.
	//
	// This member is required.
	JobDefinitionArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateModelQualityJobDefinitionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateModelQualityJobDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateModelQualityJobDefinition{}, middleware.After)
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
	if err = addOpCreateModelQualityJobDefinitionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateModelQualityJobDefinition(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateModelQualityJobDefinition(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "CreateModelQualityJobDefinition",
	}
}
