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

// BTAssemblySimulationData978 struct for BTAssemblySimulationData978
type BTAssemblySimulationData978 struct {
	BtType             *string                    `json:"btType,omitempty"`
	ImportMicroversion *string                    `json:"importMicroversion,omitempty"`
	NodeId             *string                    `json:"nodeId,omitempty"`
	ContactBehavior    *string                    `json:"contactBehavior,omitempty"`
	Loads              []BTMLoad3538              `json:"loads,omitempty"`
	LoadsByNodeId      *map[string]BTMLoad3538    `json:"loadsByNodeId,omitempty"`
	Simulations        []BTAssemblySimulation2246 `json:"simulations,omitempty"`
	StructuralLoads    []BTMLoad3538              `json:"structuralLoads,omitempty"`
}

// NewBTAssemblySimulationData978 instantiates a new BTAssemblySimulationData978 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAssemblySimulationData978() *BTAssemblySimulationData978 {
	this := BTAssemblySimulationData978{}
	return &this
}

// NewBTAssemblySimulationData978WithDefaults instantiates a new BTAssemblySimulationData978 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAssemblySimulationData978WithDefaults() *BTAssemblySimulationData978 {
	this := BTAssemblySimulationData978{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTAssemblySimulationData978) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblySimulationData978) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTAssemblySimulationData978) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTAssemblySimulationData978) SetBtType(v string) {
	o.BtType = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTAssemblySimulationData978) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblySimulationData978) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTAssemblySimulationData978) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTAssemblySimulationData978) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTAssemblySimulationData978) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblySimulationData978) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTAssemblySimulationData978) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTAssemblySimulationData978) SetNodeId(v string) {
	o.NodeId = &v
}

// GetContactBehavior returns the ContactBehavior field value if set, zero value otherwise.
func (o *BTAssemblySimulationData978) GetContactBehavior() string {
	if o == nil || o.ContactBehavior == nil {
		var ret string
		return ret
	}
	return *o.ContactBehavior
}

// GetContactBehaviorOk returns a tuple with the ContactBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblySimulationData978) GetContactBehaviorOk() (*string, bool) {
	if o == nil || o.ContactBehavior == nil {
		return nil, false
	}
	return o.ContactBehavior, true
}

// HasContactBehavior returns a boolean if a field has been set.
func (o *BTAssemblySimulationData978) HasContactBehavior() bool {
	if o != nil && o.ContactBehavior != nil {
		return true
	}

	return false
}

// SetContactBehavior gets a reference to the given string and assigns it to the ContactBehavior field.
func (o *BTAssemblySimulationData978) SetContactBehavior(v string) {
	o.ContactBehavior = &v
}

// GetLoads returns the Loads field value if set, zero value otherwise.
func (o *BTAssemblySimulationData978) GetLoads() []BTMLoad3538 {
	if o == nil || o.Loads == nil {
		var ret []BTMLoad3538
		return ret
	}
	return o.Loads
}

// GetLoadsOk returns a tuple with the Loads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblySimulationData978) GetLoadsOk() ([]BTMLoad3538, bool) {
	if o == nil || o.Loads == nil {
		return nil, false
	}
	return o.Loads, true
}

// HasLoads returns a boolean if a field has been set.
func (o *BTAssemblySimulationData978) HasLoads() bool {
	if o != nil && o.Loads != nil {
		return true
	}

	return false
}

// SetLoads gets a reference to the given []BTMLoad3538 and assigns it to the Loads field.
func (o *BTAssemblySimulationData978) SetLoads(v []BTMLoad3538) {
	o.Loads = v
}

// GetLoadsByNodeId returns the LoadsByNodeId field value if set, zero value otherwise.
func (o *BTAssemblySimulationData978) GetLoadsByNodeId() map[string]BTMLoad3538 {
	if o == nil || o.LoadsByNodeId == nil {
		var ret map[string]BTMLoad3538
		return ret
	}
	return *o.LoadsByNodeId
}

// GetLoadsByNodeIdOk returns a tuple with the LoadsByNodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblySimulationData978) GetLoadsByNodeIdOk() (*map[string]BTMLoad3538, bool) {
	if o == nil || o.LoadsByNodeId == nil {
		return nil, false
	}
	return o.LoadsByNodeId, true
}

// HasLoadsByNodeId returns a boolean if a field has been set.
func (o *BTAssemblySimulationData978) HasLoadsByNodeId() bool {
	if o != nil && o.LoadsByNodeId != nil {
		return true
	}

	return false
}

// SetLoadsByNodeId gets a reference to the given map[string]BTMLoad3538 and assigns it to the LoadsByNodeId field.
func (o *BTAssemblySimulationData978) SetLoadsByNodeId(v map[string]BTMLoad3538) {
	o.LoadsByNodeId = &v
}

// GetSimulations returns the Simulations field value if set, zero value otherwise.
func (o *BTAssemblySimulationData978) GetSimulations() []BTAssemblySimulation2246 {
	if o == nil || o.Simulations == nil {
		var ret []BTAssemblySimulation2246
		return ret
	}
	return o.Simulations
}

// GetSimulationsOk returns a tuple with the Simulations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblySimulationData978) GetSimulationsOk() ([]BTAssemblySimulation2246, bool) {
	if o == nil || o.Simulations == nil {
		return nil, false
	}
	return o.Simulations, true
}

// HasSimulations returns a boolean if a field has been set.
func (o *BTAssemblySimulationData978) HasSimulations() bool {
	if o != nil && o.Simulations != nil {
		return true
	}

	return false
}

// SetSimulations gets a reference to the given []BTAssemblySimulation2246 and assigns it to the Simulations field.
func (o *BTAssemblySimulationData978) SetSimulations(v []BTAssemblySimulation2246) {
	o.Simulations = v
}

// GetStructuralLoads returns the StructuralLoads field value if set, zero value otherwise.
func (o *BTAssemblySimulationData978) GetStructuralLoads() []BTMLoad3538 {
	if o == nil || o.StructuralLoads == nil {
		var ret []BTMLoad3538
		return ret
	}
	return o.StructuralLoads
}

// GetStructuralLoadsOk returns a tuple with the StructuralLoads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblySimulationData978) GetStructuralLoadsOk() ([]BTMLoad3538, bool) {
	if o == nil || o.StructuralLoads == nil {
		return nil, false
	}
	return o.StructuralLoads, true
}

// HasStructuralLoads returns a boolean if a field has been set.
func (o *BTAssemblySimulationData978) HasStructuralLoads() bool {
	if o != nil && o.StructuralLoads != nil {
		return true
	}

	return false
}

// SetStructuralLoads gets a reference to the given []BTMLoad3538 and assigns it to the StructuralLoads field.
func (o *BTAssemblySimulationData978) SetStructuralLoads(v []BTMLoad3538) {
	o.StructuralLoads = v
}

func (o BTAssemblySimulationData978) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.ContactBehavior != nil {
		toSerialize["contactBehavior"] = o.ContactBehavior
	}
	if o.Loads != nil {
		toSerialize["loads"] = o.Loads
	}
	if o.LoadsByNodeId != nil {
		toSerialize["loadsByNodeId"] = o.LoadsByNodeId
	}
	if o.Simulations != nil {
		toSerialize["simulations"] = o.Simulations
	}
	if o.StructuralLoads != nil {
		toSerialize["structuralLoads"] = o.StructuralLoads
	}
	return json.Marshal(toSerialize)
}

type NullableBTAssemblySimulationData978 struct {
	value *BTAssemblySimulationData978
	isSet bool
}

func (v NullableBTAssemblySimulationData978) Get() *BTAssemblySimulationData978 {
	return v.value
}

func (v *NullableBTAssemblySimulationData978) Set(val *BTAssemblySimulationData978) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAssemblySimulationData978) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAssemblySimulationData978) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAssemblySimulationData978(val *BTAssemblySimulationData978) *NullableBTAssemblySimulationData978 {
	return &NullableBTAssemblySimulationData978{value: val, isSet: true}
}

func (v NullableBTAssemblySimulationData978) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAssemblySimulationData978) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
