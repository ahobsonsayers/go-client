/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.162.14462-13ace71ec1df
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

// WebhookApiService WebhookApi service
type WebhookApiService service

type ApiCreateWebhookRequest struct {
	ctx             context.Context
	ApiService      *WebhookApiService
	bTWebhookParams *BTWebhookParams
}

func (r ApiCreateWebhookRequest) BTWebhookParams(bTWebhookParams BTWebhookParams) ApiCreateWebhookRequest {
	r.bTWebhookParams = &bTWebhookParams
	return r
}

func (r ApiCreateWebhookRequest) Execute() (*BTWebhookInfo, *http.Response, error) {
	return r.ApiService.CreateWebhookExecute(r)
}

/*
CreateWebhook Method for CreateWebhook

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateWebhookRequest
*/
func (a *WebhookApiService) CreateWebhook(ctx context.Context) ApiCreateWebhookRequest {
	return ApiCreateWebhookRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return BTWebhookInfo
func (a *WebhookApiService) CreateWebhookExecute(r ApiCreateWebhookRequest) (*BTWebhookInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTWebhookInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhookApiService.CreateWebhook")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/webhooks"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.bTWebhookParams
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
		var v BTWebhookInfo
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

type ApiGetWebhookRequest struct {
	ctx        context.Context
	ApiService *WebhookApiService
	webhookid  string
}

func (r ApiGetWebhookRequest) Execute() (*BTWebhookInfo, *http.Response, error) {
	return r.ApiService.GetWebhookExecute(r)
}

/*
GetWebhook Retrieve webhook by webhook ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param webhookid
 @return ApiGetWebhookRequest
*/
func (a *WebhookApiService) GetWebhook(ctx context.Context, webhookid string) ApiGetWebhookRequest {
	return ApiGetWebhookRequest{
		ApiService: a,
		ctx:        ctx,
		webhookid:  webhookid,
	}
}

// Execute executes the request
//  @return BTWebhookInfo
func (a *WebhookApiService) GetWebhookExecute(r ApiGetWebhookRequest) (*BTWebhookInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTWebhookInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhookApiService.GetWebhook")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/webhooks/{webhookid}"
	localVarPath = strings.Replace(localVarPath, "{"+"webhookid"+"}", url.PathEscape(parameterToString(r.webhookid, "")), -1)

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
		var v BTWebhookInfo
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

type ApiGetWebhooksRequest struct {
	ctx        context.Context
	ApiService *WebhookApiService
	company    *string
	user       *string
	offset     *int32
	limit      *int32
}

func (r ApiGetWebhooksRequest) Company(company string) ApiGetWebhooksRequest {
	r.company = &company
	return r
}

func (r ApiGetWebhooksRequest) User(user string) ApiGetWebhooksRequest {
	r.user = &user
	return r
}

func (r ApiGetWebhooksRequest) Offset(offset int32) ApiGetWebhooksRequest {
	r.offset = &offset
	return r
}

func (r ApiGetWebhooksRequest) Limit(limit int32) ApiGetWebhooksRequest {
	r.limit = &limit
	return r
}

func (r ApiGetWebhooksRequest) Execute() (*BTListResponseBTWebhookInfo, *http.Response, error) {
	return r.ApiService.GetWebhooksExecute(r)
}

/*
GetWebhooks Retrieve a list of webhooks registered for a company or user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetWebhooksRequest
*/
func (a *WebhookApiService) GetWebhooks(ctx context.Context) ApiGetWebhooksRequest {
	return ApiGetWebhooksRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return BTListResponseBTWebhookInfo
func (a *WebhookApiService) GetWebhooksExecute(r ApiGetWebhooksRequest) (*BTListResponseBTWebhookInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTListResponseBTWebhookInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhookApiService.GetWebhooks")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/webhooks"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.company != nil {
		localVarQueryParams.Add("company", parameterToString(*r.company, ""))
	}
	if r.user != nil {
		localVarQueryParams.Add("user", parameterToString(*r.user, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
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
		var v BTListResponseBTWebhookInfo
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

type ApiPingWebhookRequest struct {
	ctx        context.Context
	ApiService *WebhookApiService
	webhookid  string
}

func (r ApiPingWebhookRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.PingWebhookExecute(r)
}

/*
PingWebhook Method for PingWebhook

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param webhookid
 @return ApiPingWebhookRequest
*/
func (a *WebhookApiService) PingWebhook(ctx context.Context, webhookid string) ApiPingWebhookRequest {
	return ApiPingWebhookRequest{
		ApiService: a,
		ctx:        ctx,
		webhookid:  webhookid,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *WebhookApiService) PingWebhookExecute(r ApiPingWebhookRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhookApiService.PingWebhook")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/webhooks/{webhookid}/ping"
	localVarPath = strings.Replace(localVarPath, "{"+"webhookid"+"}", url.PathEscape(parameterToString(r.webhookid, "")), -1)

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

type ApiUnregisterWebhookRequest struct {
	ctx               context.Context
	ApiService        *WebhookApiService
	webhookid         string
	blockNotification *bool
}

func (r ApiUnregisterWebhookRequest) BlockNotification(blockNotification bool) ApiUnregisterWebhookRequest {
	r.blockNotification = &blockNotification
	return r
}

func (r ApiUnregisterWebhookRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.UnregisterWebhookExecute(r)
}

/*
UnregisterWebhook Delete webhook by webhook ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param webhookid
 @return ApiUnregisterWebhookRequest
*/
func (a *WebhookApiService) UnregisterWebhook(ctx context.Context, webhookid string) ApiUnregisterWebhookRequest {
	return ApiUnregisterWebhookRequest{
		ApiService: a,
		ctx:        ctx,
		webhookid:  webhookid,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *WebhookApiService) UnregisterWebhookExecute(r ApiUnregisterWebhookRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhookApiService.UnregisterWebhook")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/webhooks/{webhookid}"
	localVarPath = strings.Replace(localVarPath, "{"+"webhookid"+"}", url.PathEscape(parameterToString(r.webhookid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.blockNotification != nil {
		localVarQueryParams.Add("blockNotification", parameterToString(*r.blockNotification, ""))
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

type ApiUpdateWebhookRequest struct {
	ctx             context.Context
	ApiService      *WebhookApiService
	webhookid       string
	bTWebhookParams *BTWebhookParams
}

func (r ApiUpdateWebhookRequest) BTWebhookParams(bTWebhookParams BTWebhookParams) ApiUpdateWebhookRequest {
	r.bTWebhookParams = &bTWebhookParams
	return r
}

func (r ApiUpdateWebhookRequest) Execute() (*BTWebhookInfo, *http.Response, error) {
	return r.ApiService.UpdateWebhookExecute(r)
}

/*
UpdateWebhook Method for UpdateWebhook

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param webhookid
 @return ApiUpdateWebhookRequest
*/
func (a *WebhookApiService) UpdateWebhook(ctx context.Context, webhookid string) ApiUpdateWebhookRequest {
	return ApiUpdateWebhookRequest{
		ApiService: a,
		ctx:        ctx,
		webhookid:  webhookid,
	}
}

// Execute executes the request
//  @return BTWebhookInfo
func (a *WebhookApiService) UpdateWebhookExecute(r ApiUpdateWebhookRequest) (*BTWebhookInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTWebhookInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhookApiService.UpdateWebhook")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/webhooks/{webhookid}"
	localVarPath = strings.Replace(localVarPath, "{"+"webhookid"+"}", url.PathEscape(parameterToString(r.webhookid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.bTWebhookParams
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
		var v BTWebhookInfo
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
