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

// BTGeometryFilter130 struct for BTGeometryFilter130
type BTGeometryFilter130 struct {
	BtType       *string          `json:"btType,omitempty"`
	GeometryType *GBTGeometryType `json:"geometryType,omitempty"`
}

// NewBTGeometryFilter130 instantiates a new BTGeometryFilter130 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTGeometryFilter130() *BTGeometryFilter130 {
	this := BTGeometryFilter130{}
	return &this
}

// NewBTGeometryFilter130WithDefaults instantiates a new BTGeometryFilter130 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTGeometryFilter130WithDefaults() *BTGeometryFilter130 {
	this := BTGeometryFilter130{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTGeometryFilter130) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTGeometryFilter130) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTGeometryFilter130) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTGeometryFilter130) SetBtType(v string) {
	o.BtType = &v
}

// GetGeometryType returns the GeometryType field value if set, zero value otherwise.
func (o *BTGeometryFilter130) GetGeometryType() GBTGeometryType {
	if o == nil || o.GeometryType == nil {
		var ret GBTGeometryType
		return ret
	}
	return *o.GeometryType
}

// GetGeometryTypeOk returns a tuple with the GeometryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTGeometryFilter130) GetGeometryTypeOk() (*GBTGeometryType, bool) {
	if o == nil || o.GeometryType == nil {
		return nil, false
	}
	return o.GeometryType, true
}

// HasGeometryType returns a boolean if a field has been set.
func (o *BTGeometryFilter130) HasGeometryType() bool {
	if o != nil && o.GeometryType != nil {
		return true
	}

	return false
}

// SetGeometryType gets a reference to the given GBTGeometryType and assigns it to the GeometryType field.
func (o *BTGeometryFilter130) SetGeometryType(v GBTGeometryType) {
	o.GeometryType = &v
}

func (o BTGeometryFilter130) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.GeometryType != nil {
		toSerialize["geometryType"] = o.GeometryType
	}
	return json.Marshal(toSerialize)
}

type NullableBTGeometryFilter130 struct {
	value *BTGeometryFilter130
	isSet bool
}

func (v NullableBTGeometryFilter130) Get() *BTGeometryFilter130 {
	return v.value
}

func (v *NullableBTGeometryFilter130) Set(val *BTGeometryFilter130) {
	v.value = val
	v.isSet = true
}

func (v NullableBTGeometryFilter130) IsSet() bool {
	return v.isSet
}

func (v *NullableBTGeometryFilter130) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTGeometryFilter130(val *BTGeometryFilter130) *NullableBTGeometryFilter130 {
	return &NullableBTGeometryFilter130{value: val, isSet: true}
}

func (v NullableBTGeometryFilter130) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTGeometryFilter130) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
