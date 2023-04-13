/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.162.14462-13ace71ec1df
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTTableResponse1546 struct for BTTableResponse1546
type BTTableResponse1546 struct {
	SourceMicroversion *string      `json:"sourceMicroversion,omitempty"`
	Table              *BTTable1825 `json:"table,omitempty"`
}

// NewBTTableResponse1546 instantiates a new BTTableResponse1546 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTTableResponse1546() *BTTableResponse1546 {
	this := BTTableResponse1546{}
	return &this
}

// NewBTTableResponse1546WithDefaults instantiates a new BTTableResponse1546 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTTableResponse1546WithDefaults() *BTTableResponse1546 {
	this := BTTableResponse1546{}
	return &this
}

// GetSourceMicroversion returns the SourceMicroversion field value if set, zero value otherwise.
func (o *BTTableResponse1546) GetSourceMicroversion() string {
	if o == nil || o.SourceMicroversion == nil {
		var ret string
		return ret
	}
	return *o.SourceMicroversion
}

// GetSourceMicroversionOk returns a tuple with the SourceMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableResponse1546) GetSourceMicroversionOk() (*string, bool) {
	if o == nil || o.SourceMicroversion == nil {
		return nil, false
	}
	return o.SourceMicroversion, true
}

// HasSourceMicroversion returns a boolean if a field has been set.
func (o *BTTableResponse1546) HasSourceMicroversion() bool {
	if o != nil && o.SourceMicroversion != nil {
		return true
	}

	return false
}

// SetSourceMicroversion gets a reference to the given string and assigns it to the SourceMicroversion field.
func (o *BTTableResponse1546) SetSourceMicroversion(v string) {
	o.SourceMicroversion = &v
}

// GetTable returns the Table field value if set, zero value otherwise.
func (o *BTTableResponse1546) GetTable() BTTable1825 {
	if o == nil || o.Table == nil {
		var ret BTTable1825
		return ret
	}
	return *o.Table
}

// GetTableOk returns a tuple with the Table field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableResponse1546) GetTableOk() (*BTTable1825, bool) {
	if o == nil || o.Table == nil {
		return nil, false
	}
	return o.Table, true
}

// HasTable returns a boolean if a field has been set.
func (o *BTTableResponse1546) HasTable() bool {
	if o != nil && o.Table != nil {
		return true
	}

	return false
}

// SetTable gets a reference to the given BTTable1825 and assigns it to the Table field.
func (o *BTTableResponse1546) SetTable(v BTTable1825) {
	o.Table = &v
}

func (o BTTableResponse1546) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SourceMicroversion != nil {
		toSerialize["sourceMicroversion"] = o.SourceMicroversion
	}
	if o.Table != nil {
		toSerialize["table"] = o.Table
	}
	return json.Marshal(toSerialize)
}

type NullableBTTableResponse1546 struct {
	value *BTTableResponse1546
	isSet bool
}

func (v NullableBTTableResponse1546) Get() *BTTableResponse1546 {
	return v.value
}

func (v *NullableBTTableResponse1546) Set(val *BTTableResponse1546) {
	v.value = val
	v.isSet = true
}

func (v NullableBTTableResponse1546) IsSet() bool {
	return v.isSet
}

func (v *NullableBTTableResponse1546) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTTableResponse1546(val *BTTableResponse1546) *NullableBTTableResponse1546 {
	return &NullableBTTableResponse1546{value: val, isSet: true}
}

func (v NullableBTTableResponse1546) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTTableResponse1546) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
