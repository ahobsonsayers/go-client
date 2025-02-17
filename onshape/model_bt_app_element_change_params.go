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

// BTAppElementChangeParams struct for BTAppElementChangeParams
type BTAppElementChangeParams struct {
	// This base64-encoded data is treated as a full representation of the sub-element's data and all deltas previously added will no longer be returned.
	BaseContent *string `json:"baseContent,omitempty"`
	// This base64-encoded data is a delta which the application can apply to the last added baseContent data.
	Delta *string `json:"delta,omitempty"`
	// The id of the subelement to edit. The value is determined by the application.
	SubelementId string `json:"subelementId"`
}

// NewBTAppElementChangeParams instantiates a new BTAppElementChangeParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAppElementChangeParams(subelementId string) *BTAppElementChangeParams {
	this := BTAppElementChangeParams{}
	this.SubelementId = subelementId
	return &this
}

// NewBTAppElementChangeParamsWithDefaults instantiates a new BTAppElementChangeParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAppElementChangeParamsWithDefaults() *BTAppElementChangeParams {
	this := BTAppElementChangeParams{}
	return &this
}

// GetBaseContent returns the BaseContent field value if set, zero value otherwise.
func (o *BTAppElementChangeParams) GetBaseContent() string {
	if o == nil || o.BaseContent == nil {
		var ret string
		return ret
	}
	return *o.BaseContent
}

// GetBaseContentOk returns a tuple with the BaseContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementChangeParams) GetBaseContentOk() (*string, bool) {
	if o == nil || o.BaseContent == nil {
		return nil, false
	}
	return o.BaseContent, true
}

// HasBaseContent returns a boolean if a field has been set.
func (o *BTAppElementChangeParams) HasBaseContent() bool {
	if o != nil && o.BaseContent != nil {
		return true
	}

	return false
}

// SetBaseContent gets a reference to the given string and assigns it to the BaseContent field.
func (o *BTAppElementChangeParams) SetBaseContent(v string) {
	o.BaseContent = &v
}

// GetDelta returns the Delta field value if set, zero value otherwise.
func (o *BTAppElementChangeParams) GetDelta() string {
	if o == nil || o.Delta == nil {
		var ret string
		return ret
	}
	return *o.Delta
}

// GetDeltaOk returns a tuple with the Delta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementChangeParams) GetDeltaOk() (*string, bool) {
	if o == nil || o.Delta == nil {
		return nil, false
	}
	return o.Delta, true
}

// HasDelta returns a boolean if a field has been set.
func (o *BTAppElementChangeParams) HasDelta() bool {
	if o != nil && o.Delta != nil {
		return true
	}

	return false
}

// SetDelta gets a reference to the given string and assigns it to the Delta field.
func (o *BTAppElementChangeParams) SetDelta(v string) {
	o.Delta = &v
}

// GetSubelementId returns the SubelementId field value
func (o *BTAppElementChangeParams) GetSubelementId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubelementId
}

// GetSubelementIdOk returns a tuple with the SubelementId field value
// and a boolean to check if the value has been set.
func (o *BTAppElementChangeParams) GetSubelementIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubelementId, true
}

// SetSubelementId sets field value
func (o *BTAppElementChangeParams) SetSubelementId(v string) {
	o.SubelementId = v
}

func (o BTAppElementChangeParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BaseContent != nil {
		toSerialize["baseContent"] = o.BaseContent
	}
	if o.Delta != nil {
		toSerialize["delta"] = o.Delta
	}
	if true {
		toSerialize["subelementId"] = o.SubelementId
	}
	return json.Marshal(toSerialize)
}

type NullableBTAppElementChangeParams struct {
	value *BTAppElementChangeParams
	isSet bool
}

func (v NullableBTAppElementChangeParams) Get() *BTAppElementChangeParams {
	return v.value
}

func (v *NullableBTAppElementChangeParams) Set(val *BTAppElementChangeParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAppElementChangeParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAppElementChangeParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAppElementChangeParams(val *BTAppElementChangeParams) *NullableBTAppElementChangeParams {
	return &NullableBTAppElementChangeParams{value: val, isSet: true}
}

func (v NullableBTAppElementChangeParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAppElementChangeParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
