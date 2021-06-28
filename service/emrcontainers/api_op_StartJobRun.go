// Code generated by smithy-go-codegen DO NOT EDIT.

package emrcontainers

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/emrcontainers/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts a job run. A job run is a unit of work, such as a Spark jar, PySpark
// script, or SparkSQL query, that you submit to Amazon EMR on EKS.
func (c *Client) StartJobRun(ctx context.Context, params *StartJobRunInput, optFns ...func(*Options)) (*StartJobRunOutput, error) {
	if params == nil {
		params = &StartJobRunInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartJobRun", params, optFns, c.addOperationStartJobRunMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartJobRunOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartJobRunInput struct {

	// The client idempotency token of the job run request.
	//
	// This member is required.
	ClientToken *string

	// The execution role ARN for the job run.
	//
	// This member is required.
	ExecutionRoleArn *string

	// The job driver for the job run.
	//
	// This member is required.
	JobDriver *types.JobDriver

	// The Amazon EMR release version to use for the job run.
	//
	// This member is required.
	ReleaseLabel *string

	// The virtual cluster ID for which the job run request is submitted.
	//
	// This member is required.
	VirtualClusterId *string

	// The configuration overrides for the job run.
	ConfigurationOverrides *types.ConfigurationOverrides

	// The name of the job run.
	Name *string

	// The tags assigned to job runs.
	Tags map[string]string

	noSmithyDocumentSerde
}

type StartJobRunOutput struct {

	// This output lists the ARN of job run.
	Arn *string

	// This output displays the started job run ID.
	Id *string

	// This output displays the name of the started job run.
	Name *string

	// This output displays the virtual cluster ID for which the job run was submitted.
	VirtualClusterId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartJobRunMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartJobRun{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartJobRun{}, middleware.After)
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
	if err = addIdempotencyToken_opStartJobRunMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartJobRunValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartJobRun(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpStartJobRun struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartJobRun) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartJobRun) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartJobRunInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartJobRunInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opStartJobRunMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpStartJobRun{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartJobRun(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "emr-containers",
		OperationName: "StartJobRun",
	}
}
