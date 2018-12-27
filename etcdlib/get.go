package etcdlib

import "context"

func (c *client) Get(key string) (*Node, error) {
	return c.GetContext(context.Background(), key)
}

func (c *client) GetContext(ctx context.Context, key string) (*Node, error) {
	key, _, err := c.ensureKey(key)
	if err != nil {
		return nil, err
	}

	resp, err := c.keysAPI.Get(context.Background(), key)
	if err != nil {
		return nil, err
	}

	if len(resp.Kvs) == 0 {
		return nil, ErrorKeyNotFound
	}

	//TODO  为什么要用这个
	return c.createNode(resp.Kvs[0]), nil
}
