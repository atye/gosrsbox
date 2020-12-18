package inmemory

import "golang.org/x/sync/errgroup"

type Option func(c *InMemoryClient) error

func WithInit() Option {
	return func(c *InMemoryClient) error {
		var eg errgroup.Group
		eg.Go(func() error {
			items, err := c.Updater.Items()
			if err != nil {
				return err
			}
			c.Items = items
			return nil
		})
		eg.Go(func() error {
			monsters, err := c.Updater.Monsters()
			if err != nil {
				return err
			}
			c.Monsters = monsters
			return nil
		})
		eg.Go(func() error {
			prayers, err := c.Updater.Prayers()
			if err != nil {
				return err
			}
			c.Prayers = prayers
			return nil
		})

		err := eg.Wait()
		if err != nil {
			return err
		}
		return nil
	}
}

func WithUpdater(updater Updater) Option {
	return func(c *InMemoryClient) error {
		c.Updater = updater
		return nil
	}
}
