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

// GBTElementBranchStatus the model 'GBTElementBranchStatus'
type GBTElementBranchStatus string

// List of GBTElementBranchStatus
const (
	GBTElementBranchStatusCreated         GBTElementBranchStatus = "CREATED"
	GBTElementBranchStatusDeleted         GBTElementBranchStatus = "DELETED"
	GBTElementBranchStatusEdits           GBTElementBranchStatus = "EDITS"
	GBTElementBranchStatusNotOnThisBranch GBTElementBranchStatus = "NOT_ON_THIS_BRANCH"
	GBTElementBranchStatusNoChanges       GBTElementBranchStatus = "NO_CHANGES"
	GBTElementBranchStatusUnknown         GBTElementBranchStatus = "UNKNOWN"
)

// All allowed values of GBTElementBranchStatus enum
var AllowedGBTElementBranchStatusEnumValues = []GBTElementBranchStatus{
	"CREATED",
	"DELETED",
	"EDITS",
	"NOT_ON_THIS_BRANCH",
	"NO_CHANGES",
	"UNKNOWN",
}

func (v *GBTElementBranchStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GBTElementBranchStatus(value)
	for _, existing := range AllowedGBTElementBranchStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GBTElementBranchStatus", value)
}

// NewGBTElementBranchStatusFromValue returns a pointer to a valid GBTElementBranchStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGBTElementBranchStatusFromValue(v string) (*GBTElementBranchStatus, error) {
	ev := GBTElementBranchStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GBTElementBranchStatus: valid values are %v", v, AllowedGBTElementBranchStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GBTElementBranchStatus) IsValid() bool {
	for _, existing := range AllowedGBTElementBranchStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GBTElementBranchStatus value
func (v GBTElementBranchStatus) Ptr() *GBTElementBranchStatus {
	return &v
}

type NullableGBTElementBranchStatus struct {
	value *GBTElementBranchStatus
	isSet bool
}

func (v NullableGBTElementBranchStatus) Get() *GBTElementBranchStatus {
	return v.value
}

func (v *NullableGBTElementBranchStatus) Set(val *GBTElementBranchStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableGBTElementBranchStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableGBTElementBranchStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGBTElementBranchStatus(val *GBTElementBranchStatus) *NullableGBTElementBranchStatus {
	return &NullableGBTElementBranchStatus{value: val, isSet: true}
}

func (v NullableGBTElementBranchStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGBTElementBranchStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
