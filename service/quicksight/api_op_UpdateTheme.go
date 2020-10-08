// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates a theme.
func (c *Client) UpdateTheme(ctx context.Context, params *UpdateThemeInput, optFns ...func(*Options)) (*UpdateThemeOutput, error) {
	stack := middleware.NewStack("UpdateTheme", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateThemeMiddlewares(stack)
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
	addOpUpdateThemeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateTheme(options.Region), middleware.Before)
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
			OperationName: "UpdateTheme",
			Err:           err,
		}
	}
	out := result.(*UpdateThemeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateThemeInput struct {

	// The ID of the AWS account that contains the theme that you're updating.
	//
	// This member is required.
	AwsAccountId *string

	// The theme ID, defined by Amazon QuickSight, that a custom theme inherits from.
	// All themes initially inherit from a default QuickSight theme.
	//
	// This member is required.
	BaseThemeId *string

	// The ID for the theme.
	//
	// This member is required.
	ThemeId *string

	// The theme configuration, which contains the theme display properties.
	Configuration *types.ThemeConfiguration

	// The name for the theme.
	Name *string

	// A description of the theme version that you're updating Every time that you call
	// UpdateTheme, you create a new version of the theme. Each version of the theme
	// maintains a description of the version in VersionDescription.
	VersionDescription *string
}

type UpdateThemeOutput struct {

	// The Amazon Resource Name (ARN) for the theme.
	Arn *string

	// The creation status of the theme.
	CreationStatus types.ResourceStatus

	// The AWS request ID for this operation.
	RequestId *string

	// The ID for the theme.
	ThemeId *string

	// The Amazon Resource Name (ARN) for the new version of the theme.
	VersionArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateThemeMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateTheme{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateTheme{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateTheme(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "UpdateTheme",
	}
}
