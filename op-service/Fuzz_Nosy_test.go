package op_service

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

func Fuzz_Nosy_FindMonorepoRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, startDir string) {
		FindMonorepoRoot(startDir)
	})
}

func Fuzz_Nosy_FlagNameToEnvVarName__(f *testing.F) {
	f.Fuzz(func(t *testing.T, f1 string, prefix string) {
		FlagNameToEnvVarName(f1, prefix)
	})
}

func Fuzz_Nosy_FormatVersion__(f *testing.F) {
	f.Fuzz(func(t *testing.T, version string, gitCommit string, gitDate string, meta string) {
		FormatVersion(version, gitCommit, gitDate, meta)
	})
}

func Fuzz_Nosy_PrefixEnvVar__(f *testing.F) {
	f.Fuzz(func(t *testing.T, prefix string, suffix string) {
		PrefixEnvVar(prefix, suffix)
	})
}

// skipping Fuzz_Nosy_ValidateEnvVars__ because parameters include func, chan, or unsupported interface: []github.com/urfave/cli/v2.Flag

// skipping Fuzz_Nosy_WarnOnDeprecatedFlags__ because parameters include func, chan, or unsupported interface: []github.com/urfave/cli/v2.Flag

// skipping Fuzz_Nosy_cliFlagsToEnvVars__ because parameters include func, chan, or unsupported interface: []github.com/urfave/cli/v2.Flag

func Fuzz_Nosy_validateEnvVars__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var prefix string
		fill_err = tp.Fill(&prefix)
		if fill_err != nil {
			return
		}
		var providedEnvVars []string
		fill_err = tp.Fill(&providedEnvVars)
		if fill_err != nil {
			return
		}
		var definedEnvVars map[string]struct{}
		fill_err = tp.Fill(&definedEnvVars)
		if fill_err != nil {
			return
		}

		validateEnvVars(prefix, providedEnvVars, definedEnvVars)
	})
}
