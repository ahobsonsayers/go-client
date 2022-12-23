/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.157.9170-8cb36f16959d
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTTable1825 - struct for BTTable1825
type BTTable1825 struct {
	implBTTable1825 interface{}
}

// BTAssemblySimulationStructuralLoadsTable3867AsBTTable1825 is a convenience function that returns BTAssemblySimulationStructuralLoadsTable3867 wrapped in BTTable1825
func (o *BTAssemblySimulationStructuralLoadsTable3867) AsBTTable1825() *BTTable1825 {
	return &BTTable1825{o}
}

// BTAssemblySimulationTable3236AsBTTable1825 is a convenience function that returns BTAssemblySimulationTable3236 wrapped in BTTable1825
func (o *BTAssemblySimulationTable3236) AsBTTable1825() *BTTable1825 {
	return &BTTable1825{o}
}

// BTConfiguredPartPropertiesTable2740AsBTTable1825 is a convenience function that returns BTConfiguredPartPropertiesTable2740 wrapped in BTTable1825
func (o *BTConfiguredPartPropertiesTable2740) AsBTTable1825() *BTTable1825 {
	return &BTTable1825{o}
}

// BTFSTable953AsBTTable1825 is a convenience function that returns BTFSTable953 wrapped in BTTable1825
func (o *BTFSTable953) AsBTTable1825() *BTTable1825 {
	return &BTTable1825{o}
}

// BTBillOfMaterialsTable1073AsBTTable1825 is a convenience function that returns BTBillOfMaterialsTable1073 wrapped in BTTable1825
func (o *BTBillOfMaterialsTable1073) AsBTTable1825() *BTTable1825 {
	return &BTTable1825{o}
}

// NewBTTable1825 instantiates a new BTTable1825 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTTable1825() *BTTable1825 {
	this := BTTable1825{Newbase_BTTable1825()}
	return &this
}

// NewBTTable1825WithDefaults instantiates a new BTTable1825 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTTable1825WithDefaults() *BTTable1825 {
	this := BTTable1825{Newbase_BTTable1825WithDefaults()}
	return &this
}

// GetAllRowValues returns the AllRowValues field value if set, zero value otherwise.
func (o *BTTable1825) GetAllRowValues() [][]string {
	type getResult interface {
		GetAllRowValues() [][]string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetAllRowValues()
	} else {
		var de [][]string
		return de
	}
}

// GetAllRowValuesOk returns a tuple with the AllRowValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTable1825) GetAllRowValuesOk() ([][]string, bool) {
	type getResult interface {
		GetAllRowValuesOk() ([][]string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetAllRowValuesOk()
	} else {
		return nil, false
	}
}

// HasAllRowValues returns a boolean if a field has been set.
func (o *BTTable1825) HasAllRowValues() bool {
	type getResult interface {
		HasAllRowValues() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasAllRowValues()
	} else {
		return false
	}
}

// SetAllRowValues gets a reference to the given [][]string and assigns it to the AllRowValues field.
func (o *BTTable1825) SetAllRowValues(v [][]string) {
	type getResult interface {
		SetAllRowValues(v [][]string)
	}

	o.GetActualInstance().(getResult).SetAllRowValues(v)
}

// GetColumnCount returns the ColumnCount field value if set, zero value otherwise.
func (o *BTTable1825) GetColumnCount() int32 {
	type getResult interface {
		GetColumnCount() int32
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetColumnCount()
	} else {
		var de int32
		return de
	}
}

// GetColumnCountOk returns a tuple with the ColumnCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTable1825) GetColumnCountOk() (*int32, bool) {
	type getResult interface {
		GetColumnCountOk() (*int32, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetColumnCountOk()
	} else {
		return nil, false
	}
}

// HasColumnCount returns a boolean if a field has been set.
func (o *BTTable1825) HasColumnCount() bool {
	type getResult interface {
		HasColumnCount() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasColumnCount()
	} else {
		return false
	}
}

// SetColumnCount gets a reference to the given int32 and assigns it to the ColumnCount field.
func (o *BTTable1825) SetColumnCount(v int32) {
	type getResult interface {
		SetColumnCount(v int32)
	}

	o.GetActualInstance().(getResult).SetColumnCount(v)
}

// GetFrozenColumns returns the FrozenColumns field value if set, zero value otherwise.
func (o *BTTable1825) GetFrozenColumns() int32 {
	type getResult interface {
		GetFrozenColumns() int32
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetFrozenColumns()
	} else {
		var de int32
		return de
	}
}

// GetFrozenColumnsOk returns a tuple with the FrozenColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTable1825) GetFrozenColumnsOk() (*int32, bool) {
	type getResult interface {
		GetFrozenColumnsOk() (*int32, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetFrozenColumnsOk()
	} else {
		return nil, false
	}
}

// HasFrozenColumns returns a boolean if a field has been set.
func (o *BTTable1825) HasFrozenColumns() bool {
	type getResult interface {
		HasFrozenColumns() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasFrozenColumns()
	} else {
		return false
	}
}

// SetFrozenColumns gets a reference to the given int32 and assigns it to the FrozenColumns field.
func (o *BTTable1825) SetFrozenColumns(v int32) {
	type getResult interface {
		SetFrozenColumns(v int32)
	}

	o.GetActualInstance().(getResult).SetFrozenColumns(v)
}

// GetIsFailed returns the IsFailed field value if set, zero value otherwise.
func (o *BTTable1825) GetIsFailed() bool {
	type getResult interface {
		GetIsFailed() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetIsFailed()
	} else {
		var de bool
		return de
	}
}

// GetIsFailedOk returns a tuple with the IsFailed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTable1825) GetIsFailedOk() (*bool, bool) {
	type getResult interface {
		GetIsFailedOk() (*bool, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetIsFailedOk()
	} else {
		return nil, false
	}
}

// HasIsFailed returns a boolean if a field has been set.
func (o *BTTable1825) HasIsFailed() bool {
	type getResult interface {
		HasIsFailed() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasIsFailed()
	} else {
		return false
	}
}

// SetIsFailed gets a reference to the given bool and assigns it to the IsFailed field.
func (o *BTTable1825) SetIsFailed(v bool) {
	type getResult interface {
		SetIsFailed(v bool)
	}

	o.GetActualInstance().(getResult).SetIsFailed(v)
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTTable1825) GetNodeId() string {
	type getResult interface {
		GetNodeId() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetNodeId()
	} else {
		var de string
		return de
	}
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTable1825) GetNodeIdOk() (*string, bool) {
	type getResult interface {
		GetNodeIdOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetNodeIdOk()
	} else {
		return nil, false
	}
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTTable1825) HasNodeId() bool {
	type getResult interface {
		HasNodeId() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasNodeId()
	} else {
		return false
	}
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTTable1825) SetNodeId(v string) {
	type getResult interface {
		SetNodeId(v string)
	}

	o.GetActualInstance().(getResult).SetNodeId(v)
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *BTTable1825) GetReadOnly() bool {
	type getResult interface {
		GetReadOnly() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetReadOnly()
	} else {
		var de bool
		return de
	}
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTable1825) GetReadOnlyOk() (*bool, bool) {
	type getResult interface {
		GetReadOnlyOk() (*bool, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetReadOnlyOk()
	} else {
		return nil, false
	}
}

// HasReadOnly returns a boolean if a field has been set.
func (o *BTTable1825) HasReadOnly() bool {
	type getResult interface {
		HasReadOnly() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasReadOnly()
	} else {
		return false
	}
}

// SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.
func (o *BTTable1825) SetReadOnly(v bool) {
	type getResult interface {
		SetReadOnly(v bool)
	}

	o.GetActualInstance().(getResult).SetReadOnly(v)
}

// GetRowCount returns the RowCount field value if set, zero value otherwise.
func (o *BTTable1825) GetRowCount() int32 {
	type getResult interface {
		GetRowCount() int32
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetRowCount()
	} else {
		var de int32
		return de
	}
}

// GetRowCountOk returns a tuple with the RowCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTable1825) GetRowCountOk() (*int32, bool) {
	type getResult interface {
		GetRowCountOk() (*int32, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetRowCountOk()
	} else {
		return nil, false
	}
}

// HasRowCount returns a boolean if a field has been set.
func (o *BTTable1825) HasRowCount() bool {
	type getResult interface {
		HasRowCount() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasRowCount()
	} else {
		return false
	}
}

// SetRowCount gets a reference to the given int32 and assigns it to the RowCount field.
func (o *BTTable1825) SetRowCount(v int32) {
	type getResult interface {
		SetRowCount(v int32)
	}

	o.GetActualInstance().(getResult).SetRowCount(v)
}

// GetSortOrder returns the SortOrder field value if set, zero value otherwise.
func (o *BTTable1825) GetSortOrder() BTTableSortOrder4371 {
	type getResult interface {
		GetSortOrder() BTTableSortOrder4371
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetSortOrder()
	} else {
		var de BTTableSortOrder4371
		return de
	}
}

// GetSortOrderOk returns a tuple with the SortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTable1825) GetSortOrderOk() (*BTTableSortOrder4371, bool) {
	type getResult interface {
		GetSortOrderOk() (*BTTableSortOrder4371, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetSortOrderOk()
	} else {
		return nil, false
	}
}

// HasSortOrder returns a boolean if a field has been set.
func (o *BTTable1825) HasSortOrder() bool {
	type getResult interface {
		HasSortOrder() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasSortOrder()
	} else {
		return false
	}
}

// SetSortOrder gets a reference to the given BTTableSortOrder4371 and assigns it to the SortOrder field.
func (o *BTTable1825) SetSortOrder(v BTTableSortOrder4371) {
	type getResult interface {
		SetSortOrder(v BTTableSortOrder4371)
	}

	o.GetActualInstance().(getResult).SetSortOrder(v)
}

// GetTableColumns returns the TableColumns field value if set, zero value otherwise.
func (o *BTTable1825) GetTableColumns() []BTTableColumnInfo1222 {
	type getResult interface {
		GetTableColumns() []BTTableColumnInfo1222
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetTableColumns()
	} else {
		var de []BTTableColumnInfo1222
		return de
	}
}

// GetTableColumnsOk returns a tuple with the TableColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTable1825) GetTableColumnsOk() ([]BTTableColumnInfo1222, bool) {
	type getResult interface {
		GetTableColumnsOk() ([]BTTableColumnInfo1222, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetTableColumnsOk()
	} else {
		return nil, false
	}
}

// HasTableColumns returns a boolean if a field has been set.
func (o *BTTable1825) HasTableColumns() bool {
	type getResult interface {
		HasTableColumns() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasTableColumns()
	} else {
		return false
	}
}

// SetTableColumns gets a reference to the given []BTTableColumnInfo1222 and assigns it to the TableColumns field.
func (o *BTTable1825) SetTableColumns(v []BTTableColumnInfo1222) {
	type getResult interface {
		SetTableColumns(v []BTTableColumnInfo1222)
	}

	o.GetActualInstance().(getResult).SetTableColumns(v)
}

// GetTableId returns the TableId field value if set, zero value otherwise.
func (o *BTTable1825) GetTableId() string {
	type getResult interface {
		GetTableId() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetTableId()
	} else {
		var de string
		return de
	}
}

// GetTableIdOk returns a tuple with the TableId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTable1825) GetTableIdOk() (*string, bool) {
	type getResult interface {
		GetTableIdOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetTableIdOk()
	} else {
		return nil, false
	}
}

// HasTableId returns a boolean if a field has been set.
func (o *BTTable1825) HasTableId() bool {
	type getResult interface {
		HasTableId() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasTableId()
	} else {
		return false
	}
}

// SetTableId gets a reference to the given string and assigns it to the TableId field.
func (o *BTTable1825) SetTableId(v string) {
	type getResult interface {
		SetTableId(v string)
	}

	o.GetActualInstance().(getResult).SetTableId(v)
}

// GetTableRows returns the TableRows field value if set, zero value otherwise.
func (o *BTTable1825) GetTableRows() []BTTableRow1054 {
	type getResult interface {
		GetTableRows() []BTTableRow1054
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetTableRows()
	} else {
		var de []BTTableRow1054
		return de
	}
}

// GetTableRowsOk returns a tuple with the TableRows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTable1825) GetTableRowsOk() ([]BTTableRow1054, bool) {
	type getResult interface {
		GetTableRowsOk() ([]BTTableRow1054, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetTableRowsOk()
	} else {
		return nil, false
	}
}

// HasTableRows returns a boolean if a field has been set.
func (o *BTTable1825) HasTableRows() bool {
	type getResult interface {
		HasTableRows() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasTableRows()
	} else {
		return false
	}
}

// SetTableRows gets a reference to the given []BTTableRow1054 and assigns it to the TableRows field.
func (o *BTTable1825) SetTableRows(v []BTTableRow1054) {
	type getResult interface {
		SetTableRows(v []BTTableRow1054)
	}

	o.GetActualInstance().(getResult).SetTableRows(v)
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *BTTable1825) GetTitle() string {
	type getResult interface {
		GetTitle() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetTitle()
	} else {
		var de string
		return de
	}
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTable1825) GetTitleOk() (*string, bool) {
	type getResult interface {
		GetTitleOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetTitleOk()
	} else {
		return nil, false
	}
}

// HasTitle returns a boolean if a field has been set.
func (o *BTTable1825) HasTitle() bool {
	type getResult interface {
		HasTitle() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasTitle()
	} else {
		return false
	}
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *BTTable1825) SetTitle(v string) {
	type getResult interface {
		SetTitle(v string)
	}

	o.GetActualInstance().(getResult).SetTitle(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTTable1825) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTAssemblySimulationStructuralLoadsTable-3867'
	if jsonDict["btType"] == "BTAssemblySimulationStructuralLoadsTable-3867" {
		// try to unmarshal JSON data into BTAssemblySimulationStructuralLoadsTable3867
		var qr *BTAssemblySimulationStructuralLoadsTable3867
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTTable1825 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTTable1825 = nil
			return fmt.Errorf("Failed to unmarshal BTTable1825 as BTAssemblySimulationStructuralLoadsTable3867: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTAssemblySimulationTable-3236'
	if jsonDict["btType"] == "BTAssemblySimulationTable-3236" {
		// try to unmarshal JSON data into BTAssemblySimulationTable3236
		var qr *BTAssemblySimulationTable3236
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTTable1825 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTTable1825 = nil
			return fmt.Errorf("Failed to unmarshal BTTable1825 as BTAssemblySimulationTable3236: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTBillOfMaterialsTable-1073'
	if jsonDict["btType"] == "BTBillOfMaterialsTable-1073" {
		// try to unmarshal JSON data into BTBillOfMaterialsTable1073
		var qr *BTBillOfMaterialsTable1073
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTTable1825 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTTable1825 = nil
			return fmt.Errorf("Failed to unmarshal BTTable1825 as BTBillOfMaterialsTable1073: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTConfiguredPartPropertiesTable-2740'
	if jsonDict["btType"] == "BTConfiguredPartPropertiesTable-2740" {
		// try to unmarshal JSON data into BTConfiguredPartPropertiesTable2740
		var qr *BTConfiguredPartPropertiesTable2740
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTTable1825 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTTable1825 = nil
			return fmt.Errorf("Failed to unmarshal BTTable1825 as BTConfiguredPartPropertiesTable2740: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTFSTable-953'
	if jsonDict["btType"] == "BTFSTable-953" {
		// try to unmarshal JSON data into BTFSTable953
		var qr *BTFSTable953
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTTable1825 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTTable1825 = nil
			return fmt.Errorf("Failed to unmarshal BTTable1825 as BTFSTable953: %s", err.Error())
		}
	}

	var qtx *base_BTTable1825
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTTable1825 = qtx
		return nil // data stored in dst.base_BTTable1825, return on the first match
	} else {
		dst.implBTTable1825 = nil
		return fmt.Errorf("Failed to unmarshal BTTable1825 as base_BTTable1825: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTTable1825) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTTable1825) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTTable1825
}

type NullableBTTable1825 struct {
	value *BTTable1825
	isSet bool
}

func (v NullableBTTable1825) Get() *BTTable1825 {
	return v.value
}

func (v *NullableBTTable1825) Set(val *BTTable1825) {
	v.value = val
	v.isSet = true
}

func (v NullableBTTable1825) IsSet() bool {
	return v.isSet
}

func (v *NullableBTTable1825) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTTable1825(val *BTTable1825) *NullableBTTable1825 {
	return &NullableBTTable1825{value: val, isSet: true}
}

func (v NullableBTTable1825) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTTable1825) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTTable1825 struct {
	AllRowValues  [][]string              `json:"allRowValues,omitempty"`
	ColumnCount   *int32                  `json:"columnCount,omitempty"`
	FrozenColumns *int32                  `json:"frozenColumns,omitempty"`
	IsFailed      *bool                   `json:"isFailed,omitempty"`
	NodeId        *string                 `json:"nodeId,omitempty"`
	ReadOnly      *bool                   `json:"readOnly,omitempty"`
	RowCount      *int32                  `json:"rowCount,omitempty"`
	SortOrder     *BTTableSortOrder4371   `json:"sortOrder,omitempty"`
	TableColumns  []BTTableColumnInfo1222 `json:"tableColumns,omitempty"`
	TableId       *string                 `json:"tableId,omitempty"`
	TableRows     []BTTableRow1054        `json:"tableRows,omitempty"`
	Title         *string                 `json:"title,omitempty"`
}

// Newbase_BTTable1825 instantiates a new base_BTTable1825 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTTable1825() *base_BTTable1825 {
	this := base_BTTable1825{}
	return &this
}

// Newbase_BTTable1825WithDefaults instantiates a new base_BTTable1825 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTTable1825WithDefaults() *base_BTTable1825 {
	this := base_BTTable1825{}
	return &this
}

// GetAllRowValues returns the AllRowValues field value if set, zero value otherwise.
func (o *base_BTTable1825) GetAllRowValues() [][]string {
	if o == nil || o.AllRowValues == nil {
		var ret [][]string
		return ret
	}
	return o.AllRowValues
}

// GetAllRowValuesOk returns a tuple with the AllRowValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTTable1825) GetAllRowValuesOk() ([][]string, bool) {
	if o == nil || o.AllRowValues == nil {
		return nil, false
	}
	return o.AllRowValues, true
}

// HasAllRowValues returns a boolean if a field has been set.
func (o *base_BTTable1825) HasAllRowValues() bool {
	if o != nil && o.AllRowValues != nil {
		return true
	}

	return false
}

// SetAllRowValues gets a reference to the given [][]string and assigns it to the AllRowValues field.
func (o *base_BTTable1825) SetAllRowValues(v [][]string) {
	o.AllRowValues = v
}

// GetColumnCount returns the ColumnCount field value if set, zero value otherwise.
func (o *base_BTTable1825) GetColumnCount() int32 {
	if o == nil || o.ColumnCount == nil {
		var ret int32
		return ret
	}
	return *o.ColumnCount
}

// GetColumnCountOk returns a tuple with the ColumnCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTTable1825) GetColumnCountOk() (*int32, bool) {
	if o == nil || o.ColumnCount == nil {
		return nil, false
	}
	return o.ColumnCount, true
}

// HasColumnCount returns a boolean if a field has been set.
func (o *base_BTTable1825) HasColumnCount() bool {
	if o != nil && o.ColumnCount != nil {
		return true
	}

	return false
}

// SetColumnCount gets a reference to the given int32 and assigns it to the ColumnCount field.
func (o *base_BTTable1825) SetColumnCount(v int32) {
	o.ColumnCount = &v
}

// GetFrozenColumns returns the FrozenColumns field value if set, zero value otherwise.
func (o *base_BTTable1825) GetFrozenColumns() int32 {
	if o == nil || o.FrozenColumns == nil {
		var ret int32
		return ret
	}
	return *o.FrozenColumns
}

// GetFrozenColumnsOk returns a tuple with the FrozenColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTTable1825) GetFrozenColumnsOk() (*int32, bool) {
	if o == nil || o.FrozenColumns == nil {
		return nil, false
	}
	return o.FrozenColumns, true
}

// HasFrozenColumns returns a boolean if a field has been set.
func (o *base_BTTable1825) HasFrozenColumns() bool {
	if o != nil && o.FrozenColumns != nil {
		return true
	}

	return false
}

// SetFrozenColumns gets a reference to the given int32 and assigns it to the FrozenColumns field.
func (o *base_BTTable1825) SetFrozenColumns(v int32) {
	o.FrozenColumns = &v
}

// GetIsFailed returns the IsFailed field value if set, zero value otherwise.
func (o *base_BTTable1825) GetIsFailed() bool {
	if o == nil || o.IsFailed == nil {
		var ret bool
		return ret
	}
	return *o.IsFailed
}

// GetIsFailedOk returns a tuple with the IsFailed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTTable1825) GetIsFailedOk() (*bool, bool) {
	if o == nil || o.IsFailed == nil {
		return nil, false
	}
	return o.IsFailed, true
}

// HasIsFailed returns a boolean if a field has been set.
func (o *base_BTTable1825) HasIsFailed() bool {
	if o != nil && o.IsFailed != nil {
		return true
	}

	return false
}

// SetIsFailed gets a reference to the given bool and assigns it to the IsFailed field.
func (o *base_BTTable1825) SetIsFailed(v bool) {
	o.IsFailed = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *base_BTTable1825) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTTable1825) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *base_BTTable1825) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *base_BTTable1825) SetNodeId(v string) {
	o.NodeId = &v
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *base_BTTable1825) GetReadOnly() bool {
	if o == nil || o.ReadOnly == nil {
		var ret bool
		return ret
	}
	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTTable1825) GetReadOnlyOk() (*bool, bool) {
	if o == nil || o.ReadOnly == nil {
		return nil, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *base_BTTable1825) HasReadOnly() bool {
	if o != nil && o.ReadOnly != nil {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.
func (o *base_BTTable1825) SetReadOnly(v bool) {
	o.ReadOnly = &v
}

// GetRowCount returns the RowCount field value if set, zero value otherwise.
func (o *base_BTTable1825) GetRowCount() int32 {
	if o == nil || o.RowCount == nil {
		var ret int32
		return ret
	}
	return *o.RowCount
}

// GetRowCountOk returns a tuple with the RowCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTTable1825) GetRowCountOk() (*int32, bool) {
	if o == nil || o.RowCount == nil {
		return nil, false
	}
	return o.RowCount, true
}

// HasRowCount returns a boolean if a field has been set.
func (o *base_BTTable1825) HasRowCount() bool {
	if o != nil && o.RowCount != nil {
		return true
	}

	return false
}

// SetRowCount gets a reference to the given int32 and assigns it to the RowCount field.
func (o *base_BTTable1825) SetRowCount(v int32) {
	o.RowCount = &v
}

// GetSortOrder returns the SortOrder field value if set, zero value otherwise.
func (o *base_BTTable1825) GetSortOrder() BTTableSortOrder4371 {
	if o == nil || o.SortOrder == nil {
		var ret BTTableSortOrder4371
		return ret
	}
	return *o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTTable1825) GetSortOrderOk() (*BTTableSortOrder4371, bool) {
	if o == nil || o.SortOrder == nil {
		return nil, false
	}
	return o.SortOrder, true
}

// HasSortOrder returns a boolean if a field has been set.
func (o *base_BTTable1825) HasSortOrder() bool {
	if o != nil && o.SortOrder != nil {
		return true
	}

	return false
}

// SetSortOrder gets a reference to the given BTTableSortOrder4371 and assigns it to the SortOrder field.
func (o *base_BTTable1825) SetSortOrder(v BTTableSortOrder4371) {
	o.SortOrder = &v
}

// GetTableColumns returns the TableColumns field value if set, zero value otherwise.
func (o *base_BTTable1825) GetTableColumns() []BTTableColumnInfo1222 {
	if o == nil || o.TableColumns == nil {
		var ret []BTTableColumnInfo1222
		return ret
	}
	return o.TableColumns
}

// GetTableColumnsOk returns a tuple with the TableColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTTable1825) GetTableColumnsOk() ([]BTTableColumnInfo1222, bool) {
	if o == nil || o.TableColumns == nil {
		return nil, false
	}
	return o.TableColumns, true
}

// HasTableColumns returns a boolean if a field has been set.
func (o *base_BTTable1825) HasTableColumns() bool {
	if o != nil && o.TableColumns != nil {
		return true
	}

	return false
}

// SetTableColumns gets a reference to the given []BTTableColumnInfo1222 and assigns it to the TableColumns field.
func (o *base_BTTable1825) SetTableColumns(v []BTTableColumnInfo1222) {
	o.TableColumns = v
}

// GetTableId returns the TableId field value if set, zero value otherwise.
func (o *base_BTTable1825) GetTableId() string {
	if o == nil || o.TableId == nil {
		var ret string
		return ret
	}
	return *o.TableId
}

// GetTableIdOk returns a tuple with the TableId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTTable1825) GetTableIdOk() (*string, bool) {
	if o == nil || o.TableId == nil {
		return nil, false
	}
	return o.TableId, true
}

// HasTableId returns a boolean if a field has been set.
func (o *base_BTTable1825) HasTableId() bool {
	if o != nil && o.TableId != nil {
		return true
	}

	return false
}

// SetTableId gets a reference to the given string and assigns it to the TableId field.
func (o *base_BTTable1825) SetTableId(v string) {
	o.TableId = &v
}

// GetTableRows returns the TableRows field value if set, zero value otherwise.
func (o *base_BTTable1825) GetTableRows() []BTTableRow1054 {
	if o == nil || o.TableRows == nil {
		var ret []BTTableRow1054
		return ret
	}
	return o.TableRows
}

// GetTableRowsOk returns a tuple with the TableRows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTTable1825) GetTableRowsOk() ([]BTTableRow1054, bool) {
	if o == nil || o.TableRows == nil {
		return nil, false
	}
	return o.TableRows, true
}

// HasTableRows returns a boolean if a field has been set.
func (o *base_BTTable1825) HasTableRows() bool {
	if o != nil && o.TableRows != nil {
		return true
	}

	return false
}

// SetTableRows gets a reference to the given []BTTableRow1054 and assigns it to the TableRows field.
func (o *base_BTTable1825) SetTableRows(v []BTTableRow1054) {
	o.TableRows = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *base_BTTable1825) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTTable1825) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *base_BTTable1825) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *base_BTTable1825) SetTitle(v string) {
	o.Title = &v
}

func (o base_BTTable1825) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllRowValues != nil {
		toSerialize["allRowValues"] = o.AllRowValues
	}
	if o.ColumnCount != nil {
		toSerialize["columnCount"] = o.ColumnCount
	}
	if o.FrozenColumns != nil {
		toSerialize["frozenColumns"] = o.FrozenColumns
	}
	if o.IsFailed != nil {
		toSerialize["isFailed"] = o.IsFailed
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.ReadOnly != nil {
		toSerialize["readOnly"] = o.ReadOnly
	}
	if o.RowCount != nil {
		toSerialize["rowCount"] = o.RowCount
	}
	if o.SortOrder != nil {
		toSerialize["sortOrder"] = o.SortOrder
	}
	if o.TableColumns != nil {
		toSerialize["tableColumns"] = o.TableColumns
	}
	if o.TableId != nil {
		toSerialize["tableId"] = o.TableId
	}
	if o.TableRows != nil {
		toSerialize["tableRows"] = o.TableRows
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	return json.Marshal(toSerialize)
}
