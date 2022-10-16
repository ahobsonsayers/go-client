/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.155.6926-35d1d6168263
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTFeatureTypeFilter962 struct for BTFeatureTypeFilter962
type BTFeatureTypeFilter962 struct {
	BtType      *string `json:"btType,omitempty"`
	FeatureType *string `json:"featureType,omitempty"`
}

// NewBTFeatureTypeFilter962 instantiates a new BTFeatureTypeFilter962 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTFeatureTypeFilter962() *BTFeatureTypeFilter962 {
	this := BTFeatureTypeFilter962{}
	return &this
}

// NewBTFeatureTypeFilter962WithDefaults instantiates a new BTFeatureTypeFilter962 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTFeatureTypeFilter962WithDefaults() *BTFeatureTypeFilter962 {
	this := BTFeatureTypeFilter962{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTFeatureTypeFilter962) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureTypeFilter962) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTFeatureTypeFilter962) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTFeatureTypeFilter962) SetBtType(v string) {
	o.BtType = &v
}

// GetFeatureType returns the FeatureType field value if set, zero value otherwise.
func (o *BTFeatureTypeFilter962) GetFeatureType() string {
	if o == nil || o.FeatureType == nil {
		var ret string
		return ret
	}
	return *o.FeatureType
}

// GetFeatureTypeOk returns a tuple with the FeatureType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureTypeFilter962) GetFeatureTypeOk() (*string, bool) {
	if o == nil || o.FeatureType == nil {
		return nil, false
	}
	return o.FeatureType, true
}

// HasFeatureType returns a boolean if a field has been set.
func (o *BTFeatureTypeFilter962) HasFeatureType() bool {
	if o != nil && o.FeatureType != nil {
		return true
	}

	return false
}

// SetFeatureType gets a reference to the given string and assigns it to the FeatureType field.
func (o *BTFeatureTypeFilter962) SetFeatureType(v string) {
	o.FeatureType = &v
}

func (o BTFeatureTypeFilter962) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.FeatureType != nil {
		toSerialize["featureType"] = o.FeatureType
	}
	return json.Marshal(toSerialize)
}

type NullableBTFeatureTypeFilter962 struct {
	value *BTFeatureTypeFilter962
	isSet bool
}

func (v NullableBTFeatureTypeFilter962) Get() *BTFeatureTypeFilter962 {
	return v.value
}

func (v *NullableBTFeatureTypeFilter962) Set(val *BTFeatureTypeFilter962) {
	v.value = val
	v.isSet = true
}

func (v NullableBTFeatureTypeFilter962) IsSet() bool {
	return v.isSet
}

func (v *NullableBTFeatureTypeFilter962) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTFeatureTypeFilter962(val *BTFeatureTypeFilter962) *NullableBTFeatureTypeFilter962 {
	return &NullableBTFeatureTypeFilter962{value: val, isSet: true}
}

func (v NullableBTFeatureTypeFilter962) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTFeatureTypeFilter962) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
