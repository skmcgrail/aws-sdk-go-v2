// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Creates a new version of the bot based on the $LATEST version. If the $LATEST
// version of this resource hasn't changed since you created the last version,
// Amazon Lex doesn't create a new version. It returns the last created version.
// You can update only the $LATEST version of the bot. You can't update the
// numbered versions that you create with the CreateBotVersion operation. When you
// create the first version of a bot, Amazon Lex sets the version to 1. Subsequent
// versions increment by 1. For more information, see versioning-intro (). This
// operation requires permission for the lex:CreateBotVersion action.
func (c *Client) CreateBotVersion(ctx context.Context, params *CreateBotVersionInput, optFns ...func(*Options)) (*CreateBotVersionOutput, error) {
	stack := middleware.NewStack("CreateBotVersion", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateBotVersionMiddlewares(stack)
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
	addOpCreateBotVersionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateBotVersion(options.Region), middleware.Before)
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
			OperationName: "CreateBotVersion",
			Err:           err,
		}
	}
	out := result.(*CreateBotVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateBotVersionInput struct {

	// The name of the bot that you want to create a new version of. The name is case
	// sensitive.
	//
	// This member is required.
	Name *string

	// Identifies a specific revision of the $LATEST version of the bot. If you specify
	// a checksum and the $LATEST version of the bot has a different checksum, a
	// PreconditionFailedException exception is returned and Amazon Lex doesn't publish
	// a new version. If you don't specify a checksum, Amazon Lex publishes the $LATEST
	// version.
	Checksum *string
}

type CreateBotVersionOutput struct {

	// The message that Amazon Lex uses to abort a conversation. For more information,
	// see PutBot ().
	AbortStatement *types.Statement

	// Checksum identifying the version of the bot that was created.
	Checksum *string

	// For each Amazon Lex bot created with the Amazon Lex Model Building Service, you
	// must specify whether your use of Amazon Lex is related to a website, program, or
	// other application that is directed or targeted, in whole or in part, to children
	// under age 13 and subject to the Children's Online Privacy Protection Act (COPPA)
	// by specifying true or false in the childDirected field. By specifying true in
	// the childDirected field, you confirm that your use of Amazon Lex is related to a
	// website, program, or other application that is directed or targeted, in whole or
	// in part, to children under age 13 and subject to COPPA. By specifying false in
	// the childDirected field, you confirm that your use of Amazon Lex is not related
	// to a website, program, or other application that is directed or targeted, in
	// whole or in part, to children under age 13 and subject to COPPA. You may not
	// specify a default value for the childDirected field that does not accurately
	// reflect whether your use of Amazon Lex is related to a website, program, or
	// other application that is directed or targeted, in whole or in part, to children
	// under age 13 and subject to COPPA. If your use of Amazon Lex relates to a
	// website, program, or other application that is directed in whole or in part, to
	// children under age 13, you must obtain any required verifiable parental consent
	// under COPPA. For information regarding the use of Amazon Lex in connection with
	// websites, programs, or other applications that are directed or targeted, in
	// whole or in part, to children under age 13, see the Amazon Lex FAQ.
	// (https://aws.amazon.com/lex/faqs#data-security)
	ChildDirected *bool

	// The message that Amazon Lex uses when it doesn't understand the user's request.
	// For more information, see PutBot ().
	ClarificationPrompt *types.Prompt

	// The date when the bot version was created.
	CreatedDate *time.Time

	// A description of the bot.
	Description *string

	// Indicates whether utterances entered by the user should be sent to Amazon
	// Comprehend for sentiment analysis.
	DetectSentiment *bool

	// If status is FAILED, Amazon Lex provides the reason that it failed to build the
	// bot.
	FailureReason *string

	// The maximum time in seconds that Amazon Lex retains the data gathered in a
	// conversation. For more information, see PutBot ().
	IdleSessionTTLInSeconds *int32

	// An array of Intent objects. For more information, see PutBot ().
	Intents []*types.Intent

	// The date when the $LATEST version of this bot was updated.
	LastUpdatedDate *time.Time

	// Specifies the target locale for the bot.
	Locale types.Locale

	// The name of the bot.
	Name *string

	// When you send a request to create or update a bot, Amazon Lex sets the status
	// response element to BUILDING. After Amazon Lex builds the bot, it sets status to
	// READY. If Amazon Lex can't build the bot, it sets status to FAILED. Amazon Lex
	// returns the reason for the failure in the failureReason response element.
	Status types.Status

	// The version of the bot.
	Version *string

	// The Amazon Polly voice ID that Amazon Lex uses for voice interactions with the
	// user.
	VoiceId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateBotVersionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateBotVersion{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateBotVersion{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateBotVersion(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "CreateBotVersion",
	}
}
