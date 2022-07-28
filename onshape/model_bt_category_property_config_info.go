/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.151.5750-4f2542599dd4
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTCategoryPropertyConfigInfo struct for BTCategoryPropertyConfigInfo
type BTCategoryPropertyConfigInfo struct {
	ComputedPartPropertyConfig        *BTComputedPartPropertyConfig `json:"computedPartPropertyConfig,omitempty"`
	ComputedPropertyFunctionName      *string                       `json:"computedPropertyFunctionName,omitempty"`
	ComputedPropertyFunctionNamespace *string                       `json:"computedPropertyFunctionNamespace,omitempty"`
	ComputedPropertyFunctionURL       *string                       `json:"computedPropertyFunctionURL,omitempty"`
	DefaultValue                      *string                       `json:"defaultValue,omitempty"`
	DisplayName                       *string                       `json:"displayName,omitempty"`
	EnumValues                        []BTMetadataEnumValue         `json:"enumValues,omitempty"`
	MaxCount                          *int32                        `json:"maxCount,omitempty"`
	MaxDate                           *JSONTime                     `json:"maxDate,omitempty"`
	MaxLength                         *int32                        `json:"maxLength,omitempty"`
	MaxValue                          *float64                      `json:"maxValue,omitempty"`
	MinCount                          *int32                        `json:"minCount,omitempty"`
	MinDate                           *JSONTime                     `json:"minDate,omitempty"`
	MinLength                         *int32                        `json:"minLength,omitempty"`
	MinValue                          *float64                      `json:"minValue,omitempty"`
	Multiline                         *bool                         `json:"multiline,omitempty"`
	Multivalued                       *bool                         `json:"multivalued,omitempty"`
	Pattern                           *string                       `json:"pattern,omitempty"`
	PublishState                      *int32                        `json:"publishState,omitempty"`
	QuantityType                      *int32                        `json:"quantityType,omitempty"`
	Required                          *bool                         `json:"required,omitempty"`
}

// NewBTCategoryPropertyConfigInfo instantiates a new BTCategoryPropertyConfigInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTCategoryPropertyConfigInfo() *BTCategoryPropertyConfigInfo {
	this := BTCategoryPropertyConfigInfo{}
	return &this
}

// NewBTCategoryPropertyConfigInfoWithDefaults instantiates a new BTCategoryPropertyConfigInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTCategoryPropertyConfigInfoWithDefaults() *BTCategoryPropertyConfigInfo {
	this := BTCategoryPropertyConfigInfo{}
	return &this
}

// GetComputedPartPropertyConfig returns the ComputedPartPropertyConfig field value if set, zero value otherwise.
func (o *BTCategoryPropertyConfigInfo) GetComputedPartPropertyConfig() BTComputedPartPropertyConfig {
	if o == nil || o.ComputedPartPropertyConfig == nil {
		var ret BTComputedPartPropertyConfig
		return ret
	}
	return *o.ComputedPartPropertyConfig
}

// GetComputedPartPropertyConfigOk returns a tuple with the ComputedPartPropertyConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCategoryPropertyConfigInfo) GetComputedPartPropertyConfigOk() (*BTComputedPartPropertyConfig, bool) {
	if o == nil || o.ComputedPartPropertyConfig == nil {
		return nil, false
	}
	return o.ComputedPartPropertyConfig, true
}

// HasComputedPartPropertyConfig returns a boolean if a field has been set.
func (o *BTCategoryPropertyConfigInfo) HasComputedPartPropertyConfig() bool {
	if o != nil && o.ComputedPartPropertyConfig != nil {
		return true
	}

	return false
}

// SetComputedPartPropertyConfig gets a reference to the given BTComputedPartPropertyConfig and assigns it to the ComputedPartPropertyConfig field.
func (o *BTCategoryPropertyConfigInfo) SetComputedPartPropertyConfig(v BTComputedPartPropertyConfig) {
	o.ComputedPartPropertyConfig = &v
}

// GetComputedPropertyFunctionName returns the ComputedPropertyFunctionName field value if set, zero value otherwise.
func (o *BTCategoryPropertyConfigInfo) GetComputedPropertyFunctionName() string {
	if o == nil || o.ComputedPropertyFunctionName == nil {
		var ret string
		return ret
	}
	return *o.ComputedPropertyFunctionName
}

// GetComputedPropertyFunctionNameOk returns a tuple with the ComputedPropertyFunctionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCategoryPropertyConfigInfo) GetComputedPropertyFunctionNameOk() (*string, bool) {
	if o == nil || o.ComputedPropertyFunctionName == nil {
		return nil, false
	}
	return o.ComputedPropertyFunctionName, true
}

// HasComputedPropertyFunctionName returns a boolean if a field has been set.
func (o *BTCategoryPropertyConfigInfo) HasComputedPropertyFunctionName() bool {
	if o != nil && o.ComputedPropertyFunctionName != nil {
		return true
	}

	return false
}

// SetComputedPropertyFunctionName gets a reference to the given string and assigns it to the ComputedPropertyFunctionName field.
func (o *BTCategoryPropertyConfigInfo) SetComputedPropertyFunctionName(v string) {
	o.ComputedPropertyFunctionName = &v
}

// GetComputedPropertyFunctionNamespace returns the ComputedPropertyFunctionNamespace field value if set, zero value otherwise.
func (o *BTCategoryPropertyConfigInfo) GetComputedPropertyFunctionNamespace() string {
	if o == nil || o.ComputedPropertyFunctionNamespace == nil {
		var ret string
		return ret
	}
	return *o.ComputedPropertyFunctionNamespace
}

// GetComputedPropertyFunctionNamespaceOk returns a tuple with the ComputedPropertyFunctionNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCategoryPropertyConfigInfo) GetComputedPropertyFunctionNamespaceOk() (*string, bool) {
	if o == nil || o.ComputedPropertyFunctionNamespace == nil {
		return nil, false
	}
	return o.ComputedPropertyFunctionNamespace, true
}

// HasComputedPropertyFunctionNamespace returns a boolean if a field has been set.
func (o *BTCategoryPropertyConfigInfo) HasComputedPropertyFunctionNamespace() bool {
	if o != nil && o.ComputedPropertyFunctionNamespace != nil {
		return true
	}

	return false
}

// SetComputedPropertyFunctionNamespace gets a reference to the given string and assigns it to the ComputedPropertyFunctionNamespace field.
func (o *BTCategoryPropertyConfigInfo) SetComputedPropertyFunctionNamespace(v string) {
	o.ComputedPropertyFunctionNamespace = &v
}

// GetComputedPropertyFunctionURL returns the ComputedPropertyFunctionURL field value if set, zero value otherwise.
func (o *BTCategoryPropertyConfigInfo) GetComputedPropertyFunctionURL() string {
	if o == nil || o.ComputedPropertyFunctionURL == nil {
		var ret string
		return ret
	}
	return *o.ComputedPropertyFunctionURL
}

// GetComputedPropertyFunctionURLOk returns a tuple with the ComputedPropertyFunctionURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCategoryPropertyConfigInfo) GetComputedPropertyFunctionURLOk() (*string, bool) {
	if o == nil || o.ComputedPropertyFunctionURL == nil {
		return nil, false
	}
	return o.ComputedPropertyFunctionURL, true
}

// HasComputedPropertyFunctionURL returns a boolean if a field has been set.
func (o *BTCategoryPropertyConfigInfo) HasComputedPropertyFunctionURL() bool {
	if o != nil && o.ComputedPropertyFunctionURL != nil {
		return true
	}

	return false
}

// SetComputedPropertyFunctionURL gets a reference to the given string and assigns it to the ComputedPropertyFunctionURL field.
func (o *BTCategoryPropertyConfigInfo) SetComputedPropertyFunctionURL(v string) {
	o.ComputedPropertyFunctionURL = &v
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise.
func (o *BTCategoryPropertyConfigInfo) GetDefaultValue() string {
	if o == nil || o.DefaultValue == nil {
		var ret string
		return ret
	}
	return *o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCategoryPropertyConfigInfo) GetDefaultValueOk() (*string, bool) {
	if o == nil || o.DefaultValue == nil {
		return nil, false
	}
	return o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *BTCategoryPropertyConfigInfo) HasDefaultValue() bool {
	if o != nil && o.DefaultValue != nil {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given string and assigns it to the DefaultValue field.
func (o *BTCategoryPropertyConfigInfo) SetDefaultValue(v string) {
	o.DefaultValue = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *BTCategoryPropertyConfigInfo) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCategoryPropertyConfigInfo) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *BTCategoryPropertyConfigInfo) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *BTCategoryPropertyConfigInfo) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEnumValues returns the EnumValues field value if set, zero value otherwise.
func (o *BTCategoryPropertyConfigInfo) GetEnumValues() []BTMetadataEnumValue {
	if o == nil || o.EnumValues == nil {
		var ret []BTMetadataEnumValue
		return ret
	}
	return o.EnumValues
}

// GetEnumValuesOk returns a tuple with the EnumValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCategoryPropertyConfigInfo) GetEnumValuesOk() ([]BTMetadataEnumValue, bool) {
	if o == nil || o.EnumValues == nil {
		return nil, false
	}
	return o.EnumValues, true
}

// HasEnumValues returns a boolean if a field has been set.
func (o *BTCategoryPropertyConfigInfo) HasEnumValues() bool {
	if o != nil && o.EnumValues != nil {
		return true
	}

	return false
}

// SetEnumValues gets a reference to the given []BTMetadataEnumValue and assigns it to the EnumValues field.
func (o *BTCategoryPropertyConfigInfo) SetEnumValues(v []BTMetadataEnumValue) {
	o.EnumValues = v
}

// GetMaxCount returns the MaxCount field value if set, zero value otherwise.
func (o *BTCategoryPropertyConfigInfo) GetMaxCount() int32 {
	if o == nil || o.MaxCount == nil {
		var ret int32
		return ret
	}
	return *o.MaxCount
}

// GetMaxCountOk returns a tuple with the MaxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCategoryPropertyConfigInfo) GetMaxCountOk() (*int32, bool) {
	if o == nil || o.MaxCount == nil {
		return nil, false
	}
	return o.MaxCount, true
}

// HasMaxCount returns a boolean if a field has been set.
func (o *BTCategoryPropertyConfigInfo) HasMaxCount() bool {
	if o != nil && o.MaxCount != nil {
		return true
	}

	return false
}

// SetMaxCount gets a reference to the given int32 and assigns it to the MaxCount field.
func (o *BTCategoryPropertyConfigInfo) SetMaxCount(v int32) {
	o.MaxCount = &v
}

// GetMaxDate returns the MaxDate field value if set, zero value otherwise.
func (o *BTCategoryPropertyConfigInfo) GetMaxDate() JSONTime {
	if o == nil || o.MaxDate == nil {
		var ret JSONTime
		return ret
	}
	return *o.MaxDate
}

// GetMaxDateOk returns a tuple with the MaxDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCategoryPropertyConfigInfo) GetMaxDateOk() (*JSONTime, bool) {
	if o == nil || o.MaxDate == nil {
		return nil, false
	}
	return o.MaxDate, true
}

// HasMaxDate returns a boolean if a field has been set.
func (o *BTCategoryPropertyConfigInfo) HasMaxDate() bool {
	if o != nil && o.MaxDate != nil {
		return true
	}

	return false
}

// SetMaxDate gets a reference to the given JSONTime and assigns it to the MaxDate field.
func (o *BTCategoryPropertyConfigInfo) SetMaxDate(v JSONTime) {
	o.MaxDate = &v
}

// GetMaxLength returns the MaxLength field value if set, zero value otherwise.
func (o *BTCategoryPropertyConfigInfo) GetMaxLength() int32 {
	if o == nil || o.MaxLength == nil {
		var ret int32
		return ret
	}
	return *o.MaxLength
}

// GetMaxLengthOk returns a tuple with the MaxLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCategoryPropertyConfigInfo) GetMaxLengthOk() (*int32, bool) {
	if o == nil || o.MaxLength == nil {
		return nil, false
	}
	return o.MaxLength, true
}

// HasMaxLength returns a boolean if a field has been set.
func (o *BTCategoryPropertyConfigInfo) HasMaxLength() bool {
	if o != nil && o.MaxLength != nil {
		return true
	}

	return false
}

// SetMaxLength gets a reference to the given int32 and assigns it to the MaxLength field.
func (o *BTCategoryPropertyConfigInfo) SetMaxLength(v int32) {
	o.MaxLength = &v
}

// GetMaxValue returns the MaxValue field value if set, zero value otherwise.
func (o *BTCategoryPropertyConfigInfo) GetMaxValue() float64 {
	if o == nil || o.MaxValue == nil {
		var ret float64
		return ret
	}
	return *o.MaxValue
}

// GetMaxValueOk returns a tuple with the MaxValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCategoryPropertyConfigInfo) GetMaxValueOk() (*float64, bool) {
	if o == nil || o.MaxValue == nil {
		return nil, false
	}
	return o.MaxValue, true
}

// HasMaxValue returns a boolean if a field has been set.
func (o *BTCategoryPropertyConfigInfo) HasMaxValue() bool {
	if o != nil && o.MaxValue != nil {
		return true
	}

	return false
}

// SetMaxValue gets a reference to the given float64 and assigns it to the MaxValue field.
func (o *BTCategoryPropertyConfigInfo) SetMaxValue(v float64) {
	o.MaxValue = &v
}

// GetMinCount returns the MinCount field value if set, zero value otherwise.
func (o *BTCategoryPropertyConfigInfo) GetMinCount() int32 {
	if o == nil || o.MinCount == nil {
		var ret int32
		return ret
	}
	return *o.MinCount
}

// GetMinCountOk returns a tuple with the MinCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCategoryPropertyConfigInfo) GetMinCountOk() (*int32, bool) {
	if o == nil || o.MinCount == nil {
		return nil, false
	}
	return o.MinCount, true
}

// HasMinCount returns a boolean if a field has been set.
func (o *BTCategoryPropertyConfigInfo) HasMinCount() bool {
	if o != nil && o.MinCount != nil {
		return true
	}

	return false
}

// SetMinCount gets a reference to the given int32 and assigns it to the MinCount field.
func (o *BTCategoryPropertyConfigInfo) SetMinCount(v int32) {
	o.MinCount = &v
}

// GetMinDate returns the MinDate field value if set, zero value otherwise.
func (o *BTCategoryPropertyConfigInfo) GetMinDate() JSONTime {
	if o == nil || o.MinDate == nil {
		var ret JSONTime
		return ret
	}
	return *o.MinDate
}

// GetMinDateOk returns a tuple with the MinDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCategoryPropertyConfigInfo) GetMinDateOk() (*JSONTime, bool) {
	if o == nil || o.MinDate == nil {
		return nil, false
	}
	return o.MinDate, true
}

// HasMinDate returns a boolean if a field has been set.
func (o *BTCategoryPropertyConfigInfo) HasMinDate() bool {
	if o != nil && o.MinDate != nil {
		return true
	}

	return false
}

// SetMinDate gets a reference to the given JSONTime and assigns it to the MinDate field.
func (o *BTCategoryPropertyConfigInfo) SetMinDate(v JSONTime) {
	o.MinDate = &v
}

// GetMinLength returns the MinLength field value if set, zero value otherwise.
func (o *BTCategoryPropertyConfigInfo) GetMinLength() int32 {
	if o == nil || o.MinLength == nil {
		var ret int32
		return ret
	}
	return *o.MinLength
}

// GetMinLengthOk returns a tuple with the MinLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCategoryPropertyConfigInfo) GetMinLengthOk() (*int32, bool) {
	if o == nil || o.MinLength == nil {
		return nil, false
	}
	return o.MinLength, true
}

// HasMinLength returns a boolean if a field has been set.
func (o *BTCategoryPropertyConfigInfo) HasMinLength() bool {
	if o != nil && o.MinLength != nil {
		return true
	}

	return false
}

// SetMinLength gets a reference to the given int32 and assigns it to the MinLength field.
func (o *BTCategoryPropertyConfigInfo) SetMinLength(v int32) {
	o.MinLength = &v
}

// GetMinValue returns the MinValue field value if set, zero value otherwise.
func (o *BTCategoryPropertyConfigInfo) GetMinValue() float64 {
	if o == nil || o.MinValue == nil {
		var ret float64
		return ret
	}
	return *o.MinValue
}

// GetMinValueOk returns a tuple with the MinValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCategoryPropertyConfigInfo) GetMinValueOk() (*float64, bool) {
	if o == nil || o.MinValue == nil {
		return nil, false
	}
	return o.MinValue, true
}

// HasMinValue returns a boolean if a field has been set.
func (o *BTCategoryPropertyConfigInfo) HasMinValue() bool {
	if o != nil && o.MinValue != nil {
		return true
	}

	return false
}

// SetMinValue gets a reference to the given float64 and assigns it to the MinValue field.
func (o *BTCategoryPropertyConfigInfo) SetMinValue(v float64) {
	o.MinValue = &v
}

// GetMultiline returns the Multiline field value if set, zero value otherwise.
func (o *BTCategoryPropertyConfigInfo) GetMultiline() bool {
	if o == nil || o.Multiline == nil {
		var ret bool
		return ret
	}
	return *o.Multiline
}

// GetMultilineOk returns a tuple with the Multiline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCategoryPropertyConfigInfo) GetMultilineOk() (*bool, bool) {
	if o == nil || o.Multiline == nil {
		return nil, false
	}
	return o.Multiline, true
}

// HasMultiline returns a boolean if a field has been set.
func (o *BTCategoryPropertyConfigInfo) HasMultiline() bool {
	if o != nil && o.Multiline != nil {
		return true
	}

	return false
}

// SetMultiline gets a reference to the given bool and assigns it to the Multiline field.
func (o *BTCategoryPropertyConfigInfo) SetMultiline(v bool) {
	o.Multiline = &v
}

// GetMultivalued returns the Multivalued field value if set, zero value otherwise.
func (o *BTCategoryPropertyConfigInfo) GetMultivalued() bool {
	if o == nil || o.Multivalued == nil {
		var ret bool
		return ret
	}
	return *o.Multivalued
}

// GetMultivaluedOk returns a tuple with the Multivalued field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCategoryPropertyConfigInfo) GetMultivaluedOk() (*bool, bool) {
	if o == nil || o.Multivalued == nil {
		return nil, false
	}
	return o.Multivalued, true
}

// HasMultivalued returns a boolean if a field has been set.
func (o *BTCategoryPropertyConfigInfo) HasMultivalued() bool {
	if o != nil && o.Multivalued != nil {
		return true
	}

	return false
}

// SetMultivalued gets a reference to the given bool and assigns it to the Multivalued field.
func (o *BTCategoryPropertyConfigInfo) SetMultivalued(v bool) {
	o.Multivalued = &v
}

// GetPattern returns the Pattern field value if set, zero value otherwise.
func (o *BTCategoryPropertyConfigInfo) GetPattern() string {
	if o == nil || o.Pattern == nil {
		var ret string
		return ret
	}
	return *o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCategoryPropertyConfigInfo) GetPatternOk() (*string, bool) {
	if o == nil || o.Pattern == nil {
		return nil, false
	}
	return o.Pattern, true
}

// HasPattern returns a boolean if a field has been set.
func (o *BTCategoryPropertyConfigInfo) HasPattern() bool {
	if o != nil && o.Pattern != nil {
		return true
	}

	return false
}

// SetPattern gets a reference to the given string and assigns it to the Pattern field.
func (o *BTCategoryPropertyConfigInfo) SetPattern(v string) {
	o.Pattern = &v
}

// GetPublishState returns the PublishState field value if set, zero value otherwise.
func (o *BTCategoryPropertyConfigInfo) GetPublishState() int32 {
	if o == nil || o.PublishState == nil {
		var ret int32
		return ret
	}
	return *o.PublishState
}

// GetPublishStateOk returns a tuple with the PublishState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCategoryPropertyConfigInfo) GetPublishStateOk() (*int32, bool) {
	if o == nil || o.PublishState == nil {
		return nil, false
	}
	return o.PublishState, true
}

// HasPublishState returns a boolean if a field has been set.
func (o *BTCategoryPropertyConfigInfo) HasPublishState() bool {
	if o != nil && o.PublishState != nil {
		return true
	}

	return false
}

// SetPublishState gets a reference to the given int32 and assigns it to the PublishState field.
func (o *BTCategoryPropertyConfigInfo) SetPublishState(v int32) {
	o.PublishState = &v
}

// GetQuantityType returns the QuantityType field value if set, zero value otherwise.
func (o *BTCategoryPropertyConfigInfo) GetQuantityType() int32 {
	if o == nil || o.QuantityType == nil {
		var ret int32
		return ret
	}
	return *o.QuantityType
}

// GetQuantityTypeOk returns a tuple with the QuantityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCategoryPropertyConfigInfo) GetQuantityTypeOk() (*int32, bool) {
	if o == nil || o.QuantityType == nil {
		return nil, false
	}
	return o.QuantityType, true
}

// HasQuantityType returns a boolean if a field has been set.
func (o *BTCategoryPropertyConfigInfo) HasQuantityType() bool {
	if o != nil && o.QuantityType != nil {
		return true
	}

	return false
}

// SetQuantityType gets a reference to the given int32 and assigns it to the QuantityType field.
func (o *BTCategoryPropertyConfigInfo) SetQuantityType(v int32) {
	o.QuantityType = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *BTCategoryPropertyConfigInfo) GetRequired() bool {
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCategoryPropertyConfigInfo) GetRequiredOk() (*bool, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *BTCategoryPropertyConfigInfo) HasRequired() bool {
	if o != nil && o.Required != nil {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *BTCategoryPropertyConfigInfo) SetRequired(v bool) {
	o.Required = &v
}

func (o BTCategoryPropertyConfigInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ComputedPartPropertyConfig != nil {
		toSerialize["computedPartPropertyConfig"] = o.ComputedPartPropertyConfig
	}
	if o.ComputedPropertyFunctionName != nil {
		toSerialize["computedPropertyFunctionName"] = o.ComputedPropertyFunctionName
	}
	if o.ComputedPropertyFunctionNamespace != nil {
		toSerialize["computedPropertyFunctionNamespace"] = o.ComputedPropertyFunctionNamespace
	}
	if o.ComputedPropertyFunctionURL != nil {
		toSerialize["computedPropertyFunctionURL"] = o.ComputedPropertyFunctionURL
	}
	if o.DefaultValue != nil {
		toSerialize["defaultValue"] = o.DefaultValue
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.EnumValues != nil {
		toSerialize["enumValues"] = o.EnumValues
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
	if o.MaxValue != nil {
		toSerialize["maxValue"] = o.MaxValue
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
	if o.MinValue != nil {
		toSerialize["minValue"] = o.MinValue
	}
	if o.Multiline != nil {
		toSerialize["multiline"] = o.Multiline
	}
	if o.Multivalued != nil {
		toSerialize["multivalued"] = o.Multivalued
	}
	if o.Pattern != nil {
		toSerialize["pattern"] = o.Pattern
	}
	if o.PublishState != nil {
		toSerialize["publishState"] = o.PublishState
	}
	if o.QuantityType != nil {
		toSerialize["quantityType"] = o.QuantityType
	}
	if o.Required != nil {
		toSerialize["required"] = o.Required
	}
	return json.Marshal(toSerialize)
}

type NullableBTCategoryPropertyConfigInfo struct {
	value *BTCategoryPropertyConfigInfo
	isSet bool
}

func (v NullableBTCategoryPropertyConfigInfo) Get() *BTCategoryPropertyConfigInfo {
	return v.value
}

func (v *NullableBTCategoryPropertyConfigInfo) Set(val *BTCategoryPropertyConfigInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTCategoryPropertyConfigInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTCategoryPropertyConfigInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTCategoryPropertyConfigInfo(val *BTCategoryPropertyConfigInfo) *NullableBTCategoryPropertyConfigInfo {
	return &NullableBTCategoryPropertyConfigInfo{value: val, isSet: true}
}

func (v NullableBTCategoryPropertyConfigInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTCategoryPropertyConfigInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
