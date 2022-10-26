/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.155.7070-d84f850e97a2
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTGetJsonResponse2137Tree struct for BTGetJsonResponse2137Tree
type BTGetJsonResponse2137Tree struct {
	BtType               *string `json:"btType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BTGetJsonResponse2137Tree BTGetJsonResponse2137Tree

// NewBTGetJsonResponse2137Tree instantiates a new BTGetJsonResponse2137Tree object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTGetJsonResponse2137Tree() *BTGetJsonResponse2137Tree {
	this := BTGetJsonResponse2137Tree{}
	return &this
}

// NewBTGetJsonResponse2137TreeWithDefaults instantiates a new BTGetJsonResponse2137Tree object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTGetJsonResponse2137TreeWithDefaults() *BTGetJsonResponse2137Tree {
	this := BTGetJsonResponse2137Tree{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTGetJsonResponse2137Tree) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTGetJsonResponse2137Tree) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTGetJsonResponse2137Tree) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTGetJsonResponse2137Tree) SetBtType(v string) {
	o.BtType = &v
}

func (o BTGetJsonResponse2137Tree) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BTGetJsonResponse2137Tree) UnmarshalJSON(bytes []byte) (err error) {
	varBTGetJsonResponse2137Tree := _BTGetJsonResponse2137Tree{}

	if err = json.Unmarshal(bytes, &varBTGetJsonResponse2137Tree); err == nil {
		*o = BTGetJsonResponse2137Tree(varBTGetJsonResponse2137Tree)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "btType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBTGetJsonResponse2137Tree struct {
	value *BTGetJsonResponse2137Tree
	isSet bool
}

func (v NullableBTGetJsonResponse2137Tree) Get() *BTGetJsonResponse2137Tree {
	return v.value
}

func (v *NullableBTGetJsonResponse2137Tree) Set(val *BTGetJsonResponse2137Tree) {
	v.value = val
	v.isSet = true
}

func (v NullableBTGetJsonResponse2137Tree) IsSet() bool {
	return v.isSet
}

func (v *NullableBTGetJsonResponse2137Tree) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTGetJsonResponse2137Tree(val *BTGetJsonResponse2137Tree) *NullableBTGetJsonResponse2137Tree {
	return &NullableBTGetJsonResponse2137Tree{value: val, isSet: true}
}

func (v NullableBTGetJsonResponse2137Tree) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTGetJsonResponse2137Tree) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
