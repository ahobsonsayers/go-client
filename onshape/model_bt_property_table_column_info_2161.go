/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.162.14462-13ace71ec1df
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTPropertyTableColumnInfo2161 struct for BTPropertyTableColumnInfo2161
type BTPropertyTableColumnInfo2161 struct {
	Id                         *string                `json:"id,omitempty"`
	NodeId                     *string                `json:"nodeId,omitempty"`
	Specification              *BTTableColumnSpec1967 `json:"specification,omitempty"`
	BtType                     *string                `json:"btType,omitempty"`
	IsComputedAssemblyProperty *bool                  `json:"isComputedAssemblyProperty,omitempty"`
	IsComputedProperty         *bool                  `json:"isComputedProperty,omitempty"`
}

// NewBTPropertyTableColumnInfo2161 instantiates a new BTPropertyTableColumnInfo2161 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPropertyTableColumnInfo2161() *BTPropertyTableColumnInfo2161 {
	this := BTPropertyTableColumnInfo2161{}
	return &this
}

// NewBTPropertyTableColumnInfo2161WithDefaults instantiates a new BTPropertyTableColumnInfo2161 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPropertyTableColumnInfo2161WithDefaults() *BTPropertyTableColumnInfo2161 {
	this := BTPropertyTableColumnInfo2161{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTPropertyTableColumnInfo2161) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPropertyTableColumnInfo2161) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTPropertyTableColumnInfo2161) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTPropertyTableColumnInfo2161) SetId(v string) {
	o.Id = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTPropertyTableColumnInfo2161) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPropertyTableColumnInfo2161) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTPropertyTableColumnInfo2161) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTPropertyTableColumnInfo2161) SetNodeId(v string) {
	o.NodeId = &v
}

// GetSpecification returns the Specification field value if set, zero value otherwise.
func (o *BTPropertyTableColumnInfo2161) GetSpecification() BTTableColumnSpec1967 {
	if o == nil || o.Specification == nil {
		var ret BTTableColumnSpec1967
		return ret
	}
	return *o.Specification
}

// GetSpecificationOk returns a tuple with the Specification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPropertyTableColumnInfo2161) GetSpecificationOk() (*BTTableColumnSpec1967, bool) {
	if o == nil || o.Specification == nil {
		return nil, false
	}
	return o.Specification, true
}

// HasSpecification returns a boolean if a field has been set.
func (o *BTPropertyTableColumnInfo2161) HasSpecification() bool {
	if o != nil && o.Specification != nil {
		return true
	}

	return false
}

// SetSpecification gets a reference to the given BTTableColumnSpec1967 and assigns it to the Specification field.
func (o *BTPropertyTableColumnInfo2161) SetSpecification(v BTTableColumnSpec1967) {
	o.Specification = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPropertyTableColumnInfo2161) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPropertyTableColumnInfo2161) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPropertyTableColumnInfo2161) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPropertyTableColumnInfo2161) SetBtType(v string) {
	o.BtType = &v
}

// GetIsComputedAssemblyProperty returns the IsComputedAssemblyProperty field value if set, zero value otherwise.
func (o *BTPropertyTableColumnInfo2161) GetIsComputedAssemblyProperty() bool {
	if o == nil || o.IsComputedAssemblyProperty == nil {
		var ret bool
		return ret
	}
	return *o.IsComputedAssemblyProperty
}

// GetIsComputedAssemblyPropertyOk returns a tuple with the IsComputedAssemblyProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPropertyTableColumnInfo2161) GetIsComputedAssemblyPropertyOk() (*bool, bool) {
	if o == nil || o.IsComputedAssemblyProperty == nil {
		return nil, false
	}
	return o.IsComputedAssemblyProperty, true
}

// HasIsComputedAssemblyProperty returns a boolean if a field has been set.
func (o *BTPropertyTableColumnInfo2161) HasIsComputedAssemblyProperty() bool {
	if o != nil && o.IsComputedAssemblyProperty != nil {
		return true
	}

	return false
}

// SetIsComputedAssemblyProperty gets a reference to the given bool and assigns it to the IsComputedAssemblyProperty field.
func (o *BTPropertyTableColumnInfo2161) SetIsComputedAssemblyProperty(v bool) {
	o.IsComputedAssemblyProperty = &v
}

// GetIsComputedProperty returns the IsComputedProperty field value if set, zero value otherwise.
func (o *BTPropertyTableColumnInfo2161) GetIsComputedProperty() bool {
	if o == nil || o.IsComputedProperty == nil {
		var ret bool
		return ret
	}
	return *o.IsComputedProperty
}

// GetIsComputedPropertyOk returns a tuple with the IsComputedProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPropertyTableColumnInfo2161) GetIsComputedPropertyOk() (*bool, bool) {
	if o == nil || o.IsComputedProperty == nil {
		return nil, false
	}
	return o.IsComputedProperty, true
}

// HasIsComputedProperty returns a boolean if a field has been set.
func (o *BTPropertyTableColumnInfo2161) HasIsComputedProperty() bool {
	if o != nil && o.IsComputedProperty != nil {
		return true
	}

	return false
}

// SetIsComputedProperty gets a reference to the given bool and assigns it to the IsComputedProperty field.
func (o *BTPropertyTableColumnInfo2161) SetIsComputedProperty(v bool) {
	o.IsComputedProperty = &v
}

func (o BTPropertyTableColumnInfo2161) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.Specification != nil {
		toSerialize["specification"] = o.Specification
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.IsComputedAssemblyProperty != nil {
		toSerialize["isComputedAssemblyProperty"] = o.IsComputedAssemblyProperty
	}
	if o.IsComputedProperty != nil {
		toSerialize["isComputedProperty"] = o.IsComputedProperty
	}
	return json.Marshal(toSerialize)
}

type NullableBTPropertyTableColumnInfo2161 struct {
	value *BTPropertyTableColumnInfo2161
	isSet bool
}

func (v NullableBTPropertyTableColumnInfo2161) Get() *BTPropertyTableColumnInfo2161 {
	return v.value
}

func (v *NullableBTPropertyTableColumnInfo2161) Set(val *BTPropertyTableColumnInfo2161) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPropertyTableColumnInfo2161) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPropertyTableColumnInfo2161) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPropertyTableColumnInfo2161(val *BTPropertyTableColumnInfo2161) *NullableBTPropertyTableColumnInfo2161 {
	return &NullableBTPropertyTableColumnInfo2161{value: val, isSet: true}
}

func (v NullableBTPropertyTableColumnInfo2161) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPropertyTableColumnInfo2161) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
