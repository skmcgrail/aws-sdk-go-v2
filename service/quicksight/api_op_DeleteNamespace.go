// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a namespace and the users and groups that are associated with the
// namespace. This is an asynchronous process. Assets including dashboards,
// analyses, datasets and data sources are not deleted. To delete these assets, you
// use the APIs for the relevant asset.
func (c *Client) DeleteNamespace(ctx context.Context, params *DeleteNamespaceInput, optFns ...func(*Options)) (*DeleteNamespaceOutput, error) {
	stack := middleware.NewStack("DeleteNamespace", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteNamespaceMiddlewares(stack)
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
	addOpDeleteNamespaceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteNamespace(options.Region), middleware.Before)
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
			OperationName: "DeleteNamespace",
			Err:           err,
		}
	}
	out := result.(*DeleteNamespaceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteNamespaceInput struct {

	// The ID for the AWS account that you want to delete the QuickSight namespace
	// from.
	//
	// This member is required.
	AwsAccountId *string

	// The namespace that you want to delete.
	//
	// This member is required.
	Namespace *string
}

type DeleteNamespaceOutput struct {

	// The AWS request ID for this operation.
	RequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteNamespaceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteNamespace{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteNamespace{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteNamespace(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "DeleteNamespace",
	}
}
