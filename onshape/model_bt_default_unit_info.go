/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.151.5768-0013f50d23d2
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTDefaultUnitInfo Specifies which unit should be used for each supported quantity type.
type BTDefaultUnitInfo struct {
	// The quantity type.
	Key *string `json:"key,omitempty"`
	// The unit that should be used.
	Value *string `json:"value,omitempty"`
}

// NewBTDefaultUnitInfo instantiates a new BTDefaultUnitInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTDefaultUnitInfo() *BTDefaultUnitInfo {
	this := BTDefaultUnitInfo{}
	return &this
}

// NewBTDefaultUnitInfoWithDefaults instantiates a new BTDefaultUnitInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTDefaultUnitInfoWithDefaults() *BTDefaultUnitInfo {
	this := BTDefaultUnitInfo{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *BTDefaultUnitInfo) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDefaultUnitInfo) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *BTDefaultUnitInfo) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *BTDefaultUnitInfo) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *BTDefaultUnitInfo) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDefaultUnitInfo) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *BTDefaultUnitInfo) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *BTDefaultUnitInfo) SetValue(v string) {
	o.Value = &v
}

func (o BTDefaultUnitInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableBTDefaultUnitInfo struct {
	value *BTDefaultUnitInfo
	isSet bool
}

func (v NullableBTDefaultUnitInfo) Get() *BTDefaultUnitInfo {
	return v.value
}

func (v *NullableBTDefaultUnitInfo) Set(val *BTDefaultUnitInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTDefaultUnitInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTDefaultUnitInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTDefaultUnitInfo(val *BTDefaultUnitInfo) *NullableBTDefaultUnitInfo {
	return &NullableBTDefaultUnitInfo{value: val, isSet: true}
}

func (v NullableBTDefaultUnitInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTDefaultUnitInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
