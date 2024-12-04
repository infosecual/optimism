package mocks

import (
	net "net"
	"testing"
	time "time"

	control "github.com/libp2p/go-libp2p/core/control"
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

func Fuzz_Nosy_BlockingConnectionGater_BlockAddr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *BlockingConnectionGater
		fill_err = tp.Fill(&_m)
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

		_m.BlockAddr(ip)
	})
}

func Fuzz_Nosy_BlockingConnectionGater_BlockPeer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *BlockingConnectionGater
		fill_err = tp.Fill(&_m)
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

		_m.BlockPeer(p)
	})
}

func Fuzz_Nosy_BlockingConnectionGater_BlockSubnet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *BlockingConnectionGater
		fill_err = tp.Fill(&_m)
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

		_m.BlockSubnet(ipnet)
	})
}

func Fuzz_Nosy_BlockingConnectionGater_EXPECT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *BlockingConnectionGater
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

// skipping Fuzz_Nosy_BlockingConnectionGater_InterceptAccept__ because parameters include func, chan, or unsupported interface: github.com/libp2p/go-libp2p/core/network.ConnMultiaddrs

// skipping Fuzz_Nosy_BlockingConnectionGater_InterceptAddrDial__ because parameters include func, chan, or unsupported interface: github.com/multiformats/go-multiaddr.Multiaddr

func Fuzz_Nosy_BlockingConnectionGater_InterceptPeerDial__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *BlockingConnectionGater
		fill_err = tp.Fill(&_m)
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

		_m.InterceptPeerDial(p)
	})
}

// skipping Fuzz_Nosy_BlockingConnectionGater_InterceptSecured__ because parameters include func, chan, or unsupported interface: github.com/libp2p/go-libp2p/core/network.ConnMultiaddrs

// skipping Fuzz_Nosy_BlockingConnectionGater_InterceptUpgraded__ because parameters include func, chan, or unsupported interface: github.com/libp2p/go-libp2p/core/network.Conn

func Fuzz_Nosy_BlockingConnectionGater_ListBlockedAddrs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *BlockingConnectionGater
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.ListBlockedAddrs()
	})
}

func Fuzz_Nosy_BlockingConnectionGater_ListBlockedPeers__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *BlockingConnectionGater
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.ListBlockedPeers()
	})
}

func Fuzz_Nosy_BlockingConnectionGater_ListBlockedSubnets__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *BlockingConnectionGater
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.ListBlockedSubnets()
	})
}

func Fuzz_Nosy_BlockingConnectionGater_UnblockAddr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *BlockingConnectionGater
		fill_err = tp.Fill(&_m)
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

		_m.UnblockAddr(ip)
	})
}

func Fuzz_Nosy_BlockingConnectionGater_UnblockPeer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *BlockingConnectionGater
		fill_err = tp.Fill(&_m)
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

		_m.UnblockPeer(p)
	})
}

func Fuzz_Nosy_BlockingConnectionGater_UnblockSubnet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *BlockingConnectionGater
		fill_err = tp.Fill(&_m)
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

		_m.UnblockSubnet(ipnet)
	})
}

// skipping Fuzz_Nosy_BlockingConnectionGater_BlockAddr_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_BlockingConnectionGater_BlockAddr_Call_Run__ because parameters include func, chan, or unsupported interface: func(ip net.IP)

// skipping Fuzz_Nosy_BlockingConnectionGater_BlockAddr_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(net.IP) error

// skipping Fuzz_Nosy_BlockingConnectionGater_BlockPeer_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_BlockingConnectionGater_BlockPeer_Call_Run__ because parameters include func, chan, or unsupported interface: func(p github.com/libp2p/go-libp2p/core/peer.ID)

// skipping Fuzz_Nosy_BlockingConnectionGater_BlockPeer_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(github.com/libp2p/go-libp2p/core/peer.ID) error

// skipping Fuzz_Nosy_BlockingConnectionGater_BlockSubnet_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_BlockingConnectionGater_BlockSubnet_Call_Run__ because parameters include func, chan, or unsupported interface: func(ipnet *net.IPNet)

// skipping Fuzz_Nosy_BlockingConnectionGater_BlockSubnet_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(*net.IPNet) error

// skipping Fuzz_Nosy_BlockingConnectionGater_Expecter_BlockAddr__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_BlockingConnectionGater_Expecter_BlockPeer__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_BlockingConnectionGater_Expecter_BlockSubnet__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_BlockingConnectionGater_Expecter_InterceptAccept__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_BlockingConnectionGater_Expecter_InterceptAddrDial__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_BlockingConnectionGater_Expecter_InterceptPeerDial__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_BlockingConnectionGater_Expecter_InterceptSecured__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_BlockingConnectionGater_Expecter_InterceptUpgraded__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_BlockingConnectionGater_Expecter_ListBlockedAddrs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _e *BlockingConnectionGater_Expecter
		fill_err = tp.Fill(&_e)
		if fill_err != nil {
			return
		}
		if _e == nil {
			return
		}

		_e.ListBlockedAddrs()
	})
}

func Fuzz_Nosy_BlockingConnectionGater_Expecter_ListBlockedPeers__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _e *BlockingConnectionGater_Expecter
		fill_err = tp.Fill(&_e)
		if fill_err != nil {
			return
		}
		if _e == nil {
			return
		}

		_e.ListBlockedPeers()
	})
}

func Fuzz_Nosy_BlockingConnectionGater_Expecter_ListBlockedSubnets__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _e *BlockingConnectionGater_Expecter
		fill_err = tp.Fill(&_e)
		if fill_err != nil {
			return
		}
		if _e == nil {
			return
		}

		_e.ListBlockedSubnets()
	})
}

// skipping Fuzz_Nosy_BlockingConnectionGater_Expecter_UnblockAddr__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_BlockingConnectionGater_Expecter_UnblockPeer__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_BlockingConnectionGater_Expecter_UnblockSubnet__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_BlockingConnectionGater_InterceptAccept_Call_Return__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _c *BlockingConnectionGater_InterceptAccept_Call
		fill_err = tp.Fill(&_c)
		if fill_err != nil {
			return
		}
		var allow bool
		fill_err = tp.Fill(&allow)
		if fill_err != nil {
			return
		}
		if _c == nil {
			return
		}

		_c.Return(allow)
	})
}

// skipping Fuzz_Nosy_BlockingConnectionGater_InterceptAccept_Call_Run__ because parameters include func, chan, or unsupported interface: func(_a0 github.com/libp2p/go-libp2p/core/network.ConnMultiaddrs)

// skipping Fuzz_Nosy_BlockingConnectionGater_InterceptAccept_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(github.com/libp2p/go-libp2p/core/network.ConnMultiaddrs) bool

func Fuzz_Nosy_BlockingConnectionGater_InterceptAddrDial_Call_Return__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _c *BlockingConnectionGater_InterceptAddrDial_Call
		fill_err = tp.Fill(&_c)
		if fill_err != nil {
			return
		}
		var allow bool
		fill_err = tp.Fill(&allow)
		if fill_err != nil {
			return
		}
		if _c == nil {
			return
		}

		_c.Return(allow)
	})
}

// skipping Fuzz_Nosy_BlockingConnectionGater_InterceptAddrDial_Call_Run__ because parameters include func, chan, or unsupported interface: func(_a0 github.com/libp2p/go-libp2p/core/peer.ID, _a1 github.com/multiformats/go-multiaddr.Multiaddr)

// skipping Fuzz_Nosy_BlockingConnectionGater_InterceptAddrDial_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(github.com/libp2p/go-libp2p/core/peer.ID, github.com/multiformats/go-multiaddr.Multiaddr) bool

func Fuzz_Nosy_BlockingConnectionGater_InterceptPeerDial_Call_Return__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _c *BlockingConnectionGater_InterceptPeerDial_Call
		fill_err = tp.Fill(&_c)
		if fill_err != nil {
			return
		}
		var allow bool
		fill_err = tp.Fill(&allow)
		if fill_err != nil {
			return
		}
		if _c == nil {
			return
		}

		_c.Return(allow)
	})
}

// skipping Fuzz_Nosy_BlockingConnectionGater_InterceptPeerDial_Call_Run__ because parameters include func, chan, or unsupported interface: func(p github.com/libp2p/go-libp2p/core/peer.ID)

// skipping Fuzz_Nosy_BlockingConnectionGater_InterceptPeerDial_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(github.com/libp2p/go-libp2p/core/peer.ID) bool

func Fuzz_Nosy_BlockingConnectionGater_InterceptSecured_Call_Return__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _c *BlockingConnectionGater_InterceptSecured_Call
		fill_err = tp.Fill(&_c)
		if fill_err != nil {
			return
		}
		var allow bool
		fill_err = tp.Fill(&allow)
		if fill_err != nil {
			return
		}
		if _c == nil {
			return
		}

		_c.Return(allow)
	})
}

// skipping Fuzz_Nosy_BlockingConnectionGater_InterceptSecured_Call_Run__ because parameters include func, chan, or unsupported interface: func(_a0 github.com/libp2p/go-libp2p/core/network.Direction, _a1 github.com/libp2p/go-libp2p/core/peer.ID, _a2 github.com/libp2p/go-libp2p/core/network.ConnMultiaddrs)

// skipping Fuzz_Nosy_BlockingConnectionGater_InterceptSecured_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(github.com/libp2p/go-libp2p/core/network.Direction, github.com/libp2p/go-libp2p/core/peer.ID, github.com/libp2p/go-libp2p/core/network.ConnMultiaddrs) bool

func Fuzz_Nosy_BlockingConnectionGater_InterceptUpgraded_Call_Return__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _c *BlockingConnectionGater_InterceptUpgraded_Call
		fill_err = tp.Fill(&_c)
		if fill_err != nil {
			return
		}
		var allow bool
		fill_err = tp.Fill(&allow)
		if fill_err != nil {
			return
		}
		var reason control.DisconnectReason
		fill_err = tp.Fill(&reason)
		if fill_err != nil {
			return
		}
		if _c == nil {
			return
		}

		_c.Return(allow, reason)
	})
}

// skipping Fuzz_Nosy_BlockingConnectionGater_InterceptUpgraded_Call_Run__ because parameters include func, chan, or unsupported interface: func(_a0 github.com/libp2p/go-libp2p/core/network.Conn)

// skipping Fuzz_Nosy_BlockingConnectionGater_InterceptUpgraded_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(github.com/libp2p/go-libp2p/core/network.Conn) (bool, github.com/libp2p/go-libp2p/core/control.DisconnectReason)

func Fuzz_Nosy_BlockingConnectionGater_ListBlockedAddrs_Call_Return__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _c *BlockingConnectionGater_ListBlockedAddrs_Call
		fill_err = tp.Fill(&_c)
		if fill_err != nil {
			return
		}
		var _a0 []net.IP
		fill_err = tp.Fill(&_a0)
		if fill_err != nil {
			return
		}
		if _c == nil {
			return
		}

		_c.Return(_a0)
	})
}

// skipping Fuzz_Nosy_BlockingConnectionGater_ListBlockedAddrs_Call_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_BlockingConnectionGater_ListBlockedAddrs_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() []net.IP

func Fuzz_Nosy_BlockingConnectionGater_ListBlockedPeers_Call_Return__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _c *BlockingConnectionGater_ListBlockedPeers_Call
		fill_err = tp.Fill(&_c)
		if fill_err != nil {
			return
		}
		var _a0 []peer.ID
		fill_err = tp.Fill(&_a0)
		if fill_err != nil {
			return
		}
		if _c == nil {
			return
		}

		_c.Return(_a0)
	})
}

// skipping Fuzz_Nosy_BlockingConnectionGater_ListBlockedPeers_Call_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_BlockingConnectionGater_ListBlockedPeers_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() []github.com/libp2p/go-libp2p/core/peer.ID

func Fuzz_Nosy_BlockingConnectionGater_ListBlockedSubnets_Call_Return__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _c *BlockingConnectionGater_ListBlockedSubnets_Call
		fill_err = tp.Fill(&_c)
		if fill_err != nil {
			return
		}
		var _a0 []*net.IPNet
		fill_err = tp.Fill(&_a0)
		if fill_err != nil {
			return
		}
		if _c == nil {
			return
		}

		_c.Return(_a0)
	})
}

// skipping Fuzz_Nosy_BlockingConnectionGater_ListBlockedSubnets_Call_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_BlockingConnectionGater_ListBlockedSubnets_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() []*net.IPNet

// skipping Fuzz_Nosy_BlockingConnectionGater_UnblockAddr_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_BlockingConnectionGater_UnblockAddr_Call_Run__ because parameters include func, chan, or unsupported interface: func(ip net.IP)

// skipping Fuzz_Nosy_BlockingConnectionGater_UnblockAddr_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(net.IP) error

// skipping Fuzz_Nosy_BlockingConnectionGater_UnblockPeer_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_BlockingConnectionGater_UnblockPeer_Call_Run__ because parameters include func, chan, or unsupported interface: func(p github.com/libp2p/go-libp2p/core/peer.ID)

// skipping Fuzz_Nosy_BlockingConnectionGater_UnblockPeer_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(github.com/libp2p/go-libp2p/core/peer.ID) error

// skipping Fuzz_Nosy_BlockingConnectionGater_UnblockSubnet_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_BlockingConnectionGater_UnblockSubnet_Call_Run__ because parameters include func, chan, or unsupported interface: func(ipnet *net.IPNet)

// skipping Fuzz_Nosy_BlockingConnectionGater_UnblockSubnet_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(*net.IPNet) error

func Fuzz_Nosy_ExpiryStore_EXPECT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *ExpiryStore
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

func Fuzz_Nosy_ExpiryStore_GetIPBanExpiration__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *ExpiryStore
		fill_err = tp.Fill(&_m)
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

		_m.GetIPBanExpiration(ip)
	})
}

func Fuzz_Nosy_ExpiryStore_GetPeerBanExpiration__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *ExpiryStore
		fill_err = tp.Fill(&_m)
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

		_m.GetPeerBanExpiration(id)
	})
}

func Fuzz_Nosy_ExpiryStore_SetIPBanExpiration__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *ExpiryStore
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ip net.IP
		fill_err = tp.Fill(&ip)
		if fill_err != nil {
			return
		}
		var expiry time.Time
		fill_err = tp.Fill(&expiry)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.SetIPBanExpiration(ip, expiry)
	})
}

func Fuzz_Nosy_ExpiryStore_SetPeerBanExpiration__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *ExpiryStore
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var id peer.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var expiry time.Time
		fill_err = tp.Fill(&expiry)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.SetPeerBanExpiration(id, expiry)
	})
}

// skipping Fuzz_Nosy_ExpiryStore_Expecter_GetIPBanExpiration__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_ExpiryStore_Expecter_GetPeerBanExpiration__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_ExpiryStore_Expecter_SetIPBanExpiration__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_ExpiryStore_Expecter_SetPeerBanExpiration__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_ExpiryStore_GetIPBanExpiration_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_ExpiryStore_GetIPBanExpiration_Call_Run__ because parameters include func, chan, or unsupported interface: func(ip net.IP)

// skipping Fuzz_Nosy_ExpiryStore_GetIPBanExpiration_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(net.IP) (time.Time, error)

// skipping Fuzz_Nosy_ExpiryStore_GetPeerBanExpiration_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_ExpiryStore_GetPeerBanExpiration_Call_Run__ because parameters include func, chan, or unsupported interface: func(id github.com/libp2p/go-libp2p/core/peer.ID)

// skipping Fuzz_Nosy_ExpiryStore_GetPeerBanExpiration_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(github.com/libp2p/go-libp2p/core/peer.ID) (time.Time, error)

// skipping Fuzz_Nosy_ExpiryStore_SetIPBanExpiration_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_ExpiryStore_SetIPBanExpiration_Call_Run__ because parameters include func, chan, or unsupported interface: func(ip net.IP, expiry time.Time)

// skipping Fuzz_Nosy_ExpiryStore_SetIPBanExpiration_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(net.IP, time.Time) error

// skipping Fuzz_Nosy_ExpiryStore_SetPeerBanExpiration_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_ExpiryStore_SetPeerBanExpiration_Call_Run__ because parameters include func, chan, or unsupported interface: func(id github.com/libp2p/go-libp2p/core/peer.ID, expiry time.Time)

// skipping Fuzz_Nosy_ExpiryStore_SetPeerBanExpiration_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(github.com/libp2p/go-libp2p/core/peer.ID, time.Time) error

func Fuzz_Nosy_Scores_EXPECT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *Scores
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

func Fuzz_Nosy_Scores_GetPeerScore__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *Scores
		fill_err = tp.Fill(&_m)
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

		_m.GetPeerScore(id)
	})
}

// skipping Fuzz_Nosy_Scores_Expecter_GetPeerScore__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_Scores_GetPeerScore_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_Scores_GetPeerScore_Call_Run__ because parameters include func, chan, or unsupported interface: func(id github.com/libp2p/go-libp2p/core/peer.ID)

// skipping Fuzz_Nosy_Scores_GetPeerScore_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(github.com/libp2p/go-libp2p/core/peer.ID) (float64, error)
