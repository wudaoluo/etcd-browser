package etcdlib

import (
	"context"
	"crypto/tls"
	"errors"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/mvcc/mvccpb"
	"strings"
	"sync/atomic"
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

	MembersHandler() (interface{}, error)

	Watch(fn func(key,value []byte,revision int64) error)

	Close()
}

type watch struct {
	revision int64
}

func (w *watch) Store(newRevision int64) {
	atomic.StoreInt64(&w.revision,newRevision)
}

type client struct {
	ctx context.Context
	watch    *watch
	keysAPI  *clientv3.Client
	prefix   string //etcd root key
	dirValue string
	timeout  time.Duration

}



func New(ctx context.Context, endpoint []string, Prefix string,tls *tls.Config) (Clienter, error) {
	var err error

	cfg := clientv3.Config{
		Endpoints:   endpoint,
		TLS:         tls,
		DialTimeout: time.Second * 3,
	}

	c, err := clientv3.New(cfg)

	if err != nil {
		return nil, err
	}

	ctl := &client{
		ctx: ctx,
		watch: &watch{revision:0},
		keysAPI:  c,
		prefix:   Prefix,
		dirValue: DEFAULT_DIR_VALUE,
		timeout:  1 * time.Second,
	}

	err = ctl.FormatRootKey() //prefix key 如果不存在就创建它
	if err != nil {
		panic(err)
	}

	return ctl, nil
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

func (c *client) FormatRootKey() error {
	ctx, cancel := context.WithTimeout(context.TODO(), c.timeout)
	defer cancel()
	_, err := c.keysAPI.Put(ctx, c.prefix, c.dirValue)
	return err
}
