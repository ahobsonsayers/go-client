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

// BTRevisionApproverInfo struct for BTRevisionApproverInfo
type BTRevisionApproverInfo struct {
	Date *JSONTime `json:"date,omitempty"`
	Id   *string   `json:"id,omitempty"`
	Name *string   `json:"name,omitempty"`
}

// NewBTRevisionApproverInfo instantiates a new BTRevisionApproverInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTRevisionApproverInfo() *BTRevisionApproverInfo {
	this := BTRevisionApproverInfo{}
	return &this
}

// NewBTRevisionApproverInfoWithDefaults instantiates a new BTRevisionApproverInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTRevisionApproverInfoWithDefaults() *BTRevisionApproverInfo {
	this := BTRevisionApproverInfo{}
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *BTRevisionApproverInfo) GetDate() JSONTime {
	if o == nil || o.Date == nil {
		var ret JSONTime
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRevisionApproverInfo) GetDateOk() (*JSONTime, bool) {
	if o == nil || o.Date == nil {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *BTRevisionApproverInfo) HasDate() bool {
	if o != nil && o.Date != nil {
		return true
	}

	return false
}

// SetDate gets a reference to the given JSONTime and assigns it to the Date field.
func (o *BTRevisionApproverInfo) SetDate(v JSONTime) {
	o.Date = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTRevisionApproverInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRevisionApproverInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTRevisionApproverInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTRevisionApproverInfo) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTRevisionApproverInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRevisionApproverInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTRevisionApproverInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTRevisionApproverInfo) SetName(v string) {
	o.Name = &v
}

func (o BTRevisionApproverInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Date != nil {
		toSerialize["date"] = o.Date
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableBTRevisionApproverInfo struct {
	value *BTRevisionApproverInfo
	isSet bool
}

func (v NullableBTRevisionApproverInfo) Get() *BTRevisionApproverInfo {
	return v.value
}

func (v *NullableBTRevisionApproverInfo) Set(val *BTRevisionApproverInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTRevisionApproverInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTRevisionApproverInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTRevisionApproverInfo(val *BTRevisionApproverInfo) *NullableBTRevisionApproverInfo {
	return &NullableBTRevisionApproverInfo{value: val, isSet: true}
}

func (v NullableBTRevisionApproverInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTRevisionApproverInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
