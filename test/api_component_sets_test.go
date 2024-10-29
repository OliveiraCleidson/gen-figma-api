/*
Figma API

Testing ComponentSetsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"testing"

	openapiclient "github.com/oliveiracleidson/gen-figma-api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_ComponentSetsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ComponentSetsAPIService GetComponentSet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var key string

		resp, httpRes, err := apiClient.ComponentSetsAPI.GetComponentSet(context.Background(), key).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ComponentSetsAPIService GetFileComponentSets", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var fileKey string

		resp, httpRes, err := apiClient.ComponentSetsAPI.GetFileComponentSets(context.Background(), fileKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ComponentSetsAPIService GetTeamComponentSets", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var teamId string

		resp, httpRes, err := apiClient.ComponentSetsAPI.GetTeamComponentSets(context.Background(), teamId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
