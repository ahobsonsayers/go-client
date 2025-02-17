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

// RequestOptions struct for RequestOptions
type RequestOptions struct {
	ApiKey         *string `json:"apiKey,omitempty"`
	ConnectTimeout *int32  `json:"connectTimeout,omitempty"`
	IdempotencyKey *string `json:"idempotencyKey,omitempty"`
	ReadTimeout    *int32  `json:"readTimeout,omitempty"`
	StripeAccount  *string `json:"stripeAccount,omitempty"`
	StripeVersion  *string `json:"stripeVersion,omitempty"`
}

// NewRequestOptions instantiates a new RequestOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestOptions() *RequestOptions {
	this := RequestOptions{}
	return &this
}

// NewRequestOptionsWithDefaults instantiates a new RequestOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestOptionsWithDefaults() *RequestOptions {
	this := RequestOptions{}
	return &this
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *RequestOptions) GetApiKey() string {
	if o == nil || o.ApiKey == nil {
		var ret string
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestOptions) GetApiKeyOk() (*string, bool) {
	if o == nil || o.ApiKey == nil {
		return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *RequestOptions) HasApiKey() bool {
	if o != nil && o.ApiKey != nil {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given string and assigns it to the ApiKey field.
func (o *RequestOptions) SetApiKey(v string) {
	o.ApiKey = &v
}

// GetConnectTimeout returns the ConnectTimeout field value if set, zero value otherwise.
func (o *RequestOptions) GetConnectTimeout() int32 {
	if o == nil || o.ConnectTimeout == nil {
		var ret int32
		return ret
	}
	return *o.ConnectTimeout
}

// GetConnectTimeoutOk returns a tuple with the ConnectTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestOptions) GetConnectTimeoutOk() (*int32, bool) {
	if o == nil || o.ConnectTimeout == nil {
		return nil, false
	}
	return o.ConnectTimeout, true
}

// HasConnectTimeout returns a boolean if a field has been set.
func (o *RequestOptions) HasConnectTimeout() bool {
	if o != nil && o.ConnectTimeout != nil {
		return true
	}

	return false
}

// SetConnectTimeout gets a reference to the given int32 and assigns it to the ConnectTimeout field.
func (o *RequestOptions) SetConnectTimeout(v int32) {
	o.ConnectTimeout = &v
}

// GetIdempotencyKey returns the IdempotencyKey field value if set, zero value otherwise.
func (o *RequestOptions) GetIdempotencyKey() string {
	if o == nil || o.IdempotencyKey == nil {
		var ret string
		return ret
	}
	return *o.IdempotencyKey
}

// GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestOptions) GetIdempotencyKeyOk() (*string, bool) {
	if o == nil || o.IdempotencyKey == nil {
		return nil, false
	}
	return o.IdempotencyKey, true
}

// HasIdempotencyKey returns a boolean if a field has been set.
func (o *RequestOptions) HasIdempotencyKey() bool {
	if o != nil && o.IdempotencyKey != nil {
		return true
	}

	return false
}

// SetIdempotencyKey gets a reference to the given string and assigns it to the IdempotencyKey field.
func (o *RequestOptions) SetIdempotencyKey(v string) {
	o.IdempotencyKey = &v
}

// GetReadTimeout returns the ReadTimeout field value if set, zero value otherwise.
func (o *RequestOptions) GetReadTimeout() int32 {
	if o == nil || o.ReadTimeout == nil {
		var ret int32
		return ret
	}
	return *o.ReadTimeout
}

// GetReadTimeoutOk returns a tuple with the ReadTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestOptions) GetReadTimeoutOk() (*int32, bool) {
	if o == nil || o.ReadTimeout == nil {
		return nil, false
	}
	return o.ReadTimeout, true
}

// HasReadTimeout returns a boolean if a field has been set.
func (o *RequestOptions) HasReadTimeout() bool {
	if o != nil && o.ReadTimeout != nil {
		return true
	}

	return false
}

// SetReadTimeout gets a reference to the given int32 and assigns it to the ReadTimeout field.
func (o *RequestOptions) SetReadTimeout(v int32) {
	o.ReadTimeout = &v
}

// GetStripeAccount returns the StripeAccount field value if set, zero value otherwise.
func (o *RequestOptions) GetStripeAccount() string {
	if o == nil || o.StripeAccount == nil {
		var ret string
		return ret
	}
	return *o.StripeAccount
}

// GetStripeAccountOk returns a tuple with the StripeAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestOptions) GetStripeAccountOk() (*string, bool) {
	if o == nil || o.StripeAccount == nil {
		return nil, false
	}
	return o.StripeAccount, true
}

// HasStripeAccount returns a boolean if a field has been set.
func (o *RequestOptions) HasStripeAccount() bool {
	if o != nil && o.StripeAccount != nil {
		return true
	}

	return false
}

// SetStripeAccount gets a reference to the given string and assigns it to the StripeAccount field.
func (o *RequestOptions) SetStripeAccount(v string) {
	o.StripeAccount = &v
}

// GetStripeVersion returns the StripeVersion field value if set, zero value otherwise.
func (o *RequestOptions) GetStripeVersion() string {
	if o == nil || o.StripeVersion == nil {
		var ret string
		return ret
	}
	return *o.StripeVersion
}

// GetStripeVersionOk returns a tuple with the StripeVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestOptions) GetStripeVersionOk() (*string, bool) {
	if o == nil || o.StripeVersion == nil {
		return nil, false
	}
	return o.StripeVersion, true
}

// HasStripeVersion returns a boolean if a field has been set.
func (o *RequestOptions) HasStripeVersion() bool {
	if o != nil && o.StripeVersion != nil {
		return true
	}

	return false
}

// SetStripeVersion gets a reference to the given string and assigns it to the StripeVersion field.
func (o *RequestOptions) SetStripeVersion(v string) {
	o.StripeVersion = &v
}

func (o RequestOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiKey != nil {
		toSerialize["apiKey"] = o.ApiKey
	}
	if o.ConnectTimeout != nil {
		toSerialize["connectTimeout"] = o.ConnectTimeout
	}
	if o.IdempotencyKey != nil {
		toSerialize["idempotencyKey"] = o.IdempotencyKey
	}
	if o.ReadTimeout != nil {
		toSerialize["readTimeout"] = o.ReadTimeout
	}
	if o.StripeAccount != nil {
		toSerialize["stripeAccount"] = o.StripeAccount
	}
	if o.StripeVersion != nil {
		toSerialize["stripeVersion"] = o.StripeVersion
	}
	return json.Marshal(toSerialize)
}

type NullableRequestOptions struct {
	value *RequestOptions
	isSet bool
}

func (v NullableRequestOptions) Get() *RequestOptions {
	return v.value
}

func (v *NullableRequestOptions) Set(val *RequestOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestOptions(val *RequestOptions) *NullableRequestOptions {
	return &NullableRequestOptions{value: val, isSet: true}
}

func (v NullableRequestOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
