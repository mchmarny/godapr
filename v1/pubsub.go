package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	"go.opencensus.io/trace"
)

// PublishWithData invokes to specific topic with the passed in content
func (c *Client) PublishWithData(ctx trace.SpanContext, topic string, in []byte) error {
	if topic == "" {
		return errors.New("nil topic")
	}
	if in == nil {
		return errors.New("nil in content")
	}

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

	return nil
}

// Publish serializes data to JSON and invokes PublishWithData
func (c *Client) Publish(ctx trace.SpanContext, topic string, in interface{}) error {
	b, err := json.Marshal(in)
	if err != nil {
		return errors.Wrapf(err, "error serializing identity: %v", in)
	}
	return c.PublishWithData(ctx, topic, b)
}
