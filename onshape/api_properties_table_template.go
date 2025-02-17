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

// PropertiesTableTemplateApiService PropertiesTableTemplateApi service
type PropertiesTableTemplateApiService service

type ApiCreateTableTemplateRequest struct {
	ctx                             context.Context
	ApiService                      *PropertiesTableTemplateApiService
	bTPropertiesTableTemplateParams *BTPropertiesTableTemplateParams
	templateGroupId                 *string
}

func (r ApiCreateTableTemplateRequest) BTPropertiesTableTemplateParams(bTPropertiesTableTemplateParams BTPropertiesTableTemplateParams) ApiCreateTableTemplateRequest {
	r.bTPropertiesTableTemplateParams = &bTPropertiesTableTemplateParams
	return r
}

func (r ApiCreateTableTemplateRequest) TemplateGroupId(templateGroupId string) ApiCreateTableTemplateRequest {
	r.templateGroupId = &templateGroupId
	return r
}

func (r ApiCreateTableTemplateRequest) Execute() (*BTPropertiesTableTemplateInfo, *http.Response, error) {
	return r.ApiService.CreateTableTemplateExecute(r)
}

/*
CreateTableTemplate Create a new properties table template.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateTableTemplateRequest
*/
func (a *PropertiesTableTemplateApiService) CreateTableTemplate(ctx context.Context) ApiCreateTableTemplateRequest {
	return ApiCreateTableTemplateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return BTPropertiesTableTemplateInfo
func (a *PropertiesTableTemplateApiService) CreateTableTemplateExecute(r ApiCreateTableTemplateRequest) (*BTPropertiesTableTemplateInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTPropertiesTableTemplateInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PropertiesTableTemplateApiService.CreateTableTemplate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tabletemplates"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.bTPropertiesTableTemplateParams == nil {
		return localVarReturnValue, nil, reportError("bTPropertiesTableTemplateParams is required and must be specified")
	}

	if r.templateGroupId != nil {
		localVarQueryParams.Add("templateGroupId", parameterToString(*r.templateGroupId, ""))
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
	localVarPostBody = r.bTPropertiesTableTemplateParams
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
		var v BTPropertiesTableTemplateInfo
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

type ApiDeleteTableTemplateRequest struct {
	ctx        context.Context
	ApiService *PropertiesTableTemplateApiService
	tid        string
}

func (r ApiDeleteTableTemplateRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.DeleteTableTemplateExecute(r)
}

/*
DeleteTableTemplate Delete a properties table template by template ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tid The id of the template in which to perform the operation.
 @return ApiDeleteTableTemplateRequest
*/
func (a *PropertiesTableTemplateApiService) DeleteTableTemplate(ctx context.Context, tid string) ApiDeleteTableTemplateRequest {
	return ApiDeleteTableTemplateRequest{
		ApiService: a,
		ctx:        ctx,
		tid:        tid,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *PropertiesTableTemplateApiService) DeleteTableTemplateExecute(r ApiDeleteTableTemplateRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PropertiesTableTemplateApiService.DeleteTableTemplate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tabletemplates/{tid}"
	localVarPath = strings.Replace(localVarPath, "{"+"tid"+"}", url.PathEscape(parameterToString(r.tid, "")), -1)

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

type ApiGetByCompanyIdRequest struct {
	ctx             context.Context
	ApiService      *PropertiesTableTemplateApiService
	cid             string
	templateType    *BTPropertiesTableTemplateType
	onlyActive      *bool
	includeDefaults *bool
}

// Indicates filter for table templates: 0 (BOM) or 1 (Revision Table).
func (r ApiGetByCompanyIdRequest) TemplateType(templateType BTPropertiesTableTemplateType) ApiGetByCompanyIdRequest {
	r.templateType = &templateType
	return r
}

func (r ApiGetByCompanyIdRequest) OnlyActive(onlyActive bool) ApiGetByCompanyIdRequest {
	r.onlyActive = &onlyActive
	return r
}

func (r ApiGetByCompanyIdRequest) IncludeDefaults(includeDefaults bool) ApiGetByCompanyIdRequest {
	r.includeDefaults = &includeDefaults
	return r
}

func (r ApiGetByCompanyIdRequest) Execute() ([]BTPropertiesTableTemplateInfo, *http.Response, error) {
	return r.ApiService.GetByCompanyIdExecute(r)
}

/*
GetByCompanyId Retrieve the properties table templates by company ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param cid The id of the company in which to perform the operation.
 @return ApiGetByCompanyIdRequest
*/
func (a *PropertiesTableTemplateApiService) GetByCompanyId(ctx context.Context, cid string) ApiGetByCompanyIdRequest {
	return ApiGetByCompanyIdRequest{
		ApiService: a,
		ctx:        ctx,
		cid:        cid,
	}
}

// Execute executes the request
//  @return []BTPropertiesTableTemplateInfo
func (a *PropertiesTableTemplateApiService) GetByCompanyIdExecute(r ApiGetByCompanyIdRequest) ([]BTPropertiesTableTemplateInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []BTPropertiesTableTemplateInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PropertiesTableTemplateApiService.GetByCompanyId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tabletemplates/companies/{cid}"
	localVarPath = strings.Replace(localVarPath, "{"+"cid"+"}", url.PathEscape(parameterToString(r.cid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.templateType != nil {
		localVarQueryParams.Add("templateType", parameterToString(*r.templateType, ""))
	}
	if r.onlyActive != nil {
		localVarQueryParams.Add("onlyActive", parameterToString(*r.onlyActive, ""))
	}
	if r.includeDefaults != nil {
		localVarQueryParams.Add("includeDefaults", parameterToString(*r.includeDefaults, ""))
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
		var v []BTPropertiesTableTemplateInfo
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

type ApiGetByDocumentIdRequest struct {
	ctx             context.Context
	ApiService      *PropertiesTableTemplateApiService
	did             string
	templateType    *BTPropertiesTableTemplateType
	onlyActive      *bool
	includeDefaults *bool
}

// Indicates filter for table templates: 0 (BOM) or 1 (Revision Table).
func (r ApiGetByDocumentIdRequest) TemplateType(templateType BTPropertiesTableTemplateType) ApiGetByDocumentIdRequest {
	r.templateType = &templateType
	return r
}

func (r ApiGetByDocumentIdRequest) OnlyActive(onlyActive bool) ApiGetByDocumentIdRequest {
	r.onlyActive = &onlyActive
	return r
}

func (r ApiGetByDocumentIdRequest) IncludeDefaults(includeDefaults bool) ApiGetByDocumentIdRequest {
	r.includeDefaults = &includeDefaults
	return r
}

func (r ApiGetByDocumentIdRequest) Execute() ([]BTPropertiesTableTemplateInfo, *http.Response, error) {
	return r.ApiService.GetByDocumentIdExecute(r)
}

/*
GetByDocumentId Retrieve the properties table templates by document ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param did The id of the document in which to perform the operation.
 @return ApiGetByDocumentIdRequest
*/
func (a *PropertiesTableTemplateApiService) GetByDocumentId(ctx context.Context, did string) ApiGetByDocumentIdRequest {
	return ApiGetByDocumentIdRequest{
		ApiService: a,
		ctx:        ctx,
		did:        did,
	}
}

// Execute executes the request
//  @return []BTPropertiesTableTemplateInfo
func (a *PropertiesTableTemplateApiService) GetByDocumentIdExecute(r ApiGetByDocumentIdRequest) ([]BTPropertiesTableTemplateInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []BTPropertiesTableTemplateInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PropertiesTableTemplateApiService.GetByDocumentId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tabletemplates/d/{did}"
	localVarPath = strings.Replace(localVarPath, "{"+"did"+"}", url.PathEscape(parameterToString(r.did, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.templateType != nil {
		localVarQueryParams.Add("templateType", parameterToString(*r.templateType, ""))
	}
	if r.onlyActive != nil {
		localVarQueryParams.Add("onlyActive", parameterToString(*r.onlyActive, ""))
	}
	if r.includeDefaults != nil {
		localVarQueryParams.Add("includeDefaults", parameterToString(*r.includeDefaults, ""))
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
		var v []BTPropertiesTableTemplateInfo
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

type ApiGetTableTemplateRequest struct {
	ctx        context.Context
	ApiService *PropertiesTableTemplateApiService
	tid        string
}

func (r ApiGetTableTemplateRequest) Execute() (*BTPropertiesTableTemplateInfo, *http.Response, error) {
	return r.ApiService.GetTableTemplateExecute(r)
}

/*
GetTableTemplate Retrieve a properties table template by template ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tid The id of the template in which to perform the operation.
 @return ApiGetTableTemplateRequest
*/
func (a *PropertiesTableTemplateApiService) GetTableTemplate(ctx context.Context, tid string) ApiGetTableTemplateRequest {
	return ApiGetTableTemplateRequest{
		ApiService: a,
		ctx:        ctx,
		tid:        tid,
	}
}

// Execute executes the request
//  @return BTPropertiesTableTemplateInfo
func (a *PropertiesTableTemplateApiService) GetTableTemplateExecute(r ApiGetTableTemplateRequest) (*BTPropertiesTableTemplateInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTPropertiesTableTemplateInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PropertiesTableTemplateApiService.GetTableTemplate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tabletemplates/{tid}"
	localVarPath = strings.Replace(localVarPath, "{"+"tid"+"}", url.PathEscape(parameterToString(r.tid, "")), -1)

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
		var v BTPropertiesTableTemplateInfo
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
