// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Allows a schema to be updated using JSON upload. Only available for development
// schemas. See JSON Schema Format
// (https://docs.aws.amazon.com/clouddirectory/latest/developerguide/schemas_jsonformat.html#schemas_json)
// for more information.
func (c *Client) PutSchemaFromJson(ctx context.Context, params *PutSchemaFromJsonInput, optFns ...func(*Options)) (*PutSchemaFromJsonOutput, error) {
	stack := middleware.NewStack("PutSchemaFromJson", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpPutSchemaFromJsonMiddlewares(stack)
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
	addOpPutSchemaFromJsonValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutSchemaFromJson(options.Region), middleware.Before)
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
			OperationName: "PutSchemaFromJson",
			Err:           err,
		}
	}
	out := result.(*PutSchemaFromJsonOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutSchemaFromJsonInput struct {
	// The ARN of the schema to update.
	SchemaArn *string
	// The replacement JSON schema.
	Document *string
}

type PutSchemaFromJsonOutput struct {
	// The ARN of the schema to update.
	Arn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpPutSchemaFromJsonMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpPutSchemaFromJson{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpPutSchemaFromJson{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutSchemaFromJson(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "PutSchemaFromJson",
	}
}