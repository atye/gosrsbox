package api

import (
	"context"
	"fmt"
)

// GetDocument retrieves the specified file from the Static JSON API
func (c *client) GetDocument(ctx context.Context, file string, destination interface{}) error {
	err := c.doDocumentRequest(ctx, fmt.Sprintf("%s/%s", c.docsAddress, file), destination)
	if err != nil {
		return err
	}
	return nil
}
