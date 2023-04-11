/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.161.14306-351f5b17f026
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTAppElementModifyInfo struct for BTAppElementModifyInfo
type BTAppElementModifyInfo struct {
	// The latest change id for the element, after the edit was committed.
	ChangeId string `json:"changeId"`
	// The id of the edited element.
	ElementId *string `json:"elementId,omitempty"`
	// The ids of the edited elements, if multiple elements were edited.
	ElementIds []string `json:"elementIds,omitempty"`
	// The numeric code identifying the error that occurred, if one occurred.
	ErrorCode *int32 `json:"errorCode,omitempty"`
	// A human-readable value for the error that occurred, if one occurred.
	ErrorDescription *string                 `json:"errorDescription,omitempty"`
	ErrorValue       *string                 `json:"errorValue,omitempty"`
	JsonDifference   *BTDiffJsonResponse2725 `json:"jsonDifference,omitempty"`
	// The latest change id for the element, before the edit was made.
	ParentChangeId *string `json:"parentChangeId,omitempty"`
	// When committing a transaction, this field indicates if the properties of the application element were changed after the transaction was created.
	PropertyEditsMerged *bool `json:"propertyEditsMerged,omitempty"`
	// The id of the transaction in which the edit was applied.
	TransactionId *string `json:"transactionId,omitempty"`
}

// NewBTAppElementModifyInfo instantiates a new BTAppElementModifyInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAppElementModifyInfo(changeId string) *BTAppElementModifyInfo {
	this := BTAppElementModifyInfo{}
	this.ChangeId = changeId
	return &this
}

// NewBTAppElementModifyInfoWithDefaults instantiates a new BTAppElementModifyInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAppElementModifyInfoWithDefaults() *BTAppElementModifyInfo {
	this := BTAppElementModifyInfo{}
	return &this
}

// GetChangeId returns the ChangeId field value
func (o *BTAppElementModifyInfo) GetChangeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChangeId
}

// GetChangeIdOk returns a tuple with the ChangeId field value
// and a boolean to check if the value has been set.
func (o *BTAppElementModifyInfo) GetChangeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChangeId, true
}

// SetChangeId sets field value
func (o *BTAppElementModifyInfo) SetChangeId(v string) {
	o.ChangeId = v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *BTAppElementModifyInfo) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementModifyInfo) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *BTAppElementModifyInfo) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *BTAppElementModifyInfo) SetElementId(v string) {
	o.ElementId = &v
}

// GetElementIds returns the ElementIds field value if set, zero value otherwise.
func (o *BTAppElementModifyInfo) GetElementIds() []string {
	if o == nil || o.ElementIds == nil {
		var ret []string
		return ret
	}
	return o.ElementIds
}

// GetElementIdsOk returns a tuple with the ElementIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementModifyInfo) GetElementIdsOk() ([]string, bool) {
	if o == nil || o.ElementIds == nil {
		return nil, false
	}
	return o.ElementIds, true
}

// HasElementIds returns a boolean if a field has been set.
func (o *BTAppElementModifyInfo) HasElementIds() bool {
	if o != nil && o.ElementIds != nil {
		return true
	}

	return false
}

// SetElementIds gets a reference to the given []string and assigns it to the ElementIds field.
func (o *BTAppElementModifyInfo) SetElementIds(v []string) {
	o.ElementIds = v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *BTAppElementModifyInfo) GetErrorCode() int32 {
	if o == nil || o.ErrorCode == nil {
		var ret int32
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementModifyInfo) GetErrorCodeOk() (*int32, bool) {
	if o == nil || o.ErrorCode == nil {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *BTAppElementModifyInfo) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given int32 and assigns it to the ErrorCode field.
func (o *BTAppElementModifyInfo) SetErrorCode(v int32) {
	o.ErrorCode = &v
}

// GetErrorDescription returns the ErrorDescription field value if set, zero value otherwise.
func (o *BTAppElementModifyInfo) GetErrorDescription() string {
	if o == nil || o.ErrorDescription == nil {
		var ret string
		return ret
	}
	return *o.ErrorDescription
}

// GetErrorDescriptionOk returns a tuple with the ErrorDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementModifyInfo) GetErrorDescriptionOk() (*string, bool) {
	if o == nil || o.ErrorDescription == nil {
		return nil, false
	}
	return o.ErrorDescription, true
}

// HasErrorDescription returns a boolean if a field has been set.
func (o *BTAppElementModifyInfo) HasErrorDescription() bool {
	if o != nil && o.ErrorDescription != nil {
		return true
	}

	return false
}

// SetErrorDescription gets a reference to the given string and assigns it to the ErrorDescription field.
func (o *BTAppElementModifyInfo) SetErrorDescription(v string) {
	o.ErrorDescription = &v
}

// GetErrorValue returns the ErrorValue field value if set, zero value otherwise.
func (o *BTAppElementModifyInfo) GetErrorValue() string {
	if o == nil || o.ErrorValue == nil {
		var ret string
		return ret
	}
	return *o.ErrorValue
}

// GetErrorValueOk returns a tuple with the ErrorValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementModifyInfo) GetErrorValueOk() (*string, bool) {
	if o == nil || o.ErrorValue == nil {
		return nil, false
	}
	return o.ErrorValue, true
}

// HasErrorValue returns a boolean if a field has been set.
func (o *BTAppElementModifyInfo) HasErrorValue() bool {
	if o != nil && o.ErrorValue != nil {
		return true
	}

	return false
}

// SetErrorValue gets a reference to the given string and assigns it to the ErrorValue field.
func (o *BTAppElementModifyInfo) SetErrorValue(v string) {
	o.ErrorValue = &v
}

// GetJsonDifference returns the JsonDifference field value if set, zero value otherwise.
func (o *BTAppElementModifyInfo) GetJsonDifference() BTDiffJsonResponse2725 {
	if o == nil || o.JsonDifference == nil {
		var ret BTDiffJsonResponse2725
		return ret
	}
	return *o.JsonDifference
}

// GetJsonDifferenceOk returns a tuple with the JsonDifference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementModifyInfo) GetJsonDifferenceOk() (*BTDiffJsonResponse2725, bool) {
	if o == nil || o.JsonDifference == nil {
		return nil, false
	}
	return o.JsonDifference, true
}

// HasJsonDifference returns a boolean if a field has been set.
func (o *BTAppElementModifyInfo) HasJsonDifference() bool {
	if o != nil && o.JsonDifference != nil {
		return true
	}

	return false
}

// SetJsonDifference gets a reference to the given BTDiffJsonResponse2725 and assigns it to the JsonDifference field.
func (o *BTAppElementModifyInfo) SetJsonDifference(v BTDiffJsonResponse2725) {
	o.JsonDifference = &v
}

// GetParentChangeId returns the ParentChangeId field value if set, zero value otherwise.
func (o *BTAppElementModifyInfo) GetParentChangeId() string {
	if o == nil || o.ParentChangeId == nil {
		var ret string
		return ret
	}
	return *o.ParentChangeId
}

// GetParentChangeIdOk returns a tuple with the ParentChangeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementModifyInfo) GetParentChangeIdOk() (*string, bool) {
	if o == nil || o.ParentChangeId == nil {
		return nil, false
	}
	return o.ParentChangeId, true
}

// HasParentChangeId returns a boolean if a field has been set.
func (o *BTAppElementModifyInfo) HasParentChangeId() bool {
	if o != nil && o.ParentChangeId != nil {
		return true
	}

	return false
}

// SetParentChangeId gets a reference to the given string and assigns it to the ParentChangeId field.
func (o *BTAppElementModifyInfo) SetParentChangeId(v string) {
	o.ParentChangeId = &v
}

// GetPropertyEditsMerged returns the PropertyEditsMerged field value if set, zero value otherwise.
func (o *BTAppElementModifyInfo) GetPropertyEditsMerged() bool {
	if o == nil || o.PropertyEditsMerged == nil {
		var ret bool
		return ret
	}
	return *o.PropertyEditsMerged
}

// GetPropertyEditsMergedOk returns a tuple with the PropertyEditsMerged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementModifyInfo) GetPropertyEditsMergedOk() (*bool, bool) {
	if o == nil || o.PropertyEditsMerged == nil {
		return nil, false
	}
	return o.PropertyEditsMerged, true
}

// HasPropertyEditsMerged returns a boolean if a field has been set.
func (o *BTAppElementModifyInfo) HasPropertyEditsMerged() bool {
	if o != nil && o.PropertyEditsMerged != nil {
		return true
	}

	return false
}

// SetPropertyEditsMerged gets a reference to the given bool and assigns it to the PropertyEditsMerged field.
func (o *BTAppElementModifyInfo) SetPropertyEditsMerged(v bool) {
	o.PropertyEditsMerged = &v
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *BTAppElementModifyInfo) GetTransactionId() string {
	if o == nil || o.TransactionId == nil {
		var ret string
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementModifyInfo) GetTransactionIdOk() (*string, bool) {
	if o == nil || o.TransactionId == nil {
		return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *BTAppElementModifyInfo) HasTransactionId() bool {
	if o != nil && o.TransactionId != nil {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given string and assigns it to the TransactionId field.
func (o *BTAppElementModifyInfo) SetTransactionId(v string) {
	o.TransactionId = &v
}

func (o BTAppElementModifyInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["changeId"] = o.ChangeId
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.ElementIds != nil {
		toSerialize["elementIds"] = o.ElementIds
	}
	if o.ErrorCode != nil {
		toSerialize["errorCode"] = o.ErrorCode
	}
	if o.ErrorDescription != nil {
		toSerialize["errorDescription"] = o.ErrorDescription
	}
	if o.ErrorValue != nil {
		toSerialize["errorValue"] = o.ErrorValue
	}
	if o.JsonDifference != nil {
		toSerialize["jsonDifference"] = o.JsonDifference
	}
	if o.ParentChangeId != nil {
		toSerialize["parentChangeId"] = o.ParentChangeId
	}
	if o.PropertyEditsMerged != nil {
		toSerialize["propertyEditsMerged"] = o.PropertyEditsMerged
	}
	if o.TransactionId != nil {
		toSerialize["transactionId"] = o.TransactionId
	}
	return json.Marshal(toSerialize)
}

type NullableBTAppElementModifyInfo struct {
	value *BTAppElementModifyInfo
	isSet bool
}

func (v NullableBTAppElementModifyInfo) Get() *BTAppElementModifyInfo {
	return v.value
}

func (v *NullableBTAppElementModifyInfo) Set(val *BTAppElementModifyInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAppElementModifyInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAppElementModifyInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAppElementModifyInfo(val *BTAppElementModifyInfo) *NullableBTAppElementModifyInfo {
	return &NullableBTAppElementModifyInfo{value: val, isSet: true}
}

func (v NullableBTAppElementModifyInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAppElementModifyInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
