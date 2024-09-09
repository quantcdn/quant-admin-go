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


type RulesRedirectAPI interface {

	/*
	RulesRedirectCreate Method for RulesRedirectCreate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param organization
	@param project
	@return RulesRedirectAPIRulesRedirectCreateRequest
	*/
	RulesRedirectCreate(ctx context.Context, organization string, project string) RulesRedirectAPIRulesRedirectCreateRequest

	// RulesRedirectCreateExecute executes the request
	//  @return RuleRedirect
	RulesRedirectCreateExecute(r RulesRedirectAPIRulesRedirectCreateRequest) (*RuleRedirect, *http.Response, error)

	/*
	RulesRedirectDelete Method for RulesRedirectDelete

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param organization
	@param project
	@param rule
	@return RulesRedirectAPIRulesRedirectDeleteRequest
	*/
	RulesRedirectDelete(ctx context.Context, organization string, project string, rule string) RulesRedirectAPIRulesRedirectDeleteRequest

	// RulesRedirectDeleteExecute executes the request
	//  @return RuleRedirect
	RulesRedirectDeleteExecute(r RulesRedirectAPIRulesRedirectDeleteRequest) (*RuleRedirect, *http.Response, error)

	/*
	RulesRedirectList Method for RulesRedirectList

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param organization
	@param project
	@return RulesRedirectAPIRulesRedirectListRequest
	*/
	RulesRedirectList(ctx context.Context, organization string, project string) RulesRedirectAPIRulesRedirectListRequest

	// RulesRedirectListExecute executes the request
	//  @return []RuleRedirect
	RulesRedirectListExecute(r RulesRedirectAPIRulesRedirectListRequest) ([]RuleRedirect, *http.Response, error)

	/*
	RulesRedirectRead Method for RulesRedirectRead

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param organization
	@param project
	@param rule
	@return RulesRedirectAPIRulesRedirectReadRequest
	*/
	RulesRedirectRead(ctx context.Context, organization string, project string, rule string) RulesRedirectAPIRulesRedirectReadRequest

	// RulesRedirectReadExecute executes the request
	//  @return RuleRedirect
	RulesRedirectReadExecute(r RulesRedirectAPIRulesRedirectReadRequest) (*RuleRedirect, *http.Response, error)

	/*
	RulesRedirectUpdate Method for RulesRedirectUpdate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param organization
	@param project
	@param rule
	@return RulesRedirectAPIRulesRedirectUpdateRequest
	*/
	RulesRedirectUpdate(ctx context.Context, organization string, project string, rule string) RulesRedirectAPIRulesRedirectUpdateRequest

	// RulesRedirectUpdateExecute executes the request
	//  @return RuleRedirect
	RulesRedirectUpdateExecute(r RulesRedirectAPIRulesRedirectUpdateRequest) (*RuleRedirect, *http.Response, error)
}

// RulesRedirectAPIService RulesRedirectAPI service
type RulesRedirectAPIService service

type RulesRedirectAPIRulesRedirectCreateRequest struct {
	ctx context.Context
	ApiService RulesRedirectAPI
	organization string
	project string
	ruleRedirectRequest *RuleRedirectRequest
}

func (r RulesRedirectAPIRulesRedirectCreateRequest) RuleRedirectRequest(ruleRedirectRequest RuleRedirectRequest) RulesRedirectAPIRulesRedirectCreateRequest {
	r.ruleRedirectRequest = &ruleRedirectRequest
	return r
}

func (r RulesRedirectAPIRulesRedirectCreateRequest) Execute() (*RuleRedirect, *http.Response, error) {
	return r.ApiService.RulesRedirectCreateExecute(r)
}

/*
RulesRedirectCreate Method for RulesRedirectCreate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organization
 @param project
 @return RulesRedirectAPIRulesRedirectCreateRequest
*/
func (a *RulesRedirectAPIService) RulesRedirectCreate(ctx context.Context, organization string, project string) RulesRedirectAPIRulesRedirectCreateRequest {
	return RulesRedirectAPIRulesRedirectCreateRequest{
		ApiService: a,
		ctx: ctx,
		organization: organization,
		project: project,
	}
}

// Execute executes the request
//  @return RuleRedirect
func (a *RulesRedirectAPIService) RulesRedirectCreateExecute(r RulesRedirectAPIRulesRedirectCreateRequest) (*RuleRedirect, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RuleRedirect
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RulesRedirectAPIService.RulesRedirectCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organization}/projects/{project}/rules/redirect"
	localVarPath = strings.Replace(localVarPath, "{"+"organization"+"}", url.PathEscape(parameterValueToString(r.organization, "organization")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", url.PathEscape(parameterValueToString(r.project, "project")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.ruleRedirectRequest == nil {
		return localVarReturnValue, nil, reportError("ruleRedirectRequest is required and must be specified")
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
	localVarPostBody = r.ruleRedirectRequest
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

type RulesRedirectAPIRulesRedirectDeleteRequest struct {
	ctx context.Context
	ApiService RulesRedirectAPI
	organization string
	project string
	rule string
}

func (r RulesRedirectAPIRulesRedirectDeleteRequest) Execute() (*RuleRedirect, *http.Response, error) {
	return r.ApiService.RulesRedirectDeleteExecute(r)
}

/*
RulesRedirectDelete Method for RulesRedirectDelete

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organization
 @param project
 @param rule
 @return RulesRedirectAPIRulesRedirectDeleteRequest
*/
func (a *RulesRedirectAPIService) RulesRedirectDelete(ctx context.Context, organization string, project string, rule string) RulesRedirectAPIRulesRedirectDeleteRequest {
	return RulesRedirectAPIRulesRedirectDeleteRequest{
		ApiService: a,
		ctx: ctx,
		organization: organization,
		project: project,
		rule: rule,
	}
}

// Execute executes the request
//  @return RuleRedirect
func (a *RulesRedirectAPIService) RulesRedirectDeleteExecute(r RulesRedirectAPIRulesRedirectDeleteRequest) (*RuleRedirect, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RuleRedirect
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RulesRedirectAPIService.RulesRedirectDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organization}/projects/{project}/rules/redirect/{rule}"
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

type RulesRedirectAPIRulesRedirectListRequest struct {
	ctx context.Context
	ApiService RulesRedirectAPI
	organization string
	project string
}

func (r RulesRedirectAPIRulesRedirectListRequest) Execute() ([]RuleRedirect, *http.Response, error) {
	return r.ApiService.RulesRedirectListExecute(r)
}

/*
RulesRedirectList Method for RulesRedirectList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organization
 @param project
 @return RulesRedirectAPIRulesRedirectListRequest
*/
func (a *RulesRedirectAPIService) RulesRedirectList(ctx context.Context, organization string, project string) RulesRedirectAPIRulesRedirectListRequest {
	return RulesRedirectAPIRulesRedirectListRequest{
		ApiService: a,
		ctx: ctx,
		organization: organization,
		project: project,
	}
}

// Execute executes the request
//  @return []RuleRedirect
func (a *RulesRedirectAPIService) RulesRedirectListExecute(r RulesRedirectAPIRulesRedirectListRequest) ([]RuleRedirect, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []RuleRedirect
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RulesRedirectAPIService.RulesRedirectList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organization}/projects/{project}/rules/redirect"
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

type RulesRedirectAPIRulesRedirectReadRequest struct {
	ctx context.Context
	ApiService RulesRedirectAPI
	organization string
	project string
	rule string
}

func (r RulesRedirectAPIRulesRedirectReadRequest) Execute() (*RuleRedirect, *http.Response, error) {
	return r.ApiService.RulesRedirectReadExecute(r)
}

/*
RulesRedirectRead Method for RulesRedirectRead

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organization
 @param project
 @param rule
 @return RulesRedirectAPIRulesRedirectReadRequest
*/
func (a *RulesRedirectAPIService) RulesRedirectRead(ctx context.Context, organization string, project string, rule string) RulesRedirectAPIRulesRedirectReadRequest {
	return RulesRedirectAPIRulesRedirectReadRequest{
		ApiService: a,
		ctx: ctx,
		organization: organization,
		project: project,
		rule: rule,
	}
}

// Execute executes the request
//  @return RuleRedirect
func (a *RulesRedirectAPIService) RulesRedirectReadExecute(r RulesRedirectAPIRulesRedirectReadRequest) (*RuleRedirect, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RuleRedirect
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RulesRedirectAPIService.RulesRedirectRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organization}/projects/{project}/rules/redirect/{rule}"
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

type RulesRedirectAPIRulesRedirectUpdateRequest struct {
	ctx context.Context
	ApiService RulesRedirectAPI
	organization string
	project string
	rule string
	ruleRedirectRequestUpdate *RuleRedirectRequestUpdate
}

func (r RulesRedirectAPIRulesRedirectUpdateRequest) RuleRedirectRequestUpdate(ruleRedirectRequestUpdate RuleRedirectRequestUpdate) RulesRedirectAPIRulesRedirectUpdateRequest {
	r.ruleRedirectRequestUpdate = &ruleRedirectRequestUpdate
	return r
}

func (r RulesRedirectAPIRulesRedirectUpdateRequest) Execute() (*RuleRedirect, *http.Response, error) {
	return r.ApiService.RulesRedirectUpdateExecute(r)
}

/*
RulesRedirectUpdate Method for RulesRedirectUpdate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organization
 @param project
 @param rule
 @return RulesRedirectAPIRulesRedirectUpdateRequest
*/
func (a *RulesRedirectAPIService) RulesRedirectUpdate(ctx context.Context, organization string, project string, rule string) RulesRedirectAPIRulesRedirectUpdateRequest {
	return RulesRedirectAPIRulesRedirectUpdateRequest{
		ApiService: a,
		ctx: ctx,
		organization: organization,
		project: project,
		rule: rule,
	}
}

// Execute executes the request
//  @return RuleRedirect
func (a *RulesRedirectAPIService) RulesRedirectUpdateExecute(r RulesRedirectAPIRulesRedirectUpdateRequest) (*RuleRedirect, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RuleRedirect
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RulesRedirectAPIService.RulesRedirectUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organization}/projects/{project}/rules/redirect/{rule}"
	localVarPath = strings.Replace(localVarPath, "{"+"organization"+"}", url.PathEscape(parameterValueToString(r.organization, "organization")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", url.PathEscape(parameterValueToString(r.project, "project")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"rule"+"}", url.PathEscape(parameterValueToString(r.rule, "rule")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.ruleRedirectRequestUpdate == nil {
		return localVarReturnValue, nil, reportError("ruleRedirectRequestUpdate is required and must be specified")
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
	localVarPostBody = r.ruleRedirectRequestUpdate
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
