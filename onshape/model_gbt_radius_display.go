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

// GBTRadiusDisplay the model 'GBTRadiusDisplay'
type GBTRadiusDisplay string

// List of GBTRadiusDisplay
const (
	GBTRadiusDisplayNone      GBTRadiusDisplay = "NONE"
	GBTRadiusDisplayRadial    GBTRadiusDisplay = "RADIAL"
	GBTRadiusDisplayDiametral GBTRadiusDisplay = "DIAMETRAL"
	GBTRadiusDisplayUnknown   GBTRadiusDisplay = "UNKNOWN"
)

// All allowed values of GBTRadiusDisplay enum
var AllowedGBTRadiusDisplayEnumValues = []GBTRadiusDisplay{
	"NONE",
	"RADIAL",
	"DIAMETRAL",
	"UNKNOWN",
}

func (v *GBTRadiusDisplay) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GBTRadiusDisplay(value)
	for _, existing := range AllowedGBTRadiusDisplayEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GBTRadiusDisplay", value)
}

// NewGBTRadiusDisplayFromValue returns a pointer to a valid GBTRadiusDisplay
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGBTRadiusDisplayFromValue(v string) (*GBTRadiusDisplay, error) {
	ev := GBTRadiusDisplay(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GBTRadiusDisplay: valid values are %v", v, AllowedGBTRadiusDisplayEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GBTRadiusDisplay) IsValid() bool {
	for _, existing := range AllowedGBTRadiusDisplayEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GBTRadiusDisplay value
func (v GBTRadiusDisplay) Ptr() *GBTRadiusDisplay {
	return &v
}

type NullableGBTRadiusDisplay struct {
	value *GBTRadiusDisplay
	isSet bool
}

func (v NullableGBTRadiusDisplay) Get() *GBTRadiusDisplay {
	return v.value
}

func (v *NullableGBTRadiusDisplay) Set(val *GBTRadiusDisplay) {
	v.value = val
	v.isSet = true
}

func (v NullableGBTRadiusDisplay) IsSet() bool {
	return v.isSet
}

func (v *NullableGBTRadiusDisplay) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGBTRadiusDisplay(val *GBTRadiusDisplay) *NullableGBTRadiusDisplay {
	return &NullableGBTRadiusDisplay{value: val, isSet: true}
}

func (v NullableGBTRadiusDisplay) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGBTRadiusDisplay) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
