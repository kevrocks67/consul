package consul

import (
	"github.com/hashicorp/consul-net-rpc/net/rpc"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"

	"github.com/hashicorp/consul/agent/pool"
	"github.com/hashicorp/consul/agent/router"
	"github.com/hashicorp/consul/agent/rpc/middleware"
	"github.com/hashicorp/consul/agent/token"
	"github.com/hashicorp/consul/tlsutil"
)

type Deps struct {
	Logger          hclog.InterceptLogger
	TLSConfigurator *tlsutil.Configurator
	Tokens          *token.Store
	Router          *router.Router
	ConnPool        *pool.ConnPool
	GRPCConnPool    GRPCClientConner
	LeaderForwarder LeaderForwarder
	// GetNetRPCInterceptorFunc, if not nil, sets the net/rpc rpc.ServerServiceCallInterceptor on
	// the server side to record metrics around the RPC requests. If nil, no interceptor is added to
	// the rpc server.
	GetNetRPCInterceptorFunc func(recorder *middleware.RequestRecorder) rpc.ServerServiceCallInterceptor
	// NewRequestRecorderFunc provides a middleware.RequestRecorder for the server to use; it cannot be nil
	NewRequestRecorderFunc func(logger hclog.Logger) *middleware.RequestRecorder
	EnterpriseDeps
}

type GRPCClientConner interface {
	ClientConn(datacenter string) (*grpc.ClientConn, error)
	ClientConnLeader() (*grpc.ClientConn, error)
	SetGatewayResolver(func(string) string)
}

type LeaderForwarder interface {
	// UpdateLeaderAddr updates the leader address in the local DC's resolver.
	UpdateLeaderAddr(datacenter, addr string)
}
