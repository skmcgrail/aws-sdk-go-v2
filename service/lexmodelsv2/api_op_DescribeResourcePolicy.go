// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the resource policy and policy revision for a bot or bot alias.
func (c *Client) DescribeResourcePolicy(ctx context.Context, params *DescribeResourcePolicyInput, optFns ...func(*Options)) (*DescribeResourcePolicyOutput, error) {
	if params == nil {
		params = &DescribeResourcePolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeResourcePolicy", params, optFns, c.addOperationDescribeResourcePolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeResourcePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeResourcePolicyInput struct {

	// The Amazon Resource Name (ARN) of the bot or bot alias that the resource policy
	// is attached to.
	//
	// This member is required.
	ResourceArn *string

	noSmithyDocumentSerde
}

type DescribeResourcePolicyOutput struct {

	// The JSON structure that contains the resource policy. For more information about
	// the contents of a JSON policy document, see  IAM JSON policy reference
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies.html).
	Policy *string

	// The Amazon Resource Name (ARN) of the bot or bot alias that the resource policy
	// is attached to.
	ResourceArn *string

	// The current revision of the resource policy. Use the revision ID to make sure
	// that you are updating the most current version of a resource policy when you add
	// a policy statement to a resource, delete a resource, or update a resource.
	RevisionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeResourcePolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeResourcePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeResourcePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDescribeResourcePolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeResourcePolicy(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDescribeResourcePolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "DescribeResourcePolicy",
	}
}
