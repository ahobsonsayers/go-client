/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.157.8827-f82e65cdc920
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMSketchConstraint2 struct for BTMSketchConstraint2
type BTMSketchConstraint2 struct {
	BtType                              *string         `json:"btType,omitempty"`
	ConstraintType                      *string         `json:"constraintType,omitempty"`
	DrivenDimension                     *bool           `json:"drivenDimension,omitempty"`
	EntityId                            *string         `json:"entityId,omitempty"`
	EntityIdAndReplaceInDependentFields *string         `json:"entityIdAndReplaceInDependentFields,omitempty"`
	HasOffsetData1                      *bool           `json:"hasOffsetData1,omitempty"`
	HasOffsetData2                      *bool           `json:"hasOffsetData2,omitempty"`
	HasPierceParameter_                 *bool           `json:"hasPierceParameter,omitempty"`
	HelpParameters                      []float64       `json:"helpParameters,omitempty"`
	ImportMicroversion                  *string         `json:"importMicroversion,omitempty"`
	Namespace                           *string         `json:"namespace,omitempty"`
	NodeId                              *string         `json:"nodeId,omitempty"`
	OffsetDistance1                     *float64        `json:"offsetDistance1,omitempty"`
	OffsetDistance2                     *float64        `json:"offsetDistance2,omitempty"`
	OffsetOrientation1                  *bool           `json:"offsetOrientation1,omitempty"`
	OffsetOrientation2                  *bool           `json:"offsetOrientation2,omitempty"`
	Parameters                          []BTMParameter1 `json:"parameters,omitempty"`
	PierceParameter                     *float64        `json:"pierceParameter,omitempty"`
}

// NewBTMSketchConstraint2 instantiates a new BTMSketchConstraint2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMSketchConstraint2() *BTMSketchConstraint2 {
	this := BTMSketchConstraint2{}
	return &this
}

// NewBTMSketchConstraint2WithDefaults instantiates a new BTMSketchConstraint2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMSketchConstraint2WithDefaults() *BTMSketchConstraint2 {
	this := BTMSketchConstraint2{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMSketchConstraint2) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchConstraint2) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMSketchConstraint2) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMSketchConstraint2) SetBtType(v string) {
	o.BtType = &v
}

// GetConstraintType returns the ConstraintType field value if set, zero value otherwise.
func (o *BTMSketchConstraint2) GetConstraintType() string {
	if o == nil || o.ConstraintType == nil {
		var ret string
		return ret
	}
	return *o.ConstraintType
}

// GetConstraintTypeOk returns a tuple with the ConstraintType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchConstraint2) GetConstraintTypeOk() (*string, bool) {
	if o == nil || o.ConstraintType == nil {
		return nil, false
	}
	return o.ConstraintType, true
}

// HasConstraintType returns a boolean if a field has been set.
func (o *BTMSketchConstraint2) HasConstraintType() bool {
	if o != nil && o.ConstraintType != nil {
		return true
	}

	return false
}

// SetConstraintType gets a reference to the given string and assigns it to the ConstraintType field.
func (o *BTMSketchConstraint2) SetConstraintType(v string) {
	o.ConstraintType = &v
}

// GetDrivenDimension returns the DrivenDimension field value if set, zero value otherwise.
func (o *BTMSketchConstraint2) GetDrivenDimension() bool {
	if o == nil || o.DrivenDimension == nil {
		var ret bool
		return ret
	}
	return *o.DrivenDimension
}

// GetDrivenDimensionOk returns a tuple with the DrivenDimension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchConstraint2) GetDrivenDimensionOk() (*bool, bool) {
	if o == nil || o.DrivenDimension == nil {
		return nil, false
	}
	return o.DrivenDimension, true
}

// HasDrivenDimension returns a boolean if a field has been set.
func (o *BTMSketchConstraint2) HasDrivenDimension() bool {
	if o != nil && o.DrivenDimension != nil {
		return true
	}

	return false
}

// SetDrivenDimension gets a reference to the given bool and assigns it to the DrivenDimension field.
func (o *BTMSketchConstraint2) SetDrivenDimension(v bool) {
	o.DrivenDimension = &v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *BTMSketchConstraint2) GetEntityId() string {
	if o == nil || o.EntityId == nil {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchConstraint2) GetEntityIdOk() (*string, bool) {
	if o == nil || o.EntityId == nil {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *BTMSketchConstraint2) HasEntityId() bool {
	if o != nil && o.EntityId != nil {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *BTMSketchConstraint2) SetEntityId(v string) {
	o.EntityId = &v
}

// GetEntityIdAndReplaceInDependentFields returns the EntityIdAndReplaceInDependentFields field value if set, zero value otherwise.
func (o *BTMSketchConstraint2) GetEntityIdAndReplaceInDependentFields() string {
	if o == nil || o.EntityIdAndReplaceInDependentFields == nil {
		var ret string
		return ret
	}
	return *o.EntityIdAndReplaceInDependentFields
}

// GetEntityIdAndReplaceInDependentFieldsOk returns a tuple with the EntityIdAndReplaceInDependentFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchConstraint2) GetEntityIdAndReplaceInDependentFieldsOk() (*string, bool) {
	if o == nil || o.EntityIdAndReplaceInDependentFields == nil {
		return nil, false
	}
	return o.EntityIdAndReplaceInDependentFields, true
}

// HasEntityIdAndReplaceInDependentFields returns a boolean if a field has been set.
func (o *BTMSketchConstraint2) HasEntityIdAndReplaceInDependentFields() bool {
	if o != nil && o.EntityIdAndReplaceInDependentFields != nil {
		return true
	}

	return false
}

// SetEntityIdAndReplaceInDependentFields gets a reference to the given string and assigns it to the EntityIdAndReplaceInDependentFields field.
func (o *BTMSketchConstraint2) SetEntityIdAndReplaceInDependentFields(v string) {
	o.EntityIdAndReplaceInDependentFields = &v
}

// GetHasOffsetData1 returns the HasOffsetData1 field value if set, zero value otherwise.
func (o *BTMSketchConstraint2) GetHasOffsetData1() bool {
	if o == nil || o.HasOffsetData1 == nil {
		var ret bool
		return ret
	}
	return *o.HasOffsetData1
}

// GetHasOffsetData1Ok returns a tuple with the HasOffsetData1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchConstraint2) GetHasOffsetData1Ok() (*bool, bool) {
	if o == nil || o.HasOffsetData1 == nil {
		return nil, false
	}
	return o.HasOffsetData1, true
}

// HasHasOffsetData1 returns a boolean if a field has been set.
func (o *BTMSketchConstraint2) HasHasOffsetData1() bool {
	if o != nil && o.HasOffsetData1 != nil {
		return true
	}

	return false
}

// SetHasOffsetData1 gets a reference to the given bool and assigns it to the HasOffsetData1 field.
func (o *BTMSketchConstraint2) SetHasOffsetData1(v bool) {
	o.HasOffsetData1 = &v
}

// GetHasOffsetData2 returns the HasOffsetData2 field value if set, zero value otherwise.
func (o *BTMSketchConstraint2) GetHasOffsetData2() bool {
	if o == nil || o.HasOffsetData2 == nil {
		var ret bool
		return ret
	}
	return *o.HasOffsetData2
}

// GetHasOffsetData2Ok returns a tuple with the HasOffsetData2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchConstraint2) GetHasOffsetData2Ok() (*bool, bool) {
	if o == nil || o.HasOffsetData2 == nil {
		return nil, false
	}
	return o.HasOffsetData2, true
}

// HasHasOffsetData2 returns a boolean if a field has been set.
func (o *BTMSketchConstraint2) HasHasOffsetData2() bool {
	if o != nil && o.HasOffsetData2 != nil {
		return true
	}

	return false
}

// SetHasOffsetData2 gets a reference to the given bool and assigns it to the HasOffsetData2 field.
func (o *BTMSketchConstraint2) SetHasOffsetData2(v bool) {
	o.HasOffsetData2 = &v
}

// GetHasPierceParameter_ returns the HasPierceParameter_ field value if set, zero value otherwise.
func (o *BTMSketchConstraint2) GetHasPierceParameter_() bool {
	if o == nil || o.HasPierceParameter_ == nil {
		var ret bool
		return ret
	}
	return *o.HasPierceParameter_
}

// GetHasPierceParameter_Ok returns a tuple with the HasPierceParameter_ field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchConstraint2) GetHasPierceParameter_Ok() (*bool, bool) {
	if o == nil || o.HasPierceParameter_ == nil {
		return nil, false
	}
	return o.HasPierceParameter_, true
}

// HasHasPierceParameter_ returns a boolean if a field has been set.
func (o *BTMSketchConstraint2) HasHasPierceParameter_() bool {
	if o != nil && o.HasPierceParameter_ != nil {
		return true
	}

	return false
}

// SetHasPierceParameter_ gets a reference to the given bool and assigns it to the HasPierceParameter_ field.
func (o *BTMSketchConstraint2) SetHasPierceParameter_(v bool) {
	o.HasPierceParameter_ = &v
}

// GetHelpParameters returns the HelpParameters field value if set, zero value otherwise.
func (o *BTMSketchConstraint2) GetHelpParameters() []float64 {
	if o == nil || o.HelpParameters == nil {
		var ret []float64
		return ret
	}
	return o.HelpParameters
}

// GetHelpParametersOk returns a tuple with the HelpParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchConstraint2) GetHelpParametersOk() ([]float64, bool) {
	if o == nil || o.HelpParameters == nil {
		return nil, false
	}
	return o.HelpParameters, true
}

// HasHelpParameters returns a boolean if a field has been set.
func (o *BTMSketchConstraint2) HasHelpParameters() bool {
	if o != nil && o.HelpParameters != nil {
		return true
	}

	return false
}

// SetHelpParameters gets a reference to the given []float64 and assigns it to the HelpParameters field.
func (o *BTMSketchConstraint2) SetHelpParameters(v []float64) {
	o.HelpParameters = v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMSketchConstraint2) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchConstraint2) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMSketchConstraint2) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMSketchConstraint2) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *BTMSketchConstraint2) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchConstraint2) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *BTMSketchConstraint2) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *BTMSketchConstraint2) SetNamespace(v string) {
	o.Namespace = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMSketchConstraint2) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchConstraint2) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMSketchConstraint2) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMSketchConstraint2) SetNodeId(v string) {
	o.NodeId = &v
}

// GetOffsetDistance1 returns the OffsetDistance1 field value if set, zero value otherwise.
func (o *BTMSketchConstraint2) GetOffsetDistance1() float64 {
	if o == nil || o.OffsetDistance1 == nil {
		var ret float64
		return ret
	}
	return *o.OffsetDistance1
}

// GetOffsetDistance1Ok returns a tuple with the OffsetDistance1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchConstraint2) GetOffsetDistance1Ok() (*float64, bool) {
	if o == nil || o.OffsetDistance1 == nil {
		return nil, false
	}
	return o.OffsetDistance1, true
}

// HasOffsetDistance1 returns a boolean if a field has been set.
func (o *BTMSketchConstraint2) HasOffsetDistance1() bool {
	if o != nil && o.OffsetDistance1 != nil {
		return true
	}

	return false
}

// SetOffsetDistance1 gets a reference to the given float64 and assigns it to the OffsetDistance1 field.
func (o *BTMSketchConstraint2) SetOffsetDistance1(v float64) {
	o.OffsetDistance1 = &v
}

// GetOffsetDistance2 returns the OffsetDistance2 field value if set, zero value otherwise.
func (o *BTMSketchConstraint2) GetOffsetDistance2() float64 {
	if o == nil || o.OffsetDistance2 == nil {
		var ret float64
		return ret
	}
	return *o.OffsetDistance2
}

// GetOffsetDistance2Ok returns a tuple with the OffsetDistance2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchConstraint2) GetOffsetDistance2Ok() (*float64, bool) {
	if o == nil || o.OffsetDistance2 == nil {
		return nil, false
	}
	return o.OffsetDistance2, true
}

// HasOffsetDistance2 returns a boolean if a field has been set.
func (o *BTMSketchConstraint2) HasOffsetDistance2() bool {
	if o != nil && o.OffsetDistance2 != nil {
		return true
	}

	return false
}

// SetOffsetDistance2 gets a reference to the given float64 and assigns it to the OffsetDistance2 field.
func (o *BTMSketchConstraint2) SetOffsetDistance2(v float64) {
	o.OffsetDistance2 = &v
}

// GetOffsetOrientation1 returns the OffsetOrientation1 field value if set, zero value otherwise.
func (o *BTMSketchConstraint2) GetOffsetOrientation1() bool {
	if o == nil || o.OffsetOrientation1 == nil {
		var ret bool
		return ret
	}
	return *o.OffsetOrientation1
}

// GetOffsetOrientation1Ok returns a tuple with the OffsetOrientation1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchConstraint2) GetOffsetOrientation1Ok() (*bool, bool) {
	if o == nil || o.OffsetOrientation1 == nil {
		return nil, false
	}
	return o.OffsetOrientation1, true
}

// HasOffsetOrientation1 returns a boolean if a field has been set.
func (o *BTMSketchConstraint2) HasOffsetOrientation1() bool {
	if o != nil && o.OffsetOrientation1 != nil {
		return true
	}

	return false
}

// SetOffsetOrientation1 gets a reference to the given bool and assigns it to the OffsetOrientation1 field.
func (o *BTMSketchConstraint2) SetOffsetOrientation1(v bool) {
	o.OffsetOrientation1 = &v
}

// GetOffsetOrientation2 returns the OffsetOrientation2 field value if set, zero value otherwise.
func (o *BTMSketchConstraint2) GetOffsetOrientation2() bool {
	if o == nil || o.OffsetOrientation2 == nil {
		var ret bool
		return ret
	}
	return *o.OffsetOrientation2
}

// GetOffsetOrientation2Ok returns a tuple with the OffsetOrientation2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchConstraint2) GetOffsetOrientation2Ok() (*bool, bool) {
	if o == nil || o.OffsetOrientation2 == nil {
		return nil, false
	}
	return o.OffsetOrientation2, true
}

// HasOffsetOrientation2 returns a boolean if a field has been set.
func (o *BTMSketchConstraint2) HasOffsetOrientation2() bool {
	if o != nil && o.OffsetOrientation2 != nil {
		return true
	}

	return false
}

// SetOffsetOrientation2 gets a reference to the given bool and assigns it to the OffsetOrientation2 field.
func (o *BTMSketchConstraint2) SetOffsetOrientation2(v bool) {
	o.OffsetOrientation2 = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *BTMSketchConstraint2) GetParameters() []BTMParameter1 {
	if o == nil || o.Parameters == nil {
		var ret []BTMParameter1
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchConstraint2) GetParametersOk() ([]BTMParameter1, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *BTMSketchConstraint2) HasParameters() bool {
	if o != nil && o.Parameters != nil {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []BTMParameter1 and assigns it to the Parameters field.
func (o *BTMSketchConstraint2) SetParameters(v []BTMParameter1) {
	o.Parameters = v
}

// GetPierceParameter returns the PierceParameter field value if set, zero value otherwise.
func (o *BTMSketchConstraint2) GetPierceParameter() float64 {
	if o == nil || o.PierceParameter == nil {
		var ret float64
		return ret
	}
	return *o.PierceParameter
}

// GetPierceParameterOk returns a tuple with the PierceParameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchConstraint2) GetPierceParameterOk() (*float64, bool) {
	if o == nil || o.PierceParameter == nil {
		return nil, false
	}
	return o.PierceParameter, true
}

// HasPierceParameter returns a boolean if a field has been set.
func (o *BTMSketchConstraint2) HasPierceParameter() bool {
	if o != nil && o.PierceParameter != nil {
		return true
	}

	return false
}

// SetPierceParameter gets a reference to the given float64 and assigns it to the PierceParameter field.
func (o *BTMSketchConstraint2) SetPierceParameter(v float64) {
	o.PierceParameter = &v
}

func (o BTMSketchConstraint2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ConstraintType != nil {
		toSerialize["constraintType"] = o.ConstraintType
	}
	if o.DrivenDimension != nil {
		toSerialize["drivenDimension"] = o.DrivenDimension
	}
	if o.EntityId != nil {
		toSerialize["entityId"] = o.EntityId
	}
	if o.EntityIdAndReplaceInDependentFields != nil {
		toSerialize["entityIdAndReplaceInDependentFields"] = o.EntityIdAndReplaceInDependentFields
	}
	if o.HasOffsetData1 != nil {
		toSerialize["hasOffsetData1"] = o.HasOffsetData1
	}
	if o.HasOffsetData2 != nil {
		toSerialize["hasOffsetData2"] = o.HasOffsetData2
	}
	if o.HasPierceParameter_ != nil {
		toSerialize["hasPierceParameter"] = o.HasPierceParameter_
	}
	if o.HelpParameters != nil {
		toSerialize["helpParameters"] = o.HelpParameters
	}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.OffsetDistance1 != nil {
		toSerialize["offsetDistance1"] = o.OffsetDistance1
	}
	if o.OffsetDistance2 != nil {
		toSerialize["offsetDistance2"] = o.OffsetDistance2
	}
	if o.OffsetOrientation1 != nil {
		toSerialize["offsetOrientation1"] = o.OffsetOrientation1
	}
	if o.OffsetOrientation2 != nil {
		toSerialize["offsetOrientation2"] = o.OffsetOrientation2
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	if o.PierceParameter != nil {
		toSerialize["pierceParameter"] = o.PierceParameter
	}
	return json.Marshal(toSerialize)
}

type NullableBTMSketchConstraint2 struct {
	value *BTMSketchConstraint2
	isSet bool
}

func (v NullableBTMSketchConstraint2) Get() *BTMSketchConstraint2 {
	return v.value
}

func (v *NullableBTMSketchConstraint2) Set(val *BTMSketchConstraint2) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMSketchConstraint2) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMSketchConstraint2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMSketchConstraint2(val *BTMSketchConstraint2) *NullableBTMSketchConstraint2 {
	return &NullableBTMSketchConstraint2{value: val, isSet: true}
}

func (v NullableBTMSketchConstraint2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMSketchConstraint2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
