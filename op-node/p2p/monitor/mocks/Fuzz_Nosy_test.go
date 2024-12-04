package mocks

import (
	"testing"
	time "time"

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

func Fuzz_Nosy_PeerManager_BanPeer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *PeerManager
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var _a0 peer.ID
		fill_err = tp.Fill(&_a0)
		if fill_err != nil {
			return
		}
		var _a1 time.Time
		fill_err = tp.Fill(&_a1)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.BanPeer(_a0, _a1)
	})
}

func Fuzz_Nosy_PeerManager_EXPECT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *PeerManager
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

func Fuzz_Nosy_PeerManager_GetPeerScore__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *PeerManager
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

func Fuzz_Nosy_PeerManager_IsStatic__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *PeerManager
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

		_m.IsStatic(_a0)
	})
}

func Fuzz_Nosy_PeerManager_Peers__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *PeerManager
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

// skipping Fuzz_Nosy_PeerManager_BanPeer_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_PeerManager_BanPeer_Call_Run__ because parameters include func, chan, or unsupported interface: func(_a0 github.com/libp2p/go-libp2p/core/peer.ID, _a1 time.Time)

// skipping Fuzz_Nosy_PeerManager_BanPeer_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(github.com/libp2p/go-libp2p/core/peer.ID, time.Time) error

// skipping Fuzz_Nosy_PeerManager_Expecter_BanPeer__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_PeerManager_Expecter_GetPeerScore__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_PeerManager_Expecter_IsStatic__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_PeerManager_Expecter_Peers__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _e *PeerManager_Expecter
		fill_err = tp.Fill(&_e)
		if fill_err != nil {
			return
		}
		if _e == nil {
			return
		}

		_e.Peers()
	})
}

// skipping Fuzz_Nosy_PeerManager_GetPeerScore_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_PeerManager_GetPeerScore_Call_Run__ because parameters include func, chan, or unsupported interface: func(id github.com/libp2p/go-libp2p/core/peer.ID)

// skipping Fuzz_Nosy_PeerManager_GetPeerScore_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(github.com/libp2p/go-libp2p/core/peer.ID) (float64, error)

func Fuzz_Nosy_PeerManager_IsStatic_Call_Return__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _c *PeerManager_IsStatic_Call
		fill_err = tp.Fill(&_c)
		if fill_err != nil {
			return
		}
		var _a0 bool
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

// skipping Fuzz_Nosy_PeerManager_IsStatic_Call_Run__ because parameters include func, chan, or unsupported interface: func(_a0 github.com/libp2p/go-libp2p/core/peer.ID)

// skipping Fuzz_Nosy_PeerManager_IsStatic_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(github.com/libp2p/go-libp2p/core/peer.ID) bool

func Fuzz_Nosy_PeerManager_Peers_Call_Return__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _c *PeerManager_Peers_Call
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

// skipping Fuzz_Nosy_PeerManager_Peers_Call_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_PeerManager_Peers_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() []github.com/libp2p/go-libp2p/core/peer.ID
