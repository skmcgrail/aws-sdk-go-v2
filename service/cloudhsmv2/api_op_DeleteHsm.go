// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudhsmv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specified HSM. To specify an HSM, you can use its identifier (ID),
// the IP address of the HSM's elastic network interface (ENI), or the ID of the
// HSM's ENI. You need to specify only one of these values. To find these values,
// use DescribeClusters ().
func (c *Client) DeleteHsm(ctx context.Context, params *DeleteHsmInput, optFns ...func(*Options)) (*DeleteHsmOutput, error) {
	stack := middleware.NewStack("DeleteHsm", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteHsmMiddlewares(stack)
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
	addOpDeleteHsmValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteHsm(options.Region), middleware.Before)
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
			OperationName: "DeleteHsm",
			Err:           err,
		}
	}
	out := result.(*DeleteHsmOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteHsmInput struct {

	// The identifier (ID) of the cluster that contains the HSM that you are deleting.
	//
	// This member is required.
	ClusterId *string

	// The identifier (ID) of the elastic network interface (ENI) of the HSM that you
	// are deleting.
	EniId *string

	// The IP address of the elastic network interface (ENI) of the HSM that you are
	// deleting.
	EniIp *string

	// The identifier (ID) of the HSM that you are deleting.
	HsmId *string
}

type DeleteHsmOutput struct {

	// The identifier (ID) of the HSM that was deleted.
	HsmId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteHsmMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteHsm{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteHsm{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteHsm(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudhsm",
		OperationName: "DeleteHsm",
	}
}
