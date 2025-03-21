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


type CrawlerSchedulesAPI interface {

	/*
	CrawlerSchedulesCreate Method for CrawlerSchedulesCreate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param organization
	@param project
	@param crawler
	@return CrawlerSchedulesAPICrawlerSchedulesCreateRequest
	*/
	CrawlerSchedulesCreate(ctx context.Context, organization string, project string, crawler string) CrawlerSchedulesAPICrawlerSchedulesCreateRequest

	// CrawlerSchedulesCreateExecute executes the request
	//  @return CrawlerSchedule
	CrawlerSchedulesCreateExecute(r CrawlerSchedulesAPICrawlerSchedulesCreateRequest) (*CrawlerSchedule, *http.Response, error)

	/*
	CrawlerSchedulesDelete Method for CrawlerSchedulesDelete

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param organization
	@param project
	@param crawler
	@param crawlerSchedule
	@return CrawlerSchedulesAPICrawlerSchedulesDeleteRequest
	*/
	CrawlerSchedulesDelete(ctx context.Context, organization string, project string, crawler string, crawlerSchedule string) CrawlerSchedulesAPICrawlerSchedulesDeleteRequest

	// CrawlerSchedulesDeleteExecute executes the request
	//  @return CrawlerSchedule
	CrawlerSchedulesDeleteExecute(r CrawlerSchedulesAPICrawlerSchedulesDeleteRequest) (*CrawlerSchedule, *http.Response, error)

	/*
	CrawlerSchedulesList Method for CrawlerSchedulesList

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param organization
	@param project
	@param crawler
	@return CrawlerSchedulesAPICrawlerSchedulesListRequest
	*/
	CrawlerSchedulesList(ctx context.Context, organization string, project string, crawler string) CrawlerSchedulesAPICrawlerSchedulesListRequest

	// CrawlerSchedulesListExecute executes the request
	//  @return []CrawlerSchedule
	CrawlerSchedulesListExecute(r CrawlerSchedulesAPICrawlerSchedulesListRequest) ([]CrawlerSchedule, *http.Response, error)

	/*
	CrawlerSchedulesRead Method for CrawlerSchedulesRead

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param organization
	@param project
	@param crawler
	@param crawlerSchedule
	@return CrawlerSchedulesAPICrawlerSchedulesReadRequest
	*/
	CrawlerSchedulesRead(ctx context.Context, organization string, project string, crawler string, crawlerSchedule string) CrawlerSchedulesAPICrawlerSchedulesReadRequest

	// CrawlerSchedulesReadExecute executes the request
	//  @return CrawlerSchedule
	CrawlerSchedulesReadExecute(r CrawlerSchedulesAPICrawlerSchedulesReadRequest) (*CrawlerSchedule, *http.Response, error)

	/*
	CrawlerSchedulesUpdate Method for CrawlerSchedulesUpdate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param organization
	@param project
	@param crawler
	@param crawlerSchedule
	@return CrawlerSchedulesAPICrawlerSchedulesUpdateRequest
	*/
	CrawlerSchedulesUpdate(ctx context.Context, organization string, project string, crawler string, crawlerSchedule string) CrawlerSchedulesAPICrawlerSchedulesUpdateRequest

	// CrawlerSchedulesUpdateExecute executes the request
	//  @return CrawlerSchedule
	CrawlerSchedulesUpdateExecute(r CrawlerSchedulesAPICrawlerSchedulesUpdateRequest) (*CrawlerSchedule, *http.Response, error)
}

// CrawlerSchedulesAPIService CrawlerSchedulesAPI service
type CrawlerSchedulesAPIService service

type CrawlerSchedulesAPICrawlerSchedulesCreateRequest struct {
	ctx context.Context
	ApiService CrawlerSchedulesAPI
	organization string
	project string
	crawler string
	crawlerScheduleRequest *CrawlerScheduleRequest
}

func (r CrawlerSchedulesAPICrawlerSchedulesCreateRequest) CrawlerScheduleRequest(crawlerScheduleRequest CrawlerScheduleRequest) CrawlerSchedulesAPICrawlerSchedulesCreateRequest {
	r.crawlerScheduleRequest = &crawlerScheduleRequest
	return r
}

func (r CrawlerSchedulesAPICrawlerSchedulesCreateRequest) Execute() (*CrawlerSchedule, *http.Response, error) {
	return r.ApiService.CrawlerSchedulesCreateExecute(r)
}

/*
CrawlerSchedulesCreate Method for CrawlerSchedulesCreate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organization
 @param project
 @param crawler
 @return CrawlerSchedulesAPICrawlerSchedulesCreateRequest
*/
func (a *CrawlerSchedulesAPIService) CrawlerSchedulesCreate(ctx context.Context, organization string, project string, crawler string) CrawlerSchedulesAPICrawlerSchedulesCreateRequest {
	return CrawlerSchedulesAPICrawlerSchedulesCreateRequest{
		ApiService: a,
		ctx: ctx,
		organization: organization,
		project: project,
		crawler: crawler,
	}
}

// Execute executes the request
//  @return CrawlerSchedule
func (a *CrawlerSchedulesAPIService) CrawlerSchedulesCreateExecute(r CrawlerSchedulesAPICrawlerSchedulesCreateRequest) (*CrawlerSchedule, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CrawlerSchedule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CrawlerSchedulesAPIService.CrawlerSchedulesCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organization}/projects/{project}/crawlers/{crawler}/schedules"
	localVarPath = strings.Replace(localVarPath, "{"+"organization"+"}", url.PathEscape(parameterValueToString(r.organization, "organization")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", url.PathEscape(parameterValueToString(r.project, "project")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"crawler"+"}", url.PathEscape(parameterValueToString(r.crawler, "crawler")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.crawlerScheduleRequest == nil {
		return localVarReturnValue, nil, reportError("crawlerScheduleRequest is required and must be specified")
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
	localVarPostBody = r.crawlerScheduleRequest
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

type CrawlerSchedulesAPICrawlerSchedulesDeleteRequest struct {
	ctx context.Context
	ApiService CrawlerSchedulesAPI
	organization string
	project string
	crawler string
	crawlerSchedule string
}

func (r CrawlerSchedulesAPICrawlerSchedulesDeleteRequest) Execute() (*CrawlerSchedule, *http.Response, error) {
	return r.ApiService.CrawlerSchedulesDeleteExecute(r)
}

/*
CrawlerSchedulesDelete Method for CrawlerSchedulesDelete

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organization
 @param project
 @param crawler
 @param crawlerSchedule
 @return CrawlerSchedulesAPICrawlerSchedulesDeleteRequest
*/
func (a *CrawlerSchedulesAPIService) CrawlerSchedulesDelete(ctx context.Context, organization string, project string, crawler string, crawlerSchedule string) CrawlerSchedulesAPICrawlerSchedulesDeleteRequest {
	return CrawlerSchedulesAPICrawlerSchedulesDeleteRequest{
		ApiService: a,
		ctx: ctx,
		organization: organization,
		project: project,
		crawler: crawler,
		crawlerSchedule: crawlerSchedule,
	}
}

// Execute executes the request
//  @return CrawlerSchedule
func (a *CrawlerSchedulesAPIService) CrawlerSchedulesDeleteExecute(r CrawlerSchedulesAPICrawlerSchedulesDeleteRequest) (*CrawlerSchedule, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CrawlerSchedule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CrawlerSchedulesAPIService.CrawlerSchedulesDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organization}/projects/{project}/crawlers/{crawler}/schedules/{crawler_schedule}"
	localVarPath = strings.Replace(localVarPath, "{"+"organization"+"}", url.PathEscape(parameterValueToString(r.organization, "organization")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", url.PathEscape(parameterValueToString(r.project, "project")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"crawler"+"}", url.PathEscape(parameterValueToString(r.crawler, "crawler")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"crawler_schedule"+"}", url.PathEscape(parameterValueToString(r.crawlerSchedule, "crawlerSchedule")), -1)

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

type CrawlerSchedulesAPICrawlerSchedulesListRequest struct {
	ctx context.Context
	ApiService CrawlerSchedulesAPI
	organization string
	project string
	crawler string
}

func (r CrawlerSchedulesAPICrawlerSchedulesListRequest) Execute() ([]CrawlerSchedule, *http.Response, error) {
	return r.ApiService.CrawlerSchedulesListExecute(r)
}

/*
CrawlerSchedulesList Method for CrawlerSchedulesList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organization
 @param project
 @param crawler
 @return CrawlerSchedulesAPICrawlerSchedulesListRequest
*/
func (a *CrawlerSchedulesAPIService) CrawlerSchedulesList(ctx context.Context, organization string, project string, crawler string) CrawlerSchedulesAPICrawlerSchedulesListRequest {
	return CrawlerSchedulesAPICrawlerSchedulesListRequest{
		ApiService: a,
		ctx: ctx,
		organization: organization,
		project: project,
		crawler: crawler,
	}
}

// Execute executes the request
//  @return []CrawlerSchedule
func (a *CrawlerSchedulesAPIService) CrawlerSchedulesListExecute(r CrawlerSchedulesAPICrawlerSchedulesListRequest) ([]CrawlerSchedule, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []CrawlerSchedule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CrawlerSchedulesAPIService.CrawlerSchedulesList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organization}/projects/{project}/crawlers/{crawler}/schedules"
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

type CrawlerSchedulesAPICrawlerSchedulesReadRequest struct {
	ctx context.Context
	ApiService CrawlerSchedulesAPI
	organization string
	project string
	crawler string
	crawlerSchedule string
}

func (r CrawlerSchedulesAPICrawlerSchedulesReadRequest) Execute() (*CrawlerSchedule, *http.Response, error) {
	return r.ApiService.CrawlerSchedulesReadExecute(r)
}

/*
CrawlerSchedulesRead Method for CrawlerSchedulesRead

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organization
 @param project
 @param crawler
 @param crawlerSchedule
 @return CrawlerSchedulesAPICrawlerSchedulesReadRequest
*/
func (a *CrawlerSchedulesAPIService) CrawlerSchedulesRead(ctx context.Context, organization string, project string, crawler string, crawlerSchedule string) CrawlerSchedulesAPICrawlerSchedulesReadRequest {
	return CrawlerSchedulesAPICrawlerSchedulesReadRequest{
		ApiService: a,
		ctx: ctx,
		organization: organization,
		project: project,
		crawler: crawler,
		crawlerSchedule: crawlerSchedule,
	}
}

// Execute executes the request
//  @return CrawlerSchedule
func (a *CrawlerSchedulesAPIService) CrawlerSchedulesReadExecute(r CrawlerSchedulesAPICrawlerSchedulesReadRequest) (*CrawlerSchedule, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CrawlerSchedule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CrawlerSchedulesAPIService.CrawlerSchedulesRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organization}/projects/{project}/crawlers/{crawler}/schedules/{crawler_schedule}"
	localVarPath = strings.Replace(localVarPath, "{"+"organization"+"}", url.PathEscape(parameterValueToString(r.organization, "organization")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", url.PathEscape(parameterValueToString(r.project, "project")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"crawler"+"}", url.PathEscape(parameterValueToString(r.crawler, "crawler")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"crawler_schedule"+"}", url.PathEscape(parameterValueToString(r.crawlerSchedule, "crawlerSchedule")), -1)

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

type CrawlerSchedulesAPICrawlerSchedulesUpdateRequest struct {
	ctx context.Context
	ApiService CrawlerSchedulesAPI
	organization string
	project string
	crawler string
	crawlerSchedule string
	crawlerScheduleRequestUpdate *CrawlerScheduleRequestUpdate
}

func (r CrawlerSchedulesAPICrawlerSchedulesUpdateRequest) CrawlerScheduleRequestUpdate(crawlerScheduleRequestUpdate CrawlerScheduleRequestUpdate) CrawlerSchedulesAPICrawlerSchedulesUpdateRequest {
	r.crawlerScheduleRequestUpdate = &crawlerScheduleRequestUpdate
	return r
}

func (r CrawlerSchedulesAPICrawlerSchedulesUpdateRequest) Execute() (*CrawlerSchedule, *http.Response, error) {
	return r.ApiService.CrawlerSchedulesUpdateExecute(r)
}

/*
CrawlerSchedulesUpdate Method for CrawlerSchedulesUpdate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organization
 @param project
 @param crawler
 @param crawlerSchedule
 @return CrawlerSchedulesAPICrawlerSchedulesUpdateRequest
*/
func (a *CrawlerSchedulesAPIService) CrawlerSchedulesUpdate(ctx context.Context, organization string, project string, crawler string, crawlerSchedule string) CrawlerSchedulesAPICrawlerSchedulesUpdateRequest {
	return CrawlerSchedulesAPICrawlerSchedulesUpdateRequest{
		ApiService: a,
		ctx: ctx,
		organization: organization,
		project: project,
		crawler: crawler,
		crawlerSchedule: crawlerSchedule,
	}
}

// Execute executes the request
//  @return CrawlerSchedule
func (a *CrawlerSchedulesAPIService) CrawlerSchedulesUpdateExecute(r CrawlerSchedulesAPICrawlerSchedulesUpdateRequest) (*CrawlerSchedule, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CrawlerSchedule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CrawlerSchedulesAPIService.CrawlerSchedulesUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organization}/projects/{project}/crawlers/{crawler}/schedules/{crawler_schedule}"
	localVarPath = strings.Replace(localVarPath, "{"+"organization"+"}", url.PathEscape(parameterValueToString(r.organization, "organization")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", url.PathEscape(parameterValueToString(r.project, "project")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"crawler"+"}", url.PathEscape(parameterValueToString(r.crawler, "crawler")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"crawler_schedule"+"}", url.PathEscape(parameterValueToString(r.crawlerSchedule, "crawlerSchedule")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.crawlerScheduleRequestUpdate == nil {
		return localVarReturnValue, nil, reportError("crawlerScheduleRequestUpdate is required and must be specified")
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
	localVarPostBody = r.crawlerScheduleRequestUpdate
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
