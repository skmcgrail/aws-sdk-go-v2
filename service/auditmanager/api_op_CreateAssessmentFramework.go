// Code generated by smithy-go-codegen DO NOT EDIT.

package auditmanager

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/auditmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a custom framework in AWS Audit Manager.
func (c *Client) CreateAssessmentFramework(ctx context.Context, params *CreateAssessmentFrameworkInput, optFns ...func(*Options)) (*CreateAssessmentFrameworkOutput, error) {
	if params == nil {
		params = &CreateAssessmentFrameworkInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAssessmentFramework", params, optFns, c.addOperationCreateAssessmentFrameworkMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAssessmentFrameworkOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAssessmentFrameworkInput struct {

	// The control sets to be associated with the framework.
	//
	// This member is required.
	ControlSets []types.CreateAssessmentFrameworkControlSet

	// The name of the new custom framework.
	//
	// This member is required.
	Name *string

	// The compliance type that the new custom framework supports, such as CIS or
	// HIPAA.
	ComplianceType *string

	// An optional description for the new custom framework.
	Description *string

	// The tags associated with the framework.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateAssessmentFrameworkOutput struct {

	// The name of the new framework returned by the CreateAssessmentFramework API.
	Framework *types.Framework

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAssessmentFrameworkMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateAssessmentFramework{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateAssessmentFramework{}, middleware.After)
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
	if err = addOpCreateAssessmentFrameworkValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAssessmentFramework(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateAssessmentFramework(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "auditmanager",
		OperationName: "CreateAssessmentFramework",
	}
}
