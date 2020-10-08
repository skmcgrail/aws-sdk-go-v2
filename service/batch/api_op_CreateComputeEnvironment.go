// Code generated by smithy-go-codegen DO NOT EDIT.

package batch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/batch/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an AWS Batch compute environment. You can create MANAGED or UNMANAGED
// compute environments. In a managed compute environment, AWS Batch manages the
// capacity and instance types of the compute resources within the environment.
// This is based on the compute resource specification that you define or the
// launch template
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-launch-templates.html)
// that you specify when you create the compute environment. You can choose to use
// Amazon EC2 On-Demand Instances or Spot Instances in your managed compute
// environment. You can optionally set a maximum price so that Spot Instances only
// launch when the Spot Instance price is below a specified percentage of the
// On-Demand price. Multi-node parallel jobs are not supported on Spot Instances.
// In an unmanaged compute environment, you can manage your own compute resources.
// This provides more compute resource configuration options, such as using a
// custom AMI, but you must ensure that your AMI meets the Amazon ECS container
// instance AMI specification. For more information, see Container Instance AMIs
// (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/container_instance_AMIs.html)
// in the Amazon Elastic Container Service Developer Guide. After you have created
// your unmanaged compute environment, you can use the DescribeComputeEnvironments
// () operation to find the Amazon ECS cluster that is associated with it. Then,
// manually launch your container instances into that Amazon ECS cluster. For more
// information, see Launching an Amazon ECS Container Instance
// (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/launch_container_instance.html)
// in the Amazon Elastic Container Service Developer Guide. AWS Batch does not
// upgrade the AMIs in a compute environment after it is created (for example, when
// a newer version of the Amazon ECS-optimized AMI is available). You are
// responsible for the management of the guest operating system (including updates
// and security patches) and any additional application software or utilities that
// you install on the compute resources. To use a new AMI for your AWS Batch
// jobs:
//
//     * Create a new compute environment with the new AMI.
//
//     * Add the
// compute environment to an existing job queue.
//
//     * Remove the old compute
// environment from your job queue.
//
//     * Delete the old compute environment.
func (c *Client) CreateComputeEnvironment(ctx context.Context, params *CreateComputeEnvironmentInput, optFns ...func(*Options)) (*CreateComputeEnvironmentOutput, error) {
	stack := middleware.NewStack("CreateComputeEnvironment", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateComputeEnvironmentMiddlewares(stack)
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
	addOpCreateComputeEnvironmentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateComputeEnvironment(options.Region), middleware.Before)
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
			OperationName: "CreateComputeEnvironment",
			Err:           err,
		}
	}
	out := result.(*CreateComputeEnvironmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateComputeEnvironmentInput struct {

	// The name for your compute environment. Up to 128 letters (uppercase and
	// lowercase), numbers, hyphens, and underscores are allowed.
	//
	// This member is required.
	ComputeEnvironmentName *string

	// The full Amazon Resource Name (ARN) of the IAM role that allows AWS Batch to
	// make calls to other AWS services on your behalf. If your specified role has a
	// path other than /, then you must either specify the full role ARN (this is
	// recommended) or prefix the role name with the path. Depending on how you created
	// your AWS Batch service role, its ARN may contain the service-role path prefix.
	// When you only specify the name of the service role, AWS Batch assumes that your
	// ARN does not use the service-role path prefix. Because of this, we recommend
	// that you specify the full ARN of your service role when you create compute
	// environments.
	//
	// This member is required.
	ServiceRole *string

	// The type of the compute environment. For more information, see Compute
	// Environments
	// (https://docs.aws.amazon.com/batch/latest/userguide/compute_environments.html)
	// in the AWS Batch User Guide.
	//
	// This member is required.
	Type types.CEType

	// Details of the compute resources managed by the compute environment. This
	// parameter is required for managed compute environments. For more information,
	// see Compute Environments
	// (https://docs.aws.amazon.com/batch/latest/userguide/compute_environments.html)
	// in the AWS Batch User Guide.
	ComputeResources *types.ComputeResource

	// The state of the compute environment. If the state is ENABLED, then the compute
	// environment accepts jobs from a queue and can scale out automatically based on
	// queues.
	State types.CEState
}

type CreateComputeEnvironmentOutput struct {

	// The Amazon Resource Name (ARN) of the compute environment.
	ComputeEnvironmentArn *string

	// The name of the compute environment.
	ComputeEnvironmentName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateComputeEnvironmentMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateComputeEnvironment{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateComputeEnvironment{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateComputeEnvironment(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "batch",
		OperationName: "CreateComputeEnvironment",
	}
}
