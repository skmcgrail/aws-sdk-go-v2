// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// A rule that controls access to an Amazon WorkMail organization.
type AccessControlRule struct {

	// Access protocol actions to include in the rule. Valid values include ActiveSync,
	// AutoDiscover, EWS, IMAP, SMTP, WindowsOutlook, and WebMail.
	Actions []string

	// The date that the rule was created.
	DateCreated *time.Time

	// The date that the rule was modified.
	DateModified *time.Time

	// The rule description.
	Description *string

	// The rule effect.
	Effect AccessControlRuleEffect

	// IPv4 CIDR ranges to include in the rule.
	IpRanges []string

	// The rule name.
	Name *string

	// Access protocol actions to exclude from the rule. Valid values include
	// ActiveSync, AutoDiscover, EWS, IMAP, SMTP, WindowsOutlook, and WebMail.
	NotActions []string

	// IPv4 CIDR ranges to exclude from the rule.
	NotIpRanges []string

	// User IDs to exclude from the rule.
	NotUserIds []string

	// User IDs to include in the rule.
	UserIds []string

	noSmithyDocumentSerde
}

// At least one delegate must be associated to the resource to disable automatic
// replies from the resource.
type BookingOptions struct {

	// The resource's ability to automatically reply to requests. If disabled,
	// delegates must be associated to the resource.
	AutoAcceptRequests bool

	// The resource's ability to automatically decline any conflicting requests.
	AutoDeclineConflictingRequests bool

	// The resource's ability to automatically decline any recurring requests.
	AutoDeclineRecurringRequests bool

	noSmithyDocumentSerde
}

// The name of the attribute, which is one of the values defined in the
// UserAttribute enumeration.
type Delegate struct {

	// The identifier for the user or group associated as the resource's delegate.
	//
	// This member is required.
	Id *string

	// The type of the delegate: user or group.
	//
	// This member is required.
	Type MemberType

	noSmithyDocumentSerde
}

// The domain to associate with an Amazon WorkMail organization. When you configure
// a domain hosted in Amazon Route 53 (Route 53), all recommended DNS records are
// added to the organization when you create it. For more information, see Adding a
// domain (https://docs.aws.amazon.com/workmail/latest/adminguide/add_domain.html)
// in the Amazon WorkMail Administrator Guide.
type Domain struct {

	// The fully qualified domain name.
	DomainName *string

	// The hosted zone ID for a domain hosted in Route 53. Required when configuring a
	// domain hosted in Route 53.
	HostedZoneId *string

	noSmithyDocumentSerde
}

// The configuration applied to an organization's folders by its retention policy.
type FolderConfiguration struct {

	// The action to take on the folder contents at the end of the folder configuration
	// period.
	//
	// This member is required.
	Action RetentionAction

	// The folder name.
	//
	// This member is required.
	Name FolderName

	// The period of time at which the folder configuration action is applied.
	Period *int32

	noSmithyDocumentSerde
}

// The representation of an Amazon WorkMail group.
type Group struct {

	// The date indicating when the group was disabled from Amazon WorkMail use.
	DisabledDate *time.Time

	// The email of the group.
	Email *string

	// The date indicating when the group was enabled for Amazon WorkMail use.
	EnabledDate *time.Time

	// The identifier of the group.
	Id *string

	// The name of the group.
	Name *string

	// The state of the group, which can be ENABLED, DISABLED, or DELETED.
	State EntityState

	noSmithyDocumentSerde
}

// The details of a mailbox export job, including the user or resource ID
// associated with the mailbox and the S3 bucket that the mailbox contents are
// exported to.
type MailboxExportJob struct {

	// The mailbox export job description.
	Description *string

	// The mailbox export job end timestamp.
	EndTime *time.Time

	// The identifier of the user or resource associated with the mailbox.
	EntityId *string

	// The estimated progress of the mailbox export job, in percentage points.
	EstimatedProgress int32

	// The identifier of the mailbox export job.
	JobId *string

	// The name of the S3 bucket.
	S3BucketName *string

	// The path to the S3 bucket and file that the mailbox export job exports to.
	S3Path *string

	// The mailbox export job start timestamp.
	StartTime *time.Time

	// The state of the mailbox export job.
	State MailboxExportJobState

	noSmithyDocumentSerde
}

// The representation of a user or group.
type Member struct {

	// The date indicating when the member was disabled from Amazon WorkMail use.
	DisabledDate *time.Time

	// The date indicating when the member was enabled for Amazon WorkMail use.
	EnabledDate *time.Time

	// The identifier of the member.
	Id *string

	// The name of the member.
	Name *string

	// The state of the member, which can be ENABLED, DISABLED, or DELETED.
	State EntityState

	// A member can be a user or group.
	Type MemberType

	noSmithyDocumentSerde
}

// The rule that a simulated user matches.
type MobileDeviceAccessMatchedRule struct {

	// Identifier of the rule that a simulated user matches.
	MobileDeviceAccessRuleId *string

	// Name of a rule that a simulated user matches.
	Name *string

	noSmithyDocumentSerde
}

// A rule that controls access to mobile devices for an Amazon WorkMail group.
type MobileDeviceAccessRule struct {

	// The date and time at which an access rule was created.
	DateCreated *time.Time

	// The date and time at which an access rule was modified.
	DateModified *time.Time

	// The description of a mobile access rule.
	Description *string

	// Device models that a rule will match.
	DeviceModels []string

	// Device operating systems that a rule will match.
	DeviceOperatingSystems []string

	// Device types that a rule will match.
	DeviceTypes []string

	// Device user agents that a rule will match.
	DeviceUserAgents []string

	// The effect of the rule when it matches. Allowed values are ALLOW or DENY.
	Effect MobileDeviceAccessRuleEffect

	// The ID assigned to a mobile access rule.
	MobileDeviceAccessRuleId *string

	// The name of a mobile access rule.
	Name *string

	// Device models that a rule will not match. All other device models will match.
	NotDeviceModels []string

	// Device operating systems that a rule will not match. All other device types will
	// match.
	NotDeviceOperatingSystems []string

	// Device types that a rule will not match. All other device types will match.
	NotDeviceTypes []string

	// Device user agents that a rule will not match. All other device user agents will
	// match.
	NotDeviceUserAgents []string

	noSmithyDocumentSerde
}

// The representation of an organization.
type OrganizationSummary struct {

	// The alias associated with the organization.
	Alias *string

	// The default email domain associated with the organization.
	DefaultMailDomain *string

	// The error message associated with the organization. It is only present if
	// unexpected behavior has occurred with regards to the organization. It provides
	// insight or solutions regarding unexpected behavior.
	ErrorMessage *string

	// The identifier associated with the organization.
	OrganizationId *string

	// The state associated with the organization.
	State *string

	noSmithyDocumentSerde
}

// Permission granted to a user, group, or resource to access a certain aspect of
// another user, group, or resource mailbox.
type Permission struct {

	// The identifier of the user, group, or resource to which the permissions are
	// granted.
	//
	// This member is required.
	GranteeId *string

	// The type of user, group, or resource referred to in GranteeId.
	//
	// This member is required.
	GranteeType MemberType

	// The permissions granted to the grantee. SEND_AS allows the grantee to send email
	// as the owner of the mailbox (the grantee is not mentioned on these emails).
	// SEND_ON_BEHALF allows the grantee to send email on behalf of the owner of the
	// mailbox (the grantee is not mentioned as the physical sender of these emails).
	// FULL_ACCESS allows the grantee full access to the mailbox, irrespective of other
	// folder-level permissions set on the mailbox.
	//
	// This member is required.
	PermissionValues []PermissionType

	noSmithyDocumentSerde
}

// The representation of a resource.
type Resource struct {

	// The date indicating when the resource was disabled from Amazon WorkMail use.
	DisabledDate *time.Time

	// The email of the resource.
	Email *string

	// The date indicating when the resource was enabled for Amazon WorkMail use.
	EnabledDate *time.Time

	// The identifier of the resource.
	Id *string

	// The name of the resource.
	Name *string

	// The state of the resource, which can be ENABLED, DISABLED, or DELETED.
	State EntityState

	// The type of the resource: equipment or room.
	Type ResourceType

	noSmithyDocumentSerde
}

// Describes a tag applied to a resource.
type Tag struct {

	// The key of the tag.
	//
	// This member is required.
	Key *string

	// The value of the tag.
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

// The representation of an Amazon WorkMail user.
type User struct {

	// The date indicating when the user was disabled from Amazon WorkMail use.
	DisabledDate *time.Time

	// The display name of the user.
	DisplayName *string

	// The email of the user.
	Email *string

	// The date indicating when the user was enabled for Amazon WorkMail use.
	EnabledDate *time.Time

	// The identifier of the user.
	Id *string

	// The name of the user.
	Name *string

	// The state of the user, which can be ENABLED, DISABLED, or DELETED.
	State EntityState

	// The role of the user.
	UserRole UserRole

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
