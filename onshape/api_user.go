/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.166.18273-3025d52f81b7
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// UserApiService UserApi service
type UserApiService service

type ApiGetUserSettingsRequest struct {
	ctx              context.Context
	ApiService       *UserApiService
	uid              string
	includematerials *bool
}

func (r ApiGetUserSettingsRequest) Includematerials(includematerials bool) ApiGetUserSettingsRequest {
	r.includematerials = &includematerials
	return r
}

func (r ApiGetUserSettingsRequest) Execute() (*BTUserSettingsInfo, *http.Response, error) {
	return r.ApiService.GetUserSettingsExecute(r)
}

/*
GetUserSettings Retrieve user settings by user ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param uid
 @return ApiGetUserSettingsRequest
*/
func (a *UserApiService) GetUserSettings(ctx context.Context, uid string) ApiGetUserSettingsRequest {
	return ApiGetUserSettingsRequest{
		ApiService: a,
		ctx:        ctx,
		uid:        uid,
	}
}

// Execute executes the request
//  @return BTUserSettingsInfo
func (a *UserApiService) GetUserSettingsExecute(r ApiGetUserSettingsRequest) (*BTUserSettingsInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTUserSettingsInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserApiService.GetUserSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{uid}/settings"
	localVarPath = strings.Replace(localVarPath, "{"+"uid"+"}", url.PathEscape(parameterToString(r.uid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.includematerials != nil {
		localVarQueryParams.Add("includematerials", parameterToString(*r.includematerials, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;charset=UTF-8; qs=0.09"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	var _ io.Reader

	if localVarHTTPResponse.StatusCode >= 300 {
		localVarBody, _ := ioutil.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v BTUserSettingsInfo
		err = a.client.decode(&v, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))

	if err != nil {
		localVarBody, _ := ioutil.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetUserSettingsCurrentLoggedInUserRequest struct {
	ctx              context.Context
	ApiService       *UserApiService
	includematerials *bool
}

func (r ApiGetUserSettingsCurrentLoggedInUserRequest) Includematerials(includematerials bool) ApiGetUserSettingsCurrentLoggedInUserRequest {
	r.includematerials = &includematerials
	return r
}

func (r ApiGetUserSettingsCurrentLoggedInUserRequest) Execute() (*BTUserSettingsInfo, *http.Response, error) {
	return r.ApiService.GetUserSettingsCurrentLoggedInUserExecute(r)
}

/*
GetUserSettingsCurrentLoggedInUser Get user settings for the currently signed-in user if there is one, or else return the default settings.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetUserSettingsCurrentLoggedInUserRequest
*/
func (a *UserApiService) GetUserSettingsCurrentLoggedInUser(ctx context.Context) ApiGetUserSettingsCurrentLoggedInUserRequest {
	return ApiGetUserSettingsCurrentLoggedInUserRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return BTUserSettingsInfo
func (a *UserApiService) GetUserSettingsCurrentLoggedInUserExecute(r ApiGetUserSettingsCurrentLoggedInUserRequest) (*BTUserSettingsInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTUserSettingsInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserApiService.GetUserSettingsCurrentLoggedInUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/settings"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.includematerials != nil {
		localVarQueryParams.Add("includematerials", parameterToString(*r.includematerials, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;charset=UTF-8; qs=0.09"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	var _ io.Reader

	if localVarHTTPResponse.StatusCode >= 300 {
		localVarBody, _ := ioutil.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v BTUserSettingsInfo
		err = a.client.decode(&v, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))

	if err != nil {
		localVarBody, _ := ioutil.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSessionRequest struct {
	ctx           context.Context
	ApiService    *UserApiService
	bTLoginParams *BTLoginParams
}

func (r ApiSessionRequest) BTLoginParams(bTLoginParams BTLoginParams) ApiSessionRequest {
	r.bTLoginParams = &bTLoginParams
	return r
}

func (r ApiSessionRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.SessionExecute(r)
}

/*
Session Check if current user is signed-in.Information returned depends on OAuth2ReadPII scope.

Returned information depends on caller's OAuth2ReadPll scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSessionRequest
*/
func (a *UserApiService) Session(ctx context.Context) ApiSessionRequest {
	return ApiSessionRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *UserApiService) SessionExecute(r ApiSessionRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserApiService.Session")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/session"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.bTLoginParams == nil {
		return localVarReturnValue, nil, reportError("bTLoginParams is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json;charset=UTF-8; qs=0.09"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;charset=UTF-8; qs=0.09"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.bTLoginParams
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	var _ io.Reader

	if localVarHTTPResponse.StatusCode >= 300 {
		localVarBody, _ := ioutil.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v map[string]interface{}
		err = a.client.decode(&v, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))

	if err != nil {
		localVarBody, _ := ioutil.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSessionInfoRequest struct {
	ctx        context.Context
	ApiService *UserApiService
}

func (r ApiSessionInfoRequest) Execute() (*BTUserOAuth2SummaryInfo, *http.Response, error) {
	return r.ApiService.SessionInfoExecute(r)
}

/*
SessionInfo Check to see if a user is signed in to a current session.

Returned information depends on caller's OAuth2ReadPll scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSessionInfoRequest
*/
func (a *UserApiService) SessionInfo(ctx context.Context) ApiSessionInfoRequest {
	return ApiSessionInfoRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return BTUserOAuth2SummaryInfo
func (a *UserApiService) SessionInfoExecute(r ApiSessionInfoRequest) (*BTUserOAuth2SummaryInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTUserOAuth2SummaryInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserApiService.SessionInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/sessioninfo"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;charset=UTF-8; qs=0.09"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	var _ io.Reader

	if localVarHTTPResponse.StatusCode >= 300 {
		localVarBody, _ := ioutil.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v BTUserOAuth2SummaryInfo
		err = a.client.decode(&v, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))

	if err != nil {
		localVarBody, _ := ioutil.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
