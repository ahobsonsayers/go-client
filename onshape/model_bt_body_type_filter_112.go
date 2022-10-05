/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.154.6764-8cd829594e49
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTBodyTypeFilter112 struct for BTBodyTypeFilter112
type BTBodyTypeFilter112 struct {
	BtType   *string `json:"btType,omitempty"`
	BodyType *string `json:"bodyType,omitempty"`
}

// NewBTBodyTypeFilter112 instantiates a new BTBodyTypeFilter112 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTBodyTypeFilter112() *BTBodyTypeFilter112 {
	this := BTBodyTypeFilter112{}
	return &this
}

// NewBTBodyTypeFilter112WithDefaults instantiates a new BTBodyTypeFilter112 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTBodyTypeFilter112WithDefaults() *BTBodyTypeFilter112 {
	this := BTBodyTypeFilter112{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTBodyTypeFilter112) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBodyTypeFilter112) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTBodyTypeFilter112) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTBodyTypeFilter112) SetBtType(v string) {
	o.BtType = &v
}

// GetBodyType returns the BodyType field value if set, zero value otherwise.
func (o *BTBodyTypeFilter112) GetBodyType() string {
	if o == nil || o.BodyType == nil {
		var ret string
		return ret
	}
	return *o.BodyType
}

// GetBodyTypeOk returns a tuple with the BodyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBodyTypeFilter112) GetBodyTypeOk() (*string, bool) {
	if o == nil || o.BodyType == nil {
		return nil, false
	}
	return o.BodyType, true
}

// HasBodyType returns a boolean if a field has been set.
func (o *BTBodyTypeFilter112) HasBodyType() bool {
	if o != nil && o.BodyType != nil {
		return true
	}

	return false
}

// SetBodyType gets a reference to the given string and assigns it to the BodyType field.
func (o *BTBodyTypeFilter112) SetBodyType(v string) {
	o.BodyType = &v
}

func (o BTBodyTypeFilter112) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.BodyType != nil {
		toSerialize["bodyType"] = o.BodyType
	}
	return json.Marshal(toSerialize)
}

type NullableBTBodyTypeFilter112 struct {
	value *BTBodyTypeFilter112
	isSet bool
}

func (v NullableBTBodyTypeFilter112) Get() *BTBodyTypeFilter112 {
	return v.value
}

func (v *NullableBTBodyTypeFilter112) Set(val *BTBodyTypeFilter112) {
	v.value = val
	v.isSet = true
}

func (v NullableBTBodyTypeFilter112) IsSet() bool {
	return v.isSet
}

func (v *NullableBTBodyTypeFilter112) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTBodyTypeFilter112(val *BTBodyTypeFilter112) *NullableBTBodyTypeFilter112 {
	return &NullableBTBodyTypeFilter112{value: val, isSet: true}
}

func (v NullableBTBodyTypeFilter112) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTBodyTypeFilter112) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
