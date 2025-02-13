/*
QuantCDN Admin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package quantadmingo

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


type CrawlersAPI interface {

	/*
	CrawlersCreate Method for CrawlersCreate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param organization
	@param project
	@return CrawlersAPICrawlersCreateRequest
	*/
	CrawlersCreate(ctx context.Context, organization string, project string) CrawlersAPICrawlersCreateRequest

	// CrawlersCreateExecute executes the request
	//  @return Crawler
	CrawlersCreateExecute(r CrawlersAPICrawlersCreateRequest) (*Crawler, *http.Response, error)

	/*
	CrawlersDelete Method for CrawlersDelete

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param organization
	@param project
	@param crawler
	@return CrawlersAPICrawlersDeleteRequest
	*/
	CrawlersDelete(ctx context.Context, organization string, project string, crawler string) CrawlersAPICrawlersDeleteRequest

	// CrawlersDeleteExecute executes the request
	//  @return Crawler
	CrawlersDeleteExecute(r CrawlersAPICrawlersDeleteRequest) (*Crawler, *http.Response, error)

	/*
	CrawlersList Method for CrawlersList

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param organization
	@param project
	@return CrawlersAPICrawlersListRequest
	*/
	CrawlersList(ctx context.Context, organization string, project string) CrawlersAPICrawlersListRequest

	// CrawlersListExecute executes the request
	//  @return []Crawler
	CrawlersListExecute(r CrawlersAPICrawlersListRequest) ([]Crawler, *http.Response, error)

	/*
	CrawlersRead Method for CrawlersRead

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param organization
	@param project
	@param crawler
	@return CrawlersAPICrawlersReadRequest
	*/
	CrawlersRead(ctx context.Context, organization string, project string, crawler string) CrawlersAPICrawlersReadRequest

	// CrawlersReadExecute executes the request
	//  @return Crawler
	CrawlersReadExecute(r CrawlersAPICrawlersReadRequest) (*Crawler, *http.Response, error)

	/*
	CrawlersUpdate Method for CrawlersUpdate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param organization
	@param project
	@param crawler
	@return CrawlersAPICrawlersUpdateRequest
	*/
	CrawlersUpdate(ctx context.Context, organization string, project string, crawler string) CrawlersAPICrawlersUpdateRequest

	// CrawlersUpdateExecute executes the request
	//  @return Crawler
	CrawlersUpdateExecute(r CrawlersAPICrawlersUpdateRequest) (*Crawler, *http.Response, error)
}

// CrawlersAPIService CrawlersAPI service
type CrawlersAPIService service

type CrawlersAPICrawlersCreateRequest struct {
	ctx context.Context
	ApiService CrawlersAPI
	organization string
	project string
	crawlerRequest *CrawlerRequest
}

func (r CrawlersAPICrawlersCreateRequest) CrawlerRequest(crawlerRequest CrawlerRequest) CrawlersAPICrawlersCreateRequest {
	r.crawlerRequest = &crawlerRequest
	return r
}

func (r CrawlersAPICrawlersCreateRequest) Execute() (*Crawler, *http.Response, error) {
	return r.ApiService.CrawlersCreateExecute(r)
}

/*
CrawlersCreate Method for CrawlersCreate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organization
 @param project
 @return CrawlersAPICrawlersCreateRequest
*/
func (a *CrawlersAPIService) CrawlersCreate(ctx context.Context, organization string, project string) CrawlersAPICrawlersCreateRequest {
	return CrawlersAPICrawlersCreateRequest{
		ApiService: a,
		ctx: ctx,
		organization: organization,
		project: project,
	}
}

// Execute executes the request
//  @return Crawler
func (a *CrawlersAPIService) CrawlersCreateExecute(r CrawlersAPICrawlersCreateRequest) (*Crawler, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Crawler
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CrawlersAPIService.CrawlersCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organization}/projects/{project}/crawlers"
	localVarPath = strings.Replace(localVarPath, "{"+"organization"+"}", url.PathEscape(parameterValueToString(r.organization, "organization")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", url.PathEscape(parameterValueToString(r.project, "project")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.crawlerRequest == nil {
		return localVarReturnValue, nil, reportError("crawlerRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.crawlerRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type CrawlersAPICrawlersDeleteRequest struct {
	ctx context.Context
	ApiService CrawlersAPI
	organization string
	project string
	crawler string
}

func (r CrawlersAPICrawlersDeleteRequest) Execute() (*Crawler, *http.Response, error) {
	return r.ApiService.CrawlersDeleteExecute(r)
}

/*
CrawlersDelete Method for CrawlersDelete

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organization
 @param project
 @param crawler
 @return CrawlersAPICrawlersDeleteRequest
*/
func (a *CrawlersAPIService) CrawlersDelete(ctx context.Context, organization string, project string, crawler string) CrawlersAPICrawlersDeleteRequest {
	return CrawlersAPICrawlersDeleteRequest{
		ApiService: a,
		ctx: ctx,
		organization: organization,
		project: project,
		crawler: crawler,
	}
}

// Execute executes the request
//  @return Crawler
func (a *CrawlersAPIService) CrawlersDeleteExecute(r CrawlersAPICrawlersDeleteRequest) (*Crawler, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Crawler
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CrawlersAPIService.CrawlersDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organization}/projects/{project}/crawlers/{crawler}"
	localVarPath = strings.Replace(localVarPath, "{"+"organization"+"}", url.PathEscape(parameterValueToString(r.organization, "organization")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", url.PathEscape(parameterValueToString(r.project, "project")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"crawler"+"}", url.PathEscape(parameterValueToString(r.crawler, "crawler")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type CrawlersAPICrawlersListRequest struct {
	ctx context.Context
	ApiService CrawlersAPI
	organization string
	project string
}

func (r CrawlersAPICrawlersListRequest) Execute() ([]Crawler, *http.Response, error) {
	return r.ApiService.CrawlersListExecute(r)
}

/*
CrawlersList Method for CrawlersList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organization
 @param project
 @return CrawlersAPICrawlersListRequest
*/
func (a *CrawlersAPIService) CrawlersList(ctx context.Context, organization string, project string) CrawlersAPICrawlersListRequest {
	return CrawlersAPICrawlersListRequest{
		ApiService: a,
		ctx: ctx,
		organization: organization,
		project: project,
	}
}

// Execute executes the request
//  @return []Crawler
func (a *CrawlersAPIService) CrawlersListExecute(r CrawlersAPICrawlersListRequest) ([]Crawler, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Crawler
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CrawlersAPIService.CrawlersList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organization}/projects/{project}/crawlers"
	localVarPath = strings.Replace(localVarPath, "{"+"organization"+"}", url.PathEscape(parameterValueToString(r.organization, "organization")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", url.PathEscape(parameterValueToString(r.project, "project")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type CrawlersAPICrawlersReadRequest struct {
	ctx context.Context
	ApiService CrawlersAPI
	organization string
	project string
	crawler string
}

func (r CrawlersAPICrawlersReadRequest) Execute() (*Crawler, *http.Response, error) {
	return r.ApiService.CrawlersReadExecute(r)
}

/*
CrawlersRead Method for CrawlersRead

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organization
 @param project
 @param crawler
 @return CrawlersAPICrawlersReadRequest
*/
func (a *CrawlersAPIService) CrawlersRead(ctx context.Context, organization string, project string, crawler string) CrawlersAPICrawlersReadRequest {
	return CrawlersAPICrawlersReadRequest{
		ApiService: a,
		ctx: ctx,
		organization: organization,
		project: project,
		crawler: crawler,
	}
}

// Execute executes the request
//  @return Crawler
func (a *CrawlersAPIService) CrawlersReadExecute(r CrawlersAPICrawlersReadRequest) (*Crawler, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Crawler
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CrawlersAPIService.CrawlersRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organization}/projects/{project}/crawlers/{crawler}"
	localVarPath = strings.Replace(localVarPath, "{"+"organization"+"}", url.PathEscape(parameterValueToString(r.organization, "organization")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", url.PathEscape(parameterValueToString(r.project, "project")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"crawler"+"}", url.PathEscape(parameterValueToString(r.crawler, "crawler")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type CrawlersAPICrawlersUpdateRequest struct {
	ctx context.Context
	ApiService CrawlersAPI
	organization string
	project string
	crawler string
	crawlerRequestUpdate *CrawlerRequestUpdate
}

func (r CrawlersAPICrawlersUpdateRequest) CrawlerRequestUpdate(crawlerRequestUpdate CrawlerRequestUpdate) CrawlersAPICrawlersUpdateRequest {
	r.crawlerRequestUpdate = &crawlerRequestUpdate
	return r
}

func (r CrawlersAPICrawlersUpdateRequest) Execute() (*Crawler, *http.Response, error) {
	return r.ApiService.CrawlersUpdateExecute(r)
}

/*
CrawlersUpdate Method for CrawlersUpdate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organization
 @param project
 @param crawler
 @return CrawlersAPICrawlersUpdateRequest
*/
func (a *CrawlersAPIService) CrawlersUpdate(ctx context.Context, organization string, project string, crawler string) CrawlersAPICrawlersUpdateRequest {
	return CrawlersAPICrawlersUpdateRequest{
		ApiService: a,
		ctx: ctx,
		organization: organization,
		project: project,
		crawler: crawler,
	}
}

// Execute executes the request
//  @return Crawler
func (a *CrawlersAPIService) CrawlersUpdateExecute(r CrawlersAPICrawlersUpdateRequest) (*Crawler, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Crawler
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CrawlersAPIService.CrawlersUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organization}/projects/{project}/crawlers/{crawler}"
	localVarPath = strings.Replace(localVarPath, "{"+"organization"+"}", url.PathEscape(parameterValueToString(r.organization, "organization")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", url.PathEscape(parameterValueToString(r.project, "project")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"crawler"+"}", url.PathEscape(parameterValueToString(r.crawler, "crawler")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.crawlerRequestUpdate == nil {
		return localVarReturnValue, nil, reportError("crawlerRequestUpdate is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.crawlerRequestUpdate
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
