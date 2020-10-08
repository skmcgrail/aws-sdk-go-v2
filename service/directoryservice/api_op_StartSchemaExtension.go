// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Applies a schema extension to a Microsoft AD directory.
func (c *Client) StartSchemaExtension(ctx context.Context, params *StartSchemaExtensionInput, optFns ...func(*Options)) (*StartSchemaExtensionOutput, error) {
	stack := middleware.NewStack("StartSchemaExtension", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpStartSchemaExtensionMiddlewares(stack)
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
	addOpStartSchemaExtensionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartSchemaExtension(options.Region), middleware.Before)
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
			OperationName: "StartSchemaExtension",
			Err:           err,
		}
	}
	out := result.(*StartSchemaExtensionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartSchemaExtensionInput struct {

	// If true, creates a snapshot of the directory before applying the schema
	// extension.
	//
	// This member is required.
	CreateSnapshotBeforeSchemaExtension *bool

	// A description of the schema extension.
	//
	// This member is required.
	Description *string

	// The identifier of the directory for which the schema extension will be applied
	// to.
	//
	// This member is required.
	DirectoryId *string

	// The LDIF file represented as a string. To construct the LdifContent string,
	// precede each line as it would be formatted in an ldif file with \n. See the
	// example request below for more details. The file size can be no larger than 1MB.
	//
	// This member is required.
	LdifContent *string
}

type StartSchemaExtensionOutput struct {

	// The identifier of the schema extension that will be applied.
	SchemaExtensionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpStartSchemaExtensionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpStartSchemaExtension{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartSchemaExtension{}, middleware.After)
}

func newServiceMetadataMiddleware_opStartSchemaExtension(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "StartSchemaExtension",
	}
}
