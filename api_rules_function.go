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


type RulesFunctionAPI interface {

	/*
	RulesFunctionCreate Method for RulesFunctionCreate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param organization
	@param project
	@return RulesFunctionAPIRulesFunctionCreateRequest
	*/
	RulesFunctionCreate(ctx context.Context, organization string, project string) RulesFunctionAPIRulesFunctionCreateRequest

	// RulesFunctionCreateExecute executes the request
	//  @return RuleFunction
	RulesFunctionCreateExecute(r RulesFunctionAPIRulesFunctionCreateRequest) (*RuleFunction, *http.Response, error)

	/*
	RulesFunctionDelete Method for RulesFunctionDelete

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param organization
	@param project
	@param rule
	@return RulesFunctionAPIRulesFunctionDeleteRequest
	*/
	RulesFunctionDelete(ctx context.Context, organization string, project string, rule string) RulesFunctionAPIRulesFunctionDeleteRequest

	// RulesFunctionDeleteExecute executes the request
	//  @return RuleFunction
	RulesFunctionDeleteExecute(r RulesFunctionAPIRulesFunctionDeleteRequest) (*RuleFunction, *http.Response, error)

	/*
	RulesFunctionList Method for RulesFunctionList

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param organization
	@param project
	@return RulesFunctionAPIRulesFunctionListRequest
	*/
	RulesFunctionList(ctx context.Context, organization string, project string) RulesFunctionAPIRulesFunctionListRequest

	// RulesFunctionListExecute executes the request
	//  @return []RuleFunction
	RulesFunctionListExecute(r RulesFunctionAPIRulesFunctionListRequest) ([]RuleFunction, *http.Response, error)

	/*
	RulesFunctionRead Method for RulesFunctionRead

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param organization
	@param project
	@param rule
	@return RulesFunctionAPIRulesFunctionReadRequest
	*/
	RulesFunctionRead(ctx context.Context, organization string, project string, rule string) RulesFunctionAPIRulesFunctionReadRequest

	// RulesFunctionReadExecute executes the request
	//  @return RuleFunction
	RulesFunctionReadExecute(r RulesFunctionAPIRulesFunctionReadRequest) (*RuleFunction, *http.Response, error)

	/*
	RulesFunctionUpdate Method for RulesFunctionUpdate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param organization
	@param project
	@param rule
	@return RulesFunctionAPIRulesFunctionUpdateRequest
	*/
	RulesFunctionUpdate(ctx context.Context, organization string, project string, rule string) RulesFunctionAPIRulesFunctionUpdateRequest

	// RulesFunctionUpdateExecute executes the request
	//  @return RuleFunction
	RulesFunctionUpdateExecute(r RulesFunctionAPIRulesFunctionUpdateRequest) (*RuleFunction, *http.Response, error)
}

// RulesFunctionAPIService RulesFunctionAPI service
type RulesFunctionAPIService service

type RulesFunctionAPIRulesFunctionCreateRequest struct {
	ctx context.Context
	ApiService RulesFunctionAPI
	organization string
	project string
	ruleFunctionRequest *RuleFunctionRequest
}

func (r RulesFunctionAPIRulesFunctionCreateRequest) RuleFunctionRequest(ruleFunctionRequest RuleFunctionRequest) RulesFunctionAPIRulesFunctionCreateRequest {
	r.ruleFunctionRequest = &ruleFunctionRequest
	return r
}

func (r RulesFunctionAPIRulesFunctionCreateRequest) Execute() (*RuleFunction, *http.Response, error) {
	return r.ApiService.RulesFunctionCreateExecute(r)
}

/*
RulesFunctionCreate Method for RulesFunctionCreate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organization
 @param project
 @return RulesFunctionAPIRulesFunctionCreateRequest
*/
func (a *RulesFunctionAPIService) RulesFunctionCreate(ctx context.Context, organization string, project string) RulesFunctionAPIRulesFunctionCreateRequest {
	return RulesFunctionAPIRulesFunctionCreateRequest{
		ApiService: a,
		ctx: ctx,
		organization: organization,
		project: project,
	}
}

// Execute executes the request
//  @return RuleFunction
func (a *RulesFunctionAPIService) RulesFunctionCreateExecute(r RulesFunctionAPIRulesFunctionCreateRequest) (*RuleFunction, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RuleFunction
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RulesFunctionAPIService.RulesFunctionCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organization}/projects/{project}/rules/function"
	localVarPath = strings.Replace(localVarPath, "{"+"organization"+"}", url.PathEscape(parameterValueToString(r.organization, "organization")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", url.PathEscape(parameterValueToString(r.project, "project")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.ruleFunctionRequest == nil {
		return localVarReturnValue, nil, reportError("ruleFunctionRequest is required and must be specified")
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
	localVarPostBody = r.ruleFunctionRequest
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

type RulesFunctionAPIRulesFunctionDeleteRequest struct {
	ctx context.Context
	ApiService RulesFunctionAPI
	organization string
	project string
	rule string
}

func (r RulesFunctionAPIRulesFunctionDeleteRequest) Execute() (*RuleFunction, *http.Response, error) {
	return r.ApiService.RulesFunctionDeleteExecute(r)
}

/*
RulesFunctionDelete Method for RulesFunctionDelete

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organization
 @param project
 @param rule
 @return RulesFunctionAPIRulesFunctionDeleteRequest
*/
func (a *RulesFunctionAPIService) RulesFunctionDelete(ctx context.Context, organization string, project string, rule string) RulesFunctionAPIRulesFunctionDeleteRequest {
	return RulesFunctionAPIRulesFunctionDeleteRequest{
		ApiService: a,
		ctx: ctx,
		organization: organization,
		project: project,
		rule: rule,
	}
}

// Execute executes the request
//  @return RuleFunction
func (a *RulesFunctionAPIService) RulesFunctionDeleteExecute(r RulesFunctionAPIRulesFunctionDeleteRequest) (*RuleFunction, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RuleFunction
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RulesFunctionAPIService.RulesFunctionDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organization}/projects/{project}/rules/function/{rule}"
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

type RulesFunctionAPIRulesFunctionListRequest struct {
	ctx context.Context
	ApiService RulesFunctionAPI
	organization string
	project string
}

func (r RulesFunctionAPIRulesFunctionListRequest) Execute() ([]RuleFunction, *http.Response, error) {
	return r.ApiService.RulesFunctionListExecute(r)
}

/*
RulesFunctionList Method for RulesFunctionList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organization
 @param project
 @return RulesFunctionAPIRulesFunctionListRequest
*/
func (a *RulesFunctionAPIService) RulesFunctionList(ctx context.Context, organization string, project string) RulesFunctionAPIRulesFunctionListRequest {
	return RulesFunctionAPIRulesFunctionListRequest{
		ApiService: a,
		ctx: ctx,
		organization: organization,
		project: project,
	}
}

// Execute executes the request
//  @return []RuleFunction
func (a *RulesFunctionAPIService) RulesFunctionListExecute(r RulesFunctionAPIRulesFunctionListRequest) ([]RuleFunction, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []RuleFunction
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RulesFunctionAPIService.RulesFunctionList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organization}/projects/{project}/rules/function"
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

type RulesFunctionAPIRulesFunctionReadRequest struct {
	ctx context.Context
	ApiService RulesFunctionAPI
	organization string
	project string
	rule string
}

func (r RulesFunctionAPIRulesFunctionReadRequest) Execute() (*RuleFunction, *http.Response, error) {
	return r.ApiService.RulesFunctionReadExecute(r)
}

/*
RulesFunctionRead Method for RulesFunctionRead

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organization
 @param project
 @param rule
 @return RulesFunctionAPIRulesFunctionReadRequest
*/
func (a *RulesFunctionAPIService) RulesFunctionRead(ctx context.Context, organization string, project string, rule string) RulesFunctionAPIRulesFunctionReadRequest {
	return RulesFunctionAPIRulesFunctionReadRequest{
		ApiService: a,
		ctx: ctx,
		organization: organization,
		project: project,
		rule: rule,
	}
}

// Execute executes the request
//  @return RuleFunction
func (a *RulesFunctionAPIService) RulesFunctionReadExecute(r RulesFunctionAPIRulesFunctionReadRequest) (*RuleFunction, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RuleFunction
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RulesFunctionAPIService.RulesFunctionRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organization}/projects/{project}/rules/function/{rule}"
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

type RulesFunctionAPIRulesFunctionUpdateRequest struct {
	ctx context.Context
	ApiService RulesFunctionAPI
	organization string
	project string
	rule string
	ruleFunctionRequestUpdate *RuleFunctionRequestUpdate
}

func (r RulesFunctionAPIRulesFunctionUpdateRequest) RuleFunctionRequestUpdate(ruleFunctionRequestUpdate RuleFunctionRequestUpdate) RulesFunctionAPIRulesFunctionUpdateRequest {
	r.ruleFunctionRequestUpdate = &ruleFunctionRequestUpdate
	return r
}

func (r RulesFunctionAPIRulesFunctionUpdateRequest) Execute() (*RuleFunction, *http.Response, error) {
	return r.ApiService.RulesFunctionUpdateExecute(r)
}

/*
RulesFunctionUpdate Method for RulesFunctionUpdate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organization
 @param project
 @param rule
 @return RulesFunctionAPIRulesFunctionUpdateRequest
*/
func (a *RulesFunctionAPIService) RulesFunctionUpdate(ctx context.Context, organization string, project string, rule string) RulesFunctionAPIRulesFunctionUpdateRequest {
	return RulesFunctionAPIRulesFunctionUpdateRequest{
		ApiService: a,
		ctx: ctx,
		organization: organization,
		project: project,
		rule: rule,
	}
}

// Execute executes the request
//  @return RuleFunction
func (a *RulesFunctionAPIService) RulesFunctionUpdateExecute(r RulesFunctionAPIRulesFunctionUpdateRequest) (*RuleFunction, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RuleFunction
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RulesFunctionAPIService.RulesFunctionUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organization}/projects/{project}/rules/function/{rule}"
	localVarPath = strings.Replace(localVarPath, "{"+"organization"+"}", url.PathEscape(parameterValueToString(r.organization, "organization")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", url.PathEscape(parameterValueToString(r.project, "project")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"rule"+"}", url.PathEscape(parameterValueToString(r.rule, "rule")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.ruleFunctionRequestUpdate == nil {
		return localVarReturnValue, nil, reportError("ruleFunctionRequestUpdate is required and must be specified")
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
	localVarPostBody = r.ruleFunctionRequestUpdate
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
