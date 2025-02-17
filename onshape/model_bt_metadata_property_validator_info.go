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

// BTMetadataPropertyValidatorInfo struct for BTMetadataPropertyValidatorInfo
type BTMetadataPropertyValidatorInfo struct {
	Max          *float32  `json:"max,omitempty"`
	MaxCount     *int32    `json:"maxCount,omitempty"`
	MaxDate      *JSONTime `json:"maxDate,omitempty"`
	MaxLength    *int32    `json:"maxLength,omitempty"`
	Min          *float32  `json:"min,omitempty"`
	MinCount     *int32    `json:"minCount,omitempty"`
	MinDate      *JSONTime `json:"minDate,omitempty"`
	MinLength    *int32    `json:"minLength,omitempty"`
	Pattern      *string   `json:"pattern,omitempty"`
	QuantityType *int32    `json:"quantityType,omitempty"`
}

// NewBTMetadataPropertyValidatorInfo instantiates a new BTMetadataPropertyValidatorInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMetadataPropertyValidatorInfo() *BTMetadataPropertyValidatorInfo {
	this := BTMetadataPropertyValidatorInfo{}
	return &this
}

// NewBTMetadataPropertyValidatorInfoWithDefaults instantiates a new BTMetadataPropertyValidatorInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMetadataPropertyValidatorInfoWithDefaults() *BTMetadataPropertyValidatorInfo {
	this := BTMetadataPropertyValidatorInfo{}
	return &this
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *BTMetadataPropertyValidatorInfo) GetMax() float32 {
	if o == nil || o.Max == nil {
		var ret float32
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataPropertyValidatorInfo) GetMaxOk() (*float32, bool) {
	if o == nil || o.Max == nil {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *BTMetadataPropertyValidatorInfo) HasMax() bool {
	if o != nil && o.Max != nil {
		return true
	}

	return false
}

// SetMax gets a reference to the given float32 and assigns it to the Max field.
func (o *BTMetadataPropertyValidatorInfo) SetMax(v float32) {
	o.Max = &v
}

// GetMaxCount returns the MaxCount field value if set, zero value otherwise.
func (o *BTMetadataPropertyValidatorInfo) GetMaxCount() int32 {
	if o == nil || o.MaxCount == nil {
		var ret int32
		return ret
	}
	return *o.MaxCount
}

// GetMaxCountOk returns a tuple with the MaxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataPropertyValidatorInfo) GetMaxCountOk() (*int32, bool) {
	if o == nil || o.MaxCount == nil {
		return nil, false
	}
	return o.MaxCount, true
}

// HasMaxCount returns a boolean if a field has been set.
func (o *BTMetadataPropertyValidatorInfo) HasMaxCount() bool {
	if o != nil && o.MaxCount != nil {
		return true
	}

	return false
}

// SetMaxCount gets a reference to the given int32 and assigns it to the MaxCount field.
func (o *BTMetadataPropertyValidatorInfo) SetMaxCount(v int32) {
	o.MaxCount = &v
}

// GetMaxDate returns the MaxDate field value if set, zero value otherwise.
func (o *BTMetadataPropertyValidatorInfo) GetMaxDate() JSONTime {
	if o == nil || o.MaxDate == nil {
		var ret JSONTime
		return ret
	}
	return *o.MaxDate
}

// GetMaxDateOk returns a tuple with the MaxDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataPropertyValidatorInfo) GetMaxDateOk() (*JSONTime, bool) {
	if o == nil || o.MaxDate == nil {
		return nil, false
	}
	return o.MaxDate, true
}

// HasMaxDate returns a boolean if a field has been set.
func (o *BTMetadataPropertyValidatorInfo) HasMaxDate() bool {
	if o != nil && o.MaxDate != nil {
		return true
	}

	return false
}

// SetMaxDate gets a reference to the given JSONTime and assigns it to the MaxDate field.
func (o *BTMetadataPropertyValidatorInfo) SetMaxDate(v JSONTime) {
	o.MaxDate = &v
}

// GetMaxLength returns the MaxLength field value if set, zero value otherwise.
func (o *BTMetadataPropertyValidatorInfo) GetMaxLength() int32 {
	if o == nil || o.MaxLength == nil {
		var ret int32
		return ret
	}
	return *o.MaxLength
}

// GetMaxLengthOk returns a tuple with the MaxLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataPropertyValidatorInfo) GetMaxLengthOk() (*int32, bool) {
	if o == nil || o.MaxLength == nil {
		return nil, false
	}
	return o.MaxLength, true
}

// HasMaxLength returns a boolean if a field has been set.
func (o *BTMetadataPropertyValidatorInfo) HasMaxLength() bool {
	if o != nil && o.MaxLength != nil {
		return true
	}

	return false
}

// SetMaxLength gets a reference to the given int32 and assigns it to the MaxLength field.
func (o *BTMetadataPropertyValidatorInfo) SetMaxLength(v int32) {
	o.MaxLength = &v
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *BTMetadataPropertyValidatorInfo) GetMin() float32 {
	if o == nil || o.Min == nil {
		var ret float32
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataPropertyValidatorInfo) GetMinOk() (*float32, bool) {
	if o == nil || o.Min == nil {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *BTMetadataPropertyValidatorInfo) HasMin() bool {
	if o != nil && o.Min != nil {
		return true
	}

	return false
}

// SetMin gets a reference to the given float32 and assigns it to the Min field.
func (o *BTMetadataPropertyValidatorInfo) SetMin(v float32) {
	o.Min = &v
}

// GetMinCount returns the MinCount field value if set, zero value otherwise.
func (o *BTMetadataPropertyValidatorInfo) GetMinCount() int32 {
	if o == nil || o.MinCount == nil {
		var ret int32
		return ret
	}
	return *o.MinCount
}

// GetMinCountOk returns a tuple with the MinCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataPropertyValidatorInfo) GetMinCountOk() (*int32, bool) {
	if o == nil || o.MinCount == nil {
		return nil, false
	}
	return o.MinCount, true
}

// HasMinCount returns a boolean if a field has been set.
func (o *BTMetadataPropertyValidatorInfo) HasMinCount() bool {
	if o != nil && o.MinCount != nil {
		return true
	}

	return false
}

// SetMinCount gets a reference to the given int32 and assigns it to the MinCount field.
func (o *BTMetadataPropertyValidatorInfo) SetMinCount(v int32) {
	o.MinCount = &v
}

// GetMinDate returns the MinDate field value if set, zero value otherwise.
func (o *BTMetadataPropertyValidatorInfo) GetMinDate() JSONTime {
	if o == nil || o.MinDate == nil {
		var ret JSONTime
		return ret
	}
	return *o.MinDate
}

// GetMinDateOk returns a tuple with the MinDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataPropertyValidatorInfo) GetMinDateOk() (*JSONTime, bool) {
	if o == nil || o.MinDate == nil {
		return nil, false
	}
	return o.MinDate, true
}

// HasMinDate returns a boolean if a field has been set.
func (o *BTMetadataPropertyValidatorInfo) HasMinDate() bool {
	if o != nil && o.MinDate != nil {
		return true
	}

	return false
}

// SetMinDate gets a reference to the given JSONTime and assigns it to the MinDate field.
func (o *BTMetadataPropertyValidatorInfo) SetMinDate(v JSONTime) {
	o.MinDate = &v
}

// GetMinLength returns the MinLength field value if set, zero value otherwise.
func (o *BTMetadataPropertyValidatorInfo) GetMinLength() int32 {
	if o == nil || o.MinLength == nil {
		var ret int32
		return ret
	}
	return *o.MinLength
}

// GetMinLengthOk returns a tuple with the MinLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataPropertyValidatorInfo) GetMinLengthOk() (*int32, bool) {
	if o == nil || o.MinLength == nil {
		return nil, false
	}
	return o.MinLength, true
}

// HasMinLength returns a boolean if a field has been set.
func (o *BTMetadataPropertyValidatorInfo) HasMinLength() bool {
	if o != nil && o.MinLength != nil {
		return true
	}

	return false
}

// SetMinLength gets a reference to the given int32 and assigns it to the MinLength field.
func (o *BTMetadataPropertyValidatorInfo) SetMinLength(v int32) {
	o.MinLength = &v
}

// GetPattern returns the Pattern field value if set, zero value otherwise.
func (o *BTMetadataPropertyValidatorInfo) GetPattern() string {
	if o == nil || o.Pattern == nil {
		var ret string
		return ret
	}
	return *o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataPropertyValidatorInfo) GetPatternOk() (*string, bool) {
	if o == nil || o.Pattern == nil {
		return nil, false
	}
	return o.Pattern, true
}

// HasPattern returns a boolean if a field has been set.
func (o *BTMetadataPropertyValidatorInfo) HasPattern() bool {
	if o != nil && o.Pattern != nil {
		return true
	}

	return false
}

// SetPattern gets a reference to the given string and assigns it to the Pattern field.
func (o *BTMetadataPropertyValidatorInfo) SetPattern(v string) {
	o.Pattern = &v
}

// GetQuantityType returns the QuantityType field value if set, zero value otherwise.
func (o *BTMetadataPropertyValidatorInfo) GetQuantityType() int32 {
	if o == nil || o.QuantityType == nil {
		var ret int32
		return ret
	}
	return *o.QuantityType
}

// GetQuantityTypeOk returns a tuple with the QuantityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataPropertyValidatorInfo) GetQuantityTypeOk() (*int32, bool) {
	if o == nil || o.QuantityType == nil {
		return nil, false
	}
	return o.QuantityType, true
}

// HasQuantityType returns a boolean if a field has been set.
func (o *BTMetadataPropertyValidatorInfo) HasQuantityType() bool {
	if o != nil && o.QuantityType != nil {
		return true
	}

	return false
}

// SetQuantityType gets a reference to the given int32 and assigns it to the QuantityType field.
func (o *BTMetadataPropertyValidatorInfo) SetQuantityType(v int32) {
	o.QuantityType = &v
}

func (o BTMetadataPropertyValidatorInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Max != nil {
		toSerialize["max"] = o.Max
	}
	if o.MaxCount != nil {
		toSerialize["maxCount"] = o.MaxCount
	}
	if o.MaxDate != nil {
		toSerialize["maxDate"] = o.MaxDate
	}
	if o.MaxLength != nil {
		toSerialize["maxLength"] = o.MaxLength
	}
	if o.Min != nil {
		toSerialize["min"] = o.Min
	}
	if o.MinCount != nil {
		toSerialize["minCount"] = o.MinCount
	}
	if o.MinDate != nil {
		toSerialize["minDate"] = o.MinDate
	}
	if o.MinLength != nil {
		toSerialize["minLength"] = o.MinLength
	}
	if o.Pattern != nil {
		toSerialize["pattern"] = o.Pattern
	}
	if o.QuantityType != nil {
		toSerialize["quantityType"] = o.QuantityType
	}
	return json.Marshal(toSerialize)
}

type NullableBTMetadataPropertyValidatorInfo struct {
	value *BTMetadataPropertyValidatorInfo
	isSet bool
}

func (v NullableBTMetadataPropertyValidatorInfo) Get() *BTMetadataPropertyValidatorInfo {
	return v.value
}

func (v *NullableBTMetadataPropertyValidatorInfo) Set(val *BTMetadataPropertyValidatorInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMetadataPropertyValidatorInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMetadataPropertyValidatorInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMetadataPropertyValidatorInfo(val *BTMetadataPropertyValidatorInfo) *NullableBTMetadataPropertyValidatorInfo {
	return &NullableBTMetadataPropertyValidatorInfo{value: val, isSet: true}
}

func (v NullableBTMetadataPropertyValidatorInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMetadataPropertyValidatorInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
