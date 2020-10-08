// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about the configured contact methods. Specify a protocol in
// your request to return information about a specific contact method. A contact
// method is used to send you notifications about your Amazon Lightsail resources.
// You can add one email address and one mobile phone number contact method in each
// AWS Region. However, SMS text messaging is not supported in some AWS Regions,
// and SMS text messages cannot be sent to some countries/regions. For more
// information, see Notifications in Amazon Lightsail
// (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-notifications).
func (c *Client) GetContactMethods(ctx context.Context, params *GetContactMethodsInput, optFns ...func(*Options)) (*GetContactMethodsOutput, error) {
	stack := middleware.NewStack("GetContactMethods", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetContactMethodsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetContactMethods(options.Region), middleware.Before)
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
			OperationName: "GetContactMethods",
			Err:           err,
		}
	}
	out := result.(*GetContactMethodsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetContactMethodsInput struct {

	// The protocols used to send notifications, such as Email, or SMS (text
	// messaging). Specify a protocol in your request to return information about a
	// specific contact method protocol.
	Protocols []types.ContactProtocol
}

type GetContactMethodsOutput struct {

	// An array of objects that describe the contact methods.
	ContactMethods []*types.ContactMethod

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetContactMethodsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetContactMethods{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetContactMethods{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetContactMethods(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "GetContactMethods",
	}
}
