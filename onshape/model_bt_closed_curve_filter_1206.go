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

// BTClosedCurveFilter1206 struct for BTClosedCurveFilter1206
type BTClosedCurveFilter1206 struct {
	BtType   *string `json:"btType,omitempty"`
	IsClosed *bool   `json:"isClosed,omitempty"`
}

// NewBTClosedCurveFilter1206 instantiates a new BTClosedCurveFilter1206 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTClosedCurveFilter1206() *BTClosedCurveFilter1206 {
	this := BTClosedCurveFilter1206{}
	return &this
}

// NewBTClosedCurveFilter1206WithDefaults instantiates a new BTClosedCurveFilter1206 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTClosedCurveFilter1206WithDefaults() *BTClosedCurveFilter1206 {
	this := BTClosedCurveFilter1206{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTClosedCurveFilter1206) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTClosedCurveFilter1206) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTClosedCurveFilter1206) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTClosedCurveFilter1206) SetBtType(v string) {
	o.BtType = &v
}

// GetIsClosed returns the IsClosed field value if set, zero value otherwise.
func (o *BTClosedCurveFilter1206) GetIsClosed() bool {
	if o == nil || o.IsClosed == nil {
		var ret bool
		return ret
	}
	return *o.IsClosed
}

// GetIsClosedOk returns a tuple with the IsClosed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTClosedCurveFilter1206) GetIsClosedOk() (*bool, bool) {
	if o == nil || o.IsClosed == nil {
		return nil, false
	}
	return o.IsClosed, true
}

// HasIsClosed returns a boolean if a field has been set.
func (o *BTClosedCurveFilter1206) HasIsClosed() bool {
	if o != nil && o.IsClosed != nil {
		return true
	}

	return false
}

// SetIsClosed gets a reference to the given bool and assigns it to the IsClosed field.
func (o *BTClosedCurveFilter1206) SetIsClosed(v bool) {
	o.IsClosed = &v
}

func (o BTClosedCurveFilter1206) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.IsClosed != nil {
		toSerialize["isClosed"] = o.IsClosed
	}
	return json.Marshal(toSerialize)
}

type NullableBTClosedCurveFilter1206 struct {
	value *BTClosedCurveFilter1206
	isSet bool
}

func (v NullableBTClosedCurveFilter1206) Get() *BTClosedCurveFilter1206 {
	return v.value
}

func (v *NullableBTClosedCurveFilter1206) Set(val *BTClosedCurveFilter1206) {
	v.value = val
	v.isSet = true
}

func (v NullableBTClosedCurveFilter1206) IsSet() bool {
	return v.isSet
}

func (v *NullableBTClosedCurveFilter1206) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTClosedCurveFilter1206(val *BTClosedCurveFilter1206) *NullableBTClosedCurveFilter1206 {
	return &NullableBTClosedCurveFilter1206{value: val, isSet: true}
}

func (v NullableBTClosedCurveFilter1206) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTClosedCurveFilter1206) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
