// Code generated by smithy-go-codegen DO NOT EDIT.

package glacier

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	glaciercust "github.com/aws/aws-sdk-go-v2/service/glacier/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/glacier/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This operation sets and then enacts a data retrieval policy in the region
// specified in the PUT request. You can set one policy per region for an AWS
// account. The policy is enacted within a few minutes of a successful PUT
// operation. The set policy operation does not affect retrieval jobs that were in
// progress before the policy was enacted. For more information about data
// retrieval policies, see Amazon Glacier Data Retrieval Policies
// (https://docs.aws.amazon.com/amazonglacier/latest/dev/data-retrieval-policy.html).
func (c *Client) SetDataRetrievalPolicy(ctx context.Context, params *SetDataRetrievalPolicyInput, optFns ...func(*Options)) (*SetDataRetrievalPolicyOutput, error) {
	stack := middleware.NewStack("SetDataRetrievalPolicy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpSetDataRetrievalPolicyMiddlewares(stack)
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
	addOpSetDataRetrievalPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSetDataRetrievalPolicy(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	glaciercust.AddTreeHashMiddleware(stack)
	glaciercust.AddGlacierAPIVersionMiddleware(stack, ServiceAPIVersion)
	glaciercust.AddDefaultAccountIDMiddleware(stack, setDefaultAccountID)

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
			OperationName: "SetDataRetrievalPolicy",
			Err:           err,
		}
	}
	out := result.(*SetDataRetrievalPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// SetDataRetrievalPolicy input.
type SetDataRetrievalPolicyInput struct {

	// The AccountId value is the AWS account ID. This value must match the AWS account
	// ID associated with the credentials used to sign the request. You can either
	// specify an AWS account ID or optionally a single '-' (hyphen), in which case
	// Amazon Glacier uses the AWS account ID associated with the credentials used to
	// sign the request. If you specify your account ID, do not include any hyphens
	// ('-') in the ID.
	//
	// This member is required.
	AccountId *string

	// The data retrieval policy in JSON format.
	Policy *types.DataRetrievalPolicy
}

type SetDataRetrievalPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpSetDataRetrievalPolicyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpSetDataRetrievalPolicy{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpSetDataRetrievalPolicy{}, middleware.After)
}

func newServiceMetadataMiddleware_opSetDataRetrievalPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glacier",
		OperationName: "SetDataRetrievalPolicy",
	}
}
