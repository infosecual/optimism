package tools

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
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

func Fuzz_Nosy_GameCreator_CreateGame__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *GameCreator
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var outputRoot common.Hash
		fill_err = tp.Fill(&outputRoot)
		if fill_err != nil {
			return
		}
		var traceType uint64
		fill_err = tp.Fill(&traceType)
		if fill_err != nil {
			return
		}
		var l2BlockNum uint64
		fill_err = tp.Fill(&l2BlockNum)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.CreateGame(ctx, outputRoot, traceType, l2BlockNum)
	})
}
