// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates a notebook instance lifecycle configuration created with the
// CreateNotebookInstanceLifecycleConfig () API.
func (c *Client) UpdateNotebookInstanceLifecycleConfig(ctx context.Context, params *UpdateNotebookInstanceLifecycleConfigInput, optFns ...func(*Options)) (*UpdateNotebookInstanceLifecycleConfigOutput, error) {
	stack := middleware.NewStack("UpdateNotebookInstanceLifecycleConfig", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateNotebookInstanceLifecycleConfigMiddlewares(stack)
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
	addOpUpdateNotebookInstanceLifecycleConfigValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateNotebookInstanceLifecycleConfig(options.Region), middleware.Before)
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
			OperationName: "UpdateNotebookInstanceLifecycleConfig",
			Err:           err,
		}
	}
	out := result.(*UpdateNotebookInstanceLifecycleConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateNotebookInstanceLifecycleConfigInput struct {
	// The shell script that runs every time you start a notebook instance, including
	// when you create the notebook instance. The shell script must be a base64-encoded
	// string.
	OnStart []*types.NotebookInstanceLifecycleHook
	// The shell script that runs only once, when you create a notebook instance. The
	// shell script must be a base64-encoded string.
	OnCreate []*types.NotebookInstanceLifecycleHook
	// The name of the lifecycle configuration.
	NotebookInstanceLifecycleConfigName *string
}

type UpdateNotebookInstanceLifecycleConfigOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateNotebookInstanceLifecycleConfigMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateNotebookInstanceLifecycleConfig{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateNotebookInstanceLifecycleConfig{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateNotebookInstanceLifecycleConfig(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "UpdateNotebookInstanceLifecycleConfig",
	}
}