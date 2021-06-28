// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets room skill parameter details by room, skill, and parameter key ARN.
func (c *Client) GetRoomSkillParameter(ctx context.Context, params *GetRoomSkillParameterInput, optFns ...func(*Options)) (*GetRoomSkillParameterOutput, error) {
	if params == nil {
		params = &GetRoomSkillParameterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRoomSkillParameter", params, optFns, c.addOperationGetRoomSkillParameterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRoomSkillParameterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRoomSkillParameterInput struct {

	// The room skill parameter key for which to get details. Required.
	//
	// This member is required.
	ParameterKey *string

	// The ARN of the skill from which to get the room skill parameter details.
	// Required.
	//
	// This member is required.
	SkillId *string

	// The ARN of the room from which to get the room skill parameter details.
	RoomArn *string

	noSmithyDocumentSerde
}

type GetRoomSkillParameterOutput struct {

	// The details of the room skill parameter requested. Required.
	RoomSkillParameter *types.RoomSkillParameter

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetRoomSkillParameterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetRoomSkillParameter{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetRoomSkillParameter{}, middleware.After)
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
	if err = addOpGetRoomSkillParameterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetRoomSkillParameter(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetRoomSkillParameter(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "GetRoomSkillParameter",
	}
}
