package etcdlib

import (
	"context"
	"crypto/tls"
	"errors"
	e "github.com/wudaoluo/etcd-browser"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/mvcc/mvccpb"
	"go.etcd.io/etcd/pkg/transport"
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
	var tlsConfig *tls.Config
	var err error

	cnf := e.GetConfigInstance()
	if cnf.GetString("cert_file") != "" &&
		cnf.GetString("key_file") != "" &&
		cnf.GetString("ca_file") != "" {

		tlsInfo := transport.TLSInfo{
			CertFile:      cnf.GetString("cert_file"),
			KeyFile:       cnf.GetString("key_file"),
			TrustedCAFile: cnf.GetString("ca_file"),
		}
		tlsConfig, err = tlsInfo.ClientConfig()
		if err != nil {
			return nil, err
		}

	}

	cfg := clientv3.Config{
		Endpoints:   endpoint,
		TLS:         tlsConfig,
		DialTimeout: time.Second * 3,
	}

	c, err := clientv3.New(cfg)

	if err != nil {
		return nil, err
	}

	ctl := &client{
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
