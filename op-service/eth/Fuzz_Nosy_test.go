package eth

import (
	"bytes"
	"io"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
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

func Fuzz_Nosy_APIBlobSidecar_BlobSidecar__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sc *APIBlobSidecar
		fill_err = tp.Fill(&sc)
		if fill_err != nil {
			return
		}
		if sc == nil {
			return
		}

		sc.BlobSidecar()
	})
}

func Fuzz_Nosy_AccountResult_Verify__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var res *AccountResult
		fill_err = tp.Fill(&res)
		if fill_err != nil {
			return
		}
		var stateRoot common.Hash
		fill_err = tp.Fill(&stateRoot)
		if fill_err != nil {
			return
		}
		if res == nil {
			return
		}

		res.Verify(stateRoot)
	})
}

func Fuzz_Nosy_Blob_Clear__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Blob
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.Clear()
	})
}

func Fuzz_Nosy_Blob_ComputeKZGCommitment__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Blob
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.ComputeKZGCommitment()
	})
}

func Fuzz_Nosy_Blob_FromData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Blob
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var d2 hexutil.Bytes
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.FromData(d2)
	})
}

func Fuzz_Nosy_Blob_KZGBlob__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Blob
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.KZGBlob()
	})
}

func Fuzz_Nosy_Blob_MarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Blob
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.MarshalText()
	})
}

func Fuzz_Nosy_Blob_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Blob
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.String()
	})
}

func Fuzz_Nosy_Blob_TerminalString__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Blob
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.TerminalString()
	})
}

func Fuzz_Nosy_Blob_ToData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Blob
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.ToData()
	})
}

func Fuzz_Nosy_Blob_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Blob
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var text []byte
		fill_err = tp.Fill(&text)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.UnmarshalJSON(text)
	})
}

func Fuzz_Nosy_Blob_UnmarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Blob
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var text []byte
		fill_err = tp.Fill(&text)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.UnmarshalText(text)
	})
}

func Fuzz_Nosy_Blob_decodeFieldElement__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Blob
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var opos int
		fill_err = tp.Fill(&opos)
		if fill_err != nil {
			return
		}
		var ipos int
		fill_err = tp.Fill(&ipos)
		if fill_err != nil {
			return
		}
		var output []byte
		fill_err = tp.Fill(&output)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.decodeFieldElement(opos, ipos, output)
	})
}

func Fuzz_Nosy_Bytes256_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Bytes256
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var text []byte
		fill_err = tp.Fill(&text)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.UnmarshalJSON(text)
	})
}

func Fuzz_Nosy_Bytes256_UnmarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Bytes256
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var text []byte
		fill_err = tp.Fill(&text)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.UnmarshalText(text)
	})
}

func Fuzz_Nosy_Bytes32_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Bytes32
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var text []byte
		fill_err = tp.Fill(&text)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.UnmarshalJSON(text)
	})
}

func Fuzz_Nosy_Bytes32_UnmarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Bytes32
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var text []byte
		fill_err = tp.Fill(&text)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.UnmarshalText(text)
	})
}

func Fuzz_Nosy_Bytes48_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Bytes48
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var text []byte
		fill_err = tp.Fill(&text)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.UnmarshalJSON(text)
	})
}

func Fuzz_Nosy_Bytes48_UnmarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Bytes48
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var text []byte
		fill_err = tp.Fill(&text)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.UnmarshalText(text)
	})
}

func Fuzz_Nosy_Bytes8_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Bytes8
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var text []byte
		fill_err = tp.Fill(&text)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.UnmarshalJSON(text)
	})
}

func Fuzz_Nosy_Bytes8_UnmarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Bytes8
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var text []byte
		fill_err = tp.Fill(&text)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.UnmarshalText(text)
	})
}

func Fuzz_Nosy_Bytes96_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Bytes96
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var text []byte
		fill_err = tp.Fill(&text)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.UnmarshalJSON(text)
	})
}

func Fuzz_Nosy_Bytes96_UnmarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Bytes96
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var text []byte
		fill_err = tp.Fill(&text)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.UnmarshalText(text)
	})
}

func Fuzz_Nosy_BytesMax32_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BytesMax32
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var text []byte
		fill_err = tp.Fill(&text)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.UnmarshalJSON(text)
	})
}

func Fuzz_Nosy_BytesMax32_UnmarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BytesMax32
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var text []byte
		fill_err = tp.Fill(&text)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.UnmarshalText(text)
	})
}

func Fuzz_Nosy_ExecutionPayload_CanyonBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bl *types.Block
		fill_err = tp.Fill(&bl)
		if fill_err != nil {
			return
		}
		var shanghaiTime *uint64
		fill_err = tp.Fill(&shanghaiTime)
		if fill_err != nil {
			return
		}
		if bl == nil || shanghaiTime == nil {
			return
		}

		payload, err := BlockAsPayload(bl, shanghaiTime)
		if err != nil {
			return
		}
		payload.CanyonBlock()
	})
}

func Fuzz_Nosy_ExecutionPayload_ID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bl *types.Block
		fill_err = tp.Fill(&bl)
		if fill_err != nil {
			return
		}
		var shanghaiTime *uint64
		fill_err = tp.Fill(&shanghaiTime)
		if fill_err != nil {
			return
		}
		if bl == nil || shanghaiTime == nil {
			return
		}

		payload, err := BlockAsPayload(bl, shanghaiTime)
		if err != nil {
			return
		}
		payload.ID()
	})
}

func Fuzz_Nosy_ExecutionPayload_MarshalSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bl *types.Block
		fill_err = tp.Fill(&bl)
		if fill_err != nil {
			return
		}
		var shanghaiTime *uint64
		fill_err = tp.Fill(&shanghaiTime)
		if fill_err != nil {
			return
		}
		var w io.Writer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if bl == nil || shanghaiTime == nil {
			return
		}

		payload, err := BlockAsPayload(bl, shanghaiTime)
		if err != nil {
			return
		}
		payload.MarshalSSZ(w)
	})
}

func Fuzz_Nosy_ExecutionPayload_ParentID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bl *types.Block
		fill_err = tp.Fill(&bl)
		if fill_err != nil {
			return
		}
		var shanghaiTime *uint64
		fill_err = tp.Fill(&shanghaiTime)
		if fill_err != nil {
			return
		}
		if bl == nil || shanghaiTime == nil {
			return
		}

		payload, err := BlockAsPayload(bl, shanghaiTime)
		if err != nil {
			return
		}
		payload.ParentID()
	})
}

func Fuzz_Nosy_ExecutionPayload_SizeSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bl *types.Block
		fill_err = tp.Fill(&bl)
		if fill_err != nil {
			return
		}
		var shanghaiTime *uint64
		fill_err = tp.Fill(&shanghaiTime)
		if fill_err != nil {
			return
		}
		if bl == nil || shanghaiTime == nil {
			return
		}

		payload, err := BlockAsPayload(bl, shanghaiTime)
		if err != nil {
			return
		}
		payload.SizeSSZ()
	})
}

func Fuzz_Nosy_ExecutionPayload_UnmarshalSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bl *types.Block
		fill_err = tp.Fill(&bl)
		if fill_err != nil {
			return
		}
		var shanghaiTime *uint64
		fill_err = tp.Fill(&shanghaiTime)
		if fill_err != nil {
			return
		}
		var version BlockVersion
		fill_err = tp.Fill(&version)
		if fill_err != nil {
			return
		}
		var scope uint32
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		var r io.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if bl == nil || shanghaiTime == nil {
			return
		}

		payload, err := BlockAsPayload(bl, shanghaiTime)
		if err != nil {
			return
		}
		payload.UnmarshalSSZ(version, scope, r)
	})
}

func Fuzz_Nosy_ExecutionPayload_inferVersion__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bl *types.Block
		fill_err = tp.Fill(&bl)
		if fill_err != nil {
			return
		}
		var shanghaiTime *uint64
		fill_err = tp.Fill(&shanghaiTime)
		if fill_err != nil {
			return
		}
		if bl == nil || shanghaiTime == nil {
			return
		}

		payload, err := BlockAsPayload(bl, shanghaiTime)
		if err != nil {
			return
		}
		payload.inferVersion()
	})
}

func Fuzz_Nosy_ExecutionPayload_transactionSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bl *types.Block
		fill_err = tp.Fill(&bl)
		if fill_err != nil {
			return
		}
		var shanghaiTime *uint64
		fill_err = tp.Fill(&shanghaiTime)
		if fill_err != nil {
			return
		}
		if bl == nil || shanghaiTime == nil {
			return
		}

		payload, err := BlockAsPayload(bl, shanghaiTime)
		if err != nil {
			return
		}
		payload.transactionSize()
	})
}

func Fuzz_Nosy_ExecutionPayload_withdrawalSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bl *types.Block
		fill_err = tp.Fill(&bl)
		if fill_err != nil {
			return
		}
		var shanghaiTime *uint64
		fill_err = tp.Fill(&shanghaiTime)
		if fill_err != nil {
			return
		}
		if bl == nil || shanghaiTime == nil {
			return
		}

		payload, err := BlockAsPayload(bl, shanghaiTime)
		if err != nil {
			return
		}
		payload.withdrawalSize()
	})
}

func Fuzz_Nosy_ExecutionPayloadEnvelope_CheckBlockHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bl *types.Block
		fill_err = tp.Fill(&bl)
		if fill_err != nil {
			return
		}
		var shanghaiTime *uint64
		fill_err = tp.Fill(&shanghaiTime)
		if fill_err != nil {
			return
		}
		if bl == nil || shanghaiTime == nil {
			return
		}

		envelope, err := BlockAsPayloadEnv(bl, shanghaiTime)
		if err != nil {
			return
		}
		envelope.CheckBlockHash()
	})
}

func Fuzz_Nosy_ExecutionPayloadEnvelope_MarshalSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bl *types.Block
		fill_err = tp.Fill(&bl)
		if fill_err != nil {
			return
		}
		var shanghaiTime *uint64
		fill_err = tp.Fill(&shanghaiTime)
		if fill_err != nil {
			return
		}
		var w io.Writer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if bl == nil || shanghaiTime == nil {
			return
		}

		envelope, err := BlockAsPayloadEnv(bl, shanghaiTime)
		if err != nil {
			return
		}
		envelope.MarshalSSZ(w)
	})
}

func Fuzz_Nosy_ExecutionPayloadEnvelope_UnmarshalSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bl *types.Block
		fill_err = tp.Fill(&bl)
		if fill_err != nil {
			return
		}
		var shanghaiTime *uint64
		fill_err = tp.Fill(&shanghaiTime)
		if fill_err != nil {
			return
		}
		var scope uint32
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		var r io.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if bl == nil || shanghaiTime == nil {
			return
		}

		envelope, err := BlockAsPayloadEnv(bl, shanghaiTime)
		if err != nil {
			return
		}
		envelope.UnmarshalSSZ(scope, r)
	})
}

func Fuzz_Nosy_OutputV0_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte) {
		o, err := unmarshalOutputV0(d1)
		if err != nil {
			return
		}
		o.Marshal()
	})
}

func Fuzz_Nosy_OutputV0_Version__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte) {
		o, err := unmarshalOutputV0(d1)
		if err != nil {
			return
		}
		o.Version()
	})
}

func Fuzz_Nosy_PayloadAttributes_IsDepositsOnly__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *PayloadAttributes
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if a == nil {
			return
		}

		a.IsDepositsOnly()
	})
}

func Fuzz_Nosy_PayloadAttributes_WithDepositsOnly__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *PayloadAttributes
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if a == nil {
			return
		}

		a.WithDepositsOnly()
	})
}

func Fuzz_Nosy_SystemConfig_EcotoneScalars__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sysCfg *SystemConfig
		fill_err = tp.Fill(&sysCfg)
		if fill_err != nil {
			return
		}
		if sysCfg == nil {
			return
		}

		sysCfg.EcotoneScalars()
	})
}

func Fuzz_Nosy_Uint64String_UnmarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v *Uint64String
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if v == nil {
			return
		}

		v.UnmarshalText(b)
	})
}

func Fuzz_Nosy_BlockID_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *types.Header
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		id := HeaderBlockID(h)
		id.String()
	})
}

func Fuzz_Nosy_BlockID_TerminalString__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *types.Header
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		id := HeaderBlockID(h)
		id.TerminalString()
	})
}

func Fuzz_Nosy_BlockInfo_BaseFee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *types.Header
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		_x1 := HeaderBlockInfo(h)
		_x1.BaseFee()
	})
}

func Fuzz_Nosy_BlockInfo_BlobBaseFee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *types.Header
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		_x1 := HeaderBlockInfo(h)
		_x1.BlobBaseFee()
	})
}

func Fuzz_Nosy_BlockInfo_Coinbase__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *types.Header
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		_x1 := HeaderBlockInfo(h)
		_x1.Coinbase()
	})
}

func Fuzz_Nosy_BlockInfo_GasLimit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *types.Header
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		_x1 := HeaderBlockInfo(h)
		_x1.GasLimit()
	})
}

func Fuzz_Nosy_BlockInfo_GasUsed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *types.Header
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		_x1 := HeaderBlockInfo(h)
		_x1.GasUsed()
	})
}

func Fuzz_Nosy_BlockInfo_Hash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *types.Header
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		_x1 := HeaderBlockInfo(h)
		_x1.Hash()
	})
}

func Fuzz_Nosy_BlockInfo_HeaderRLP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *types.Header
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		_x1 := HeaderBlockInfo(h)
		_x1.HeaderRLP()
	})
}

func Fuzz_Nosy_BlockInfo_MixDigest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *types.Header
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		_x1 := HeaderBlockInfo(h)
		_x1.MixDigest()
	})
}

func Fuzz_Nosy_BlockInfo_NumberU64__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *types.Header
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		_x1 := HeaderBlockInfo(h)
		_x1.NumberU64()
	})
}

func Fuzz_Nosy_BlockInfo_ParentBeaconRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *types.Header
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		_x1 := HeaderBlockInfo(h)
		_x1.ParentBeaconRoot()
	})
}

func Fuzz_Nosy_BlockInfo_ParentHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *types.Header
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		_x1 := HeaderBlockInfo(h)
		_x1.ParentHash()
	})
}

func Fuzz_Nosy_BlockInfo_ReceiptHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *types.Header
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		_x1 := HeaderBlockInfo(h)
		_x1.ReceiptHash()
	})
}

func Fuzz_Nosy_BlockInfo_Root__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *types.Header
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		_x1 := HeaderBlockInfo(h)
		_x1.Root()
	})
}

func Fuzz_Nosy_BlockInfo_Time__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *types.Header
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		_x1 := HeaderBlockInfo(h)
		_x1.Time()
	})
}

func Fuzz_Nosy_BlockLabel_Arg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var label BlockLabel
		fill_err = tp.Fill(&label)
		if fill_err != nil {
			return
		}

		label.Arg()
	})
}

func Fuzz_Nosy_BlockLabel_CheckID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 BlockLabel
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var id BlockID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}

		_x1.CheckID(id)
	})
}

func Fuzz_Nosy_BlockVersion_HasBlobProperties__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v BlockVersion
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}

		v.HasBlobProperties()
	})
}

func Fuzz_Nosy_BlockVersion_HasParentBeaconBlockRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v BlockVersion
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}

		v.HasParentBeaconBlockRoot()
	})
}

func Fuzz_Nosy_BlockVersion_HasWithdrawals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v BlockVersion
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}

		v.HasWithdrawals()
	})
}

func Fuzz_Nosy_Bytes256_MarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b Bytes256
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		b.MarshalText()
	})
}

func Fuzz_Nosy_Bytes256_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b Bytes256
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		b.String()
	})
}

func Fuzz_Nosy_Bytes256_TerminalString__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b Bytes256
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		b.TerminalString()
	})
}

func Fuzz_Nosy_Bytes32_MarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b Bytes32
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		b.MarshalText()
	})
}

func Fuzz_Nosy_Bytes32_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b Bytes32
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		b.String()
	})
}

func Fuzz_Nosy_Bytes32_TerminalString__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b Bytes32
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		b.TerminalString()
	})
}

func Fuzz_Nosy_Bytes48_MarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b Bytes48
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		b.MarshalText()
	})
}

func Fuzz_Nosy_Bytes48_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b Bytes48
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		b.String()
	})
}

func Fuzz_Nosy_Bytes48_TerminalString__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b Bytes48
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		b.TerminalString()
	})
}

func Fuzz_Nosy_Bytes8_MarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b Bytes8
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		b.MarshalText()
	})
}

func Fuzz_Nosy_Bytes8_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b Bytes8
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		b.String()
	})
}

func Fuzz_Nosy_Bytes8_TerminalString__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b Bytes8
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		b.TerminalString()
	})
}

func Fuzz_Nosy_Bytes96_MarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b Bytes96
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		b.MarshalText()
	})
}

func Fuzz_Nosy_Bytes96_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b Bytes96
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		b.String()
	})
}

func Fuzz_Nosy_Bytes96_TerminalString__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b Bytes96
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		b.TerminalString()
	})
}

func Fuzz_Nosy_BytesMax32_MarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b BytesMax32
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		b.MarshalText()
	})
}

func Fuzz_Nosy_BytesMax32_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b BytesMax32
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		b.String()
	})
}

func Fuzz_Nosy_ErrorCode_IsEngineError__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c ErrorCode
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}

		c.IsEngineError()
	})
}

func Fuzz_Nosy_ErrorCode_IsGenericRPCError__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c ErrorCode
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}

		c.IsGenericRPCError()
	})
}

func Fuzz_Nosy_InputError_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ie InputError
		fill_err = tp.Fill(&ie)
		if fill_err != nil {
			return
		}

		ie.Error()
	})
}

func Fuzz_Nosy_InputError_ErrorCode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ie InputError
		fill_err = tp.Fill(&ie)
		if fill_err != nil {
			return
		}

		ie.ErrorCode()
	})
}

// skipping Fuzz_Nosy_InputError_Is__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_InputError_Unwrap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ie InputError
		fill_err = tp.Fill(&ie)
		if fill_err != nil {
			return
		}

		ie.Unwrap()
	})
}

func Fuzz_Nosy_L1BlockRef_ID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id L1BlockRef
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}

		id.ID()
	})
}

func Fuzz_Nosy_L1BlockRef_ParentID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id L1BlockRef
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}

		id.ParentID()
	})
}

func Fuzz_Nosy_L1BlockRef_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id L1BlockRef
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}

		id.String()
	})
}

func Fuzz_Nosy_L1BlockRef_TerminalString__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id L1BlockRef
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}

		id.TerminalString()
	})
}

// skipping Fuzz_Nosy_L1BlockRefsSource_L1BlockRefByLabel__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/eth.L1BlockRefsSource

// skipping Fuzz_Nosy_L1Client_HeaderByNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/eth.L1Client

// skipping Fuzz_Nosy_L1Client_NonceAt__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/eth.L1Client

func Fuzz_Nosy_L2BlockRef_BlockRef__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id L2BlockRef
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}

		id.BlockRef()
	})
}

func Fuzz_Nosy_L2BlockRef_ID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id L2BlockRef
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}

		id.ID()
	})
}

func Fuzz_Nosy_L2BlockRef_ParentID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id L2BlockRef
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}

		id.ParentID()
	})
}

func Fuzz_Nosy_L2BlockRef_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id L2BlockRef
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}

		id.String()
	})
}

func Fuzz_Nosy_L2BlockRef_TerminalString__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id L2BlockRef
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}

		id.TerminalString()
	})
}

// skipping Fuzz_Nosy_NewHeadSource_SubscribeNewHead__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/eth.NewHeadSource

// skipping Fuzz_Nosy_NumberAndHash_Hash__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/eth.NumberAndHash

// skipping Fuzz_Nosy_NumberAndHash_NumberU64__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/eth.NumberAndHash

func Fuzz_Nosy_Output_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte) {
		_x1, err := UnmarshalOutput(d1)
		if err != nil {
			return
		}
		_x1.Marshal()
	})
}

func Fuzz_Nosy_Output_Version__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte) {
		_x1, err := UnmarshalOutput(d1)
		if err != nil {
			return
		}
		_x1.Version()
	})
}

func Fuzz_Nosy_SystemConfig_MarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sysCfg SystemConfig
		fill_err = tp.Fill(&sysCfg)
		if fill_err != nil {
			return
		}

		sysCfg.MarshalJSON()
	})
}

func Fuzz_Nosy_Uint64String_MarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v Uint64String
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}

		v.MarshalText()
	})
}

func Fuzz_Nosy_blockInfo_BlobBaseFee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b blockInfo
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		b.BlobBaseFee()
	})
}

func Fuzz_Nosy_blockInfo_HeaderRLP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b blockInfo
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		b.HeaderRLP()
	})
}

func Fuzz_Nosy_blockInfo_ParentBeaconRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b blockInfo
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		b.ParentBeaconRoot()
	})
}

func Fuzz_Nosy_headerBlockInfo_BaseFee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h headerBlockInfo
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}

		h.BaseFee()
	})
}

func Fuzz_Nosy_headerBlockInfo_BlobBaseFee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h headerBlockInfo
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}

		h.BlobBaseFee()
	})
}

func Fuzz_Nosy_headerBlockInfo_Coinbase__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h headerBlockInfo
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}

		h.Coinbase()
	})
}

func Fuzz_Nosy_headerBlockInfo_GasLimit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h headerBlockInfo
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}

		h.GasLimit()
	})
}

func Fuzz_Nosy_headerBlockInfo_GasUsed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h headerBlockInfo
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}

		h.GasUsed()
	})
}

func Fuzz_Nosy_headerBlockInfo_HeaderRLP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h headerBlockInfo
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}

		h.HeaderRLP()
	})
}

func Fuzz_Nosy_headerBlockInfo_MixDigest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h headerBlockInfo
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}

		h.MixDigest()
	})
}

func Fuzz_Nosy_headerBlockInfo_NumberU64__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h headerBlockInfo
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}

		h.NumberU64()
	})
}

func Fuzz_Nosy_headerBlockInfo_ParentBeaconRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h headerBlockInfo
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}

		h.ParentBeaconRoot()
	})
}

func Fuzz_Nosy_headerBlockInfo_ParentHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h headerBlockInfo
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}

		h.ParentHash()
	})
}

func Fuzz_Nosy_headerBlockInfo_ReceiptHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h headerBlockInfo
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}

		h.ReceiptHash()
	})
}

func Fuzz_Nosy_headerBlockInfo_Root__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h headerBlockInfo
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}

		h.Root()
	})
}

func Fuzz_Nosy_headerBlockInfo_Time__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h headerBlockInfo
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}

		h.Time()
	})
}

func Fuzz_Nosy_rawTransactions_EncodeIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s rawTransactions
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var i int
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		var w *bytes.Buffer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		s.EncodeIndex(i, w)
	})
}

func Fuzz_Nosy_rawTransactions_Len__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s rawTransactions
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}

		s.Len()
	})
}

// skipping Fuzz_Nosy_CheckRecentTxs__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/eth.L1Client

func Fuzz_Nosy_DecodeRawReceipts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var block BlockID
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		var rawReceipts []hexutil.Bytes
		fill_err = tp.Fill(&rawReceipts)
		if fill_err != nil {
			return
		}
		var txHashes []common.Hash
		fill_err = tp.Fill(&txHashes)
		if fill_err != nil {
			return
		}

		DecodeRawReceipts(block, rawReceipts, txHashes)
	})
}

func Fuzz_Nosy_DecodeTransactions__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d1 []hexutil.Bytes
		fill_err = tp.Fill(&d1)
		if fill_err != nil {
			return
		}

		DecodeTransactions(d1)
	})
}

func Fuzz_Nosy_EncodeReceipts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var elems []*types.Receipt
		fill_err = tp.Fill(&elems)
		if fill_err != nil {
			return
		}

		EncodeReceipts(elems)
	})
}

func Fuzz_Nosy_EncodeScalar__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var scalars EcotoneScalars
		fill_err = tp.Fill(&scalars)
		if fill_err != nil {
			return
		}

		EncodeScalar(scalars)
	})
}

func Fuzz_Nosy_EncodeTransactions__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var elems []*types.Transaction
		fill_err = tp.Fill(&elems)
		if fill_err != nil {
			return
		}

		EncodeTransactions(elems)
	})
}

func Fuzz_Nosy_TransactionsToHashes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var elems []*types.Transaction
		fill_err = tp.Fill(&elems)
		if fill_err != nil {
			return
		}

		TransactionsToHashes(elems)
	})
}

func Fuzz_Nosy_WeiToEther__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var wei *big.Int
		fill_err = tp.Fill(&wei)
		if fill_err != nil {
			return
		}
		if wei == nil {
			return
		}

		WeiToEther(wei)
	})
}

func Fuzz_Nosy_executionPayloadFixedPart__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var version BlockVersion
		fill_err = tp.Fill(&version)
		if fill_err != nil {
			return
		}

		executionPayloadFixedPart(version)
	})
}

func Fuzz_Nosy_jsonMarshalHolocene__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sysCfg SystemConfig
		fill_err = tp.Fill(&sysCfg)
		if fill_err != nil {
			return
		}

		jsonMarshalHolocene(sysCfg)
	})
}

func Fuzz_Nosy_jsonMarshalPreHolocene__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sysCfg SystemConfig
		fill_err = tp.Fill(&sysCfg)
		if fill_err != nil {
			return
		}

		jsonMarshalPreHolocene(sysCfg)
	})
}

func Fuzz_Nosy_marshalBytes32LE__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var out []byte
		fill_err = tp.Fill(&out)
		if fill_err != nil {
			return
		}
		var z *hexutil.U256
		fill_err = tp.Fill(&z)
		if fill_err != nil {
			return
		}
		if z == nil {
			return
		}

		marshalBytes32LE(out, z)
	})
}

func Fuzz_Nosy_marshalTransactions__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var out []byte
		fill_err = tp.Fill(&out)
		if fill_err != nil {
			return
		}
		var txs []hexutil.Bytes
		fill_err = tp.Fill(&txs)
		if fill_err != nil {
			return
		}

		marshalTransactions(out, txs)
	})
}

func Fuzz_Nosy_marshalWithdrawals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var out []byte
		fill_err = tp.Fill(&out)
		if fill_err != nil {
			return
		}
		var withdrawals types.Withdrawals
		fill_err = tp.Fill(&withdrawals)
		if fill_err != nil {
			return
		}

		marshalWithdrawals(out, withdrawals)
	})
}

func Fuzz_Nosy_reassembleBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, opos int, encodedByte []byte, output []byte) {
		reassembleBytes(opos, encodedByte, output)
	})
}

func Fuzz_Nosy_unmarshalBytes32LE__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var in []byte
		fill_err = tp.Fill(&in)
		if fill_err != nil {
			return
		}
		var z *hexutil.U256
		fill_err = tp.Fill(&z)
		if fill_err != nil {
			return
		}
		if z == nil {
			return
		}

		unmarshalBytes32LE(in, z)
	})
}

func Fuzz_Nosy_unmarshalTransactions__(f *testing.F) {
	f.Fuzz(func(t *testing.T, in []byte) {
		unmarshalTransactions(in)
	})
}
