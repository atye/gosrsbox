package client

import (
	"log"
	"time"

	"golang.org/x/sync/errgroup"
)

type Option func(c *client) error

func WithInit() Option {
	return func(c *client) error {
		var eg errgroup.Group
		eg.Go(func() error {
			items, err := c.source.Items()
			if err != nil {
				return err
			}
			c.items = items
			return nil
		})
		eg.Go(func() error {
			monsters, err := c.source.Monsters()
			if err != nil {
				return err
			}
			c.monsters = monsters
			return nil
		})
		eg.Go(func() error {
			prayers, err := c.source.Prayers()
			if err != nil {
				return err
			}
			c.prayers = prayers
			return nil
		})

		err := eg.Wait()
		if err != nil {
			return err
		}
		return nil
	}
}

func WithSource(source source) Option {
	return func(c *client) error {
		c.source = source
		return nil
	}
}

func WithOptionLogging(logger *log.Logger, option Option, name string) Option {
	return func(c *client) error {
		now := time.Now()
		defer func() {
			logger.Printf("%s took %v", name, time.Since(now))
		}()
		return option(c)
	}
}
