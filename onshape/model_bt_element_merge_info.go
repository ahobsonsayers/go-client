/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.161.13452-78145f970001
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTElementMergeInfo struct for BTElementMergeInfo
type BTElementMergeInfo struct {
	BranchPointElementName    *string                 `json:"branchPointElementName,omitempty"`
	BranchPointElementPath    []string                `json:"branchPointElementPath,omitempty"`
	DependentElementMergeInfo *BTElementMergeInfo     `json:"dependentElementMergeInfo,omitempty"`
	ElementDataType           *string                 `json:"elementDataType,omitempty"`
	ElementId                 *string                 `json:"elementId,omitempty"`
	ElementType               *string                 `json:"elementType,omitempty"`
	Mergeable                 *bool                   `json:"mergeable,omitempty"`
	SourceElementName         *string                 `json:"sourceElementName,omitempty"`
	SourceElementPath         []string                `json:"sourceElementPath,omitempty"`
	SourceElementStatus       *string                 `json:"sourceElementStatus,omitempty"`
	SourceModifiedAt          *JSONTime               `json:"sourceModifiedAt,omitempty"`
	SourceModifiedBy          *BTUserBasicSummaryInfo `json:"sourceModifiedBy,omitempty"`
	TargetElementName         *string                 `json:"targetElementName,omitempty"`
	TargetElementPath         []string                `json:"targetElementPath,omitempty"`
	TargetElementStatus       *string                 `json:"targetElementStatus,omitempty"`
	TargetModifiedAt          *JSONTime               `json:"targetModifiedAt,omitempty"`
	TargetModifiedBy          *BTUserBasicSummaryInfo `json:"targetModifiedBy,omitempty"`
}

// NewBTElementMergeInfo instantiates a new BTElementMergeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTElementMergeInfo() *BTElementMergeInfo {
	this := BTElementMergeInfo{}
	return &this
}

// NewBTElementMergeInfoWithDefaults instantiates a new BTElementMergeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTElementMergeInfoWithDefaults() *BTElementMergeInfo {
	this := BTElementMergeInfo{}
	return &this
}

// GetBranchPointElementName returns the BranchPointElementName field value if set, zero value otherwise.
func (o *BTElementMergeInfo) GetBranchPointElementName() string {
	if o == nil || o.BranchPointElementName == nil {
		var ret string
		return ret
	}
	return *o.BranchPointElementName
}

// GetBranchPointElementNameOk returns a tuple with the BranchPointElementName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTElementMergeInfo) GetBranchPointElementNameOk() (*string, bool) {
	if o == nil || o.BranchPointElementName == nil {
		return nil, false
	}
	return o.BranchPointElementName, true
}

// HasBranchPointElementName returns a boolean if a field has been set.
func (o *BTElementMergeInfo) HasBranchPointElementName() bool {
	if o != nil && o.BranchPointElementName != nil {
		return true
	}

	return false
}

// SetBranchPointElementName gets a reference to the given string and assigns it to the BranchPointElementName field.
func (o *BTElementMergeInfo) SetBranchPointElementName(v string) {
	o.BranchPointElementName = &v
}

// GetBranchPointElementPath returns the BranchPointElementPath field value if set, zero value otherwise.
func (o *BTElementMergeInfo) GetBranchPointElementPath() []string {
	if o == nil || o.BranchPointElementPath == nil {
		var ret []string
		return ret
	}
	return o.BranchPointElementPath
}

// GetBranchPointElementPathOk returns a tuple with the BranchPointElementPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTElementMergeInfo) GetBranchPointElementPathOk() ([]string, bool) {
	if o == nil || o.BranchPointElementPath == nil {
		return nil, false
	}
	return o.BranchPointElementPath, true
}

// HasBranchPointElementPath returns a boolean if a field has been set.
func (o *BTElementMergeInfo) HasBranchPointElementPath() bool {
	if o != nil && o.BranchPointElementPath != nil {
		return true
	}

	return false
}

// SetBranchPointElementPath gets a reference to the given []string and assigns it to the BranchPointElementPath field.
func (o *BTElementMergeInfo) SetBranchPointElementPath(v []string) {
	o.BranchPointElementPath = v
}

// GetDependentElementMergeInfo returns the DependentElementMergeInfo field value if set, zero value otherwise.
func (o *BTElementMergeInfo) GetDependentElementMergeInfo() BTElementMergeInfo {
	if o == nil || o.DependentElementMergeInfo == nil {
		var ret BTElementMergeInfo
		return ret
	}
	return *o.DependentElementMergeInfo
}

// GetDependentElementMergeInfoOk returns a tuple with the DependentElementMergeInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTElementMergeInfo) GetDependentElementMergeInfoOk() (*BTElementMergeInfo, bool) {
	if o == nil || o.DependentElementMergeInfo == nil {
		return nil, false
	}
	return o.DependentElementMergeInfo, true
}

// HasDependentElementMergeInfo returns a boolean if a field has been set.
func (o *BTElementMergeInfo) HasDependentElementMergeInfo() bool {
	if o != nil && o.DependentElementMergeInfo != nil {
		return true
	}

	return false
}

// SetDependentElementMergeInfo gets a reference to the given BTElementMergeInfo and assigns it to the DependentElementMergeInfo field.
func (o *BTElementMergeInfo) SetDependentElementMergeInfo(v BTElementMergeInfo) {
	o.DependentElementMergeInfo = &v
}

// GetElementDataType returns the ElementDataType field value if set, zero value otherwise.
func (o *BTElementMergeInfo) GetElementDataType() string {
	if o == nil || o.ElementDataType == nil {
		var ret string
		return ret
	}
	return *o.ElementDataType
}

// GetElementDataTypeOk returns a tuple with the ElementDataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTElementMergeInfo) GetElementDataTypeOk() (*string, bool) {
	if o == nil || o.ElementDataType == nil {
		return nil, false
	}
	return o.ElementDataType, true
}

// HasElementDataType returns a boolean if a field has been set.
func (o *BTElementMergeInfo) HasElementDataType() bool {
	if o != nil && o.ElementDataType != nil {
		return true
	}

	return false
}

// SetElementDataType gets a reference to the given string and assigns it to the ElementDataType field.
func (o *BTElementMergeInfo) SetElementDataType(v string) {
	o.ElementDataType = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *BTElementMergeInfo) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTElementMergeInfo) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *BTElementMergeInfo) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *BTElementMergeInfo) SetElementId(v string) {
	o.ElementId = &v
}

// GetElementType returns the ElementType field value if set, zero value otherwise.
func (o *BTElementMergeInfo) GetElementType() string {
	if o == nil || o.ElementType == nil {
		var ret string
		return ret
	}
	return *o.ElementType
}

// GetElementTypeOk returns a tuple with the ElementType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTElementMergeInfo) GetElementTypeOk() (*string, bool) {
	if o == nil || o.ElementType == nil {
		return nil, false
	}
	return o.ElementType, true
}

// HasElementType returns a boolean if a field has been set.
func (o *BTElementMergeInfo) HasElementType() bool {
	if o != nil && o.ElementType != nil {
		return true
	}

	return false
}

// SetElementType gets a reference to the given string and assigns it to the ElementType field.
func (o *BTElementMergeInfo) SetElementType(v string) {
	o.ElementType = &v
}

// GetMergeable returns the Mergeable field value if set, zero value otherwise.
func (o *BTElementMergeInfo) GetMergeable() bool {
	if o == nil || o.Mergeable == nil {
		var ret bool
		return ret
	}
	return *o.Mergeable
}

// GetMergeableOk returns a tuple with the Mergeable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTElementMergeInfo) GetMergeableOk() (*bool, bool) {
	if o == nil || o.Mergeable == nil {
		return nil, false
	}
	return o.Mergeable, true
}

// HasMergeable returns a boolean if a field has been set.
func (o *BTElementMergeInfo) HasMergeable() bool {
	if o != nil && o.Mergeable != nil {
		return true
	}

	return false
}

// SetMergeable gets a reference to the given bool and assigns it to the Mergeable field.
func (o *BTElementMergeInfo) SetMergeable(v bool) {
	o.Mergeable = &v
}

// GetSourceElementName returns the SourceElementName field value if set, zero value otherwise.
func (o *BTElementMergeInfo) GetSourceElementName() string {
	if o == nil || o.SourceElementName == nil {
		var ret string
		return ret
	}
	return *o.SourceElementName
}

// GetSourceElementNameOk returns a tuple with the SourceElementName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTElementMergeInfo) GetSourceElementNameOk() (*string, bool) {
	if o == nil || o.SourceElementName == nil {
		return nil, false
	}
	return o.SourceElementName, true
}

// HasSourceElementName returns a boolean if a field has been set.
func (o *BTElementMergeInfo) HasSourceElementName() bool {
	if o != nil && o.SourceElementName != nil {
		return true
	}

	return false
}

// SetSourceElementName gets a reference to the given string and assigns it to the SourceElementName field.
func (o *BTElementMergeInfo) SetSourceElementName(v string) {
	o.SourceElementName = &v
}

// GetSourceElementPath returns the SourceElementPath field value if set, zero value otherwise.
func (o *BTElementMergeInfo) GetSourceElementPath() []string {
	if o == nil || o.SourceElementPath == nil {
		var ret []string
		return ret
	}
	return o.SourceElementPath
}

// GetSourceElementPathOk returns a tuple with the SourceElementPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTElementMergeInfo) GetSourceElementPathOk() ([]string, bool) {
	if o == nil || o.SourceElementPath == nil {
		return nil, false
	}
	return o.SourceElementPath, true
}

// HasSourceElementPath returns a boolean if a field has been set.
func (o *BTElementMergeInfo) HasSourceElementPath() bool {
	if o != nil && o.SourceElementPath != nil {
		return true
	}

	return false
}

// SetSourceElementPath gets a reference to the given []string and assigns it to the SourceElementPath field.
func (o *BTElementMergeInfo) SetSourceElementPath(v []string) {
	o.SourceElementPath = v
}

// GetSourceElementStatus returns the SourceElementStatus field value if set, zero value otherwise.
func (o *BTElementMergeInfo) GetSourceElementStatus() string {
	if o == nil || o.SourceElementStatus == nil {
		var ret string
		return ret
	}
	return *o.SourceElementStatus
}

// GetSourceElementStatusOk returns a tuple with the SourceElementStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTElementMergeInfo) GetSourceElementStatusOk() (*string, bool) {
	if o == nil || o.SourceElementStatus == nil {
		return nil, false
	}
	return o.SourceElementStatus, true
}

// HasSourceElementStatus returns a boolean if a field has been set.
func (o *BTElementMergeInfo) HasSourceElementStatus() bool {
	if o != nil && o.SourceElementStatus != nil {
		return true
	}

	return false
}

// SetSourceElementStatus gets a reference to the given string and assigns it to the SourceElementStatus field.
func (o *BTElementMergeInfo) SetSourceElementStatus(v string) {
	o.SourceElementStatus = &v
}

// GetSourceModifiedAt returns the SourceModifiedAt field value if set, zero value otherwise.
func (o *BTElementMergeInfo) GetSourceModifiedAt() JSONTime {
	if o == nil || o.SourceModifiedAt == nil {
		var ret JSONTime
		return ret
	}
	return *o.SourceModifiedAt
}

// GetSourceModifiedAtOk returns a tuple with the SourceModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTElementMergeInfo) GetSourceModifiedAtOk() (*JSONTime, bool) {
	if o == nil || o.SourceModifiedAt == nil {
		return nil, false
	}
	return o.SourceModifiedAt, true
}

// HasSourceModifiedAt returns a boolean if a field has been set.
func (o *BTElementMergeInfo) HasSourceModifiedAt() bool {
	if o != nil && o.SourceModifiedAt != nil {
		return true
	}

	return false
}

// SetSourceModifiedAt gets a reference to the given JSONTime and assigns it to the SourceModifiedAt field.
func (o *BTElementMergeInfo) SetSourceModifiedAt(v JSONTime) {
	o.SourceModifiedAt = &v
}

// GetSourceModifiedBy returns the SourceModifiedBy field value if set, zero value otherwise.
func (o *BTElementMergeInfo) GetSourceModifiedBy() BTUserBasicSummaryInfo {
	if o == nil || o.SourceModifiedBy == nil {
		var ret BTUserBasicSummaryInfo
		return ret
	}
	return *o.SourceModifiedBy
}

// GetSourceModifiedByOk returns a tuple with the SourceModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTElementMergeInfo) GetSourceModifiedByOk() (*BTUserBasicSummaryInfo, bool) {
	if o == nil || o.SourceModifiedBy == nil {
		return nil, false
	}
	return o.SourceModifiedBy, true
}

// HasSourceModifiedBy returns a boolean if a field has been set.
func (o *BTElementMergeInfo) HasSourceModifiedBy() bool {
	if o != nil && o.SourceModifiedBy != nil {
		return true
	}

	return false
}

// SetSourceModifiedBy gets a reference to the given BTUserBasicSummaryInfo and assigns it to the SourceModifiedBy field.
func (o *BTElementMergeInfo) SetSourceModifiedBy(v BTUserBasicSummaryInfo) {
	o.SourceModifiedBy = &v
}

// GetTargetElementName returns the TargetElementName field value if set, zero value otherwise.
func (o *BTElementMergeInfo) GetTargetElementName() string {
	if o == nil || o.TargetElementName == nil {
		var ret string
		return ret
	}
	return *o.TargetElementName
}

// GetTargetElementNameOk returns a tuple with the TargetElementName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTElementMergeInfo) GetTargetElementNameOk() (*string, bool) {
	if o == nil || o.TargetElementName == nil {
		return nil, false
	}
	return o.TargetElementName, true
}

// HasTargetElementName returns a boolean if a field has been set.
func (o *BTElementMergeInfo) HasTargetElementName() bool {
	if o != nil && o.TargetElementName != nil {
		return true
	}

	return false
}

// SetTargetElementName gets a reference to the given string and assigns it to the TargetElementName field.
func (o *BTElementMergeInfo) SetTargetElementName(v string) {
	o.TargetElementName = &v
}

// GetTargetElementPath returns the TargetElementPath field value if set, zero value otherwise.
func (o *BTElementMergeInfo) GetTargetElementPath() []string {
	if o == nil || o.TargetElementPath == nil {
		var ret []string
		return ret
	}
	return o.TargetElementPath
}

// GetTargetElementPathOk returns a tuple with the TargetElementPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTElementMergeInfo) GetTargetElementPathOk() ([]string, bool) {
	if o == nil || o.TargetElementPath == nil {
		return nil, false
	}
	return o.TargetElementPath, true
}

// HasTargetElementPath returns a boolean if a field has been set.
func (o *BTElementMergeInfo) HasTargetElementPath() bool {
	if o != nil && o.TargetElementPath != nil {
		return true
	}

	return false
}

// SetTargetElementPath gets a reference to the given []string and assigns it to the TargetElementPath field.
func (o *BTElementMergeInfo) SetTargetElementPath(v []string) {
	o.TargetElementPath = v
}

// GetTargetElementStatus returns the TargetElementStatus field value if set, zero value otherwise.
func (o *BTElementMergeInfo) GetTargetElementStatus() string {
	if o == nil || o.TargetElementStatus == nil {
		var ret string
		return ret
	}
	return *o.TargetElementStatus
}

// GetTargetElementStatusOk returns a tuple with the TargetElementStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTElementMergeInfo) GetTargetElementStatusOk() (*string, bool) {
	if o == nil || o.TargetElementStatus == nil {
		return nil, false
	}
	return o.TargetElementStatus, true
}

// HasTargetElementStatus returns a boolean if a field has been set.
func (o *BTElementMergeInfo) HasTargetElementStatus() bool {
	if o != nil && o.TargetElementStatus != nil {
		return true
	}

	return false
}

// SetTargetElementStatus gets a reference to the given string and assigns it to the TargetElementStatus field.
func (o *BTElementMergeInfo) SetTargetElementStatus(v string) {
	o.TargetElementStatus = &v
}

// GetTargetModifiedAt returns the TargetModifiedAt field value if set, zero value otherwise.
func (o *BTElementMergeInfo) GetTargetModifiedAt() JSONTime {
	if o == nil || o.TargetModifiedAt == nil {
		var ret JSONTime
		return ret
	}
	return *o.TargetModifiedAt
}

// GetTargetModifiedAtOk returns a tuple with the TargetModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTElementMergeInfo) GetTargetModifiedAtOk() (*JSONTime, bool) {
	if o == nil || o.TargetModifiedAt == nil {
		return nil, false
	}
	return o.TargetModifiedAt, true
}

// HasTargetModifiedAt returns a boolean if a field has been set.
func (o *BTElementMergeInfo) HasTargetModifiedAt() bool {
	if o != nil && o.TargetModifiedAt != nil {
		return true
	}

	return false
}

// SetTargetModifiedAt gets a reference to the given JSONTime and assigns it to the TargetModifiedAt field.
func (o *BTElementMergeInfo) SetTargetModifiedAt(v JSONTime) {
	o.TargetModifiedAt = &v
}

// GetTargetModifiedBy returns the TargetModifiedBy field value if set, zero value otherwise.
func (o *BTElementMergeInfo) GetTargetModifiedBy() BTUserBasicSummaryInfo {
	if o == nil || o.TargetModifiedBy == nil {
		var ret BTUserBasicSummaryInfo
		return ret
	}
	return *o.TargetModifiedBy
}

// GetTargetModifiedByOk returns a tuple with the TargetModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTElementMergeInfo) GetTargetModifiedByOk() (*BTUserBasicSummaryInfo, bool) {
	if o == nil || o.TargetModifiedBy == nil {
		return nil, false
	}
	return o.TargetModifiedBy, true
}

// HasTargetModifiedBy returns a boolean if a field has been set.
func (o *BTElementMergeInfo) HasTargetModifiedBy() bool {
	if o != nil && o.TargetModifiedBy != nil {
		return true
	}

	return false
}

// SetTargetModifiedBy gets a reference to the given BTUserBasicSummaryInfo and assigns it to the TargetModifiedBy field.
func (o *BTElementMergeInfo) SetTargetModifiedBy(v BTUserBasicSummaryInfo) {
	o.TargetModifiedBy = &v
}

func (o BTElementMergeInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BranchPointElementName != nil {
		toSerialize["branchPointElementName"] = o.BranchPointElementName
	}
	if o.BranchPointElementPath != nil {
		toSerialize["branchPointElementPath"] = o.BranchPointElementPath
	}
	if o.DependentElementMergeInfo != nil {
		toSerialize["dependentElementMergeInfo"] = o.DependentElementMergeInfo
	}
	if o.ElementDataType != nil {
		toSerialize["elementDataType"] = o.ElementDataType
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.ElementType != nil {
		toSerialize["elementType"] = o.ElementType
	}
	if o.Mergeable != nil {
		toSerialize["mergeable"] = o.Mergeable
	}
	if o.SourceElementName != nil {
		toSerialize["sourceElementName"] = o.SourceElementName
	}
	if o.SourceElementPath != nil {
		toSerialize["sourceElementPath"] = o.SourceElementPath
	}
	if o.SourceElementStatus != nil {
		toSerialize["sourceElementStatus"] = o.SourceElementStatus
	}
	if o.SourceModifiedAt != nil {
		toSerialize["sourceModifiedAt"] = o.SourceModifiedAt
	}
	if o.SourceModifiedBy != nil {
		toSerialize["sourceModifiedBy"] = o.SourceModifiedBy
	}
	if o.TargetElementName != nil {
		toSerialize["targetElementName"] = o.TargetElementName
	}
	if o.TargetElementPath != nil {
		toSerialize["targetElementPath"] = o.TargetElementPath
	}
	if o.TargetElementStatus != nil {
		toSerialize["targetElementStatus"] = o.TargetElementStatus
	}
	if o.TargetModifiedAt != nil {
		toSerialize["targetModifiedAt"] = o.TargetModifiedAt
	}
	if o.TargetModifiedBy != nil {
		toSerialize["targetModifiedBy"] = o.TargetModifiedBy
	}
	return json.Marshal(toSerialize)
}

type NullableBTElementMergeInfo struct {
	value *BTElementMergeInfo
	isSet bool
}

func (v NullableBTElementMergeInfo) Get() *BTElementMergeInfo {
	return v.value
}

func (v *NullableBTElementMergeInfo) Set(val *BTElementMergeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTElementMergeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTElementMergeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTElementMergeInfo(val *BTElementMergeInfo) *NullableBTElementMergeInfo {
	return &NullableBTElementMergeInfo{value: val, isSet: true}
}

func (v NullableBTElementMergeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTElementMergeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
