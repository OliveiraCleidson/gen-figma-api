/*
Figma API

Testing WebhooksAPIService

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

func Test_openapi_WebhooksAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test WebhooksAPIService DeleteWebhook", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var webhookId string

		resp, httpRes, err := apiClient.WebhooksAPI.DeleteWebhook(context.Background(), webhookId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WebhooksAPIService GetTeamWebhooks", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var teamId string

		resp, httpRes, err := apiClient.WebhooksAPI.GetTeamWebhooks(context.Background(), teamId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WebhooksAPIService GetWebhook", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var webhookId string

		resp, httpRes, err := apiClient.WebhooksAPI.GetWebhook(context.Background(), webhookId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WebhooksAPIService GetWebhookRequests", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var webhookId string

		resp, httpRes, err := apiClient.WebhooksAPI.GetWebhookRequests(context.Background(), webhookId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WebhooksAPIService PostWebhook", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.WebhooksAPI.PostWebhook(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WebhooksAPIService PutWebhook", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var webhookId string

		resp, httpRes, err := apiClient.WebhooksAPI.PutWebhook(context.Background(), webhookId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}