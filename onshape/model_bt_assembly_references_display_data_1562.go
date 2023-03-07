/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.160.12410-b0c73c1032e8
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTAssemblyReferencesDisplayData1562 struct for BTAssemblyReferencesDisplayData1562
type BTAssemblyReferencesDisplayData1562 struct {
	AssemblyReferences                     []BTSingleReferenceDisplayData1943            `json:"assemblyReferences,omitempty"`
	ElementId                              *string                                       `json:"elementId,omitempty"`
	FromFullElementId                      *BTFullElementId756                           `json:"fromFullElementId,omitempty"`
	FullElementId                          *BTFullElementId756                           `json:"fullElementId,omitempty"`
	Incremental                            *bool                                         `json:"incremental,omitempty"`
	InstanceCount                          *int32                                        `json:"instanceCount,omitempty"`
	KeepFromMicroversion                   *bool                                         `json:"keepFromMicroversion,omitempty"`
	MicroversionId                         *BTMicroversionId366                          `json:"microversionId,omitempty"`
	MicroversionIdAndConfigurationInterval *BTMicroversionIdAndConfigurationInterval2364 `json:"microversionIdAndConfigurationInterval,omitempty"`
	MicroversionInterval                   *BTMicroversionIdInterval367                  `json:"microversionInterval,omitempty"`
	VersionForRasterization                *BTElementDisplayData326                      `json:"versionForRasterization,omitempty"`
}

// NewBTAssemblyReferencesDisplayData1562 instantiates a new BTAssemblyReferencesDisplayData1562 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAssemblyReferencesDisplayData1562() *BTAssemblyReferencesDisplayData1562 {
	this := BTAssemblyReferencesDisplayData1562{}
	return &this
}

// NewBTAssemblyReferencesDisplayData1562WithDefaults instantiates a new BTAssemblyReferencesDisplayData1562 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAssemblyReferencesDisplayData1562WithDefaults() *BTAssemblyReferencesDisplayData1562 {
	this := BTAssemblyReferencesDisplayData1562{}
	return &this
}

// GetAssemblyReferences returns the AssemblyReferences field value if set, zero value otherwise.
func (o *BTAssemblyReferencesDisplayData1562) GetAssemblyReferences() []BTSingleReferenceDisplayData1943 {
	if o == nil || o.AssemblyReferences == nil {
		var ret []BTSingleReferenceDisplayData1943
		return ret
	}
	return o.AssemblyReferences
}

// GetAssemblyReferencesOk returns a tuple with the AssemblyReferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyReferencesDisplayData1562) GetAssemblyReferencesOk() ([]BTSingleReferenceDisplayData1943, bool) {
	if o == nil || o.AssemblyReferences == nil {
		return nil, false
	}
	return o.AssemblyReferences, true
}

// HasAssemblyReferences returns a boolean if a field has been set.
func (o *BTAssemblyReferencesDisplayData1562) HasAssemblyReferences() bool {
	if o != nil && o.AssemblyReferences != nil {
		return true
	}

	return false
}

// SetAssemblyReferences gets a reference to the given []BTSingleReferenceDisplayData1943 and assigns it to the AssemblyReferences field.
func (o *BTAssemblyReferencesDisplayData1562) SetAssemblyReferences(v []BTSingleReferenceDisplayData1943) {
	o.AssemblyReferences = v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *BTAssemblyReferencesDisplayData1562) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyReferencesDisplayData1562) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *BTAssemblyReferencesDisplayData1562) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *BTAssemblyReferencesDisplayData1562) SetElementId(v string) {
	o.ElementId = &v
}

// GetFromFullElementId returns the FromFullElementId field value if set, zero value otherwise.
func (o *BTAssemblyReferencesDisplayData1562) GetFromFullElementId() BTFullElementId756 {
	if o == nil || o.FromFullElementId == nil {
		var ret BTFullElementId756
		return ret
	}
	return *o.FromFullElementId
}

// GetFromFullElementIdOk returns a tuple with the FromFullElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyReferencesDisplayData1562) GetFromFullElementIdOk() (*BTFullElementId756, bool) {
	if o == nil || o.FromFullElementId == nil {
		return nil, false
	}
	return o.FromFullElementId, true
}

// HasFromFullElementId returns a boolean if a field has been set.
func (o *BTAssemblyReferencesDisplayData1562) HasFromFullElementId() bool {
	if o != nil && o.FromFullElementId != nil {
		return true
	}

	return false
}

// SetFromFullElementId gets a reference to the given BTFullElementId756 and assigns it to the FromFullElementId field.
func (o *BTAssemblyReferencesDisplayData1562) SetFromFullElementId(v BTFullElementId756) {
	o.FromFullElementId = &v
}

// GetFullElementId returns the FullElementId field value if set, zero value otherwise.
func (o *BTAssemblyReferencesDisplayData1562) GetFullElementId() BTFullElementId756 {
	if o == nil || o.FullElementId == nil {
		var ret BTFullElementId756
		return ret
	}
	return *o.FullElementId
}

// GetFullElementIdOk returns a tuple with the FullElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyReferencesDisplayData1562) GetFullElementIdOk() (*BTFullElementId756, bool) {
	if o == nil || o.FullElementId == nil {
		return nil, false
	}
	return o.FullElementId, true
}

// HasFullElementId returns a boolean if a field has been set.
func (o *BTAssemblyReferencesDisplayData1562) HasFullElementId() bool {
	if o != nil && o.FullElementId != nil {
		return true
	}

	return false
}

// SetFullElementId gets a reference to the given BTFullElementId756 and assigns it to the FullElementId field.
func (o *BTAssemblyReferencesDisplayData1562) SetFullElementId(v BTFullElementId756) {
	o.FullElementId = &v
}

// GetIncremental returns the Incremental field value if set, zero value otherwise.
func (o *BTAssemblyReferencesDisplayData1562) GetIncremental() bool {
	if o == nil || o.Incremental == nil {
		var ret bool
		return ret
	}
	return *o.Incremental
}

// GetIncrementalOk returns a tuple with the Incremental field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyReferencesDisplayData1562) GetIncrementalOk() (*bool, bool) {
	if o == nil || o.Incremental == nil {
		return nil, false
	}
	return o.Incremental, true
}

// HasIncremental returns a boolean if a field has been set.
func (o *BTAssemblyReferencesDisplayData1562) HasIncremental() bool {
	if o != nil && o.Incremental != nil {
		return true
	}

	return false
}

// SetIncremental gets a reference to the given bool and assigns it to the Incremental field.
func (o *BTAssemblyReferencesDisplayData1562) SetIncremental(v bool) {
	o.Incremental = &v
}

// GetInstanceCount returns the InstanceCount field value if set, zero value otherwise.
func (o *BTAssemblyReferencesDisplayData1562) GetInstanceCount() int32 {
	if o == nil || o.InstanceCount == nil {
		var ret int32
		return ret
	}
	return *o.InstanceCount
}

// GetInstanceCountOk returns a tuple with the InstanceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyReferencesDisplayData1562) GetInstanceCountOk() (*int32, bool) {
	if o == nil || o.InstanceCount == nil {
		return nil, false
	}
	return o.InstanceCount, true
}

// HasInstanceCount returns a boolean if a field has been set.
func (o *BTAssemblyReferencesDisplayData1562) HasInstanceCount() bool {
	if o != nil && o.InstanceCount != nil {
		return true
	}

	return false
}

// SetInstanceCount gets a reference to the given int32 and assigns it to the InstanceCount field.
func (o *BTAssemblyReferencesDisplayData1562) SetInstanceCount(v int32) {
	o.InstanceCount = &v
}

// GetKeepFromMicroversion returns the KeepFromMicroversion field value if set, zero value otherwise.
func (o *BTAssemblyReferencesDisplayData1562) GetKeepFromMicroversion() bool {
	if o == nil || o.KeepFromMicroversion == nil {
		var ret bool
		return ret
	}
	return *o.KeepFromMicroversion
}

// GetKeepFromMicroversionOk returns a tuple with the KeepFromMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyReferencesDisplayData1562) GetKeepFromMicroversionOk() (*bool, bool) {
	if o == nil || o.KeepFromMicroversion == nil {
		return nil, false
	}
	return o.KeepFromMicroversion, true
}

// HasKeepFromMicroversion returns a boolean if a field has been set.
func (o *BTAssemblyReferencesDisplayData1562) HasKeepFromMicroversion() bool {
	if o != nil && o.KeepFromMicroversion != nil {
		return true
	}

	return false
}

// SetKeepFromMicroversion gets a reference to the given bool and assigns it to the KeepFromMicroversion field.
func (o *BTAssemblyReferencesDisplayData1562) SetKeepFromMicroversion(v bool) {
	o.KeepFromMicroversion = &v
}

// GetMicroversionId returns the MicroversionId field value if set, zero value otherwise.
func (o *BTAssemblyReferencesDisplayData1562) GetMicroversionId() BTMicroversionId366 {
	if o == nil || o.MicroversionId == nil {
		var ret BTMicroversionId366
		return ret
	}
	return *o.MicroversionId
}

// GetMicroversionIdOk returns a tuple with the MicroversionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyReferencesDisplayData1562) GetMicroversionIdOk() (*BTMicroversionId366, bool) {
	if o == nil || o.MicroversionId == nil {
		return nil, false
	}
	return o.MicroversionId, true
}

// HasMicroversionId returns a boolean if a field has been set.
func (o *BTAssemblyReferencesDisplayData1562) HasMicroversionId() bool {
	if o != nil && o.MicroversionId != nil {
		return true
	}

	return false
}

// SetMicroversionId gets a reference to the given BTMicroversionId366 and assigns it to the MicroversionId field.
func (o *BTAssemblyReferencesDisplayData1562) SetMicroversionId(v BTMicroversionId366) {
	o.MicroversionId = &v
}

// GetMicroversionIdAndConfigurationInterval returns the MicroversionIdAndConfigurationInterval field value if set, zero value otherwise.
func (o *BTAssemblyReferencesDisplayData1562) GetMicroversionIdAndConfigurationInterval() BTMicroversionIdAndConfigurationInterval2364 {
	if o == nil || o.MicroversionIdAndConfigurationInterval == nil {
		var ret BTMicroversionIdAndConfigurationInterval2364
		return ret
	}
	return *o.MicroversionIdAndConfigurationInterval
}

// GetMicroversionIdAndConfigurationIntervalOk returns a tuple with the MicroversionIdAndConfigurationInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyReferencesDisplayData1562) GetMicroversionIdAndConfigurationIntervalOk() (*BTMicroversionIdAndConfigurationInterval2364, bool) {
	if o == nil || o.MicroversionIdAndConfigurationInterval == nil {
		return nil, false
	}
	return o.MicroversionIdAndConfigurationInterval, true
}

// HasMicroversionIdAndConfigurationInterval returns a boolean if a field has been set.
func (o *BTAssemblyReferencesDisplayData1562) HasMicroversionIdAndConfigurationInterval() bool {
	if o != nil && o.MicroversionIdAndConfigurationInterval != nil {
		return true
	}

	return false
}

// SetMicroversionIdAndConfigurationInterval gets a reference to the given BTMicroversionIdAndConfigurationInterval2364 and assigns it to the MicroversionIdAndConfigurationInterval field.
func (o *BTAssemblyReferencesDisplayData1562) SetMicroversionIdAndConfigurationInterval(v BTMicroversionIdAndConfigurationInterval2364) {
	o.MicroversionIdAndConfigurationInterval = &v
}

// GetMicroversionInterval returns the MicroversionInterval field value if set, zero value otherwise.
func (o *BTAssemblyReferencesDisplayData1562) GetMicroversionInterval() BTMicroversionIdInterval367 {
	if o == nil || o.MicroversionInterval == nil {
		var ret BTMicroversionIdInterval367
		return ret
	}
	return *o.MicroversionInterval
}

// GetMicroversionIntervalOk returns a tuple with the MicroversionInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyReferencesDisplayData1562) GetMicroversionIntervalOk() (*BTMicroversionIdInterval367, bool) {
	if o == nil || o.MicroversionInterval == nil {
		return nil, false
	}
	return o.MicroversionInterval, true
}

// HasMicroversionInterval returns a boolean if a field has been set.
func (o *BTAssemblyReferencesDisplayData1562) HasMicroversionInterval() bool {
	if o != nil && o.MicroversionInterval != nil {
		return true
	}

	return false
}

// SetMicroversionInterval gets a reference to the given BTMicroversionIdInterval367 and assigns it to the MicroversionInterval field.
func (o *BTAssemblyReferencesDisplayData1562) SetMicroversionInterval(v BTMicroversionIdInterval367) {
	o.MicroversionInterval = &v
}

// GetVersionForRasterization returns the VersionForRasterization field value if set, zero value otherwise.
func (o *BTAssemblyReferencesDisplayData1562) GetVersionForRasterization() BTElementDisplayData326 {
	if o == nil || o.VersionForRasterization == nil {
		var ret BTElementDisplayData326
		return ret
	}
	return *o.VersionForRasterization
}

// GetVersionForRasterizationOk returns a tuple with the VersionForRasterization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyReferencesDisplayData1562) GetVersionForRasterizationOk() (*BTElementDisplayData326, bool) {
	if o == nil || o.VersionForRasterization == nil {
		return nil, false
	}
	return o.VersionForRasterization, true
}

// HasVersionForRasterization returns a boolean if a field has been set.
func (o *BTAssemblyReferencesDisplayData1562) HasVersionForRasterization() bool {
	if o != nil && o.VersionForRasterization != nil {
		return true
	}

	return false
}

// SetVersionForRasterization gets a reference to the given BTElementDisplayData326 and assigns it to the VersionForRasterization field.
func (o *BTAssemblyReferencesDisplayData1562) SetVersionForRasterization(v BTElementDisplayData326) {
	o.VersionForRasterization = &v
}

func (o BTAssemblyReferencesDisplayData1562) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AssemblyReferences != nil {
		toSerialize["assemblyReferences"] = o.AssemblyReferences
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.FromFullElementId != nil {
		toSerialize["fromFullElementId"] = o.FromFullElementId
	}
	if o.FullElementId != nil {
		toSerialize["fullElementId"] = o.FullElementId
	}
	if o.Incremental != nil {
		toSerialize["incremental"] = o.Incremental
	}
	if o.InstanceCount != nil {
		toSerialize["instanceCount"] = o.InstanceCount
	}
	if o.KeepFromMicroversion != nil {
		toSerialize["keepFromMicroversion"] = o.KeepFromMicroversion
	}
	if o.MicroversionId != nil {
		toSerialize["microversionId"] = o.MicroversionId
	}
	if o.MicroversionIdAndConfigurationInterval != nil {
		toSerialize["microversionIdAndConfigurationInterval"] = o.MicroversionIdAndConfigurationInterval
	}
	if o.MicroversionInterval != nil {
		toSerialize["microversionInterval"] = o.MicroversionInterval
	}
	if o.VersionForRasterization != nil {
		toSerialize["versionForRasterization"] = o.VersionForRasterization
	}
	return json.Marshal(toSerialize)
}

type NullableBTAssemblyReferencesDisplayData1562 struct {
	value *BTAssemblyReferencesDisplayData1562
	isSet bool
}

func (v NullableBTAssemblyReferencesDisplayData1562) Get() *BTAssemblyReferencesDisplayData1562 {
	return v.value
}

func (v *NullableBTAssemblyReferencesDisplayData1562) Set(val *BTAssemblyReferencesDisplayData1562) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAssemblyReferencesDisplayData1562) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAssemblyReferencesDisplayData1562) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAssemblyReferencesDisplayData1562(val *BTAssemblyReferencesDisplayData1562) *NullableBTAssemblyReferencesDisplayData1562 {
	return &NullableBTAssemblyReferencesDisplayData1562{value: val, isSet: true}
}

func (v NullableBTAssemblyReferencesDisplayData1562) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAssemblyReferencesDisplayData1562) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
