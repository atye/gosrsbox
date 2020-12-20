package client

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/sync/errgroup"
)

// GetJSONFile retrieves the specified file from the Static JSON API.
// This is a good option if you want to dump enitire database .
func (c *client) GetJSONFiles(ctx context.Context, files []string, destinations ...interface{}) error {
	if len(destinations) < len(files) {
		return fmt.Errorf("not enough interfaces provided")
	}

	var eg errgroup.Group
	for i := range destinations {
		i := i
		eg.Go(func() error {
			code, err := c.doJSONDocsRequest(ctx, fmt.Sprintf("%s/%s", c.docsAddress, files[i]), destinations[i])
			if err != nil {
				return err
			}
			if code != http.StatusOK {
				return fmt.Errorf("expected status 200/OK, got: %d", code)
			}
			return nil
		})
	}
	return eg.Wait()
}
