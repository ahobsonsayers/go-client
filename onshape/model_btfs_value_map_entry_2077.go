/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.151.5884-a8034368dd2e
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTFSValueMapEntry2077 struct for BTFSValueMapEntry2077
type BTFSValueMapEntry2077 struct {
	Key   *BTFSValue1888 `json:"key,omitempty"`
	Value *BTFSValue1888 `json:"value,omitempty"`
}

// NewBTFSValueMapEntry2077 instantiates a new BTFSValueMapEntry2077 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTFSValueMapEntry2077() *BTFSValueMapEntry2077 {
	this := BTFSValueMapEntry2077{}
	return &this
}

// NewBTFSValueMapEntry2077WithDefaults instantiates a new BTFSValueMapEntry2077 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTFSValueMapEntry2077WithDefaults() *BTFSValueMapEntry2077 {
	this := BTFSValueMapEntry2077{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *BTFSValueMapEntry2077) GetKey() BTFSValue1888 {
	if o == nil || o.Key == nil {
		var ret BTFSValue1888
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFSValueMapEntry2077) GetKeyOk() (*BTFSValue1888, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *BTFSValueMapEntry2077) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given BTFSValue1888 and assigns it to the Key field.
func (o *BTFSValueMapEntry2077) SetKey(v BTFSValue1888) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *BTFSValueMapEntry2077) GetValue() BTFSValue1888 {
	if o == nil || o.Value == nil {
		var ret BTFSValue1888
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFSValueMapEntry2077) GetValueOk() (*BTFSValue1888, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *BTFSValueMapEntry2077) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given BTFSValue1888 and assigns it to the Value field.
func (o *BTFSValueMapEntry2077) SetValue(v BTFSValue1888) {
	o.Value = &v
}

func (o BTFSValueMapEntry2077) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableBTFSValueMapEntry2077 struct {
	value *BTFSValueMapEntry2077
	isSet bool
}

func (v NullableBTFSValueMapEntry2077) Get() *BTFSValueMapEntry2077 {
	return v.value
}

func (v *NullableBTFSValueMapEntry2077) Set(val *BTFSValueMapEntry2077) {
	v.value = val
	v.isSet = true
}

func (v NullableBTFSValueMapEntry2077) IsSet() bool {
	return v.isSet
}

func (v *NullableBTFSValueMapEntry2077) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTFSValueMapEntry2077(val *BTFSValueMapEntry2077) *NullableBTFSValueMapEntry2077 {
	return &NullableBTFSValueMapEntry2077{value: val, isSet: true}
}

func (v NullableBTFSValueMapEntry2077) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTFSValueMapEntry2077) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
