// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// A home region control is an object that specifies the home region for an
// account, with some additional information. It contains a target (always of type
// ACCOUNT), an ID, and a time at which the home region was set.
type HomeRegionControl struct {

	// A unique identifier that's generated for each home region control. It's always a
	// string that begins with "hrc-" followed by 12 lowercase letters and numbers.
	ControlId *string

	// The AWS Region that's been set as home region. For example, "us-west-2" or
	// "eu-central-1" are valid home regions.
	HomeRegion *string

	// A timestamp representing the time when the customer called
	// CreateHomeregionControl and set the home region for the account.
	RequestedTime *time.Time

	// The target parameter specifies the identifier to which the home region is
	// applied, which is always an ACCOUNT. It applies the home region to the current
	// ACCOUNT.
	Target *Target

	noSmithyDocumentSerde
}

// The target parameter specifies the identifier to which the home region is
// applied, which is always an ACCOUNT. It applies the home region to the current
// ACCOUNT.
type Target struct {

	// The target type is always an ACCOUNT.
	//
	// This member is required.
	Type TargetType

	// The TargetID is a 12-character identifier of the ACCOUNT for which the control
	// was created. (This must be the current account.)
	Id *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
