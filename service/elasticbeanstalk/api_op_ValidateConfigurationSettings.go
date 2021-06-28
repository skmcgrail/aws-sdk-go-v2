// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Takes a set of configuration settings and either a configuration template or
// environment, and determines whether those values are valid. This action returns
// a list of messages indicating any errors or warnings associated with the
// selection of option values.
func (c *Client) ValidateConfigurationSettings(ctx context.Context, params *ValidateConfigurationSettingsInput, optFns ...func(*Options)) (*ValidateConfigurationSettingsOutput, error) {
	if params == nil {
		params = &ValidateConfigurationSettingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ValidateConfigurationSettings", params, optFns, c.addOperationValidateConfigurationSettingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ValidateConfigurationSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A list of validation messages for a specified configuration template.
type ValidateConfigurationSettingsInput struct {

	// The name of the application that the configuration template or environment
	// belongs to.
	//
	// This member is required.
	ApplicationName *string

	// A list of the options and desired values to evaluate.
	//
	// This member is required.
	OptionSettings []types.ConfigurationOptionSetting

	// The name of the environment to validate the settings against. Condition: You
	// cannot specify both this and a configuration template name.
	EnvironmentName *string

	// The name of the configuration template to validate the settings against.
	// Condition: You cannot specify both this and an environment name.
	TemplateName *string

	noSmithyDocumentSerde
}

// Provides a list of validation messages.
type ValidateConfigurationSettingsOutput struct {

	// A list of ValidationMessage.
	Messages []types.ValidationMessage

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationValidateConfigurationSettingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpValidateConfigurationSettings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpValidateConfigurationSettings{}, middleware.After)
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
	if err = addOpValidateConfigurationSettingsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opValidateConfigurationSettings(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opValidateConfigurationSettings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticbeanstalk",
		OperationName: "ValidateConfigurationSettings",
	}
}
