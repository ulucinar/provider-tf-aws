/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AccessControlTranslationObservation struct {
}

type AccessControlTranslationParameters struct {

	// +kubebuilder:validation:Required
	Owner *string `json:"owner" tf:"owner,omitempty"`
}

type ApplyServerSideEncryptionByDefaultObservation struct {
}

type ApplyServerSideEncryptionByDefaultParameters struct {

	// The AWS KMS master key ID used for the SSE-KMS encryption. This can only be used when you set the value of sse_algorithm as aws:kms. The default aws/s3 AWS KMS master key is used if this element is absent while the sse_algorithm is aws:kms.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-aws/apis/kms/v1alpha2.Key
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-jet-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	KMSMasterKeyID *string `json:"kmsMasterKeyId,omitempty" tf:"kms_master_key_id,omitempty"`

	// Reference to a Key in kms to populate kmsMasterKeyId.
	// +kubebuilder:validation:Optional
	KMSMasterKeyIDRef *v1.Reference `json:"kmsMasterKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsMasterKeyId.
	// +kubebuilder:validation:Optional
	KMSMasterKeyIDSelector *v1.Selector `json:"kmsMasterKeyIdSelector,omitempty" tf:"-"`

	// The server-side encryption algorithm to use. Valid values are AES256 and aws:kms
	// +kubebuilder:validation:Required
	SseAlgorithm *string `json:"sseAlgorithm" tf:"sse_algorithm,omitempty"`
}

type BucketObservation struct {

	// The bucket domain name. Will be of format bucketname.s3.amazonaws.com.
	BucketDomainName *string `json:"bucketDomainName,omitempty" tf:"bucket_domain_name,omitempty"`

	// The bucket region-specific domain name. The bucket domain name including the region name, please refer here for format. Note: The AWS CloudFront allows specifying S3 region-specific endpoint when creating S3 origin, it will prevent redirect issues from CloudFront to S3 Origin URL.
	BucketRegionalDomainName *string `json:"bucketRegionalDomainName,omitempty" tf:"bucket_regional_domain_name,omitempty"`

	// The name of the bucket.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string. This is used to create Route 53 alias records.
	WebsiteDomain *string `json:"websiteDomain,omitempty" tf:"website_domain,omitempty"`

	// The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
	WebsiteEndpoint *string `json:"websiteEndpoint,omitempty" tf:"website_endpoint,omitempty"`
}

type BucketParameters struct {

	// The canned ACL to apply. Valid values are private, public-read, public-read-write, aws-exec-read, authenticated-read, and log-delivery-write. Defaults to private.  Conflicts with grant. Use the resource aws_s3_bucket_acl instead.
	// +kubebuilder:validation:Optional
	ACL *string `json:"acl,omitempty" tf:"acl,omitempty"`

	// Sets the accelerate configuration of an existing bucket. Can be Enabled or Suspended. Cannot be used in cn-north-1 or us-gov-west-1.
	// Use the resource aws_s3_bucket_accelerate_configuration instead.
	// +kubebuilder:validation:Optional
	AccelerationStatus *string `json:"accelerationStatus,omitempty" tf:"acceleration_status,omitempty"`

	// The ARN of the bucket. Will be of format arn:aws:s3:::bucketname.
	// +kubebuilder:validation:Optional
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// A rule of Cross-Origin Resource Sharing. See CORS rule below for details. Use the resource aws_s3_bucket_cors_configuration instead.
	// +kubebuilder:validation:Optional
	CorsRule []CorsRuleParameters `json:"corsRule,omitempty" tf:"cors_rule,omitempty"`

	// A boolean that indicates all objects (including any locked objects) should be deleted from the bucket so that the bucket can be destroyed without error. These objects are not recoverable.
	// +kubebuilder:validation:Optional
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// An ACL policy grant. See Grant below for details. Conflicts with acl. Use the resource aws_s3_bucket_acl instead.
	// +kubebuilder:validation:Optional
	Grant []GrantParameters `json:"grant,omitempty" tf:"grant,omitempty"`

	// The Route 53 Hosted Zone ID for this bucket's region.
	// +kubebuilder:validation:Optional
	HostedZoneID *string `json:"hostedZoneId,omitempty" tf:"hosted_zone_id,omitempty"`

	// A configuration of object lifecycle management. See Lifecycle Rule below for details.
	// Use the resource aws_s3_bucket_lifecycle_configuration instead.
	// +kubebuilder:validation:Optional
	LifecycleRule []LifecycleRuleParameters `json:"lifecycleRule,omitempty" tf:"lifecycle_rule,omitempty"`

	// A configuration of S3 bucket logging parameters. See Logging below for details.
	// Use the resource aws_s3_bucket_logging instead.
	// +kubebuilder:validation:Optional
	Logging []LoggingParameters `json:"logging,omitempty" tf:"logging,omitempty"`

	// A configuration of S3 object locking. See Object Lock Configuration below for details.
	// Use the object_lock_enabled parameter and the resource aws_s3_bucket_object_lock_configuration instead.
	// +kubebuilder:validation:Optional
	ObjectLockConfiguration []ObjectLockConfigurationParameters `json:"objectLockConfiguration,omitempty" tf:"object_lock_configuration,omitempty"`

	// Indicates whether this bucket has an Object Lock configuration enabled. Valid values are true or false.
	// +kubebuilder:validation:Optional
	ObjectLockEnabled *bool `json:"objectLockEnabled,omitempty" tf:"object_lock_enabled,omitempty"`

	// A valid bucket policy JSON document. In this case, please make sure you use the verbose/specific version of the policy.
	// Use the resource aws_s3_bucket_policy instead.
	// +kubebuilder:validation:Optional
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// The AWS region this bucket resides in.
	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// A configuration of replication configuration. See Replication Configuration below for details.
	// Use the resource aws_s3_bucket_replication_configuration instead.
	// +kubebuilder:validation:Optional
	ReplicationConfiguration []ReplicationConfigurationParameters `json:"replicationConfiguration,omitempty" tf:"replication_configuration,omitempty"`

	// Specifies who should bear the cost of Amazon S3 data transfer.
	// Can be either BucketOwner or Requester. By default, the owner of the S3 bucket would incur the costs of any data transfer.
	// See Requester Pays Buckets developer guide for more information.
	// Use the resource aws_s3_bucket_request_payment_configuration instead.
	// +kubebuilder:validation:Optional
	RequestPayer *string `json:"requestPayer,omitempty" tf:"request_payer,omitempty"`

	// A configuration of server-side encryption configuration. See Server Side Encryption Configuration below for details.
	// Use the resource aws_s3_bucket_server_side_encryption_configuration instead.
	// +kubebuilder:validation:Optional
	ServerSideEncryptionConfiguration []ServerSideEncryptionConfigurationParameters `json:"serverSideEncryptionConfiguration,omitempty" tf:"server_side_encryption_configuration,omitempty"`

	// A map of tags to assign to the bucket. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A configuration of the S3 bucket versioning state. See Versioning below for details. Use the resource aws_s3_bucket_versioning instead.
	// +kubebuilder:validation:Optional
	Versioning []VersioningParameters `json:"versioning,omitempty" tf:"versioning,omitempty"`

	// A configuration of the S3 bucket website. See Website below for details.
	// Use the resource aws_s3_bucket_website_configuration instead.
	// +kubebuilder:validation:Optional
	Website []WebsiteParameters `json:"website,omitempty" tf:"website,omitempty"`
}

type CorsRuleObservation struct {
}

type CorsRuleParameters struct {

	// List of headers allowed.
	// +kubebuilder:validation:Optional
	AllowedHeaders []*string `json:"allowedHeaders,omitempty" tf:"allowed_headers,omitempty"`

	// One or more HTTP methods that you allow the origin to execute. Can be GET, PUT, POST, DELETE or HEAD.
	// +kubebuilder:validation:Required
	AllowedMethods []*string `json:"allowedMethods" tf:"allowed_methods,omitempty"`

	// One or more origins you want customers to be able to access the bucket from.
	// +kubebuilder:validation:Required
	AllowedOrigins []*string `json:"allowedOrigins" tf:"allowed_origins,omitempty"`

	// One or more headers in the response that you want customers to be able to access from their applications (for example, from a JavaScript XMLHttpRequest object).
	// +kubebuilder:validation:Optional
	ExposeHeaders []*string `json:"exposeHeaders,omitempty" tf:"expose_headers,omitempty"`

	// Specifies time in seconds that browser can cache the response for a preflight request.
	// +kubebuilder:validation:Optional
	MaxAgeSeconds *float64 `json:"maxAgeSeconds,omitempty" tf:"max_age_seconds,omitempty"`
}

type DefaultRetentionObservation struct {
}

type DefaultRetentionParameters struct {

	// The number of days that you want to specify for the default retention period.
	// +kubebuilder:validation:Optional
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// The default Object Lock retention mode you want to apply to new objects placed in this bucket. Valid values are GOVERNANCE and COMPLIANCE.
	// +kubebuilder:validation:Required
	Mode *string `json:"mode" tf:"mode,omitempty"`

	// The number of years that you want to specify for the default retention period.
	// +kubebuilder:validation:Optional
	Years *float64 `json:"years,omitempty" tf:"years,omitempty"`
}

type DestinationObservation struct {
}

type DestinationParameters struct {

	// Specifies the overrides to use for object owners on replication. Must be used in conjunction with account_id owner override configuration.
	// +kubebuilder:validation:Optional
	AccessControlTranslation []AccessControlTranslationParameters `json:"accessControlTranslation,omitempty" tf:"access_control_translation,omitempty"`

	// The Account ID to use for overriding the object owner on replication. Must be used in conjunction with access_control_translation override configuration.
	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// The ARN of the S3 bucket where you want Amazon S3 to store replicas of the object identified by the rule.
	// +kubebuilder:validation:Required
	Bucket *string `json:"bucket" tf:"bucket,omitempty"`

	// Enables replication metrics  (documented below).
	// +kubebuilder:validation:Optional
	Metrics []MetricsParameters `json:"metrics,omitempty" tf:"metrics,omitempty"`

	// Destination KMS encryption key ARN for SSE-KMS replication. Must be used in conjunction with
	// sse_kms_encrypted_objects source selection criteria.
	// +kubebuilder:validation:Optional
	ReplicaKMSKeyID *string `json:"replicaKmsKeyId,omitempty" tf:"replica_kms_key_id,omitempty"`

	// Enables S3 Replication Time Control (S3 RTC) (documented below).
	// +kubebuilder:validation:Optional
	ReplicationTime []ReplicationTimeParameters `json:"replicationTime,omitempty" tf:"replication_time,omitempty"`

	// Specifies the Amazon S3 storage class to which you want the object to transition.
	// +kubebuilder:validation:Optional
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`
}

type ExpirationObservation struct {
}

type ExpirationParameters struct {

	// Specifies the date after which you want the corresponding action to take effect.
	// +kubebuilder:validation:Optional
	Date *string `json:"date,omitempty" tf:"date,omitempty"`

	// Specifies the number of days after object creation when the specific rule action takes effect.
	// +kubebuilder:validation:Optional
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// On a versioned bucket (versioning-enabled or versioning-suspended bucket), you can add this element in the lifecycle configuration to direct Amazon S3 to delete expired object delete markers. This cannot be specified with Days or Date in a Lifecycle Expiration Policy.
	// +kubebuilder:validation:Optional
	ExpiredObjectDeleteMarker *bool `json:"expiredObjectDeleteMarker,omitempty" tf:"expired_object_delete_marker,omitempty"`
}

type FilterObservation struct {
}

type FilterParameters struct {

	// Object keyname prefix that identifies subset of objects to which the rule applies. Must be less than or equal to 1024 characters in length.
	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// A map of tags that identifies subset of objects to which the rule applies.
	// The rule applies only to objects having all the tags in its tagset.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type GrantObservation struct {
}

type GrantParameters struct {

	// Canonical user id to grant for. Used only when type is CanonicalUser.
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// List of permissions to apply for grantee. Valid values are READ, WRITE, READ_ACP, WRITE_ACP, FULL_CONTROL.
	// +kubebuilder:validation:Required
	Permissions []*string `json:"permissions" tf:"permissions,omitempty"`

	// Type of grantee to apply for. Valid values are CanonicalUser and Group. AmazonCustomerByEmail is not supported.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// Uri address to grant for. Used only when type is Group.
	// +kubebuilder:validation:Optional
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`
}

type LifecycleRuleObservation struct {
}

type LifecycleRuleParameters struct {

	// Specifies the number of days after initiating a multipart upload when the multipart upload must be completed.
	// +kubebuilder:validation:Optional
	AbortIncompleteMultipartUploadDays *float64 `json:"abortIncompleteMultipartUploadDays,omitempty" tf:"abort_incomplete_multipart_upload_days,omitempty"`

	// Specifies lifecycle rule status.
	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// Specifies a period in the object's expire. See Expiration below for details.
	// +kubebuilder:validation:Optional
	Expiration []ExpirationParameters `json:"expiration,omitempty" tf:"expiration,omitempty"`

	// Unique identifier for the rule. Must be less than or equal to 255 characters in length.
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies when noncurrent object versions expire. See Noncurrent Version Expiration below for details.
	// +kubebuilder:validation:Optional
	NoncurrentVersionExpiration []NoncurrentVersionExpirationParameters `json:"noncurrentVersionExpiration,omitempty" tf:"noncurrent_version_expiration,omitempty"`

	// Specifies when noncurrent object versions transitions. See Noncurrent Version Transition below for details.
	// +kubebuilder:validation:Optional
	NoncurrentVersionTransition []NoncurrentVersionTransitionParameters `json:"noncurrentVersionTransition,omitempty" tf:"noncurrent_version_transition,omitempty"`

	// Object key prefix identifying one or more objects to which the rule applies.
	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// Specifies object tags key and value.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies a period in the object's transitions. See Transition below for details.
	// +kubebuilder:validation:Optional
	Transition []TransitionParameters `json:"transition,omitempty" tf:"transition,omitempty"`
}

type LoggingObservation struct {
}

type LoggingParameters struct {

	// The name of the bucket that will receive the log objects.
	// +kubebuilder:validation:Required
	TargetBucket *string `json:"targetBucket" tf:"target_bucket,omitempty"`

	// To specify a key prefix for log objects.
	// +kubebuilder:validation:Optional
	TargetPrefix *string `json:"targetPrefix,omitempty" tf:"target_prefix,omitempty"`
}

type MetricsObservation struct {
}

type MetricsParameters struct {

	// Threshold within which objects are to be replicated. The only valid value is 15.
	// +kubebuilder:validation:Optional
	Minutes *float64 `json:"minutes,omitempty" tf:"minutes,omitempty"`

	// The status of the rule. Either Enabled or Disabled. The rule is ignored if status is not Enabled.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type NoncurrentVersionExpirationObservation struct {
}

type NoncurrentVersionExpirationParameters struct {

	// Specifies the number of days after object creation when the specific rule action takes effect.
	// +kubebuilder:validation:Optional
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`
}

type NoncurrentVersionTransitionObservation struct {
}

type NoncurrentVersionTransitionParameters struct {

	// Specifies the number of days after object creation when the specific rule action takes effect.
	// +kubebuilder:validation:Optional
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// Specifies the Amazon S3 storage class to which you want the object to transition.
	// +kubebuilder:validation:Required
	StorageClass *string `json:"storageClass" tf:"storage_class,omitempty"`
}

type ObjectLockConfigurationObservation struct {
}

type ObjectLockConfigurationParameters struct {

	// Indicates whether this bucket has an Object Lock configuration enabled. Valid value is Enabled. Use the top-level argument object_lock_enabled instead.
	// +kubebuilder:validation:Optional
	ObjectLockEnabled *string `json:"objectLockEnabled,omitempty" tf:"object_lock_enabled,omitempty"`

	// The Object Lock rule in place for this bucket (documented below).
	// +kubebuilder:validation:Optional
	Rule []RuleParameters `json:"rule,omitempty" tf:"rule,omitempty"`
}

type ReplicationConfigurationObservation struct {
}

type ReplicationConfigurationParameters struct {

	// The ARN of the IAM role for Amazon S3 to assume when replicating the objects.
	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`

	// Specifies the rules managing the replication (documented below).
	// +kubebuilder:validation:Required
	Rules []RulesParameters `json:"rules" tf:"rules,omitempty"`
}

type ReplicationTimeObservation struct {
}

type ReplicationTimeParameters struct {

	// Threshold within which objects are to be replicated. The only valid value is 15.
	// +kubebuilder:validation:Optional
	Minutes *float64 `json:"minutes,omitempty" tf:"minutes,omitempty"`

	// The status of the rule. Either Enabled or Disabled. The rule is ignored if status is not Enabled.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type RuleObservation struct {
}

type RuleParameters struct {

	// The default retention period that you want to apply to new objects placed in this bucket (documented below).
	// +kubebuilder:validation:Required
	DefaultRetention []DefaultRetentionParameters `json:"defaultRetention" tf:"default_retention,omitempty"`
}

type RulesObservation struct {
}

type RulesParameters struct {

	// Whether delete markers are replicated. The only valid value is Enabled. To disable, omit this argument. This argument is only valid with V2 replication configurations (i.e., when filter is used).
	// +kubebuilder:validation:Optional
	DeleteMarkerReplicationStatus *string `json:"deleteMarkerReplicationStatus,omitempty" tf:"delete_marker_replication_status,omitempty"`

	// Specifies the destination for the rule (documented below).
	// +kubebuilder:validation:Required
	Destination []DestinationParameters `json:"destination" tf:"destination,omitempty"`

	// Filter that identifies subset of objects to which the replication rule applies (documented below).
	// +kubebuilder:validation:Optional
	Filter []FilterParameters `json:"filter,omitempty" tf:"filter,omitempty"`

	// Unique identifier for the rule. Must be less than or equal to 255 characters in length.
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Object keyname prefix identifying one or more objects to which the rule applies. Must be less than or equal to 1024 characters in length.
	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// The priority associated with the rule. Priority should only be set if filter is configured. If not provided, defaults to 0. Priority must be unique between multiple rules.
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// Specifies special object selection criteria (documented below).
	// +kubebuilder:validation:Optional
	SourceSelectionCriteria []SourceSelectionCriteriaParameters `json:"sourceSelectionCriteria,omitempty" tf:"source_selection_criteria,omitempty"`

	// The status of the rule. Either Enabled or Disabled. The rule is ignored if status is not Enabled.
	// +kubebuilder:validation:Required
	Status *string `json:"status" tf:"status,omitempty"`
}

type ServerSideEncryptionConfigurationObservation struct {
}

type ServerSideEncryptionConfigurationParameters struct {

	// A single object for server-side encryption by default configuration. (documented below)
	// +kubebuilder:validation:Required
	Rule []ServerSideEncryptionConfigurationRuleParameters `json:"rule" tf:"rule,omitempty"`
}

type ServerSideEncryptionConfigurationRuleObservation struct {
}

type ServerSideEncryptionConfigurationRuleParameters struct {

	// A single object for setting server-side encryption by default. (documented below)
	// +kubebuilder:validation:Required
	ApplyServerSideEncryptionByDefault []ApplyServerSideEncryptionByDefaultParameters `json:"applyServerSideEncryptionByDefault" tf:"apply_server_side_encryption_by_default,omitempty"`

	// Whether or not to use Amazon S3 Bucket Keys for SSE-KMS.
	// +kubebuilder:validation:Optional
	BucketKeyEnabled *bool `json:"bucketKeyEnabled,omitempty" tf:"bucket_key_enabled,omitempty"`
}

type SourceSelectionCriteriaObservation struct {
}

type SourceSelectionCriteriaParameters struct {

	// Match SSE-KMS encrypted objects (documented below). If specified, replica_kms_key_id
	// in destination must be specified as well.
	// +kubebuilder:validation:Optional
	SseKMSEncryptedObjects []SseKMSEncryptedObjectsParameters `json:"sseKmsEncryptedObjects,omitempty" tf:"sse_kms_encrypted_objects,omitempty"`
}

type SseKMSEncryptedObjectsObservation struct {
}

type SseKMSEncryptedObjectsParameters struct {

	// Enable versioning. Once you version-enable a bucket, it can never return to an unversioned state. You can, however, suspend versioning on that bucket.
	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`
}

type TransitionObservation struct {
}

type TransitionParameters struct {

	// Specifies the date after which you want the corresponding action to take effect.
	// +kubebuilder:validation:Optional
	Date *string `json:"date,omitempty" tf:"date,omitempty"`

	// Specifies the number of days after object creation when the specific rule action takes effect.
	// +kubebuilder:validation:Optional
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// Specifies the Amazon S3 storage class to which you want the object to transition.
	// +kubebuilder:validation:Required
	StorageClass *string `json:"storageClass" tf:"storage_class,omitempty"`
}

type VersioningObservation struct {
}

type VersioningParameters struct {

	// Enable versioning. Once you version-enable a bucket, it can never return to an unversioned state. You can, however, suspend versioning on that bucket.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Enable MFA delete for either Change the versioning state of your bucket or Permanently delete an object version. Default is false. This cannot be used to toggle this setting but is available to allow managed buckets to reflect the state in AWS
	// +kubebuilder:validation:Optional
	MfaDelete *bool `json:"mfaDelete,omitempty" tf:"mfa_delete,omitempty"`
}

type WebsiteObservation struct {
}

type WebsiteParameters struct {

	// An absolute path to the document to return in case of a 4XX error.
	// +kubebuilder:validation:Optional
	ErrorDocument *string `json:"errorDocument,omitempty" tf:"error_document,omitempty"`

	// Amazon S3 returns this index document when requests are made to the root domain or any of the subfolders.
	// +kubebuilder:validation:Optional
	IndexDocument *string `json:"indexDocument,omitempty" tf:"index_document,omitempty"`

	// A hostname to redirect all website requests for this bucket to. Hostname can optionally be prefixed with a protocol (http:// or https://) to use when redirecting requests. The default is the protocol that is used in the original request.
	// +kubebuilder:validation:Optional
	RedirectAllRequestsTo *string `json:"redirectAllRequestsTo,omitempty" tf:"redirect_all_requests_to,omitempty"`

	// A json array containing routing rules
	// describing redirect behavior and when redirects are applied.
	// +kubebuilder:validation:Optional
	RoutingRules *string `json:"routingRules,omitempty" tf:"routing_rules,omitempty"`
}

// BucketSpec defines the desired state of Bucket
type BucketSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketParameters `json:"forProvider"`
}

// BucketStatus defines the observed state of Bucket.
type BucketStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Bucket is the Schema for the Buckets API. Provides a S3 bucket resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type Bucket struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BucketSpec   `json:"spec"`
	Status            BucketStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketList contains a list of Buckets
type BucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Bucket `json:"items"`
}

// Repository type metadata.
var (
	Bucket_Kind             = "Bucket"
	Bucket_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Bucket_Kind}.String()
	Bucket_KindAPIVersion   = Bucket_Kind + "." + CRDGroupVersion.String()
	Bucket_GroupVersionKind = CRDGroupVersion.WithKind(Bucket_Kind)
)

func init() {
	SchemeBuilder.Register(&Bucket{}, &BucketList{})
}
