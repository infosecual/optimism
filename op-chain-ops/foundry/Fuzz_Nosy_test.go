package foundry

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
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

func Fuzz_Nosy_Artifact_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string, d2 []byte) {
		a, err := ReadArtifact(path)
		if err != nil {
			return
		}
		a.UnmarshalJSON(d2)
	})
}

func Fuzz_Nosy_ArtifactsFS_ListArtifacts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dirPath string) {
		af := OpenArtifactsDir(dirPath)
		af.ListArtifacts()
	})
}

func Fuzz_Nosy_ArtifactsFS_ListContracts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dirPath string, name string) {
		af := OpenArtifactsDir(dirPath)
		af.ListContracts(name)
	})
}

func Fuzz_Nosy_ArtifactsFS_ReadArtifact__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dirPath string, name string, contract string) {
		af := OpenArtifactsDir(dirPath)
		af.ReadArtifact(name, contract)
	})
}

func Fuzz_Nosy_ForgeAllocs_Copy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, allocsPath string) {
		d, err := LoadForgeAllocs(allocsPath)
		if err != nil {
			return
		}
		d.Copy()
	})
}

// skipping Fuzz_Nosy_ForgeAllocs_FromState__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-chain-ops/foundry.StateDB

func Fuzz_Nosy_ForgeAllocs_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, allocsPath string, b []byte) {
		d, err := LoadForgeAllocs(allocsPath)
		if err != nil {
			return
		}
		d.UnmarshalJSON(b)
	})
}

func Fuzz_Nosy_LinkableBytecode_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var lb *LinkableBytecode
		fill_err = tp.Fill(&lb)
		if fill_err != nil {
			return
		}
		var d2 []byte
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		if lb == nil {
			return
		}

		lb.UnmarshalJSON(d2)
	})
}

func Fuzz_Nosy_SourceMapFS_ReadSourceIDs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SourceMapFS
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var path string
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		var contract string
		fill_err = tp.Fill(&contract)
		if fill_err != nil {
			return
		}
		var compilerVersion string
		fill_err = tp.Fill(&compilerVersion)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.ReadSourceIDs(path, contract, compilerVersion)
	})
}

func Fuzz_Nosy_SourceMapFS_SourceMap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SourceMapFS
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var artifact *Artifact
		fill_err = tp.Fill(&artifact)
		if fill_err != nil {
			return
		}
		var contract string
		fill_err = tp.Fill(&contract)
		if fill_err != nil {
			return
		}
		if s == nil || artifact == nil {
			return
		}

		s.SourceMap(artifact, contract)
	})
}

func Fuzz_Nosy_SourceMapFS_readBuild__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SourceMapFS
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var buildInfoPath string
		fill_err = tp.Fill(&buildInfoPath)
		if fill_err != nil {
			return
		}
		var id string
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.readBuild(buildInfoPath, id)
	})
}

func Fuzz_Nosy_SourceMapFS_readBuildCache__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SourceMapFS
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.readBuildCache()
	})
}

func Fuzz_Nosy_forgeAllocsDump_OnAccount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *forgeAllocsDump
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var address *common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var account state.DumpAccount
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		if d == nil || address == nil {
			return
		}

		d.OnAccount(address, account)
	})
}

func Fuzz_Nosy_forgeAllocsDump_OnRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *forgeAllocsDump
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.OnRoot(hash)
	})
}

func Fuzz_Nosy_Artifact_MarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string) {
		a, err := ReadArtifact(path)
		if err != nil {
			return
		}
		a.MarshalJSON()
	})
}

func Fuzz_Nosy_ForgeAllocs_MarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, allocsPath string) {
		d, err := LoadForgeAllocs(allocsPath)
		if err != nil {
			return
		}
		d.MarshalJSON()
	})
}

func Fuzz_Nosy_LinkableBytecode_MarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var lb LinkableBytecode
		fill_err = tp.Fill(&lb)
		if fill_err != nil {
			return
		}

		lb.MarshalText()
	})
}

// skipping Fuzz_Nosy_StateDB_DumpToCollector__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-chain-ops/foundry.StateDB
