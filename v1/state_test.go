package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opencensus.io/trace"
)

func TestSaveState(t *testing.T) {
	c := NewClient()
	assert.NotNil(t, c)
	ctx := trace.SpanContext{}

	err := c.SaveState(ctx, "", "", nil)
	assert.NotNil(t, err)

	err = c.SaveState(ctx, "store-name", "", nil)
	assert.NotNil(t, err)

	err = c.SaveState(ctx, "store-name", "key", nil)
	assert.NotNil(t, err)

	err = c.SaveStateWithData(ctx, "store-name", nil)
	assert.NotNil(t, err)
}

func TestGetState(t *testing.T) {
	c := NewClient()
	assert.NotNil(t, c)
	ctx := trace.SpanContext{}

	b, err := c.GetState(ctx, "", "")
	assert.NotNil(t, err)
	assert.Nil(t, b)

	b, err = c.GetState(ctx, "store-name", "")
	assert.NotNil(t, err)
	assert.Nil(t, b)

	b, err = c.GetState(ctx, "", "key")
	assert.NotNil(t, err)
	assert.Nil(t, b)

	b, err = c.GetStateWithOptions(ctx, "store-name", "key", nil)
	assert.NotNil(t, err)
	assert.Nil(t, b)
}

func TestDeleteState(t *testing.T) {
	c := NewClient()
	assert.NotNil(t, c)
	ctx := trace.SpanContext{}

	err := c.DeleteState(ctx, "", "")
	assert.NotNil(t, err)

	err = c.DeleteState(ctx, "store-name", "")
	assert.NotNil(t, err)

	err = c.DeleteState(ctx, "", "key")
	assert.NotNil(t, err)

	err = c.DeleteStateWithOptions(ctx, "store-name", "key", nil)
	assert.NotNil(t, err)
}
