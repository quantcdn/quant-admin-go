/*
QuantCDN Admin API

Testing ProjectsAPIService

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

func Test_quant-admin-go_ProjectsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ProjectsAPIService ProjectsCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organization string

		resp, httpRes, err := apiClient.ProjectsAPI.ProjectsCreate(context.Background(), organization).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService ProjectsDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organization string
		var project string

		httpRes, err := apiClient.ProjectsAPI.ProjectsDelete(context.Background(), organization, project).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService ProjectsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organization string

		resp, httpRes, err := apiClient.ProjectsAPI.ProjectsList(context.Background(), organization).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService ProjectsRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organization string
		var project string

		resp, httpRes, err := apiClient.ProjectsAPI.ProjectsRead(context.Background(), organization, project).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService ProjectsUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organization string
		var project string

		resp, httpRes, err := apiClient.ProjectsAPI.ProjectsUpdate(context.Background(), organization, project).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
