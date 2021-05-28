// +build e2e

package test

import (
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetComment(t *testing.T) {
	client := resty.New()
	response, err := client.R().Get(BASE_URL + "/api/comment")

	if err != nil {
		t.Fail()
	}

	assert.Equal(t, 200, response.StatusCode())
}

func TestPostComment(t *testing.T) {
	client := resty.New()
	response, err := client.
		R().
		SetBody(`{"slug": "/", "author": "12345", "body": "hello world"}`).
		Post(BASE_URL + "/api/comment")
	assert.NoError(t, err)

	assert.Equal(t, 200, response.StatusCode())
}