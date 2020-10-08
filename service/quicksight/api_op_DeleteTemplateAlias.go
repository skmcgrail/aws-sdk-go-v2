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

// Deletes the item that the specified template alias points to. If you provide a
// specific alias, you delete the version of the template that the alias points to.
func (c *Client) DeleteTemplateAlias(ctx context.Context, params *DeleteTemplateAliasInput, optFns ...func(*Options)) (*DeleteTemplateAliasOutput, error) {
	stack := middleware.NewStack("DeleteTemplateAlias", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteTemplateAliasMiddlewares(stack)
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
	addOpDeleteTemplateAliasValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteTemplateAlias(options.Region), middleware.Before)
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
			OperationName: "DeleteTemplateAlias",
			Err:           err,
		}
	}
	out := result.(*DeleteTemplateAliasOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteTemplateAliasInput struct {

	// The name for the template alias. To delete a specific alias, you delete the
	// version that the alias points to. You can specify the alias name, or specify the
	// latest version of the template by providing the keyword $LATEST in the AliasName
	// parameter.
	//
	// This member is required.
	AliasName *string

	// The ID of the AWS account that contains the item to delete.
	//
	// This member is required.
	AwsAccountId *string

	// The ID for the template that the specified alias is for.
	//
	// This member is required.
	TemplateId *string
}

type DeleteTemplateAliasOutput struct {

	// The name for the template alias.
	AliasName *string

	// The Amazon Resource Name (ARN) of the template you want to delete.
	Arn *string

	// The AWS request ID for this operation.
	RequestId *string

	// An ID for the template associated with the deletion.
	TemplateId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteTemplateAliasMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteTemplateAlias{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteTemplateAlias{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteTemplateAlias(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "DeleteTemplateAlias",
	}
}
