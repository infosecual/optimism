package monitor

import (
	"testing"

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

func Fuzz_Nosy_PeerMonitor_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *PeerMonitor
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Start()
	})
}

func Fuzz_Nosy_PeerMonitor_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *PeerMonitor
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Stop()
	})
}

// skipping Fuzz_Nosy_PeerMonitor_background__ because parameters include func, chan, or unsupported interface: func() error

func Fuzz_Nosy_PeerMonitor_checkNextPeer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *PeerMonitor
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.checkNextPeer()
	})
}

// skipping Fuzz_Nosy_PeerManager_BanPeer__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p/monitor.PeerManager

// skipping Fuzz_Nosy_PeerManager_GetPeerScore__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p/monitor.PeerManager

// skipping Fuzz_Nosy_PeerManager_IsStatic__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p/monitor.PeerManager

// skipping Fuzz_Nosy_PeerManager_Peers__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p/monitor.PeerManager
