/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.161.13200-ff216a970a02
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTAppElementCommitTransactionParams struct for BTAppElementCommitTransactionParams
type BTAppElementCommitTransactionParams struct {
	AllowMerge     *bool    `json:"allowMerge,omitempty"`
	Description    *string  `json:"description,omitempty"`
	ReturnError    *bool    `json:"returnError,omitempty"`
	TransactionIds []string `json:"transactionIds,omitempty"`
}

// NewBTAppElementCommitTransactionParams instantiates a new BTAppElementCommitTransactionParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAppElementCommitTransactionParams() *BTAppElementCommitTransactionParams {
	this := BTAppElementCommitTransactionParams{}
	return &this
}

// NewBTAppElementCommitTransactionParamsWithDefaults instantiates a new BTAppElementCommitTransactionParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAppElementCommitTransactionParamsWithDefaults() *BTAppElementCommitTransactionParams {
	this := BTAppElementCommitTransactionParams{}
	return &this
}

// GetAllowMerge returns the AllowMerge field value if set, zero value otherwise.
func (o *BTAppElementCommitTransactionParams) GetAllowMerge() bool {
	if o == nil || o.AllowMerge == nil {
		var ret bool
		return ret
	}
	return *o.AllowMerge
}

// GetAllowMergeOk returns a tuple with the AllowMerge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementCommitTransactionParams) GetAllowMergeOk() (*bool, bool) {
	if o == nil || o.AllowMerge == nil {
		return nil, false
	}
	return o.AllowMerge, true
}

// HasAllowMerge returns a boolean if a field has been set.
func (o *BTAppElementCommitTransactionParams) HasAllowMerge() bool {
	if o != nil && o.AllowMerge != nil {
		return true
	}

	return false
}

// SetAllowMerge gets a reference to the given bool and assigns it to the AllowMerge field.
func (o *BTAppElementCommitTransactionParams) SetAllowMerge(v bool) {
	o.AllowMerge = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BTAppElementCommitTransactionParams) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementCommitTransactionParams) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BTAppElementCommitTransactionParams) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BTAppElementCommitTransactionParams) SetDescription(v string) {
	o.Description = &v
}

// GetReturnError returns the ReturnError field value if set, zero value otherwise.
func (o *BTAppElementCommitTransactionParams) GetReturnError() bool {
	if o == nil || o.ReturnError == nil {
		var ret bool
		return ret
	}
	return *o.ReturnError
}

// GetReturnErrorOk returns a tuple with the ReturnError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementCommitTransactionParams) GetReturnErrorOk() (*bool, bool) {
	if o == nil || o.ReturnError == nil {
		return nil, false
	}
	return o.ReturnError, true
}

// HasReturnError returns a boolean if a field has been set.
func (o *BTAppElementCommitTransactionParams) HasReturnError() bool {
	if o != nil && o.ReturnError != nil {
		return true
	}

	return false
}

// SetReturnError gets a reference to the given bool and assigns it to the ReturnError field.
func (o *BTAppElementCommitTransactionParams) SetReturnError(v bool) {
	o.ReturnError = &v
}

// GetTransactionIds returns the TransactionIds field value if set, zero value otherwise.
func (o *BTAppElementCommitTransactionParams) GetTransactionIds() []string {
	if o == nil || o.TransactionIds == nil {
		var ret []string
		return ret
	}
	return o.TransactionIds
}

// GetTransactionIdsOk returns a tuple with the TransactionIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementCommitTransactionParams) GetTransactionIdsOk() ([]string, bool) {
	if o == nil || o.TransactionIds == nil {
		return nil, false
	}
	return o.TransactionIds, true
}

// HasTransactionIds returns a boolean if a field has been set.
func (o *BTAppElementCommitTransactionParams) HasTransactionIds() bool {
	if o != nil && o.TransactionIds != nil {
		return true
	}

	return false
}

// SetTransactionIds gets a reference to the given []string and assigns it to the TransactionIds field.
func (o *BTAppElementCommitTransactionParams) SetTransactionIds(v []string) {
	o.TransactionIds = v
}

func (o BTAppElementCommitTransactionParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowMerge != nil {
		toSerialize["allowMerge"] = o.AllowMerge
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ReturnError != nil {
		toSerialize["returnError"] = o.ReturnError
	}
	if o.TransactionIds != nil {
		toSerialize["transactionIds"] = o.TransactionIds
	}
	return json.Marshal(toSerialize)
}

type NullableBTAppElementCommitTransactionParams struct {
	value *BTAppElementCommitTransactionParams
	isSet bool
}

func (v NullableBTAppElementCommitTransactionParams) Get() *BTAppElementCommitTransactionParams {
	return v.value
}

func (v *NullableBTAppElementCommitTransactionParams) Set(val *BTAppElementCommitTransactionParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAppElementCommitTransactionParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAppElementCommitTransactionParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAppElementCommitTransactionParams(val *BTAppElementCommitTransactionParams) *NullableBTAppElementCommitTransactionParams {
	return &NullableBTAppElementCommitTransactionParams{value: val, isSet: true}
}

func (v NullableBTAppElementCommitTransactionParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAppElementCommitTransactionParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
