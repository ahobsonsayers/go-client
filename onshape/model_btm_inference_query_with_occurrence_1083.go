/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.153.6344-db89a80956dd
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMInferenceQueryWithOccurrence1083 struct for BTMInferenceQueryWithOccurrence1083
type BTMInferenceQueryWithOccurrence1083 struct {
	BtType                *string                    `json:"btType,omitempty"`
	DeterministicIdList   *BTMIndividualQueryBase139 `json:"deterministicIdList,omitempty"`
	DeterministicIds      []string                   `json:"deterministicIds,omitempty"`
	FullPathAsString      *string                    `json:"fullPathAsString,omitempty"`
	ImportMicroversion    *string                    `json:"importMicroversion,omitempty"`
	NodeId                *string                    `json:"nodeId,omitempty"`
	Occurrence            *BTOccurrence74            `json:"occurrence,omitempty"`
	Path                  []string                   `json:"path,omitempty"`
	Query                 *BTMIndividualQueryBase139 `json:"query,omitempty"`
	QueryString           *string                    `json:"queryString,omitempty"`
	EntityQuery           *string                    `json:"entityQuery,omitempty"`
	InferenceType         *string                    `json:"inferenceType,omitempty"`
	SecondDeterministicId *string                    `json:"secondDeterministicId,omitempty"`
	SecondEntityQuery     *string                    `json:"secondEntityQuery,omitempty"`
}

// NewBTMInferenceQueryWithOccurrence1083 instantiates a new BTMInferenceQueryWithOccurrence1083 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMInferenceQueryWithOccurrence1083() *BTMInferenceQueryWithOccurrence1083 {
	this := BTMInferenceQueryWithOccurrence1083{}
	return &this
}

// NewBTMInferenceQueryWithOccurrence1083WithDefaults instantiates a new BTMInferenceQueryWithOccurrence1083 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMInferenceQueryWithOccurrence1083WithDefaults() *BTMInferenceQueryWithOccurrence1083 {
	this := BTMInferenceQueryWithOccurrence1083{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMInferenceQueryWithOccurrence1083) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMInferenceQueryWithOccurrence1083) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMInferenceQueryWithOccurrence1083) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMInferenceQueryWithOccurrence1083) SetBtType(v string) {
	o.BtType = &v
}

// GetDeterministicIdList returns the DeterministicIdList field value if set, zero value otherwise.
func (o *BTMInferenceQueryWithOccurrence1083) GetDeterministicIdList() BTMIndividualQueryBase139 {
	if o == nil || o.DeterministicIdList == nil {
		var ret BTMIndividualQueryBase139
		return ret
	}
	return *o.DeterministicIdList
}

// GetDeterministicIdListOk returns a tuple with the DeterministicIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMInferenceQueryWithOccurrence1083) GetDeterministicIdListOk() (*BTMIndividualQueryBase139, bool) {
	if o == nil || o.DeterministicIdList == nil {
		return nil, false
	}
	return o.DeterministicIdList, true
}

// HasDeterministicIdList returns a boolean if a field has been set.
func (o *BTMInferenceQueryWithOccurrence1083) HasDeterministicIdList() bool {
	if o != nil && o.DeterministicIdList != nil {
		return true
	}

	return false
}

// SetDeterministicIdList gets a reference to the given BTMIndividualQueryBase139 and assigns it to the DeterministicIdList field.
func (o *BTMInferenceQueryWithOccurrence1083) SetDeterministicIdList(v BTMIndividualQueryBase139) {
	o.DeterministicIdList = &v
}

// GetDeterministicIds returns the DeterministicIds field value if set, zero value otherwise.
func (o *BTMInferenceQueryWithOccurrence1083) GetDeterministicIds() []string {
	if o == nil || o.DeterministicIds == nil {
		var ret []string
		return ret
	}
	return o.DeterministicIds
}

// GetDeterministicIdsOk returns a tuple with the DeterministicIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMInferenceQueryWithOccurrence1083) GetDeterministicIdsOk() ([]string, bool) {
	if o == nil || o.DeterministicIds == nil {
		return nil, false
	}
	return o.DeterministicIds, true
}

// HasDeterministicIds returns a boolean if a field has been set.
func (o *BTMInferenceQueryWithOccurrence1083) HasDeterministicIds() bool {
	if o != nil && o.DeterministicIds != nil {
		return true
	}

	return false
}

// SetDeterministicIds gets a reference to the given []string and assigns it to the DeterministicIds field.
func (o *BTMInferenceQueryWithOccurrence1083) SetDeterministicIds(v []string) {
	o.DeterministicIds = v
}

// GetFullPathAsString returns the FullPathAsString field value if set, zero value otherwise.
func (o *BTMInferenceQueryWithOccurrence1083) GetFullPathAsString() string {
	if o == nil || o.FullPathAsString == nil {
		var ret string
		return ret
	}
	return *o.FullPathAsString
}

// GetFullPathAsStringOk returns a tuple with the FullPathAsString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMInferenceQueryWithOccurrence1083) GetFullPathAsStringOk() (*string, bool) {
	if o == nil || o.FullPathAsString == nil {
		return nil, false
	}
	return o.FullPathAsString, true
}

// HasFullPathAsString returns a boolean if a field has been set.
func (o *BTMInferenceQueryWithOccurrence1083) HasFullPathAsString() bool {
	if o != nil && o.FullPathAsString != nil {
		return true
	}

	return false
}

// SetFullPathAsString gets a reference to the given string and assigns it to the FullPathAsString field.
func (o *BTMInferenceQueryWithOccurrence1083) SetFullPathAsString(v string) {
	o.FullPathAsString = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMInferenceQueryWithOccurrence1083) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMInferenceQueryWithOccurrence1083) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMInferenceQueryWithOccurrence1083) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMInferenceQueryWithOccurrence1083) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMInferenceQueryWithOccurrence1083) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMInferenceQueryWithOccurrence1083) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMInferenceQueryWithOccurrence1083) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMInferenceQueryWithOccurrence1083) SetNodeId(v string) {
	o.NodeId = &v
}

// GetOccurrence returns the Occurrence field value if set, zero value otherwise.
func (o *BTMInferenceQueryWithOccurrence1083) GetOccurrence() BTOccurrence74 {
	if o == nil || o.Occurrence == nil {
		var ret BTOccurrence74
		return ret
	}
	return *o.Occurrence
}

// GetOccurrenceOk returns a tuple with the Occurrence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMInferenceQueryWithOccurrence1083) GetOccurrenceOk() (*BTOccurrence74, bool) {
	if o == nil || o.Occurrence == nil {
		return nil, false
	}
	return o.Occurrence, true
}

// HasOccurrence returns a boolean if a field has been set.
func (o *BTMInferenceQueryWithOccurrence1083) HasOccurrence() bool {
	if o != nil && o.Occurrence != nil {
		return true
	}

	return false
}

// SetOccurrence gets a reference to the given BTOccurrence74 and assigns it to the Occurrence field.
func (o *BTMInferenceQueryWithOccurrence1083) SetOccurrence(v BTOccurrence74) {
	o.Occurrence = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *BTMInferenceQueryWithOccurrence1083) GetPath() []string {
	if o == nil || o.Path == nil {
		var ret []string
		return ret
	}
	return o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMInferenceQueryWithOccurrence1083) GetPathOk() ([]string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *BTMInferenceQueryWithOccurrence1083) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given []string and assigns it to the Path field.
func (o *BTMInferenceQueryWithOccurrence1083) SetPath(v []string) {
	o.Path = v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *BTMInferenceQueryWithOccurrence1083) GetQuery() BTMIndividualQueryBase139 {
	if o == nil || o.Query == nil {
		var ret BTMIndividualQueryBase139
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMInferenceQueryWithOccurrence1083) GetQueryOk() (*BTMIndividualQueryBase139, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *BTMInferenceQueryWithOccurrence1083) HasQuery() bool {
	if o != nil && o.Query != nil {
		return true
	}

	return false
}

// SetQuery gets a reference to the given BTMIndividualQueryBase139 and assigns it to the Query field.
func (o *BTMInferenceQueryWithOccurrence1083) SetQuery(v BTMIndividualQueryBase139) {
	o.Query = &v
}

// GetQueryString returns the QueryString field value if set, zero value otherwise.
func (o *BTMInferenceQueryWithOccurrence1083) GetQueryString() string {
	if o == nil || o.QueryString == nil {
		var ret string
		return ret
	}
	return *o.QueryString
}

// GetQueryStringOk returns a tuple with the QueryString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMInferenceQueryWithOccurrence1083) GetQueryStringOk() (*string, bool) {
	if o == nil || o.QueryString == nil {
		return nil, false
	}
	return o.QueryString, true
}

// HasQueryString returns a boolean if a field has been set.
func (o *BTMInferenceQueryWithOccurrence1083) HasQueryString() bool {
	if o != nil && o.QueryString != nil {
		return true
	}

	return false
}

// SetQueryString gets a reference to the given string and assigns it to the QueryString field.
func (o *BTMInferenceQueryWithOccurrence1083) SetQueryString(v string) {
	o.QueryString = &v
}

// GetEntityQuery returns the EntityQuery field value if set, zero value otherwise.
func (o *BTMInferenceQueryWithOccurrence1083) GetEntityQuery() string {
	if o == nil || o.EntityQuery == nil {
		var ret string
		return ret
	}
	return *o.EntityQuery
}

// GetEntityQueryOk returns a tuple with the EntityQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMInferenceQueryWithOccurrence1083) GetEntityQueryOk() (*string, bool) {
	if o == nil || o.EntityQuery == nil {
		return nil, false
	}
	return o.EntityQuery, true
}

// HasEntityQuery returns a boolean if a field has been set.
func (o *BTMInferenceQueryWithOccurrence1083) HasEntityQuery() bool {
	if o != nil && o.EntityQuery != nil {
		return true
	}

	return false
}

// SetEntityQuery gets a reference to the given string and assigns it to the EntityQuery field.
func (o *BTMInferenceQueryWithOccurrence1083) SetEntityQuery(v string) {
	o.EntityQuery = &v
}

// GetInferenceType returns the InferenceType field value if set, zero value otherwise.
func (o *BTMInferenceQueryWithOccurrence1083) GetInferenceType() string {
	if o == nil || o.InferenceType == nil {
		var ret string
		return ret
	}
	return *o.InferenceType
}

// GetInferenceTypeOk returns a tuple with the InferenceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMInferenceQueryWithOccurrence1083) GetInferenceTypeOk() (*string, bool) {
	if o == nil || o.InferenceType == nil {
		return nil, false
	}
	return o.InferenceType, true
}

// HasInferenceType returns a boolean if a field has been set.
func (o *BTMInferenceQueryWithOccurrence1083) HasInferenceType() bool {
	if o != nil && o.InferenceType != nil {
		return true
	}

	return false
}

// SetInferenceType gets a reference to the given string and assigns it to the InferenceType field.
func (o *BTMInferenceQueryWithOccurrence1083) SetInferenceType(v string) {
	o.InferenceType = &v
}

// GetSecondDeterministicId returns the SecondDeterministicId field value if set, zero value otherwise.
func (o *BTMInferenceQueryWithOccurrence1083) GetSecondDeterministicId() string {
	if o == nil || o.SecondDeterministicId == nil {
		var ret string
		return ret
	}
	return *o.SecondDeterministicId
}

// GetSecondDeterministicIdOk returns a tuple with the SecondDeterministicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMInferenceQueryWithOccurrence1083) GetSecondDeterministicIdOk() (*string, bool) {
	if o == nil || o.SecondDeterministicId == nil {
		return nil, false
	}
	return o.SecondDeterministicId, true
}

// HasSecondDeterministicId returns a boolean if a field has been set.
func (o *BTMInferenceQueryWithOccurrence1083) HasSecondDeterministicId() bool {
	if o != nil && o.SecondDeterministicId != nil {
		return true
	}

	return false
}

// SetSecondDeterministicId gets a reference to the given string and assigns it to the SecondDeterministicId field.
func (o *BTMInferenceQueryWithOccurrence1083) SetSecondDeterministicId(v string) {
	o.SecondDeterministicId = &v
}

// GetSecondEntityQuery returns the SecondEntityQuery field value if set, zero value otherwise.
func (o *BTMInferenceQueryWithOccurrence1083) GetSecondEntityQuery() string {
	if o == nil || o.SecondEntityQuery == nil {
		var ret string
		return ret
	}
	return *o.SecondEntityQuery
}

// GetSecondEntityQueryOk returns a tuple with the SecondEntityQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMInferenceQueryWithOccurrence1083) GetSecondEntityQueryOk() (*string, bool) {
	if o == nil || o.SecondEntityQuery == nil {
		return nil, false
	}
	return o.SecondEntityQuery, true
}

// HasSecondEntityQuery returns a boolean if a field has been set.
func (o *BTMInferenceQueryWithOccurrence1083) HasSecondEntityQuery() bool {
	if o != nil && o.SecondEntityQuery != nil {
		return true
	}

	return false
}

// SetSecondEntityQuery gets a reference to the given string and assigns it to the SecondEntityQuery field.
func (o *BTMInferenceQueryWithOccurrence1083) SetSecondEntityQuery(v string) {
	o.SecondEntityQuery = &v
}

func (o BTMInferenceQueryWithOccurrence1083) MarshalJSON() ([]byte, error) {
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
	if o.FullPathAsString != nil {
		toSerialize["fullPathAsString"] = o.FullPathAsString
	}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.Occurrence != nil {
		toSerialize["occurrence"] = o.Occurrence
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.QueryString != nil {
		toSerialize["queryString"] = o.QueryString
	}
	if o.EntityQuery != nil {
		toSerialize["entityQuery"] = o.EntityQuery
	}
	if o.InferenceType != nil {
		toSerialize["inferenceType"] = o.InferenceType
	}
	if o.SecondDeterministicId != nil {
		toSerialize["secondDeterministicId"] = o.SecondDeterministicId
	}
	if o.SecondEntityQuery != nil {
		toSerialize["secondEntityQuery"] = o.SecondEntityQuery
	}
	return json.Marshal(toSerialize)
}

type NullableBTMInferenceQueryWithOccurrence1083 struct {
	value *BTMInferenceQueryWithOccurrence1083
	isSet bool
}

func (v NullableBTMInferenceQueryWithOccurrence1083) Get() *BTMInferenceQueryWithOccurrence1083 {
	return v.value
}

func (v *NullableBTMInferenceQueryWithOccurrence1083) Set(val *BTMInferenceQueryWithOccurrence1083) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMInferenceQueryWithOccurrence1083) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMInferenceQueryWithOccurrence1083) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMInferenceQueryWithOccurrence1083(val *BTMInferenceQueryWithOccurrence1083) *NullableBTMInferenceQueryWithOccurrence1083 {
	return &NullableBTMInferenceQueryWithOccurrence1083{value: val, isSet: true}
}

func (v NullableBTMInferenceQueryWithOccurrence1083) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMInferenceQueryWithOccurrence1083) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
