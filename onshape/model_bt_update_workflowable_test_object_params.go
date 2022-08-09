/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.151.5944-34bf93ccfb05
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTUpdateWorkflowableTestObjectParams struct for BTUpdateWorkflowableTestObjectParams
type BTUpdateWorkflowableTestObjectParams struct {
	Properties []BTPropertyValueParam `json:"properties,omitempty"`
}

// NewBTUpdateWorkflowableTestObjectParams instantiates a new BTUpdateWorkflowableTestObjectParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTUpdateWorkflowableTestObjectParams() *BTUpdateWorkflowableTestObjectParams {
	this := BTUpdateWorkflowableTestObjectParams{}
	return &this
}

// NewBTUpdateWorkflowableTestObjectParamsWithDefaults instantiates a new BTUpdateWorkflowableTestObjectParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTUpdateWorkflowableTestObjectParamsWithDefaults() *BTUpdateWorkflowableTestObjectParams {
	this := BTUpdateWorkflowableTestObjectParams{}
	return &this
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *BTUpdateWorkflowableTestObjectParams) GetProperties() []BTPropertyValueParam {
	if o == nil || o.Properties == nil {
		var ret []BTPropertyValueParam
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUpdateWorkflowableTestObjectParams) GetPropertiesOk() ([]BTPropertyValueParam, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *BTUpdateWorkflowableTestObjectParams) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []BTPropertyValueParam and assigns it to the Properties field.
func (o *BTUpdateWorkflowableTestObjectParams) SetProperties(v []BTPropertyValueParam) {
	o.Properties = v
}

func (o BTUpdateWorkflowableTestObjectParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	return json.Marshal(toSerialize)
}

type NullableBTUpdateWorkflowableTestObjectParams struct {
	value *BTUpdateWorkflowableTestObjectParams
	isSet bool
}

func (v NullableBTUpdateWorkflowableTestObjectParams) Get() *BTUpdateWorkflowableTestObjectParams {
	return v.value
}

func (v *NullableBTUpdateWorkflowableTestObjectParams) Set(val *BTUpdateWorkflowableTestObjectParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTUpdateWorkflowableTestObjectParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTUpdateWorkflowableTestObjectParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTUpdateWorkflowableTestObjectParams(val *BTUpdateWorkflowableTestObjectParams) *NullableBTUpdateWorkflowableTestObjectParams {
	return &NullableBTUpdateWorkflowableTestObjectParams{value: val, isSet: true}
}

func (v NullableBTUpdateWorkflowableTestObjectParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTUpdateWorkflowableTestObjectParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
