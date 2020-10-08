// Code generated by smithy-go-codegen DO NOT EDIT.

package workdocs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes all the permissions from the specified resource.
func (c *Client) RemoveAllResourcePermissions(ctx context.Context, params *RemoveAllResourcePermissionsInput, optFns ...func(*Options)) (*RemoveAllResourcePermissionsOutput, error) {
	stack := middleware.NewStack("RemoveAllResourcePermissions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpRemoveAllResourcePermissionsMiddlewares(stack)
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
	addOpRemoveAllResourcePermissionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRemoveAllResourcePermissions(options.Region), middleware.Before)
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
			OperationName: "RemoveAllResourcePermissions",
			Err:           err,
		}
	}
	out := result.(*RemoveAllResourcePermissionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RemoveAllResourcePermissionsInput struct {

	// The ID of the resource.
	//
	// This member is required.
	ResourceId *string

	// Amazon WorkDocs authentication token. Not required when using AWS administrator
	// credentials to access the API.
	AuthenticationToken *string
}

type RemoveAllResourcePermissionsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpRemoveAllResourcePermissionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpRemoveAllResourcePermissions{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpRemoveAllResourcePermissions{}, middleware.After)
}

func newServiceMetadataMiddleware_opRemoveAllResourcePermissions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workdocs",
		OperationName: "RemoveAllResourcePermissions",
	}
}
