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

// BTDocumentVersionElementIds1897 struct for BTDocumentVersionElementIds1897
type BTDocumentVersionElementIds1897 struct {
	BtType     *string `json:"btType,omitempty"`
	DocumentId *string `json:"documentId,omitempty"`
	ElementId  *string `json:"elementId,omitempty"`
	VersionId  *string `json:"versionId,omitempty"`
}

// NewBTDocumentVersionElementIds1897 instantiates a new BTDocumentVersionElementIds1897 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTDocumentVersionElementIds1897() *BTDocumentVersionElementIds1897 {
	this := BTDocumentVersionElementIds1897{}
	return &this
}

// NewBTDocumentVersionElementIds1897WithDefaults instantiates a new BTDocumentVersionElementIds1897 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTDocumentVersionElementIds1897WithDefaults() *BTDocumentVersionElementIds1897 {
	this := BTDocumentVersionElementIds1897{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTDocumentVersionElementIds1897) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentVersionElementIds1897) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTDocumentVersionElementIds1897) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTDocumentVersionElementIds1897) SetBtType(v string) {
	o.BtType = &v
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *BTDocumentVersionElementIds1897) GetDocumentId() string {
	if o == nil || o.DocumentId == nil {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentVersionElementIds1897) GetDocumentIdOk() (*string, bool) {
	if o == nil || o.DocumentId == nil {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *BTDocumentVersionElementIds1897) HasDocumentId() bool {
	if o != nil && o.DocumentId != nil {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *BTDocumentVersionElementIds1897) SetDocumentId(v string) {
	o.DocumentId = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *BTDocumentVersionElementIds1897) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentVersionElementIds1897) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *BTDocumentVersionElementIds1897) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *BTDocumentVersionElementIds1897) SetElementId(v string) {
	o.ElementId = &v
}

// GetVersionId returns the VersionId field value if set, zero value otherwise.
func (o *BTDocumentVersionElementIds1897) GetVersionId() string {
	if o == nil || o.VersionId == nil {
		var ret string
		return ret
	}
	return *o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentVersionElementIds1897) GetVersionIdOk() (*string, bool) {
	if o == nil || o.VersionId == nil {
		return nil, false
	}
	return o.VersionId, true
}

// HasVersionId returns a boolean if a field has been set.
func (o *BTDocumentVersionElementIds1897) HasVersionId() bool {
	if o != nil && o.VersionId != nil {
		return true
	}

	return false
}

// SetVersionId gets a reference to the given string and assigns it to the VersionId field.
func (o *BTDocumentVersionElementIds1897) SetVersionId(v string) {
	o.VersionId = &v
}

func (o BTDocumentVersionElementIds1897) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.DocumentId != nil {
		toSerialize["documentId"] = o.DocumentId
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.VersionId != nil {
		toSerialize["versionId"] = o.VersionId
	}
	return json.Marshal(toSerialize)
}

type NullableBTDocumentVersionElementIds1897 struct {
	value *BTDocumentVersionElementIds1897
	isSet bool
}

func (v NullableBTDocumentVersionElementIds1897) Get() *BTDocumentVersionElementIds1897 {
	return v.value
}

func (v *NullableBTDocumentVersionElementIds1897) Set(val *BTDocumentVersionElementIds1897) {
	v.value = val
	v.isSet = true
}

func (v NullableBTDocumentVersionElementIds1897) IsSet() bool {
	return v.isSet
}

func (v *NullableBTDocumentVersionElementIds1897) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTDocumentVersionElementIds1897(val *BTDocumentVersionElementIds1897) *NullableBTDocumentVersionElementIds1897 {
	return &NullableBTDocumentVersionElementIds1897{value: val, isSet: true}
}

func (v NullableBTDocumentVersionElementIds1897) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTDocumentVersionElementIds1897) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
