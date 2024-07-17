# Go API client for openapi

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.0.0
- Package version: 1.0.0
- Generator version: 7.6.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import openapi "github.com/quantcdn/quant-admin-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `openapi.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `openapi.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `openapi.ContextOperationServerIndices` and `openapi.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://dashboard.quantcdn.io/api/v2*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*CrawlersAPI* | [**CrawlersCreate**](docs/CrawlersAPI.md#crawlerscreate) | **Post** /organizations/{organization}/projects/{project}/crawlers | 
*CrawlersAPI* | [**CrawlersDelete**](docs/CrawlersAPI.md#crawlersdelete) | **Delete** /organizations/{organization}/projects/{project}/crawlers/{crawler} | 
*CrawlersAPI* | [**CrawlersList**](docs/CrawlersAPI.md#crawlerslist) | **Get** /organizations/{organization}/projects/{project}/crawlers | 
*CrawlersAPI* | [**CrawlersRead**](docs/CrawlersAPI.md#crawlersread) | **Get** /organizations/{organization}/projects/{project}/crawlers/{crawler} | 
*CrawlersAPI* | [**CrawlersUpdate**](docs/CrawlersAPI.md#crawlersupdate) | **Patch** /organizations/{organization}/projects/{project}/crawlers/{crawler} | 
*DomainsAPI* | [**DomainsCreate**](docs/DomainsAPI.md#domainscreate) | **Post** /organizations/{organization}/projects/{project}/domains | 
*DomainsAPI* | [**DomainsDelete**](docs/DomainsAPI.md#domainsdelete) | **Delete** /organizations/{organization}/projects/{project}/domains/{domain} | 
*DomainsAPI* | [**DomainsList**](docs/DomainsAPI.md#domainslist) | **Get** /organizations/{organization}/projects/{project}/domains | 
*DomainsAPI* | [**DomainsRead**](docs/DomainsAPI.md#domainsread) | **Get** /organizations/{organization}/projects/{project}/domains/{domain} | 
*DomainsAPI* | [**DomainsUpdate**](docs/DomainsAPI.md#domainsupdate) | **Patch** /organizations/{organization}/projects/{project}/domains/{domain} | 
*HeadersAPI* | [**HeadersCreate**](docs/HeadersAPI.md#headerscreate) | **Post** /organizations/{organization}/projects/{project}/custom-headers | 
*HeadersAPI* | [**HeadersDelete**](docs/HeadersAPI.md#headersdelete) | **Delete** /organizations/{organization}/projects/{project}/custom-headers | 
*HeadersAPI* | [**HeadersList**](docs/HeadersAPI.md#headerslist) | **Get** /organizations/{organization}/projects/{project}/custom-headers | 
*OrganizationsAPI* | [**OrganizationsList**](docs/OrganizationsAPI.md#organizationslist) | **Get** /organizations | 
*OrganizationsAPI* | [**OrganizationsRead**](docs/OrganizationsAPI.md#organizationsread) | **Get** /organizations/{organization} | 
*ProjectsAPI* | [**ProjectsCreate**](docs/ProjectsAPI.md#projectscreate) | **Post** /organizations/{organization}/projects | 
*ProjectsAPI* | [**ProjectsDelete**](docs/ProjectsAPI.md#projectsdelete) | **Delete** /organizations/{organization}/projects/{project} | 
*ProjectsAPI* | [**ProjectsList**](docs/ProjectsAPI.md#projectslist) | **Get** /organizations/{organization}/projects | 
*ProjectsAPI* | [**ProjectsRead**](docs/ProjectsAPI.md#projectsread) | **Get** /organizations/{organization}/projects/{project} | 
*ProjectsAPI* | [**ProjectsUpdate**](docs/ProjectsAPI.md#projectsupdate) | **Patch** /organizations/{organization}/projects/{project} | 
*RulesProxyAPI* | [**RulesProxyCreate**](docs/RulesProxyAPI.md#rulesproxycreate) | **Post** /organizations/{organization}/projects/{project}/rules/proxy | 
*RulesProxyAPI* | [**RulesProxyDelete**](docs/RulesProxyAPI.md#rulesproxydelete) | **Delete** /organizations/{organization}/projects/{project}/rules/proxy/{rule} | 
*RulesProxyAPI* | [**RulesProxyList**](docs/RulesProxyAPI.md#rulesproxylist) | **Get** /organizations/{organization}/projects/{project}/rules/proxy | 
*RulesProxyAPI* | [**RulesProxyRead**](docs/RulesProxyAPI.md#rulesproxyread) | **Get** /organizations/{organization}/projects/{project}/rules/proxy/{rule} | 
*RulesProxyAPI* | [**RulesProxyUpdate**](docs/RulesProxyAPI.md#rulesproxyupdate) | **Patch** /organizations/{organization}/projects/{project}/rules/proxy/{rule} | 
*RulesRedirectAPI* | [**RulesRedirectCreate**](docs/RulesRedirectAPI.md#rulesredirectcreate) | **Post** /organizations/{organization}/projects/{project}/rules/redirect | 
*RulesRedirectAPI* | [**RulesRedirectDelete**](docs/RulesRedirectAPI.md#rulesredirectdelete) | **Delete** /organizations/{organization}/projects/{project}/rules/redirect/{rule} | 
*RulesRedirectAPI* | [**RulesRedirectList**](docs/RulesRedirectAPI.md#rulesredirectlist) | **Get** /organizations/{organization}/projects/{project}/rules/redirect | 
*RulesRedirectAPI* | [**RulesRedirectRead**](docs/RulesRedirectAPI.md#rulesredirectread) | **Get** /organizations/{organization}/projects/{project}/rules/redirect/{rule} | 
*RulesRedirectAPI* | [**RulesRedirectUpdate**](docs/RulesRedirectAPI.md#rulesredirectupdate) | **Patch** /organizations/{organization}/projects/{project}/rules/redirect/{rule} | 


## Documentation For Models

 - [Crawler](docs/Crawler.md)
 - [CrawlerRequest](docs/CrawlerRequest.md)
 - [CrawlerRequestUpdate](docs/CrawlerRequestUpdate.md)
 - [Domain](docs/Domain.md)
 - [DomainRequest](docs/DomainRequest.md)
 - [DomainRequestUpdate](docs/DomainRequestUpdate.md)
 - [Error](docs/Error.md)
 - [HeadersCreateRequest](docs/HeadersCreateRequest.md)
 - [HeadersDeleteRequest](docs/HeadersDeleteRequest.md)
 - [Organization](docs/Organization.md)
 - [Project](docs/Project.md)
 - [ProjectRequest](docs/ProjectRequest.md)
 - [ProjectRequestUpdate](docs/ProjectRequestUpdate.md)
 - [RuleProxy](docs/RuleProxy.md)
 - [RuleProxyAction](docs/RuleProxyAction.md)
 - [RuleProxyActionNotifyConfig](docs/RuleProxyActionNotifyConfig.md)
 - [RuleProxyRequest](docs/RuleProxyRequest.md)
 - [RuleProxyRequestUpdate](docs/RuleProxyRequestUpdate.md)
 - [RuleRedirect](docs/RuleRedirect.md)
 - [RuleRedirectAction](docs/RuleRedirectAction.md)
 - [RuleRedirectRequest](docs/RuleRedirectRequest.md)
 - [RuleRedirectRequestUpdate](docs/RuleRedirectRequestUpdate.md)
 - [Threshold](docs/Threshold.md)
 - [WAFConfig](docs/WAFConfig.md)
 - [WAFConfigBlockLists](docs/WAFConfigBlockLists.md)
 - [WAFConfigHttpbl](docs/WAFConfigHttpbl.md)
 - [WAFConfigUpdate](docs/WAFConfigUpdate.md)
 - [WAFConfigUpdateHttpbl](docs/WAFConfigUpdateHttpbl.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### BearerAuth

- **Type**: HTTP Bearer token authentication

Example

```go
auth := context.WithValue(context.Background(), openapi.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



