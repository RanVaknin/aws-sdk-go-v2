// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/verifiedpermissions/types"
)

func ExampleAttributeValue_outputUsage() {
	var union types.AttributeValue
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.AttributeValueMemberBoolean:
		_ = v.Value // Value is bool

	case *types.AttributeValueMemberEntityIdentifier:
		_ = v.Value // Value is types.EntityIdentifier

	case *types.AttributeValueMemberLong:
		_ = v.Value // Value is int64

	case *types.AttributeValueMemberRecord:
		_ = v.Value // Value is map[string]types.AttributeValue

	case *types.AttributeValueMemberSet:
		_ = v.Value // Value is []types.AttributeValue

	case *types.AttributeValueMemberString:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.EntityIdentifier
var _ *string
var _ map[string]types.AttributeValue
var _ *bool
var _ *int64
var _ []types.AttributeValue

func ExampleConfiguration_outputUsage() {
	var union types.Configuration
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ConfigurationMemberCognitoUserPoolConfiguration:
		_ = v.Value // Value is types.CognitoUserPoolConfiguration

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.CognitoUserPoolConfiguration

func ExampleContextDefinition_outputUsage() {
	var union types.ContextDefinition
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ContextDefinitionMemberContextMap:
		_ = v.Value // Value is map[string]types.AttributeValue

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ map[string]types.AttributeValue

func ExampleEntitiesDefinition_outputUsage() {
	var union types.EntitiesDefinition
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.EntitiesDefinitionMemberEntityList:
		_ = v.Value // Value is []types.EntityItem

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ []types.EntityItem

func ExampleEntityReference_outputUsage() {
	var union types.EntityReference
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.EntityReferenceMemberIdentifier:
		_ = v.Value // Value is types.EntityIdentifier

	case *types.EntityReferenceMemberUnspecified:
		_ = v.Value // Value is bool

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.EntityIdentifier
var _ *bool

func ExamplePolicyDefinition_outputUsage() {
	var union types.PolicyDefinition
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.PolicyDefinitionMemberStatic:
		_ = v.Value // Value is types.StaticPolicyDefinition

	case *types.PolicyDefinitionMemberTemplateLinked:
		_ = v.Value // Value is types.TemplateLinkedPolicyDefinition

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.StaticPolicyDefinition
var _ *types.TemplateLinkedPolicyDefinition

func ExamplePolicyDefinitionDetail_outputUsage() {
	var union types.PolicyDefinitionDetail
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.PolicyDefinitionDetailMemberStatic:
		_ = v.Value // Value is types.StaticPolicyDefinitionDetail

	case *types.PolicyDefinitionDetailMemberTemplateLinked:
		_ = v.Value // Value is types.TemplateLinkedPolicyDefinitionDetail

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.StaticPolicyDefinitionDetail
var _ *types.TemplateLinkedPolicyDefinitionDetail

func ExamplePolicyDefinitionItem_outputUsage() {
	var union types.PolicyDefinitionItem
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.PolicyDefinitionItemMemberStatic:
		_ = v.Value // Value is types.StaticPolicyDefinitionItem

	case *types.PolicyDefinitionItemMemberTemplateLinked:
		_ = v.Value // Value is types.TemplateLinkedPolicyDefinitionItem

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.TemplateLinkedPolicyDefinitionItem
var _ *types.StaticPolicyDefinitionItem

func ExampleSchemaDefinition_outputUsage() {
	var union types.SchemaDefinition
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.SchemaDefinitionMemberCedarJson:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string

func ExampleUpdateConfiguration_outputUsage() {
	var union types.UpdateConfiguration
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.UpdateConfigurationMemberCognitoUserPoolConfiguration:
		_ = v.Value // Value is types.UpdateCognitoUserPoolConfiguration

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.UpdateCognitoUserPoolConfiguration

func ExampleUpdatePolicyDefinition_outputUsage() {
	var union types.UpdatePolicyDefinition
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.UpdatePolicyDefinitionMemberStatic:
		_ = v.Value // Value is types.UpdateStaticPolicyDefinition

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.UpdateStaticPolicyDefinition