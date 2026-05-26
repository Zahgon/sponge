package rpcclient

import (
	"sync"

	"google.golang.org/grpc"
)

var (
	serverNameExampleConn *grpc.ClientConn
	serverNameExampleOnce sync.Once
)

// NewServerNameExampleRPCConn instantiate rpc client connection
func NewServerNameExampleRPCConn() { _ = "STUB: not implemented"; return }

// if service discovery is not used, connect directly to the rpc service using the ip and port

// using service discovery
//discoverOption, discoveryEndpoint := discoverService(cfg, grpcClientCfg)
//if discoverOption != nil {
//	isUseDiscover = true
//	endpoint = discoveryEndpoint
//	cliOptions = append(cliOptions, discoverOption)
//	cliOptions = append(cliOptions, grpccli.WithEnableLoadBalance()) // load balance
//}

// secure

// token

// GetServerNameExampleRPCConn get client conn
func GetServerNameExampleRPCConn() *grpc.ClientConn { _ = "STUB: not implemented"; return nil }

// CloseServerNameExampleRPCConn Close tears down the ClientConn and all underlying connections.
func CloseServerNameExampleRPCConn() error { _ = "STUB: not implemented"; return nil }

// discovery service with consul or etcd or nacos, select one of them to use
//func discoverService(cfg *config.Config, grpcClientCfg config.GrpcClient) (grpccli.Option, string) {
//	var (
//		endpoint      string
//		grpcCliOption grpccli.Option
//	)
//
//	switch grpcClientCfg.RegistryDiscoveryType {
//	case "consul":
//		endpoint = "discovery:///" + grpcClientCfg.Name // format: discovery:///serverName
//		cli, err := consulcli.Init(cfg.Consul.Addr, consulcli.WithWaitTime(time.Second*5))
//		if err != nil {
//			panic(fmt.Sprintf("consulcli.Init error: %v, addr: %s", err, cfg.Consul.Addr))
//		}
//		iDiscovery := consul.New(cli)
//		grpcCliOption = grpccli.WithDiscovery(iDiscovery)
//
//	case "etcd":
//		endpoint = "discovery:///" + grpcClientCfg.Name // format: discovery:///serverName
//		cli, err := etcdcli.Init(cfg.Etcd.Addrs, etcdcli.WithDialTimeout(time.Second*5))
//		if err != nil {
//			panic(fmt.Sprintf("etcdcli.Init error: %v, addr: %v", err, cfg.Etcd.Addrs))
//		}
//		iDiscovery := etcd.New(cli)
//		grpcCliOption = grpccli.WithDiscovery(iDiscovery)
//
//	case "nacos":
//		endpoint = "discovery:///" + grpcClientCfg.Name + ".grpc" // format: discovery:///serverName.scheme
//		cli, err := nacoscli.NewNamingClient(
//			cfg.NacosRd.IPAddr,
//			cfg.NacosRd.Port,
//			cfg.NacosRd.NamespaceID)
//		if err != nil {
//			panic(fmt.Sprintf("nacoscli.NewNamingClient error: %v, ipAddr: %s, port: %d",
//				err, cfg.NacosRd.IPAddr, cfg.NacosRd.Port))
//		}
//		iDiscovery := nacos.New(cli)
//		grpcCliOption = grpccli.WithDiscovery(iDiscovery)
//	}
//
//	return grpcCliOption, endpoint
//}
