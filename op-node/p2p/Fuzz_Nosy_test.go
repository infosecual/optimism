package p2p

import (
	"context"
	"crypto/ecdsa"
	"io"
	"math/big"
	"net"
	"testing"
	"time"

	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	pubsub_pb "github.com/libp2p/go-libp2p-pubsub/pb"
	"github.com/libp2p/go-libp2p/core/peer"
	go_fuzz_utils "github.com/trailofbits/go-fuzz-utils"
)

func GetTypeProvider(data []byte) (*go_fuzz_utils.TypeProvider, error) {
	tp, err := go_fuzz_utils.NewTypeProvider(data)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsStringBounds(0, 1024)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsSliceBounds(0, 4096)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsBiases(0, 0, 0, 0)
	if err != nil {
		return nil, err
	}
	return tp, nil
}

func Fuzz_Nosy_APIBackend_BlockAddr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *APIBackend
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var ip net.IP
		fill_err = tp.Fill(&ip)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.BlockAddr(_x2, ip)
	})
}

func Fuzz_Nosy_APIBackend_BlockPeer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *APIBackend
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var id peer.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.BlockPeer(_x2, id)
	})
}

func Fuzz_Nosy_APIBackend_BlockSubnet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *APIBackend
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var ipnet *net.IPNet
		fill_err = tp.Fill(&ipnet)
		if fill_err != nil {
			return
		}
		if s == nil || ipnet == nil {
			return
		}

		s.BlockSubnet(_x2, ipnet)
	})
}

func Fuzz_Nosy_APIBackend_ConnectPeer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *APIBackend
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var addr string
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.ConnectPeer(ctx, addr)
	})
}

func Fuzz_Nosy_APIBackend_DisconnectPeer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *APIBackend
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var id peer.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.DisconnectPeer(_x2, id)
	})
}

func Fuzz_Nosy_APIBackend_DiscoveryTable__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *APIBackend
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.DiscoveryTable(_x2)
	})
}

func Fuzz_Nosy_APIBackend_ListBlockedAddrs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *APIBackend
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.ListBlockedAddrs(_x2)
	})
}

func Fuzz_Nosy_APIBackend_ListBlockedPeers__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *APIBackend
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.ListBlockedPeers(_x2)
	})
}

func Fuzz_Nosy_APIBackend_ListBlockedSubnets__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *APIBackend
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.ListBlockedSubnets(_x2)
	})
}

func Fuzz_Nosy_APIBackend_PeerStats__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *APIBackend
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.PeerStats(_x2)
	})
}

func Fuzz_Nosy_APIBackend_Peers__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *APIBackend
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var connected bool
		fill_err = tp.Fill(&connected)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Peers(ctx, connected)
	})
}

func Fuzz_Nosy_APIBackend_ProtectPeer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *APIBackend
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var id peer.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.ProtectPeer(_x2, id)
	})
}

func Fuzz_Nosy_APIBackend_Self__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *APIBackend
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Self(ctx)
	})
}

func Fuzz_Nosy_APIBackend_UnblockAddr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *APIBackend
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var ip net.IP
		fill_err = tp.Fill(&ip)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.UnblockAddr(_x2, ip)
	})
}

func Fuzz_Nosy_APIBackend_UnblockPeer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *APIBackend
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var id peer.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.UnblockPeer(_x2, id)
	})
}

func Fuzz_Nosy_APIBackend_UnblockSubnet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *APIBackend
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var ipnet *net.IPNet
		fill_err = tp.Fill(&ipnet)
		if fill_err != nil {
			return
		}
		if s == nil || ipnet == nil {
			return
		}

		s.UnblockSubnet(_x2, ipnet)
	})
}

func Fuzz_Nosy_APIBackend_UnprotectPeer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *APIBackend
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var id peer.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.UnprotectPeer(_x2, id)
	})
}

func Fuzz_Nosy_Client_BlockAddr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var ip net.IP
		fill_err = tp.Fill(&ip)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c1 := NewClient(c)
		c1.BlockAddr(ctx, ip)
	})
}

func Fuzz_Nosy_Client_BlockPeer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var p peer.ID
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c1 := NewClient(c)
		c1.BlockPeer(ctx, p)
	})
}

func Fuzz_Nosy_Client_BlockSubnet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var ipnet *net.IPNet
		fill_err = tp.Fill(&ipnet)
		if fill_err != nil {
			return
		}
		if c == nil || ipnet == nil {
			return
		}

		c1 := NewClient(c)
		c1.BlockSubnet(ctx, ipnet)
	})
}

func Fuzz_Nosy_Client_ConnectPeer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var addr string
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c1 := NewClient(c)
		c1.ConnectPeer(ctx, addr)
	})
}

func Fuzz_Nosy_Client_DisconnectPeer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var id peer.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c1 := NewClient(c)
		c1.DisconnectPeer(ctx, id)
	})
}

func Fuzz_Nosy_Client_DiscoveryTable__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c1 := NewClient(c)
		c1.DiscoveryTable(ctx)
	})
}

func Fuzz_Nosy_Client_ListBlockedAddrs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c1 := NewClient(c)
		c1.ListBlockedAddrs(ctx)
	})
}

func Fuzz_Nosy_Client_ListBlockedPeers__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c1 := NewClient(c)
		c1.ListBlockedPeers(ctx)
	})
}

func Fuzz_Nosy_Client_ListBlockedSubnets__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c1 := NewClient(c)
		c1.ListBlockedSubnets(ctx)
	})
}

func Fuzz_Nosy_Client_PeerStats__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c1 := NewClient(c)
		c1.PeerStats(ctx)
	})
}

func Fuzz_Nosy_Client_Peers__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var connected bool
		fill_err = tp.Fill(&connected)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c1 := NewClient(c)
		c1.Peers(ctx, connected)
	})
}

func Fuzz_Nosy_Client_ProtectPeer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var p peer.ID
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c1 := NewClient(c)
		c1.ProtectPeer(ctx, p)
	})
}

func Fuzz_Nosy_Client_Self__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c1 := NewClient(c)
		c1.Self(ctx)
	})
}

func Fuzz_Nosy_Client_UnblockAddr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var ip net.IP
		fill_err = tp.Fill(&ip)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c1 := NewClient(c)
		c1.UnblockAddr(ctx, ip)
	})
}

func Fuzz_Nosy_Client_UnblockPeer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var p peer.ID
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c1 := NewClient(c)
		c1.UnblockPeer(ctx, p)
	})
}

func Fuzz_Nosy_Client_UnblockSubnet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var ipnet *net.IPNet
		fill_err = tp.Fill(&ipnet)
		if fill_err != nil {
			return
		}
		if c == nil || ipnet == nil {
			return
		}

		c1 := NewClient(c)
		c1.UnblockSubnet(ctx, ipnet)
	})
}

func Fuzz_Nosy_Client_UnprotectPeer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var p peer.ID
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c1 := NewClient(c)
		c1.UnprotectPeer(ctx, p)
	})
}

func Fuzz_Nosy_Config_BanDuration__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var conf *Config
		fill_err = tp.Fill(&conf)
		if fill_err != nil {
			return
		}
		if conf == nil {
			return
		}

		conf.BanDuration()
	})
}

func Fuzz_Nosy_Config_BanPeers__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var conf *Config
		fill_err = tp.Fill(&conf)
		if fill_err != nil {
			return
		}
		if conf == nil {
			return
		}

		conf.BanPeers()
	})
}

func Fuzz_Nosy_Config_BanThreshold__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var conf *Config
		fill_err = tp.Fill(&conf)
		if fill_err != nil {
			return
		}
		if conf == nil {
			return
		}

		conf.BanThreshold()
	})
}

func Fuzz_Nosy_Config_Check__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var conf *Config
		fill_err = tp.Fill(&conf)
		if fill_err != nil {
			return
		}
		if conf == nil {
			return
		}

		conf.Check()
	})
}

func Fuzz_Nosy_Config_ConfigureGossip__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Config
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var rollupCfg *rollup.Config
		fill_err = tp.Fill(&rollupCfg)
		if fill_err != nil {
			return
		}
		if p == nil || rollupCfg == nil {
			return
		}

		p.ConfigureGossip(rollupCfg)
	})
}

func Fuzz_Nosy_Config_Disabled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var conf *Config
		fill_err = tp.Fill(&conf)
		if fill_err != nil {
			return
		}
		if conf == nil {
			return
		}

		conf.Disabled()
	})
}

// skipping Fuzz_Nosy_Config_Discovery__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

// skipping Fuzz_Nosy_Config_Host__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

func Fuzz_Nosy_Config_PeerScoringParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var conf *Config
		fill_err = tp.Fill(&conf)
		if fill_err != nil {
			return
		}
		if conf == nil {
			return
		}

		conf.PeerScoringParams()
	})
}

func Fuzz_Nosy_Config_ReqRespSyncEnabled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var conf *Config
		fill_err = tp.Fill(&conf)
		if fill_err != nil {
			return
		}
		if conf == nil {
			return
		}

		conf.ReqRespSyncEnabled()
	})
}

func Fuzz_Nosy_Config_TargetPeers__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var conf *Config
		fill_err = tp.Fill(&conf)
		if fill_err != nil {
			return
		}
		if conf == nil {
			return
		}

		conf.TargetPeers()
	})
}

func Fuzz_Nosy_LocalSigner_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var priv *ecdsa.PrivateKey
		fill_err = tp.Fill(&priv)
		if fill_err != nil {
			return
		}
		if priv == nil {
			return
		}

		s := NewLocalSigner(priv)
		s.Close()
	})
}

func Fuzz_Nosy_LocalSigner_Sign__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var priv *ecdsa.PrivateKey
		fill_err = tp.Fill(&priv)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var domain [32]byte
		fill_err = tp.Fill(&domain)
		if fill_err != nil {
			return
		}
		var chainID *big.Int
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var encodedMsg []byte
		fill_err = tp.Fill(&encodedMsg)
		if fill_err != nil {
			return
		}
		if priv == nil || chainID == nil {
			return
		}

		s := NewLocalSigner(priv)
		s.Sign(ctx, domain, chainID, encodedMsg)
	})
}

func Fuzz_Nosy_NodeP2P_AltSyncEnabled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *NodeP2P
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.AltSyncEnabled()
	})
}

func Fuzz_Nosy_NodeP2P_BanIP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *NodeP2P
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var ip net.IP
		fill_err = tp.Fill(&ip)
		if fill_err != nil {
			return
		}
		var expiration time.Time
		fill_err = tp.Fill(&expiration)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.BanIP(ip, expiration)
	})
}

func Fuzz_Nosy_NodeP2P_BanPeer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *NodeP2P
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var id peer.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var expiration time.Time
		fill_err = tp.Fill(&expiration)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.BanPeer(id, expiration)
	})
}

func Fuzz_Nosy_NodeP2P_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *NodeP2P
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.Close()
	})
}

func Fuzz_Nosy_NodeP2P_ConnectionGater__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *NodeP2P
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.ConnectionGater()
	})
}

func Fuzz_Nosy_NodeP2P_ConnectionManager__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *NodeP2P
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.ConnectionManager()
	})
}

// skipping Fuzz_Nosy_NodeP2P_DiscoveryProcess__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

func Fuzz_Nosy_NodeP2P_Dv5Local__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *NodeP2P
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.Dv5Local()
	})
}

func Fuzz_Nosy_NodeP2P_Dv5Udp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *NodeP2P
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.Dv5Udp()
	})
}

func Fuzz_Nosy_NodeP2P_GetPeerScore__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *NodeP2P
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var id peer.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.GetPeerScore(id)
	})
}

func Fuzz_Nosy_NodeP2P_GossipOut__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *NodeP2P
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.GossipOut()
	})
}

func Fuzz_Nosy_NodeP2P_GossipSub__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *NodeP2P
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.GossipSub()
	})
}

func Fuzz_Nosy_NodeP2P_Host__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *NodeP2P
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.Host()
	})
}

func Fuzz_Nosy_NodeP2P_IsStatic__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *NodeP2P
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var id peer.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.IsStatic(id)
	})
}

func Fuzz_Nosy_NodeP2P_Peers__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *NodeP2P
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.Peers()
	})
}

func Fuzz_Nosy_NodeP2P_RequestL2Range__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *NodeP2P
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var start eth.L2BlockRef
		fill_err = tp.Fill(&start)
		if fill_err != nil {
			return
		}
		var end eth.L2BlockRef
		fill_err = tp.Fill(&end)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.RequestL2Range(ctx, start, end)
	})
}

// skipping Fuzz_Nosy_NodeP2P_init__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

func Fuzz_Nosy_NoopApplicationScorer_ApplicationScore__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *NoopApplicationScorer
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var _x2 peer.ID
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.ApplicationScore(_x2)
	})
}

func Fuzz_Nosy_NoopApplicationScorer_onRejectedPayload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *NoopApplicationScorer
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var _x2 peer.ID
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.onRejectedPayload(_x2)
	})
}

func Fuzz_Nosy_NoopApplicationScorer_onResponseError__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *NoopApplicationScorer
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var _x2 peer.ID
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.onResponseError(_x2)
	})
}

func Fuzz_Nosy_NoopApplicationScorer_onValidResponse__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *NoopApplicationScorer
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var _x2 peer.ID
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.onValidResponse(_x2)
	})
}

func Fuzz_Nosy_NoopApplicationScorer_start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *NoopApplicationScorer
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.start()
	})
}

func Fuzz_Nosy_NoopApplicationScorer_stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *NoopApplicationScorer
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.stop()
	})
}

func Fuzz_Nosy_OpStackENRData_DecodeRLP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *OpStackENRData
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		var s *rlp.Stream
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if o == nil || s == nil {
			return
		}

		o.DecodeRLP(s)
	})
}

func Fuzz_Nosy_OpStackENRData_ENRKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *OpStackENRData
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		if o == nil {
			return
		}

		o.ENRKey()
	})
}

func Fuzz_Nosy_OpStackENRData_EncodeRLP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *OpStackENRData
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		var w io.Writer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if o == nil {
			return
		}

		o.EncodeRLP(w)
	})
}

func Fuzz_Nosy_PingService_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *PingService
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Close()
	})
}

func Fuzz_Nosy_PingService_pingPeer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *PingService
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var peerID peer.ID
		fill_err = tp.Fill(&peerID)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.pingPeer(ctx, peerID)
	})
}

func Fuzz_Nosy_PingService_pingPeers__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *PingService
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.pingPeers()
	})
}

func Fuzz_Nosy_PingService_pingPeersBackground__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *PingService
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.pingPeersBackground()
	})
}

func Fuzz_Nosy_Prepared_BanDuration__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Prepared
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.BanDuration()
	})
}

func Fuzz_Nosy_Prepared_BanPeers__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Prepared
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.BanPeers()
	})
}

func Fuzz_Nosy_Prepared_BanThreshold__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Prepared
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.BanThreshold()
	})
}

func Fuzz_Nosy_Prepared_Check__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Prepared
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Check()
	})
}

func Fuzz_Nosy_Prepared_ConfigureGossip__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Prepared
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var rollupCfg *rollup.Config
		fill_err = tp.Fill(&rollupCfg)
		if fill_err != nil {
			return
		}
		if p == nil || rollupCfg == nil {
			return
		}

		p.ConfigureGossip(rollupCfg)
	})
}

func Fuzz_Nosy_Prepared_Disabled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Prepared
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Disabled()
	})
}

// skipping Fuzz_Nosy_Prepared_Discovery__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

// skipping Fuzz_Nosy_Prepared_Host__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

func Fuzz_Nosy_Prepared_PeerScoringParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Prepared
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.PeerScoringParams()
	})
}

func Fuzz_Nosy_Prepared_ReqRespSyncEnabled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Prepared
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.ReqRespSyncEnabled()
	})
}

func Fuzz_Nosy_Prepared_TargetPeers__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Prepared
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.TargetPeers()
	})
}

func Fuzz_Nosy_PreparedSigner_SetupSigner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *PreparedSigner
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.SetupSigner(ctx)
	})
}

// skipping Fuzz_Nosy_ReqRespServer_HandleSyncRequest__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

// skipping Fuzz_Nosy_ReqRespServer_handleSyncRequest__ because parameters include func, chan, or unsupported interface: github.com/libp2p/go-libp2p/core/network.Stream

func Fuzz_Nosy_Secp256k1_DecodeRLP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v *Secp256k1
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var s *rlp.Stream
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if v == nil || s == nil {
			return
		}

		v.DecodeRLP(s)
	})
}

func Fuzz_Nosy_SyncClient_AddPeer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SyncClient
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var id peer.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.AddPeer(id)
	})
}

func Fuzz_Nosy_SyncClient_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SyncClient
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Close()
	})
}

func Fuzz_Nosy_SyncClient_RemovePeer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SyncClient
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var id peer.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.RemovePeer(id)
	})
}

func Fuzz_Nosy_SyncClient_RequestL2Range__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SyncClient
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var start eth.L2BlockRef
		fill_err = tp.Fill(&start)
		if fill_err != nil {
			return
		}
		var end eth.L2BlockRef
		fill_err = tp.Fill(&end)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.RequestL2Range(ctx, start, end)
	})
}

func Fuzz_Nosy_SyncClient_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SyncClient
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Start()
	})
}

func Fuzz_Nosy_SyncClient_doRequest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SyncClient
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var id peer.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var expectedBlockNum uint64
		fill_err = tp.Fill(&expectedBlockNum)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.doRequest(ctx, id, expectedBlockNum)
	})
}

func Fuzz_Nosy_SyncClient_isInFlight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SyncClient
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var num uint64
		fill_err = tp.Fill(&num)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.isInFlight(ctx, num)
	})
}

func Fuzz_Nosy_SyncClient_mainLoop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SyncClient
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.mainLoop()
	})
}

func Fuzz_Nosy_SyncClient_onQuarantineEvict__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SyncClient
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var key common.Hash
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var value syncResult
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.onQuarantineEvict(key, value)
	})
}

func Fuzz_Nosy_SyncClient_onRangeRequest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SyncClient
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var req rangeRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.onRangeRequest(ctx, req)
	})
}

func Fuzz_Nosy_SyncClient_onResult__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SyncClient
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var res syncResult
		fill_err = tp.Fill(&res)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.onResult(ctx, res)
	})
}

func Fuzz_Nosy_SyncClient_peerLoop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SyncClient
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var id peer.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.peerLoop(ctx, id)
	})
}

func Fuzz_Nosy_SyncClient_promote__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SyncClient
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var res syncResult
		fill_err = tp.Fill(&res)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.promote(ctx, res)
	})
}

func Fuzz_Nosy_SyncClient_tryPromote__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SyncClient
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var h common.Hash
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.tryPromote(h)
	})
}

func Fuzz_Nosy_blockTopic_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bt *blockTopic
		fill_err = tp.Fill(&bt)
		if fill_err != nil {
			return
		}
		if bt == nil {
			return
		}

		bt.Close()
	})
}

func Fuzz_Nosy_extraHost_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *extraHost
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.Close()
	})
}

func Fuzz_Nosy_extraHost_ConnectionGater__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *extraHost
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.ConnectionGater()
	})
}

func Fuzz_Nosy_extraHost_ConnectionManager__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *extraHost
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.ConnectionManager()
	})
}

func Fuzz_Nosy_extraHost_IsStatic__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *extraHost
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var peerID peer.ID
		fill_err = tp.Fill(&peerID)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.IsStatic(peerID)
	})
}

func Fuzz_Nosy_extraHost_SyncOnlyReqToStatic__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *extraHost
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.SyncOnlyReqToStatic()
	})
}

func Fuzz_Nosy_extraHost_dialStaticPeer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *extraHost
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var addr *peer.AddrInfo
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if e == nil || addr == nil {
			return
		}

		e.dialStaticPeer(ctx, addr)
	})
}

func Fuzz_Nosy_extraHost_initStaticPeers__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *extraHost
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.initStaticPeers()
	})
}

func Fuzz_Nosy_extraHost_monitorStaticPeers__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *extraHost
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.monitorStaticPeers()
	})
}

func Fuzz_Nosy_gossipTracer_Trace__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *gossipTracer
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var evt *pubsub_pb.TraceEvent
		fill_err = tp.Fill(&evt)
		if fill_err != nil {
			return
		}
		if g == nil || evt == nil {
			return
		}

		g.Trace(evt)
	})
}

// skipping Fuzz_Nosy_notifications_ClosedStream__ because parameters include func, chan, or unsupported interface: github.com/libp2p/go-libp2p/core/network.Network

// skipping Fuzz_Nosy_notifications_Connected__ because parameters include func, chan, or unsupported interface: github.com/libp2p/go-libp2p/core/network.Network

// skipping Fuzz_Nosy_notifications_Disconnected__ because parameters include func, chan, or unsupported interface: github.com/libp2p/go-libp2p/core/network.Network

// skipping Fuzz_Nosy_notifications_Listen__ because parameters include func, chan, or unsupported interface: github.com/libp2p/go-libp2p/core/network.Network

// skipping Fuzz_Nosy_notifications_ListenClose__ because parameters include func, chan, or unsupported interface: github.com/libp2p/go-libp2p/core/network.Network

// skipping Fuzz_Nosy_notifications_OpenedStream__ because parameters include func, chan, or unsupported interface: github.com/libp2p/go-libp2p/core/network.Network

func Fuzz_Nosy_peerApplicationScorer_ApplicationScore__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *peerApplicationScorer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var id peer.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.ApplicationScore(id)
	})
}

func Fuzz_Nosy_peerApplicationScorer_decayConnectedPeerScores__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *peerApplicationScorer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.decayConnectedPeerScores()
	})
}

func Fuzz_Nosy_peerApplicationScorer_decayScores__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *peerApplicationScorer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var id peer.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.decayScores(id)
	})
}

func Fuzz_Nosy_peerApplicationScorer_onRejectedPayload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *peerApplicationScorer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var id peer.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.onRejectedPayload(id)
	})
}

func Fuzz_Nosy_peerApplicationScorer_onResponseError__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *peerApplicationScorer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var id peer.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.onResponseError(id)
	})
}

func Fuzz_Nosy_peerApplicationScorer_onValidResponse__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *peerApplicationScorer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var id peer.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.onValidResponse(id)
	})
}

func Fuzz_Nosy_peerApplicationScorer_start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *peerApplicationScorer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.start()
	})
}

func Fuzz_Nosy_peerApplicationScorer_stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *peerApplicationScorer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.stop()
	})
}

func Fuzz_Nosy_publisher_AllBlockTopicsPeers__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *publisher
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.AllBlockTopicsPeers()
	})
}

func Fuzz_Nosy_publisher_BlocksTopicV1Peers__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *publisher
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.BlocksTopicV1Peers()
	})
}

func Fuzz_Nosy_publisher_BlocksTopicV2Peers__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *publisher
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.BlocksTopicV2Peers()
	})
}

func Fuzz_Nosy_publisher_BlocksTopicV3Peers__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *publisher
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.BlocksTopicV3Peers()
	})
}

func Fuzz_Nosy_publisher_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *publisher
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Close()
	})
}

// skipping Fuzz_Nosy_publisher_PublishL2Payload__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.Signer

func Fuzz_Nosy_requestIdMap_delete__(f *testing.F) {
	f.Fuzz(func(t *testing.T, key uint64) {
		r := newRequestIdMap()
		r.delete(key)
	})
}

func Fuzz_Nosy_requestIdMap_get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, key uint64) {
		r := newRequestIdMap()
		r.get(key)
	})
}

func Fuzz_Nosy_requestIdMap_set__(f *testing.F) {
	f.Fuzz(func(t *testing.T, key uint64, value bool) {
		r := newRequestIdMap()
		r.set(key, value)
	})
}

func Fuzz_Nosy_scorer_ApplicationScore__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *scorer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var id peer.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.ApplicationScore(id)
	})
}

func Fuzz_Nosy_scorer_SnapshotHook__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *scorer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.SnapshotHook()
	})
}

func Fuzz_Nosy_seenBlocks_hasSeen__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sb *seenBlocks
		fill_err = tp.Fill(&sb)
		if fill_err != nil {
			return
		}
		var h common.Hash
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if sb == nil {
			return
		}

		sb.hasSeen(h)
	})
}

func Fuzz_Nosy_seenBlocks_markSeen__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sb *seenBlocks
		fill_err = tp.Fill(&sb)
		if fill_err != nil {
			return
		}
		var h common.Hash
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if sb == nil {
			return
		}

		sb.markSeen(h)
	})
}

// skipping Fuzz_Nosy_API_BlockAddr__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.API

// skipping Fuzz_Nosy_API_BlockPeer__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.API

// skipping Fuzz_Nosy_API_BlockSubnet__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.API

// skipping Fuzz_Nosy_API_ConnectPeer__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.API

// skipping Fuzz_Nosy_API_DisconnectPeer__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.API

// skipping Fuzz_Nosy_API_DiscoveryTable__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.API

// skipping Fuzz_Nosy_API_ListBlockedAddrs__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.API

// skipping Fuzz_Nosy_API_ListBlockedPeers__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.API

// skipping Fuzz_Nosy_API_ListBlockedSubnets__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.API

// skipping Fuzz_Nosy_API_PeerStats__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.API

// skipping Fuzz_Nosy_API_Peers__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.API

// skipping Fuzz_Nosy_API_ProtectPeer__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.API

// skipping Fuzz_Nosy_API_Self__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.API

// skipping Fuzz_Nosy_API_UnblockAddr__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.API

// skipping Fuzz_Nosy_API_UnblockPeer__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.API

// skipping Fuzz_Nosy_API_UnblockSubnet__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.API

// skipping Fuzz_Nosy_API_UnprotectPeer__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.API

// skipping Fuzz_Nosy_ApplicationScorer_ApplicationScore__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.ApplicationScorer

// skipping Fuzz_Nosy_ApplicationScorer_onRejectedPayload__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.ApplicationScorer

// skipping Fuzz_Nosy_ApplicationScorer_onResponseError__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.ApplicationScorer

// skipping Fuzz_Nosy_ApplicationScorer_onValidResponse__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.ApplicationScorer

// skipping Fuzz_Nosy_ApplicationScorer_start__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.ApplicationScorer

// skipping Fuzz_Nosy_ApplicationScorer_stop__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.ApplicationScorer

// skipping Fuzz_Nosy_ExtraHostFeatures_ConnectionGater__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.ExtraHostFeatures

// skipping Fuzz_Nosy_ExtraHostFeatures_ConnectionManager__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.ExtraHostFeatures

// skipping Fuzz_Nosy_ExtraHostFeatures_IsStatic__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.ExtraHostFeatures

// skipping Fuzz_Nosy_ExtraHostFeatures_SyncOnlyReqToStatic__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.ExtraHostFeatures

// skipping Fuzz_Nosy_GossipIn_OnUnsafeL2Payload__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.GossipIn

// skipping Fuzz_Nosy_GossipMetricer_RecordGossipEvent__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.GossipMetricer

// skipping Fuzz_Nosy_GossipOut_Close__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.GossipOut

// skipping Fuzz_Nosy_GossipOut_PublishL2Payload__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.GossipOut

// skipping Fuzz_Nosy_GossipRuntimeConfig_P2PSequencerAddress__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.GossipRuntimeConfig

// skipping Fuzz_Nosy_GossipSetupConfigurables_ConfigureGossip__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.GossipSetupConfigurables

// skipping Fuzz_Nosy_GossipSetupConfigurables_PeerScoringParams__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.GossipSetupConfigurables

// skipping Fuzz_Nosy_GossipTopicInfo_AllBlockTopicsPeers__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.GossipTopicInfo

// skipping Fuzz_Nosy_GossipTopicInfo_BlocksTopicV1Peers__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.GossipTopicInfo

// skipping Fuzz_Nosy_GossipTopicInfo_BlocksTopicV2Peers__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.GossipTopicInfo

// skipping Fuzz_Nosy_GossipTopicInfo_BlocksTopicV3Peers__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.GossipTopicInfo

// skipping Fuzz_Nosy_HostNewStream_NewStream__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.HostNewStream

// skipping Fuzz_Nosy_L2Chain_PayloadByNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.L2Chain

// skipping Fuzz_Nosy_Node_ConnectionGater__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.Node

// skipping Fuzz_Nosy_Node_ConnectionManager__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.Node

// skipping Fuzz_Nosy_Node_Dv5Local__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.Node

// skipping Fuzz_Nosy_Node_Dv5Udp__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.Node

// skipping Fuzz_Nosy_Node_GossipOut__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.Node

// skipping Fuzz_Nosy_Node_GossipSub__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.Node

// skipping Fuzz_Nosy_Node_Host__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.Node

// skipping Fuzz_Nosy_NotificationsMetricer_DecPeerCount__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.NotificationsMetricer

// skipping Fuzz_Nosy_NotificationsMetricer_DecStreamCount__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.NotificationsMetricer

// skipping Fuzz_Nosy_NotificationsMetricer_IncPeerCount__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.NotificationsMetricer

// skipping Fuzz_Nosy_NotificationsMetricer_IncStreamCount__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.NotificationsMetricer

// skipping Fuzz_Nosy_Peerstore_PeerInfo__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.Peerstore

// skipping Fuzz_Nosy_Peerstore_Peers__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.Peerstore

// skipping Fuzz_Nosy_Peerstore_SetScore__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.Peerstore

// skipping Fuzz_Nosy_ReqRespServerMetrics_ServerPayloadByNumberEvent__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.ReqRespServerMetrics

// skipping Fuzz_Nosy_ScoreBook_GetPeerScores__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.ScoreBook

// skipping Fuzz_Nosy_ScoreBook_SetScore__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.ScoreBook

// skipping Fuzz_Nosy_ScoreMetrics_SetPeerScores__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.ScoreMetrics

// skipping Fuzz_Nosy_Scorer_ApplicationScore__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.Scorer

// skipping Fuzz_Nosy_Scorer_SnapshotHook__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.Scorer

func Fuzz_Nosy_Secp256k1_ENRKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v Secp256k1
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}

		v.ENRKey()
	})
}

func Fuzz_Nosy_Secp256k1_EncodeRLP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v Secp256k1
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var w io.Writer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}

		v.EncodeRLP(w)
	})
}

// skipping Fuzz_Nosy_SetupP2P_BanDuration__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.SetupP2P

// skipping Fuzz_Nosy_SetupP2P_BanPeers__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.SetupP2P

// skipping Fuzz_Nosy_SetupP2P_BanThreshold__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.SetupP2P

// skipping Fuzz_Nosy_SetupP2P_Check__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.SetupP2P

// skipping Fuzz_Nosy_SetupP2P_Disabled__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.SetupP2P

// skipping Fuzz_Nosy_SetupP2P_Discovery__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.SetupP2P

// skipping Fuzz_Nosy_SetupP2P_Host__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.SetupP2P

// skipping Fuzz_Nosy_SetupP2P_ReqRespSyncEnabled__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.SetupP2P

// skipping Fuzz_Nosy_SetupP2P_TargetPeers__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.SetupP2P

// skipping Fuzz_Nosy_Signer_Sign__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.Signer

// skipping Fuzz_Nosy_SignerSetup_SetupSigner__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.SignerSetup

// skipping Fuzz_Nosy_SyncClientMetrics_ClientPayloadByNumberEvent__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.SyncClientMetrics

// skipping Fuzz_Nosy_SyncClientMetrics_PayloadsQuarantineSize__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.SyncClientMetrics

// skipping Fuzz_Nosy_SyncPeerScorer_onRejectedPayload__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.SyncPeerScorer

// skipping Fuzz_Nosy_SyncPeerScorer_onResponseError__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.SyncPeerScorer

// skipping Fuzz_Nosy_SyncPeerScorer_onValidResponse__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.SyncPeerScorer

func Fuzz_Nosy_requestResultErr_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r requestResultErr
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}

		r.Error()
	})
}

func Fuzz_Nosy_requestResultErr_ResultCode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r requestResultErr
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}

		r.ResultCode()
	})
}

// skipping Fuzz_Nosy_ConfigurePeerScoring__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p.GossipSetupConfigurables

// skipping Fuzz_Nosy_FilterEnodes__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

// skipping Fuzz_Nosy_FindActiveTCPPort__ because parameters include func, chan, or unsupported interface: github.com/libp2p/go-libp2p/core/host.Host

// skipping Fuzz_Nosy_LogTopicEvents__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

func Fuzz_Nosy_ScoreDecay__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var duration time.Duration
		fill_err = tp.Fill(&duration)
		if fill_err != nil {
			return
		}
		var slot time.Duration
		fill_err = tp.Fill(&slot)
		if fill_err != nil {
			return
		}

		ScoreDecay(duration, slot)
	})
}

func Fuzz_Nosy_blocksTopicV1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *rollup.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		blocksTopicV1(cfg)
	})
}

func Fuzz_Nosy_blocksTopicV2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *rollup.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		blocksTopicV2(cfg)
	})
}

func Fuzz_Nosy_blocksTopicV3__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *rollup.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		blocksTopicV3(cfg)
	})
}

func Fuzz_Nosy_combinePeers__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var allPeers [][]peer.ID
		fill_err = tp.Fill(&allPeers)
		if fill_err != nil {
			return
		}

		combinePeers(allPeers...)
	})
}

func Fuzz_Nosy_enrToAddrInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *enode.Node
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		enrToAddrInfo(r)
	})
}

func Fuzz_Nosy_inMeshCap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var slot time.Duration
		fill_err = tp.Fill(&slot)
		if fill_err != nil {
			return
		}

		inMeshCap(slot)
	})
}

// skipping Fuzz_Nosy_panicGuard__ because parameters include func, chan, or unsupported interface: func(T, S, U) error

func Fuzz_Nosy_prefixRPC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, method string) {
		prefixRPC(method)
	})
}

func Fuzz_Nosy_validationResultString__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v pubsub.ValidationResult
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}

		validationResultString(v)
	})
}
