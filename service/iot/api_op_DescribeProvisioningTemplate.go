// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns information about a fleet provisioning template.
func (c *Client) DescribeProvisioningTemplate(ctx context.Context, params *DescribeProvisioningTemplateInput, optFns ...func(*Options)) (*DescribeProvisioningTemplateOutput, error) {
	stack := middleware.NewStack("DescribeProvisioningTemplate", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeProvisioningTemplateMiddlewares(stack)
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
	addOpDescribeProvisioningTemplateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeProvisioningTemplate(options.Region), middleware.Before)
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
			OperationName: "DescribeProvisioningTemplate",
			Err:           err,
		}
	}
	out := result.(*DescribeProvisioningTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeProvisioningTemplateInput struct {
	// The name of the fleet provisioning template.
	TemplateName *string
}

type DescribeProvisioningTemplateOutput struct {
	// The description of the fleet provisioning template.
	Description *string
	// The date when the fleet provisioning template was last modified.
	LastModifiedDate *time.Time
	// The ARN of the fleet provisioning template.
	TemplateArn *string
	// The date when the fleet provisioning template was created.
	CreationDate *time.Time
	// The ARN of the role associated with the provisioning template. This IoT role
	// grants permission to provision a device.
	ProvisioningRoleArn *string
	// True if the fleet provisioning template is enabled, otherwise false.
	Enabled *bool
	// The name of the fleet provisioning template.
	TemplateName *string
	// Gets information about a pre-provisioned hook.
	PreProvisioningHook *types.ProvisioningHook
	// The JSON formatted contents of the fleet provisioning template.
	TemplateBody *string
	// The default fleet template version ID.
	DefaultVersionId *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeProvisioningTemplateMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeProvisioningTemplate{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeProvisioningTemplate{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeProvisioningTemplate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "DescribeProvisioningTemplate",
	}
}