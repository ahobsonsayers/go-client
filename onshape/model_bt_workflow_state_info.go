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

// BTWorkflowStateInfo struct for BTWorkflowStateInfo
type BTWorkflowStateInfo struct {
	ApproverSourceProperty *string  `json:"approverSourceProperty,omitempty"`
	DisplayName            *string  `json:"displayName,omitempty"`
	EditPermissions        []string `json:"editPermissions,omitempty"`
	EditableProperties     []string `json:"editableProperties,omitempty"`
	Name                   *string  `json:"name,omitempty"`
	NonEditableProperties  []string `json:"nonEditableProperties,omitempty"`
	NotifierSourceProperty *string  `json:"notifierSourceProperty,omitempty"`
	RequiredItemProperties []string `json:"requiredItemProperties,omitempty"`
	RequiredProperties     []string `json:"requiredProperties,omitempty"`
}

// NewBTWorkflowStateInfo instantiates a new BTWorkflowStateInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTWorkflowStateInfo() *BTWorkflowStateInfo {
	this := BTWorkflowStateInfo{}
	return &this
}

// NewBTWorkflowStateInfoWithDefaults instantiates a new BTWorkflowStateInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTWorkflowStateInfoWithDefaults() *BTWorkflowStateInfo {
	this := BTWorkflowStateInfo{}
	return &this
}

// GetApproverSourceProperty returns the ApproverSourceProperty field value if set, zero value otherwise.
func (o *BTWorkflowStateInfo) GetApproverSourceProperty() string {
	if o == nil || o.ApproverSourceProperty == nil {
		var ret string
		return ret
	}
	return *o.ApproverSourceProperty
}

// GetApproverSourcePropertyOk returns a tuple with the ApproverSourceProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowStateInfo) GetApproverSourcePropertyOk() (*string, bool) {
	if o == nil || o.ApproverSourceProperty == nil {
		return nil, false
	}
	return o.ApproverSourceProperty, true
}

// HasApproverSourceProperty returns a boolean if a field has been set.
func (o *BTWorkflowStateInfo) HasApproverSourceProperty() bool {
	if o != nil && o.ApproverSourceProperty != nil {
		return true
	}

	return false
}

// SetApproverSourceProperty gets a reference to the given string and assigns it to the ApproverSourceProperty field.
func (o *BTWorkflowStateInfo) SetApproverSourceProperty(v string) {
	o.ApproverSourceProperty = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *BTWorkflowStateInfo) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowStateInfo) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *BTWorkflowStateInfo) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *BTWorkflowStateInfo) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEditPermissions returns the EditPermissions field value if set, zero value otherwise.
func (o *BTWorkflowStateInfo) GetEditPermissions() []string {
	if o == nil || o.EditPermissions == nil {
		var ret []string
		return ret
	}
	return o.EditPermissions
}

// GetEditPermissionsOk returns a tuple with the EditPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowStateInfo) GetEditPermissionsOk() ([]string, bool) {
	if o == nil || o.EditPermissions == nil {
		return nil, false
	}
	return o.EditPermissions, true
}

// HasEditPermissions returns a boolean if a field has been set.
func (o *BTWorkflowStateInfo) HasEditPermissions() bool {
	if o != nil && o.EditPermissions != nil {
		return true
	}

	return false
}

// SetEditPermissions gets a reference to the given []string and assigns it to the EditPermissions field.
func (o *BTWorkflowStateInfo) SetEditPermissions(v []string) {
	o.EditPermissions = v
}

// GetEditableProperties returns the EditableProperties field value if set, zero value otherwise.
func (o *BTWorkflowStateInfo) GetEditableProperties() []string {
	if o == nil || o.EditableProperties == nil {
		var ret []string
		return ret
	}
	return o.EditableProperties
}

// GetEditablePropertiesOk returns a tuple with the EditableProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowStateInfo) GetEditablePropertiesOk() ([]string, bool) {
	if o == nil || o.EditableProperties == nil {
		return nil, false
	}
	return o.EditableProperties, true
}

// HasEditableProperties returns a boolean if a field has been set.
func (o *BTWorkflowStateInfo) HasEditableProperties() bool {
	if o != nil && o.EditableProperties != nil {
		return true
	}

	return false
}

// SetEditableProperties gets a reference to the given []string and assigns it to the EditableProperties field.
func (o *BTWorkflowStateInfo) SetEditableProperties(v []string) {
	o.EditableProperties = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTWorkflowStateInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowStateInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTWorkflowStateInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTWorkflowStateInfo) SetName(v string) {
	o.Name = &v
}

// GetNonEditableProperties returns the NonEditableProperties field value if set, zero value otherwise.
func (o *BTWorkflowStateInfo) GetNonEditableProperties() []string {
	if o == nil || o.NonEditableProperties == nil {
		var ret []string
		return ret
	}
	return o.NonEditableProperties
}

// GetNonEditablePropertiesOk returns a tuple with the NonEditableProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowStateInfo) GetNonEditablePropertiesOk() ([]string, bool) {
	if o == nil || o.NonEditableProperties == nil {
		return nil, false
	}
	return o.NonEditableProperties, true
}

// HasNonEditableProperties returns a boolean if a field has been set.
func (o *BTWorkflowStateInfo) HasNonEditableProperties() bool {
	if o != nil && o.NonEditableProperties != nil {
		return true
	}

	return false
}

// SetNonEditableProperties gets a reference to the given []string and assigns it to the NonEditableProperties field.
func (o *BTWorkflowStateInfo) SetNonEditableProperties(v []string) {
	o.NonEditableProperties = v
}

// GetNotifierSourceProperty returns the NotifierSourceProperty field value if set, zero value otherwise.
func (o *BTWorkflowStateInfo) GetNotifierSourceProperty() string {
	if o == nil || o.NotifierSourceProperty == nil {
		var ret string
		return ret
	}
	return *o.NotifierSourceProperty
}

// GetNotifierSourcePropertyOk returns a tuple with the NotifierSourceProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowStateInfo) GetNotifierSourcePropertyOk() (*string, bool) {
	if o == nil || o.NotifierSourceProperty == nil {
		return nil, false
	}
	return o.NotifierSourceProperty, true
}

// HasNotifierSourceProperty returns a boolean if a field has been set.
func (o *BTWorkflowStateInfo) HasNotifierSourceProperty() bool {
	if o != nil && o.NotifierSourceProperty != nil {
		return true
	}

	return false
}

// SetNotifierSourceProperty gets a reference to the given string and assigns it to the NotifierSourceProperty field.
func (o *BTWorkflowStateInfo) SetNotifierSourceProperty(v string) {
	o.NotifierSourceProperty = &v
}

// GetRequiredItemProperties returns the RequiredItemProperties field value if set, zero value otherwise.
func (o *BTWorkflowStateInfo) GetRequiredItemProperties() []string {
	if o == nil || o.RequiredItemProperties == nil {
		var ret []string
		return ret
	}
	return o.RequiredItemProperties
}

// GetRequiredItemPropertiesOk returns a tuple with the RequiredItemProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowStateInfo) GetRequiredItemPropertiesOk() ([]string, bool) {
	if o == nil || o.RequiredItemProperties == nil {
		return nil, false
	}
	return o.RequiredItemProperties, true
}

// HasRequiredItemProperties returns a boolean if a field has been set.
func (o *BTWorkflowStateInfo) HasRequiredItemProperties() bool {
	if o != nil && o.RequiredItemProperties != nil {
		return true
	}

	return false
}

// SetRequiredItemProperties gets a reference to the given []string and assigns it to the RequiredItemProperties field.
func (o *BTWorkflowStateInfo) SetRequiredItemProperties(v []string) {
	o.RequiredItemProperties = v
}

// GetRequiredProperties returns the RequiredProperties field value if set, zero value otherwise.
func (o *BTWorkflowStateInfo) GetRequiredProperties() []string {
	if o == nil || o.RequiredProperties == nil {
		var ret []string
		return ret
	}
	return o.RequiredProperties
}

// GetRequiredPropertiesOk returns a tuple with the RequiredProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowStateInfo) GetRequiredPropertiesOk() ([]string, bool) {
	if o == nil || o.RequiredProperties == nil {
		return nil, false
	}
	return o.RequiredProperties, true
}

// HasRequiredProperties returns a boolean if a field has been set.
func (o *BTWorkflowStateInfo) HasRequiredProperties() bool {
	if o != nil && o.RequiredProperties != nil {
		return true
	}

	return false
}

// SetRequiredProperties gets a reference to the given []string and assigns it to the RequiredProperties field.
func (o *BTWorkflowStateInfo) SetRequiredProperties(v []string) {
	o.RequiredProperties = v
}

func (o BTWorkflowStateInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApproverSourceProperty != nil {
		toSerialize["approverSourceProperty"] = o.ApproverSourceProperty
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.EditPermissions != nil {
		toSerialize["editPermissions"] = o.EditPermissions
	}
	if o.EditableProperties != nil {
		toSerialize["editableProperties"] = o.EditableProperties
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.NonEditableProperties != nil {
		toSerialize["nonEditableProperties"] = o.NonEditableProperties
	}
	if o.NotifierSourceProperty != nil {
		toSerialize["notifierSourceProperty"] = o.NotifierSourceProperty
	}
	if o.RequiredItemProperties != nil {
		toSerialize["requiredItemProperties"] = o.RequiredItemProperties
	}
	if o.RequiredProperties != nil {
		toSerialize["requiredProperties"] = o.RequiredProperties
	}
	return json.Marshal(toSerialize)
}

type NullableBTWorkflowStateInfo struct {
	value *BTWorkflowStateInfo
	isSet bool
}

func (v NullableBTWorkflowStateInfo) Get() *BTWorkflowStateInfo {
	return v.value
}

func (v *NullableBTWorkflowStateInfo) Set(val *BTWorkflowStateInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTWorkflowStateInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTWorkflowStateInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTWorkflowStateInfo(val *BTWorkflowStateInfo) *NullableBTWorkflowStateInfo {
	return &NullableBTWorkflowStateInfo{value: val, isSet: true}
}

func (v NullableBTWorkflowStateInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTWorkflowStateInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
