// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes an Input Security Group
func (c *Client) DeleteInputSecurityGroup(ctx context.Context, params *DeleteInputSecurityGroupInput, optFns ...func(*Options)) (*DeleteInputSecurityGroupOutput, error) {
	stack := middleware.NewStack("DeleteInputSecurityGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteInputSecurityGroupMiddlewares(stack)
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
	addOpDeleteInputSecurityGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteInputSecurityGroup(options.Region), middleware.Before)
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
			OperationName: "DeleteInputSecurityGroup",
			Err:           err,
		}
	}
	out := result.(*DeleteInputSecurityGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Placeholder documentation for DeleteInputSecurityGroupRequest
type DeleteInputSecurityGroupInput struct {

	// The Input Security Group to delete
	//
	// This member is required.
	InputSecurityGroupId *string
}

// Placeholder documentation for DeleteInputSecurityGroupResponse
type DeleteInputSecurityGroupOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteInputSecurityGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteInputSecurityGroup{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteInputSecurityGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteInputSecurityGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "medialive",
		OperationName: "DeleteInputSecurityGroup",
	}
}
