// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// A top level container for a table. Databases and tables are the fundamental
// management concepts in Amazon Timestream. All tables in a database are encrypted
// with the same KMS key.
type Database struct {

	// The Amazon Resource Name that uniquely identifies this database.
	Arn *string

	// The time when the database was created, calculated from the Unix epoch time.
	CreationTime *time.Time

	// The name of the Timestream database.
	DatabaseName *string

	// The identifier of the KMS key used to encrypt the data stored in the database.
	KmsKeyId *string

	// The last time that this database was updated.
	LastUpdatedTime *time.Time

	// The total number of tables found within a Timestream database.
	TableCount int64

	noSmithyDocumentSerde
}

// Dimension represents the meta data attributes of the time series. For example,
// the name and availability zone of an EC2 instance or the name of the
// manufacturer of a wind turbine are dimensions.
type Dimension struct {

	// Dimension represents the meta data attributes of the time series. For example,
	// the name and availability zone of an EC2 instance or the name of the
	// manufacturer of a wind turbine are dimensions. For constraints on Dimension
	// names, see Naming Constraints
	// (https://docs.aws.amazon.com/timestream/latest/developerguide/ts-limits.html#limits.naming).
	//
	// This member is required.
	Name *string

	// The value of the dimension.
	//
	// This member is required.
	Value *string

	// The data type of the dimension for the time series data point.
	DimensionValueType DimensionValueType

	noSmithyDocumentSerde
}

// Represents an available endpoint against which to make API calls agaisnt, as
// well as the TTL for that endpoint.
type Endpoint struct {

	// An endpoint address.
	//
	// This member is required.
	Address *string

	// The TTL for the endpoint, in minutes.
	//
	// This member is required.
	CachePeriodInMinutes int64

	noSmithyDocumentSerde
}

// Record represents a time series data point being written into Timestream. Each
// record contains an array of dimensions. Dimensions represent the meta data
// attributes of a time series data point such as the instance name or availability
// zone of an EC2 instance. A record also contains the measure name which is the
// name of the measure being collected for example the CPU utilization of an EC2
// instance. A record also contains the measure value and the value type which is
// the data type of the measure value. In addition, the record contains the
// timestamp when the measure was collected that the timestamp unit which
// represents the granularity of the timestamp.
type Record struct {

	// Contains the list of dimensions for time series data points.
	Dimensions []Dimension

	// Measure represents the data attribute of the time series. For example, the CPU
	// utilization of an EC2 instance or the RPM of a wind turbine are measures.
	MeasureName *string

	// Contains the measure value for the time series data point.
	MeasureValue *string

	// Contains the data type of the measure value for the time series data point.
	MeasureValueType MeasureValueType

	// Contains the time at which the measure value for the data point was collected.
	// The time value plus the unit provides the time elapsed since the epoch. For
	// example, if the time value is 12345 and the unit is ms, then 12345 ms have
	// elapsed since the epoch.
	Time *string

	// The granularity of the timestamp unit. It indicates if the time value is in
	// seconds, milliseconds, nanoseconds or other supported values.
	TimeUnit TimeUnit

	// 64-bit attribute used for record updates. Write requests for duplicate data with
	// a higher version number will update the existing measure value and version. In
	// cases where the measure value is the same, Version will still be updated .
	// Default value is to 1.
	Version int64

	noSmithyDocumentSerde
}

// Records that were not successfully inserted into Timestream due to data
// validation issues that must be resolved prior to reinserting time series data
// into the system.
type RejectedRecord struct {

	// The existing version of the record. This value is populated in scenarios where
	// an identical record exists with a higher version than the version in the write
	// request.
	ExistingVersion int64

	// The reason why a record was not successfully inserted into Timestream. Possible
	// causes of failure include:
	//
	// * Records with duplicate data where there are
	// multiple records with the same dimensions, timestamps, and measure names but
	// different measure values.
	//
	// * Records with timestamps that lie outside the
	// retention duration of the memory store When the retention window is updated, you
	// will receive a RejectedRecords exception if you immediately try to ingest data
	// within the new window. To avoid a RejectedRecords exception, wait until the
	// duration of the new window to ingest new data. For further information, see
	// Best Practices for Configuring Timestream
	// (https://docs.aws.amazon.com/timestream/latest/developerguide/best-practices.html#configuration)
	// and the explanation of how storage works in Timestream
	// (https://docs.aws.amazon.com/timestream/latest/developerguide/storage.html).
	//
	// *
	// Records with dimensions or measures that exceed the Timestream defined
	// limits.
	//
	// For more information, see Access Management
	// (https://docs.aws.amazon.com/timestream/latest/developerguide/ts-limits.html) in
	// the Timestream Developer Guide.
	Reason *string

	// The index of the record in the input request for WriteRecords. Indexes begin
	// with 0.
	RecordIndex int32

	noSmithyDocumentSerde
}

// Retention properties contain the duration for which your time series data must
// be stored in the magnetic store and the memory store.
type RetentionProperties struct {

	// The duration for which data must be stored in the magnetic store.
	//
	// This member is required.
	MagneticStoreRetentionPeriodInDays int64

	// The duration for which data must be stored in the memory store.
	//
	// This member is required.
	MemoryStoreRetentionPeriodInHours int64

	noSmithyDocumentSerde
}

// Table represents a database table in Timestream. Tables contain one or more
// related time series. You can modify the retention duration of the memory store
// and the magnetic store for a table.
type Table struct {

	// The Amazon Resource Name that uniquely identifies this table.
	Arn *string

	// The time when the Timestream table was created.
	CreationTime *time.Time

	// The name of the Timestream database that contains this table.
	DatabaseName *string

	// The time when the Timestream table was last updated.
	LastUpdatedTime *time.Time

	// The retention duration for the memory store and magnetic store.
	RetentionProperties *RetentionProperties

	// The name of the Timestream table.
	TableName *string

	// The current state of the table:
	//
	// * DELETING - The table is being deleted.
	//
	// *
	// ACTIVE - The table is ready for use.
	TableStatus TableStatus

	noSmithyDocumentSerde
}

// A tag is a label that you assign to a Timestream database and/or table. Each tag
// consists of a key and an optional value, both of which you define. Tags enable
// you to categorize databases and/or tables, for example, by purpose, owner, or
// environment.
type Tag struct {

	// The key of the tag. Tag keys are case sensitive.
	//
	// This member is required.
	Key *string

	// The value of the tag. Tag values are case-sensitive and can be null.
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
