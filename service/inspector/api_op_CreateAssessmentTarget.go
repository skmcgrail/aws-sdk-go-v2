// Code generated by smithy-go-codegen DO NOT EDIT.

package inspector

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new assessment target using the ARN of the resource group that is
// generated by CreateResourceGroup (). If resourceGroupArn is not specified, all
// EC2 instances in the current AWS account and region are included in the
// assessment target. If the service-linked role
// (https://docs.aws.amazon.com/inspector/latest/userguide/inspector_slr.html)
// isn’t already registered, this action also creates and registers a
// service-linked role to grant Amazon Inspector access to AWS Services needed to
// perform security assessments. You can create up to 50 assessment targets per AWS
// account. You can run up to 500 concurrent agents per AWS account. For more
// information, see  Amazon Inspector Assessment Targets
// (https://docs.aws.amazon.com/inspector/latest/userguide/inspector_applications.html).
func (c *Client) CreateAssessmentTarget(ctx context.Context, params *CreateAssessmentTargetInput, optFns ...func(*Options)) (*CreateAssessmentTargetOutput, error) {
	stack := middleware.NewStack("CreateAssessmentTarget", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateAssessmentTargetMiddlewares(stack)
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
	addOpCreateAssessmentTargetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAssessmentTarget(options.Region), middleware.Before)
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
			OperationName: "CreateAssessmentTarget",
			Err:           err,
		}
	}
	out := result.(*CreateAssessmentTargetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAssessmentTargetInput struct {

	// The user-defined name that identifies the assessment target that you want to
	// create. The name must be unique within the AWS account.
	//
	// This member is required.
	AssessmentTargetName *string

	// The ARN that specifies the resource group that is used to create the assessment
	// target. If resourceGroupArn is not specified, all EC2 instances in the current
	// AWS account and region are included in the assessment target.
	ResourceGroupArn *string
}

type CreateAssessmentTargetOutput struct {

	// The ARN that specifies the assessment target that is created.
	//
	// This member is required.
	AssessmentTargetArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateAssessmentTargetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateAssessmentTarget{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateAssessmentTarget{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateAssessmentTarget(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "inspector",
		OperationName: "CreateAssessmentTarget",
	}
}
