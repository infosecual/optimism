package opnode

import (
	"context"
	"testing"

	"github.com/ethereum-optimism/optimism/op-service/eth"
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

func Fuzz_Nosy_FnTracer_OnNewL1Head__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *FnTracer
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var sig eth.L1BlockRef
		fill_err = tp.Fill(&sig)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.OnNewL1Head(ctx, sig)
	})
}

func Fuzz_Nosy_FnTracer_OnPublishL2Payload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *FnTracer
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var payload *eth.ExecutionPayloadEnvelope
		fill_err = tp.Fill(&payload)
		if fill_err != nil {
			return
		}
		if n == nil || payload == nil {
			return
		}

		n.OnPublishL2Payload(ctx, payload)
	})
}

func Fuzz_Nosy_FnTracer_OnUnsafeL2Payload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *FnTracer
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var from peer.ID
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		var payload *eth.ExecutionPayloadEnvelope
		fill_err = tp.Fill(&payload)
		if fill_err != nil {
			return
		}
		if n == nil || payload == nil {
			return
		}

		n.OnUnsafeL2Payload(ctx, from, payload)
	})
}

func Fuzz_Nosy_Opnode_P2P__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *Opnode
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		if o == nil {
			return
		}

		o.P2P()
	})
}

func Fuzz_Nosy_Opnode_RuntimeConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *Opnode
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		if o == nil {
			return
		}

		o.RuntimeConfig()
	})
}

func Fuzz_Nosy_Opnode_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *Opnode
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if o == nil {
			return
		}

		o.Stop(ctx)
	})
}

func Fuzz_Nosy_Opnode_Stopped__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *Opnode
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		if o == nil {
			return
		}

		o.Stopped()
	})
}

func Fuzz_Nosy_Opnode_UserRPC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *Opnode
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		if o == nil {
			return
		}

		o.UserRPC()
	})
}
