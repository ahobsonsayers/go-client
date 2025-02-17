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

// BTSketchObjectFilter184 struct for BTSketchObjectFilter184
type BTSketchObjectFilter184 struct {
	BtType         *string              `json:"btType,omitempty"`
	IsSketchObject *bool                `json:"isSketchObject,omitempty"`
	ObjectType     *GBTSketchObjectType `json:"objectType,omitempty"`
}

// NewBTSketchObjectFilter184 instantiates a new BTSketchObjectFilter184 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTSketchObjectFilter184() *BTSketchObjectFilter184 {
	this := BTSketchObjectFilter184{}
	return &this
}

// NewBTSketchObjectFilter184WithDefaults instantiates a new BTSketchObjectFilter184 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTSketchObjectFilter184WithDefaults() *BTSketchObjectFilter184 {
	this := BTSketchObjectFilter184{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTSketchObjectFilter184) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchObjectFilter184) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTSketchObjectFilter184) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTSketchObjectFilter184) SetBtType(v string) {
	o.BtType = &v
}

// GetIsSketchObject returns the IsSketchObject field value if set, zero value otherwise.
func (o *BTSketchObjectFilter184) GetIsSketchObject() bool {
	if o == nil || o.IsSketchObject == nil {
		var ret bool
		return ret
	}
	return *o.IsSketchObject
}

// GetIsSketchObjectOk returns a tuple with the IsSketchObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchObjectFilter184) GetIsSketchObjectOk() (*bool, bool) {
	if o == nil || o.IsSketchObject == nil {
		return nil, false
	}
	return o.IsSketchObject, true
}

// HasIsSketchObject returns a boolean if a field has been set.
func (o *BTSketchObjectFilter184) HasIsSketchObject() bool {
	if o != nil && o.IsSketchObject != nil {
		return true
	}

	return false
}

// SetIsSketchObject gets a reference to the given bool and assigns it to the IsSketchObject field.
func (o *BTSketchObjectFilter184) SetIsSketchObject(v bool) {
	o.IsSketchObject = &v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *BTSketchObjectFilter184) GetObjectType() GBTSketchObjectType {
	if o == nil || o.ObjectType == nil {
		var ret GBTSketchObjectType
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchObjectFilter184) GetObjectTypeOk() (*GBTSketchObjectType, bool) {
	if o == nil || o.ObjectType == nil {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *BTSketchObjectFilter184) HasObjectType() bool {
	if o != nil && o.ObjectType != nil {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given GBTSketchObjectType and assigns it to the ObjectType field.
func (o *BTSketchObjectFilter184) SetObjectType(v GBTSketchObjectType) {
	o.ObjectType = &v
}

func (o BTSketchObjectFilter184) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.IsSketchObject != nil {
		toSerialize["isSketchObject"] = o.IsSketchObject
	}
	if o.ObjectType != nil {
		toSerialize["objectType"] = o.ObjectType
	}
	return json.Marshal(toSerialize)
}

type NullableBTSketchObjectFilter184 struct {
	value *BTSketchObjectFilter184
	isSet bool
}

func (v NullableBTSketchObjectFilter184) Get() *BTSketchObjectFilter184 {
	return v.value
}

func (v *NullableBTSketchObjectFilter184) Set(val *BTSketchObjectFilter184) {
	v.value = val
	v.isSet = true
}

func (v NullableBTSketchObjectFilter184) IsSet() bool {
	return v.isSet
}

func (v *NullableBTSketchObjectFilter184) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTSketchObjectFilter184(val *BTSketchObjectFilter184) *NullableBTSketchObjectFilter184 {
	return &NullableBTSketchObjectFilter184{value: val, isSet: true}
}

func (v NullableBTSketchObjectFilter184) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTSketchObjectFilter184) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
