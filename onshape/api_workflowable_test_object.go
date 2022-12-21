/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.157.9035-1a253bfe8c38
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

// WorkflowableTestObjectApiService WorkflowableTestObjectApi service
type WorkflowableTestObjectApiService service

type ApiCreateWorkflowableTestObjectRequest struct {
	ctx        context.Context
	ApiService *WorkflowableTestObjectApiService
	wfid       string
}

func (r ApiCreateWorkflowableTestObjectRequest) Execute() (*BTWorkflowableTestObjectInfo, *http.Response, error) {
	return r.ApiService.CreateWorkflowableTestObjectExecute(r)
}

/*
CreateWorkflowableTestObject Update workflowable test object by workflow ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param wfid
 @return ApiCreateWorkflowableTestObjectRequest
*/
func (a *WorkflowableTestObjectApiService) CreateWorkflowableTestObject(ctx context.Context, wfid string) ApiCreateWorkflowableTestObjectRequest {
	return ApiCreateWorkflowableTestObjectRequest{
		ApiService: a,
		ctx:        ctx,
		wfid:       wfid,
	}
}

// Execute executes the request
//  @return BTWorkflowableTestObjectInfo
func (a *WorkflowableTestObjectApiService) CreateWorkflowableTestObjectExecute(r ApiCreateWorkflowableTestObjectRequest) (*BTWorkflowableTestObjectInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTWorkflowableTestObjectInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkflowableTestObjectApiService.CreateWorkflowableTestObject")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/workflowabletestobject/testobject/{wfid}"
	localVarPath = strings.Replace(localVarPath, "{"+"wfid"+"}", url.PathEscape(parameterToString(r.wfid, "")), -1)

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
		var v BTWorkflowableTestObjectInfo
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

type ApiGetWorkflowableTestObjectRequest struct {
	ctx        context.Context
	ApiService *WorkflowableTestObjectApiService
	oid        string
}

func (r ApiGetWorkflowableTestObjectRequest) Execute() (*BTWorkflowableTestObjectInfo, *http.Response, error) {
	return r.ApiService.GetWorkflowableTestObjectExecute(r)
}

/*
GetWorkflowableTestObject Retrieve workflowable test object by object ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param oid
 @return ApiGetWorkflowableTestObjectRequest
*/
func (a *WorkflowableTestObjectApiService) GetWorkflowableTestObject(ctx context.Context, oid string) ApiGetWorkflowableTestObjectRequest {
	return ApiGetWorkflowableTestObjectRequest{
		ApiService: a,
		ctx:        ctx,
		oid:        oid,
	}
}

// Execute executes the request
//  @return BTWorkflowableTestObjectInfo
func (a *WorkflowableTestObjectApiService) GetWorkflowableTestObjectExecute(r ApiGetWorkflowableTestObjectRequest) (*BTWorkflowableTestObjectInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTWorkflowableTestObjectInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkflowableTestObjectApiService.GetWorkflowableTestObject")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/workflowabletestobject/{oid}"
	localVarPath = strings.Replace(localVarPath, "{"+"oid"+"}", url.PathEscape(parameterToString(r.oid, "")), -1)

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
		var v BTWorkflowableTestObjectInfo
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

type ApiTransitionWorkflowableTestObjectRequest struct {
	ctx        context.Context
	ApiService *WorkflowableTestObjectApiService
	oid        string
	transition string
}

func (r ApiTransitionWorkflowableTestObjectRequest) Execute() (*BTWorkflowableTestObjectInfo, *http.Response, error) {
	return r.ApiService.TransitionWorkflowableTestObjectExecute(r)
}

/*
TransitionWorkflowableTestObject Update workflowable test object transition by object ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param oid
 @param transition
 @return ApiTransitionWorkflowableTestObjectRequest
*/
func (a *WorkflowableTestObjectApiService) TransitionWorkflowableTestObject(ctx context.Context, oid string, transition string) ApiTransitionWorkflowableTestObjectRequest {
	return ApiTransitionWorkflowableTestObjectRequest{
		ApiService: a,
		ctx:        ctx,
		oid:        oid,
		transition: transition,
	}
}

// Execute executes the request
//  @return BTWorkflowableTestObjectInfo
func (a *WorkflowableTestObjectApiService) TransitionWorkflowableTestObjectExecute(r ApiTransitionWorkflowableTestObjectRequest) (*BTWorkflowableTestObjectInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTWorkflowableTestObjectInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkflowableTestObjectApiService.TransitionWorkflowableTestObject")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/workflowabletestobject/{oid}/{transition}"
	localVarPath = strings.Replace(localVarPath, "{"+"oid"+"}", url.PathEscape(parameterToString(r.oid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"transition"+"}", url.PathEscape(parameterToString(r.transition, "")), -1)

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
		var v BTWorkflowableTestObjectInfo
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

type ApiUpdateWorkflowableTestObjectRequest struct {
	ctx                                  context.Context
	ApiService                           *WorkflowableTestObjectApiService
	oid                                  string
	bTUpdateWorkflowableTestObjectParams *BTUpdateWorkflowableTestObjectParams
}

func (r ApiUpdateWorkflowableTestObjectRequest) BTUpdateWorkflowableTestObjectParams(bTUpdateWorkflowableTestObjectParams BTUpdateWorkflowableTestObjectParams) ApiUpdateWorkflowableTestObjectRequest {
	r.bTUpdateWorkflowableTestObjectParams = &bTUpdateWorkflowableTestObjectParams
	return r
}

func (r ApiUpdateWorkflowableTestObjectRequest) Execute() (*BTWorkflowableTestObjectInfo, *http.Response, error) {
	return r.ApiService.UpdateWorkflowableTestObjectExecute(r)
}

/*
UpdateWorkflowableTestObject Update workflowable test object by object ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param oid
 @return ApiUpdateWorkflowableTestObjectRequest
*/
func (a *WorkflowableTestObjectApiService) UpdateWorkflowableTestObject(ctx context.Context, oid string) ApiUpdateWorkflowableTestObjectRequest {
	return ApiUpdateWorkflowableTestObjectRequest{
		ApiService: a,
		ctx:        ctx,
		oid:        oid,
	}
}

// Execute executes the request
//  @return BTWorkflowableTestObjectInfo
func (a *WorkflowableTestObjectApiService) UpdateWorkflowableTestObjectExecute(r ApiUpdateWorkflowableTestObjectRequest) (*BTWorkflowableTestObjectInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTWorkflowableTestObjectInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkflowableTestObjectApiService.UpdateWorkflowableTestObject")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/workflowabletestobject/{oid}"
	localVarPath = strings.Replace(localVarPath, "{"+"oid"+"}", url.PathEscape(parameterToString(r.oid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.bTUpdateWorkflowableTestObjectParams == nil {
		return localVarReturnValue, nil, reportError("bTUpdateWorkflowableTestObjectParams is required and must be specified")
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
	localVarPostBody = r.bTUpdateWorkflowableTestObjectParams
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
		var v BTWorkflowableTestObjectInfo
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
