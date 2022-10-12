/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.154.6870-e8e79a24dc2c
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTKeyMouseValuesInfo struct for BTKeyMouseValuesInfo
type BTKeyMouseValuesInfo struct {
	Keys         []string `json:"keys,omitempty"`
	MouseButtons []string `json:"mouseButtons,omitempty"`
}

// NewBTKeyMouseValuesInfo instantiates a new BTKeyMouseValuesInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTKeyMouseValuesInfo() *BTKeyMouseValuesInfo {
	this := BTKeyMouseValuesInfo{}
	return &this
}

// NewBTKeyMouseValuesInfoWithDefaults instantiates a new BTKeyMouseValuesInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTKeyMouseValuesInfoWithDefaults() *BTKeyMouseValuesInfo {
	this := BTKeyMouseValuesInfo{}
	return &this
}

// GetKeys returns the Keys field value if set, zero value otherwise.
func (o *BTKeyMouseValuesInfo) GetKeys() []string {
	if o == nil || o.Keys == nil {
		var ret []string
		return ret
	}
	return o.Keys
}

// GetKeysOk returns a tuple with the Keys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTKeyMouseValuesInfo) GetKeysOk() ([]string, bool) {
	if o == nil || o.Keys == nil {
		return nil, false
	}
	return o.Keys, true
}

// HasKeys returns a boolean if a field has been set.
func (o *BTKeyMouseValuesInfo) HasKeys() bool {
	if o != nil && o.Keys != nil {
		return true
	}

	return false
}

// SetKeys gets a reference to the given []string and assigns it to the Keys field.
func (o *BTKeyMouseValuesInfo) SetKeys(v []string) {
	o.Keys = v
}

// GetMouseButtons returns the MouseButtons field value if set, zero value otherwise.
func (o *BTKeyMouseValuesInfo) GetMouseButtons() []string {
	if o == nil || o.MouseButtons == nil {
		var ret []string
		return ret
	}
	return o.MouseButtons
}

// GetMouseButtonsOk returns a tuple with the MouseButtons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTKeyMouseValuesInfo) GetMouseButtonsOk() ([]string, bool) {
	if o == nil || o.MouseButtons == nil {
		return nil, false
	}
	return o.MouseButtons, true
}

// HasMouseButtons returns a boolean if a field has been set.
func (o *BTKeyMouseValuesInfo) HasMouseButtons() bool {
	if o != nil && o.MouseButtons != nil {
		return true
	}

	return false
}

// SetMouseButtons gets a reference to the given []string and assigns it to the MouseButtons field.
func (o *BTKeyMouseValuesInfo) SetMouseButtons(v []string) {
	o.MouseButtons = v
}

func (o BTKeyMouseValuesInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Keys != nil {
		toSerialize["keys"] = o.Keys
	}
	if o.MouseButtons != nil {
		toSerialize["mouseButtons"] = o.MouseButtons
	}
	return json.Marshal(toSerialize)
}

type NullableBTKeyMouseValuesInfo struct {
	value *BTKeyMouseValuesInfo
	isSet bool
}

func (v NullableBTKeyMouseValuesInfo) Get() *BTKeyMouseValuesInfo {
	return v.value
}

func (v *NullableBTKeyMouseValuesInfo) Set(val *BTKeyMouseValuesInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTKeyMouseValuesInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTKeyMouseValuesInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTKeyMouseValuesInfo(val *BTKeyMouseValuesInfo) *NullableBTKeyMouseValuesInfo {
	return &NullableBTKeyMouseValuesInfo{value: val, isSet: true}
}

func (v NullableBTKeyMouseValuesInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTKeyMouseValuesInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
