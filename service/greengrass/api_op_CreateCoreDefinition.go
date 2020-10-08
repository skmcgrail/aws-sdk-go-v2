// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrass

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a core definition. You may provide the initial version of the core
// definition now or use ''CreateCoreDefinitionVersion'' at a later time.
// Greengrass groups must each contain exactly one Greengrass core.
func (c *Client) CreateCoreDefinition(ctx context.Context, params *CreateCoreDefinitionInput, optFns ...func(*Options)) (*CreateCoreDefinitionOutput, error) {
	stack := middleware.NewStack("CreateCoreDefinition", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateCoreDefinitionMiddlewares(stack)
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
	addOpCreateCoreDefinitionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCoreDefinition(options.Region), middleware.Before)
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
			OperationName: "CreateCoreDefinition",
			Err:           err,
		}
	}
	out := result.(*CreateCoreDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Information needed to create a core definition.
type CreateCoreDefinitionInput struct {

	// A client token used to correlate requests and responses.
	AmznClientToken *string

	// Information about the initial version of the core definition.
	InitialVersion *types.CoreDefinitionVersion

	// The name of the core definition.
	Name *string

	// Tag(s) to add to the new resource.
	Tags map[string]*string
}

type CreateCoreDefinitionOutput struct {

	// The ARN of the definition.
	Arn *string

	// The time, in milliseconds since the epoch, when the definition was created.
	CreationTimestamp *string

	// The ID of the definition.
	Id *string

	// The time, in milliseconds since the epoch, when the definition was last updated.
	LastUpdatedTimestamp *string

	// The ID of the latest version associated with the definition.
	LatestVersion *string

	// The ARN of the latest version associated with the definition.
	LatestVersionArn *string

	// The name of the definition.
	Name *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateCoreDefinitionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateCoreDefinition{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateCoreDefinition{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateCoreDefinition(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "CreateCoreDefinition",
	}
}
