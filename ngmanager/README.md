# Go API client for ngmanager

This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./ngmanager"
```

## Documentation for API Endpoints

All URIs are relative to *http://app.harness.io*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccountProjectApi* | [**CreateAccountScopedProject**](docs/AccountProjectApi.md#createaccountscopedproject) | **Post** /v1/projects | Create a project
*AccountProjectApi* | [**DeleteAccountScopedProject**](docs/AccountProjectApi.md#deleteaccountscopedproject) | **Delete** /v1/projects/{project} | Delete a project
*AccountProjectApi* | [**GetAccountScopedProject**](docs/AccountProjectApi.md#getaccountscopedproject) | **Get** /v1/projects/{project} | Retrieve a project
*AccountProjectApi* | [**GetAccountScopedProjects**](docs/AccountProjectApi.md#getaccountscopedprojects) | **Get** /v1/projects | List projects
*AccountProjectApi* | [**UpdateAccountScopedProject**](docs/AccountProjectApi.md#updateaccountscopedproject) | **Put** /v1/projects/{project} | Update a project
*AccountSecretApi* | [**CreateAccountScopedSecret**](docs/AccountSecretApi.md#createaccountscopedsecret) | **Post** /v1/secrets | Create a secret
*AccountSecretApi* | [**DeleteAccountScopedSecret**](docs/AccountSecretApi.md#deleteaccountscopedsecret) | **Delete** /v1/secrets/{secret} | Deletes a secret
*AccountSecretApi* | [**GetAccountScopedSecret**](docs/AccountSecretApi.md#getaccountscopedsecret) | **Get** /v1/secrets/{secret} | Retrieve a secret
*AccountSecretApi* | [**GetAccountScopedSecrets**](docs/AccountSecretApi.md#getaccountscopedsecrets) | **Get** /v1/secrets | List secrets
*AccountSecretApi* | [**UpdateAccountScopedSecret**](docs/AccountSecretApi.md#updateaccountscopedsecret) | **Put** /v1/secrets/{secret} | Update a secret
*AccountSecretApi* | [**ValidateUniqueAccountScopedSecretSlug**](docs/AccountSecretApi.md#validateuniqueaccountscopedsecretslug) | **Head** /v1/secrets/{secret} | Validate unique secret slug
*OrgProjectApi* | [**CreateOrgScopedProject**](docs/OrgProjectApi.md#createorgscopedproject) | **Post** /v1/orgs/{org}/projects | Creates a project
*OrgProjectApi* | [**DeleteOrgScopedProject**](docs/OrgProjectApi.md#deleteorgscopedproject) | **Delete** /v1/orgs/{org}/projects/{project} | Delete a project
*OrgProjectApi* | [**GetOrgScopedProject**](docs/OrgProjectApi.md#getorgscopedproject) | **Get** /v1/orgs/{org}/projects/{project} | Retrieve a project
*OrgProjectApi* | [**GetOrgScopedProjects**](docs/OrgProjectApi.md#getorgscopedprojects) | **Get** /v1/orgs/{org}/projects | List projects
*OrgProjectApi* | [**UpdateOrgScopedProject**](docs/OrgProjectApi.md#updateorgscopedproject) | **Put** /v1/orgs/{org}/projects/{project} | Update a project
*OrgSecretApi* | [**CreateOrgScopedSecret**](docs/OrgSecretApi.md#createorgscopedsecret) | **Post** /v1/orgs/{org}/secrets | Create a secret
*OrgSecretApi* | [**DeleteOrgScopedSecret**](docs/OrgSecretApi.md#deleteorgscopedsecret) | **Delete** /v1/org/{org}/secrets/{secret} | Delete a secret
*OrgSecretApi* | [**GetOrgScopedSecret**](docs/OrgSecretApi.md#getorgscopedsecret) | **Get** /v1/org/{org}/secrets/{secret} | Retrieve a secret
*OrgSecretApi* | [**GetOrgScopedSecrets**](docs/OrgSecretApi.md#getorgscopedsecrets) | **Get** /v1/orgs/{org}/secrets | List secrets
*OrgSecretApi* | [**UpdateOrgScopedSecret**](docs/OrgSecretApi.md#updateorgscopedsecret) | **Put** /v1/org/{org}/secrets/{secret} | Update a secret
*OrgSecretApi* | [**ValidateUniqueOrgScopedSecretSlug**](docs/OrgSecretApi.md#validateuniqueorgscopedsecretslug) | **Head** /v1/org/{org}/secrets/{secret} | Validate unique secret slug
*OrganizationApi* | [**CreateOrganization**](docs/OrganizationApi.md#createorganization) | **Post** /v1/orgs | Create an organization
*OrganizationApi* | [**DeleteOrganization**](docs/OrganizationApi.md#deleteorganization) | **Delete** /v1/orgs/{org} | Delete an organization
*OrganizationApi* | [**GetOrganization**](docs/OrganizationApi.md#getorganization) | **Get** /v1/orgs/{org} | Retrieve an organization
*OrganizationApi* | [**GetOrganizations**](docs/OrganizationApi.md#getorganizations) | **Get** /v1/orgs | List organizations
*OrganizationApi* | [**UpdateOrganization**](docs/OrganizationApi.md#updateorganization) | **Put** /v1/orgs/{org} | Update an organization
*ProjectSecretApi* | [**CreateProjectScopedSecret**](docs/ProjectSecretApi.md#createprojectscopedsecret) | **Post** /v1/orgs/{org}/projects/{project}/secrets | Create a secret
*ProjectSecretApi* | [**DeleteProjectScopedSecret**](docs/ProjectSecretApi.md#deleteprojectscopedsecret) | **Delete** /v1/org/{org}/projects/{project}/secrets/{secret} | Delete a secret
*ProjectSecretApi* | [**GetProjectScopedSecret**](docs/ProjectSecretApi.md#getprojectscopedsecret) | **Get** /v1/org/{org}/projects/{project}/secrets/{secret} | Retrieve a secret
*ProjectSecretApi* | [**GetProjectScopedSecrets**](docs/ProjectSecretApi.md#getprojectscopedsecrets) | **Get** /v1/orgs/{org}/projects/{project}/secrets | List secrets
*ProjectSecretApi* | [**UpdateProjectScopedSecret**](docs/ProjectSecretApi.md#updateprojectscopedsecret) | **Put** /v1/org/{org}/projects/{project}/secrets/{secret} | Update a secret
*ProjectSecretApi* | [**ValidateUniqueProjectScopedSecretSlug**](docs/ProjectSecretApi.md#validateuniqueprojectscopedsecretslug) | **Head** /v1/org/{org}/projects/{project}/secrets/{secret} | Validate unique secret slug

## Documentation For Models

 - [CreateOrganizationRequest](docs/CreateOrganizationRequest.md)
 - [CreateProjectRequest](docs/CreateProjectRequest.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [ModuleType](docs/ModuleType.md)
 - [OneOfSecretSpec](docs/OneOfSecretSpec.md)
 - [Organization](docs/Organization.md)
 - [OrganizationResponse](docs/OrganizationResponse.md)
 - [Project](docs/Project.md)
 - [ProjectResponse](docs/ProjectResponse.md)
 - [Secret](docs/Secret.md)
 - [SecretFileSpec](docs/SecretFileSpec.md)
 - [SecretRequest](docs/SecretRequest.md)
 - [SecretResponse](docs/SecretResponse.md)
 - [SecretSpec](docs/SecretSpec.md)
 - [SecretTextSpec](docs/SecretTextSpec.md)
 - [SshKerberosTgtKeyTabFileSpec](docs/SshKerberosTgtKeyTabFileSpec.md)
 - [SshKerberosTgtPasswordSpec](docs/SshKerberosTgtPasswordSpec.md)
 - [SshKeyPathSpec](docs/SshKeyPathSpec.md)
 - [SshKeyReferenceSpec](docs/SshKeyReferenceSpec.md)
 - [SshPasswordSpec](docs/SshPasswordSpec.md)
 - [UpdateOrganizationRequest](docs/UpdateOrganizationRequest.md)
 - [UpdateProjectRequest](docs/UpdateProjectRequest.md)
 - [ValidateSecretSlugResponse](docs/ValidateSecretSlugResponse.md)
 - [WinRmNtlmSpec](docs/WinRmNtlmSpec.md)
 - [WinRmTgtKeyTabFileSpec](docs/WinRmTgtKeyTabFileSpec.md)
 - [WinRmTgtPasswordSpec](docs/WinRmTgtPasswordSpec.md)

## Documentation For Authorization

## x-api-key
- **Type**: API key 

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
	Key: "APIKEY",
	Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```

## Author

contact@harness.io