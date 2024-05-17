/*
QuantCDN Admin API

Testing RulesProxyAPIService

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

func Test_openapi_RulesProxyAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RulesProxyAPIService RulesProxyCreate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organization string
		var project string

		resp, httpRes, err := apiClient.RulesProxyAPI.RulesProxyCreate(context.Background(), organization, project).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RulesProxyAPIService RulesProxyDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organization string
		var project string
		var rule string

		resp, httpRes, err := apiClient.RulesProxyAPI.RulesProxyDelete(context.Background(), organization, project, rule).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RulesProxyAPIService RulesProxyList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organization string
		var project string

		resp, httpRes, err := apiClient.RulesProxyAPI.RulesProxyList(context.Background(), organization, project).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RulesProxyAPIService RulesProxyRead", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organization string
		var project string
		var rule string

		resp, httpRes, err := apiClient.RulesProxyAPI.RulesProxyRead(context.Background(), organization, project, rule).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
