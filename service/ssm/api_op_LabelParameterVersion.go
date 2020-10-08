// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// A parameter label is a user-defined alias to help you manage different versions
// of a parameter. When you modify a parameter, Systems Manager automatically saves
// a new version and increments the version number by one. A label can help you
// remember the purpose of a parameter when there are multiple versions. Parameter
// labels have the following requirements and restrictions.
//
//     * A version of a
// parameter can have a maximum of 10 labels.
//
//     * You can't attach the same
// label to different versions of the same parameter. For example, if version 1 has
// the label Production, then you can't attach Production to version 2.
//
//     * You
// can move a label from one version of a parameter to another.
//
//     * You can't
// create a label when you create a new parameter. You must attach a label to a
// specific version of a parameter.
//
//     * You can't delete a parameter label. If
// you no longer want to use a parameter label, then you must move it to a
// different version of a parameter.
//
//     * A label can have a maximum of 100
// characters.
//
//     * Labels can contain letters (case sensitive), numbers, periods
// (.), hyphens (-), or underscores (_).
//
//     * Labels can't begin with a number,
// "aws," or "ssm" (not case sensitive). If a label fails to meet these
// requirements, then the label is not associated with a parameter and the system
// displays it in the list of InvalidLabels.
func (c *Client) LabelParameterVersion(ctx context.Context, params *LabelParameterVersionInput, optFns ...func(*Options)) (*LabelParameterVersionOutput, error) {
	stack := middleware.NewStack("LabelParameterVersion", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpLabelParameterVersionMiddlewares(stack)
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
	addOpLabelParameterVersionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opLabelParameterVersion(options.Region), middleware.Before)
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
			OperationName: "LabelParameterVersion",
			Err:           err,
		}
	}
	out := result.(*LabelParameterVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type LabelParameterVersionInput struct {

	// One or more labels to attach to the specified parameter version.
	//
	// This member is required.
	Labels []*string

	// The parameter name on which you want to attach one or more labels.
	//
	// This member is required.
	Name *string

	// The specific version of the parameter on which you want to attach one or more
	// labels. If no version is specified, the system attaches the label to the latest
	// version.
	ParameterVersion *int64
}

type LabelParameterVersionOutput struct {

	// The label does not meet the requirements. For information about parameter label
	// requirements, see Labeling parameters
	// (https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-paramstore-labels.html)
	// in the AWS Systems Manager User Guide.
	InvalidLabels []*string

	// The version of the parameter that has been labeled.
	ParameterVersion *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpLabelParameterVersionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpLabelParameterVersion{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpLabelParameterVersion{}, middleware.After)
}

func newServiceMetadataMiddleware_opLabelParameterVersion(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "LabelParameterVersion",
	}
}
