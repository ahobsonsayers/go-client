/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.151.5750-4f2542599dd4
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTGetJsonPathsResponse1544 struct for BTGetJsonPathsResponse1544
type BTGetJsonPathsResponse1544 struct {
	ChangeId *string             `json:"changeId,omitempty"`
	Results  [][]BTJsonMatch2290 `json:"results,omitempty"`
}

// NewBTGetJsonPathsResponse1544 instantiates a new BTGetJsonPathsResponse1544 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTGetJsonPathsResponse1544() *BTGetJsonPathsResponse1544 {
	this := BTGetJsonPathsResponse1544{}
	return &this
}

// NewBTGetJsonPathsResponse1544WithDefaults instantiates a new BTGetJsonPathsResponse1544 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTGetJsonPathsResponse1544WithDefaults() *BTGetJsonPathsResponse1544 {
	this := BTGetJsonPathsResponse1544{}
	return &this
}

// GetChangeId returns the ChangeId field value if set, zero value otherwise.
func (o *BTGetJsonPathsResponse1544) GetChangeId() string {
	if o == nil || o.ChangeId == nil {
		var ret string
		return ret
	}
	return *o.ChangeId
}

// GetChangeIdOk returns a tuple with the ChangeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTGetJsonPathsResponse1544) GetChangeIdOk() (*string, bool) {
	if o == nil || o.ChangeId == nil {
		return nil, false
	}
	return o.ChangeId, true
}

// HasChangeId returns a boolean if a field has been set.
func (o *BTGetJsonPathsResponse1544) HasChangeId() bool {
	if o != nil && o.ChangeId != nil {
		return true
	}

	return false
}

// SetChangeId gets a reference to the given string and assigns it to the ChangeId field.
func (o *BTGetJsonPathsResponse1544) SetChangeId(v string) {
	o.ChangeId = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *BTGetJsonPathsResponse1544) GetResults() [][]BTJsonMatch2290 {
	if o == nil || o.Results == nil {
		var ret [][]BTJsonMatch2290
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTGetJsonPathsResponse1544) GetResultsOk() ([][]BTJsonMatch2290, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *BTGetJsonPathsResponse1544) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given [][]BTJsonMatch2290 and assigns it to the Results field.
func (o *BTGetJsonPathsResponse1544) SetResults(v [][]BTJsonMatch2290) {
	o.Results = v
}

func (o BTGetJsonPathsResponse1544) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ChangeId != nil {
		toSerialize["changeId"] = o.ChangeId
	}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableBTGetJsonPathsResponse1544 struct {
	value *BTGetJsonPathsResponse1544
	isSet bool
}

func (v NullableBTGetJsonPathsResponse1544) Get() *BTGetJsonPathsResponse1544 {
	return v.value
}

func (v *NullableBTGetJsonPathsResponse1544) Set(val *BTGetJsonPathsResponse1544) {
	v.value = val
	v.isSet = true
}

func (v NullableBTGetJsonPathsResponse1544) IsSet() bool {
	return v.isSet
}

func (v *NullableBTGetJsonPathsResponse1544) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTGetJsonPathsResponse1544(val *BTGetJsonPathsResponse1544) *NullableBTGetJsonPathsResponse1544 {
	return &NullableBTGetJsonPathsResponse1544{value: val, isSet: true}
}

func (v NullableBTGetJsonPathsResponse1544) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTGetJsonPathsResponse1544) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
