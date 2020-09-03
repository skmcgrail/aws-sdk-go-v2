// Code generated by smithy-go-codegen DO NOT EDIT.

package codedeploy

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the deployment groups for an application registered with the IAM user or
// AWS account.
func (c *Client) ListDeploymentGroups(ctx context.Context, params *ListDeploymentGroupsInput, optFns ...func(*Options)) (*ListDeploymentGroupsOutput, error) {
	stack := middleware.NewStack("ListDeploymentGroups", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListDeploymentGroupsMiddlewares(stack)
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
	addOpListDeploymentGroupsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListDeploymentGroups(options.Region), middleware.Before)

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
			OperationName: "ListDeploymentGroups",
			Err:           err,
		}
	}
	out := result.(*ListDeploymentGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a ListDeploymentGroups operation.
type ListDeploymentGroupsInput struct {
	// The name of an AWS CodeDeploy application associated with the IAM user or AWS
	// account.
	ApplicationName *string
	// An identifier returned from the previous list deployment groups call. It can be
	// used to return the next set of deployment groups in the list.
	NextToken *string
}

// Represents the output of a ListDeploymentGroups operation.
type ListDeploymentGroupsOutput struct {
	// The application name.
	ApplicationName *string
	// If a large amount of information is returned, an identifier is also returned. It
	// can be used in a subsequent list deployment groups call to return the next set
	// of deployment groups in the list.
	NextToken *string
	// A list of deployment group names.
	DeploymentGroups []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListDeploymentGroupsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListDeploymentGroups{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListDeploymentGroups{}, middleware.After)
}

func newServiceMetadataMiddleware_opListDeploymentGroups(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codedeploy",
		OperationName: "ListDeploymentGroups",
	}
}