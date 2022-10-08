/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.154.6820-1bef41f9cc03
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTConstructionObjectFilter113 struct for BTConstructionObjectFilter113
type BTConstructionObjectFilter113 struct {
	BtType         *string `json:"btType,omitempty"`
	IsConstruction *bool   `json:"isConstruction,omitempty"`
}

// NewBTConstructionObjectFilter113 instantiates a new BTConstructionObjectFilter113 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTConstructionObjectFilter113() *BTConstructionObjectFilter113 {
	this := BTConstructionObjectFilter113{}
	return &this
}

// NewBTConstructionObjectFilter113WithDefaults instantiates a new BTConstructionObjectFilter113 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTConstructionObjectFilter113WithDefaults() *BTConstructionObjectFilter113 {
	this := BTConstructionObjectFilter113{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTConstructionObjectFilter113) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConstructionObjectFilter113) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTConstructionObjectFilter113) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTConstructionObjectFilter113) SetBtType(v string) {
	o.BtType = &v
}

// GetIsConstruction returns the IsConstruction field value if set, zero value otherwise.
func (o *BTConstructionObjectFilter113) GetIsConstruction() bool {
	if o == nil || o.IsConstruction == nil {
		var ret bool
		return ret
	}
	return *o.IsConstruction
}

// GetIsConstructionOk returns a tuple with the IsConstruction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConstructionObjectFilter113) GetIsConstructionOk() (*bool, bool) {
	if o == nil || o.IsConstruction == nil {
		return nil, false
	}
	return o.IsConstruction, true
}

// HasIsConstruction returns a boolean if a field has been set.
func (o *BTConstructionObjectFilter113) HasIsConstruction() bool {
	if o != nil && o.IsConstruction != nil {
		return true
	}

	return false
}

// SetIsConstruction gets a reference to the given bool and assigns it to the IsConstruction field.
func (o *BTConstructionObjectFilter113) SetIsConstruction(v bool) {
	o.IsConstruction = &v
}

func (o BTConstructionObjectFilter113) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.IsConstruction != nil {
		toSerialize["isConstruction"] = o.IsConstruction
	}
	return json.Marshal(toSerialize)
}

type NullableBTConstructionObjectFilter113 struct {
	value *BTConstructionObjectFilter113
	isSet bool
}

func (v NullableBTConstructionObjectFilter113) Get() *BTConstructionObjectFilter113 {
	return v.value
}

func (v *NullableBTConstructionObjectFilter113) Set(val *BTConstructionObjectFilter113) {
	v.value = val
	v.isSet = true
}

func (v NullableBTConstructionObjectFilter113) IsSet() bool {
	return v.isSet
}

func (v *NullableBTConstructionObjectFilter113) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTConstructionObjectFilter113(val *BTConstructionObjectFilter113) *NullableBTConstructionObjectFilter113 {
	return &NullableBTConstructionObjectFilter113{value: val, isSet: true}
}

func (v NullableBTConstructionObjectFilter113) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTConstructionObjectFilter113) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
