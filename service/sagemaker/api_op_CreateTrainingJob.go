// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Starts a model training job. After training completes, Amazon SageMaker saves
// the resulting model artifacts to an Amazon S3 location that you specify. If you
// choose to host your model using Amazon SageMaker hosting services, you can use
// the resulting model artifacts as part of the model. You can also use the
// artifacts in a machine learning service other than Amazon SageMaker, provided
// that you know how to use them for inferences.  </p> <p>In the request body, you
// provide the following: </p> <ul> <li> <p> <code>AlgorithmSpecification</code> -
// Identifies the training algorithm to use. </p> </li> <li> <p>
// <code>HyperParameters</code> - Specify these algorithm-specific parameters to
// enable the estimation of model parameters during training. Hyperparameters can
// be tuned to optimize this learning process. For a list of hyperparameters for
// each training algorithm provided by Amazon SageMaker, see <a
// href="https://docs.aws.amazon.com/sagemaker/latest/dg/algos.html">Algorithms</a>.
// </p> </li> <li> <p> <code>InputDataConfig</code> - Describes the training
// dataset and the Amazon S3, EFS, or FSx location where it is stored.</p> </li>
// <li> <p> <code>OutputDataConfig</code> - Identifies the Amazon S3 bucket where
// you want Amazon SageMaker to save the results of model training. </p> <p></p>
// </li> <li> <p> <code>ResourceConfig</code> - Identifies the resources, ML
// compute instances, and ML storage volumes to deploy for model training. In
// distributed training, you specify more than one instance. </p> </li> <li> <p>
// <code>EnableManagedSpotTraining</code> - Optimize the cost of training machine
// learning models by up to 80% by using Amazon EC2 Spot instances. For more
// information, see <a
// href="https://docs.aws.amazon.com/sagemaker/latest/dg/model-managed-spot-training.html">Managed
// Spot Training</a>. </p> </li> <li> <p> <code>RoleARN</code> - The Amazon
// Resource Number (ARN) that Amazon SageMaker assumes to perform tasks on your
// behalf during model training. You must grant this role the necessary permissions
// so that Amazon SageMaker can successfully complete model training. </p> </li>
// <li> <p> <code>StoppingCondition</code> - To help cap training costs, use
// <code>MaxRuntimeInSeconds</code> to set a time limit for training. Use
// <code>MaxWaitTimeInSeconds</code> to specify how long you are willing to wait
// for a managed spot training job to complete. </p> </li> </ul> <p> For more
// information about Amazon SageMaker, see <a
// href="https://docs.aws.amazon.com/sagemaker/latest/dg/how-it-works.html">How It
// Works</a>. </p>
func (c *Client) CreateTrainingJob(ctx context.Context, params *CreateTrainingJobInput, optFns ...func(*Options)) (*CreateTrainingJobOutput, error) {
	stack := middleware.NewStack("CreateTrainingJob", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateTrainingJobMiddlewares(stack)
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
	addOpCreateTrainingJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTrainingJob(options.Region), middleware.Before)
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
			OperationName: "CreateTrainingJob",
			Err:           err,
		}
	}
	out := result.(*CreateTrainingJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTrainingJobInput struct {

	// The registry path of the Docker image that contains the training algorithm and
	// algorithm-specific metadata, including the input mode. For more information
	// about algorithms provided by Amazon SageMaker, see Algorithms
	// (https://docs.aws.amazon.com/sagemaker/latest/dg/algos.html). For information
	// about providing your own algorithms, see Using Your Own Algorithms with Amazon
	// SageMaker
	// (https://docs.aws.amazon.com/sagemaker/latest/dg/your-algorithms.html).
	//
	// This member is required.
	AlgorithmSpecification *types.AlgorithmSpecification

	// Specifies the path to the S3 location where you want to store model artifacts.
	// Amazon SageMaker creates subfolders for the artifacts.
	//
	// This member is required.
	OutputDataConfig *types.OutputDataConfig

	// The resources, including the ML compute instances and ML storage volumes, to use
	// for model training. ML storage volumes store model artifacts and incremental
	// states. Training algorithms might also use ML storage volumes for scratch space.
	// If you want Amazon SageMaker to use the ML storage volume to store the training
	// data, choose File as the TrainingInputMode in the algorithm specification. For
	// distributed training algorithms, specify an instance count greater than 1.
	//
	// This member is required.
	ResourceConfig *types.ResourceConfig

	// The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume
	// to perform tasks on your behalf. During model training, Amazon SageMaker needs
	// your permission to read input data from an S3 bucket, download a Docker image
	// that contains training code, write model artifacts to an S3 bucket, write logs
	// to Amazon CloudWatch Logs, and publish metrics to Amazon CloudWatch. You grant
	// permissions for all of these tasks to an IAM role. For more information, see
	// Amazon SageMaker Roles
	// (https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-roles.html). To be
	// able to pass this role to Amazon SageMaker, the caller of this API must have the
	// iam:PassRole permission.
	//
	// This member is required.
	RoleArn *string

	// Specifies a limit to how long a model training job can run. When the job reaches
	// the time limit, Amazon SageMaker ends the training job. Use this API to cap
	// model training costs. To stop a job, Amazon SageMaker sends the algorithm the
	// SIGTERM signal, which delays job termination for 120 seconds. Algorithms can use
	// this 120-second window to save the model artifacts, so the results of training
	// are not lost.
	//
	// This member is required.
	StoppingCondition *types.StoppingCondition

	// The name of the training job. The name must be unique within an AWS Region in an
	// AWS account.
	//
	// This member is required.
	TrainingJobName *string

	// Contains information about the output location for managed spot training
	// checkpoint data.
	CheckpointConfig *types.CheckpointConfig

	// Configuration information for the debug hook parameters, collection
	// configuration, and storage paths.
	DebugHookConfig *types.DebugHookConfig

	// Configuration information for debugging rules.
	DebugRuleConfigurations []*types.DebugRuleConfiguration

	// To encrypt all communications between ML compute instances in distributed
	// training, choose True. Encryption provides greater security for distributed
	// training, but training might take longer. How long it takes depends on the
	// amount of communication between compute instances, especially if you use a deep
	// learning algorithm in distributed training. For more information, see Protect
	// Communications Between ML Compute Instances in a Distributed Training Job
	// (https://docs.aws.amazon.com/sagemaker/latest/dg/train-encrypt.html).
	EnableInterContainerTrafficEncryption *bool

	// To train models using managed spot training, choose True. Managed spot training
	// provides a fully managed and scalable infrastructure for training machine
	// learning models. this option is useful when training jobs can be interrupted and
	// when there is flexibility when the training job is run. The complete and
	// intermediate results of jobs are stored in an Amazon S3 bucket, and can be used
	// as a starting point to train models incrementally. Amazon SageMaker provides
	// metrics and logs in CloudWatch. They can be used to see when managed spot
	// training jobs are running, interrupted, resumed, or completed.
	EnableManagedSpotTraining *bool

	// Isolates the training container. No inbound or outbound network calls can be
	// made, except for calls between peers within a training cluster for distributed
	// training. If you enable network isolation for training jobs that are configured
	// to use a VPC, Amazon SageMaker downloads and uploads customer data and model
	// artifacts through the specified VPC, but the training container does not have
	// network access.
	EnableNetworkIsolation *bool

	// Associates a SageMaker job as a trial component with an experiment and trial.
	// Specified when you call the following APIs:
	//
	//     * CreateProcessingJob ()
	//
	//     *
	// CreateTrainingJob ()
	//
	//     * CreateTransformJob ()
	ExperimentConfig *types.ExperimentConfig

	// Algorithm-specific parameters that influence the quality of the model. You set
	// hyperparameters before you start the learning process. For a list of
	// hyperparameters for each training algorithm provided by Amazon SageMaker, see
	// Algorithms (https://docs.aws.amazon.com/sagemaker/latest/dg/algos.html). You can
	// specify a maximum of 100 hyperparameters. Each hyperparameter is a key-value
	// pair. Each key and value is limited to 256 characters, as specified by the
	// Length Constraint.
	HyperParameters map[string]*string

	// An array of Channel objects. Each channel is a named input source.
	// InputDataConfig
	//     describes the input data and its location. </p>
	// <p>Algorithms can accept input data from one or more channels. For example, an
	// algorithm might have two channels of input data, <code>training_data</code> and
	// <code>validation_data</code>. The configuration for each channel provides the
	// S3, EFS, or FSx location where the input data is stored. It also provides
	// information about the stored data: the MIME type, compression method, and
	// whether the data is wrapped in RecordIO format. </p> <p>Depending on the input
	// mode that the algorithm supports, Amazon SageMaker either copies input data
	// files from an S3 bucket to a local directory in the Docker container, or makes
	// it available as input streams. For example, if you specify an EFS location,
	// input data files will be made available as input streams. They do not need to be
	// downloaded.</p>
	InputDataConfig []*types.Channel

	// An array of key-value pairs. For more information, see Using Cost Allocation
	// Tags
	// (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-what)
	// in the AWS Billing and Cost Management User Guide.  </p>
	Tags []*types.Tag

	// Configuration of storage locations for TensorBoard output.
	TensorBoardOutputConfig *types.TensorBoardOutputConfig

	// A VpcConfig () object that specifies the VPC that you want your training job to
	// connect to. Control access to and from your training container by configuring
	// the VPC. For more information, see Protect Training Jobs by Using an Amazon
	// Virtual Private Cloud
	// (https://docs.aws.amazon.com/sagemaker/latest/dg/train-vpc.html).
	VpcConfig *types.VpcConfig
}

type CreateTrainingJobOutput struct {

	// The Amazon Resource Name (ARN) of the training job.
	//
	// This member is required.
	TrainingJobArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateTrainingJobMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateTrainingJob{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateTrainingJob{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateTrainingJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "CreateTrainingJob",
	}
}
