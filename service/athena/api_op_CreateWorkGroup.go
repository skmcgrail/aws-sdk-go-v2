// Code generated by smithy-go-codegen DO NOT EDIT.

package athena

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a workgroup with the specified name.
func (c *Client) CreateWorkGroup(ctx context.Context, params *CreateWorkGroupInput, optFns ...func(*Options)) (*CreateWorkGroupOutput, error) {
	stack := middleware.NewStack("CreateWorkGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateWorkGroupMiddlewares(stack)
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
	addOpCreateWorkGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateWorkGroup(options.Region), middleware.Before)
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
			OperationName: "CreateWorkGroup",
			Err:           err,
		}
	}
	out := result.(*CreateWorkGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateWorkGroupInput struct {

	// The workgroup name.
	//
	// This member is required.
	Name *string

	// The configuration for the workgroup, which includes the location in Amazon S3
	// where query results are stored, the encryption configuration, if any, used for
	// encrypting query results, whether the Amazon CloudWatch Metrics are enabled for
	// the workgroup, the limit for the amount of bytes scanned (cutoff) per query, if
	// it is specified, and whether workgroup's settings (specified with
	// EnforceWorkGroupConfiguration) in the WorkGroupConfiguration override
	// client-side settings. See WorkGroupConfiguration$EnforceWorkGroupConfiguration
	// ().
	Configuration *types.WorkGroupConfiguration

	// The workgroup description.
	Description *string

	// A list of comma separated tags to add to the workgroup that is created.
	Tags []*types.Tag
}

type CreateWorkGroupOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateWorkGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateWorkGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateWorkGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateWorkGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "athena",
		OperationName: "CreateWorkGroup",
	}
}
