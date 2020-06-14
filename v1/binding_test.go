package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opencensus.io/trace"
)

func TestInvokeBinding(t *testing.T) {
	c := NewClient()
	assert.NotNil(t, c)
	ctx := trace.SpanContext{}

	b, err := c.InvokeBinding(ctx, "", "")
	assert.NotNil(t, err)
	assert.Nil(t, b)

	b, err = c.InvokeBinding(ctx, "binding-name", "")
	assert.NotNil(t, err)
	assert.Nil(t, b)

	b, err = c.InvokeBinding(ctx, "", "operation-name")
	assert.NotNil(t, err)
	assert.Nil(t, b)

	b, err = c.InvokeBindingWithData(ctx, "binding-name", nil)
	assert.NotNil(t, err)
	assert.Nil(t, b)
}
