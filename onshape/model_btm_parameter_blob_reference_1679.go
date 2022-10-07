/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.154.6797-6824f8df9946
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMParameterBlobReference1679 struct for BTMParameterBlobReference1679
type BTMParameterBlobReference1679 struct {
	ImportMicroversion *string       `json:"importMicroversion,omitempty"`
	NodeId             *string       `json:"nodeId,omitempty"`
	ParameterId        *string       `json:"parameterId,omitempty"`
	BlobImport         *BTMImport136 `json:"blobImport,omitempty"`
	BtType             *string       `json:"btType,omitempty"`
	Namespace          *string       `json:"namespace,omitempty"`
}

// NewBTMParameterBlobReference1679 instantiates a new BTMParameterBlobReference1679 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMParameterBlobReference1679() *BTMParameterBlobReference1679 {
	this := BTMParameterBlobReference1679{}
	return &this
}

// NewBTMParameterBlobReference1679WithDefaults instantiates a new BTMParameterBlobReference1679 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMParameterBlobReference1679WithDefaults() *BTMParameterBlobReference1679 {
	this := BTMParameterBlobReference1679{}
	return &this
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMParameterBlobReference1679) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterBlobReference1679) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMParameterBlobReference1679) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMParameterBlobReference1679) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMParameterBlobReference1679) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterBlobReference1679) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMParameterBlobReference1679) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMParameterBlobReference1679) SetNodeId(v string) {
	o.NodeId = &v
}

// GetParameterId returns the ParameterId field value if set, zero value otherwise.
func (o *BTMParameterBlobReference1679) GetParameterId() string {
	if o == nil || o.ParameterId == nil {
		var ret string
		return ret
	}
	return *o.ParameterId
}

// GetParameterIdOk returns a tuple with the ParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterBlobReference1679) GetParameterIdOk() (*string, bool) {
	if o == nil || o.ParameterId == nil {
		return nil, false
	}
	return o.ParameterId, true
}

// HasParameterId returns a boolean if a field has been set.
func (o *BTMParameterBlobReference1679) HasParameterId() bool {
	if o != nil && o.ParameterId != nil {
		return true
	}

	return false
}

// SetParameterId gets a reference to the given string and assigns it to the ParameterId field.
func (o *BTMParameterBlobReference1679) SetParameterId(v string) {
	o.ParameterId = &v
}

// GetBlobImport returns the BlobImport field value if set, zero value otherwise.
func (o *BTMParameterBlobReference1679) GetBlobImport() BTMImport136 {
	if o == nil || o.BlobImport == nil {
		var ret BTMImport136
		return ret
	}
	return *o.BlobImport
}

// GetBlobImportOk returns a tuple with the BlobImport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterBlobReference1679) GetBlobImportOk() (*BTMImport136, bool) {
	if o == nil || o.BlobImport == nil {
		return nil, false
	}
	return o.BlobImport, true
}

// HasBlobImport returns a boolean if a field has been set.
func (o *BTMParameterBlobReference1679) HasBlobImport() bool {
	if o != nil && o.BlobImport != nil {
		return true
	}

	return false
}

// SetBlobImport gets a reference to the given BTMImport136 and assigns it to the BlobImport field.
func (o *BTMParameterBlobReference1679) SetBlobImport(v BTMImport136) {
	o.BlobImport = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMParameterBlobReference1679) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterBlobReference1679) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMParameterBlobReference1679) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMParameterBlobReference1679) SetBtType(v string) {
	o.BtType = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *BTMParameterBlobReference1679) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterBlobReference1679) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *BTMParameterBlobReference1679) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *BTMParameterBlobReference1679) SetNamespace(v string) {
	o.Namespace = &v
}

func (o BTMParameterBlobReference1679) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.ParameterId != nil {
		toSerialize["parameterId"] = o.ParameterId
	}
	if o.BlobImport != nil {
		toSerialize["blobImport"] = o.BlobImport
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	return json.Marshal(toSerialize)
}

type NullableBTMParameterBlobReference1679 struct {
	value *BTMParameterBlobReference1679
	isSet bool
}

func (v NullableBTMParameterBlobReference1679) Get() *BTMParameterBlobReference1679 {
	return v.value
}

func (v *NullableBTMParameterBlobReference1679) Set(val *BTMParameterBlobReference1679) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMParameterBlobReference1679) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMParameterBlobReference1679) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMParameterBlobReference1679(val *BTMParameterBlobReference1679) *NullableBTMParameterBlobReference1679 {
	return &NullableBTMParameterBlobReference1679{value: val, isSet: true}
}

func (v NullableBTMParameterBlobReference1679) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMParameterBlobReference1679) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
