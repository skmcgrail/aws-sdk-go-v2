// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Associates a skill with a skill group.
func (c *Client) AssociateSkillWithSkillGroup(ctx context.Context, params *AssociateSkillWithSkillGroupInput, optFns ...func(*Options)) (*AssociateSkillWithSkillGroupOutput, error) {
	stack := middleware.NewStack("AssociateSkillWithSkillGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAssociateSkillWithSkillGroupMiddlewares(stack)
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
	addOpAssociateSkillWithSkillGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateSkillWithSkillGroup(options.Region), middleware.Before)
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
			OperationName: "AssociateSkillWithSkillGroup",
			Err:           err,
		}
	}
	out := result.(*AssociateSkillWithSkillGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateSkillWithSkillGroupInput struct {

	// The unique identifier of the skill.
	//
	// This member is required.
	SkillId *string

	// The ARN of the skill group to associate the skill to. Required.
	SkillGroupArn *string
}

type AssociateSkillWithSkillGroupOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAssociateSkillWithSkillGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAssociateSkillWithSkillGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociateSkillWithSkillGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opAssociateSkillWithSkillGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "AssociateSkillWithSkillGroup",
	}
}
