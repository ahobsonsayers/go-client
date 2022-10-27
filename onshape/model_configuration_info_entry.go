/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.155.7076-cd7f519b38e7
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// ConfigurationInfoEntry struct for ConfigurationInfoEntry
type ConfigurationInfoEntry struct {
	Explicit                         *bool   `json:"explicit,omitempty"`
	ParameterAbbreviatedDisplayValue *string `json:"parameterAbbreviatedDisplayValue,omitempty"`
	ParameterDisplayValue            *string `json:"parameterDisplayValue,omitempty"`
	ParameterId                      *string `json:"parameterId,omitempty"`
	ParameterName                    *string `json:"parameterName,omitempty"`
	ParameterType                    *int32  `json:"parameterType,omitempty"`
	ParameterValue                   *string `json:"parameterValue,omitempty"`
}

// NewConfigurationInfoEntry instantiates a new ConfigurationInfoEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigurationInfoEntry() *ConfigurationInfoEntry {
	this := ConfigurationInfoEntry{}
	return &this
}

// NewConfigurationInfoEntryWithDefaults instantiates a new ConfigurationInfoEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigurationInfoEntryWithDefaults() *ConfigurationInfoEntry {
	this := ConfigurationInfoEntry{}
	return &this
}

// GetExplicit returns the Explicit field value if set, zero value otherwise.
func (o *ConfigurationInfoEntry) GetExplicit() bool {
	if o == nil || o.Explicit == nil {
		var ret bool
		return ret
	}
	return *o.Explicit
}

// GetExplicitOk returns a tuple with the Explicit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationInfoEntry) GetExplicitOk() (*bool, bool) {
	if o == nil || o.Explicit == nil {
		return nil, false
	}
	return o.Explicit, true
}

// HasExplicit returns a boolean if a field has been set.
func (o *ConfigurationInfoEntry) HasExplicit() bool {
	if o != nil && o.Explicit != nil {
		return true
	}

	return false
}

// SetExplicit gets a reference to the given bool and assigns it to the Explicit field.
func (o *ConfigurationInfoEntry) SetExplicit(v bool) {
	o.Explicit = &v
}

// GetParameterAbbreviatedDisplayValue returns the ParameterAbbreviatedDisplayValue field value if set, zero value otherwise.
func (o *ConfigurationInfoEntry) GetParameterAbbreviatedDisplayValue() string {
	if o == nil || o.ParameterAbbreviatedDisplayValue == nil {
		var ret string
		return ret
	}
	return *o.ParameterAbbreviatedDisplayValue
}

// GetParameterAbbreviatedDisplayValueOk returns a tuple with the ParameterAbbreviatedDisplayValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationInfoEntry) GetParameterAbbreviatedDisplayValueOk() (*string, bool) {
	if o == nil || o.ParameterAbbreviatedDisplayValue == nil {
		return nil, false
	}
	return o.ParameterAbbreviatedDisplayValue, true
}

// HasParameterAbbreviatedDisplayValue returns a boolean if a field has been set.
func (o *ConfigurationInfoEntry) HasParameterAbbreviatedDisplayValue() bool {
	if o != nil && o.ParameterAbbreviatedDisplayValue != nil {
		return true
	}

	return false
}

// SetParameterAbbreviatedDisplayValue gets a reference to the given string and assigns it to the ParameterAbbreviatedDisplayValue field.
func (o *ConfigurationInfoEntry) SetParameterAbbreviatedDisplayValue(v string) {
	o.ParameterAbbreviatedDisplayValue = &v
}

// GetParameterDisplayValue returns the ParameterDisplayValue field value if set, zero value otherwise.
func (o *ConfigurationInfoEntry) GetParameterDisplayValue() string {
	if o == nil || o.ParameterDisplayValue == nil {
		var ret string
		return ret
	}
	return *o.ParameterDisplayValue
}

// GetParameterDisplayValueOk returns a tuple with the ParameterDisplayValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationInfoEntry) GetParameterDisplayValueOk() (*string, bool) {
	if o == nil || o.ParameterDisplayValue == nil {
		return nil, false
	}
	return o.ParameterDisplayValue, true
}

// HasParameterDisplayValue returns a boolean if a field has been set.
func (o *ConfigurationInfoEntry) HasParameterDisplayValue() bool {
	if o != nil && o.ParameterDisplayValue != nil {
		return true
	}

	return false
}

// SetParameterDisplayValue gets a reference to the given string and assigns it to the ParameterDisplayValue field.
func (o *ConfigurationInfoEntry) SetParameterDisplayValue(v string) {
	o.ParameterDisplayValue = &v
}

// GetParameterId returns the ParameterId field value if set, zero value otherwise.
func (o *ConfigurationInfoEntry) GetParameterId() string {
	if o == nil || o.ParameterId == nil {
		var ret string
		return ret
	}
	return *o.ParameterId
}

// GetParameterIdOk returns a tuple with the ParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationInfoEntry) GetParameterIdOk() (*string, bool) {
	if o == nil || o.ParameterId == nil {
		return nil, false
	}
	return o.ParameterId, true
}

// HasParameterId returns a boolean if a field has been set.
func (o *ConfigurationInfoEntry) HasParameterId() bool {
	if o != nil && o.ParameterId != nil {
		return true
	}

	return false
}

// SetParameterId gets a reference to the given string and assigns it to the ParameterId field.
func (o *ConfigurationInfoEntry) SetParameterId(v string) {
	o.ParameterId = &v
}

// GetParameterName returns the ParameterName field value if set, zero value otherwise.
func (o *ConfigurationInfoEntry) GetParameterName() string {
	if o == nil || o.ParameterName == nil {
		var ret string
		return ret
	}
	return *o.ParameterName
}

// GetParameterNameOk returns a tuple with the ParameterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationInfoEntry) GetParameterNameOk() (*string, bool) {
	if o == nil || o.ParameterName == nil {
		return nil, false
	}
	return o.ParameterName, true
}

// HasParameterName returns a boolean if a field has been set.
func (o *ConfigurationInfoEntry) HasParameterName() bool {
	if o != nil && o.ParameterName != nil {
		return true
	}

	return false
}

// SetParameterName gets a reference to the given string and assigns it to the ParameterName field.
func (o *ConfigurationInfoEntry) SetParameterName(v string) {
	o.ParameterName = &v
}

// GetParameterType returns the ParameterType field value if set, zero value otherwise.
func (o *ConfigurationInfoEntry) GetParameterType() int32 {
	if o == nil || o.ParameterType == nil {
		var ret int32
		return ret
	}
	return *o.ParameterType
}

// GetParameterTypeOk returns a tuple with the ParameterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationInfoEntry) GetParameterTypeOk() (*int32, bool) {
	if o == nil || o.ParameterType == nil {
		return nil, false
	}
	return o.ParameterType, true
}

// HasParameterType returns a boolean if a field has been set.
func (o *ConfigurationInfoEntry) HasParameterType() bool {
	if o != nil && o.ParameterType != nil {
		return true
	}

	return false
}

// SetParameterType gets a reference to the given int32 and assigns it to the ParameterType field.
func (o *ConfigurationInfoEntry) SetParameterType(v int32) {
	o.ParameterType = &v
}

// GetParameterValue returns the ParameterValue field value if set, zero value otherwise.
func (o *ConfigurationInfoEntry) GetParameterValue() string {
	if o == nil || o.ParameterValue == nil {
		var ret string
		return ret
	}
	return *o.ParameterValue
}

// GetParameterValueOk returns a tuple with the ParameterValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationInfoEntry) GetParameterValueOk() (*string, bool) {
	if o == nil || o.ParameterValue == nil {
		return nil, false
	}
	return o.ParameterValue, true
}

// HasParameterValue returns a boolean if a field has been set.
func (o *ConfigurationInfoEntry) HasParameterValue() bool {
	if o != nil && o.ParameterValue != nil {
		return true
	}

	return false
}

// SetParameterValue gets a reference to the given string and assigns it to the ParameterValue field.
func (o *ConfigurationInfoEntry) SetParameterValue(v string) {
	o.ParameterValue = &v
}

func (o ConfigurationInfoEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Explicit != nil {
		toSerialize["explicit"] = o.Explicit
	}
	if o.ParameterAbbreviatedDisplayValue != nil {
		toSerialize["parameterAbbreviatedDisplayValue"] = o.ParameterAbbreviatedDisplayValue
	}
	if o.ParameterDisplayValue != nil {
		toSerialize["parameterDisplayValue"] = o.ParameterDisplayValue
	}
	if o.ParameterId != nil {
		toSerialize["parameterId"] = o.ParameterId
	}
	if o.ParameterName != nil {
		toSerialize["parameterName"] = o.ParameterName
	}
	if o.ParameterType != nil {
		toSerialize["parameterType"] = o.ParameterType
	}
	if o.ParameterValue != nil {
		toSerialize["parameterValue"] = o.ParameterValue
	}
	return json.Marshal(toSerialize)
}

type NullableConfigurationInfoEntry struct {
	value *ConfigurationInfoEntry
	isSet bool
}

func (v NullableConfigurationInfoEntry) Get() *ConfigurationInfoEntry {
	return v.value
}

func (v *NullableConfigurationInfoEntry) Set(val *ConfigurationInfoEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigurationInfoEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigurationInfoEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigurationInfoEntry(val *ConfigurationInfoEntry) *NullableConfigurationInfoEntry {
	return &NullableConfigurationInfoEntry{value: val, isSet: true}
}

func (v NullableConfigurationInfoEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigurationInfoEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
