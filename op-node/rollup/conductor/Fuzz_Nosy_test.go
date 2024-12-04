package conductor

import (
	"context"
	"testing"

	"github.com/ethereum-optimism/optimism/op-service/eth"
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

func Fuzz_Nosy_NoOpConductor_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *NoOpConductor
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.Close()
	})
}

func Fuzz_Nosy_NoOpConductor_CommitUnsafePayload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *NoOpConductor
		fill_err = tp.Fill(&c)
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
		if c == nil || payload == nil {
			return
		}

		c.CommitUnsafePayload(ctx, payload)
	})
}

func Fuzz_Nosy_NoOpConductor_Enabled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *NoOpConductor
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

		c.Enabled(ctx)
	})
}

func Fuzz_Nosy_NoOpConductor_Leader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *NoOpConductor
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

		c.Leader(ctx)
	})
}

func Fuzz_Nosy_NoOpConductor_OverrideLeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *NoOpConductor
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

		c.OverrideLeader(ctx)
	})
}

// skipping Fuzz_Nosy_SequencerConductor_Close__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/conductor.SequencerConductor

// skipping Fuzz_Nosy_SequencerConductor_CommitUnsafePayload__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/conductor.SequencerConductor

// skipping Fuzz_Nosy_SequencerConductor_Enabled__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/conductor.SequencerConductor

// skipping Fuzz_Nosy_SequencerConductor_Leader__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/conductor.SequencerConductor

// skipping Fuzz_Nosy_SequencerConductor_OverrideLeader__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/conductor.SequencerConductor
