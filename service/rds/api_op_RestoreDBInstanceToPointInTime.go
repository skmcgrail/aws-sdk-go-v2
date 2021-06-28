// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Restores a DB instance to an arbitrary point in time. You can restore to any
// point in time before the time identified by the LatestRestorableTime property.
// You can restore to a point up to the number of days specified by the
// BackupRetentionPeriod property. The target database is created with most of the
// original configuration, but in a system-selected Availability Zone, with the
// default security group, the default subnet group, and the default DB parameter
// group. By default, the new DB instance is created as a single-AZ deployment
// except when the instance is a SQL Server instance that has an option group that
// is associated with mirroring; in this case, the instance becomes a mirrored
// deployment and not a single-AZ deployment. This command doesn't apply to Aurora
// MySQL and Aurora PostgreSQL. For Aurora, use RestoreDBClusterToPointInTime.
func (c *Client) RestoreDBInstanceToPointInTime(ctx context.Context, params *RestoreDBInstanceToPointInTimeInput, optFns ...func(*Options)) (*RestoreDBInstanceToPointInTimeOutput, error) {
	if params == nil {
		params = &RestoreDBInstanceToPointInTimeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RestoreDBInstanceToPointInTime", params, optFns, c.addOperationRestoreDBInstanceToPointInTimeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RestoreDBInstanceToPointInTimeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type RestoreDBInstanceToPointInTimeInput struct {

	// The name of the new DB instance to be created. Constraints:
	//
	// * Must contain from
	// 1 to 63 letters, numbers, or hyphens
	//
	// * First character must be a letter
	//
	// *
	// Can't end with a hyphen or contain two consecutive hyphens
	//
	// This member is required.
	TargetDBInstanceIdentifier *string

	// A value that indicates whether minor version upgrades are applied automatically
	// to the DB instance during the maintenance window.
	AutoMinorVersionUpgrade *bool

	// The Availability Zone (AZ) where the DB instance will be created. Default: A
	// random, system-chosen Availability Zone. Constraint: You can't specify the
	// AvailabilityZone parameter if the DB instance is a Multi-AZ deployment. Example:
	// us-east-1a
	AvailabilityZone *string

	// A value that indicates whether to copy all tags from the restored DB instance to
	// snapshots of the DB instance. By default, tags are not copied.
	CopyTagsToSnapshot *bool

	// The compute and memory capacity of the Amazon RDS DB instance, for example,
	// db.m4.large. Not all DB instance classes are available in all Amazon Web
	// Services Regions, or for all database engines. For the full list of DB instance
	// classes, and availability for your engine, see DB Instance Class
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html)
	// in the Amazon RDS User Guide. Default: The same DBInstanceClass as the original
	// DB instance.
	DBInstanceClass *string

	// The database name for the restored DB instance. This parameter isn't used for
	// the MySQL or MariaDB engines.
	DBName *string

	// The name of the DB parameter group to associate with this DB instance. If you do
	// not specify a value for DBParameterGroupName, then the default DBParameterGroup
	// for the specified DB engine is used. Constraints:
	//
	// * If supplied, must match the
	// name of an existing DBParameterGroup.
	//
	// * Must be 1 to 255 letters, numbers, or
	// hyphens.
	//
	// * First character must be a letter.
	//
	// * Can't end with a hyphen or
	// contain two consecutive hyphens.
	DBParameterGroupName *string

	// The DB subnet group name to use for the new instance. Constraints: If supplied,
	// must match the name of an existing DBSubnetGroup. Example: mySubnetgroup
	DBSubnetGroupName *string

	// A value that indicates whether the DB instance has deletion protection enabled.
	// The database can't be deleted when deletion protection is enabled. By default,
	// deletion protection is disabled. For more information, see  Deleting a DB
	// Instance
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_DeleteInstance.html).
	DeletionProtection *bool

	// Specify the Active Directory directory ID to restore the DB instance in. The
	// domain must be created prior to this operation. Currently, only MySQL, Microsoft
	// SQL Server, Oracle, and PostgreSQL DB instances can be created in an Active
	// Directory Domain. For more information, see  Kerberos Authentication
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/kerberos-authentication.html)
	// in the Amazon RDS User Guide.
	Domain *string

	// Specify the name of the IAM role to be used when making API calls to the
	// Directory Service.
	DomainIAMRoleName *string

	// The list of logs that the restored DB instance is to export to CloudWatch Logs.
	// The values in the list depend on the DB engine being used. For more information,
	// see Publishing Database Logs to Amazon CloudWatch Logs
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_LogAccess.html#USER_LogAccess.Procedural.UploadtoCloudWatch)
	// in the Amazon RDS User Guide.
	EnableCloudwatchLogsExports []string

	// A value that indicates whether to enable a customer-owned IP address (CoIP) for
	// an RDS on Outposts DB instance. A CoIP provides local or external connectivity
	// to resources in your Outpost subnets through your on-premises network. For some
	// use cases, a CoIP can provide lower latency for connections to the DB instance
	// from outside of its virtual private cloud (VPC) on your local network. For more
	// information about RDS on Outposts, see Working with Amazon RDS on Amazon Web
	// Services Outposts
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-on-outposts.html) in
	// the Amazon RDS User Guide. For more information about CoIPs, see Customer-owned
	// IP addresses
	// (https://docs.aws.amazon.com/outposts/latest/userguide/outposts-networking-components.html#ip-addressing)
	// in the Amazon Web Services Outposts User Guide.
	EnableCustomerOwnedIp *bool

	// A value that indicates whether to enable mapping of Amazon Web Services Identity
	// and Access Management (IAM) accounts to database accounts. By default, mapping
	// is disabled. For more information about IAM database authentication, see  IAM
	// Database Authentication for MySQL and PostgreSQL
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.IAMDBAuth.html)
	// in the Amazon RDS User Guide.
	EnableIAMDatabaseAuthentication *bool

	// The database engine to use for the new instance. Default: The same as source
	// Constraint: Must be compatible with the engine of the source Valid Values:
	//
	// *
	// mariadb
	//
	// * mysql
	//
	// * oracle-ee
	//
	// * oracle-ee-cdb
	//
	// * oracle-se2
	//
	// *
	// oracle-se2-cdb
	//
	// * postgres
	//
	// * sqlserver-ee
	//
	// * sqlserver-se
	//
	// * sqlserver-ex
	//
	// *
	// sqlserver-web
	Engine *string

	// The amount of Provisioned IOPS (input/output operations per second) to be
	// initially allocated for the DB instance. Constraints: Must be an integer greater
	// than 1000. SQL Server Setting the IOPS value for the SQL Server database engine
	// isn't supported.
	Iops *int32

	// License model information for the restored DB instance. Default: Same as source.
	// Valid values: license-included | bring-your-own-license | general-public-license
	LicenseModel *string

	// The upper limit to which Amazon RDS can automatically scale the storage of the
	// DB instance. For more information about this setting, including limitations that
	// apply to it, see  Managing capacity automatically with Amazon RDS storage
	// autoscaling
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PIOPS.StorageTypes.html#USER_PIOPS.Autoscaling)
	// in the Amazon RDS User Guide.
	MaxAllocatedStorage *int32

	// A value that indicates whether the DB instance is a Multi-AZ deployment.
	// Constraint: You can't specify the AvailabilityZone parameter if the DB instance
	// is a Multi-AZ deployment.
	MultiAZ *bool

	// The name of the option group to be used for the restored DB instance. Permanent
	// options, such as the TDE option for Oracle Advanced Security TDE, can't be
	// removed from an option group, and that option group can't be removed from a DB
	// instance once it is associated with a DB instance
	OptionGroupName *string

	// The port number on which the database accepts connections. Constraints: Value
	// must be 1150-65535 Default: The same port as the original DB instance.
	Port *int32

	// The number of CPU cores and the number of threads per core for the DB instance
	// class of the DB instance.
	ProcessorFeatures []types.ProcessorFeature

	// A value that indicates whether the DB instance is publicly accessible. When the
	// DB instance is publicly accessible, its DNS endpoint resolves to the private IP
	// address from within the DB instance's VPC, and to the public IP address from
	// outside of the DB instance's VPC. Access to the DB instance is ultimately
	// controlled by the security group it uses, and that public access is not
	// permitted if the security group assigned to the DB instance doesn't permit it.
	// When the DB instance isn't publicly accessible, it is an internal DB instance
	// with a DNS name that resolves to a private IP address. For more information, see
	// CreateDBInstance.
	PubliclyAccessible *bool

	// The date and time to restore from. Valid Values: Value must be a time in
	// Universal Coordinated Time (UTC) format Constraints:
	//
	// * Must be before the
	// latest restorable time for the DB instance
	//
	// * Can't be specified if the
	// UseLatestRestorableTime parameter is enabled
	//
	// Example: 2009-09-07T23:45:00Z
	RestoreTime *time.Time

	// The Amazon Resource Name (ARN) of the replicated automated backups from which to
	// restore, for example,
	// arn:aws:rds:useast-1:123456789012:auto-backup:ab-L2IJCEXJP7XQ7HOJ4SIEXAMPLE.
	SourceDBInstanceAutomatedBackupsArn *string

	// The identifier of the source DB instance from which to restore. Constraints:
	//
	// *
	// Must match the identifier of an existing DB instance.
	SourceDBInstanceIdentifier *string

	// The resource ID of the source DB instance from which to restore.
	SourceDbiResourceId *string

	// Specifies the storage type to be associated with the DB instance. Valid values:
	// standard | gp2 | io1 If you specify io1, you must also include a value for the
	// Iops parameter. Default: io1 if the Iops parameter is specified, otherwise gp2
	StorageType *string

	// A list of tags. For more information, see Tagging Amazon RDS Resources
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html) in
	// the Amazon RDS User Guide.
	Tags []types.Tag

	// The ARN from the key store with which to associate the instance for TDE
	// encryption.
	TdeCredentialArn *string

	// The password for the given ARN from the key store in order to access the device.
	TdeCredentialPassword *string

	// A value that indicates whether the DB instance class of the DB instance uses its
	// default processor features.
	UseDefaultProcessorFeatures *bool

	// A value that indicates whether the DB instance is restored from the latest
	// backup time. By default, the DB instance isn't restored from the latest backup
	// time. Constraints: Can't be specified if the RestoreTime parameter is provided.
	UseLatestRestorableTime bool

	// A list of EC2 VPC security groups to associate with this DB instance. Default:
	// The default EC2 VPC security group for the DB subnet group's VPC.
	VpcSecurityGroupIds []string

	noSmithyDocumentSerde
}

type RestoreDBInstanceToPointInTimeOutput struct {

	// Contains the details of an Amazon RDS DB instance. This data type is used as a
	// response element in the DescribeDBInstances action.
	DBInstance *types.DBInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRestoreDBInstanceToPointInTimeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpRestoreDBInstanceToPointInTime{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpRestoreDBInstanceToPointInTime{}, middleware.After)
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
	if err = addOpRestoreDBInstanceToPointInTimeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRestoreDBInstanceToPointInTime(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRestoreDBInstanceToPointInTime(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "RestoreDBInstanceToPointInTime",
	}
}
