// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AttributeMatchingModel string

// Enum values for AttributeMatchingModel
const (
	AttributeMatchingModelOneToOne   AttributeMatchingModel = "ONE_TO_ONE"
	AttributeMatchingModelManyToMany AttributeMatchingModel = "MANY_TO_MANY"
)

// Values returns all known values for AttributeMatchingModel. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AttributeMatchingModel) Values() []AttributeMatchingModel {
	return []AttributeMatchingModel{
		"ONE_TO_ONE",
		"MANY_TO_MANY",
	}
}

type IncrementalRunType string

// Enum values for IncrementalRunType
const (
	IncrementalRunTypeImmediate IncrementalRunType = "IMMEDIATE"
)

// Values returns all known values for IncrementalRunType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (IncrementalRunType) Values() []IncrementalRunType {
	return []IncrementalRunType{
		"IMMEDIATE",
	}
}

type JobStatus string

// Enum values for JobStatus
const (
	JobStatusRunning   JobStatus = "RUNNING"
	JobStatusSucceeded JobStatus = "SUCCEEDED"
	JobStatusFailed    JobStatus = "FAILED"
	JobStatusQueued    JobStatus = "QUEUED"
)

// Values returns all known values for JobStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (JobStatus) Values() []JobStatus {
	return []JobStatus{
		"RUNNING",
		"SUCCEEDED",
		"FAILED",
		"QUEUED",
	}
}

type ResolutionType string

// Enum values for ResolutionType
const (
	ResolutionTypeRuleMatching ResolutionType = "RULE_MATCHING"
	ResolutionTypeMlMatching   ResolutionType = "ML_MATCHING"
)

// Values returns all known values for ResolutionType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ResolutionType) Values() []ResolutionType {
	return []ResolutionType{
		"RULE_MATCHING",
		"ML_MATCHING",
	}
}

type SchemaAttributeType string

// Enum values for SchemaAttributeType
const (
	SchemaAttributeTypeName              SchemaAttributeType = "NAME"
	SchemaAttributeTypeNameFirst         SchemaAttributeType = "NAME_FIRST"
	SchemaAttributeTypeNameMiddle        SchemaAttributeType = "NAME_MIDDLE"
	SchemaAttributeTypeNameLast          SchemaAttributeType = "NAME_LAST"
	SchemaAttributeTypeAddress           SchemaAttributeType = "ADDRESS"
	SchemaAttributeTypeAddressStreet1    SchemaAttributeType = "ADDRESS_STREET1"
	SchemaAttributeTypeAddressStreet2    SchemaAttributeType = "ADDRESS_STREET2"
	SchemaAttributeTypeAddressStreet3    SchemaAttributeType = "ADDRESS_STREET3"
	SchemaAttributeTypeAddressCity       SchemaAttributeType = "ADDRESS_CITY"
	SchemaAttributeTypeAddressState      SchemaAttributeType = "ADDRESS_STATE"
	SchemaAttributeTypeAddressCountry    SchemaAttributeType = "ADDRESS_COUNTRY"
	SchemaAttributeTypeAddressPostalcode SchemaAttributeType = "ADDRESS_POSTALCODE"
	SchemaAttributeTypePhone             SchemaAttributeType = "PHONE"
	SchemaAttributeTypePhoneNumber       SchemaAttributeType = "PHONE_NUMBER"
	SchemaAttributeTypePhoneCountrycode  SchemaAttributeType = "PHONE_COUNTRYCODE"
	SchemaAttributeTypeEmailAddress      SchemaAttributeType = "EMAIL_ADDRESS"
	SchemaAttributeTypeUniqueId          SchemaAttributeType = "UNIQUE_ID"
	SchemaAttributeTypeDate              SchemaAttributeType = "DATE"
	SchemaAttributeTypeString            SchemaAttributeType = "STRING"
)

// Values returns all known values for SchemaAttributeType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SchemaAttributeType) Values() []SchemaAttributeType {
	return []SchemaAttributeType{
		"NAME",
		"NAME_FIRST",
		"NAME_MIDDLE",
		"NAME_LAST",
		"ADDRESS",
		"ADDRESS_STREET1",
		"ADDRESS_STREET2",
		"ADDRESS_STREET3",
		"ADDRESS_CITY",
		"ADDRESS_STATE",
		"ADDRESS_COUNTRY",
		"ADDRESS_POSTALCODE",
		"PHONE",
		"PHONE_NUMBER",
		"PHONE_COUNTRYCODE",
		"EMAIL_ADDRESS",
		"UNIQUE_ID",
		"DATE",
		"STRING",
	}
}