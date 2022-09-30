/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.154.6700-f2963ce28904
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTApplicationElementThumbnailParamsArray struct for BTApplicationElementThumbnailParamsArray
type BTApplicationElementThumbnailParamsArray struct {
	Thumbnails []BTApplicationElementThumbnailParams `json:"thumbnails,omitempty"`
}

// NewBTApplicationElementThumbnailParamsArray instantiates a new BTApplicationElementThumbnailParamsArray object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTApplicationElementThumbnailParamsArray() *BTApplicationElementThumbnailParamsArray {
	this := BTApplicationElementThumbnailParamsArray{}
	return &this
}

// NewBTApplicationElementThumbnailParamsArrayWithDefaults instantiates a new BTApplicationElementThumbnailParamsArray object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTApplicationElementThumbnailParamsArrayWithDefaults() *BTApplicationElementThumbnailParamsArray {
	this := BTApplicationElementThumbnailParamsArray{}
	return &this
}

// GetThumbnails returns the Thumbnails field value if set, zero value otherwise.
func (o *BTApplicationElementThumbnailParamsArray) GetThumbnails() []BTApplicationElementThumbnailParams {
	if o == nil || o.Thumbnails == nil {
		var ret []BTApplicationElementThumbnailParams
		return ret
	}
	return o.Thumbnails
}

// GetThumbnailsOk returns a tuple with the Thumbnails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTApplicationElementThumbnailParamsArray) GetThumbnailsOk() ([]BTApplicationElementThumbnailParams, bool) {
	if o == nil || o.Thumbnails == nil {
		return nil, false
	}
	return o.Thumbnails, true
}

// HasThumbnails returns a boolean if a field has been set.
func (o *BTApplicationElementThumbnailParamsArray) HasThumbnails() bool {
	if o != nil && o.Thumbnails != nil {
		return true
	}

	return false
}

// SetThumbnails gets a reference to the given []BTApplicationElementThumbnailParams and assigns it to the Thumbnails field.
func (o *BTApplicationElementThumbnailParamsArray) SetThumbnails(v []BTApplicationElementThumbnailParams) {
	o.Thumbnails = v
}

func (o BTApplicationElementThumbnailParamsArray) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Thumbnails != nil {
		toSerialize["thumbnails"] = o.Thumbnails
	}
	return json.Marshal(toSerialize)
}

type NullableBTApplicationElementThumbnailParamsArray struct {
	value *BTApplicationElementThumbnailParamsArray
	isSet bool
}

func (v NullableBTApplicationElementThumbnailParamsArray) Get() *BTApplicationElementThumbnailParamsArray {
	return v.value
}

func (v *NullableBTApplicationElementThumbnailParamsArray) Set(val *BTApplicationElementThumbnailParamsArray) {
	v.value = val
	v.isSet = true
}

func (v NullableBTApplicationElementThumbnailParamsArray) IsSet() bool {
	return v.isSet
}

func (v *NullableBTApplicationElementThumbnailParamsArray) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTApplicationElementThumbnailParamsArray(val *BTApplicationElementThumbnailParamsArray) *NullableBTApplicationElementThumbnailParamsArray {
	return &NullableBTApplicationElementThumbnailParamsArray{value: val, isSet: true}
}

func (v NullableBTApplicationElementThumbnailParamsArray) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTApplicationElementThumbnailParamsArray) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
