/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.161.13200-ff216a970a02
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTConfigurationUpdateCall2933 struct for BTConfigurationUpdateCall2933
type BTConfigurationUpdateCall2933 struct {
	ConfigurationParameters []BTMConfigurationParameter819 `json:"configurationParameters,omitempty"`
	CurrentConfiguration    []BTMParameter1                `json:"currentConfiguration,omitempty"`
	LibraryVersion          *int32                         `json:"libraryVersion,omitempty"`
	MicroversionSkew        *bool                          `json:"microversionSkew,omitempty"`
	RejectMicroversionSkew  *bool                          `json:"rejectMicroversionSkew,omitempty"`
	SerializationVersion    *string                        `json:"serializationVersion,omitempty"`
	SourceMicroversion      *string                        `json:"sourceMicroversion,omitempty"`
}

// NewBTConfigurationUpdateCall2933 instantiates a new BTConfigurationUpdateCall2933 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTConfigurationUpdateCall2933() *BTConfigurationUpdateCall2933 {
	this := BTConfigurationUpdateCall2933{}
	return &this
}

// NewBTConfigurationUpdateCall2933WithDefaults instantiates a new BTConfigurationUpdateCall2933 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTConfigurationUpdateCall2933WithDefaults() *BTConfigurationUpdateCall2933 {
	this := BTConfigurationUpdateCall2933{}
	return &this
}

// GetConfigurationParameters returns the ConfigurationParameters field value if set, zero value otherwise.
func (o *BTConfigurationUpdateCall2933) GetConfigurationParameters() []BTMConfigurationParameter819 {
	if o == nil || o.ConfigurationParameters == nil {
		var ret []BTMConfigurationParameter819
		return ret
	}
	return o.ConfigurationParameters
}

// GetConfigurationParametersOk returns a tuple with the ConfigurationParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfigurationUpdateCall2933) GetConfigurationParametersOk() ([]BTMConfigurationParameter819, bool) {
	if o == nil || o.ConfigurationParameters == nil {
		return nil, false
	}
	return o.ConfigurationParameters, true
}

// HasConfigurationParameters returns a boolean if a field has been set.
func (o *BTConfigurationUpdateCall2933) HasConfigurationParameters() bool {
	if o != nil && o.ConfigurationParameters != nil {
		return true
	}

	return false
}

// SetConfigurationParameters gets a reference to the given []BTMConfigurationParameter819 and assigns it to the ConfigurationParameters field.
func (o *BTConfigurationUpdateCall2933) SetConfigurationParameters(v []BTMConfigurationParameter819) {
	o.ConfigurationParameters = v
}

// GetCurrentConfiguration returns the CurrentConfiguration field value if set, zero value otherwise.
func (o *BTConfigurationUpdateCall2933) GetCurrentConfiguration() []BTMParameter1 {
	if o == nil || o.CurrentConfiguration == nil {
		var ret []BTMParameter1
		return ret
	}
	return o.CurrentConfiguration
}

// GetCurrentConfigurationOk returns a tuple with the CurrentConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfigurationUpdateCall2933) GetCurrentConfigurationOk() ([]BTMParameter1, bool) {
	if o == nil || o.CurrentConfiguration == nil {
		return nil, false
	}
	return o.CurrentConfiguration, true
}

// HasCurrentConfiguration returns a boolean if a field has been set.
func (o *BTConfigurationUpdateCall2933) HasCurrentConfiguration() bool {
	if o != nil && o.CurrentConfiguration != nil {
		return true
	}

	return false
}

// SetCurrentConfiguration gets a reference to the given []BTMParameter1 and assigns it to the CurrentConfiguration field.
func (o *BTConfigurationUpdateCall2933) SetCurrentConfiguration(v []BTMParameter1) {
	o.CurrentConfiguration = v
}

// GetLibraryVersion returns the LibraryVersion field value if set, zero value otherwise.
func (o *BTConfigurationUpdateCall2933) GetLibraryVersion() int32 {
	if o == nil || o.LibraryVersion == nil {
		var ret int32
		return ret
	}
	return *o.LibraryVersion
}

// GetLibraryVersionOk returns a tuple with the LibraryVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfigurationUpdateCall2933) GetLibraryVersionOk() (*int32, bool) {
	if o == nil || o.LibraryVersion == nil {
		return nil, false
	}
	return o.LibraryVersion, true
}

// HasLibraryVersion returns a boolean if a field has been set.
func (o *BTConfigurationUpdateCall2933) HasLibraryVersion() bool {
	if o != nil && o.LibraryVersion != nil {
		return true
	}

	return false
}

// SetLibraryVersion gets a reference to the given int32 and assigns it to the LibraryVersion field.
func (o *BTConfigurationUpdateCall2933) SetLibraryVersion(v int32) {
	o.LibraryVersion = &v
}

// GetMicroversionSkew returns the MicroversionSkew field value if set, zero value otherwise.
func (o *BTConfigurationUpdateCall2933) GetMicroversionSkew() bool {
	if o == nil || o.MicroversionSkew == nil {
		var ret bool
		return ret
	}
	return *o.MicroversionSkew
}

// GetMicroversionSkewOk returns a tuple with the MicroversionSkew field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfigurationUpdateCall2933) GetMicroversionSkewOk() (*bool, bool) {
	if o == nil || o.MicroversionSkew == nil {
		return nil, false
	}
	return o.MicroversionSkew, true
}

// HasMicroversionSkew returns a boolean if a field has been set.
func (o *BTConfigurationUpdateCall2933) HasMicroversionSkew() bool {
	if o != nil && o.MicroversionSkew != nil {
		return true
	}

	return false
}

// SetMicroversionSkew gets a reference to the given bool and assigns it to the MicroversionSkew field.
func (o *BTConfigurationUpdateCall2933) SetMicroversionSkew(v bool) {
	o.MicroversionSkew = &v
}

// GetRejectMicroversionSkew returns the RejectMicroversionSkew field value if set, zero value otherwise.
func (o *BTConfigurationUpdateCall2933) GetRejectMicroversionSkew() bool {
	if o == nil || o.RejectMicroversionSkew == nil {
		var ret bool
		return ret
	}
	return *o.RejectMicroversionSkew
}

// GetRejectMicroversionSkewOk returns a tuple with the RejectMicroversionSkew field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfigurationUpdateCall2933) GetRejectMicroversionSkewOk() (*bool, bool) {
	if o == nil || o.RejectMicroversionSkew == nil {
		return nil, false
	}
	return o.RejectMicroversionSkew, true
}

// HasRejectMicroversionSkew returns a boolean if a field has been set.
func (o *BTConfigurationUpdateCall2933) HasRejectMicroversionSkew() bool {
	if o != nil && o.RejectMicroversionSkew != nil {
		return true
	}

	return false
}

// SetRejectMicroversionSkew gets a reference to the given bool and assigns it to the RejectMicroversionSkew field.
func (o *BTConfigurationUpdateCall2933) SetRejectMicroversionSkew(v bool) {
	o.RejectMicroversionSkew = &v
}

// GetSerializationVersion returns the SerializationVersion field value if set, zero value otherwise.
func (o *BTConfigurationUpdateCall2933) GetSerializationVersion() string {
	if o == nil || o.SerializationVersion == nil {
		var ret string
		return ret
	}
	return *o.SerializationVersion
}

// GetSerializationVersionOk returns a tuple with the SerializationVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfigurationUpdateCall2933) GetSerializationVersionOk() (*string, bool) {
	if o == nil || o.SerializationVersion == nil {
		return nil, false
	}
	return o.SerializationVersion, true
}

// HasSerializationVersion returns a boolean if a field has been set.
func (o *BTConfigurationUpdateCall2933) HasSerializationVersion() bool {
	if o != nil && o.SerializationVersion != nil {
		return true
	}

	return false
}

// SetSerializationVersion gets a reference to the given string and assigns it to the SerializationVersion field.
func (o *BTConfigurationUpdateCall2933) SetSerializationVersion(v string) {
	o.SerializationVersion = &v
}

// GetSourceMicroversion returns the SourceMicroversion field value if set, zero value otherwise.
func (o *BTConfigurationUpdateCall2933) GetSourceMicroversion() string {
	if o == nil || o.SourceMicroversion == nil {
		var ret string
		return ret
	}
	return *o.SourceMicroversion
}

// GetSourceMicroversionOk returns a tuple with the SourceMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfigurationUpdateCall2933) GetSourceMicroversionOk() (*string, bool) {
	if o == nil || o.SourceMicroversion == nil {
		return nil, false
	}
	return o.SourceMicroversion, true
}

// HasSourceMicroversion returns a boolean if a field has been set.
func (o *BTConfigurationUpdateCall2933) HasSourceMicroversion() bool {
	if o != nil && o.SourceMicroversion != nil {
		return true
	}

	return false
}

// SetSourceMicroversion gets a reference to the given string and assigns it to the SourceMicroversion field.
func (o *BTConfigurationUpdateCall2933) SetSourceMicroversion(v string) {
	o.SourceMicroversion = &v
}

func (o BTConfigurationUpdateCall2933) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ConfigurationParameters != nil {
		toSerialize["configurationParameters"] = o.ConfigurationParameters
	}
	if o.CurrentConfiguration != nil {
		toSerialize["currentConfiguration"] = o.CurrentConfiguration
	}
	if o.LibraryVersion != nil {
		toSerialize["libraryVersion"] = o.LibraryVersion
	}
	if o.MicroversionSkew != nil {
		toSerialize["microversionSkew"] = o.MicroversionSkew
	}
	if o.RejectMicroversionSkew != nil {
		toSerialize["rejectMicroversionSkew"] = o.RejectMicroversionSkew
	}
	if o.SerializationVersion != nil {
		toSerialize["serializationVersion"] = o.SerializationVersion
	}
	if o.SourceMicroversion != nil {
		toSerialize["sourceMicroversion"] = o.SourceMicroversion
	}
	return json.Marshal(toSerialize)
}

type NullableBTConfigurationUpdateCall2933 struct {
	value *BTConfigurationUpdateCall2933
	isSet bool
}

func (v NullableBTConfigurationUpdateCall2933) Get() *BTConfigurationUpdateCall2933 {
	return v.value
}

func (v *NullableBTConfigurationUpdateCall2933) Set(val *BTConfigurationUpdateCall2933) {
	v.value = val
	v.isSet = true
}

func (v NullableBTConfigurationUpdateCall2933) IsSet() bool {
	return v.isSet
}

func (v *NullableBTConfigurationUpdateCall2933) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTConfigurationUpdateCall2933(val *BTConfigurationUpdateCall2933) *NullableBTConfigurationUpdateCall2933 {
	return &NullableBTConfigurationUpdateCall2933{value: val, isSet: true}
}

func (v NullableBTConfigurationUpdateCall2933) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTConfigurationUpdateCall2933) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
