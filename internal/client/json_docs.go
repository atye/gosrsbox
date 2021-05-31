package client

import (
	"context"
	"encoding/json"
	"fmt"
)

// GetDocument retrieves the specified file from the Static JSON API and
// unmarshals the response into the destination.
func (c *APIClient) GetDocument(ctx context.Context, file string, destination interface{}) error {
	resp, err := c.doDocumentRequest(ctx, fmt.Sprintf("%s/%s", c.docsAddress, file))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(destination)
	if err != nil {
		return err
	}
	return nil
}
