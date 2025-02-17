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

// GBTMateType the model 'GBTMateType'
type GBTMateType string

// List of GBTMateType
const (
	GBTMateTypeSlider      GBTMateType = "SLIDER"
	GBTMateTypeCylindrical GBTMateType = "CYLINDRICAL"
	GBTMateTypeRevolute    GBTMateType = "REVOLUTE"
	GBTMateTypePinSlot     GBTMateType = "PIN_SLOT"
	GBTMateTypePlanar      GBTMateType = "PLANAR"
	GBTMateTypeBall        GBTMateType = "BALL"
	GBTMateTypeFastened    GBTMateType = "FASTENED"
	GBTMateTypeParallel    GBTMateType = "PARALLEL"
	GBTMateTypeUnknown     GBTMateType = "UNKNOWN"
)

// All allowed values of GBTMateType enum
var AllowedGBTMateTypeEnumValues = []GBTMateType{
	"SLIDER",
	"CYLINDRICAL",
	"REVOLUTE",
	"PIN_SLOT",
	"PLANAR",
	"BALL",
	"FASTENED",
	"PARALLEL",
	"UNKNOWN",
}

func (v *GBTMateType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GBTMateType(value)
	for _, existing := range AllowedGBTMateTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GBTMateType", value)
}

// NewGBTMateTypeFromValue returns a pointer to a valid GBTMateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGBTMateTypeFromValue(v string) (*GBTMateType, error) {
	ev := GBTMateType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GBTMateType: valid values are %v", v, AllowedGBTMateTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GBTMateType) IsValid() bool {
	for _, existing := range AllowedGBTMateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GBTMateType value
func (v GBTMateType) Ptr() *GBTMateType {
	return &v
}

type NullableGBTMateType struct {
	value *GBTMateType
	isSet bool
}

func (v NullableGBTMateType) Get() *GBTMateType {
	return v.value
}

func (v *NullableGBTMateType) Set(val *GBTMateType) {
	v.value = val
	v.isSet = true
}

func (v NullableGBTMateType) IsSet() bool {
	return v.isSet
}

func (v *NullableGBTMateType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGBTMateType(val *GBTMateType) *NullableGBTMateType {
	return &NullableGBTMateType{value: val, isSet: true}
}

func (v NullableGBTMateType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGBTMateType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
