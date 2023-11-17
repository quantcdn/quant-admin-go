/*
Quant administration API

Testing DomainsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"testing"

	openapiclient "github.com/quantcdn/quant-admin-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_DomainsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DomainsAPIService OrganizationsOrganizationProjectsProjectDomainsDomainDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organization interface{}
		var project interface{}
		var domain interface{}

		resp, httpRes, err := apiClient.DomainsAPI.OrganizationsOrganizationProjectsProjectDomainsDomainDelete(context.Background(), organization, project, domain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DomainsAPIService OrganizationsOrganizationProjectsProjectDomainsDomainGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organization interface{}
		var project interface{}
		var domain interface{}

		resp, httpRes, err := apiClient.DomainsAPI.OrganizationsOrganizationProjectsProjectDomainsDomainGet(context.Background(), organization, project, domain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DomainsAPIService OrganizationsOrganizationProjectsProjectDomainsDomainPatch", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organization interface{}
		var project interface{}
		var domain interface{}

		resp, httpRes, err := apiClient.DomainsAPI.OrganizationsOrganizationProjectsProjectDomainsDomainPatch(context.Background(), organization, project, domain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DomainsAPIService OrganizationsOrganizationProjectsProjectDomainsGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organization interface{}
		var project interface{}

		resp, httpRes, err := apiClient.DomainsAPI.OrganizationsOrganizationProjectsProjectDomainsGet(context.Background(), organization, project).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DomainsAPIService OrganizationsOrganizationProjectsProjectDomainsPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organization interface{}
		var project interface{}

		resp, httpRes, err := apiClient.DomainsAPI.OrganizationsOrganizationProjectsProjectDomainsPost(context.Background(), organization, project).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
