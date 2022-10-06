/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.154.6779-364d99ceba76
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTExportModelVertex858 struct for BTExportModelVertex858
type BTExportModelVertex858 struct {
	Id    *string        `json:"id,omitempty"`
	Point *BTVector3d389 `json:"point,omitempty"`
}

// NewBTExportModelVertex858 instantiates a new BTExportModelVertex858 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTExportModelVertex858() *BTExportModelVertex858 {
	this := BTExportModelVertex858{}
	return &this
}

// NewBTExportModelVertex858WithDefaults instantiates a new BTExportModelVertex858 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTExportModelVertex858WithDefaults() *BTExportModelVertex858 {
	this := BTExportModelVertex858{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTExportModelVertex858) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelVertex858) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTExportModelVertex858) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTExportModelVertex858) SetId(v string) {
	o.Id = &v
}

// GetPoint returns the Point field value if set, zero value otherwise.
func (o *BTExportModelVertex858) GetPoint() BTVector3d389 {
	if o == nil || o.Point == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.Point
}

// GetPointOk returns a tuple with the Point field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelVertex858) GetPointOk() (*BTVector3d389, bool) {
	if o == nil || o.Point == nil {
		return nil, false
	}
	return o.Point, true
}

// HasPoint returns a boolean if a field has been set.
func (o *BTExportModelVertex858) HasPoint() bool {
	if o != nil && o.Point != nil {
		return true
	}

	return false
}

// SetPoint gets a reference to the given BTVector3d389 and assigns it to the Point field.
func (o *BTExportModelVertex858) SetPoint(v BTVector3d389) {
	o.Point = &v
}

func (o BTExportModelVertex858) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Point != nil {
		toSerialize["point"] = o.Point
	}
	return json.Marshal(toSerialize)
}

type NullableBTExportModelVertex858 struct {
	value *BTExportModelVertex858
	isSet bool
}

func (v NullableBTExportModelVertex858) Get() *BTExportModelVertex858 {
	return v.value
}

func (v *NullableBTExportModelVertex858) Set(val *BTExportModelVertex858) {
	v.value = val
	v.isSet = true
}

func (v NullableBTExportModelVertex858) IsSet() bool {
	return v.isSet
}

func (v *NullableBTExportModelVertex858) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTExportModelVertex858(val *BTExportModelVertex858) *NullableBTExportModelVertex858 {
	return &NullableBTExportModelVertex858{value: val, isSet: true}
}

func (v NullableBTExportModelVertex858) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTExportModelVertex858) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
