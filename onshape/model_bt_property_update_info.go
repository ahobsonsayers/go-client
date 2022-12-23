/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.157.9170-8cb36f16959d
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTPropertyUpdateInfo struct for BTPropertyUpdateInfo
type BTPropertyUpdateInfo struct {
	ErrorMessage *string                 `json:"errorMessage,omitempty"`
	ItemHref     *string                 `json:"itemHref,omitempty"`
	NewValue     *map[string]interface{} `json:"newValue,omitempty"`
	OldValue     *map[string]interface{} `json:"oldValue,omitempty"`
	PropertyId   *string                 `json:"propertyId,omitempty"`
}

// NewBTPropertyUpdateInfo instantiates a new BTPropertyUpdateInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPropertyUpdateInfo() *BTPropertyUpdateInfo {
	this := BTPropertyUpdateInfo{}
	return &this
}

// NewBTPropertyUpdateInfoWithDefaults instantiates a new BTPropertyUpdateInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPropertyUpdateInfoWithDefaults() *BTPropertyUpdateInfo {
	this := BTPropertyUpdateInfo{}
	return &this
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *BTPropertyUpdateInfo) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPropertyUpdateInfo) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *BTPropertyUpdateInfo) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage != nil {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *BTPropertyUpdateInfo) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetItemHref returns the ItemHref field value if set, zero value otherwise.
func (o *BTPropertyUpdateInfo) GetItemHref() string {
	if o == nil || o.ItemHref == nil {
		var ret string
		return ret
	}
	return *o.ItemHref
}

// GetItemHrefOk returns a tuple with the ItemHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPropertyUpdateInfo) GetItemHrefOk() (*string, bool) {
	if o == nil || o.ItemHref == nil {
		return nil, false
	}
	return o.ItemHref, true
}

// HasItemHref returns a boolean if a field has been set.
func (o *BTPropertyUpdateInfo) HasItemHref() bool {
	if o != nil && o.ItemHref != nil {
		return true
	}

	return false
}

// SetItemHref gets a reference to the given string and assigns it to the ItemHref field.
func (o *BTPropertyUpdateInfo) SetItemHref(v string) {
	o.ItemHref = &v
}

// GetNewValue returns the NewValue field value if set, zero value otherwise.
func (o *BTPropertyUpdateInfo) GetNewValue() map[string]interface{} {
	if o == nil || o.NewValue == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.NewValue
}

// GetNewValueOk returns a tuple with the NewValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPropertyUpdateInfo) GetNewValueOk() (*map[string]interface{}, bool) {
	if o == nil || o.NewValue == nil {
		return nil, false
	}
	return o.NewValue, true
}

// HasNewValue returns a boolean if a field has been set.
func (o *BTPropertyUpdateInfo) HasNewValue() bool {
	if o != nil && o.NewValue != nil {
		return true
	}

	return false
}

// SetNewValue gets a reference to the given map[string]interface{} and assigns it to the NewValue field.
func (o *BTPropertyUpdateInfo) SetNewValue(v map[string]interface{}) {
	o.NewValue = &v
}

// GetOldValue returns the OldValue field value if set, zero value otherwise.
func (o *BTPropertyUpdateInfo) GetOldValue() map[string]interface{} {
	if o == nil || o.OldValue == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.OldValue
}

// GetOldValueOk returns a tuple with the OldValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPropertyUpdateInfo) GetOldValueOk() (*map[string]interface{}, bool) {
	if o == nil || o.OldValue == nil {
		return nil, false
	}
	return o.OldValue, true
}

// HasOldValue returns a boolean if a field has been set.
func (o *BTPropertyUpdateInfo) HasOldValue() bool {
	if o != nil && o.OldValue != nil {
		return true
	}

	return false
}

// SetOldValue gets a reference to the given map[string]interface{} and assigns it to the OldValue field.
func (o *BTPropertyUpdateInfo) SetOldValue(v map[string]interface{}) {
	o.OldValue = &v
}

// GetPropertyId returns the PropertyId field value if set, zero value otherwise.
func (o *BTPropertyUpdateInfo) GetPropertyId() string {
	if o == nil || o.PropertyId == nil {
		var ret string
		return ret
	}
	return *o.PropertyId
}

// GetPropertyIdOk returns a tuple with the PropertyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPropertyUpdateInfo) GetPropertyIdOk() (*string, bool) {
	if o == nil || o.PropertyId == nil {
		return nil, false
	}
	return o.PropertyId, true
}

// HasPropertyId returns a boolean if a field has been set.
func (o *BTPropertyUpdateInfo) HasPropertyId() bool {
	if o != nil && o.PropertyId != nil {
		return true
	}

	return false
}

// SetPropertyId gets a reference to the given string and assigns it to the PropertyId field.
func (o *BTPropertyUpdateInfo) SetPropertyId(v string) {
	o.PropertyId = &v
}

func (o BTPropertyUpdateInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ErrorMessage != nil {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if o.ItemHref != nil {
		toSerialize["itemHref"] = o.ItemHref
	}
	if o.NewValue != nil {
		toSerialize["newValue"] = o.NewValue
	}
	if o.OldValue != nil {
		toSerialize["oldValue"] = o.OldValue
	}
	if o.PropertyId != nil {
		toSerialize["propertyId"] = o.PropertyId
	}
	return json.Marshal(toSerialize)
}

type NullableBTPropertyUpdateInfo struct {
	value *BTPropertyUpdateInfo
	isSet bool
}

func (v NullableBTPropertyUpdateInfo) Get() *BTPropertyUpdateInfo {
	return v.value
}

func (v *NullableBTPropertyUpdateInfo) Set(val *BTPropertyUpdateInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPropertyUpdateInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPropertyUpdateInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPropertyUpdateInfo(val *BTPropertyUpdateInfo) *NullableBTPropertyUpdateInfo {
	return &NullableBTPropertyUpdateInfo{value: val, isSet: true}
}

func (v NullableBTPropertyUpdateInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPropertyUpdateInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
