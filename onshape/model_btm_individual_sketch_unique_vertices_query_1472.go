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

// BTMIndividualSketchUniqueVerticesQuery1472 struct for BTMIndividualSketchUniqueVerticesQuery1472
type BTMIndividualSketchUniqueVerticesQuery1472 struct {
	BtType              *string                    `json:"btType,omitempty"`
	DeterministicIdList *BTMIndividualQueryBase139 `json:"deterministicIdList,omitempty"`
	DeterministicIds    []string                   `json:"deterministicIds,omitempty"`
	ImportMicroversion  *string                    `json:"importMicroversion,omitempty"`
	NodeId              *string                    `json:"nodeId,omitempty"`
	Query               *BTMIndividualQueryBase139 `json:"query,omitempty"`
	QueryString         *string                    `json:"queryString,omitempty"`
	PersistentQuery     *BTPStatement269           `json:"persistentQuery,omitempty"`
	QueryStatement      *BTPStatement269           `json:"queryStatement,omitempty"`
	VariableName        *BTMIndividualQuery138     `json:"variableName,omitempty"`
	FeatureId           *string                    `json:"featureId,omitempty"`
}

// NewBTMIndividualSketchUniqueVerticesQuery1472 instantiates a new BTMIndividualSketchUniqueVerticesQuery1472 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMIndividualSketchUniqueVerticesQuery1472() *BTMIndividualSketchUniqueVerticesQuery1472 {
	this := BTMIndividualSketchUniqueVerticesQuery1472{}
	return &this
}

// NewBTMIndividualSketchUniqueVerticesQuery1472WithDefaults instantiates a new BTMIndividualSketchUniqueVerticesQuery1472 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMIndividualSketchUniqueVerticesQuery1472WithDefaults() *BTMIndividualSketchUniqueVerticesQuery1472 {
	this := BTMIndividualSketchUniqueVerticesQuery1472{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMIndividualSketchUniqueVerticesQuery1472) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualSketchUniqueVerticesQuery1472) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMIndividualSketchUniqueVerticesQuery1472) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMIndividualSketchUniqueVerticesQuery1472) SetBtType(v string) {
	o.BtType = &v
}

// GetDeterministicIdList returns the DeterministicIdList field value if set, zero value otherwise.
func (o *BTMIndividualSketchUniqueVerticesQuery1472) GetDeterministicIdList() BTMIndividualQueryBase139 {
	if o == nil || o.DeterministicIdList == nil {
		var ret BTMIndividualQueryBase139
		return ret
	}
	return *o.DeterministicIdList
}

// GetDeterministicIdListOk returns a tuple with the DeterministicIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualSketchUniqueVerticesQuery1472) GetDeterministicIdListOk() (*BTMIndividualQueryBase139, bool) {
	if o == nil || o.DeterministicIdList == nil {
		return nil, false
	}
	return o.DeterministicIdList, true
}

// HasDeterministicIdList returns a boolean if a field has been set.
func (o *BTMIndividualSketchUniqueVerticesQuery1472) HasDeterministicIdList() bool {
	if o != nil && o.DeterministicIdList != nil {
		return true
	}

	return false
}

// SetDeterministicIdList gets a reference to the given BTMIndividualQueryBase139 and assigns it to the DeterministicIdList field.
func (o *BTMIndividualSketchUniqueVerticesQuery1472) SetDeterministicIdList(v BTMIndividualQueryBase139) {
	o.DeterministicIdList = &v
}

// GetDeterministicIds returns the DeterministicIds field value if set, zero value otherwise.
func (o *BTMIndividualSketchUniqueVerticesQuery1472) GetDeterministicIds() []string {
	if o == nil || o.DeterministicIds == nil {
		var ret []string
		return ret
	}
	return o.DeterministicIds
}

// GetDeterministicIdsOk returns a tuple with the DeterministicIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualSketchUniqueVerticesQuery1472) GetDeterministicIdsOk() ([]string, bool) {
	if o == nil || o.DeterministicIds == nil {
		return nil, false
	}
	return o.DeterministicIds, true
}

// HasDeterministicIds returns a boolean if a field has been set.
func (o *BTMIndividualSketchUniqueVerticesQuery1472) HasDeterministicIds() bool {
	if o != nil && o.DeterministicIds != nil {
		return true
	}

	return false
}

// SetDeterministicIds gets a reference to the given []string and assigns it to the DeterministicIds field.
func (o *BTMIndividualSketchUniqueVerticesQuery1472) SetDeterministicIds(v []string) {
	o.DeterministicIds = v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMIndividualSketchUniqueVerticesQuery1472) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualSketchUniqueVerticesQuery1472) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMIndividualSketchUniqueVerticesQuery1472) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMIndividualSketchUniqueVerticesQuery1472) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMIndividualSketchUniqueVerticesQuery1472) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualSketchUniqueVerticesQuery1472) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMIndividualSketchUniqueVerticesQuery1472) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMIndividualSketchUniqueVerticesQuery1472) SetNodeId(v string) {
	o.NodeId = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *BTMIndividualSketchUniqueVerticesQuery1472) GetQuery() BTMIndividualQueryBase139 {
	if o == nil || o.Query == nil {
		var ret BTMIndividualQueryBase139
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualSketchUniqueVerticesQuery1472) GetQueryOk() (*BTMIndividualQueryBase139, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *BTMIndividualSketchUniqueVerticesQuery1472) HasQuery() bool {
	if o != nil && o.Query != nil {
		return true
	}

	return false
}

// SetQuery gets a reference to the given BTMIndividualQueryBase139 and assigns it to the Query field.
func (o *BTMIndividualSketchUniqueVerticesQuery1472) SetQuery(v BTMIndividualQueryBase139) {
	o.Query = &v
}

// GetQueryString returns the QueryString field value if set, zero value otherwise.
func (o *BTMIndividualSketchUniqueVerticesQuery1472) GetQueryString() string {
	if o == nil || o.QueryString == nil {
		var ret string
		return ret
	}
	return *o.QueryString
}

// GetQueryStringOk returns a tuple with the QueryString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualSketchUniqueVerticesQuery1472) GetQueryStringOk() (*string, bool) {
	if o == nil || o.QueryString == nil {
		return nil, false
	}
	return o.QueryString, true
}

// HasQueryString returns a boolean if a field has been set.
func (o *BTMIndividualSketchUniqueVerticesQuery1472) HasQueryString() bool {
	if o != nil && o.QueryString != nil {
		return true
	}

	return false
}

// SetQueryString gets a reference to the given string and assigns it to the QueryString field.
func (o *BTMIndividualSketchUniqueVerticesQuery1472) SetQueryString(v string) {
	o.QueryString = &v
}

// GetPersistentQuery returns the PersistentQuery field value if set, zero value otherwise.
func (o *BTMIndividualSketchUniqueVerticesQuery1472) GetPersistentQuery() BTPStatement269 {
	if o == nil || o.PersistentQuery == nil {
		var ret BTPStatement269
		return ret
	}
	return *o.PersistentQuery
}

// GetPersistentQueryOk returns a tuple with the PersistentQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualSketchUniqueVerticesQuery1472) GetPersistentQueryOk() (*BTPStatement269, bool) {
	if o == nil || o.PersistentQuery == nil {
		return nil, false
	}
	return o.PersistentQuery, true
}

// HasPersistentQuery returns a boolean if a field has been set.
func (o *BTMIndividualSketchUniqueVerticesQuery1472) HasPersistentQuery() bool {
	if o != nil && o.PersistentQuery != nil {
		return true
	}

	return false
}

// SetPersistentQuery gets a reference to the given BTPStatement269 and assigns it to the PersistentQuery field.
func (o *BTMIndividualSketchUniqueVerticesQuery1472) SetPersistentQuery(v BTPStatement269) {
	o.PersistentQuery = &v
}

// GetQueryStatement returns the QueryStatement field value if set, zero value otherwise.
func (o *BTMIndividualSketchUniqueVerticesQuery1472) GetQueryStatement() BTPStatement269 {
	if o == nil || o.QueryStatement == nil {
		var ret BTPStatement269
		return ret
	}
	return *o.QueryStatement
}

// GetQueryStatementOk returns a tuple with the QueryStatement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualSketchUniqueVerticesQuery1472) GetQueryStatementOk() (*BTPStatement269, bool) {
	if o == nil || o.QueryStatement == nil {
		return nil, false
	}
	return o.QueryStatement, true
}

// HasQueryStatement returns a boolean if a field has been set.
func (o *BTMIndividualSketchUniqueVerticesQuery1472) HasQueryStatement() bool {
	if o != nil && o.QueryStatement != nil {
		return true
	}

	return false
}

// SetQueryStatement gets a reference to the given BTPStatement269 and assigns it to the QueryStatement field.
func (o *BTMIndividualSketchUniqueVerticesQuery1472) SetQueryStatement(v BTPStatement269) {
	o.QueryStatement = &v
}

// GetVariableName returns the VariableName field value if set, zero value otherwise.
func (o *BTMIndividualSketchUniqueVerticesQuery1472) GetVariableName() BTMIndividualQuery138 {
	if o == nil || o.VariableName == nil {
		var ret BTMIndividualQuery138
		return ret
	}
	return *o.VariableName
}

// GetVariableNameOk returns a tuple with the VariableName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualSketchUniqueVerticesQuery1472) GetVariableNameOk() (*BTMIndividualQuery138, bool) {
	if o == nil || o.VariableName == nil {
		return nil, false
	}
	return o.VariableName, true
}

// HasVariableName returns a boolean if a field has been set.
func (o *BTMIndividualSketchUniqueVerticesQuery1472) HasVariableName() bool {
	if o != nil && o.VariableName != nil {
		return true
	}

	return false
}

// SetVariableName gets a reference to the given BTMIndividualQuery138 and assigns it to the VariableName field.
func (o *BTMIndividualSketchUniqueVerticesQuery1472) SetVariableName(v BTMIndividualQuery138) {
	o.VariableName = &v
}

// GetFeatureId returns the FeatureId field value if set, zero value otherwise.
func (o *BTMIndividualSketchUniqueVerticesQuery1472) GetFeatureId() string {
	if o == nil || o.FeatureId == nil {
		var ret string
		return ret
	}
	return *o.FeatureId
}

// GetFeatureIdOk returns a tuple with the FeatureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualSketchUniqueVerticesQuery1472) GetFeatureIdOk() (*string, bool) {
	if o == nil || o.FeatureId == nil {
		return nil, false
	}
	return o.FeatureId, true
}

// HasFeatureId returns a boolean if a field has been set.
func (o *BTMIndividualSketchUniqueVerticesQuery1472) HasFeatureId() bool {
	if o != nil && o.FeatureId != nil {
		return true
	}

	return false
}

// SetFeatureId gets a reference to the given string and assigns it to the FeatureId field.
func (o *BTMIndividualSketchUniqueVerticesQuery1472) SetFeatureId(v string) {
	o.FeatureId = &v
}

func (o BTMIndividualSketchUniqueVerticesQuery1472) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.DeterministicIdList != nil {
		toSerialize["deterministicIdList"] = o.DeterministicIdList
	}
	if o.DeterministicIds != nil {
		toSerialize["deterministicIds"] = o.DeterministicIds
	}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.QueryString != nil {
		toSerialize["queryString"] = o.QueryString
	}
	if o.PersistentQuery != nil {
		toSerialize["persistentQuery"] = o.PersistentQuery
	}
	if o.QueryStatement != nil {
		toSerialize["queryStatement"] = o.QueryStatement
	}
	if o.VariableName != nil {
		toSerialize["variableName"] = o.VariableName
	}
	if o.FeatureId != nil {
		toSerialize["featureId"] = o.FeatureId
	}
	return json.Marshal(toSerialize)
}

type NullableBTMIndividualSketchUniqueVerticesQuery1472 struct {
	value *BTMIndividualSketchUniqueVerticesQuery1472
	isSet bool
}

func (v NullableBTMIndividualSketchUniqueVerticesQuery1472) Get() *BTMIndividualSketchUniqueVerticesQuery1472 {
	return v.value
}

func (v *NullableBTMIndividualSketchUniqueVerticesQuery1472) Set(val *BTMIndividualSketchUniqueVerticesQuery1472) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMIndividualSketchUniqueVerticesQuery1472) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMIndividualSketchUniqueVerticesQuery1472) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMIndividualSketchUniqueVerticesQuery1472(val *BTMIndividualSketchUniqueVerticesQuery1472) *NullableBTMIndividualSketchUniqueVerticesQuery1472 {
	return &NullableBTMIndividualSketchUniqueVerticesQuery1472{value: val, isSet: true}
}

func (v NullableBTMIndividualSketchUniqueVerticesQuery1472) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMIndividualSketchUniqueVerticesQuery1472) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
