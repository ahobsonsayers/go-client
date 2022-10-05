/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.154.6764-8cd829594e49
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTExportTessellatedBody3398 - struct for BTExportTessellatedBody3398
type BTExportTessellatedBody3398 struct {
	implBTExportTessellatedBody3398 interface{}
}

// BTExportTessellatedEdgesBody890AsBTExportTessellatedBody3398 is a convenience function that returns BTExportTessellatedEdgesBody890 wrapped in BTExportTessellatedBody3398
func (o *BTExportTessellatedEdgesBody890) AsBTExportTessellatedBody3398() *BTExportTessellatedBody3398 {
	return &BTExportTessellatedBody3398{o}
}

// BTExportTessellatedFacesBody1321AsBTExportTessellatedBody3398 is a convenience function that returns BTExportTessellatedFacesBody1321 wrapped in BTExportTessellatedBody3398
func (o *BTExportTessellatedFacesBody1321) AsBTExportTessellatedBody3398() *BTExportTessellatedBody3398 {
	return &BTExportTessellatedBody3398{o}
}

// NewBTExportTessellatedBody3398 instantiates a new BTExportTessellatedBody3398 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTExportTessellatedBody3398() *BTExportTessellatedBody3398 {
	this := BTExportTessellatedBody3398{Newbase_BTExportTessellatedBody3398()}
	return &this
}

// NewBTExportTessellatedBody3398WithDefaults instantiates a new BTExportTessellatedBody3398 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTExportTessellatedBody3398WithDefaults() *BTExportTessellatedBody3398 {
	this := BTExportTessellatedBody3398{Newbase_BTExportTessellatedBody3398WithDefaults()}
	return &this
}

// GetConstituents returns the Constituents field value if set, zero value otherwise.
func (o *BTExportTessellatedBody3398) GetConstituents() []string {
	type getResult interface {
		GetConstituents() []string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetConstituents()
	} else {
		var de []string
		return de
	}
}

// GetConstituentsOk returns a tuple with the Constituents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportTessellatedBody3398) GetConstituentsOk() ([]string, bool) {
	type getResult interface {
		GetConstituentsOk() ([]string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetConstituentsOk()
	} else {
		return nil, false
	}
}

// HasConstituents returns a boolean if a field has been set.
func (o *BTExportTessellatedBody3398) HasConstituents() bool {
	type getResult interface {
		HasConstituents() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasConstituents()
	} else {
		return false
	}
}

// SetConstituents gets a reference to the given []string and assigns it to the Constituents field.
func (o *BTExportTessellatedBody3398) SetConstituents(v []string) {
	type getResult interface {
		SetConstituents(v []string)
	}

	o.GetActualInstance().(getResult).SetConstituents(v)
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTExportTessellatedBody3398) GetId() string {
	type getResult interface {
		GetId() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetId()
	} else {
		var de string
		return de
	}
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportTessellatedBody3398) GetIdOk() (*string, bool) {
	type getResult interface {
		GetIdOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetIdOk()
	} else {
		return nil, false
	}
}

// HasId returns a boolean if a field has been set.
func (o *BTExportTessellatedBody3398) HasId() bool {
	type getResult interface {
		HasId() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasId()
	} else {
		return false
	}
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTExportTessellatedBody3398) SetId(v string) {
	type getResult interface {
		SetId(v string)
	}

	o.GetActualInstance().(getResult).SetId(v)
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTExportTessellatedBody3398) GetName() string {
	type getResult interface {
		GetName() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetName()
	} else {
		var de string
		return de
	}
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportTessellatedBody3398) GetNameOk() (*string, bool) {
	type getResult interface {
		GetNameOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetNameOk()
	} else {
		return nil, false
	}
}

// HasName returns a boolean if a field has been set.
func (o *BTExportTessellatedBody3398) HasName() bool {
	type getResult interface {
		HasName() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasName()
	} else {
		return false
	}
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTExportTessellatedBody3398) SetName(v string) {
	type getResult interface {
		SetName(v string)
	}

	o.GetActualInstance().(getResult).SetName(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTExportTessellatedBody3398) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTExportTessellatedEdgesBody-890'
	if jsonDict["btType"] == "BTExportTessellatedEdgesBody-890" {
		// try to unmarshal JSON data into BTExportTessellatedEdgesBody890
		var qr *BTExportTessellatedEdgesBody890
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTExportTessellatedBody3398 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTExportTessellatedBody3398 = nil
			return fmt.Errorf("Failed to unmarshal BTExportTessellatedBody3398 as BTExportTessellatedEdgesBody890: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTExportTessellatedFacesBody-1321'
	if jsonDict["btType"] == "BTExportTessellatedFacesBody-1321" {
		// try to unmarshal JSON data into BTExportTessellatedFacesBody1321
		var qr *BTExportTessellatedFacesBody1321
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTExportTessellatedBody3398 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTExportTessellatedBody3398 = nil
			return fmt.Errorf("Failed to unmarshal BTExportTessellatedBody3398 as BTExportTessellatedFacesBody1321: %s", err.Error())
		}
	}

	var qtx *base_BTExportTessellatedBody3398
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTExportTessellatedBody3398 = qtx
		return nil // data stored in dst.base_BTExportTessellatedBody3398, return on the first match
	} else {
		dst.implBTExportTessellatedBody3398 = nil
		return fmt.Errorf("Failed to unmarshal BTExportTessellatedBody3398 as base_BTExportTessellatedBody3398: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTExportTessellatedBody3398) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTExportTessellatedBody3398) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTExportTessellatedBody3398
}

type NullableBTExportTessellatedBody3398 struct {
	value *BTExportTessellatedBody3398
	isSet bool
}

func (v NullableBTExportTessellatedBody3398) Get() *BTExportTessellatedBody3398 {
	return v.value
}

func (v *NullableBTExportTessellatedBody3398) Set(val *BTExportTessellatedBody3398) {
	v.value = val
	v.isSet = true
}

func (v NullableBTExportTessellatedBody3398) IsSet() bool {
	return v.isSet
}

func (v *NullableBTExportTessellatedBody3398) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTExportTessellatedBody3398(val *BTExportTessellatedBody3398) *NullableBTExportTessellatedBody3398 {
	return &NullableBTExportTessellatedBody3398{value: val, isSet: true}
}

func (v NullableBTExportTessellatedBody3398) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTExportTessellatedBody3398) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTExportTessellatedBody3398 struct {
	Constituents []string `json:"constituents,omitempty"`
	Id           *string  `json:"id,omitempty"`
	Name         *string  `json:"name,omitempty"`
}

// Newbase_BTExportTessellatedBody3398 instantiates a new base_BTExportTessellatedBody3398 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTExportTessellatedBody3398() *base_BTExportTessellatedBody3398 {
	this := base_BTExportTessellatedBody3398{}
	return &this
}

// Newbase_BTExportTessellatedBody3398WithDefaults instantiates a new base_BTExportTessellatedBody3398 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTExportTessellatedBody3398WithDefaults() *base_BTExportTessellatedBody3398 {
	this := base_BTExportTessellatedBody3398{}
	return &this
}

// GetConstituents returns the Constituents field value if set, zero value otherwise.
func (o *base_BTExportTessellatedBody3398) GetConstituents() []string {
	if o == nil || o.Constituents == nil {
		var ret []string
		return ret
	}
	return o.Constituents
}

// GetConstituentsOk returns a tuple with the Constituents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTExportTessellatedBody3398) GetConstituentsOk() ([]string, bool) {
	if o == nil || o.Constituents == nil {
		return nil, false
	}
	return o.Constituents, true
}

// HasConstituents returns a boolean if a field has been set.
func (o *base_BTExportTessellatedBody3398) HasConstituents() bool {
	if o != nil && o.Constituents != nil {
		return true
	}

	return false
}

// SetConstituents gets a reference to the given []string and assigns it to the Constituents field.
func (o *base_BTExportTessellatedBody3398) SetConstituents(v []string) {
	o.Constituents = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *base_BTExportTessellatedBody3398) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTExportTessellatedBody3398) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *base_BTExportTessellatedBody3398) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *base_BTExportTessellatedBody3398) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *base_BTExportTessellatedBody3398) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTExportTessellatedBody3398) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *base_BTExportTessellatedBody3398) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *base_BTExportTessellatedBody3398) SetName(v string) {
	o.Name = &v
}

func (o base_BTExportTessellatedBody3398) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Constituents != nil {
		toSerialize["constituents"] = o.Constituents
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}
