/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.155.7489-3fe01473c812
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTModelProperties1258 struct for BTModelProperties1258
type BTModelProperties1258 struct {
	BtType            *string                  `json:"btType,omitempty"`
	NodeId            *string                  `json:"nodeId,omitempty"`
	SubPartProperties []BTOnePartProperties230 `json:"subPartProperties,omitempty"`
}

// NewBTModelProperties1258 instantiates a new BTModelProperties1258 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTModelProperties1258() *BTModelProperties1258 {
	this := BTModelProperties1258{}
	return &this
}

// NewBTModelProperties1258WithDefaults instantiates a new BTModelProperties1258 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTModelProperties1258WithDefaults() *BTModelProperties1258 {
	this := BTModelProperties1258{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTModelProperties1258) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTModelProperties1258) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTModelProperties1258) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTModelProperties1258) SetBtType(v string) {
	o.BtType = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTModelProperties1258) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTModelProperties1258) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTModelProperties1258) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTModelProperties1258) SetNodeId(v string) {
	o.NodeId = &v
}

// GetSubPartProperties returns the SubPartProperties field value if set, zero value otherwise.
func (o *BTModelProperties1258) GetSubPartProperties() []BTOnePartProperties230 {
	if o == nil || o.SubPartProperties == nil {
		var ret []BTOnePartProperties230
		return ret
	}
	return o.SubPartProperties
}

// GetSubPartPropertiesOk returns a tuple with the SubPartProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTModelProperties1258) GetSubPartPropertiesOk() ([]BTOnePartProperties230, bool) {
	if o == nil || o.SubPartProperties == nil {
		return nil, false
	}
	return o.SubPartProperties, true
}

// HasSubPartProperties returns a boolean if a field has been set.
func (o *BTModelProperties1258) HasSubPartProperties() bool {
	if o != nil && o.SubPartProperties != nil {
		return true
	}

	return false
}

// SetSubPartProperties gets a reference to the given []BTOnePartProperties230 and assigns it to the SubPartProperties field.
func (o *BTModelProperties1258) SetSubPartProperties(v []BTOnePartProperties230) {
	o.SubPartProperties = v
}

func (o BTModelProperties1258) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.SubPartProperties != nil {
		toSerialize["subPartProperties"] = o.SubPartProperties
	}
	return json.Marshal(toSerialize)
}

type NullableBTModelProperties1258 struct {
	value *BTModelProperties1258
	isSet bool
}

func (v NullableBTModelProperties1258) Get() *BTModelProperties1258 {
	return v.value
}

func (v *NullableBTModelProperties1258) Set(val *BTModelProperties1258) {
	v.value = val
	v.isSet = true
}

func (v NullableBTModelProperties1258) IsSet() bool {
	return v.isSet
}

func (v *NullableBTModelProperties1258) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTModelProperties1258(val *BTModelProperties1258) *NullableBTModelProperties1258 {
	return &NullableBTModelProperties1258{value: val, isSet: true}
}

func (v NullableBTModelProperties1258) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTModelProperties1258) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
