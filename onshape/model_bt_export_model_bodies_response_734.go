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

// BTExportModelBodiesResponse734 struct for BTExportModelBodiesResponse734
type BTExportModelBodiesResponse734 struct {
	Bodies                     []BTExportModelBody1272                 `json:"bodies,omitempty"`
	BtType                     *string                                 `json:"btType,omitempty"`
	ErrorEnum                  *GBTErrorStringEnum                     `json:"errorEnum,omitempty"`
	MicroversionId             *BTMicroversionId366                    `json:"microversionId,omitempty"`
	NodeIdToReferencedProperty *map[string]BTExportModelProperties3216 `json:"nodeIdToReferencedProperty,omitempty"`
}

// NewBTExportModelBodiesResponse734 instantiates a new BTExportModelBodiesResponse734 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTExportModelBodiesResponse734() *BTExportModelBodiesResponse734 {
	this := BTExportModelBodiesResponse734{}
	return &this
}

// NewBTExportModelBodiesResponse734WithDefaults instantiates a new BTExportModelBodiesResponse734 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTExportModelBodiesResponse734WithDefaults() *BTExportModelBodiesResponse734 {
	this := BTExportModelBodiesResponse734{}
	return &this
}

// GetBodies returns the Bodies field value if set, zero value otherwise.
func (o *BTExportModelBodiesResponse734) GetBodies() []BTExportModelBody1272 {
	if o == nil || o.Bodies == nil {
		var ret []BTExportModelBody1272
		return ret
	}
	return o.Bodies
}

// GetBodiesOk returns a tuple with the Bodies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelBodiesResponse734) GetBodiesOk() ([]BTExportModelBody1272, bool) {
	if o == nil || o.Bodies == nil {
		return nil, false
	}
	return o.Bodies, true
}

// HasBodies returns a boolean if a field has been set.
func (o *BTExportModelBodiesResponse734) HasBodies() bool {
	if o != nil && o.Bodies != nil {
		return true
	}

	return false
}

// SetBodies gets a reference to the given []BTExportModelBody1272 and assigns it to the Bodies field.
func (o *BTExportModelBodiesResponse734) SetBodies(v []BTExportModelBody1272) {
	o.Bodies = v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTExportModelBodiesResponse734) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelBodiesResponse734) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTExportModelBodiesResponse734) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTExportModelBodiesResponse734) SetBtType(v string) {
	o.BtType = &v
}

// GetErrorEnum returns the ErrorEnum field value if set, zero value otherwise.
func (o *BTExportModelBodiesResponse734) GetErrorEnum() GBTErrorStringEnum {
	if o == nil || o.ErrorEnum == nil {
		var ret GBTErrorStringEnum
		return ret
	}
	return *o.ErrorEnum
}

// GetErrorEnumOk returns a tuple with the ErrorEnum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelBodiesResponse734) GetErrorEnumOk() (*GBTErrorStringEnum, bool) {
	if o == nil || o.ErrorEnum == nil {
		return nil, false
	}
	return o.ErrorEnum, true
}

// HasErrorEnum returns a boolean if a field has been set.
func (o *BTExportModelBodiesResponse734) HasErrorEnum() bool {
	if o != nil && o.ErrorEnum != nil {
		return true
	}

	return false
}

// SetErrorEnum gets a reference to the given GBTErrorStringEnum and assigns it to the ErrorEnum field.
func (o *BTExportModelBodiesResponse734) SetErrorEnum(v GBTErrorStringEnum) {
	o.ErrorEnum = &v
}

// GetMicroversionId returns the MicroversionId field value if set, zero value otherwise.
func (o *BTExportModelBodiesResponse734) GetMicroversionId() BTMicroversionId366 {
	if o == nil || o.MicroversionId == nil {
		var ret BTMicroversionId366
		return ret
	}
	return *o.MicroversionId
}

// GetMicroversionIdOk returns a tuple with the MicroversionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelBodiesResponse734) GetMicroversionIdOk() (*BTMicroversionId366, bool) {
	if o == nil || o.MicroversionId == nil {
		return nil, false
	}
	return o.MicroversionId, true
}

// HasMicroversionId returns a boolean if a field has been set.
func (o *BTExportModelBodiesResponse734) HasMicroversionId() bool {
	if o != nil && o.MicroversionId != nil {
		return true
	}

	return false
}

// SetMicroversionId gets a reference to the given BTMicroversionId366 and assigns it to the MicroversionId field.
func (o *BTExportModelBodiesResponse734) SetMicroversionId(v BTMicroversionId366) {
	o.MicroversionId = &v
}

// GetNodeIdToReferencedProperty returns the NodeIdToReferencedProperty field value if set, zero value otherwise.
func (o *BTExportModelBodiesResponse734) GetNodeIdToReferencedProperty() map[string]BTExportModelProperties3216 {
	if o == nil || o.NodeIdToReferencedProperty == nil {
		var ret map[string]BTExportModelProperties3216
		return ret
	}
	return *o.NodeIdToReferencedProperty
}

// GetNodeIdToReferencedPropertyOk returns a tuple with the NodeIdToReferencedProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelBodiesResponse734) GetNodeIdToReferencedPropertyOk() (*map[string]BTExportModelProperties3216, bool) {
	if o == nil || o.NodeIdToReferencedProperty == nil {
		return nil, false
	}
	return o.NodeIdToReferencedProperty, true
}

// HasNodeIdToReferencedProperty returns a boolean if a field has been set.
func (o *BTExportModelBodiesResponse734) HasNodeIdToReferencedProperty() bool {
	if o != nil && o.NodeIdToReferencedProperty != nil {
		return true
	}

	return false
}

// SetNodeIdToReferencedProperty gets a reference to the given map[string]BTExportModelProperties3216 and assigns it to the NodeIdToReferencedProperty field.
func (o *BTExportModelBodiesResponse734) SetNodeIdToReferencedProperty(v map[string]BTExportModelProperties3216) {
	o.NodeIdToReferencedProperty = &v
}

func (o BTExportModelBodiesResponse734) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Bodies != nil {
		toSerialize["bodies"] = o.Bodies
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ErrorEnum != nil {
		toSerialize["errorEnum"] = o.ErrorEnum
	}
	if o.MicroversionId != nil {
		toSerialize["microversionId"] = o.MicroversionId
	}
	if o.NodeIdToReferencedProperty != nil {
		toSerialize["nodeIdToReferencedProperty"] = o.NodeIdToReferencedProperty
	}
	return json.Marshal(toSerialize)
}

type NullableBTExportModelBodiesResponse734 struct {
	value *BTExportModelBodiesResponse734
	isSet bool
}

func (v NullableBTExportModelBodiesResponse734) Get() *BTExportModelBodiesResponse734 {
	return v.value
}

func (v *NullableBTExportModelBodiesResponse734) Set(val *BTExportModelBodiesResponse734) {
	v.value = val
	v.isSet = true
}

func (v NullableBTExportModelBodiesResponse734) IsSet() bool {
	return v.isSet
}

func (v *NullableBTExportModelBodiesResponse734) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTExportModelBodiesResponse734(val *BTExportModelBodiesResponse734) *NullableBTExportModelBodiesResponse734 {
	return &NullableBTExportModelBodiesResponse734{value: val, isSet: true}
}

func (v NullableBTExportModelBodiesResponse734) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTExportModelBodiesResponse734) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
