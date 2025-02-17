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

// BTStringMinimumLengthPattern895 struct for BTStringMinimumLengthPattern895
type BTStringMinimumLengthPattern895 struct {
	BtType                        *string `json:"btType,omitempty"`
	ErrorMessage                  *string `json:"errorMessage,omitempty"`
	ShouldResetValueWhenConfirmed *bool   `json:"shouldResetValueWhenConfirmed,omitempty"`
	MinimumLength                 *int32  `json:"minimumLength,omitempty"`
}

// NewBTStringMinimumLengthPattern895 instantiates a new BTStringMinimumLengthPattern895 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTStringMinimumLengthPattern895() *BTStringMinimumLengthPattern895 {
	this := BTStringMinimumLengthPattern895{}
	return &this
}

// NewBTStringMinimumLengthPattern895WithDefaults instantiates a new BTStringMinimumLengthPattern895 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTStringMinimumLengthPattern895WithDefaults() *BTStringMinimumLengthPattern895 {
	this := BTStringMinimumLengthPattern895{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTStringMinimumLengthPattern895) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTStringMinimumLengthPattern895) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTStringMinimumLengthPattern895) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTStringMinimumLengthPattern895) SetBtType(v string) {
	o.BtType = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *BTStringMinimumLengthPattern895) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTStringMinimumLengthPattern895) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *BTStringMinimumLengthPattern895) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage != nil {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *BTStringMinimumLengthPattern895) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetShouldResetValueWhenConfirmed returns the ShouldResetValueWhenConfirmed field value if set, zero value otherwise.
func (o *BTStringMinimumLengthPattern895) GetShouldResetValueWhenConfirmed() bool {
	if o == nil || o.ShouldResetValueWhenConfirmed == nil {
		var ret bool
		return ret
	}
	return *o.ShouldResetValueWhenConfirmed
}

// GetShouldResetValueWhenConfirmedOk returns a tuple with the ShouldResetValueWhenConfirmed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTStringMinimumLengthPattern895) GetShouldResetValueWhenConfirmedOk() (*bool, bool) {
	if o == nil || o.ShouldResetValueWhenConfirmed == nil {
		return nil, false
	}
	return o.ShouldResetValueWhenConfirmed, true
}

// HasShouldResetValueWhenConfirmed returns a boolean if a field has been set.
func (o *BTStringMinimumLengthPattern895) HasShouldResetValueWhenConfirmed() bool {
	if o != nil && o.ShouldResetValueWhenConfirmed != nil {
		return true
	}

	return false
}

// SetShouldResetValueWhenConfirmed gets a reference to the given bool and assigns it to the ShouldResetValueWhenConfirmed field.
func (o *BTStringMinimumLengthPattern895) SetShouldResetValueWhenConfirmed(v bool) {
	o.ShouldResetValueWhenConfirmed = &v
}

// GetMinimumLength returns the MinimumLength field value if set, zero value otherwise.
func (o *BTStringMinimumLengthPattern895) GetMinimumLength() int32 {
	if o == nil || o.MinimumLength == nil {
		var ret int32
		return ret
	}
	return *o.MinimumLength
}

// GetMinimumLengthOk returns a tuple with the MinimumLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTStringMinimumLengthPattern895) GetMinimumLengthOk() (*int32, bool) {
	if o == nil || o.MinimumLength == nil {
		return nil, false
	}
	return o.MinimumLength, true
}

// HasMinimumLength returns a boolean if a field has been set.
func (o *BTStringMinimumLengthPattern895) HasMinimumLength() bool {
	if o != nil && o.MinimumLength != nil {
		return true
	}

	return false
}

// SetMinimumLength gets a reference to the given int32 and assigns it to the MinimumLength field.
func (o *BTStringMinimumLengthPattern895) SetMinimumLength(v int32) {
	o.MinimumLength = &v
}

func (o BTStringMinimumLengthPattern895) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ErrorMessage != nil {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if o.ShouldResetValueWhenConfirmed != nil {
		toSerialize["shouldResetValueWhenConfirmed"] = o.ShouldResetValueWhenConfirmed
	}
	if o.MinimumLength != nil {
		toSerialize["minimumLength"] = o.MinimumLength
	}
	return json.Marshal(toSerialize)
}

type NullableBTStringMinimumLengthPattern895 struct {
	value *BTStringMinimumLengthPattern895
	isSet bool
}

func (v NullableBTStringMinimumLengthPattern895) Get() *BTStringMinimumLengthPattern895 {
	return v.value
}

func (v *NullableBTStringMinimumLengthPattern895) Set(val *BTStringMinimumLengthPattern895) {
	v.value = val
	v.isSet = true
}

func (v NullableBTStringMinimumLengthPattern895) IsSet() bool {
	return v.isSet
}

func (v *NullableBTStringMinimumLengthPattern895) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTStringMinimumLengthPattern895(val *BTStringMinimumLengthPattern895) *NullableBTStringMinimumLengthPattern895 {
	return &NullableBTStringMinimumLengthPattern895{value: val, isSet: true}
}

func (v NullableBTStringMinimumLengthPattern895) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTStringMinimumLengthPattern895) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
