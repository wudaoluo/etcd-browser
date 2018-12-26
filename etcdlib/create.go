package etcdlib

import "context"

//创建key
func (c *client) Create(key, value string) error {
	return c.put(context.Background(), key, value, true)
}

func (c *client) CreateContext(ctx context.Context, key, value string) error {
	return c.put(ctx, key, value, true)
}

//创建dir
func (c *client) CreateDir(key string) error {
	return c.Create(key, c.dirValue)
}

func (c *client) CreateDirContext(ctx context.Context, key string) error {
	return c.CreateContext(ctx, key, c.dirValue)
}
