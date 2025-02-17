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

// BTTableCellParameterWithValue2122 struct for BTTableCellParameterWithValue2122
type BTTableCellParameterWithValue2122 struct {
	BtType        *string           `json:"btType,omitempty"`
	IsEverVisible *bool             `json:"isEverVisible,omitempty"`
	IsReadOnly    *bool             `json:"isReadOnly,omitempty"`
	Error         *string           `json:"error,omitempty"`
	OverrideSpec  *BTParameterSpec6 `json:"overrideSpec,omitempty"`
	Parameter     *BTMParameter1    `json:"parameter,omitempty"`
	Value         *BTFSValue1888    `json:"value,omitempty"`
}

// NewBTTableCellParameterWithValue2122 instantiates a new BTTableCellParameterWithValue2122 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTTableCellParameterWithValue2122() *BTTableCellParameterWithValue2122 {
	this := BTTableCellParameterWithValue2122{}
	return &this
}

// NewBTTableCellParameterWithValue2122WithDefaults instantiates a new BTTableCellParameterWithValue2122 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTTableCellParameterWithValue2122WithDefaults() *BTTableCellParameterWithValue2122 {
	this := BTTableCellParameterWithValue2122{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTTableCellParameterWithValue2122) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableCellParameterWithValue2122) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTTableCellParameterWithValue2122) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTTableCellParameterWithValue2122) SetBtType(v string) {
	o.BtType = &v
}

// GetIsEverVisible returns the IsEverVisible field value if set, zero value otherwise.
func (o *BTTableCellParameterWithValue2122) GetIsEverVisible() bool {
	if o == nil || o.IsEverVisible == nil {
		var ret bool
		return ret
	}
	return *o.IsEverVisible
}

// GetIsEverVisibleOk returns a tuple with the IsEverVisible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableCellParameterWithValue2122) GetIsEverVisibleOk() (*bool, bool) {
	if o == nil || o.IsEverVisible == nil {
		return nil, false
	}
	return o.IsEverVisible, true
}

// HasIsEverVisible returns a boolean if a field has been set.
func (o *BTTableCellParameterWithValue2122) HasIsEverVisible() bool {
	if o != nil && o.IsEverVisible != nil {
		return true
	}

	return false
}

// SetIsEverVisible gets a reference to the given bool and assigns it to the IsEverVisible field.
func (o *BTTableCellParameterWithValue2122) SetIsEverVisible(v bool) {
	o.IsEverVisible = &v
}

// GetIsReadOnly returns the IsReadOnly field value if set, zero value otherwise.
func (o *BTTableCellParameterWithValue2122) GetIsReadOnly() bool {
	if o == nil || o.IsReadOnly == nil {
		var ret bool
		return ret
	}
	return *o.IsReadOnly
}

// GetIsReadOnlyOk returns a tuple with the IsReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableCellParameterWithValue2122) GetIsReadOnlyOk() (*bool, bool) {
	if o == nil || o.IsReadOnly == nil {
		return nil, false
	}
	return o.IsReadOnly, true
}

// HasIsReadOnly returns a boolean if a field has been set.
func (o *BTTableCellParameterWithValue2122) HasIsReadOnly() bool {
	if o != nil && o.IsReadOnly != nil {
		return true
	}

	return false
}

// SetIsReadOnly gets a reference to the given bool and assigns it to the IsReadOnly field.
func (o *BTTableCellParameterWithValue2122) SetIsReadOnly(v bool) {
	o.IsReadOnly = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *BTTableCellParameterWithValue2122) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableCellParameterWithValue2122) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *BTTableCellParameterWithValue2122) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *BTTableCellParameterWithValue2122) SetError(v string) {
	o.Error = &v
}

// GetOverrideSpec returns the OverrideSpec field value if set, zero value otherwise.
func (o *BTTableCellParameterWithValue2122) GetOverrideSpec() BTParameterSpec6 {
	if o == nil || o.OverrideSpec == nil {
		var ret BTParameterSpec6
		return ret
	}
	return *o.OverrideSpec
}

// GetOverrideSpecOk returns a tuple with the OverrideSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableCellParameterWithValue2122) GetOverrideSpecOk() (*BTParameterSpec6, bool) {
	if o == nil || o.OverrideSpec == nil {
		return nil, false
	}
	return o.OverrideSpec, true
}

// HasOverrideSpec returns a boolean if a field has been set.
func (o *BTTableCellParameterWithValue2122) HasOverrideSpec() bool {
	if o != nil && o.OverrideSpec != nil {
		return true
	}

	return false
}

// SetOverrideSpec gets a reference to the given BTParameterSpec6 and assigns it to the OverrideSpec field.
func (o *BTTableCellParameterWithValue2122) SetOverrideSpec(v BTParameterSpec6) {
	o.OverrideSpec = &v
}

// GetParameter returns the Parameter field value if set, zero value otherwise.
func (o *BTTableCellParameterWithValue2122) GetParameter() BTMParameter1 {
	if o == nil || o.Parameter == nil {
		var ret BTMParameter1
		return ret
	}
	return *o.Parameter
}

// GetParameterOk returns a tuple with the Parameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableCellParameterWithValue2122) GetParameterOk() (*BTMParameter1, bool) {
	if o == nil || o.Parameter == nil {
		return nil, false
	}
	return o.Parameter, true
}

// HasParameter returns a boolean if a field has been set.
func (o *BTTableCellParameterWithValue2122) HasParameter() bool {
	if o != nil && o.Parameter != nil {
		return true
	}

	return false
}

// SetParameter gets a reference to the given BTMParameter1 and assigns it to the Parameter field.
func (o *BTTableCellParameterWithValue2122) SetParameter(v BTMParameter1) {
	o.Parameter = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *BTTableCellParameterWithValue2122) GetValue() BTFSValue1888 {
	if o == nil || o.Value == nil {
		var ret BTFSValue1888
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableCellParameterWithValue2122) GetValueOk() (*BTFSValue1888, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *BTTableCellParameterWithValue2122) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given BTFSValue1888 and assigns it to the Value field.
func (o *BTTableCellParameterWithValue2122) SetValue(v BTFSValue1888) {
	o.Value = &v
}

func (o BTTableCellParameterWithValue2122) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.IsEverVisible != nil {
		toSerialize["isEverVisible"] = o.IsEverVisible
	}
	if o.IsReadOnly != nil {
		toSerialize["isReadOnly"] = o.IsReadOnly
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.OverrideSpec != nil {
		toSerialize["overrideSpec"] = o.OverrideSpec
	}
	if o.Parameter != nil {
		toSerialize["parameter"] = o.Parameter
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableBTTableCellParameterWithValue2122 struct {
	value *BTTableCellParameterWithValue2122
	isSet bool
}

func (v NullableBTTableCellParameterWithValue2122) Get() *BTTableCellParameterWithValue2122 {
	return v.value
}

func (v *NullableBTTableCellParameterWithValue2122) Set(val *BTTableCellParameterWithValue2122) {
	v.value = val
	v.isSet = true
}

func (v NullableBTTableCellParameterWithValue2122) IsSet() bool {
	return v.isSet
}

func (v *NullableBTTableCellParameterWithValue2122) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTTableCellParameterWithValue2122(val *BTTableCellParameterWithValue2122) *NullableBTTableCellParameterWithValue2122 {
	return &NullableBTTableCellParameterWithValue2122{value: val, isSet: true}
}

func (v NullableBTTableCellParameterWithValue2122) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTTableCellParameterWithValue2122) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
