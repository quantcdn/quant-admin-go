/*
QuantCDN Admin API

Testing OrganizationsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/quantcdn/quant-admin-go"
)

func Test_openapi_OrganizationsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OrganizationsAPIService OrganizationsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.OrganizationsAPI.OrganizationsList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrganizationsAPIService OrganizationsRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organization string

		resp, httpRes, err := apiClient.OrganizationsAPI.OrganizationsRead(context.Background(), organization).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
