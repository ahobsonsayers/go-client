/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.166.18273-3025d52f81b7
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// GBTAppElementAssociativeDataType the model 'GBTAppElementAssociativeDataType'
type GBTAppElementAssociativeDataType string

// List of GBTAppElementAssociativeDataType
const (
	GBTAppElementAssociativeDataTypeOnshapeDrawingView     GBTAppElementAssociativeDataType = "ONSHAPE_DRAWING_VIEW"
	GBTAppElementAssociativeDataTypeModelTopology          GBTAppElementAssociativeDataType = "MODEL_TOPOLOGY"
	GBTAppElementAssociativeDataTypeModelDefinitionFeature GBTAppElementAssociativeDataType = "MODEL_DEFINITION_FEATURE"
	GBTAppElementAssociativeDataTypeModelDefinitionEntity  GBTAppElementAssociativeDataType = "MODEL_DEFINITION_ENTITY"
	GBTAppElementAssociativeDataTypeUnknown                GBTAppElementAssociativeDataType = "UNKNOWN"
)

// All allowed values of GBTAppElementAssociativeDataType enum
var AllowedGBTAppElementAssociativeDataTypeEnumValues = []GBTAppElementAssociativeDataType{
	"ONSHAPE_DRAWING_VIEW",
	"MODEL_TOPOLOGY",
	"MODEL_DEFINITION_FEATURE",
	"MODEL_DEFINITION_ENTITY",
	"UNKNOWN",
}

func (v *GBTAppElementAssociativeDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GBTAppElementAssociativeDataType(value)
	for _, existing := range AllowedGBTAppElementAssociativeDataTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GBTAppElementAssociativeDataType", value)
}

// NewGBTAppElementAssociativeDataTypeFromValue returns a pointer to a valid GBTAppElementAssociativeDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGBTAppElementAssociativeDataTypeFromValue(v string) (*GBTAppElementAssociativeDataType, error) {
	ev := GBTAppElementAssociativeDataType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GBTAppElementAssociativeDataType: valid values are %v", v, AllowedGBTAppElementAssociativeDataTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GBTAppElementAssociativeDataType) IsValid() bool {
	for _, existing := range AllowedGBTAppElementAssociativeDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GBTAppElementAssociativeDataType value
func (v GBTAppElementAssociativeDataType) Ptr() *GBTAppElementAssociativeDataType {
	return &v
}

type NullableGBTAppElementAssociativeDataType struct {
	value *GBTAppElementAssociativeDataType
	isSet bool
}

func (v NullableGBTAppElementAssociativeDataType) Get() *GBTAppElementAssociativeDataType {
	return v.value
}

func (v *NullableGBTAppElementAssociativeDataType) Set(val *GBTAppElementAssociativeDataType) {
	v.value = val
	v.isSet = true
}

func (v NullableGBTAppElementAssociativeDataType) IsSet() bool {
	return v.isSet
}

func (v *NullableGBTAppElementAssociativeDataType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGBTAppElementAssociativeDataType(val *GBTAppElementAssociativeDataType) *NullableGBTAppElementAssociativeDataType {
	return &NullableGBTAppElementAssociativeDataType{value: val, isSet: true}
}

func (v NullableGBTAppElementAssociativeDataType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGBTAppElementAssociativeDataType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
