/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.154.6672-3324fec9d980
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// RoleMapEntry struct for RoleMapEntry
type RoleMapEntry struct {
	Identities []BTIdentityInfo `json:"identities,omitempty"`
	Role       *BTRbacRoleInfo  `json:"role,omitempty"`
}

// NewRoleMapEntry instantiates a new RoleMapEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleMapEntry() *RoleMapEntry {
	this := RoleMapEntry{}
	return &this
}

// NewRoleMapEntryWithDefaults instantiates a new RoleMapEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleMapEntryWithDefaults() *RoleMapEntry {
	this := RoleMapEntry{}
	return &this
}

// GetIdentities returns the Identities field value if set, zero value otherwise.
func (o *RoleMapEntry) GetIdentities() []BTIdentityInfo {
	if o == nil || o.Identities == nil {
		var ret []BTIdentityInfo
		return ret
	}
	return o.Identities
}

// GetIdentitiesOk returns a tuple with the Identities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMapEntry) GetIdentitiesOk() ([]BTIdentityInfo, bool) {
	if o == nil || o.Identities == nil {
		return nil, false
	}
	return o.Identities, true
}

// HasIdentities returns a boolean if a field has been set.
func (o *RoleMapEntry) HasIdentities() bool {
	if o != nil && o.Identities != nil {
		return true
	}

	return false
}

// SetIdentities gets a reference to the given []BTIdentityInfo and assigns it to the Identities field.
func (o *RoleMapEntry) SetIdentities(v []BTIdentityInfo) {
	o.Identities = v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *RoleMapEntry) GetRole() BTRbacRoleInfo {
	if o == nil || o.Role == nil {
		var ret BTRbacRoleInfo
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMapEntry) GetRoleOk() (*BTRbacRoleInfo, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *RoleMapEntry) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given BTRbacRoleInfo and assigns it to the Role field.
func (o *RoleMapEntry) SetRole(v BTRbacRoleInfo) {
	o.Role = &v
}

func (o RoleMapEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Identities != nil {
		toSerialize["identities"] = o.Identities
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	return json.Marshal(toSerialize)
}

type NullableRoleMapEntry struct {
	value *RoleMapEntry
	isSet bool
}

func (v NullableRoleMapEntry) Get() *RoleMapEntry {
	return v.value
}

func (v *NullableRoleMapEntry) Set(val *RoleMapEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleMapEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleMapEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleMapEntry(val *RoleMapEntry) *NullableRoleMapEntry {
	return &NullableRoleMapEntry{value: val, isSet: true}
}

func (v NullableRoleMapEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleMapEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
