/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.161.14306-351f5b17f026
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTDocumentOpenEventParams struct for BTDocumentOpenEventParams
type BTDocumentOpenEventParams struct {
	JsonType   *string `json:"jsonType,omitempty"`
	DocumentId *string `json:"documentId,omitempty"`
}

// NewBTDocumentOpenEventParams instantiates a new BTDocumentOpenEventParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTDocumentOpenEventParams() *BTDocumentOpenEventParams {
	this := BTDocumentOpenEventParams{}
	return &this
}

// NewBTDocumentOpenEventParamsWithDefaults instantiates a new BTDocumentOpenEventParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTDocumentOpenEventParamsWithDefaults() *BTDocumentOpenEventParams {
	this := BTDocumentOpenEventParams{}
	return &this
}

// GetJsonType returns the JsonType field value if set, zero value otherwise.
func (o *BTDocumentOpenEventParams) GetJsonType() string {
	if o == nil || o.JsonType == nil {
		var ret string
		return ret
	}
	return *o.JsonType
}

// GetJsonTypeOk returns a tuple with the JsonType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentOpenEventParams) GetJsonTypeOk() (*string, bool) {
	if o == nil || o.JsonType == nil {
		return nil, false
	}
	return o.JsonType, true
}

// HasJsonType returns a boolean if a field has been set.
func (o *BTDocumentOpenEventParams) HasJsonType() bool {
	if o != nil && o.JsonType != nil {
		return true
	}

	return false
}

// SetJsonType gets a reference to the given string and assigns it to the JsonType field.
func (o *BTDocumentOpenEventParams) SetJsonType(v string) {
	o.JsonType = &v
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *BTDocumentOpenEventParams) GetDocumentId() string {
	if o == nil || o.DocumentId == nil {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentOpenEventParams) GetDocumentIdOk() (*string, bool) {
	if o == nil || o.DocumentId == nil {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *BTDocumentOpenEventParams) HasDocumentId() bool {
	if o != nil && o.DocumentId != nil {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *BTDocumentOpenEventParams) SetDocumentId(v string) {
	o.DocumentId = &v
}

func (o BTDocumentOpenEventParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.JsonType != nil {
		toSerialize["jsonType"] = o.JsonType
	}
	if o.DocumentId != nil {
		toSerialize["documentId"] = o.DocumentId
	}
	return json.Marshal(toSerialize)
}

type NullableBTDocumentOpenEventParams struct {
	value *BTDocumentOpenEventParams
	isSet bool
}

func (v NullableBTDocumentOpenEventParams) Get() *BTDocumentOpenEventParams {
	return v.value
}

func (v *NullableBTDocumentOpenEventParams) Set(val *BTDocumentOpenEventParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTDocumentOpenEventParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTDocumentOpenEventParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTDocumentOpenEventParams(val *BTDocumentOpenEventParams) *NullableBTDocumentOpenEventParams {
	return &NullableBTDocumentOpenEventParams{value: val, isSet: true}
}

func (v NullableBTDocumentOpenEventParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTDocumentOpenEventParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
