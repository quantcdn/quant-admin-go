/*
QuantCDN Admin API

Testing KVItemsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package quant-admin-go

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/quantcdn/quant-admin-go"
)

func Test_quant-admin-go_KVItemsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test KVItemsAPIService KVItemsCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organization string
		var project string
		var store string

		resp, httpRes, err := apiClient.KVItemsAPI.KVItemsCreate(context.Background(), organization, project, store).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KVItemsAPIService KVItemsDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organization string
		var project string
		var store string
		var item string

		resp, httpRes, err := apiClient.KVItemsAPI.KVItemsDelete(context.Background(), organization, project, store, item).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KVItemsAPIService KVItemsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organization string
		var project string
		var store string

		resp, httpRes, err := apiClient.KVItemsAPI.KVItemsList(context.Background(), organization, project, store).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KVItemsAPIService KVItemsRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organization string
		var project string
		var store string
		var item string

		resp, httpRes, err := apiClient.KVItemsAPI.KVItemsRead(context.Background(), organization, project, store, item).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KVItemsAPIService KVItemsUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organization string
		var project string
		var store string
		var item string

		resp, httpRes, err := apiClient.KVItemsAPI.KVItemsUpdate(context.Background(), organization, project, store, item).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
