// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes a task definition. You can specify a family and revision to find
// information about a specific task definition, or you can simply specify the
// family to find the latest ACTIVE revision in that family. You can only describe
// INACTIVE task definitions while an active task or service references them.
func (c *Client) DescribeTaskDefinition(ctx context.Context, params *DescribeTaskDefinitionInput, optFns ...func(*Options)) (*DescribeTaskDefinitionOutput, error) {
	stack := middleware.NewStack("DescribeTaskDefinition", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeTaskDefinitionMiddlewares(stack)
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
	addOpDescribeTaskDefinitionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTaskDefinition(options.Region), middleware.Before)
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
			OperationName: "DescribeTaskDefinition",
			Err:           err,
		}
	}
	out := result.(*DescribeTaskDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTaskDefinitionInput struct {
	// The family for the latest ACTIVE revision, family and revision (family:revision)
	// for a specific revision in the family, or full Amazon Resource Name (ARN) of the
	// task definition to describe.
	TaskDefinition *string
	// Specifies whether to see the resource tags for the task definition. If TAGS is
	// specified, the tags are included in the response. If this field is omitted, tags
	// are not included in the response.
	Include []types.TaskDefinitionField
}

type DescribeTaskDefinitionOutput struct {
	// The full task definition description.
	TaskDefinition *types.TaskDefinition
	// The metadata that is applied to the task definition to help you categorize and
	// organize them. Each tag consists of a key and an optional value, both of which
	// you define. The following basic restrictions apply to tags:
	//
	//     * Maximum
	// number of tags per resource - 50
	//
	//     * For each resource, each tag key must be
	// unique, and each tag key can have only one value.
	//
	//     * Maximum key length -
	// 128 Unicode characters in UTF-8
	//
	//     * Maximum value length - 256 Unicode
	// characters in UTF-8
	//
	//     * If your tagging schema is used across multiple
	// services and resources, remember that other services may have restrictions on
	// allowed characters. Generally allowed characters are: letters, numbers, and
	// spaces representable in UTF-8, and the following characters: + - = . _ : / @.
	//
	//
	// * Tag keys and values are case-sensitive.
	//
	//     * Do not use aws:, AWS:, or any
	// upper or lowercase combination of such as a prefix for either keys or values as
	// it is reserved for AWS use. You cannot edit or delete tag keys or values with
	// this prefix. Tags with this prefix do not count against your tags per resource
	// limit.
	Tags []*types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeTaskDefinitionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeTaskDefinition{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeTaskDefinition{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeTaskDefinition(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecs",
		OperationName: "DescribeTaskDefinition",
	}
}