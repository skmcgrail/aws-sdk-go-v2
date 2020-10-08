// Code generated by smithy-go-codegen DO NOT EDIT.

package detective

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new behavior graph for the calling account, and sets that account as
// the master account. This operation is called by the account that is enabling
// Detective. Before you try to enable Detective, make sure that your account has
// been enrolled in Amazon GuardDuty for at least 48 hours. If you do not meet this
// requirement, you cannot enable Detective. If you do meet the GuardDuty
// prerequisite, then when you make the request to enable Detective, it checks
// whether your data volume is within the Detective quota. If it exceeds the quota,
// then you cannot enable Detective. The operation also enables Detective for the
// calling account in the currently selected Region. It returns the ARN of the new
// behavior graph. CreateGraph triggers a process to create the corresponding data
// tables for the new behavior graph. An account can only be the master account for
// one behavior graph within a Region. If the same account calls CreateGraph with
// the same master account, it always returns the same behavior graph ARN. It does
// not create a new behavior graph.
func (c *Client) CreateGraph(ctx context.Context, params *CreateGraphInput, optFns ...func(*Options)) (*CreateGraphOutput, error) {
	stack := middleware.NewStack("CreateGraph", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateGraphMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateGraph(options.Region), middleware.Before)
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
			OperationName: "CreateGraph",
			Err:           err,
		}
	}
	out := result.(*CreateGraphOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateGraphInput struct {
}

type CreateGraphOutput struct {

	// The ARN of the new behavior graph.
	GraphArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateGraphMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateGraph{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateGraph{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateGraph(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "detective",
		OperationName: "CreateGraph",
	}
}
