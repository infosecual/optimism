package gating

import (
	"testing"

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

// skipping Fuzz_Nosy_ExpiryConnectionGater_InterceptAccept__ because parameters include func, chan, or unsupported interface: github.com/libp2p/go-libp2p/core/network.ConnMultiaddrs

// skipping Fuzz_Nosy_ExpiryConnectionGater_InterceptAddrDial__ because parameters include func, chan, or unsupported interface: github.com/multiformats/go-multiaddr.Multiaddr

func Fuzz_Nosy_ExpiryConnectionGater_InterceptPeerDial__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *ExpiryConnectionGater
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var p peer.ID
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.InterceptPeerDial(p)
	})
}

// skipping Fuzz_Nosy_ExpiryConnectionGater_InterceptSecured__ because parameters include func, chan, or unsupported interface: github.com/libp2p/go-libp2p/core/network.ConnMultiaddrs

func Fuzz_Nosy_ExpiryConnectionGater_UnblockPeer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *ExpiryConnectionGater
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var p peer.ID
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.UnblockPeer(p)
	})
}

// skipping Fuzz_Nosy_ExpiryConnectionGater_addrBanExpiryCheck__ because parameters include func, chan, or unsupported interface: github.com/multiformats/go-multiaddr.Multiaddr

func Fuzz_Nosy_ExpiryConnectionGater_peerBanExpiryCheck__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *ExpiryConnectionGater
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var p peer.ID
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.peerBanExpiryCheck(p)
	})
}

// skipping Fuzz_Nosy_MeteredConnectionGater_InterceptAccept__ because parameters include func, chan, or unsupported interface: github.com/libp2p/go-libp2p/core/network.ConnMultiaddrs

// skipping Fuzz_Nosy_MeteredConnectionGater_InterceptAddrDial__ because parameters include func, chan, or unsupported interface: github.com/multiformats/go-multiaddr.Multiaddr

func Fuzz_Nosy_MeteredConnectionGater_InterceptPeerDial__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *MeteredConnectionGater
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var p peer.ID
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.InterceptPeerDial(p)
	})
}

// skipping Fuzz_Nosy_MeteredConnectionGater_InterceptSecured__ because parameters include func, chan, or unsupported interface: github.com/libp2p/go-libp2p/core/network.ConnMultiaddrs

// skipping Fuzz_Nosy_ScoringConnectionGater_InterceptAddrDial__ because parameters include func, chan, or unsupported interface: github.com/multiformats/go-multiaddr.Multiaddr

func Fuzz_Nosy_ScoringConnectionGater_InterceptPeerDial__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *ScoringConnectionGater
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var p peer.ID
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.InterceptPeerDial(p)
	})
}

// skipping Fuzz_Nosy_ScoringConnectionGater_InterceptSecured__ because parameters include func, chan, or unsupported interface: github.com/libp2p/go-libp2p/core/network.ConnMultiaddrs

func Fuzz_Nosy_ScoringConnectionGater_checkScore__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *ScoringConnectionGater
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var p peer.ID
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.checkScore(p)
	})
}

// skipping Fuzz_Nosy_BlockingConnectionGater_BlockAddr__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p/gating.BlockingConnectionGater

// skipping Fuzz_Nosy_BlockingConnectionGater_BlockPeer__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p/gating.BlockingConnectionGater

// skipping Fuzz_Nosy_BlockingConnectionGater_BlockSubnet__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p/gating.BlockingConnectionGater

// skipping Fuzz_Nosy_BlockingConnectionGater_ListBlockedAddrs__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p/gating.BlockingConnectionGater

// skipping Fuzz_Nosy_BlockingConnectionGater_ListBlockedPeers__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p/gating.BlockingConnectionGater

// skipping Fuzz_Nosy_BlockingConnectionGater_ListBlockedSubnets__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p/gating.BlockingConnectionGater

// skipping Fuzz_Nosy_BlockingConnectionGater_UnblockAddr__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p/gating.BlockingConnectionGater

// skipping Fuzz_Nosy_BlockingConnectionGater_UnblockPeer__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p/gating.BlockingConnectionGater

// skipping Fuzz_Nosy_BlockingConnectionGater_UnblockSubnet__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p/gating.BlockingConnectionGater

// skipping Fuzz_Nosy_ConnectionGaterMetrics_RecordAccept__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p/gating.ConnectionGaterMetrics

// skipping Fuzz_Nosy_ConnectionGaterMetrics_RecordDial__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p/gating.ConnectionGaterMetrics

// skipping Fuzz_Nosy_Scores_GetPeerScore__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p/gating.Scores

// skipping Fuzz_Nosy_UnbanMetrics_RecordIPUnban__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p/gating.UnbanMetrics

// skipping Fuzz_Nosy_UnbanMetrics_RecordPeerUnban__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p/gating.UnbanMetrics
