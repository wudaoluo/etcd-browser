package etcdlib

import (
	"context"
	"go.etcd.io/etcd/clientv3"
)

func (c *client) Delete(key string) error {
	return c.DeleteContext(context.Background(), key)
}

func (c *client) DeleteContext(ctx context.Context, key string) error {
	key, _, err := c.ensureKey(key)
	if err != nil {
		return err
	}

	dir := key + "/"

	txn := c.keysAPI.Txn(ctx)

	//TODO 为啥这么写
	txn.If(
		clientv3.Compare(clientv3.Value(key), "=", c.dirValue),
	).Then(
		clientv3.OpDelete(key),
		clientv3.OpDelete(dir, clientv3.WithPrefix()),
	).Else(
		clientv3.OpDelete(key),
	)

	_, err = txn.Commit()
	return err

}
