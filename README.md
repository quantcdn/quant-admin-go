# Go API client for openapi

Provides programmatic interface for projects within QuantCDN

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.2.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import openapi "github.com/quantcdn/quant-admin-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
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

All URIs are relative to *https://dashboard.quantcdn.io/api/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**CreateDomain**](docs/DefaultApi.md#createdomain) | **Post** /domain/create | Create domaim
*DefaultApi* | [**CreateProject**](docs/DefaultApi.md#createproject) | **Post** /project/create | Create project
*DefaultApi* | [**DeleteDomain**](docs/DefaultApi.md#deletedomain) | **Delete** /domain/delete | Delete domain
*DefaultApi* | [**DeleteProject**](docs/DefaultApi.md#deleteproject) | **Delete** /project/delete | Delete project
*DefaultApi* | [**EditProject**](docs/DefaultApi.md#editproject) | **Patch** /project/edit | Edit project
*DefaultApi* | [**GetCrawlRun**](docs/DefaultApi.md#getcrawlrun) | **Get** /crawl/runs/{runId} | Get crawl run status
*DefaultApi* | [**GetCrawlers**](docs/DefaultApi.md#getcrawlers) | **Get** /crawl/configs | Get crawl configs
*DefaultApi* | [**GetOrganisations**](docs/DefaultApi.md#getorganisations) | **Get** /organisations | Get organisations
*DefaultApi* | [**GetProject**](docs/DefaultApi.md#getproject) | **Get** /project | Get project
*DefaultApi* | [**GetProjects**](docs/DefaultApi.md#getprojects) | **Get** /projects | Get projects
*DefaultApi* | [**RunCrawler**](docs/DefaultApi.md#runcrawler) | **Post** /crawl/run/{uuid} | Run crawl config
*DefaultApi* | [**SslRenewDomain**](docs/DefaultApi.md#sslrenewdomain) | **Post** /domain/ssl-renew | Renew domain SSL certificate


## Documentation For Models

 - [CrawlerConfig](docs/CrawlerConfig.md)
 - [CrawlerConfigResponse](docs/CrawlerConfigResponse.md)
 - [CrawlerConfigResponseData](docs/CrawlerConfigResponseData.md)
 - [DomainCreate](docs/DomainCreate.md)
 - [DomainSSLRenew](docs/DomainSSLRenew.md)
 - [Error](docs/Error.md)
 - [GetCrawlRun200Response](docs/GetCrawlRun200Response.md)
 - [GetCrawlRun200ResponseData](docs/GetCrawlRun200ResponseData.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [InlineResponse200Data](docs/InlineResponse200Data.md)
 - [Message](docs/Message.md)
 - [MessageData](docs/MessageData.md)
 - [Organisation](docs/Organisation.md)
 - [Organisations](docs/Organisations.md)
 - [OrganisationsData](docs/OrganisationsData.md)
 - [Project](docs/Project.md)
 - [ProjectConfig](docs/ProjectConfig.md)
 - [ProjectCreate](docs/ProjectCreate.md)
 - [ProjectDomains](docs/ProjectDomains.md)
 - [ProjectEdit](docs/ProjectEdit.md)
 - [ProjectResponse](docs/ProjectResponse.md)
 - [ProjectResponseData](docs/ProjectResponseData.md)
 - [Projects](docs/Projects.md)
 - [ProjectsData](docs/ProjectsData.md)
 - [ProjectsDataProjects](docs/ProjectsDataProjects.md)
 - [RunCrawlerRequest](docs/RunCrawlerRequest.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### ApiTokenAuth

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "BEARER_TOKEN_STRING")
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



