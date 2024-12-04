package helpers

import (
	"testing"
	"github.com/ethereum-optimism/optimism/op-e2e/actions/proofs"
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
	
	// skipping Fuzz_Nosy_L2FaultProofEnv_RunFaultProofProgram__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_TestMatrix[cfg any]_AddTestCase__ because parameters include func, chan, or unsupported interface: cfg

func Fuzz_Nosy_TestMatrix[cfg any]_Run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
	if t1 == nil {
		return
	}

	suite := NewMatrix()
	suite.Run(t1)
	})
}

// skipping Fuzz_Nosy_TestMatrix[cfg any]_WithCommonInputParams__ because parameters include func, chan, or unsupported interface: []github.com/ethereum-optimism/optimism/op-e2e/actions/proofs/helpers.FixtureInputParam

func Fuzz_Nosy_NewForkMatrix__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var forks []*Hardfork
		fill_err = tp.Fill(&forks)
		if fill_err != nil {
			return
		}

	NewForkMatrix(forks...)
	})
}

func Fuzz_Nosy_convertToKebabCase__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input string) {
	convertToKebabCase(input)
	})
}

// skipping Fuzz_Nosy_tryDumpTestFixture__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

