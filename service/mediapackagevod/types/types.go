// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
)

// A MediaPackage VOD Asset resource.
type AssetShallow struct {

	// The ARN of the Asset.
	Arn *string

	// The time the Asset was initially submitted for Ingest.
	CreatedAt *string

	// The unique identifier for the Asset.
	Id *string

	// The ID of the PackagingGroup for the Asset.
	PackagingGroupId *string

	// The resource ID to include in SPEKE key requests.
	ResourceId *string

	// ARN of the source object in S3.
	SourceArn *string

	// The IAM role ARN used to access the source S3 bucket.
	SourceRoleArn *string

	// A collection of tags associated with a resource
	Tags map[string]string

	noSmithyDocumentSerde
}

// CDN Authorization credentials
type Authorization struct {

	// The Amazon Resource Name (ARN) for the secret in AWS Secrets Manager that is
	// used for CDN authorization.
	//
	// This member is required.
	CdnIdentifierSecret *string

	// The Amazon Resource Name (ARN) for the IAM role that allows MediaPackage to
	// communicate with AWS Secrets Manager.
	//
	// This member is required.
	SecretsRoleArn *string

	noSmithyDocumentSerde
}

// A CMAF encryption configuration.
type CmafEncryption struct {

	// A configuration for accessing an external Secure Packager and Encoder Key
	// Exchange (SPEKE) service that will provide encryption keys.
	//
	// This member is required.
	SpekeKeyProvider *SpekeKeyProvider

	// An optional 128-bit, 16-byte hex value represented by a 32-character string,
	// used in conjunction with the key for encrypting blocks. If you don't specify a
	// value, then MediaPackage creates the constant initialization vector (IV).
	ConstantInitializationVector *string

	noSmithyDocumentSerde
}

// A CMAF packaging configuration.
type CmafPackage struct {

	// A list of HLS manifest configurations.
	//
	// This member is required.
	HlsManifests []HlsManifest

	// A CMAF encryption configuration.
	Encryption *CmafEncryption

	// When includeEncoderConfigurationInSegments is set to true, MediaPackage places
	// your encoder's Sequence Parameter Set (SPS), Picture Parameter Set (PPS), and
	// Video Parameter Set (VPS) metadata in every video segment instead of in the init
	// fragment. This lets you use different SPS/PPS/VPS settings for your assets
	// during content playback.
	IncludeEncoderConfigurationInSegments bool

	// Duration (in seconds) of each fragment. Actual fragments will be rounded to the
	// nearest multiple of the source fragment duration.
	SegmentDurationSeconds int32

	noSmithyDocumentSerde
}

// A Dynamic Adaptive Streaming over HTTP (DASH) encryption configuration.
type DashEncryption struct {

	// A configuration for accessing an external Secure Packager and Encoder Key
	// Exchange (SPEKE) service that will provide encryption keys.
	//
	// This member is required.
	SpekeKeyProvider *SpekeKeyProvider

	noSmithyDocumentSerde
}

// A DASH manifest configuration.
type DashManifest struct {

	// Determines the position of some tags in the Media Presentation Description
	// (MPD). When set to FULL, elements like SegmentTemplate and ContentProtection are
	// included in each Representation. When set to COMPACT, duplicate elements are
	// combined and presented at the AdaptationSet level.
	ManifestLayout ManifestLayout

	// An optional string to include in the name of the manifest.
	ManifestName *string

	// Minimum duration (in seconds) that a player will buffer media before starting
	// the presentation.
	MinBufferTimeSeconds int32

	// The Dynamic Adaptive Streaming over HTTP (DASH) profile type. When set to
	// "HBBTV_1_5", HbbTV 1.5 compliant output is enabled.
	Profile Profile

	// A StreamSelection configuration.
	StreamSelection *StreamSelection

	noSmithyDocumentSerde
}

// A Dynamic Adaptive Streaming over HTTP (DASH) packaging configuration.
type DashPackage struct {

	// A list of DASH manifest configurations.
	//
	// This member is required.
	DashManifests []DashManifest

	// A Dynamic Adaptive Streaming over HTTP (DASH) encryption configuration.
	Encryption *DashEncryption

	// When includeEncoderConfigurationInSegments is set to true, MediaPackage places
	// your encoder's Sequence Parameter Set (SPS), Picture Parameter Set (PPS), and
	// Video Parameter Set (VPS) metadata in every video segment instead of in the init
	// fragment. This lets you use different SPS/PPS/VPS settings for your assets
	// during content playback.
	IncludeEncoderConfigurationInSegments bool

	// A list of triggers that controls when the outgoing Dynamic Adaptive Streaming
	// over HTTP (DASH) Media Presentation Description (MPD) will be partitioned into
	// multiple periods. If empty, the content will not be partitioned into more than
	// one period. If the list contains "ADS", new periods will be created where the
	// Asset contains SCTE-35 ad markers.
	PeriodTriggers []PeriodTriggersElement

	// Duration (in seconds) of each segment. Actual segments will be rounded to the
	// nearest multiple of the source segment duration.
	SegmentDurationSeconds int32

	// Determines the type of SegmentTemplate included in the Media Presentation
	// Description (MPD). When set to NUMBER_WITH_TIMELINE, a full timeline is
	// presented in each SegmentTemplate, with $Number$ media URLs. When set to
	// TIME_WITH_TIMELINE, a full timeline is presented in each SegmentTemplate, with
	// $Time$ media URLs. When set to NUMBER_WITH_DURATION, only a duration is included
	// in each SegmentTemplate, with $Number$ media URLs.
	SegmentTemplateFormat SegmentTemplateFormat

	noSmithyDocumentSerde
}

// Configure egress access logging.
type EgressAccessLogs struct {

	// Customize the log group name.
	LogGroupName *string

	noSmithyDocumentSerde
}

// The endpoint URL used to access an Asset using one PackagingConfiguration.
type EgressEndpoint struct {

	// The ID of the PackagingConfiguration being applied to the Asset.
	PackagingConfigurationId *string

	// The URL of the parent manifest for the repackaged Asset.
	Url *string

	noSmithyDocumentSerde
}

// An HTTP Live Streaming (HLS) encryption configuration.
type HlsEncryption struct {

	// A configuration for accessing an external Secure Packager and Encoder Key
	// Exchange (SPEKE) service that will provide encryption keys.
	//
	// This member is required.
	SpekeKeyProvider *SpekeKeyProvider

	// A constant initialization vector for encryption (optional). When not specified
	// the initialization vector will be periodically rotated.
	ConstantInitializationVector *string

	// The encryption method to use.
	EncryptionMethod EncryptionMethod

	noSmithyDocumentSerde
}

// An HTTP Live Streaming (HLS) manifest configuration.
type HlsManifest struct {

	// This setting controls how ad markers are included in the packaged
	// OriginEndpoint. "NONE" will omit all SCTE-35 ad markers from the output.
	// "PASSTHROUGH" causes the manifest to contain a copy of the SCTE-35 ad markers
	// (comments) taken directly from the input HTTP Live Streaming (HLS) manifest.
	// "SCTE35_ENHANCED" generates ad markers and blackout tags based on SCTE-35
	// messages in the input source.
	AdMarkers AdMarkers

	// When enabled, an I-Frame only stream will be included in the output.
	IncludeIframeOnlyStream bool

	// An optional string to include in the name of the manifest.
	ManifestName *string

	// The interval (in seconds) between each EXT-X-PROGRAM-DATE-TIME tag inserted into
	// manifests. Additionally, when an interval is specified ID3Timed Metadata
	// messages will be generated every 5 seconds using the ingest time of the content.
	// If the interval is not specified, or set to 0, then no EXT-X-PROGRAM-DATE-TIME
	// tags will be inserted into manifests and no ID3Timed Metadata messages will be
	// generated. Note that irrespective of this parameter, if any ID3 Timed Metadata
	// is found in HTTP Live Streaming (HLS) input, it will be passed through to HLS
	// output.
	ProgramDateTimeIntervalSeconds int32

	// When enabled, the EXT-X-KEY tag will be repeated in output manifests.
	RepeatExtXKey bool

	// A StreamSelection configuration.
	StreamSelection *StreamSelection

	noSmithyDocumentSerde
}

// An HTTP Live Streaming (HLS) packaging configuration.
type HlsPackage struct {

	// A list of HLS manifest configurations.
	//
	// This member is required.
	HlsManifests []HlsManifest

	// An HTTP Live Streaming (HLS) encryption configuration.
	Encryption *HlsEncryption

	// Duration (in seconds) of each fragment. Actual fragments will be rounded to the
	// nearest multiple of the source fragment duration.
	SegmentDurationSeconds int32

	// When enabled, audio streams will be placed in rendition groups in the output.
	UseAudioRenditionGroup bool

	noSmithyDocumentSerde
}

// A Microsoft Smooth Streaming (MSS) encryption configuration.
type MssEncryption struct {

	// A configuration for accessing an external Secure Packager and Encoder Key
	// Exchange (SPEKE) service that will provide encryption keys.
	//
	// This member is required.
	SpekeKeyProvider *SpekeKeyProvider

	noSmithyDocumentSerde
}

// A Microsoft Smooth Streaming (MSS) manifest configuration.
type MssManifest struct {

	// An optional string to include in the name of the manifest.
	ManifestName *string

	// A StreamSelection configuration.
	StreamSelection *StreamSelection

	noSmithyDocumentSerde
}

// A Microsoft Smooth Streaming (MSS) PackagingConfiguration.
type MssPackage struct {

	// A list of MSS manifest configurations.
	//
	// This member is required.
	MssManifests []MssManifest

	// A Microsoft Smooth Streaming (MSS) encryption configuration.
	Encryption *MssEncryption

	// The duration (in seconds) of each segment.
	SegmentDurationSeconds int32

	noSmithyDocumentSerde
}

// A MediaPackage VOD PackagingConfiguration resource.
type PackagingConfiguration struct {

	// The ARN of the PackagingConfiguration.
	Arn *string

	// A CMAF packaging configuration.
	CmafPackage *CmafPackage

	// A Dynamic Adaptive Streaming over HTTP (DASH) packaging configuration.
	DashPackage *DashPackage

	// An HTTP Live Streaming (HLS) packaging configuration.
	HlsPackage *HlsPackage

	// The ID of the PackagingConfiguration.
	Id *string

	// A Microsoft Smooth Streaming (MSS) PackagingConfiguration.
	MssPackage *MssPackage

	// The ID of a PackagingGroup.
	PackagingGroupId *string

	// A collection of tags associated with a resource
	Tags map[string]string

	noSmithyDocumentSerde
}

// A MediaPackage VOD PackagingGroup resource.
type PackagingGroup struct {

	// The ARN of the PackagingGroup.
	Arn *string

	// CDN Authorization credentials
	Authorization *Authorization

	// The fully qualified domain name for Assets in the PackagingGroup.
	DomainName *string

	// Configure egress access logging.
	EgressAccessLogs *EgressAccessLogs

	// The ID of the PackagingGroup.
	Id *string

	// A collection of tags associated with a resource
	Tags map[string]string

	noSmithyDocumentSerde
}

// A configuration for accessing an external Secure Packager and Encoder Key
// Exchange (SPEKE) service that will provide encryption keys.
type SpekeKeyProvider struct {

	// An Amazon Resource Name (ARN) of an IAM role that AWS Elemental MediaPackage
	// will assume when accessing the key provider service.
	//
	// This member is required.
	RoleArn *string

	// The system IDs to include in key requests.
	//
	// This member is required.
	SystemIds []string

	// The URL of the external key provider service.
	//
	// This member is required.
	Url *string

	noSmithyDocumentSerde
}

// A StreamSelection configuration.
type StreamSelection struct {

	// The maximum video bitrate (bps) to include in output.
	MaxVideoBitsPerSecond int32

	// The minimum video bitrate (bps) to include in output.
	MinVideoBitsPerSecond int32

	// A directive that determines the order of streams in the output.
	StreamOrder StreamOrder

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
