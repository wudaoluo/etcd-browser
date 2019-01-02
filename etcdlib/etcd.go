package etcdlib

import (
	"context"
	"errors"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/mvcc/mvccpb"
	"strings"
	"time"
)

const (
	DEFAULT_DIR_VALUE = "etcdv3_dir_$2H#%gRe3*t"
)

var (
	ErrorInvalidRootKey = errors.New("root key should not be empty or end with /")
	ErrorInvalidKey     = errors.New("key should start with /")
	ErrorPutKey         = errors.New("key is not under a directory or key is a directory or key is not empty")
	ErrorKeyNotFound    = errors.New("key has not been set")
	ErrorListKey        = errors.New("can only list a directory")
)

type Clienter interface {
	Get(key string) (*Node, error)
	GetContext(ctx context.Context, key string) (*Node, error)

	List(key string) ([]*Node, error)
	ListContext(ctx context.Context, key string) ([]*Node, error)

	Put(key, value string) error
	PutContext(ctx context.Context, key, value string) error

	Create(key, value string) error
	CreateContext(ctx context.Context, key, value string) error

	CreateDir(key string) error
	CreateDirContext(ctx context.Context, key string) error

	Delete(key string) error
	DeleteContext(ctx context.Context, key string) error
}

type client struct {
	keysAPI  *clientv3.Client
	prefix   string //etcd root key
	dirValue string
	timeout  time.Duration
}

func New(endpoint []string, Prefix string) (Clienter, error) {
	cfg := clientv3.Config{
		Endpoints:   endpoint,
		DialTimeout: time.Second * 1,
	}

	c, err := clientv3.New(cfg)

	if err != nil {
		return nil, err
	}

	return &client{
		keysAPI:  c,
		prefix:   Prefix,
		dirValue: DEFAULT_DIR_VALUE,
		timeout:  3 * time.Second,
	}, nil
}

type Node struct {
	*mvccpb.KeyValue
	IsDir bool `json:"is_dir"`
}

func (c *client) createNode(kv *mvccpb.KeyValue) *Node {
	// remove rootKey prefix
	kv.Key = []byte(c.trimRootKey(string(kv.Key)))
	return &Node{
		KeyValue: kv,
		IsDir:    c.isDir(kv.Value),
	}
}

func (c *client) isDir(value []byte) bool {
	return string(value) == c.dirValue
}

func (c *client) trimRootKey(key string) string {
	return strings.TrimPrefix(key, c.prefix)
}


func (c *client) Close() {
	c.Close()
}
