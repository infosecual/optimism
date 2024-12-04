package mocks

import (
	context "context"
	net "net"
	"testing"

	store "github.com/ethereum-optimism/optimism/op-node/p2p/store"
	peer "github.com/libp2p/go-libp2p/core/peer"
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

func Fuzz_Nosy_API_BlockAddr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *API
		fill_err = tp.Fill(&_m)
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
		if _m == nil {
			return
		}

		_m.BlockAddr(ctx, ip)
	})
}

func Fuzz_Nosy_API_BlockPeer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *API
		fill_err = tp.Fill(&_m)
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
		if _m == nil {
			return
		}

		_m.BlockPeer(ctx, p)
	})
}

func Fuzz_Nosy_API_BlockSubnet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *API
		fill_err = tp.Fill(&_m)
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
		if _m == nil || ipnet == nil {
			return
		}

		_m.BlockSubnet(ctx, ipnet)
	})
}

func Fuzz_Nosy_API_ConnectPeer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *API
		fill_err = tp.Fill(&_m)
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
		if _m == nil {
			return
		}

		_m.ConnectPeer(ctx, addr)
	})
}

func Fuzz_Nosy_API_DisconnectPeer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *API
		fill_err = tp.Fill(&_m)
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
		if _m == nil {
			return
		}

		_m.DisconnectPeer(ctx, id)
	})
}

func Fuzz_Nosy_API_DiscoveryTable__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *API
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.DiscoveryTable(ctx)
	})
}

func Fuzz_Nosy_API_EXPECT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *API
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.EXPECT()
	})
}

func Fuzz_Nosy_API_ListBlockedAddrs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *API
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.ListBlockedAddrs(ctx)
	})
}

func Fuzz_Nosy_API_ListBlockedPeers__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *API
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.ListBlockedPeers(ctx)
	})
}

func Fuzz_Nosy_API_ListBlockedSubnets__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *API
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.ListBlockedSubnets(ctx)
	})
}

func Fuzz_Nosy_API_PeerStats__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *API
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.PeerStats(ctx)
	})
}

func Fuzz_Nosy_API_Peers__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *API
		fill_err = tp.Fill(&_m)
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
		if _m == nil {
			return
		}

		_m.Peers(ctx, connected)
	})
}

func Fuzz_Nosy_API_ProtectPeer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *API
		fill_err = tp.Fill(&_m)
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
		if _m == nil {
			return
		}

		_m.ProtectPeer(ctx, p)
	})
}

func Fuzz_Nosy_API_Self__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *API
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.Self(ctx)
	})
}

func Fuzz_Nosy_API_UnblockAddr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *API
		fill_err = tp.Fill(&_m)
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
		if _m == nil {
			return
		}

		_m.UnblockAddr(ctx, ip)
	})
}

func Fuzz_Nosy_API_UnblockPeer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *API
		fill_err = tp.Fill(&_m)
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
		if _m == nil {
			return
		}

		_m.UnblockPeer(ctx, p)
	})
}

func Fuzz_Nosy_API_UnblockSubnet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *API
		fill_err = tp.Fill(&_m)
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
		if _m == nil || ipnet == nil {
			return
		}

		_m.UnblockSubnet(ctx, ipnet)
	})
}

func Fuzz_Nosy_API_UnprotectPeer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *API
		fill_err = tp.Fill(&_m)
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
		if _m == nil {
			return
		}

		_m.UnprotectPeer(ctx, p)
	})
}

// skipping Fuzz_Nosy_API_BlockAddr_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_API_BlockAddr_Call_Run__ because parameters include func, chan, or unsupported interface: func(ctx context.Context, ip net.IP)

// skipping Fuzz_Nosy_API_BlockAddr_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(context.Context, net.IP) error

// skipping Fuzz_Nosy_API_BlockPeer_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_API_BlockPeer_Call_Run__ because parameters include func, chan, or unsupported interface: func(ctx context.Context, p github.com/libp2p/go-libp2p/core/peer.ID)

// skipping Fuzz_Nosy_API_BlockPeer_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(context.Context, github.com/libp2p/go-libp2p/core/peer.ID) error

// skipping Fuzz_Nosy_API_BlockSubnet_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_API_BlockSubnet_Call_Run__ because parameters include func, chan, or unsupported interface: func(ctx context.Context, ipnet *net.IPNet)

// skipping Fuzz_Nosy_API_BlockSubnet_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(context.Context, *net.IPNet) error

// skipping Fuzz_Nosy_API_ConnectPeer_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_API_ConnectPeer_Call_Run__ because parameters include func, chan, or unsupported interface: func(ctx context.Context, addr string)

// skipping Fuzz_Nosy_API_ConnectPeer_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(context.Context, string) error

// skipping Fuzz_Nosy_API_DisconnectPeer_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_API_DisconnectPeer_Call_Run__ because parameters include func, chan, or unsupported interface: func(ctx context.Context, id github.com/libp2p/go-libp2p/core/peer.ID)

// skipping Fuzz_Nosy_API_DisconnectPeer_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(context.Context, github.com/libp2p/go-libp2p/core/peer.ID) error

// skipping Fuzz_Nosy_API_DiscoveryTable_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_API_DiscoveryTable_Call_Run__ because parameters include func, chan, or unsupported interface: func(ctx context.Context)

// skipping Fuzz_Nosy_API_DiscoveryTable_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(context.Context) ([]*github.com/ethereum/go-ethereum/p2p/enode.Node, error)

// skipping Fuzz_Nosy_API_Expecter_BlockAddr__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_API_Expecter_BlockPeer__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_API_Expecter_BlockSubnet__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_API_Expecter_ConnectPeer__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_API_Expecter_DisconnectPeer__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_API_Expecter_DiscoveryTable__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_API_Expecter_ListBlockedAddrs__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_API_Expecter_ListBlockedPeers__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_API_Expecter_ListBlockedSubnets__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_API_Expecter_PeerStats__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_API_Expecter_Peers__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_API_Expecter_ProtectPeer__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_API_Expecter_Self__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_API_Expecter_UnblockAddr__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_API_Expecter_UnblockPeer__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_API_Expecter_UnblockSubnet__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_API_Expecter_UnprotectPeer__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_API_ListBlockedAddrs_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_API_ListBlockedAddrs_Call_Run__ because parameters include func, chan, or unsupported interface: func(ctx context.Context)

// skipping Fuzz_Nosy_API_ListBlockedAddrs_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(context.Context) ([]net.IP, error)

// skipping Fuzz_Nosy_API_ListBlockedPeers_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_API_ListBlockedPeers_Call_Run__ because parameters include func, chan, or unsupported interface: func(ctx context.Context)

// skipping Fuzz_Nosy_API_ListBlockedPeers_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(context.Context) ([]github.com/libp2p/go-libp2p/core/peer.ID, error)

// skipping Fuzz_Nosy_API_ListBlockedSubnets_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_API_ListBlockedSubnets_Call_Run__ because parameters include func, chan, or unsupported interface: func(ctx context.Context)

// skipping Fuzz_Nosy_API_ListBlockedSubnets_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(context.Context) ([]*net.IPNet, error)

// skipping Fuzz_Nosy_API_PeerStats_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_API_PeerStats_Call_Run__ because parameters include func, chan, or unsupported interface: func(ctx context.Context)

// skipping Fuzz_Nosy_API_PeerStats_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(context.Context) (*github.com/ethereum-optimism/optimism/op-node/p2p.PeerStats, error)

// skipping Fuzz_Nosy_API_Peers_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_API_Peers_Call_Run__ because parameters include func, chan, or unsupported interface: func(ctx context.Context, connected bool)

// skipping Fuzz_Nosy_API_Peers_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(context.Context, bool) (*github.com/ethereum-optimism/optimism/op-node/p2p.PeerDump, error)

// skipping Fuzz_Nosy_API_ProtectPeer_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_API_ProtectPeer_Call_Run__ because parameters include func, chan, or unsupported interface: func(ctx context.Context, p github.com/libp2p/go-libp2p/core/peer.ID)

// skipping Fuzz_Nosy_API_ProtectPeer_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(context.Context, github.com/libp2p/go-libp2p/core/peer.ID) error

// skipping Fuzz_Nosy_API_Self_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_API_Self_Call_Run__ because parameters include func, chan, or unsupported interface: func(ctx context.Context)

// skipping Fuzz_Nosy_API_Self_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(context.Context) (*github.com/ethereum-optimism/optimism/op-node/p2p.PeerInfo, error)

// skipping Fuzz_Nosy_API_UnblockAddr_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_API_UnblockAddr_Call_Run__ because parameters include func, chan, or unsupported interface: func(ctx context.Context, ip net.IP)

// skipping Fuzz_Nosy_API_UnblockAddr_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(context.Context, net.IP) error

// skipping Fuzz_Nosy_API_UnblockPeer_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_API_UnblockPeer_Call_Run__ because parameters include func, chan, or unsupported interface: func(ctx context.Context, p github.com/libp2p/go-libp2p/core/peer.ID)

// skipping Fuzz_Nosy_API_UnblockPeer_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(context.Context, github.com/libp2p/go-libp2p/core/peer.ID) error

// skipping Fuzz_Nosy_API_UnblockSubnet_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_API_UnblockSubnet_Call_Run__ because parameters include func, chan, or unsupported interface: func(ctx context.Context, ipnet *net.IPNet)

// skipping Fuzz_Nosy_API_UnblockSubnet_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(context.Context, *net.IPNet) error

// skipping Fuzz_Nosy_API_UnprotectPeer_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_API_UnprotectPeer_Call_Run__ because parameters include func, chan, or unsupported interface: func(ctx context.Context, p github.com/libp2p/go-libp2p/core/peer.ID)

// skipping Fuzz_Nosy_API_UnprotectPeer_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(context.Context, github.com/libp2p/go-libp2p/core/peer.ID) error

func Fuzz_Nosy_GossipMetricer_RecordGossipEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *GossipMetricer
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var evType int32
		fill_err = tp.Fill(&evType)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.RecordGossipEvent(evType)
	})
}

func Fuzz_Nosy_PeerGater_IsBlocked__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *PeerGater
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var _a0 peer.ID
		fill_err = tp.Fill(&_a0)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.IsBlocked(_a0)
	})
}

func Fuzz_Nosy_PeerGater_Update__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *PeerGater
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var _a0 peer.ID
		fill_err = tp.Fill(&_a0)
		if fill_err != nil {
			return
		}
		var _a1 float64
		fill_err = tp.Fill(&_a1)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.Update(_a0, _a1)
	})
}

func Fuzz_Nosy_Peerstore_PeerInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *Peerstore
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var _a0 peer.ID
		fill_err = tp.Fill(&_a0)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.PeerInfo(_a0)
	})
}

func Fuzz_Nosy_Peerstore_Peers__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *Peerstore
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.Peers()
	})
}

// skipping Fuzz_Nosy_Peerstore_SetScore__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p/store.ScoreDiff

func Fuzz_Nosy_ScoreMetrics_SetPeerScores__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *ScoreMetrics
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var _a0 []store.PeerScores
		fill_err = tp.Fill(&_a0)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.SetPeerScores(_a0)
	})
}

// skipping Fuzz_Nosy_mockConstructorTestingTNewPeerGater_Cleanup__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p/mocks.mockConstructorTestingTNewPeerGater
