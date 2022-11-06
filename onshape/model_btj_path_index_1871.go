/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.156.7203-6065bde47225
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTJPathIndex1871 Identifies a value in a json array. For insert and move edit destinations, -1 can be used to indicate 'end'.
type BTJPathIndex1871 struct {
	BtType *string `json:"btType,omitempty"`
	Index  *int32  `json:"index,omitempty"`
}

// NewBTJPathIndex1871 instantiates a new BTJPathIndex1871 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTJPathIndex1871() *BTJPathIndex1871 {
	this := BTJPathIndex1871{}
	return &this
}

// NewBTJPathIndex1871WithDefaults instantiates a new BTJPathIndex1871 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTJPathIndex1871WithDefaults() *BTJPathIndex1871 {
	this := BTJPathIndex1871{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTJPathIndex1871) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTJPathIndex1871) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTJPathIndex1871) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTJPathIndex1871) SetBtType(v string) {
	o.BtType = &v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *BTJPathIndex1871) GetIndex() int32 {
	if o == nil || o.Index == nil {
		var ret int32
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTJPathIndex1871) GetIndexOk() (*int32, bool) {
	if o == nil || o.Index == nil {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *BTJPathIndex1871) HasIndex() bool {
	if o != nil && o.Index != nil {
		return true
	}

	return false
}

// SetIndex gets a reference to the given int32 and assigns it to the Index field.
func (o *BTJPathIndex1871) SetIndex(v int32) {
	o.Index = &v
}

func (o BTJPathIndex1871) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Index != nil {
		toSerialize["index"] = o.Index
	}
	return json.Marshal(toSerialize)
}

type NullableBTJPathIndex1871 struct {
	value *BTJPathIndex1871
	isSet bool
}

func (v NullableBTJPathIndex1871) Get() *BTJPathIndex1871 {
	return v.value
}

func (v *NullableBTJPathIndex1871) Set(val *BTJPathIndex1871) {
	v.value = val
	v.isSet = true
}

func (v NullableBTJPathIndex1871) IsSet() bool {
	return v.isSet
}

func (v *NullableBTJPathIndex1871) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTJPathIndex1871(val *BTJPathIndex1871) *NullableBTJPathIndex1871 {
	return &NullableBTJPathIndex1871{value: val, isSet: true}
}

func (v NullableBTJPathIndex1871) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTJPathIndex1871) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
