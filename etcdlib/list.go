package etcdlib

import (
	"context"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/mvcc/mvccpb"
	"strings"
)

func (c *client) List(key string) ([]*Node, error) {
	return c.ListContext(context.Background(), key)
}

// list key 是一个目录
func (c *client) ListContext(ctx context.Context, key string) ([]*Node, error) {
	key, _, err := c.ensureKey(key)
	if err != nil {
		return nil, err
	}

	dir := key + "/"

	txn := c.keysAPI.Txn(ctx)
	txn.If(
		clientv3.Compare(clientv3.Value(key), "=", c.dirValue),
	).Then(
		clientv3.OpGet(dir, clientv3.WithPrefix()),
	)
	txnResp, err := txn.Commit()
	if err != nil {
		return nil, err
	}

	if !txnResp.Succeeded {
		return nil, ErrorListKey
	} else {
		if len(txnResp.Responses) > 0 {
			rangeResp := txnResp.Responses[0].GetResponseRange()
			return c.list(dir, rangeResp.Kvs)
		} else {
			// empty directory
			return []*Node{}, nil
		}
	}
}

// pick key/value under the dir
func (c *client) list(dir string, kvs []*mvccpb.KeyValue) ([]*Node, error) {
	nodes := []*Node{}
	for _, kv := range kvs {
		name := strings.TrimPrefix(string(kv.Key), dir)
		if strings.Contains(name, "/") {
			// secondary directory
			continue
		}
		nodes = append(nodes, c.createNode(kv))
	}
	return nodes, nil
}
