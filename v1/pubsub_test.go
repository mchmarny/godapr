package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opencensus.io/trace"
)

func TestPublish(t *testing.T) {
	c := NewClient()
	assert.NotNil(t, c)
	ctx := trace.SpanContext{}

	err := c.Publish(ctx, "", nil)
	assert.NotNil(t, err)

	err = c.Publish(ctx, "topic-name", nil)
	assert.NotNil(t, err)

	err = c.Publish(ctx, "", &StateData{})
	assert.NotNil(t, err)

	err = c.PublishWithData(ctx, "", nil)
	assert.NotNil(t, err)
}
