/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.155.7076-cd7f519b38e7
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTAliasEntryParams struct for BTAliasEntryParams
type BTAliasEntryParams struct {
	Email  *string `json:"email,omitempty"`
	TeamId *string `json:"teamId,omitempty"`
	UserId *string `json:"userId,omitempty"`
}

// NewBTAliasEntryParams instantiates a new BTAliasEntryParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAliasEntryParams() *BTAliasEntryParams {
	this := BTAliasEntryParams{}
	return &this
}

// NewBTAliasEntryParamsWithDefaults instantiates a new BTAliasEntryParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAliasEntryParamsWithDefaults() *BTAliasEntryParams {
	this := BTAliasEntryParams{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *BTAliasEntryParams) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAliasEntryParams) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *BTAliasEntryParams) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *BTAliasEntryParams) SetEmail(v string) {
	o.Email = &v
}

// GetTeamId returns the TeamId field value if set, zero value otherwise.
func (o *BTAliasEntryParams) GetTeamId() string {
	if o == nil || o.TeamId == nil {
		var ret string
		return ret
	}
	return *o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAliasEntryParams) GetTeamIdOk() (*string, bool) {
	if o == nil || o.TeamId == nil {
		return nil, false
	}
	return o.TeamId, true
}

// HasTeamId returns a boolean if a field has been set.
func (o *BTAliasEntryParams) HasTeamId() bool {
	if o != nil && o.TeamId != nil {
		return true
	}

	return false
}

// SetTeamId gets a reference to the given string and assigns it to the TeamId field.
func (o *BTAliasEntryParams) SetTeamId(v string) {
	o.TeamId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *BTAliasEntryParams) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAliasEntryParams) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *BTAliasEntryParams) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *BTAliasEntryParams) SetUserId(v string) {
	o.UserId = &v
}

func (o BTAliasEntryParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.TeamId != nil {
		toSerialize["teamId"] = o.TeamId
	}
	if o.UserId != nil {
		toSerialize["userId"] = o.UserId
	}
	return json.Marshal(toSerialize)
}

type NullableBTAliasEntryParams struct {
	value *BTAliasEntryParams
	isSet bool
}

func (v NullableBTAliasEntryParams) Get() *BTAliasEntryParams {
	return v.value
}

func (v *NullableBTAliasEntryParams) Set(val *BTAliasEntryParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAliasEntryParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAliasEntryParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAliasEntryParams(val *BTAliasEntryParams) *NullableBTAliasEntryParams {
	return &NullableBTAliasEntryParams{value: val, isSet: true}
}

func (v NullableBTAliasEntryParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAliasEntryParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
