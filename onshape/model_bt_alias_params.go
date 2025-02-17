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

// BTAliasParams struct for BTAliasParams
type BTAliasParams struct {
	Additions   []BTAliasEntryParams `json:"additions,omitempty"`
	Description *string              `json:"description,omitempty"`
	Entries     []BTAliasEntryParams `json:"entries,omitempty"`
	Name        *string              `json:"name,omitempty"`
	Removals    []BTAliasEntryParams `json:"removals,omitempty"`
}

// NewBTAliasParams instantiates a new BTAliasParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAliasParams() *BTAliasParams {
	this := BTAliasParams{}
	return &this
}

// NewBTAliasParamsWithDefaults instantiates a new BTAliasParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAliasParamsWithDefaults() *BTAliasParams {
	this := BTAliasParams{}
	return &this
}

// GetAdditions returns the Additions field value if set, zero value otherwise.
func (o *BTAliasParams) GetAdditions() []BTAliasEntryParams {
	if o == nil || o.Additions == nil {
		var ret []BTAliasEntryParams
		return ret
	}
	return o.Additions
}

// GetAdditionsOk returns a tuple with the Additions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAliasParams) GetAdditionsOk() ([]BTAliasEntryParams, bool) {
	if o == nil || o.Additions == nil {
		return nil, false
	}
	return o.Additions, true
}

// HasAdditions returns a boolean if a field has been set.
func (o *BTAliasParams) HasAdditions() bool {
	if o != nil && o.Additions != nil {
		return true
	}

	return false
}

// SetAdditions gets a reference to the given []BTAliasEntryParams and assigns it to the Additions field.
func (o *BTAliasParams) SetAdditions(v []BTAliasEntryParams) {
	o.Additions = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BTAliasParams) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAliasParams) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BTAliasParams) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BTAliasParams) SetDescription(v string) {
	o.Description = &v
}

// GetEntries returns the Entries field value if set, zero value otherwise.
func (o *BTAliasParams) GetEntries() []BTAliasEntryParams {
	if o == nil || o.Entries == nil {
		var ret []BTAliasEntryParams
		return ret
	}
	return o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAliasParams) GetEntriesOk() ([]BTAliasEntryParams, bool) {
	if o == nil || o.Entries == nil {
		return nil, false
	}
	return o.Entries, true
}

// HasEntries returns a boolean if a field has been set.
func (o *BTAliasParams) HasEntries() bool {
	if o != nil && o.Entries != nil {
		return true
	}

	return false
}

// SetEntries gets a reference to the given []BTAliasEntryParams and assigns it to the Entries field.
func (o *BTAliasParams) SetEntries(v []BTAliasEntryParams) {
	o.Entries = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTAliasParams) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAliasParams) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTAliasParams) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTAliasParams) SetName(v string) {
	o.Name = &v
}

// GetRemovals returns the Removals field value if set, zero value otherwise.
func (o *BTAliasParams) GetRemovals() []BTAliasEntryParams {
	if o == nil || o.Removals == nil {
		var ret []BTAliasEntryParams
		return ret
	}
	return o.Removals
}

// GetRemovalsOk returns a tuple with the Removals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAliasParams) GetRemovalsOk() ([]BTAliasEntryParams, bool) {
	if o == nil || o.Removals == nil {
		return nil, false
	}
	return o.Removals, true
}

// HasRemovals returns a boolean if a field has been set.
func (o *BTAliasParams) HasRemovals() bool {
	if o != nil && o.Removals != nil {
		return true
	}

	return false
}

// SetRemovals gets a reference to the given []BTAliasEntryParams and assigns it to the Removals field.
func (o *BTAliasParams) SetRemovals(v []BTAliasEntryParams) {
	o.Removals = v
}

func (o BTAliasParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Additions != nil {
		toSerialize["additions"] = o.Additions
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Entries != nil {
		toSerialize["entries"] = o.Entries
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Removals != nil {
		toSerialize["removals"] = o.Removals
	}
	return json.Marshal(toSerialize)
}

type NullableBTAliasParams struct {
	value *BTAliasParams
	isSet bool
}

func (v NullableBTAliasParams) Get() *BTAliasParams {
	return v.value
}

func (v *NullableBTAliasParams) Set(val *BTAliasParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAliasParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAliasParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAliasParams(val *BTAliasParams) *NullableBTAliasParams {
	return &NullableBTAliasParams{value: val, isSet: true}
}

func (v NullableBTAliasParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAliasParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
