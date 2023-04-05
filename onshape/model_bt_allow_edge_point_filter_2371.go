/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.161.13995-cdd961a1a6ad
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTAllowEdgePointFilter2371 struct for BTAllowEdgePointFilter2371
type BTAllowEdgePointFilter2371 struct {
	BtType          *string `json:"btType,omitempty"`
	AllowsEdgePoint *bool   `json:"allowsEdgePoint,omitempty"`
}

// NewBTAllowEdgePointFilter2371 instantiates a new BTAllowEdgePointFilter2371 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAllowEdgePointFilter2371() *BTAllowEdgePointFilter2371 {
	this := BTAllowEdgePointFilter2371{}
	return &this
}

// NewBTAllowEdgePointFilter2371WithDefaults instantiates a new BTAllowEdgePointFilter2371 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAllowEdgePointFilter2371WithDefaults() *BTAllowEdgePointFilter2371 {
	this := BTAllowEdgePointFilter2371{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTAllowEdgePointFilter2371) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAllowEdgePointFilter2371) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTAllowEdgePointFilter2371) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTAllowEdgePointFilter2371) SetBtType(v string) {
	o.BtType = &v
}

// GetAllowsEdgePoint returns the AllowsEdgePoint field value if set, zero value otherwise.
func (o *BTAllowEdgePointFilter2371) GetAllowsEdgePoint() bool {
	if o == nil || o.AllowsEdgePoint == nil {
		var ret bool
		return ret
	}
	return *o.AllowsEdgePoint
}

// GetAllowsEdgePointOk returns a tuple with the AllowsEdgePoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAllowEdgePointFilter2371) GetAllowsEdgePointOk() (*bool, bool) {
	if o == nil || o.AllowsEdgePoint == nil {
		return nil, false
	}
	return o.AllowsEdgePoint, true
}

// HasAllowsEdgePoint returns a boolean if a field has been set.
func (o *BTAllowEdgePointFilter2371) HasAllowsEdgePoint() bool {
	if o != nil && o.AllowsEdgePoint != nil {
		return true
	}

	return false
}

// SetAllowsEdgePoint gets a reference to the given bool and assigns it to the AllowsEdgePoint field.
func (o *BTAllowEdgePointFilter2371) SetAllowsEdgePoint(v bool) {
	o.AllowsEdgePoint = &v
}

func (o BTAllowEdgePointFilter2371) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.AllowsEdgePoint != nil {
		toSerialize["allowsEdgePoint"] = o.AllowsEdgePoint
	}
	return json.Marshal(toSerialize)
}

type NullableBTAllowEdgePointFilter2371 struct {
	value *BTAllowEdgePointFilter2371
	isSet bool
}

func (v NullableBTAllowEdgePointFilter2371) Get() *BTAllowEdgePointFilter2371 {
	return v.value
}

func (v *NullableBTAllowEdgePointFilter2371) Set(val *BTAllowEdgePointFilter2371) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAllowEdgePointFilter2371) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAllowEdgePointFilter2371) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAllowEdgePointFilter2371(val *BTAllowEdgePointFilter2371) *NullableBTAllowEdgePointFilter2371 {
	return &NullableBTAllowEdgePointFilter2371{value: val, isSet: true}
}

func (v NullableBTAllowEdgePointFilter2371) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAllowEdgePointFilter2371) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
