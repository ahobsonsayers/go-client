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

// BTComputedConfigurationInputSpec2525 struct for BTComputedConfigurationInputSpec2525
type BTComputedConfigurationInputSpec2525 struct {
	BtType  *string `json:"btType,omitempty"`
	InputId *string `json:"inputId,omitempty"`
}

// NewBTComputedConfigurationInputSpec2525 instantiates a new BTComputedConfigurationInputSpec2525 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTComputedConfigurationInputSpec2525() *BTComputedConfigurationInputSpec2525 {
	this := BTComputedConfigurationInputSpec2525{}
	return &this
}

// NewBTComputedConfigurationInputSpec2525WithDefaults instantiates a new BTComputedConfigurationInputSpec2525 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTComputedConfigurationInputSpec2525WithDefaults() *BTComputedConfigurationInputSpec2525 {
	this := BTComputedConfigurationInputSpec2525{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTComputedConfigurationInputSpec2525) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTComputedConfigurationInputSpec2525) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTComputedConfigurationInputSpec2525) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTComputedConfigurationInputSpec2525) SetBtType(v string) {
	o.BtType = &v
}

// GetInputId returns the InputId field value if set, zero value otherwise.
func (o *BTComputedConfigurationInputSpec2525) GetInputId() string {
	if o == nil || o.InputId == nil {
		var ret string
		return ret
	}
	return *o.InputId
}

// GetInputIdOk returns a tuple with the InputId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTComputedConfigurationInputSpec2525) GetInputIdOk() (*string, bool) {
	if o == nil || o.InputId == nil {
		return nil, false
	}
	return o.InputId, true
}

// HasInputId returns a boolean if a field has been set.
func (o *BTComputedConfigurationInputSpec2525) HasInputId() bool {
	if o != nil && o.InputId != nil {
		return true
	}

	return false
}

// SetInputId gets a reference to the given string and assigns it to the InputId field.
func (o *BTComputedConfigurationInputSpec2525) SetInputId(v string) {
	o.InputId = &v
}

func (o BTComputedConfigurationInputSpec2525) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.InputId != nil {
		toSerialize["inputId"] = o.InputId
	}
	return json.Marshal(toSerialize)
}

type NullableBTComputedConfigurationInputSpec2525 struct {
	value *BTComputedConfigurationInputSpec2525
	isSet bool
}

func (v NullableBTComputedConfigurationInputSpec2525) Get() *BTComputedConfigurationInputSpec2525 {
	return v.value
}

func (v *NullableBTComputedConfigurationInputSpec2525) Set(val *BTComputedConfigurationInputSpec2525) {
	v.value = val
	v.isSet = true
}

func (v NullableBTComputedConfigurationInputSpec2525) IsSet() bool {
	return v.isSet
}

func (v *NullableBTComputedConfigurationInputSpec2525) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTComputedConfigurationInputSpec2525(val *BTComputedConfigurationInputSpec2525) *NullableBTComputedConfigurationInputSpec2525 {
	return &NullableBTComputedConfigurationInputSpec2525{value: val, isSet: true}
}

func (v NullableBTComputedConfigurationInputSpec2525) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTComputedConfigurationInputSpec2525) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
