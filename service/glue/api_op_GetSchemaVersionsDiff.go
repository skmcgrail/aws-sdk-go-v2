// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Fetches the schema version difference in the specified difference type between
// two stored schema versions in the Schema Registry. This API allows you to
// compare two schema versions between two schema definitions under the same
// schema.
func (c *Client) GetSchemaVersionsDiff(ctx context.Context, params *GetSchemaVersionsDiffInput, optFns ...func(*Options)) (*GetSchemaVersionsDiffOutput, error) {
	if params == nil {
		params = &GetSchemaVersionsDiffInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSchemaVersionsDiff", params, optFns, c.addOperationGetSchemaVersionsDiffMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSchemaVersionsDiffOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSchemaVersionsDiffInput struct {

	// The first of the two schema versions to be compared.
	//
	// This member is required.
	FirstSchemaVersionNumber *types.SchemaVersionNumber

	// Refers to SYNTAX_DIFF, which is the currently supported diff type.
	//
	// This member is required.
	SchemaDiffType types.SchemaDiffType

	// This is a wrapper structure to contain schema identity fields. The structure
	// contains:
	//
	// * SchemaId$SchemaArn: The Amazon Resource Name (ARN) of the schema.
	// One of SchemaArn or SchemaName has to be provided.
	//
	// * SchemaId$SchemaName: The
	// name of the schema. One of SchemaArn or SchemaName has to be provided.
	//
	// This member is required.
	SchemaId *types.SchemaId

	// The second of the two schema versions to be compared.
	//
	// This member is required.
	SecondSchemaVersionNumber *types.SchemaVersionNumber

	noSmithyDocumentSerde
}

type GetSchemaVersionsDiffOutput struct {

	// The difference between schemas as a string in JsonPatch format.
	Diff *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSchemaVersionsDiffMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetSchemaVersionsDiff{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetSchemaVersionsDiff{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetSchemaVersionsDiffValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSchemaVersionsDiff(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetSchemaVersionsDiff(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "GetSchemaVersionsDiff",
	}
}
