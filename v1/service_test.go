package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opencensus.io/trace"
)

func TestInvokeService(t *testing.T) {
	c := NewClient()
	assert.NotNil(t, c)
	ctx := trace.SpanContext{}

	b, err := c.InvokeService(ctx, "", "")
	assert.NotNil(t, err)
	assert.Nil(t, b)

	b, err = c.InvokeService(ctx, "service-name", "")
	assert.NotNil(t, err)
	assert.Nil(t, b)

	b, err = c.InvokeService(ctx, "", "method-name")
	assert.NotNil(t, err)
	assert.Nil(t, b)

	b, err = c.InvokeServiceWithIdentity(ctx, "service-name", "method-name", nil)
	assert.NotNil(t, err)
	assert.Nil(t, b)
}
