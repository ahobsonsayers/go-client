/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.152.6290-b55936bc8c5a
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTVariableParams struct for BTVariableParams
type BTVariableParams struct {
	// Variable description
	Description *string `json:"description,omitempty"`
	// Variable definition expression
	Expression string `json:"expression"`
	// Variable name
	Name string `json:"name"`
	// VariableType name, from FeatureScript VariableType
	Type string `json:"type"`
}

// NewBTVariableParams instantiates a new BTVariableParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTVariableParams(expression string, name string, type_ string) *BTVariableParams {
	this := BTVariableParams{}
	this.Expression = expression
	this.Name = name
	this.Type = type_
	return &this
}

// NewBTVariableParamsWithDefaults instantiates a new BTVariableParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTVariableParamsWithDefaults() *BTVariableParams {
	this := BTVariableParams{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BTVariableParams) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTVariableParams) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BTVariableParams) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BTVariableParams) SetDescription(v string) {
	o.Description = &v
}

// GetExpression returns the Expression field value
func (o *BTVariableParams) GetExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *BTVariableParams) GetExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expression, true
}

// SetExpression sets field value
func (o *BTVariableParams) SetExpression(v string) {
	o.Expression = v
}

// GetName returns the Name field value
func (o *BTVariableParams) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BTVariableParams) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BTVariableParams) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *BTVariableParams) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BTVariableParams) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BTVariableParams) SetType(v string) {
	o.Type = v
}

func (o BTVariableParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["expression"] = o.Expression
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableBTVariableParams struct {
	value *BTVariableParams
	isSet bool
}

func (v NullableBTVariableParams) Get() *BTVariableParams {
	return v.value
}

func (v *NullableBTVariableParams) Set(val *BTVariableParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTVariableParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTVariableParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTVariableParams(val *BTVariableParams) *NullableBTVariableParams {
	return &NullableBTVariableParams{value: val, isSet: true}
}

func (v NullableBTVariableParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTVariableParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
