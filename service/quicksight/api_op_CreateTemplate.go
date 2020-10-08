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

// Creates a template from an existing QuickSight analysis or template. You can use
// the resulting template to create a dashboard. A template is an entity in
// QuickSight that encapsulates the metadata required to create an analysis and
// that you can use to create s dashboard. A template adds a layer of abstraction
// by using placeholders to replace the dataset associated with the analysis. You
// can use templates to create dashboards by replacing dataset placeholders with
// datasets that follow the same schema that was used to create the source analysis
// and template.
func (c *Client) CreateTemplate(ctx context.Context, params *CreateTemplateInput, optFns ...func(*Options)) (*CreateTemplateOutput, error) {
	stack := middleware.NewStack("CreateTemplate", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateTemplateMiddlewares(stack)
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
	addOpCreateTemplateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTemplate(options.Region), middleware.Before)
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
			OperationName: "CreateTemplate",
			Err:           err,
		}
	}
	out := result.(*CreateTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTemplateInput struct {

	// The ID for the AWS account that the group is in. Currently, you use the ID for
	// the AWS account that contains your Amazon QuickSight account.
	//
	// This member is required.
	AwsAccountId *string

	// The entity that you are using as a source when you create the template. In
	// SourceEntity, you specify the type of object you're using as source:
	// SourceTemplate for a template or SourceAnalysis for an analysis. Both of these
	// require an Amazon Resource Name (ARN). For SourceTemplate, specify the ARN of
	// the source template. For SourceAnalysis, specify the ARN of the source analysis.
	// The SourceTemplate ARN can contain any AWS Account and any QuickSight-supported
	// AWS Region. Use the DataSetReferences entity within SourceTemplate or
	// SourceAnalysis to list the replacement datasets for the placeholders listed in
	// the original. The schema in each dataset must match its placeholder.
	//
	// This member is required.
	SourceEntity *types.TemplateSourceEntity

	// An ID for the template that you want to create. This template is unique per AWS
	// Region in each AWS account.
	//
	// This member is required.
	TemplateId *string

	// A display name for the template.
	Name *string

	// A list of resource permissions to be set on the template.
	Permissions []*types.ResourcePermission

	// Contains a map of the key-value pairs for the resource tag or tags assigned to
	// the resource.
	Tags []*types.Tag

	// A description of the current template version being created. This API operation
	// creates the first version of the template. Every time UpdateTemplate is called,
	// a new version is created. Each version of the template maintains a description
	// of the version in the VersionDescription field.
	VersionDescription *string
}

type CreateTemplateOutput struct {

	// The ARN for the template.
	Arn *string

	// The template creation status.
	CreationStatus types.ResourceStatus

	// The AWS request ID for this operation.
	RequestId *string

	// The ID of the template.
	TemplateId *string

	// The ARN for the template, including the version information of the first
	// version.
	VersionArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateTemplateMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateTemplate{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateTemplate{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateTemplate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "CreateTemplate",
	}
}
