package etcdlib

import (
	"context"
	"go.etcd.io/etcd/clientv3"
)

func (c *client) Put(key, value string) error {
	return c.put(context.Background(), key, value, false)
}

func (c *client) PutContext(ctx context.Context, key, value string) error {
	return c.put(ctx, key, value, false)
}

func (c *client) put(ctx context.Context, key, value string, mustEmpty bool) error {
	// parentKey should be directory
	// key should not be directory
	key, parentKey, err := c.ensureKey(key)
	if err != nil {
		return err
	}

	cmp := []clientv3.Cmp{
		clientv3.Compare(clientv3.Value(parentKey), "=", c.dirValue),
	}

	if mustEmpty {
		//TODO clientv3.Compare 的作用是啥啊
		cmp = append(cmp, clientv3.Compare(clientv3.Version(key), "=", 0))
	} else {
		cmp = append(cmp, clientv3.Compare(clientv3.Value(key), "!=", c.dirValue))
	}

	txn := c.keysAPI.Txn(ctx)
	txn.If(
		cmp...,
	).Then(
		clientv3.OpPut(key, value),
	)

	txnResp, err := txn.Commit()
	if err != nil {
		return err
	}

	if !txnResp.Succeeded {
		return ErrorPutKey
	}


	return nil
}
