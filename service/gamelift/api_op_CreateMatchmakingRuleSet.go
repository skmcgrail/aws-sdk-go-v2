// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new rule set for FlexMatch matchmaking. A rule set describes the type
// of match to create, such as the number and size of teams. It also sets the
// parameters for acceptable player matches, such as minimum skill level or
// character type. A rule set is used by a MatchmakingConfiguration. To create a
// matchmaking rule set, provide unique rule set name and the rule set body in JSON
// format. Rule sets must be defined in the same Region as the matchmaking
// configuration they are used with. Since matchmaking rule sets cannot be edited,
// it is a good idea to check the rule set syntax using ValidateMatchmakingRuleSet
// before creating a new rule set. Learn more
//
// * Build a rule set
// (https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-rulesets.html)
//
// *
// Design a matchmaker
// (https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-configuration.html)
//
// *
// Matchmaking with FlexMatch
// (https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-intro.html)
//
// Related
// actions CreateMatchmakingConfiguration | DescribeMatchmakingConfigurations |
// UpdateMatchmakingConfiguration | DeleteMatchmakingConfiguration |
// CreateMatchmakingRuleSet | DescribeMatchmakingRuleSets |
// ValidateMatchmakingRuleSet | DeleteMatchmakingRuleSet | All APIs by task
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html#reference-awssdk-resources-fleets)
func (c *Client) CreateMatchmakingRuleSet(ctx context.Context, params *CreateMatchmakingRuleSetInput, optFns ...func(*Options)) (*CreateMatchmakingRuleSetOutput, error) {
	if params == nil {
		params = &CreateMatchmakingRuleSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateMatchmakingRuleSet", params, optFns, c.addOperationCreateMatchmakingRuleSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateMatchmakingRuleSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request operation.
type CreateMatchmakingRuleSetInput struct {

	// A unique identifier for the matchmaking rule set. A matchmaking configuration
	// identifies the rule set it uses by this name value. Note that the rule set name
	// is different from the optional name field in the rule set body.
	//
	// This member is required.
	Name *string

	// A collection of matchmaking rules, formatted as a JSON string. Comments are not
	// allowed in JSON, but most elements support a description field.
	//
	// This member is required.
	RuleSetBody *string

	// A list of labels to assign to the new matchmaking rule set resource. Tags are
	// developer-defined key-value pairs. Tagging AWS resources are useful for resource
	// management, access management and cost allocation. For more information, see
	// Tagging AWS Resources
	// (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the AWS
	// General Reference. Once the resource is created, you can use TagResource,
	// UntagResource, and ListTagsForResource to add, remove, and view tags. The
	// maximum tag limit may be lower than stated. See the AWS General Reference for
	// actual tagging limits.
	Tags []types.Tag

	noSmithyDocumentSerde
}

// Represents the returned data in response to a request operation.
type CreateMatchmakingRuleSetOutput struct {

	// The newly created matchmaking rule set.
	//
	// This member is required.
	RuleSet *types.MatchmakingRuleSet

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateMatchmakingRuleSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateMatchmakingRuleSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateMatchmakingRuleSet{}, middleware.After)
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
	if err = addOpCreateMatchmakingRuleSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateMatchmakingRuleSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateMatchmakingRuleSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "CreateMatchmakingRuleSet",
	}
}
