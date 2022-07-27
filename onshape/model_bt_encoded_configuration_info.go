/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.151.5726-0a2a2c07c0aa
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTEncodedConfigurationInfo struct for BTEncodedConfigurationInfo
type BTEncodedConfigurationInfo struct {
	EncodedId  *string `json:"encodedId,omitempty"`
	QueryParam *string `json:"queryParam,omitempty"`
}

// NewBTEncodedConfigurationInfo instantiates a new BTEncodedConfigurationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTEncodedConfigurationInfo() *BTEncodedConfigurationInfo {
	this := BTEncodedConfigurationInfo{}
	return &this
}

// NewBTEncodedConfigurationInfoWithDefaults instantiates a new BTEncodedConfigurationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTEncodedConfigurationInfoWithDefaults() *BTEncodedConfigurationInfo {
	this := BTEncodedConfigurationInfo{}
	return &this
}

// GetEncodedId returns the EncodedId field value if set, zero value otherwise.
func (o *BTEncodedConfigurationInfo) GetEncodedId() string {
	if o == nil || o.EncodedId == nil {
		var ret string
		return ret
	}
	return *o.EncodedId
}

// GetEncodedIdOk returns a tuple with the EncodedId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEncodedConfigurationInfo) GetEncodedIdOk() (*string, bool) {
	if o == nil || o.EncodedId == nil {
		return nil, false
	}
	return o.EncodedId, true
}

// HasEncodedId returns a boolean if a field has been set.
func (o *BTEncodedConfigurationInfo) HasEncodedId() bool {
	if o != nil && o.EncodedId != nil {
		return true
	}

	return false
}

// SetEncodedId gets a reference to the given string and assigns it to the EncodedId field.
func (o *BTEncodedConfigurationInfo) SetEncodedId(v string) {
	o.EncodedId = &v
}

// GetQueryParam returns the QueryParam field value if set, zero value otherwise.
func (o *BTEncodedConfigurationInfo) GetQueryParam() string {
	if o == nil || o.QueryParam == nil {
		var ret string
		return ret
	}
	return *o.QueryParam
}

// GetQueryParamOk returns a tuple with the QueryParam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEncodedConfigurationInfo) GetQueryParamOk() (*string, bool) {
	if o == nil || o.QueryParam == nil {
		return nil, false
	}
	return o.QueryParam, true
}

// HasQueryParam returns a boolean if a field has been set.
func (o *BTEncodedConfigurationInfo) HasQueryParam() bool {
	if o != nil && o.QueryParam != nil {
		return true
	}

	return false
}

// SetQueryParam gets a reference to the given string and assigns it to the QueryParam field.
func (o *BTEncodedConfigurationInfo) SetQueryParam(v string) {
	o.QueryParam = &v
}

func (o BTEncodedConfigurationInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EncodedId != nil {
		toSerialize["encodedId"] = o.EncodedId
	}
	if o.QueryParam != nil {
		toSerialize["queryParam"] = o.QueryParam
	}
	return json.Marshal(toSerialize)
}

type NullableBTEncodedConfigurationInfo struct {
	value *BTEncodedConfigurationInfo
	isSet bool
}

func (v NullableBTEncodedConfigurationInfo) Get() *BTEncodedConfigurationInfo {
	return v.value
}

func (v *NullableBTEncodedConfigurationInfo) Set(val *BTEncodedConfigurationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTEncodedConfigurationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTEncodedConfigurationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTEncodedConfigurationInfo(val *BTEncodedConfigurationInfo) *NullableBTEncodedConfigurationInfo {
	return &NullableBTEncodedConfigurationInfo{value: val, isSet: true}
}

func (v NullableBTEncodedConfigurationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTEncodedConfigurationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
