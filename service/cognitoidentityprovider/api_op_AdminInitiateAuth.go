// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Initiates the authentication flow, as an administrator. Calling this action
// requires developer credentials.
func (c *Client) AdminInitiateAuth(ctx context.Context, params *AdminInitiateAuthInput, optFns ...func(*Options)) (*AdminInitiateAuthOutput, error) {
	stack := middleware.NewStack("AdminInitiateAuth", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAdminInitiateAuthMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpAdminInitiateAuthValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAdminInitiateAuth(options.Region), middleware.Before)
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
			OperationName: "AdminInitiateAuth",
			Err:           err,
		}
	}
	out := result.(*AdminInitiateAuthOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Initiates the authorization request, as an administrator.
type AdminInitiateAuthInput struct {
	// The authentication flow for this call to execute. The API action will depend on
	// this value. For example:
	//
	//     * REFRESH_TOKEN_AUTH will take in a valid refresh
	// token and return new tokens.
	//
	//     * USER_SRP_AUTH will take in USERNAME and
	// SRP_A and return the SRP variables to be used for next challenge execution.
	//
	//
	// * USER_PASSWORD_AUTH will take in USERNAME and PASSWORD and return the next
	// challenge or tokens.
	//
	// Valid values include:
	//
	//     * USER_SRP_AUTH: Authentication
	// flow for the Secure Remote Password (SRP) protocol.
	//
	//     *
	// REFRESH_TOKEN_AUTH/REFRESH_TOKEN: Authentication flow for refreshing the access
	// token and ID token by supplying a valid refresh token.
	//
	//     * CUSTOM_AUTH:
	// Custom authentication flow.
	//
	//     * ADMIN_NO_SRP_AUTH: Non-SRP authentication
	// flow; you can pass in the USERNAME and PASSWORD directly if the flow is enabled
	// for calling the app client.
	//
	//     * USER_PASSWORD_AUTH: Non-SRP authentication
	// flow; USERNAME and PASSWORD are passed directly. If a user migration Lambda
	// trigger is set, this flow will invoke the user migration Lambda if the USERNAME
	// is not found in the user pool.
	//
	//     * ADMIN_USER_PASSWORD_AUTH: Admin-based user
	// password authentication. This replaces the ADMIN_NO_SRP_AUTH authentication
	// flow. In this flow, Cognito receives the password in the request instead of
	// using the SRP process to verify passwords.
	AuthFlow types.AuthFlowType
	// The ID of the Amazon Cognito user pool.
	UserPoolId *string
	// A map of custom key-value pairs that you can provide as input for certain custom
	// workflows that this action triggers. You create custom workflows by assigning
	// AWS Lambda functions to user pool triggers. When you use the AdminInitiateAuth
	// API action, Amazon Cognito invokes the AWS Lambda functions that are specified
	// for various triggers. The ClientMetadata value is passed as input to the
	// functions for only the following triggers:
	//
	//     * Pre signup
	//
	//     * Pre
	// authentication
	//
	//     * User migration
	//
	//     <p>When Amazon Cognito invokes the
	// functions for these triggers, it passes a JSON payload, which the function
	// receives as input. This payload contains a <code>validationData</code>
	// attribute, which provides the data that you assigned to the ClientMetadata
	// parameter in your AdminInitiateAuth request. In your function code in AWS
	// Lambda, you can process the <code>validationData</code> value to enhance your
	// workflow for your specific needs.</p> <p>When you use the AdminInitiateAuth API
	// action, Amazon Cognito also invokes the functions for the following triggers,
	// but it does not provide the ClientMetadata value as input:</p> <ul> <li> <p>Post
	// authentication</p> </li> <li> <p>Custom message</p> </li> <li> <p>Pre token
	// generation</p> </li> <li> <p>Create auth challenge</p> </li> <li> <p>Define auth
	// challenge</p> </li> <li> <p>Verify auth challenge</p> </li> </ul> <p>For more
	// information, see <a
	// href="https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools-working-with-aws-lambda-triggers.html">Customizing
	// User Pool Workflows with Lambda Triggers</a> in the <i>Amazon Cognito Developer
	// Guide</i>.</p> <note> <p>Take the following limitations into consideration when
	// you use the ClientMetadata parameter:</p> <ul> <li> <p>Amazon Cognito does not
	// store the ClientMetadata value. This data is available only to AWS Lambda
	// triggers that are assigned to a user pool to support custom workflows. If your
	// user pool configuration does not include triggers, the ClientMetadata parameter
	// serves no purpose.</p> </li> <li> <p>Amazon Cognito does not validate the
	// ClientMetadata value.</p> </li> <li> <p>Amazon Cognito does not encrypt the the
	// ClientMetadata value, so don't use it to provide sensitive information.</p>
	// </li> </ul> </note>
	ClientMetadata map[string]*string
	// The analytics metadata for collecting Amazon Pinpoint metrics for
	// AdminInitiateAuth calls.
	AnalyticsMetadata *types.AnalyticsMetadataType
	// Contextual data such as the user's device fingerprint, IP address, or location
	// used for evaluating the risk of an unexpected event by Amazon Cognito advanced
	// security.
	ContextData *types.ContextDataType
	// The authentication parameters. These are inputs corresponding to the AuthFlow
	// that you are invoking. The required values depend on the value of AuthFlow:
	//
	//
	// * For USER_SRP_AUTH: USERNAME (required), SRP_A (required), SECRET_HASH
	// (required if the app client is configured with a client secret), DEVICE_KEY
	//
	//
	// * For REFRESH_TOKEN_AUTH/REFRESH_TOKEN: REFRESH_TOKEN (required), SECRET_HASH
	// (required if the app client is configured with a client secret), DEVICE_KEY
	//
	//
	// * For ADMIN_NO_SRP_AUTH: USERNAME (required), SECRET_HASH (if app client is
	// configured with client secret), PASSWORD (required), DEVICE_KEY
	//
	//     * For
	// CUSTOM_AUTH: USERNAME (required), SECRET_HASH (if app client is configured with
	// client secret), DEVICE_KEY
	AuthParameters map[string]*string
	// The app client ID.
	ClientId *string
}

// Initiates the authentication response, as an administrator.
type AdminInitiateAuthOutput struct {
	// The result of the authentication response. This is only returned if the caller
	// does not need to pass another challenge. If the caller does need to pass another
	// challenge before it gets tokens, ChallengeName, ChallengeParameters, and Session
	// are returned.
	AuthenticationResult *types.AuthenticationResultType
	// The name of the challenge which you are responding to with this call. This is
	// returned to you in the AdminInitiateAuth response if you need to pass another
	// challenge.
	//
	//     * MFA_SETUP: If MFA is required, users who do not have at least
	// one of the MFA methods set up are presented with an MFA_SETUP challenge. The
	// user must set up at least one MFA type to continue to authenticate.
	//
	//     *
	// SELECT_MFA_TYPE: Selects the MFA type. Valid MFA options are SMS_MFA for text
	// SMS MFA, and SOFTWARE_TOKEN_MFA for TOTP software token MFA.
	//
	//     * SMS_MFA:
	// Next challenge is to supply an SMS_MFA_CODE, delivered via SMS.
	//
	//     *
	// PASSWORD_VERIFIER: Next challenge is to supply PASSWORD_CLAIM_SIGNATURE,
	// PASSWORD_CLAIM_SECRET_BLOCK, and TIMESTAMP after the client-side SRP
	// calculations.
	//
	//     * CUSTOM_CHALLENGE: This is returned if your custom
	// authentication flow determines that the user should pass another challenge
	// before tokens are issued.
	//
	//     * DEVICE_SRP_AUTH: If device tracking was enabled
	// on your user pool and the previous challenges were passed, this challenge is
	// returned so that Amazon Cognito can start tracking this device.
	//
	//     *
	// DEVICE_PASSWORD_VERIFIER: Similar to PASSWORD_VERIFIER, but for devices only.
	//
	//
	// * ADMIN_NO_SRP_AUTH: This is returned if you need to authenticate with USERNAME
	// and PASSWORD directly. An app client must be enabled to use this flow.
	//
	//     *
	// NEW_PASSWORD_REQUIRED: For users which are required to change their passwords
	// after successful first login. This challenge should be passed with NEW_PASSWORD
	// and any other required attributes.
	ChallengeName types.ChallengeNameType
	// The session which should be passed both ways in challenge-response calls to the
	// service. If AdminInitiateAuth or AdminRespondToAuthChallenge API call determines
	// that the caller needs to go through another challenge, they return a session
	// with other challenge parameters. This session should be passed as it is to the
	// next AdminRespondToAuthChallenge API call.
	Session *string
	// The challenge parameters. These are returned to you in the AdminInitiateAuth
	// response if you need to pass another challenge. The responses in this parameter
	// should be used to compute inputs to the next call (AdminRespondToAuthChallenge).
	// All challenges require USERNAME and SECRET_HASH (if applicable). The value of
	// the USER_ID_FOR_SRP attribute will be the user's actual username, not an alias
	// (such as email address or phone number), even if you specified an alias in your
	// call to AdminInitiateAuth. This is because, in the AdminRespondToAuthChallenge
	// API ChallengeResponses, the USERNAME attribute cannot be an alias.
	ChallengeParameters map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAdminInitiateAuthMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAdminInitiateAuth{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAdminInitiateAuth{}, middleware.After)
}

func newServiceMetadataMiddleware_opAdminInitiateAuth(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "AdminInitiateAuth",
	}
}