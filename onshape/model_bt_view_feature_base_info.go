/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.154.6764-8cd829594e49
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTViewFeatureBaseInfo struct for BTViewFeatureBaseInfo
type BTViewFeatureBaseInfo struct {
	Id   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewBTViewFeatureBaseInfo instantiates a new BTViewFeatureBaseInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTViewFeatureBaseInfo() *BTViewFeatureBaseInfo {
	this := BTViewFeatureBaseInfo{}
	return &this
}

// NewBTViewFeatureBaseInfoWithDefaults instantiates a new BTViewFeatureBaseInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTViewFeatureBaseInfoWithDefaults() *BTViewFeatureBaseInfo {
	this := BTViewFeatureBaseInfo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTViewFeatureBaseInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTViewFeatureBaseInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTViewFeatureBaseInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTViewFeatureBaseInfo) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTViewFeatureBaseInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTViewFeatureBaseInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTViewFeatureBaseInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTViewFeatureBaseInfo) SetName(v string) {
	o.Name = &v
}

func (o BTViewFeatureBaseInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableBTViewFeatureBaseInfo struct {
	value *BTViewFeatureBaseInfo
	isSet bool
}

func (v NullableBTViewFeatureBaseInfo) Get() *BTViewFeatureBaseInfo {
	return v.value
}

func (v *NullableBTViewFeatureBaseInfo) Set(val *BTViewFeatureBaseInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTViewFeatureBaseInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTViewFeatureBaseInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTViewFeatureBaseInfo(val *BTViewFeatureBaseInfo) *NullableBTViewFeatureBaseInfo {
	return &NullableBTViewFeatureBaseInfo{value: val, isSet: true}
}

func (v NullableBTViewFeatureBaseInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTViewFeatureBaseInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
