/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.154.6859-c8a2bd7d8338
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTAssemblyDefinitionInfo struct for BTAssemblyDefinitionInfo
type BTAssemblyDefinitionInfo struct {
	PartStudioFeatures []BTAssemblyPsFeatureInfo `json:"partStudioFeatures,omitempty"`
	Parts              []BTAssemblyPartInfo      `json:"parts,omitempty"`
	RootAssembly       *BTRootAssemblyInfo       `json:"rootAssembly,omitempty"`
	SubAssemblies      []BTSubAssemblyInfo       `json:"subAssemblies,omitempty"`
}

// NewBTAssemblyDefinitionInfo instantiates a new BTAssemblyDefinitionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAssemblyDefinitionInfo() *BTAssemblyDefinitionInfo {
	this := BTAssemblyDefinitionInfo{}
	return &this
}

// NewBTAssemblyDefinitionInfoWithDefaults instantiates a new BTAssemblyDefinitionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAssemblyDefinitionInfoWithDefaults() *BTAssemblyDefinitionInfo {
	this := BTAssemblyDefinitionInfo{}
	return &this
}

// GetPartStudioFeatures returns the PartStudioFeatures field value if set, zero value otherwise.
func (o *BTAssemblyDefinitionInfo) GetPartStudioFeatures() []BTAssemblyPsFeatureInfo {
	if o == nil || o.PartStudioFeatures == nil {
		var ret []BTAssemblyPsFeatureInfo
		return ret
	}
	return o.PartStudioFeatures
}

// GetPartStudioFeaturesOk returns a tuple with the PartStudioFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyDefinitionInfo) GetPartStudioFeaturesOk() ([]BTAssemblyPsFeatureInfo, bool) {
	if o == nil || o.PartStudioFeatures == nil {
		return nil, false
	}
	return o.PartStudioFeatures, true
}

// HasPartStudioFeatures returns a boolean if a field has been set.
func (o *BTAssemblyDefinitionInfo) HasPartStudioFeatures() bool {
	if o != nil && o.PartStudioFeatures != nil {
		return true
	}

	return false
}

// SetPartStudioFeatures gets a reference to the given []BTAssemblyPsFeatureInfo and assigns it to the PartStudioFeatures field.
func (o *BTAssemblyDefinitionInfo) SetPartStudioFeatures(v []BTAssemblyPsFeatureInfo) {
	o.PartStudioFeatures = v
}

// GetParts returns the Parts field value if set, zero value otherwise.
func (o *BTAssemblyDefinitionInfo) GetParts() []BTAssemblyPartInfo {
	if o == nil || o.Parts == nil {
		var ret []BTAssemblyPartInfo
		return ret
	}
	return o.Parts
}

// GetPartsOk returns a tuple with the Parts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyDefinitionInfo) GetPartsOk() ([]BTAssemblyPartInfo, bool) {
	if o == nil || o.Parts == nil {
		return nil, false
	}
	return o.Parts, true
}

// HasParts returns a boolean if a field has been set.
func (o *BTAssemblyDefinitionInfo) HasParts() bool {
	if o != nil && o.Parts != nil {
		return true
	}

	return false
}

// SetParts gets a reference to the given []BTAssemblyPartInfo and assigns it to the Parts field.
func (o *BTAssemblyDefinitionInfo) SetParts(v []BTAssemblyPartInfo) {
	o.Parts = v
}

// GetRootAssembly returns the RootAssembly field value if set, zero value otherwise.
func (o *BTAssemblyDefinitionInfo) GetRootAssembly() BTRootAssemblyInfo {
	if o == nil || o.RootAssembly == nil {
		var ret BTRootAssemblyInfo
		return ret
	}
	return *o.RootAssembly
}

// GetRootAssemblyOk returns a tuple with the RootAssembly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyDefinitionInfo) GetRootAssemblyOk() (*BTRootAssemblyInfo, bool) {
	if o == nil || o.RootAssembly == nil {
		return nil, false
	}
	return o.RootAssembly, true
}

// HasRootAssembly returns a boolean if a field has been set.
func (o *BTAssemblyDefinitionInfo) HasRootAssembly() bool {
	if o != nil && o.RootAssembly != nil {
		return true
	}

	return false
}

// SetRootAssembly gets a reference to the given BTRootAssemblyInfo and assigns it to the RootAssembly field.
func (o *BTAssemblyDefinitionInfo) SetRootAssembly(v BTRootAssemblyInfo) {
	o.RootAssembly = &v
}

// GetSubAssemblies returns the SubAssemblies field value if set, zero value otherwise.
func (o *BTAssemblyDefinitionInfo) GetSubAssemblies() []BTSubAssemblyInfo {
	if o == nil || o.SubAssemblies == nil {
		var ret []BTSubAssemblyInfo
		return ret
	}
	return o.SubAssemblies
}

// GetSubAssembliesOk returns a tuple with the SubAssemblies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyDefinitionInfo) GetSubAssembliesOk() ([]BTSubAssemblyInfo, bool) {
	if o == nil || o.SubAssemblies == nil {
		return nil, false
	}
	return o.SubAssemblies, true
}

// HasSubAssemblies returns a boolean if a field has been set.
func (o *BTAssemblyDefinitionInfo) HasSubAssemblies() bool {
	if o != nil && o.SubAssemblies != nil {
		return true
	}

	return false
}

// SetSubAssemblies gets a reference to the given []BTSubAssemblyInfo and assigns it to the SubAssemblies field.
func (o *BTAssemblyDefinitionInfo) SetSubAssemblies(v []BTSubAssemblyInfo) {
	o.SubAssemblies = v
}

func (o BTAssemblyDefinitionInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PartStudioFeatures != nil {
		toSerialize["partStudioFeatures"] = o.PartStudioFeatures
	}
	if o.Parts != nil {
		toSerialize["parts"] = o.Parts
	}
	if o.RootAssembly != nil {
		toSerialize["rootAssembly"] = o.RootAssembly
	}
	if o.SubAssemblies != nil {
		toSerialize["subAssemblies"] = o.SubAssemblies
	}
	return json.Marshal(toSerialize)
}

type NullableBTAssemblyDefinitionInfo struct {
	value *BTAssemblyDefinitionInfo
	isSet bool
}

func (v NullableBTAssemblyDefinitionInfo) Get() *BTAssemblyDefinitionInfo {
	return v.value
}

func (v *NullableBTAssemblyDefinitionInfo) Set(val *BTAssemblyDefinitionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAssemblyDefinitionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAssemblyDefinitionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAssemblyDefinitionInfo(val *BTAssemblyDefinitionInfo) *NullableBTAssemblyDefinitionInfo {
	return &NullableBTAssemblyDefinitionInfo{value: val, isSet: true}
}

func (v NullableBTAssemblyDefinitionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAssemblyDefinitionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
