/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.151.5809-984b9f882afe
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMetadataPropertyUpdateParams struct for BTMetadataPropertyUpdateParams
type BTMetadataPropertyUpdateParams struct {
	// The id of the property that should be edited. This can be retrieved from MetadataCategory:getCategoryProperties.
	PropertyId *string `json:"propertyId,omitempty"`
	// The new value for the property.
	Value *map[string]interface{} `json:"value,omitempty"`
}

// NewBTMetadataPropertyUpdateParams instantiates a new BTMetadataPropertyUpdateParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMetadataPropertyUpdateParams() *BTMetadataPropertyUpdateParams {
	this := BTMetadataPropertyUpdateParams{}
	return &this
}

// NewBTMetadataPropertyUpdateParamsWithDefaults instantiates a new BTMetadataPropertyUpdateParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMetadataPropertyUpdateParamsWithDefaults() *BTMetadataPropertyUpdateParams {
	this := BTMetadataPropertyUpdateParams{}
	return &this
}

// GetPropertyId returns the PropertyId field value if set, zero value otherwise.
func (o *BTMetadataPropertyUpdateParams) GetPropertyId() string {
	if o == nil || o.PropertyId == nil {
		var ret string
		return ret
	}
	return *o.PropertyId
}

// GetPropertyIdOk returns a tuple with the PropertyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataPropertyUpdateParams) GetPropertyIdOk() (*string, bool) {
	if o == nil || o.PropertyId == nil {
		return nil, false
	}
	return o.PropertyId, true
}

// HasPropertyId returns a boolean if a field has been set.
func (o *BTMetadataPropertyUpdateParams) HasPropertyId() bool {
	if o != nil && o.PropertyId != nil {
		return true
	}

	return false
}

// SetPropertyId gets a reference to the given string and assigns it to the PropertyId field.
func (o *BTMetadataPropertyUpdateParams) SetPropertyId(v string) {
	o.PropertyId = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *BTMetadataPropertyUpdateParams) GetValue() map[string]interface{} {
	if o == nil || o.Value == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataPropertyUpdateParams) GetValueOk() (*map[string]interface{}, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *BTMetadataPropertyUpdateParams) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given map[string]interface{} and assigns it to the Value field.
func (o *BTMetadataPropertyUpdateParams) SetValue(v map[string]interface{}) {
	o.Value = &v
}

func (o BTMetadataPropertyUpdateParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PropertyId != nil {
		toSerialize["propertyId"] = o.PropertyId
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableBTMetadataPropertyUpdateParams struct {
	value *BTMetadataPropertyUpdateParams
	isSet bool
}

func (v NullableBTMetadataPropertyUpdateParams) Get() *BTMetadataPropertyUpdateParams {
	return v.value
}

func (v *NullableBTMetadataPropertyUpdateParams) Set(val *BTMetadataPropertyUpdateParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMetadataPropertyUpdateParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMetadataPropertyUpdateParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMetadataPropertyUpdateParams(val *BTMetadataPropertyUpdateParams) *NullableBTMetadataPropertyUpdateParams {
	return &NullableBTMetadataPropertyUpdateParams{value: val, isSet: true}
}

func (v NullableBTMetadataPropertyUpdateParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMetadataPropertyUpdateParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
