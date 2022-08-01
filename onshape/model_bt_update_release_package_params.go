/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.151.5812-3c52042ec720
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTUpdateReleasePackageParams struct for BTUpdateReleasePackageParams
type BTUpdateReleasePackageParams struct {
	Empty      *bool                        `json:"empty,omitempty"`
	ItemIds    []string                     `json:"itemIds,omitempty"`
	Items      []BTReleasePackageItemParams `json:"items,omitempty"`
	Properties []BTPropertyValueParam       `json:"properties,omitempty"`
}

// NewBTUpdateReleasePackageParams instantiates a new BTUpdateReleasePackageParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTUpdateReleasePackageParams() *BTUpdateReleasePackageParams {
	this := BTUpdateReleasePackageParams{}
	return &this
}

// NewBTUpdateReleasePackageParamsWithDefaults instantiates a new BTUpdateReleasePackageParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTUpdateReleasePackageParamsWithDefaults() *BTUpdateReleasePackageParams {
	this := BTUpdateReleasePackageParams{}
	return &this
}

// GetEmpty returns the Empty field value if set, zero value otherwise.
func (o *BTUpdateReleasePackageParams) GetEmpty() bool {
	if o == nil || o.Empty == nil {
		var ret bool
		return ret
	}
	return *o.Empty
}

// GetEmptyOk returns a tuple with the Empty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUpdateReleasePackageParams) GetEmptyOk() (*bool, bool) {
	if o == nil || o.Empty == nil {
		return nil, false
	}
	return o.Empty, true
}

// HasEmpty returns a boolean if a field has been set.
func (o *BTUpdateReleasePackageParams) HasEmpty() bool {
	if o != nil && o.Empty != nil {
		return true
	}

	return false
}

// SetEmpty gets a reference to the given bool and assigns it to the Empty field.
func (o *BTUpdateReleasePackageParams) SetEmpty(v bool) {
	o.Empty = &v
}

// GetItemIds returns the ItemIds field value if set, zero value otherwise.
func (o *BTUpdateReleasePackageParams) GetItemIds() []string {
	if o == nil || o.ItemIds == nil {
		var ret []string
		return ret
	}
	return o.ItemIds
}

// GetItemIdsOk returns a tuple with the ItemIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUpdateReleasePackageParams) GetItemIdsOk() ([]string, bool) {
	if o == nil || o.ItemIds == nil {
		return nil, false
	}
	return o.ItemIds, true
}

// HasItemIds returns a boolean if a field has been set.
func (o *BTUpdateReleasePackageParams) HasItemIds() bool {
	if o != nil && o.ItemIds != nil {
		return true
	}

	return false
}

// SetItemIds gets a reference to the given []string and assigns it to the ItemIds field.
func (o *BTUpdateReleasePackageParams) SetItemIds(v []string) {
	o.ItemIds = v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *BTUpdateReleasePackageParams) GetItems() []BTReleasePackageItemParams {
	if o == nil || o.Items == nil {
		var ret []BTReleasePackageItemParams
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUpdateReleasePackageParams) GetItemsOk() ([]BTReleasePackageItemParams, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *BTUpdateReleasePackageParams) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []BTReleasePackageItemParams and assigns it to the Items field.
func (o *BTUpdateReleasePackageParams) SetItems(v []BTReleasePackageItemParams) {
	o.Items = v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *BTUpdateReleasePackageParams) GetProperties() []BTPropertyValueParam {
	if o == nil || o.Properties == nil {
		var ret []BTPropertyValueParam
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUpdateReleasePackageParams) GetPropertiesOk() ([]BTPropertyValueParam, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *BTUpdateReleasePackageParams) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []BTPropertyValueParam and assigns it to the Properties field.
func (o *BTUpdateReleasePackageParams) SetProperties(v []BTPropertyValueParam) {
	o.Properties = v
}

func (o BTUpdateReleasePackageParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Empty != nil {
		toSerialize["empty"] = o.Empty
	}
	if o.ItemIds != nil {
		toSerialize["itemIds"] = o.ItemIds
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	return json.Marshal(toSerialize)
}

type NullableBTUpdateReleasePackageParams struct {
	value *BTUpdateReleasePackageParams
	isSet bool
}

func (v NullableBTUpdateReleasePackageParams) Get() *BTUpdateReleasePackageParams {
	return v.value
}

func (v *NullableBTUpdateReleasePackageParams) Set(val *BTUpdateReleasePackageParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTUpdateReleasePackageParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTUpdateReleasePackageParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTUpdateReleasePackageParams(val *BTUpdateReleasePackageParams) *NullableBTUpdateReleasePackageParams {
	return &NullableBTUpdateReleasePackageParams{value: val, isSet: true}
}

func (v NullableBTUpdateReleasePackageParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTUpdateReleasePackageParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
