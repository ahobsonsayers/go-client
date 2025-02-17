/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.166.18273-3025d52f81b7
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTLinkToLatestDocumentParams struct for BTLinkToLatestDocumentParams
type BTLinkToLatestDocumentParams struct {
	Elements []string `json:"elements,omitempty"`
}

// NewBTLinkToLatestDocumentParams instantiates a new BTLinkToLatestDocumentParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTLinkToLatestDocumentParams() *BTLinkToLatestDocumentParams {
	this := BTLinkToLatestDocumentParams{}
	return &this
}

// NewBTLinkToLatestDocumentParamsWithDefaults instantiates a new BTLinkToLatestDocumentParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTLinkToLatestDocumentParamsWithDefaults() *BTLinkToLatestDocumentParams {
	this := BTLinkToLatestDocumentParams{}
	return &this
}

// GetElements returns the Elements field value if set, zero value otherwise.
func (o *BTLinkToLatestDocumentParams) GetElements() []string {
	if o == nil || o.Elements == nil {
		var ret []string
		return ret
	}
	return o.Elements
}

// GetElementsOk returns a tuple with the Elements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLinkToLatestDocumentParams) GetElementsOk() ([]string, bool) {
	if o == nil || o.Elements == nil {
		return nil, false
	}
	return o.Elements, true
}

// HasElements returns a boolean if a field has been set.
func (o *BTLinkToLatestDocumentParams) HasElements() bool {
	if o != nil && o.Elements != nil {
		return true
	}

	return false
}

// SetElements gets a reference to the given []string and assigns it to the Elements field.
func (o *BTLinkToLatestDocumentParams) SetElements(v []string) {
	o.Elements = v
}

func (o BTLinkToLatestDocumentParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Elements != nil {
		toSerialize["elements"] = o.Elements
	}
	return json.Marshal(toSerialize)
}

type NullableBTLinkToLatestDocumentParams struct {
	value *BTLinkToLatestDocumentParams
	isSet bool
}

func (v NullableBTLinkToLatestDocumentParams) Get() *BTLinkToLatestDocumentParams {
	return v.value
}

func (v *NullableBTLinkToLatestDocumentParams) Set(val *BTLinkToLatestDocumentParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTLinkToLatestDocumentParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTLinkToLatestDocumentParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTLinkToLatestDocumentParams(val *BTLinkToLatestDocumentParams) *NullableBTLinkToLatestDocumentParams {
	return &NullableBTLinkToLatestDocumentParams{value: val, isSet: true}
}

func (v NullableBTLinkToLatestDocumentParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTLinkToLatestDocumentParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
