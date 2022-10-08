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

// BTMicroversionIdAndConfiguration2338 struct for BTMicroversionIdAndConfiguration2338
type BTMicroversionIdAndConfiguration2338 struct {
	BtType                          *string                   `json:"btType,omitempty"`
	CacheKey                        *string                   `json:"cacheKey,omitempty"`
	ConfigurationParameterIdToValue *map[string]BTFSValue1888 `json:"configurationParameterIdToValue,omitempty"`
	Deleted                         *bool                     `json:"deleted,omitempty"`
	Description                     *string                   `json:"description,omitempty"`
	Microversion                    *BTMicroversionId366      `json:"microversion,omitempty"`
}

// NewBTMicroversionIdAndConfiguration2338 instantiates a new BTMicroversionIdAndConfiguration2338 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMicroversionIdAndConfiguration2338() *BTMicroversionIdAndConfiguration2338 {
	this := BTMicroversionIdAndConfiguration2338{}
	return &this
}

// NewBTMicroversionIdAndConfiguration2338WithDefaults instantiates a new BTMicroversionIdAndConfiguration2338 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMicroversionIdAndConfiguration2338WithDefaults() *BTMicroversionIdAndConfiguration2338 {
	this := BTMicroversionIdAndConfiguration2338{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMicroversionIdAndConfiguration2338) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMicroversionIdAndConfiguration2338) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMicroversionIdAndConfiguration2338) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMicroversionIdAndConfiguration2338) SetBtType(v string) {
	o.BtType = &v
}

// GetCacheKey returns the CacheKey field value if set, zero value otherwise.
func (o *BTMicroversionIdAndConfiguration2338) GetCacheKey() string {
	if o == nil || o.CacheKey == nil {
		var ret string
		return ret
	}
	return *o.CacheKey
}

// GetCacheKeyOk returns a tuple with the CacheKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMicroversionIdAndConfiguration2338) GetCacheKeyOk() (*string, bool) {
	if o == nil || o.CacheKey == nil {
		return nil, false
	}
	return o.CacheKey, true
}

// HasCacheKey returns a boolean if a field has been set.
func (o *BTMicroversionIdAndConfiguration2338) HasCacheKey() bool {
	if o != nil && o.CacheKey != nil {
		return true
	}

	return false
}

// SetCacheKey gets a reference to the given string and assigns it to the CacheKey field.
func (o *BTMicroversionIdAndConfiguration2338) SetCacheKey(v string) {
	o.CacheKey = &v
}

// GetConfigurationParameterIdToValue returns the ConfigurationParameterIdToValue field value if set, zero value otherwise.
func (o *BTMicroversionIdAndConfiguration2338) GetConfigurationParameterIdToValue() map[string]BTFSValue1888 {
	if o == nil || o.ConfigurationParameterIdToValue == nil {
		var ret map[string]BTFSValue1888
		return ret
	}
	return *o.ConfigurationParameterIdToValue
}

// GetConfigurationParameterIdToValueOk returns a tuple with the ConfigurationParameterIdToValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMicroversionIdAndConfiguration2338) GetConfigurationParameterIdToValueOk() (*map[string]BTFSValue1888, bool) {
	if o == nil || o.ConfigurationParameterIdToValue == nil {
		return nil, false
	}
	return o.ConfigurationParameterIdToValue, true
}

// HasConfigurationParameterIdToValue returns a boolean if a field has been set.
func (o *BTMicroversionIdAndConfiguration2338) HasConfigurationParameterIdToValue() bool {
	if o != nil && o.ConfigurationParameterIdToValue != nil {
		return true
	}

	return false
}

// SetConfigurationParameterIdToValue gets a reference to the given map[string]BTFSValue1888 and assigns it to the ConfigurationParameterIdToValue field.
func (o *BTMicroversionIdAndConfiguration2338) SetConfigurationParameterIdToValue(v map[string]BTFSValue1888) {
	o.ConfigurationParameterIdToValue = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *BTMicroversionIdAndConfiguration2338) GetDeleted() bool {
	if o == nil || o.Deleted == nil {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMicroversionIdAndConfiguration2338) GetDeletedOk() (*bool, bool) {
	if o == nil || o.Deleted == nil {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *BTMicroversionIdAndConfiguration2338) HasDeleted() bool {
	if o != nil && o.Deleted != nil {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *BTMicroversionIdAndConfiguration2338) SetDeleted(v bool) {
	o.Deleted = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BTMicroversionIdAndConfiguration2338) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMicroversionIdAndConfiguration2338) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BTMicroversionIdAndConfiguration2338) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BTMicroversionIdAndConfiguration2338) SetDescription(v string) {
	o.Description = &v
}

// GetMicroversion returns the Microversion field value if set, zero value otherwise.
func (o *BTMicroversionIdAndConfiguration2338) GetMicroversion() BTMicroversionId366 {
	if o == nil || o.Microversion == nil {
		var ret BTMicroversionId366
		return ret
	}
	return *o.Microversion
}

// GetMicroversionOk returns a tuple with the Microversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMicroversionIdAndConfiguration2338) GetMicroversionOk() (*BTMicroversionId366, bool) {
	if o == nil || o.Microversion == nil {
		return nil, false
	}
	return o.Microversion, true
}

// HasMicroversion returns a boolean if a field has been set.
func (o *BTMicroversionIdAndConfiguration2338) HasMicroversion() bool {
	if o != nil && o.Microversion != nil {
		return true
	}

	return false
}

// SetMicroversion gets a reference to the given BTMicroversionId366 and assigns it to the Microversion field.
func (o *BTMicroversionIdAndConfiguration2338) SetMicroversion(v BTMicroversionId366) {
	o.Microversion = &v
}

func (o BTMicroversionIdAndConfiguration2338) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.CacheKey != nil {
		toSerialize["cacheKey"] = o.CacheKey
	}
	if o.ConfigurationParameterIdToValue != nil {
		toSerialize["configurationParameterIdToValue"] = o.ConfigurationParameterIdToValue
	}
	if o.Deleted != nil {
		toSerialize["deleted"] = o.Deleted
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Microversion != nil {
		toSerialize["microversion"] = o.Microversion
	}
	return json.Marshal(toSerialize)
}

type NullableBTMicroversionIdAndConfiguration2338 struct {
	value *BTMicroversionIdAndConfiguration2338
	isSet bool
}

func (v NullableBTMicroversionIdAndConfiguration2338) Get() *BTMicroversionIdAndConfiguration2338 {
	return v.value
}

func (v *NullableBTMicroversionIdAndConfiguration2338) Set(val *BTMicroversionIdAndConfiguration2338) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMicroversionIdAndConfiguration2338) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMicroversionIdAndConfiguration2338) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMicroversionIdAndConfiguration2338(val *BTMicroversionIdAndConfiguration2338) *NullableBTMicroversionIdAndConfiguration2338 {
	return &NullableBTMicroversionIdAndConfiguration2338{value: val, isSet: true}
}

func (v NullableBTMicroversionIdAndConfiguration2338) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMicroversionIdAndConfiguration2338) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
