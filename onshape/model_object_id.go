/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.164.16955-b4ecd192bba6
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// ObjectId struct for ObjectId
type ObjectId struct {
	Date      *JSONTime `json:"date,omitempty"`
	Timestamp *int32    `json:"timestamp,omitempty"`
}

// NewObjectId instantiates a new ObjectId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectId() *ObjectId {
	this := ObjectId{}
	return &this
}

// NewObjectIdWithDefaults instantiates a new ObjectId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectIdWithDefaults() *ObjectId {
	this := ObjectId{}
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *ObjectId) GetDate() JSONTime {
	if o == nil || o.Date == nil {
		var ret JSONTime
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectId) GetDateOk() (*JSONTime, bool) {
	if o == nil || o.Date == nil {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *ObjectId) HasDate() bool {
	if o != nil && o.Date != nil {
		return true
	}

	return false
}

// SetDate gets a reference to the given JSONTime and assigns it to the Date field.
func (o *ObjectId) SetDate(v JSONTime) {
	o.Date = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *ObjectId) GetTimestamp() int32 {
	if o == nil || o.Timestamp == nil {
		var ret int32
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectId) GetTimestampOk() (*int32, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ObjectId) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int32 and assigns it to the Timestamp field.
func (o *ObjectId) SetTimestamp(v int32) {
	o.Timestamp = &v
}

func (o ObjectId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Date != nil {
		toSerialize["date"] = o.Date
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	return json.Marshal(toSerialize)
}

type NullableObjectId struct {
	value *ObjectId
	isSet bool
}

func (v NullableObjectId) Get() *ObjectId {
	return v.value
}

func (v *NullableObjectId) Set(val *ObjectId) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectId) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectId(val *ObjectId) *NullableObjectId {
	return &NullableObjectId{value: val, isSet: true}
}

func (v NullableObjectId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
