/*
QuantCDN Admin API

Testing KVAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package quantadmingo

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/quantcdn/quant-admin-go"
)

func Test_quantadmingo_KVAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test KVAPIService KVCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organization string
		var project string

		resp, httpRes, err := apiClient.KVAPI.KVCreate(context.Background(), organization, project).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KVAPIService KVDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organization string
		var project string
		var store string

		resp, httpRes, err := apiClient.KVAPI.KVDelete(context.Background(), organization, project, store).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KVAPIService KVList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organization string
		var project string

		resp, httpRes, err := apiClient.KVAPI.KVList(context.Background(), organization, project).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KVAPIService KVRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organization string
		var project string
		var store string

		resp, httpRes, err := apiClient.KVAPI.KVRead(context.Background(), organization, project, store).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KVAPIService KVUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organization string
		var project string
		var store string

		resp, httpRes, err := apiClient.KVAPI.KVUpdate(context.Background(), organization, project, store).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
