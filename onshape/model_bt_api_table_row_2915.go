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

// BTApiTableRow2915 struct for BTApiTableRow2915
type BTApiTableRow2915 struct {
	BtType          *string            `json:"btType,omitempty"`
	Callout         *string            `json:"callout,omitempty"`
	ColumnIdToValue *map[string]string `json:"columnIdToValue,omitempty"`
	EntityIds       []string           `json:"entityIds,omitempty"`
}

// NewBTApiTableRow2915 instantiates a new BTApiTableRow2915 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTApiTableRow2915() *BTApiTableRow2915 {
	this := BTApiTableRow2915{}
	return &this
}

// NewBTApiTableRow2915WithDefaults instantiates a new BTApiTableRow2915 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTApiTableRow2915WithDefaults() *BTApiTableRow2915 {
	this := BTApiTableRow2915{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTApiTableRow2915) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTApiTableRow2915) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTApiTableRow2915) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTApiTableRow2915) SetBtType(v string) {
	o.BtType = &v
}

// GetCallout returns the Callout field value if set, zero value otherwise.
func (o *BTApiTableRow2915) GetCallout() string {
	if o == nil || o.Callout == nil {
		var ret string
		return ret
	}
	return *o.Callout
}

// GetCalloutOk returns a tuple with the Callout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTApiTableRow2915) GetCalloutOk() (*string, bool) {
	if o == nil || o.Callout == nil {
		return nil, false
	}
	return o.Callout, true
}

// HasCallout returns a boolean if a field has been set.
func (o *BTApiTableRow2915) HasCallout() bool {
	if o != nil && o.Callout != nil {
		return true
	}

	return false
}

// SetCallout gets a reference to the given string and assigns it to the Callout field.
func (o *BTApiTableRow2915) SetCallout(v string) {
	o.Callout = &v
}

// GetColumnIdToValue returns the ColumnIdToValue field value if set, zero value otherwise.
func (o *BTApiTableRow2915) GetColumnIdToValue() map[string]string {
	if o == nil || o.ColumnIdToValue == nil {
		var ret map[string]string
		return ret
	}
	return *o.ColumnIdToValue
}

// GetColumnIdToValueOk returns a tuple with the ColumnIdToValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTApiTableRow2915) GetColumnIdToValueOk() (*map[string]string, bool) {
	if o == nil || o.ColumnIdToValue == nil {
		return nil, false
	}
	return o.ColumnIdToValue, true
}

// HasColumnIdToValue returns a boolean if a field has been set.
func (o *BTApiTableRow2915) HasColumnIdToValue() bool {
	if o != nil && o.ColumnIdToValue != nil {
		return true
	}

	return false
}

// SetColumnIdToValue gets a reference to the given map[string]string and assigns it to the ColumnIdToValue field.
func (o *BTApiTableRow2915) SetColumnIdToValue(v map[string]string) {
	o.ColumnIdToValue = &v
}

// GetEntityIds returns the EntityIds field value if set, zero value otherwise.
func (o *BTApiTableRow2915) GetEntityIds() []string {
	if o == nil || o.EntityIds == nil {
		var ret []string
		return ret
	}
	return o.EntityIds
}

// GetEntityIdsOk returns a tuple with the EntityIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTApiTableRow2915) GetEntityIdsOk() ([]string, bool) {
	if o == nil || o.EntityIds == nil {
		return nil, false
	}
	return o.EntityIds, true
}

// HasEntityIds returns a boolean if a field has been set.
func (o *BTApiTableRow2915) HasEntityIds() bool {
	if o != nil && o.EntityIds != nil {
		return true
	}

	return false
}

// SetEntityIds gets a reference to the given []string and assigns it to the EntityIds field.
func (o *BTApiTableRow2915) SetEntityIds(v []string) {
	o.EntityIds = v
}

func (o BTApiTableRow2915) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Callout != nil {
		toSerialize["callout"] = o.Callout
	}
	if o.ColumnIdToValue != nil {
		toSerialize["columnIdToValue"] = o.ColumnIdToValue
	}
	if o.EntityIds != nil {
		toSerialize["entityIds"] = o.EntityIds
	}
	return json.Marshal(toSerialize)
}

type NullableBTApiTableRow2915 struct {
	value *BTApiTableRow2915
	isSet bool
}

func (v NullableBTApiTableRow2915) Get() *BTApiTableRow2915 {
	return v.value
}

func (v *NullableBTApiTableRow2915) Set(val *BTApiTableRow2915) {
	v.value = val
	v.isSet = true
}

func (v NullableBTApiTableRow2915) IsSet() bool {
	return v.isSet
}

func (v *NullableBTApiTableRow2915) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTApiTableRow2915(val *BTApiTableRow2915) *NullableBTApiTableRow2915 {
	return &NullableBTApiTableRow2915{value: val, isSet: true}
}

func (v NullableBTApiTableRow2915) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTApiTableRow2915) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
