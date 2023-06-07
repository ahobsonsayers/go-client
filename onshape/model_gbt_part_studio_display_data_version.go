/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.164.16955-b4ecd192bba6
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// GBTPartStudioDisplayDataVersion the model 'GBTPartStudioDisplayDataVersion'
type GBTPartStudioDisplayDataVersion string

// List of GBTPartStudioDisplayDataVersion
const (
	GBTPartStudioDisplayDataVersionV0OriginalVersion             GBTPartStudioDisplayDataVersion = "V0_ORIGINAL_VERSION"
	GBTPartStudioDisplayDataVersionV1SmoothEdgesRenderingOptions GBTPartStudioDisplayDataVersion = "V1_SMOOTH_EDGES_RENDERING_OPTIONS"
	GBTPartStudioDisplayDataVersionV2SmoothEdgesToleranceChanged GBTPartStudioDisplayDataVersion = "V2_SMOOTH_EDGES_TOLERANCE_CHANGED"
	GBTPartStudioDisplayDataVersionUnknown                       GBTPartStudioDisplayDataVersion = "UNKNOWN"
)

// All allowed values of GBTPartStudioDisplayDataVersion enum
var AllowedGBTPartStudioDisplayDataVersionEnumValues = []GBTPartStudioDisplayDataVersion{
	"V0_ORIGINAL_VERSION",
	"V1_SMOOTH_EDGES_RENDERING_OPTIONS",
	"V2_SMOOTH_EDGES_TOLERANCE_CHANGED",
	"UNKNOWN",
}

func (v *GBTPartStudioDisplayDataVersion) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GBTPartStudioDisplayDataVersion(value)
	for _, existing := range AllowedGBTPartStudioDisplayDataVersionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GBTPartStudioDisplayDataVersion", value)
}

// NewGBTPartStudioDisplayDataVersionFromValue returns a pointer to a valid GBTPartStudioDisplayDataVersion
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGBTPartStudioDisplayDataVersionFromValue(v string) (*GBTPartStudioDisplayDataVersion, error) {
	ev := GBTPartStudioDisplayDataVersion(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GBTPartStudioDisplayDataVersion: valid values are %v", v, AllowedGBTPartStudioDisplayDataVersionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GBTPartStudioDisplayDataVersion) IsValid() bool {
	for _, existing := range AllowedGBTPartStudioDisplayDataVersionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GBTPartStudioDisplayDataVersion value
func (v GBTPartStudioDisplayDataVersion) Ptr() *GBTPartStudioDisplayDataVersion {
	return &v
}

type NullableGBTPartStudioDisplayDataVersion struct {
	value *GBTPartStudioDisplayDataVersion
	isSet bool
}

func (v NullableGBTPartStudioDisplayDataVersion) Get() *GBTPartStudioDisplayDataVersion {
	return v.value
}

func (v *NullableGBTPartStudioDisplayDataVersion) Set(val *GBTPartStudioDisplayDataVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableGBTPartStudioDisplayDataVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableGBTPartStudioDisplayDataVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGBTPartStudioDisplayDataVersion(val *GBTPartStudioDisplayDataVersion) *NullableGBTPartStudioDisplayDataVersion {
	return &NullableGBTPartStudioDisplayDataVersion{value: val, isSet: true}
}

func (v NullableGBTPartStudioDisplayDataVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGBTPartStudioDisplayDataVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
