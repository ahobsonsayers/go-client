/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.155.6922-11b87e04d539
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMSketchImageEntity763 struct for BTMSketchImageEntity763
type BTMSketchImageEntity763 struct {
	BtType                              *string         `json:"btType,omitempty"`
	ControlBoxIds                       []string        `json:"controlBoxIds,omitempty"`
	EntityId                            *string         `json:"entityId,omitempty"`
	EntityIdAndReplaceInDependentFields *string         `json:"entityIdAndReplaceInDependentFields,omitempty"`
	ImportMicroversion                  *string         `json:"importMicroversion,omitempty"`
	IsConstruction                      *bool           `json:"isConstruction,omitempty"`
	IsFromEndpointSplineHandle          *bool           `json:"isFromEndpointSplineHandle,omitempty"`
	IsFromSplineControlPolygon          *bool           `json:"isFromSplineControlPolygon,omitempty"`
	IsFromSplineHandle                  *bool           `json:"isFromSplineHandle,omitempty"`
	Namespace                           *string         `json:"namespace,omitempty"`
	NodeId                              *string         `json:"nodeId,omitempty"`
	Parameters                          []BTMParameter1 `json:"parameters,omitempty"`
	AspectRatio                         *float64        `json:"aspectRatio,omitempty"`
	OriginX                             *float64        `json:"originX,omitempty"`
	OriginY                             *float64        `json:"originY,omitempty"`
	XaxisX                              *float64        `json:"xaxisX,omitempty"`
	XaxisY                              *float64        `json:"xaxisY,omitempty"`
}

// NewBTMSketchImageEntity763 instantiates a new BTMSketchImageEntity763 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMSketchImageEntity763() *BTMSketchImageEntity763 {
	this := BTMSketchImageEntity763{}
	return &this
}

// NewBTMSketchImageEntity763WithDefaults instantiates a new BTMSketchImageEntity763 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMSketchImageEntity763WithDefaults() *BTMSketchImageEntity763 {
	this := BTMSketchImageEntity763{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMSketchImageEntity763) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchImageEntity763) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMSketchImageEntity763) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMSketchImageEntity763) SetBtType(v string) {
	o.BtType = &v
}

// GetControlBoxIds returns the ControlBoxIds field value if set, zero value otherwise.
func (o *BTMSketchImageEntity763) GetControlBoxIds() []string {
	if o == nil || o.ControlBoxIds == nil {
		var ret []string
		return ret
	}
	return o.ControlBoxIds
}

// GetControlBoxIdsOk returns a tuple with the ControlBoxIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchImageEntity763) GetControlBoxIdsOk() ([]string, bool) {
	if o == nil || o.ControlBoxIds == nil {
		return nil, false
	}
	return o.ControlBoxIds, true
}

// HasControlBoxIds returns a boolean if a field has been set.
func (o *BTMSketchImageEntity763) HasControlBoxIds() bool {
	if o != nil && o.ControlBoxIds != nil {
		return true
	}

	return false
}

// SetControlBoxIds gets a reference to the given []string and assigns it to the ControlBoxIds field.
func (o *BTMSketchImageEntity763) SetControlBoxIds(v []string) {
	o.ControlBoxIds = v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *BTMSketchImageEntity763) GetEntityId() string {
	if o == nil || o.EntityId == nil {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchImageEntity763) GetEntityIdOk() (*string, bool) {
	if o == nil || o.EntityId == nil {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *BTMSketchImageEntity763) HasEntityId() bool {
	if o != nil && o.EntityId != nil {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *BTMSketchImageEntity763) SetEntityId(v string) {
	o.EntityId = &v
}

// GetEntityIdAndReplaceInDependentFields returns the EntityIdAndReplaceInDependentFields field value if set, zero value otherwise.
func (o *BTMSketchImageEntity763) GetEntityIdAndReplaceInDependentFields() string {
	if o == nil || o.EntityIdAndReplaceInDependentFields == nil {
		var ret string
		return ret
	}
	return *o.EntityIdAndReplaceInDependentFields
}

// GetEntityIdAndReplaceInDependentFieldsOk returns a tuple with the EntityIdAndReplaceInDependentFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchImageEntity763) GetEntityIdAndReplaceInDependentFieldsOk() (*string, bool) {
	if o == nil || o.EntityIdAndReplaceInDependentFields == nil {
		return nil, false
	}
	return o.EntityIdAndReplaceInDependentFields, true
}

// HasEntityIdAndReplaceInDependentFields returns a boolean if a field has been set.
func (o *BTMSketchImageEntity763) HasEntityIdAndReplaceInDependentFields() bool {
	if o != nil && o.EntityIdAndReplaceInDependentFields != nil {
		return true
	}

	return false
}

// SetEntityIdAndReplaceInDependentFields gets a reference to the given string and assigns it to the EntityIdAndReplaceInDependentFields field.
func (o *BTMSketchImageEntity763) SetEntityIdAndReplaceInDependentFields(v string) {
	o.EntityIdAndReplaceInDependentFields = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMSketchImageEntity763) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchImageEntity763) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMSketchImageEntity763) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMSketchImageEntity763) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetIsConstruction returns the IsConstruction field value if set, zero value otherwise.
func (o *BTMSketchImageEntity763) GetIsConstruction() bool {
	if o == nil || o.IsConstruction == nil {
		var ret bool
		return ret
	}
	return *o.IsConstruction
}

// GetIsConstructionOk returns a tuple with the IsConstruction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchImageEntity763) GetIsConstructionOk() (*bool, bool) {
	if o == nil || o.IsConstruction == nil {
		return nil, false
	}
	return o.IsConstruction, true
}

// HasIsConstruction returns a boolean if a field has been set.
func (o *BTMSketchImageEntity763) HasIsConstruction() bool {
	if o != nil && o.IsConstruction != nil {
		return true
	}

	return false
}

// SetIsConstruction gets a reference to the given bool and assigns it to the IsConstruction field.
func (o *BTMSketchImageEntity763) SetIsConstruction(v bool) {
	o.IsConstruction = &v
}

// GetIsFromEndpointSplineHandle returns the IsFromEndpointSplineHandle field value if set, zero value otherwise.
func (o *BTMSketchImageEntity763) GetIsFromEndpointSplineHandle() bool {
	if o == nil || o.IsFromEndpointSplineHandle == nil {
		var ret bool
		return ret
	}
	return *o.IsFromEndpointSplineHandle
}

// GetIsFromEndpointSplineHandleOk returns a tuple with the IsFromEndpointSplineHandle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchImageEntity763) GetIsFromEndpointSplineHandleOk() (*bool, bool) {
	if o == nil || o.IsFromEndpointSplineHandle == nil {
		return nil, false
	}
	return o.IsFromEndpointSplineHandle, true
}

// HasIsFromEndpointSplineHandle returns a boolean if a field has been set.
func (o *BTMSketchImageEntity763) HasIsFromEndpointSplineHandle() bool {
	if o != nil && o.IsFromEndpointSplineHandle != nil {
		return true
	}

	return false
}

// SetIsFromEndpointSplineHandle gets a reference to the given bool and assigns it to the IsFromEndpointSplineHandle field.
func (o *BTMSketchImageEntity763) SetIsFromEndpointSplineHandle(v bool) {
	o.IsFromEndpointSplineHandle = &v
}

// GetIsFromSplineControlPolygon returns the IsFromSplineControlPolygon field value if set, zero value otherwise.
func (o *BTMSketchImageEntity763) GetIsFromSplineControlPolygon() bool {
	if o == nil || o.IsFromSplineControlPolygon == nil {
		var ret bool
		return ret
	}
	return *o.IsFromSplineControlPolygon
}

// GetIsFromSplineControlPolygonOk returns a tuple with the IsFromSplineControlPolygon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchImageEntity763) GetIsFromSplineControlPolygonOk() (*bool, bool) {
	if o == nil || o.IsFromSplineControlPolygon == nil {
		return nil, false
	}
	return o.IsFromSplineControlPolygon, true
}

// HasIsFromSplineControlPolygon returns a boolean if a field has been set.
func (o *BTMSketchImageEntity763) HasIsFromSplineControlPolygon() bool {
	if o != nil && o.IsFromSplineControlPolygon != nil {
		return true
	}

	return false
}

// SetIsFromSplineControlPolygon gets a reference to the given bool and assigns it to the IsFromSplineControlPolygon field.
func (o *BTMSketchImageEntity763) SetIsFromSplineControlPolygon(v bool) {
	o.IsFromSplineControlPolygon = &v
}

// GetIsFromSplineHandle returns the IsFromSplineHandle field value if set, zero value otherwise.
func (o *BTMSketchImageEntity763) GetIsFromSplineHandle() bool {
	if o == nil || o.IsFromSplineHandle == nil {
		var ret bool
		return ret
	}
	return *o.IsFromSplineHandle
}

// GetIsFromSplineHandleOk returns a tuple with the IsFromSplineHandle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchImageEntity763) GetIsFromSplineHandleOk() (*bool, bool) {
	if o == nil || o.IsFromSplineHandle == nil {
		return nil, false
	}
	return o.IsFromSplineHandle, true
}

// HasIsFromSplineHandle returns a boolean if a field has been set.
func (o *BTMSketchImageEntity763) HasIsFromSplineHandle() bool {
	if o != nil && o.IsFromSplineHandle != nil {
		return true
	}

	return false
}

// SetIsFromSplineHandle gets a reference to the given bool and assigns it to the IsFromSplineHandle field.
func (o *BTMSketchImageEntity763) SetIsFromSplineHandle(v bool) {
	o.IsFromSplineHandle = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *BTMSketchImageEntity763) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchImageEntity763) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *BTMSketchImageEntity763) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *BTMSketchImageEntity763) SetNamespace(v string) {
	o.Namespace = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMSketchImageEntity763) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchImageEntity763) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMSketchImageEntity763) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMSketchImageEntity763) SetNodeId(v string) {
	o.NodeId = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *BTMSketchImageEntity763) GetParameters() []BTMParameter1 {
	if o == nil || o.Parameters == nil {
		var ret []BTMParameter1
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchImageEntity763) GetParametersOk() ([]BTMParameter1, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *BTMSketchImageEntity763) HasParameters() bool {
	if o != nil && o.Parameters != nil {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []BTMParameter1 and assigns it to the Parameters field.
func (o *BTMSketchImageEntity763) SetParameters(v []BTMParameter1) {
	o.Parameters = v
}

// GetAspectRatio returns the AspectRatio field value if set, zero value otherwise.
func (o *BTMSketchImageEntity763) GetAspectRatio() float64 {
	if o == nil || o.AspectRatio == nil {
		var ret float64
		return ret
	}
	return *o.AspectRatio
}

// GetAspectRatioOk returns a tuple with the AspectRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchImageEntity763) GetAspectRatioOk() (*float64, bool) {
	if o == nil || o.AspectRatio == nil {
		return nil, false
	}
	return o.AspectRatio, true
}

// HasAspectRatio returns a boolean if a field has been set.
func (o *BTMSketchImageEntity763) HasAspectRatio() bool {
	if o != nil && o.AspectRatio != nil {
		return true
	}

	return false
}

// SetAspectRatio gets a reference to the given float64 and assigns it to the AspectRatio field.
func (o *BTMSketchImageEntity763) SetAspectRatio(v float64) {
	o.AspectRatio = &v
}

// GetOriginX returns the OriginX field value if set, zero value otherwise.
func (o *BTMSketchImageEntity763) GetOriginX() float64 {
	if o == nil || o.OriginX == nil {
		var ret float64
		return ret
	}
	return *o.OriginX
}

// GetOriginXOk returns a tuple with the OriginX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchImageEntity763) GetOriginXOk() (*float64, bool) {
	if o == nil || o.OriginX == nil {
		return nil, false
	}
	return o.OriginX, true
}

// HasOriginX returns a boolean if a field has been set.
func (o *BTMSketchImageEntity763) HasOriginX() bool {
	if o != nil && o.OriginX != nil {
		return true
	}

	return false
}

// SetOriginX gets a reference to the given float64 and assigns it to the OriginX field.
func (o *BTMSketchImageEntity763) SetOriginX(v float64) {
	o.OriginX = &v
}

// GetOriginY returns the OriginY field value if set, zero value otherwise.
func (o *BTMSketchImageEntity763) GetOriginY() float64 {
	if o == nil || o.OriginY == nil {
		var ret float64
		return ret
	}
	return *o.OriginY
}

// GetOriginYOk returns a tuple with the OriginY field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchImageEntity763) GetOriginYOk() (*float64, bool) {
	if o == nil || o.OriginY == nil {
		return nil, false
	}
	return o.OriginY, true
}

// HasOriginY returns a boolean if a field has been set.
func (o *BTMSketchImageEntity763) HasOriginY() bool {
	if o != nil && o.OriginY != nil {
		return true
	}

	return false
}

// SetOriginY gets a reference to the given float64 and assigns it to the OriginY field.
func (o *BTMSketchImageEntity763) SetOriginY(v float64) {
	o.OriginY = &v
}

// GetXaxisX returns the XaxisX field value if set, zero value otherwise.
func (o *BTMSketchImageEntity763) GetXaxisX() float64 {
	if o == nil || o.XaxisX == nil {
		var ret float64
		return ret
	}
	return *o.XaxisX
}

// GetXaxisXOk returns a tuple with the XaxisX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchImageEntity763) GetXaxisXOk() (*float64, bool) {
	if o == nil || o.XaxisX == nil {
		return nil, false
	}
	return o.XaxisX, true
}

// HasXaxisX returns a boolean if a field has been set.
func (o *BTMSketchImageEntity763) HasXaxisX() bool {
	if o != nil && o.XaxisX != nil {
		return true
	}

	return false
}

// SetXaxisX gets a reference to the given float64 and assigns it to the XaxisX field.
func (o *BTMSketchImageEntity763) SetXaxisX(v float64) {
	o.XaxisX = &v
}

// GetXaxisY returns the XaxisY field value if set, zero value otherwise.
func (o *BTMSketchImageEntity763) GetXaxisY() float64 {
	if o == nil || o.XaxisY == nil {
		var ret float64
		return ret
	}
	return *o.XaxisY
}

// GetXaxisYOk returns a tuple with the XaxisY field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchImageEntity763) GetXaxisYOk() (*float64, bool) {
	if o == nil || o.XaxisY == nil {
		return nil, false
	}
	return o.XaxisY, true
}

// HasXaxisY returns a boolean if a field has been set.
func (o *BTMSketchImageEntity763) HasXaxisY() bool {
	if o != nil && o.XaxisY != nil {
		return true
	}

	return false
}

// SetXaxisY gets a reference to the given float64 and assigns it to the XaxisY field.
func (o *BTMSketchImageEntity763) SetXaxisY(v float64) {
	o.XaxisY = &v
}

func (o BTMSketchImageEntity763) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ControlBoxIds != nil {
		toSerialize["controlBoxIds"] = o.ControlBoxIds
	}
	if o.EntityId != nil {
		toSerialize["entityId"] = o.EntityId
	}
	if o.EntityIdAndReplaceInDependentFields != nil {
		toSerialize["entityIdAndReplaceInDependentFields"] = o.EntityIdAndReplaceInDependentFields
	}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.IsConstruction != nil {
		toSerialize["isConstruction"] = o.IsConstruction
	}
	if o.IsFromEndpointSplineHandle != nil {
		toSerialize["isFromEndpointSplineHandle"] = o.IsFromEndpointSplineHandle
	}
	if o.IsFromSplineControlPolygon != nil {
		toSerialize["isFromSplineControlPolygon"] = o.IsFromSplineControlPolygon
	}
	if o.IsFromSplineHandle != nil {
		toSerialize["isFromSplineHandle"] = o.IsFromSplineHandle
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	if o.AspectRatio != nil {
		toSerialize["aspectRatio"] = o.AspectRatio
	}
	if o.OriginX != nil {
		toSerialize["originX"] = o.OriginX
	}
	if o.OriginY != nil {
		toSerialize["originY"] = o.OriginY
	}
	if o.XaxisX != nil {
		toSerialize["xaxisX"] = o.XaxisX
	}
	if o.XaxisY != nil {
		toSerialize["xaxisY"] = o.XaxisY
	}
	return json.Marshal(toSerialize)
}

type NullableBTMSketchImageEntity763 struct {
	value *BTMSketchImageEntity763
	isSet bool
}

func (v NullableBTMSketchImageEntity763) Get() *BTMSketchImageEntity763 {
	return v.value
}

func (v *NullableBTMSketchImageEntity763) Set(val *BTMSketchImageEntity763) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMSketchImageEntity763) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMSketchImageEntity763) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMSketchImageEntity763(val *BTMSketchImageEntity763) *NullableBTMSketchImageEntity763 {
	return &NullableBTMSketchImageEntity763{value: val, isSet: true}
}

func (v NullableBTMSketchImageEntity763) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMSketchImageEntity763) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
