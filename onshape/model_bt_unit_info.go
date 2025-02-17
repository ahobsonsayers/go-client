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
)

// BTUnitInfo struct for BTUnitInfo
type BTUnitInfo struct {
	DefaultUnits *BTDefaultUnitsInfo `json:"defaultUnits,omitempty"`
	// Specifies the display precision for every supported unit.
	UnitsDisplayPrecision *map[string]int32 `json:"unitsDisplayPrecision,omitempty"`
}

// NewBTUnitInfo instantiates a new BTUnitInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTUnitInfo() *BTUnitInfo {
	this := BTUnitInfo{}
	return &this
}

// NewBTUnitInfoWithDefaults instantiates a new BTUnitInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTUnitInfoWithDefaults() *BTUnitInfo {
	this := BTUnitInfo{}
	return &this
}

// GetDefaultUnits returns the DefaultUnits field value if set, zero value otherwise.
func (o *BTUnitInfo) GetDefaultUnits() BTDefaultUnitsInfo {
	if o == nil || o.DefaultUnits == nil {
		var ret BTDefaultUnitsInfo
		return ret
	}
	return *o.DefaultUnits
}

// GetDefaultUnitsOk returns a tuple with the DefaultUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUnitInfo) GetDefaultUnitsOk() (*BTDefaultUnitsInfo, bool) {
	if o == nil || o.DefaultUnits == nil {
		return nil, false
	}
	return o.DefaultUnits, true
}

// HasDefaultUnits returns a boolean if a field has been set.
func (o *BTUnitInfo) HasDefaultUnits() bool {
	if o != nil && o.DefaultUnits != nil {
		return true
	}

	return false
}

// SetDefaultUnits gets a reference to the given BTDefaultUnitsInfo and assigns it to the DefaultUnits field.
func (o *BTUnitInfo) SetDefaultUnits(v BTDefaultUnitsInfo) {
	o.DefaultUnits = &v
}

// GetUnitsDisplayPrecision returns the UnitsDisplayPrecision field value if set, zero value otherwise.
func (o *BTUnitInfo) GetUnitsDisplayPrecision() map[string]int32 {
	if o == nil || o.UnitsDisplayPrecision == nil {
		var ret map[string]int32
		return ret
	}
	return *o.UnitsDisplayPrecision
}

// GetUnitsDisplayPrecisionOk returns a tuple with the UnitsDisplayPrecision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUnitInfo) GetUnitsDisplayPrecisionOk() (*map[string]int32, bool) {
	if o == nil || o.UnitsDisplayPrecision == nil {
		return nil, false
	}
	return o.UnitsDisplayPrecision, true
}

// HasUnitsDisplayPrecision returns a boolean if a field has been set.
func (o *BTUnitInfo) HasUnitsDisplayPrecision() bool {
	if o != nil && o.UnitsDisplayPrecision != nil {
		return true
	}

	return false
}

// SetUnitsDisplayPrecision gets a reference to the given map[string]int32 and assigns it to the UnitsDisplayPrecision field.
func (o *BTUnitInfo) SetUnitsDisplayPrecision(v map[string]int32) {
	o.UnitsDisplayPrecision = &v
}

func (o BTUnitInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DefaultUnits != nil {
		toSerialize["defaultUnits"] = o.DefaultUnits
	}
	if o.UnitsDisplayPrecision != nil {
		toSerialize["unitsDisplayPrecision"] = o.UnitsDisplayPrecision
	}
	return json.Marshal(toSerialize)
}

type NullableBTUnitInfo struct {
	value *BTUnitInfo
	isSet bool
}

func (v NullableBTUnitInfo) Get() *BTUnitInfo {
	return v.value
}

func (v *NullableBTUnitInfo) Set(val *BTUnitInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTUnitInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTUnitInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTUnitInfo(val *BTUnitInfo) *NullableBTUnitInfo {
	return &NullableBTUnitInfo{value: val, isSet: true}
}

func (v NullableBTUnitInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTUnitInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
