// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehendmedical

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/comprehendmedical/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Starts an asynchronous job to detect protected health information (PHI). Use the
// DescribePHIDetectionJob operation to track the status of a job.
func (c *Client) StartPHIDetectionJob(ctx context.Context, params *StartPHIDetectionJobInput, optFns ...func(*Options)) (*StartPHIDetectionJobOutput, error) {
	stack := middleware.NewStack("StartPHIDetectionJob", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpStartPHIDetectionJobMiddlewares(stack)
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
	addIdempotencyToken_opStartPHIDetectionJobMiddleware(stack, options)
	addOpStartPHIDetectionJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartPHIDetectionJob(options.Region), middleware.Before)
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
			OperationName: "StartPHIDetectionJob",
			Err:           err,
		}
	}
	out := result.(*StartPHIDetectionJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartPHIDetectionJobInput struct {

	// The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM)
	// role that grants Amazon Comprehend Medical read access to your input data. For
	// more information, see  Role-Based Permissions Required for Asynchronous
	// Operations
	// (https://docs.aws.amazon.com/comprehend/latest/dg/access-control-managing-permissions-med.html#auth-role-permissions-med).
	//
	// This member is required.
	DataAccessRoleArn *string

	// Specifies the format and location of the input data for the job.
	//
	// This member is required.
	InputDataConfig *types.InputDataConfig

	// The language of the input documents. All documents must be in the same language.
	//
	// This member is required.
	LanguageCode types.LanguageCode

	// Specifies where to send the output files.
	//
	// This member is required.
	OutputDataConfig *types.OutputDataConfig

	// A unique identifier for the request. If you don't set the client request token,
	// Amazon Comprehend Medical generates one.
	ClientRequestToken *string

	// The identifier of the job.
	JobName *string

	// An AWS Key Management Service key to encrypt your output files. If you do not
	// specify a key, the files are written in plain text.
	KMSKey *string
}

type StartPHIDetectionJobOutput struct {

	// The identifier generated for the job. To get the status of a job, use this
	// identifier with the DescribePHIDetectionJob operation.
	JobId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpStartPHIDetectionJobMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpStartPHIDetectionJob{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartPHIDetectionJob{}, middleware.After)
}

type idempotencyToken_initializeOpStartPHIDetectionJob struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartPHIDetectionJob) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartPHIDetectionJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartPHIDetectionJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartPHIDetectionJobInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opStartPHIDetectionJobMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpStartPHIDetectionJob{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartPHIDetectionJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehendmedical",
		OperationName: "StartPHIDetectionJob",
	}
}
