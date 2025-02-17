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

// BTRevertUnchangedElementParams struct for BTRevertUnchangedElementParams
type BTRevertUnchangedElementParams struct {
	Configuration *string  `json:"configuration,omitempty"`
	ElementId     *string  `json:"elementId,omitempty"`
	ReferenceIds  []string `json:"referenceIds,omitempty"`
}

// NewBTRevertUnchangedElementParams instantiates a new BTRevertUnchangedElementParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTRevertUnchangedElementParams() *BTRevertUnchangedElementParams {
	this := BTRevertUnchangedElementParams{}
	return &this
}

// NewBTRevertUnchangedElementParamsWithDefaults instantiates a new BTRevertUnchangedElementParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTRevertUnchangedElementParamsWithDefaults() *BTRevertUnchangedElementParams {
	this := BTRevertUnchangedElementParams{}
	return &this
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *BTRevertUnchangedElementParams) GetConfiguration() string {
	if o == nil || o.Configuration == nil {
		var ret string
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRevertUnchangedElementParams) GetConfigurationOk() (*string, bool) {
	if o == nil || o.Configuration == nil {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *BTRevertUnchangedElementParams) HasConfiguration() bool {
	if o != nil && o.Configuration != nil {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given string and assigns it to the Configuration field.
func (o *BTRevertUnchangedElementParams) SetConfiguration(v string) {
	o.Configuration = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *BTRevertUnchangedElementParams) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRevertUnchangedElementParams) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *BTRevertUnchangedElementParams) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *BTRevertUnchangedElementParams) SetElementId(v string) {
	o.ElementId = &v
}

// GetReferenceIds returns the ReferenceIds field value if set, zero value otherwise.
func (o *BTRevertUnchangedElementParams) GetReferenceIds() []string {
	if o == nil || o.ReferenceIds == nil {
		var ret []string
		return ret
	}
	return o.ReferenceIds
}

// GetReferenceIdsOk returns a tuple with the ReferenceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRevertUnchangedElementParams) GetReferenceIdsOk() ([]string, bool) {
	if o == nil || o.ReferenceIds == nil {
		return nil, false
	}
	return o.ReferenceIds, true
}

// HasReferenceIds returns a boolean if a field has been set.
func (o *BTRevertUnchangedElementParams) HasReferenceIds() bool {
	if o != nil && o.ReferenceIds != nil {
		return true
	}

	return false
}

// SetReferenceIds gets a reference to the given []string and assigns it to the ReferenceIds field.
func (o *BTRevertUnchangedElementParams) SetReferenceIds(v []string) {
	o.ReferenceIds = v
}

func (o BTRevertUnchangedElementParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Configuration != nil {
		toSerialize["configuration"] = o.Configuration
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.ReferenceIds != nil {
		toSerialize["referenceIds"] = o.ReferenceIds
	}
	return json.Marshal(toSerialize)
}

type NullableBTRevertUnchangedElementParams struct {
	value *BTRevertUnchangedElementParams
	isSet bool
}

func (v NullableBTRevertUnchangedElementParams) Get() *BTRevertUnchangedElementParams {
	return v.value
}

func (v *NullableBTRevertUnchangedElementParams) Set(val *BTRevertUnchangedElementParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTRevertUnchangedElementParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTRevertUnchangedElementParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTRevertUnchangedElementParams(val *BTRevertUnchangedElementParams) *NullableBTRevertUnchangedElementParams {
	return &NullableBTRevertUnchangedElementParams{value: val, isSet: true}
}

func (v NullableBTRevertUnchangedElementParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTRevertUnchangedElementParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
