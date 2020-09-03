// Code generated by smithy-go-codegen DO NOT EDIT.

package mobile

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mobile/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Delets a project in AWS Mobile Hub.
func (c *Client) DeleteProject(ctx context.Context, params *DeleteProjectInput, optFns ...func(*Options)) (*DeleteProjectOutput, error) {
	stack := middleware.NewStack("DeleteProject", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteProjectMiddlewares(stack)
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
	addOpDeleteProjectValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteProject(options.Region), middleware.Before)

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
			OperationName: "DeleteProject",
			Err:           err,
		}
	}
	out := result.(*DeleteProjectOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request structure used to request a project be deleted.
type DeleteProjectInput struct {
	// Unique project identifier.
	ProjectId *string
}

// Result structure used in response to request to delete a project.
type DeleteProjectOutput struct {
	// Resources which were deleted.
	DeletedResources []*types.Resource
	// Resources which were not deleted, due to a risk of losing potentially important
	// data or files.
	OrphanedResources []*types.Resource

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteProjectMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteProject{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteProject{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteProject(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "awsmobilehubservice",
		OperationName: "DeleteProject",
	}
}