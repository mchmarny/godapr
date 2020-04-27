package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClientCreation(t *testing.T) {
	url := "http://localhost:3500"
	c := NewClient()
	assert.NotNil(t, c)
	assert.Equal(t, c.BaseURL, url)

	c1 := NewClientWithURL(url)
	assert.NotNil(t, c1)
	assert.Equal(t, c.BaseURL, url)

	//TODO: Test creation with env var set
}
