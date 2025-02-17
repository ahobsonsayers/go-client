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

// BTAssemblyPatternInfo Pattern description.
type BTAssemblyPatternInfo struct {
	// Id of the pattern.
	Id *string `json:"id,omitempty"`
	// Name of the pattern.
	Name *string `json:"name,omitempty"`
	// Mapping of seed to pattern instance ids.
	SeedToPatternInstances *map[string][]string `json:"seedToPatternInstances,omitempty"`
	// If pattern is suppressed.
	Suppressed *bool           `json:"suppressed,omitempty"`
	Type       *GBTPatternType `json:"type,omitempty"`
}

// NewBTAssemblyPatternInfo instantiates a new BTAssemblyPatternInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAssemblyPatternInfo() *BTAssemblyPatternInfo {
	this := BTAssemblyPatternInfo{}
	return &this
}

// NewBTAssemblyPatternInfoWithDefaults instantiates a new BTAssemblyPatternInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAssemblyPatternInfoWithDefaults() *BTAssemblyPatternInfo {
	this := BTAssemblyPatternInfo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTAssemblyPatternInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyPatternInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTAssemblyPatternInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTAssemblyPatternInfo) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTAssemblyPatternInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyPatternInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTAssemblyPatternInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTAssemblyPatternInfo) SetName(v string) {
	o.Name = &v
}

// GetSeedToPatternInstances returns the SeedToPatternInstances field value if set, zero value otherwise.
func (o *BTAssemblyPatternInfo) GetSeedToPatternInstances() map[string][]string {
	if o == nil || o.SeedToPatternInstances == nil {
		var ret map[string][]string
		return ret
	}
	return *o.SeedToPatternInstances
}

// GetSeedToPatternInstancesOk returns a tuple with the SeedToPatternInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyPatternInfo) GetSeedToPatternInstancesOk() (*map[string][]string, bool) {
	if o == nil || o.SeedToPatternInstances == nil {
		return nil, false
	}
	return o.SeedToPatternInstances, true
}

// HasSeedToPatternInstances returns a boolean if a field has been set.
func (o *BTAssemblyPatternInfo) HasSeedToPatternInstances() bool {
	if o != nil && o.SeedToPatternInstances != nil {
		return true
	}

	return false
}

// SetSeedToPatternInstances gets a reference to the given map[string][]string and assigns it to the SeedToPatternInstances field.
func (o *BTAssemblyPatternInfo) SetSeedToPatternInstances(v map[string][]string) {
	o.SeedToPatternInstances = &v
}

// GetSuppressed returns the Suppressed field value if set, zero value otherwise.
func (o *BTAssemblyPatternInfo) GetSuppressed() bool {
	if o == nil || o.Suppressed == nil {
		var ret bool
		return ret
	}
	return *o.Suppressed
}

// GetSuppressedOk returns a tuple with the Suppressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyPatternInfo) GetSuppressedOk() (*bool, bool) {
	if o == nil || o.Suppressed == nil {
		return nil, false
	}
	return o.Suppressed, true
}

// HasSuppressed returns a boolean if a field has been set.
func (o *BTAssemblyPatternInfo) HasSuppressed() bool {
	if o != nil && o.Suppressed != nil {
		return true
	}

	return false
}

// SetSuppressed gets a reference to the given bool and assigns it to the Suppressed field.
func (o *BTAssemblyPatternInfo) SetSuppressed(v bool) {
	o.Suppressed = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BTAssemblyPatternInfo) GetType() GBTPatternType {
	if o == nil || o.Type == nil {
		var ret GBTPatternType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyPatternInfo) GetTypeOk() (*GBTPatternType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BTAssemblyPatternInfo) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given GBTPatternType and assigns it to the Type field.
func (o *BTAssemblyPatternInfo) SetType(v GBTPatternType) {
	o.Type = &v
}

func (o BTAssemblyPatternInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.SeedToPatternInstances != nil {
		toSerialize["seedToPatternInstances"] = o.SeedToPatternInstances
	}
	if o.Suppressed != nil {
		toSerialize["suppressed"] = o.Suppressed
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableBTAssemblyPatternInfo struct {
	value *BTAssemblyPatternInfo
	isSet bool
}

func (v NullableBTAssemblyPatternInfo) Get() *BTAssemblyPatternInfo {
	return v.value
}

func (v *NullableBTAssemblyPatternInfo) Set(val *BTAssemblyPatternInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAssemblyPatternInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAssemblyPatternInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAssemblyPatternInfo(val *BTAssemblyPatternInfo) *NullableBTAssemblyPatternInfo {
	return &NullableBTAssemblyPatternInfo{value: val, isSet: true}
}

func (v NullableBTAssemblyPatternInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAssemblyPatternInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
