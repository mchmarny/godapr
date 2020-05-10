package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	"go.opencensus.io/trace"
)

// PublishWithData invokes to specific topic with the passed in content
func (c *Client) PublishWithData(ctx context.Context, topic string, in []byte) error {
	ctx, span := trace.StartSpan(ctx, "publish")
	defer span.End()

	url := fmt.Sprintf("%s/v1.0/publish/%s", c.url, topic)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(in))
	if err != nil {
		return errors.Wrapf(err, "error creating publish request: %s", url)
	}

	_, status, err := c.exec(ctx, req)
	if err != nil {
		return errors.Wrapf(err, "error executing: %+v", req)
	}

	if status != http.StatusOK {
		return fmt.Errorf("invalid response code to %s: %d", url, status)
	}

	span.Annotate([]trace.Attribute{
		trace.StringAttribute("topic", topic),
	}, "Published to topic")

	return nil
}

// Publish serializes data to JSON and invokes PublishWithData
func (c *Client) Publish(ctx context.Context, topic string, data interface{}) error {
	b, _ := json.Marshal(data)
	return c.PublishWithData(ctx, topic, b)
}
