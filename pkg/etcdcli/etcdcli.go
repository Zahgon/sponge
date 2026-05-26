// Package etcdcli is use for connecting to the etcd service
package etcdcli

import (
	clientv3 "go.etcd.io/etcd/client/v3"
)

// Init connecting to the etcd service
// Note: If the WithConfig(*clientv3.Config) parameter is set, the endpoints parameter is ignored!
func Init(endpoints []string, opts ...Option) (*clientv3.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
