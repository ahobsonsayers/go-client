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

// BTConfiguredPartPropertiesTable2740 struct for BTConfiguredPartPropertiesTable2740
type BTConfiguredPartPropertiesTable2740 struct {
	AllRowValues         [][]string              `json:"allRowValues,omitempty"`
	ColumnCount          *int32                  `json:"columnCount,omitempty"`
	FrozenColumns        *int32                  `json:"frozenColumns,omitempty"`
	IsFailed             *bool                   `json:"isFailed,omitempty"`
	NodeId               *string                 `json:"nodeId,omitempty"`
	ReadOnly             *bool                   `json:"readOnly,omitempty"`
	RowCount             *int32                  `json:"rowCount,omitempty"`
	SortOrder            *BTTableSortOrder4371   `json:"sortOrder,omitempty"`
	TableColumns         []BTTableColumnInfo1222 `json:"tableColumns,omitempty"`
	TableId              *string                 `json:"tableId,omitempty"`
	TableRows            []BTTableRow1054        `json:"tableRows,omitempty"`
	Title                *string                 `json:"title,omitempty"`
	BtType               *string                 `json:"btType,omitempty"`
	PartDeterministicId  *string                 `json:"partDeterministicId,omitempty"`
	PartDeterministicIds []string                `json:"partDeterministicIds,omitempty"`
	PropertyNodeId       *string                 `json:"propertyNodeId,omitempty"`
}

// NewBTConfiguredPartPropertiesTable2740 instantiates a new BTConfiguredPartPropertiesTable2740 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTConfiguredPartPropertiesTable2740() *BTConfiguredPartPropertiesTable2740 {
	this := BTConfiguredPartPropertiesTable2740{}
	return &this
}

// NewBTConfiguredPartPropertiesTable2740WithDefaults instantiates a new BTConfiguredPartPropertiesTable2740 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTConfiguredPartPropertiesTable2740WithDefaults() *BTConfiguredPartPropertiesTable2740 {
	this := BTConfiguredPartPropertiesTable2740{}
	return &this
}

// GetAllRowValues returns the AllRowValues field value if set, zero value otherwise.
func (o *BTConfiguredPartPropertiesTable2740) GetAllRowValues() [][]string {
	if o == nil || o.AllRowValues == nil {
		var ret [][]string
		return ret
	}
	return o.AllRowValues
}

// GetAllRowValuesOk returns a tuple with the AllRowValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredPartPropertiesTable2740) GetAllRowValuesOk() ([][]string, bool) {
	if o == nil || o.AllRowValues == nil {
		return nil, false
	}
	return o.AllRowValues, true
}

// HasAllRowValues returns a boolean if a field has been set.
func (o *BTConfiguredPartPropertiesTable2740) HasAllRowValues() bool {
	if o != nil && o.AllRowValues != nil {
		return true
	}

	return false
}

// SetAllRowValues gets a reference to the given [][]string and assigns it to the AllRowValues field.
func (o *BTConfiguredPartPropertiesTable2740) SetAllRowValues(v [][]string) {
	o.AllRowValues = v
}

// GetColumnCount returns the ColumnCount field value if set, zero value otherwise.
func (o *BTConfiguredPartPropertiesTable2740) GetColumnCount() int32 {
	if o == nil || o.ColumnCount == nil {
		var ret int32
		return ret
	}
	return *o.ColumnCount
}

// GetColumnCountOk returns a tuple with the ColumnCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredPartPropertiesTable2740) GetColumnCountOk() (*int32, bool) {
	if o == nil || o.ColumnCount == nil {
		return nil, false
	}
	return o.ColumnCount, true
}

// HasColumnCount returns a boolean if a field has been set.
func (o *BTConfiguredPartPropertiesTable2740) HasColumnCount() bool {
	if o != nil && o.ColumnCount != nil {
		return true
	}

	return false
}

// SetColumnCount gets a reference to the given int32 and assigns it to the ColumnCount field.
func (o *BTConfiguredPartPropertiesTable2740) SetColumnCount(v int32) {
	o.ColumnCount = &v
}

// GetFrozenColumns returns the FrozenColumns field value if set, zero value otherwise.
func (o *BTConfiguredPartPropertiesTable2740) GetFrozenColumns() int32 {
	if o == nil || o.FrozenColumns == nil {
		var ret int32
		return ret
	}
	return *o.FrozenColumns
}

// GetFrozenColumnsOk returns a tuple with the FrozenColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredPartPropertiesTable2740) GetFrozenColumnsOk() (*int32, bool) {
	if o == nil || o.FrozenColumns == nil {
		return nil, false
	}
	return o.FrozenColumns, true
}

// HasFrozenColumns returns a boolean if a field has been set.
func (o *BTConfiguredPartPropertiesTable2740) HasFrozenColumns() bool {
	if o != nil && o.FrozenColumns != nil {
		return true
	}

	return false
}

// SetFrozenColumns gets a reference to the given int32 and assigns it to the FrozenColumns field.
func (o *BTConfiguredPartPropertiesTable2740) SetFrozenColumns(v int32) {
	o.FrozenColumns = &v
}

// GetIsFailed returns the IsFailed field value if set, zero value otherwise.
func (o *BTConfiguredPartPropertiesTable2740) GetIsFailed() bool {
	if o == nil || o.IsFailed == nil {
		var ret bool
		return ret
	}
	return *o.IsFailed
}

// GetIsFailedOk returns a tuple with the IsFailed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredPartPropertiesTable2740) GetIsFailedOk() (*bool, bool) {
	if o == nil || o.IsFailed == nil {
		return nil, false
	}
	return o.IsFailed, true
}

// HasIsFailed returns a boolean if a field has been set.
func (o *BTConfiguredPartPropertiesTable2740) HasIsFailed() bool {
	if o != nil && o.IsFailed != nil {
		return true
	}

	return false
}

// SetIsFailed gets a reference to the given bool and assigns it to the IsFailed field.
func (o *BTConfiguredPartPropertiesTable2740) SetIsFailed(v bool) {
	o.IsFailed = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTConfiguredPartPropertiesTable2740) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredPartPropertiesTable2740) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTConfiguredPartPropertiesTable2740) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTConfiguredPartPropertiesTable2740) SetNodeId(v string) {
	o.NodeId = &v
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *BTConfiguredPartPropertiesTable2740) GetReadOnly() bool {
	if o == nil || o.ReadOnly == nil {
		var ret bool
		return ret
	}
	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredPartPropertiesTable2740) GetReadOnlyOk() (*bool, bool) {
	if o == nil || o.ReadOnly == nil {
		return nil, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *BTConfiguredPartPropertiesTable2740) HasReadOnly() bool {
	if o != nil && o.ReadOnly != nil {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.
func (o *BTConfiguredPartPropertiesTable2740) SetReadOnly(v bool) {
	o.ReadOnly = &v
}

// GetRowCount returns the RowCount field value if set, zero value otherwise.
func (o *BTConfiguredPartPropertiesTable2740) GetRowCount() int32 {
	if o == nil || o.RowCount == nil {
		var ret int32
		return ret
	}
	return *o.RowCount
}

// GetRowCountOk returns a tuple with the RowCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredPartPropertiesTable2740) GetRowCountOk() (*int32, bool) {
	if o == nil || o.RowCount == nil {
		return nil, false
	}
	return o.RowCount, true
}

// HasRowCount returns a boolean if a field has been set.
func (o *BTConfiguredPartPropertiesTable2740) HasRowCount() bool {
	if o != nil && o.RowCount != nil {
		return true
	}

	return false
}

// SetRowCount gets a reference to the given int32 and assigns it to the RowCount field.
func (o *BTConfiguredPartPropertiesTable2740) SetRowCount(v int32) {
	o.RowCount = &v
}

// GetSortOrder returns the SortOrder field value if set, zero value otherwise.
func (o *BTConfiguredPartPropertiesTable2740) GetSortOrder() BTTableSortOrder4371 {
	if o == nil || o.SortOrder == nil {
		var ret BTTableSortOrder4371
		return ret
	}
	return *o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredPartPropertiesTable2740) GetSortOrderOk() (*BTTableSortOrder4371, bool) {
	if o == nil || o.SortOrder == nil {
		return nil, false
	}
	return o.SortOrder, true
}

// HasSortOrder returns a boolean if a field has been set.
func (o *BTConfiguredPartPropertiesTable2740) HasSortOrder() bool {
	if o != nil && o.SortOrder != nil {
		return true
	}

	return false
}

// SetSortOrder gets a reference to the given BTTableSortOrder4371 and assigns it to the SortOrder field.
func (o *BTConfiguredPartPropertiesTable2740) SetSortOrder(v BTTableSortOrder4371) {
	o.SortOrder = &v
}

// GetTableColumns returns the TableColumns field value if set, zero value otherwise.
func (o *BTConfiguredPartPropertiesTable2740) GetTableColumns() []BTTableColumnInfo1222 {
	if o == nil || o.TableColumns == nil {
		var ret []BTTableColumnInfo1222
		return ret
	}
	return o.TableColumns
}

// GetTableColumnsOk returns a tuple with the TableColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredPartPropertiesTable2740) GetTableColumnsOk() ([]BTTableColumnInfo1222, bool) {
	if o == nil || o.TableColumns == nil {
		return nil, false
	}
	return o.TableColumns, true
}

// HasTableColumns returns a boolean if a field has been set.
func (o *BTConfiguredPartPropertiesTable2740) HasTableColumns() bool {
	if o != nil && o.TableColumns != nil {
		return true
	}

	return false
}

// SetTableColumns gets a reference to the given []BTTableColumnInfo1222 and assigns it to the TableColumns field.
func (o *BTConfiguredPartPropertiesTable2740) SetTableColumns(v []BTTableColumnInfo1222) {
	o.TableColumns = v
}

// GetTableId returns the TableId field value if set, zero value otherwise.
func (o *BTConfiguredPartPropertiesTable2740) GetTableId() string {
	if o == nil || o.TableId == nil {
		var ret string
		return ret
	}
	return *o.TableId
}

// GetTableIdOk returns a tuple with the TableId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredPartPropertiesTable2740) GetTableIdOk() (*string, bool) {
	if o == nil || o.TableId == nil {
		return nil, false
	}
	return o.TableId, true
}

// HasTableId returns a boolean if a field has been set.
func (o *BTConfiguredPartPropertiesTable2740) HasTableId() bool {
	if o != nil && o.TableId != nil {
		return true
	}

	return false
}

// SetTableId gets a reference to the given string and assigns it to the TableId field.
func (o *BTConfiguredPartPropertiesTable2740) SetTableId(v string) {
	o.TableId = &v
}

// GetTableRows returns the TableRows field value if set, zero value otherwise.
func (o *BTConfiguredPartPropertiesTable2740) GetTableRows() []BTTableRow1054 {
	if o == nil || o.TableRows == nil {
		var ret []BTTableRow1054
		return ret
	}
	return o.TableRows
}

// GetTableRowsOk returns a tuple with the TableRows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredPartPropertiesTable2740) GetTableRowsOk() ([]BTTableRow1054, bool) {
	if o == nil || o.TableRows == nil {
		return nil, false
	}
	return o.TableRows, true
}

// HasTableRows returns a boolean if a field has been set.
func (o *BTConfiguredPartPropertiesTable2740) HasTableRows() bool {
	if o != nil && o.TableRows != nil {
		return true
	}

	return false
}

// SetTableRows gets a reference to the given []BTTableRow1054 and assigns it to the TableRows field.
func (o *BTConfiguredPartPropertiesTable2740) SetTableRows(v []BTTableRow1054) {
	o.TableRows = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *BTConfiguredPartPropertiesTable2740) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredPartPropertiesTable2740) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *BTConfiguredPartPropertiesTable2740) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *BTConfiguredPartPropertiesTable2740) SetTitle(v string) {
	o.Title = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTConfiguredPartPropertiesTable2740) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredPartPropertiesTable2740) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTConfiguredPartPropertiesTable2740) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTConfiguredPartPropertiesTable2740) SetBtType(v string) {
	o.BtType = &v
}

// GetPartDeterministicId returns the PartDeterministicId field value if set, zero value otherwise.
func (o *BTConfiguredPartPropertiesTable2740) GetPartDeterministicId() string {
	if o == nil || o.PartDeterministicId == nil {
		var ret string
		return ret
	}
	return *o.PartDeterministicId
}

// GetPartDeterministicIdOk returns a tuple with the PartDeterministicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredPartPropertiesTable2740) GetPartDeterministicIdOk() (*string, bool) {
	if o == nil || o.PartDeterministicId == nil {
		return nil, false
	}
	return o.PartDeterministicId, true
}

// HasPartDeterministicId returns a boolean if a field has been set.
func (o *BTConfiguredPartPropertiesTable2740) HasPartDeterministicId() bool {
	if o != nil && o.PartDeterministicId != nil {
		return true
	}

	return false
}

// SetPartDeterministicId gets a reference to the given string and assigns it to the PartDeterministicId field.
func (o *BTConfiguredPartPropertiesTable2740) SetPartDeterministicId(v string) {
	o.PartDeterministicId = &v
}

// GetPartDeterministicIds returns the PartDeterministicIds field value if set, zero value otherwise.
func (o *BTConfiguredPartPropertiesTable2740) GetPartDeterministicIds() []string {
	if o == nil || o.PartDeterministicIds == nil {
		var ret []string
		return ret
	}
	return o.PartDeterministicIds
}

// GetPartDeterministicIdsOk returns a tuple with the PartDeterministicIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredPartPropertiesTable2740) GetPartDeterministicIdsOk() ([]string, bool) {
	if o == nil || o.PartDeterministicIds == nil {
		return nil, false
	}
	return o.PartDeterministicIds, true
}

// HasPartDeterministicIds returns a boolean if a field has been set.
func (o *BTConfiguredPartPropertiesTable2740) HasPartDeterministicIds() bool {
	if o != nil && o.PartDeterministicIds != nil {
		return true
	}

	return false
}

// SetPartDeterministicIds gets a reference to the given []string and assigns it to the PartDeterministicIds field.
func (o *BTConfiguredPartPropertiesTable2740) SetPartDeterministicIds(v []string) {
	o.PartDeterministicIds = v
}

// GetPropertyNodeId returns the PropertyNodeId field value if set, zero value otherwise.
func (o *BTConfiguredPartPropertiesTable2740) GetPropertyNodeId() string {
	if o == nil || o.PropertyNodeId == nil {
		var ret string
		return ret
	}
	return *o.PropertyNodeId
}

// GetPropertyNodeIdOk returns a tuple with the PropertyNodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredPartPropertiesTable2740) GetPropertyNodeIdOk() (*string, bool) {
	if o == nil || o.PropertyNodeId == nil {
		return nil, false
	}
	return o.PropertyNodeId, true
}

// HasPropertyNodeId returns a boolean if a field has been set.
func (o *BTConfiguredPartPropertiesTable2740) HasPropertyNodeId() bool {
	if o != nil && o.PropertyNodeId != nil {
		return true
	}

	return false
}

// SetPropertyNodeId gets a reference to the given string and assigns it to the PropertyNodeId field.
func (o *BTConfiguredPartPropertiesTable2740) SetPropertyNodeId(v string) {
	o.PropertyNodeId = &v
}

func (o BTConfiguredPartPropertiesTable2740) MarshalJSON() ([]byte, error) {
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
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.PartDeterministicId != nil {
		toSerialize["partDeterministicId"] = o.PartDeterministicId
	}
	if o.PartDeterministicIds != nil {
		toSerialize["partDeterministicIds"] = o.PartDeterministicIds
	}
	if o.PropertyNodeId != nil {
		toSerialize["propertyNodeId"] = o.PropertyNodeId
	}
	return json.Marshal(toSerialize)
}

type NullableBTConfiguredPartPropertiesTable2740 struct {
	value *BTConfiguredPartPropertiesTable2740
	isSet bool
}

func (v NullableBTConfiguredPartPropertiesTable2740) Get() *BTConfiguredPartPropertiesTable2740 {
	return v.value
}

func (v *NullableBTConfiguredPartPropertiesTable2740) Set(val *BTConfiguredPartPropertiesTable2740) {
	v.value = val
	v.isSet = true
}

func (v NullableBTConfiguredPartPropertiesTable2740) IsSet() bool {
	return v.isSet
}

func (v *NullableBTConfiguredPartPropertiesTable2740) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTConfiguredPartPropertiesTable2740(val *BTConfiguredPartPropertiesTable2740) *NullableBTConfiguredPartPropertiesTable2740 {
	return &NullableBTConfiguredPartPropertiesTable2740{value: val, isSet: true}
}

func (v NullableBTConfiguredPartPropertiesTable2740) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTConfiguredPartPropertiesTable2740) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
