// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Describes a bulk thing provisioning task.
func (c *Client) DescribeThingRegistrationTask(ctx context.Context, params *DescribeThingRegistrationTaskInput, optFns ...func(*Options)) (*DescribeThingRegistrationTaskOutput, error) {
	stack := middleware.NewStack("DescribeThingRegistrationTask", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeThingRegistrationTaskMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpDescribeThingRegistrationTaskValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeThingRegistrationTask(options.Region), middleware.Before)
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
			OperationName: "DescribeThingRegistrationTask",
			Err:           err,
		}
	}
	out := result.(*DescribeThingRegistrationTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeThingRegistrationTaskInput struct {
	// The task ID.
	TaskId *string
}

type DescribeThingRegistrationTaskOutput struct {
	// The input file key.
	InputFileKey *string
	// The date when the task was last modified.
	LastModifiedDate *time.Time
	// The status of the bulk thing provisioning task.
	Status types.Status
	// The role ARN that grants access to the input file bucket.
	RoleArn *string
	// The number of things successfully provisioned.
	SuccessCount *int32
	// The S3 bucket that contains the input file.
	InputFileBucket *string
	// The task's template.
	TemplateBody *string
	// The message.
	Message *string
	// The number of things that failed to be provisioned.
	FailureCount *int32
	// The progress of the bulk provisioning task expressed as a percentage.
	PercentageProgress *int32
	// The task ID.
	TaskId *string
	// The task creation date.
	CreationDate *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeThingRegistrationTaskMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeThingRegistrationTask{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeThingRegistrationTask{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeThingRegistrationTask(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "DescribeThingRegistrationTask",
	}
}