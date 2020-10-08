// Code generated by smithy-go-codegen DO NOT EDIT.

package inspector

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/inspector/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an assessment template for the assessment target that is specified by
// the ARN of the assessment target. If the service-linked role
// (https://docs.aws.amazon.com/inspector/latest/userguide/inspector_slr.html)
// isn’t already registered, this action also creates and registers a
// service-linked role to grant Amazon Inspector access to AWS Services needed to
// perform security assessments.
func (c *Client) CreateAssessmentTemplate(ctx context.Context, params *CreateAssessmentTemplateInput, optFns ...func(*Options)) (*CreateAssessmentTemplateOutput, error) {
	stack := middleware.NewStack("CreateAssessmentTemplate", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateAssessmentTemplateMiddlewares(stack)
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
	addOpCreateAssessmentTemplateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAssessmentTemplate(options.Region), middleware.Before)
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
			OperationName: "CreateAssessmentTemplate",
			Err:           err,
		}
	}
	out := result.(*CreateAssessmentTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAssessmentTemplateInput struct {

	// The ARN that specifies the assessment target for which you want to create the
	// assessment template.
	//
	// This member is required.
	AssessmentTargetArn *string

	// The user-defined name that identifies the assessment template that you want to
	// create. You can create several assessment templates for an assessment target.
	// The names of the assessment templates that correspond to a particular assessment
	// target must be unique.
	//
	// This member is required.
	AssessmentTemplateName *string

	// The duration of the assessment run in seconds.
	//
	// This member is required.
	DurationInSeconds *int32

	// The ARNs that specify the rules packages that you want to attach to the
	// assessment template.
	//
	// This member is required.
	RulesPackageArns []*string

	// The user-defined attributes that are assigned to every finding that is generated
	// by the assessment run that uses this assessment template. An attribute is a key
	// and value pair (an Attribute () object). Within an assessment template, each key
	// must be unique.
	UserAttributesForFindings []*types.Attribute
}

type CreateAssessmentTemplateOutput struct {

	// The ARN that specifies the assessment template that is created.
	//
	// This member is required.
	AssessmentTemplateArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateAssessmentTemplateMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateAssessmentTemplate{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateAssessmentTemplate{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateAssessmentTemplate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "inspector",
		OperationName: "CreateAssessmentTemplate",
	}
}
