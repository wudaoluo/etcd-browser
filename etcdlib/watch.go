package etcdlib

import (
	"bytes"
	"github.com/ThreeKing2018/goutil/golog"
	"github.com/wudaoluo/etcd-browser/util"
	"go.etcd.io/etcd/clientv3"
	"time"
)

var delim = []byte{byte('/')}

const WATCH_EVENT_DELETE = "DELETE"

func (c *client) Watch(fn func(key, reallyKey, value []byte, revision int64,t string) error) {
	rch := c.keysAPI.Watch(c.ctx, c.prefix, clientv3.WithPrefix(),
		clientv3.WithCreatedNotify(), clientv3.WithCreatedNotify())

	go func() {
		for {

			for wresp := range rch {
				switch {
				case wresp.CompactRevision > c.watch.revision:
					c.watch.Store(wresp.CompactRevision)
					golog.Debugf("Watch to '%s' updated to %d by CompactRevision", c.prefix, wresp.CompactRevision)
				case wresp.Header.Revision > c.watch.revision:
					c.watch.Store(wresp.CompactRevision)
					golog.Debugf("Watch to '%s' updated to %d by header revision", c.prefix, wresp.Header.GetRevision())
				}

				if err := wresp.Err(); err != nil {
					golog.Errorw("watch err")
				}

				for _, ev := range wresp.Events {
					reallyKey := bytes.TrimPrefix(ev.Kv.Key,[]byte(c.prefix))
					if ev.Type.String() == WATCH_EVENT_DELETE  {
						ev.Kv.Key = getParentKey(ev.Kv.Key, delim)
					}

					err := fn(append(ev.Kv.Key, util.Int64ToBytes(ev.Kv.ModRevision)...),
						reallyKey, ev.Kv.Value, ev.Kv.ModRevision,ev.Type.String())

					if err != nil {
						golog.Errorw("faild",
							"type", ev.Type,
							"key", string(ev.Kv.Key),
							"reallyKey",string(reallyKey),
							"value", ev.Kv.Value,
							"createRevisoin", ev.Kv.CreateRevision,
							"modrevision", ev.Kv.ModRevision,
							"err", err)
					} else {
						golog.Debugw("success",
							"type", ev.Type,
							"key", string(ev.Kv.Key),
							"reallyKey",string(reallyKey),
							"value", string(ev.Kv.Value),
							"createRevisoin", ev.Kv.CreateRevision,
							"modrevision", ev.Kv.ModRevision)
					}
				}

			}

			golog.Warnf("Watch to '%s' stopped at revision %d", c.prefix, c.watch.revision)

			time.Sleep(time.Duration(1) * time.Second)
			// Start from next revision so we are not missing anything

			//TODO 需要测试
			if c.watch.revision > 0 {
				rch = c.keysAPI.Watch(c.ctx, c.prefix, clientv3.WithPrefix(),
					clientv3.WithRev(c.watch.revision+1))
			} else {
				// Start from the latest revision
				rch = c.keysAPI.Watch(c.ctx, c.prefix, clientv3.WithPrefix(),
					clientv3.WithCreatedNotify())
			}
		}
	}()

	golog.Warn("退出etcd v3 watch")

}

func getParentKey(key, delim []byte) []byte {
	b := bytes.Split(key, delim)
	if len(b) > 1 {
		key = bytes.Join(b[:len(b)-1], delim)
	}

	return key
}
