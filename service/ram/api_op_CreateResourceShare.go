// Code generated by smithy-go-codegen DO NOT EDIT.

package ram

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a resource share.
func (c *Client) CreateResourceShare(ctx context.Context, params *CreateResourceShareInput, optFns ...func(*Options)) (*CreateResourceShareOutput, error) {
	if params == nil {
		params = &CreateResourceShareInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateResourceShare", params, optFns, c.addOperationCreateResourceShareMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateResourceShareOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateResourceShareInput struct {

	// The name of the resource share.
	//
	// This member is required.
	Name *string

	// Indicates whether principals outside your AWS organization can be associated
	// with a resource share.
	AllowExternalPrincipals *bool

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientToken *string

	// The ARNs of the permissions to associate with the resource share. If you do not
	// specify an ARN for the permission, AWS RAM automatically attaches the default
	// version of the permission for each resource type.
	PermissionArns []string

	// The principals to associate with the resource share. The possible values are IDs
	// of AWS accounts, the ARN of an OU or organization from AWS Organizations.
	Principals []string

	// The Amazon Resource Names (ARN) of the resources to associate with the resource
	// share.
	ResourceArns []string

	// One or more tags.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateResourceShareOutput struct {

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientToken *string

	// Information about the resource share.
	ResourceShare *types.ResourceShare

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateResourceShareMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateResourceShare{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateResourceShare{}, middleware.After)
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
	if err = addOpCreateResourceShareValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateResourceShare(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateResourceShare(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ram",
		OperationName: "CreateResourceShare",
	}
}
