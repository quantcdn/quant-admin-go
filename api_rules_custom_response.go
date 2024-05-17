/*
QuantCDN Admin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// RulesCustomResponseAPIService RulesCustomResponseAPI service
type RulesCustomResponseAPIService service

type ApiRulesCustomResponseCreateRequest struct {
	ctx context.Context
	ApiService *RulesCustomResponseAPIService
	organization string
	project string
	ruleCustomResponseRequest *RuleCustomResponseRequest
}

func (r ApiRulesCustomResponseCreateRequest) RuleCustomResponseRequest(ruleCustomResponseRequest RuleCustomResponseRequest) ApiRulesCustomResponseCreateRequest {
	r.ruleCustomResponseRequest = &ruleCustomResponseRequest
	return r
}

func (r ApiRulesCustomResponseCreateRequest) Execute() (*Rule, *http.Response, error) {
	return r.ApiService.RulesCustomResponseCreateExecute(r)
}

/*
RulesCustomResponseCreate Method for RulesCustomResponseCreate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organization
 @param project
 @return ApiRulesCustomResponseCreateRequest
*/
func (a *RulesCustomResponseAPIService) RulesCustomResponseCreate(ctx context.Context, organization string, project string) ApiRulesCustomResponseCreateRequest {
	return ApiRulesCustomResponseCreateRequest{
		ApiService: a,
		ctx: ctx,
		organization: organization,
		project: project,
	}
}

// Execute executes the request
//  @return Rule
func (a *RulesCustomResponseAPIService) RulesCustomResponseCreateExecute(r ApiRulesCustomResponseCreateRequest) (*Rule, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Rule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RulesCustomResponseAPIService.RulesCustomResponseCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organization}/projects/{project}/rules/custom-response"
	localVarPath = strings.Replace(localVarPath, "{"+"organization"+"}", url.PathEscape(parameterValueToString(r.organization, "organization")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", url.PathEscape(parameterValueToString(r.project, "project")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.ruleCustomResponseRequest == nil {
		return localVarReturnValue, nil, reportError("ruleCustomResponseRequest is required and must be specified")
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
	localVarPostBody = r.ruleCustomResponseRequest
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

type ApiRulesCustomResponseDeleteRequest struct {
	ctx context.Context
	ApiService *RulesCustomResponseAPIService
	organization string
	project string
	rule string
}

func (r ApiRulesCustomResponseDeleteRequest) Execute() (*Rule, *http.Response, error) {
	return r.ApiService.RulesCustomResponseDeleteExecute(r)
}

/*
RulesCustomResponseDelete Method for RulesCustomResponseDelete

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organization
 @param project
 @param rule
 @return ApiRulesCustomResponseDeleteRequest
*/
func (a *RulesCustomResponseAPIService) RulesCustomResponseDelete(ctx context.Context, organization string, project string, rule string) ApiRulesCustomResponseDeleteRequest {
	return ApiRulesCustomResponseDeleteRequest{
		ApiService: a,
		ctx: ctx,
		organization: organization,
		project: project,
		rule: rule,
	}
}

// Execute executes the request
//  @return Rule
func (a *RulesCustomResponseAPIService) RulesCustomResponseDeleteExecute(r ApiRulesCustomResponseDeleteRequest) (*Rule, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Rule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RulesCustomResponseAPIService.RulesCustomResponseDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organization}/projects/{project}/rules/custom-response/{rule}"
	localVarPath = strings.Replace(localVarPath, "{"+"organization"+"}", url.PathEscape(parameterValueToString(r.organization, "organization")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", url.PathEscape(parameterValueToString(r.project, "project")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"rule"+"}", url.PathEscape(parameterValueToString(r.rule, "rule")), -1)

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

type ApiRulesCustomResponseListRequest struct {
	ctx context.Context
	ApiService *RulesCustomResponseAPIService
	organization string
	project string
}

func (r ApiRulesCustomResponseListRequest) Execute() ([]Rule, *http.Response, error) {
	return r.ApiService.RulesCustomResponseListExecute(r)
}

/*
RulesCustomResponseList Method for RulesCustomResponseList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organization
 @param project
 @return ApiRulesCustomResponseListRequest
*/
func (a *RulesCustomResponseAPIService) RulesCustomResponseList(ctx context.Context, organization string, project string) ApiRulesCustomResponseListRequest {
	return ApiRulesCustomResponseListRequest{
		ApiService: a,
		ctx: ctx,
		organization: organization,
		project: project,
	}
}

// Execute executes the request
//  @return []Rule
func (a *RulesCustomResponseAPIService) RulesCustomResponseListExecute(r ApiRulesCustomResponseListRequest) ([]Rule, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Rule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RulesCustomResponseAPIService.RulesCustomResponseList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organization}/projects/{project}/rules/custom-response"
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

type ApiRulesCustomResponseReadRequest struct {
	ctx context.Context
	ApiService *RulesCustomResponseAPIService
	organization string
	project string
	rule string
}

func (r ApiRulesCustomResponseReadRequest) Execute() (*Rule, *http.Response, error) {
	return r.ApiService.RulesCustomResponseReadExecute(r)
}

/*
RulesCustomResponseRead Method for RulesCustomResponseRead

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organization
 @param project
 @param rule
 @return ApiRulesCustomResponseReadRequest
*/
func (a *RulesCustomResponseAPIService) RulesCustomResponseRead(ctx context.Context, organization string, project string, rule string) ApiRulesCustomResponseReadRequest {
	return ApiRulesCustomResponseReadRequest{
		ApiService: a,
		ctx: ctx,
		organization: organization,
		project: project,
		rule: rule,
	}
}

// Execute executes the request
//  @return Rule
func (a *RulesCustomResponseAPIService) RulesCustomResponseReadExecute(r ApiRulesCustomResponseReadRequest) (*Rule, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Rule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RulesCustomResponseAPIService.RulesCustomResponseRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organization}/projects/{project}/rules/custom-response/{rule}"
	localVarPath = strings.Replace(localVarPath, "{"+"organization"+"}", url.PathEscape(parameterValueToString(r.organization, "organization")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", url.PathEscape(parameterValueToString(r.project, "project")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"rule"+"}", url.PathEscape(parameterValueToString(r.rule, "rule")), -1)

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

type ApiRulesCustomResponseUpdateRequest struct {
	ctx context.Context
	ApiService *RulesCustomResponseAPIService
	organization string
	project string
	rule string
	ruleCustomResponseRequest *RuleCustomResponseRequest
}

func (r ApiRulesCustomResponseUpdateRequest) RuleCustomResponseRequest(ruleCustomResponseRequest RuleCustomResponseRequest) ApiRulesCustomResponseUpdateRequest {
	r.ruleCustomResponseRequest = &ruleCustomResponseRequest
	return r
}

func (r ApiRulesCustomResponseUpdateRequest) Execute() (*Rule, *http.Response, error) {
	return r.ApiService.RulesCustomResponseUpdateExecute(r)
}

/*
RulesCustomResponseUpdate Method for RulesCustomResponseUpdate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organization
 @param project
 @param rule
 @return ApiRulesCustomResponseUpdateRequest
*/
func (a *RulesCustomResponseAPIService) RulesCustomResponseUpdate(ctx context.Context, organization string, project string, rule string) ApiRulesCustomResponseUpdateRequest {
	return ApiRulesCustomResponseUpdateRequest{
		ApiService: a,
		ctx: ctx,
		organization: organization,
		project: project,
		rule: rule,
	}
}

// Execute executes the request
//  @return Rule
func (a *RulesCustomResponseAPIService) RulesCustomResponseUpdateExecute(r ApiRulesCustomResponseUpdateRequest) (*Rule, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Rule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RulesCustomResponseAPIService.RulesCustomResponseUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organization}/projects/{project}/rules/custom-response/{rule}"
	localVarPath = strings.Replace(localVarPath, "{"+"organization"+"}", url.PathEscape(parameterValueToString(r.organization, "organization")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", url.PathEscape(parameterValueToString(r.project, "project")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"rule"+"}", url.PathEscape(parameterValueToString(r.rule, "rule")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.ruleCustomResponseRequest == nil {
		return localVarReturnValue, nil, reportError("ruleCustomResponseRequest is required and must be specified")
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
	localVarPostBody = r.ruleCustomResponseRequest
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
