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

// Lines struct for Lines
type Lines struct {
}

// NewLines instantiates a new Lines object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLines() *Lines {
	this := Lines{}
	return &this
}

// NewLinesWithDefaults instantiates a new Lines object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinesWithDefaults() *Lines {
	this := Lines{}
	return &this
}

func (o Lines) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	return json.Marshal(toSerialize)
}

type NullableLines struct {
	value *Lines
	isSet bool
}

func (v NullableLines) Get() *Lines {
	return v.value
}

func (v *NullableLines) Set(val *Lines) {
	v.value = val
	v.isSet = true
}

func (v NullableLines) IsSet() bool {
	return v.isSet
}

func (v *NullableLines) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLines(val *Lines) *NullableLines {
	return &NullableLines{value: val, isSet: true}
}

func (v NullableLines) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLines) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
