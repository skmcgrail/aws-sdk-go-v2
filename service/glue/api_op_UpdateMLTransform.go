// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an existing machine learning transform. Call this operation to tune the
// algorithm parameters to achieve better results. After calling this operation,
// you can call the StartMLEvaluationTaskRun operation to assess how well your new
// parameters achieved your goals (such as improving the quality of your machine
// learning transform, or making it more cost-effective).
func (c *Client) UpdateMLTransform(ctx context.Context, params *UpdateMLTransformInput, optFns ...func(*Options)) (*UpdateMLTransformOutput, error) {
	if params == nil {
		params = &UpdateMLTransformInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateMLTransform", params, optFns, c.addOperationUpdateMLTransformMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateMLTransformOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateMLTransformInput struct {

	// A unique identifier that was generated when the transform was created.
	//
	// This member is required.
	TransformId *string

	// A description of the transform. The default is an empty string.
	Description *string

	// This value determines which version of Glue this machine learning transform is
	// compatible with. Glue 1.0 is recommended for most customers. If the value is not
	// set, the Glue compatibility defaults to Glue 0.9. For more information, see Glue
	// Versions
	// (https://docs.aws.amazon.com/glue/latest/dg/release-notes.html#release-notes-versions)
	// in the developer guide.
	GlueVersion *string

	// The number of Glue data processing units (DPUs) that are allocated to task runs
	// for this transform. You can allocate from 2 to 100 DPUs; the default is 10. A
	// DPU is a relative measure of processing power that consists of 4 vCPUs of
	// compute capacity and 16 GB of memory. For more information, see the Glue pricing
	// page (https://aws.amazon.com/glue/pricing/). When the WorkerType field is set to
	// a value other than Standard, the MaxCapacity field is set automatically and
	// becomes read-only.
	MaxCapacity *float64

	// The maximum number of times to retry a task for this transform after a task run
	// fails.
	MaxRetries *int32

	// The unique name that you gave the transform when you created it.
	Name *string

	// The number of workers of a defined workerType that are allocated when this task
	// runs.
	NumberOfWorkers *int32

	// The configuration parameters that are specific to the transform type (algorithm)
	// used. Conditionally dependent on the transform type.
	Parameters *types.TransformParameters

	// The name or Amazon Resource Name (ARN) of the IAM role with the required
	// permissions.
	Role *string

	// The timeout for a task run for this transform in minutes. This is the maximum
	// time that a task run for this transform can consume resources before it is
	// terminated and enters TIMEOUT status. The default is 2,880 minutes (48 hours).
	Timeout *int32

	// The type of predefined worker that is allocated when this task runs. Accepts a
	// value of Standard, G.1X, or G.2X.
	//
	// * For the Standard worker type, each worker
	// provides 4 vCPU, 16 GB of memory and a 50GB disk, and 2 executors per worker.
	//
	// *
	// For the G.1X worker type, each worker provides 4 vCPU, 16 GB of memory and a
	// 64GB disk, and 1 executor per worker.
	//
	// * For the G.2X worker type, each worker
	// provides 8 vCPU, 32 GB of memory and a 128GB disk, and 1 executor per worker.
	WorkerType types.WorkerType

	noSmithyDocumentSerde
}

type UpdateMLTransformOutput struct {

	// The unique identifier for the transform that was updated.
	TransformId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateMLTransformMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateMLTransform{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateMLTransform{}, middleware.After)
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
	if err = addOpUpdateMLTransformValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateMLTransform(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateMLTransform(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "UpdateMLTransform",
	}
}
