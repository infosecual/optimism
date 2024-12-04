package types

import (
	"math/big"
	"testing"
	"time"

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

func Fuzz_Nosy_ClaimData_ValueBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *ClaimData
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.ValueBytes()
	})
}

func Fuzz_Nosy_PreimageOracleData_GetIdent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, key []byte, d2 []byte, offset uint32) {
		p := NewPreimageOracleData(key, d2, offset)
		p.GetIdent()
	})
}

func Fuzz_Nosy_PreimageOracleData_GetPrecompileAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, key []byte, d2 []byte, offset uint32) {
		p := NewPreimageOracleData(key, d2, offset)
		p.GetPrecompileAddress()
	})
}

func Fuzz_Nosy_PreimageOracleData_GetPrecompileInput__(f *testing.F) {
	f.Fuzz(func(t *testing.T, key []byte, d2 []byte, offset uint32) {
		p := NewPreimageOracleData(key, d2, offset)
		p.GetPrecompileInput()
	})
}

func Fuzz_Nosy_PreimageOracleData_GetPrecompileRequiredGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, key []byte, d2 []byte, offset uint32) {
		p := NewPreimageOracleData(key, d2, offset)
		p.GetPrecompileRequiredGas()
	})
}

func Fuzz_Nosy_PreimageOracleData_GetPreimageWithSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, key []byte, d2 []byte, offset uint32) {
		p := NewPreimageOracleData(key, d2, offset)
		p.GetPreimageWithSize()
	})
}

func Fuzz_Nosy_PreimageOracleData_GetPreimageWithoutSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, key []byte, d2 []byte, offset uint32) {
		p := NewPreimageOracleData(key, d2, offset)
		p.GetPreimageWithoutSize()
	})
}

func Fuzz_Nosy_TraceType_Clone__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *TraceType
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Clone()
	})
}

func Fuzz_Nosy_TraceType_Set__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *TraceType
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var value string
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Set(value)
	})
}

func Fuzz_Nosy_gameState_AgreeWithClaimLevel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var claims []Claim
		fill_err = tp.Fill(&claims)
		if fill_err != nil {
			return
		}
		var depth Depth
		fill_err = tp.Fill(&depth)
		if fill_err != nil {
			return
		}
		var claim Claim
		fill_err = tp.Fill(&claim)
		if fill_err != nil {
			return
		}
		var agreeWithRootClaim bool
		fill_err = tp.Fill(&agreeWithRootClaim)
		if fill_err != nil {
			return
		}

		g := NewGameState(claims, depth)
		g.AgreeWithClaimLevel(claim, agreeWithRootClaim)
	})
}

func Fuzz_Nosy_gameState_AncestorWithTraceIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var claims []Claim
		fill_err = tp.Fill(&claims)
		if fill_err != nil {
			return
		}
		var depth Depth
		fill_err = tp.Fill(&depth)
		if fill_err != nil {
			return
		}
		var claim Claim
		fill_err = tp.Fill(&claim)
		if fill_err != nil {
			return
		}
		var idx *big.Int
		fill_err = tp.Fill(&idx)
		if fill_err != nil {
			return
		}
		if idx == nil {
			return
		}

		g := NewGameState(claims, depth)
		g.AncestorWithTraceIndex(claim, idx)
	})
}

func Fuzz_Nosy_gameState_ChessClock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var claims []Claim
		fill_err = tp.Fill(&claims)
		if fill_err != nil {
			return
		}
		var depth Depth
		fill_err = tp.Fill(&depth)
		if fill_err != nil {
			return
		}
		var now time.Time
		fill_err = tp.Fill(&now)
		if fill_err != nil {
			return
		}
		var claim Claim
		fill_err = tp.Fill(&claim)
		if fill_err != nil {
			return
		}

		g := NewGameState(claims, depth)
		g.ChessClock(now, claim)
	})
}

func Fuzz_Nosy_gameState_Claims__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var claims []Claim
		fill_err = tp.Fill(&claims)
		if fill_err != nil {
			return
		}
		var depth Depth
		fill_err = tp.Fill(&depth)
		if fill_err != nil {
			return
		}

		g := NewGameState(claims, depth)
		g.Claims()
	})
}

func Fuzz_Nosy_gameState_DefendsParent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var claims []Claim
		fill_err = tp.Fill(&claims)
		if fill_err != nil {
			return
		}
		var depth Depth
		fill_err = tp.Fill(&depth)
		if fill_err != nil {
			return
		}
		var claim Claim
		fill_err = tp.Fill(&claim)
		if fill_err != nil {
			return
		}

		g := NewGameState(claims, depth)
		g.DefendsParent(claim)
	})
}

func Fuzz_Nosy_gameState_GetParent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var claims []Claim
		fill_err = tp.Fill(&claims)
		if fill_err != nil {
			return
		}
		var depth Depth
		fill_err = tp.Fill(&depth)
		if fill_err != nil {
			return
		}
		var claim Claim
		fill_err = tp.Fill(&claim)
		if fill_err != nil {
			return
		}

		g := NewGameState(claims, depth)
		g.GetParent(claim)
	})
}

func Fuzz_Nosy_gameState_IsDuplicate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var claims []Claim
		fill_err = tp.Fill(&claims)
		if fill_err != nil {
			return
		}
		var depth Depth
		fill_err = tp.Fill(&depth)
		if fill_err != nil {
			return
		}
		var claim Claim
		fill_err = tp.Fill(&claim)
		if fill_err != nil {
			return
		}

		g := NewGameState(claims, depth)
		g.IsDuplicate(claim)
	})
}

func Fuzz_Nosy_gameState_MaxDepth__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var claims []Claim
		fill_err = tp.Fill(&claims)
		if fill_err != nil {
			return
		}
		var depth Depth
		fill_err = tp.Fill(&depth)
		if fill_err != nil {
			return
		}

		g := NewGameState(claims, depth)
		g.MaxDepth()
	})
}

func Fuzz_Nosy_gameState_getParent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var claims []Claim
		fill_err = tp.Fill(&claims)
		if fill_err != nil {
			return
		}
		var depth Depth
		fill_err = tp.Fill(&depth)
		if fill_err != nil {
			return
		}
		var claim Claim
		fill_err = tp.Fill(&claim)
		if fill_err != nil {
			return
		}

		g := NewGameState(claims, depth)
		g.getParent(claim)
	})
}

func Fuzz_Nosy_ActionType_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a ActionType
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}

		a.String()
	})
}

func Fuzz_Nosy_Claim_ID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c Claim
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}

		c.ID()
	})
}

func Fuzz_Nosy_Claim_IsRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c Claim
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}

		c.IsRoot()
	})
}

// skipping Fuzz_Nosy_ClockReader_Now__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/types.ClockReader

// skipping Fuzz_Nosy_Game_AgreeWithClaimLevel__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/types.Game

// skipping Fuzz_Nosy_Game_AncestorWithTraceIndex__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/types.Game

// skipping Fuzz_Nosy_Game_ChessClock__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/types.Game

// skipping Fuzz_Nosy_Game_Claims__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/types.Game

// skipping Fuzz_Nosy_Game_DefendsParent__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/types.Game

// skipping Fuzz_Nosy_Game_GetParent__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/types.Game

// skipping Fuzz_Nosy_Game_IsDuplicate__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/types.Game

// skipping Fuzz_Nosy_Game_MaxDepth__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/types.Game

func Fuzz_Nosy_GameType_MarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 GameType
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}

		t1.MarshalText()
	})
}

func Fuzz_Nosy_GameType_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 GameType
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}

		t1.String()
	})
}

func Fuzz_Nosy_Position_Attack__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var depth Depth
		fill_err = tp.Fill(&depth)
		if fill_err != nil {
			return
		}
		var indexAtDepth *big.Int
		fill_err = tp.Fill(&indexAtDepth)
		if fill_err != nil {
			return
		}
		if indexAtDepth == nil {
			return
		}

		p := NewPosition(depth, indexAtDepth)
		p.Attack()
	})
}

func Fuzz_Nosy_Position_Defend__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var depth Depth
		fill_err = tp.Fill(&depth)
		if fill_err != nil {
			return
		}
		var indexAtDepth *big.Int
		fill_err = tp.Fill(&indexAtDepth)
		if fill_err != nil {
			return
		}
		if indexAtDepth == nil {
			return
		}

		p := NewPosition(depth, indexAtDepth)
		p.Defend()
	})
}

func Fuzz_Nosy_Position_Depth__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var depth Depth
		fill_err = tp.Fill(&depth)
		if fill_err != nil {
			return
		}
		var indexAtDepth *big.Int
		fill_err = tp.Fill(&indexAtDepth)
		if fill_err != nil {
			return
		}
		if indexAtDepth == nil {
			return
		}

		p := NewPosition(depth, indexAtDepth)
		p.Depth()
	})
}

func Fuzz_Nosy_Position_IndexAtDepth__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var depth Depth
		fill_err = tp.Fill(&depth)
		if fill_err != nil {
			return
		}
		var indexAtDepth *big.Int
		fill_err = tp.Fill(&indexAtDepth)
		if fill_err != nil {
			return
		}
		if indexAtDepth == nil {
			return
		}

		p := NewPosition(depth, indexAtDepth)
		p.IndexAtDepth()
	})
}

func Fuzz_Nosy_Position_IsRootPosition__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var depth Depth
		fill_err = tp.Fill(&depth)
		if fill_err != nil {
			return
		}
		var indexAtDepth *big.Int
		fill_err = tp.Fill(&indexAtDepth)
		if fill_err != nil {
			return
		}
		if indexAtDepth == nil {
			return
		}

		p := NewPosition(depth, indexAtDepth)
		p.IsRootPosition()
	})
}

func Fuzz_Nosy_Position_MoveRight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var depth Depth
		fill_err = tp.Fill(&depth)
		if fill_err != nil {
			return
		}
		var indexAtDepth *big.Int
		fill_err = tp.Fill(&indexAtDepth)
		if fill_err != nil {
			return
		}
		if indexAtDepth == nil {
			return
		}

		p := NewPosition(depth, indexAtDepth)
		p.MoveRight()
	})
}

func Fuzz_Nosy_Position_Print__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var depth Depth
		fill_err = tp.Fill(&depth)
		if fill_err != nil {
			return
		}
		var indexAtDepth *big.Int
		fill_err = tp.Fill(&indexAtDepth)
		if fill_err != nil {
			return
		}
		var maxDepth Depth
		fill_err = tp.Fill(&maxDepth)
		if fill_err != nil {
			return
		}
		if indexAtDepth == nil {
			return
		}

		p := NewPosition(depth, indexAtDepth)
		p.Print(maxDepth)
	})
}

func Fuzz_Nosy_Position_RelativeToAncestorAtDepth__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var depth Depth
		fill_err = tp.Fill(&depth)
		if fill_err != nil {
			return
		}
		var indexAtDepth *big.Int
		fill_err = tp.Fill(&indexAtDepth)
		if fill_err != nil {
			return
		}
		var ancestor Depth
		fill_err = tp.Fill(&ancestor)
		if fill_err != nil {
			return
		}
		if indexAtDepth == nil {
			return
		}

		p := NewPosition(depth, indexAtDepth)
		p.RelativeToAncestorAtDepth(ancestor)
	})
}

func Fuzz_Nosy_Position_RightOf__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var depth Depth
		fill_err = tp.Fill(&depth)
		if fill_err != nil {
			return
		}
		var indexAtDepth *big.Int
		fill_err = tp.Fill(&indexAtDepth)
		if fill_err != nil {
			return
		}
		var parent Position
		fill_err = tp.Fill(&parent)
		if fill_err != nil {
			return
		}
		if indexAtDepth == nil {
			return
		}

		p := NewPosition(depth, indexAtDepth)
		p.RightOf(parent)
	})
}

func Fuzz_Nosy_Position_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var depth Depth
		fill_err = tp.Fill(&depth)
		if fill_err != nil {
			return
		}
		var indexAtDepth *big.Int
		fill_err = tp.Fill(&indexAtDepth)
		if fill_err != nil {
			return
		}
		if indexAtDepth == nil {
			return
		}

		p := NewPosition(depth, indexAtDepth)
		p.String()
	})
}

func Fuzz_Nosy_Position_ToGIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var depth Depth
		fill_err = tp.Fill(&depth)
		if fill_err != nil {
			return
		}
		var indexAtDepth *big.Int
		fill_err = tp.Fill(&indexAtDepth)
		if fill_err != nil {
			return
		}
		if indexAtDepth == nil {
			return
		}

		p := NewPosition(depth, indexAtDepth)
		p.ToGIndex()
	})
}

func Fuzz_Nosy_Position_TraceIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var depth Depth
		fill_err = tp.Fill(&depth)
		if fill_err != nil {
			return
		}
		var indexAtDepth *big.Int
		fill_err = tp.Fill(&indexAtDepth)
		if fill_err != nil {
			return
		}
		var maxDepth Depth
		fill_err = tp.Fill(&maxDepth)
		if fill_err != nil {
			return
		}
		if indexAtDepth == nil {
			return
		}

		p := NewPosition(depth, indexAtDepth)
		p.TraceIndex(maxDepth)
	})
}

func Fuzz_Nosy_Position_lshIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var depth Depth
		fill_err = tp.Fill(&depth)
		if fill_err != nil {
			return
		}
		var indexAtDepth *big.Int
		fill_err = tp.Fill(&indexAtDepth)
		if fill_err != nil {
			return
		}
		var amount Depth
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if indexAtDepth == nil {
			return
		}

		p := NewPosition(depth, indexAtDepth)
		p.lshIndex(amount)
	})
}

func Fuzz_Nosy_Position_move__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var depth Depth
		fill_err = tp.Fill(&depth)
		if fill_err != nil {
			return
		}
		var indexAtDepth *big.Int
		fill_err = tp.Fill(&indexAtDepth)
		if fill_err != nil {
			return
		}
		var right bool
		fill_err = tp.Fill(&right)
		if fill_err != nil {
			return
		}
		if indexAtDepth == nil {
			return
		}

		p := NewPosition(depth, indexAtDepth)
		p.move(right)
	})
}

func Fuzz_Nosy_Position_parent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var depth Depth
		fill_err = tp.Fill(&depth)
		if fill_err != nil {
			return
		}
		var indexAtDepth *big.Int
		fill_err = tp.Fill(&indexAtDepth)
		if fill_err != nil {
			return
		}
		if indexAtDepth == nil {
			return
		}

		p := NewPosition(depth, indexAtDepth)
		p.parent()
	})
}

func Fuzz_Nosy_Position_parentIndexAtDepth__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var depth Depth
		fill_err = tp.Fill(&depth)
		if fill_err != nil {
			return
		}
		var indexAtDepth *big.Int
		fill_err = tp.Fill(&indexAtDepth)
		if fill_err != nil {
			return
		}
		if indexAtDepth == nil {
			return
		}

		p := NewPosition(depth, indexAtDepth)
		p.parentIndexAtDepth()
	})
}

// skipping Fuzz_Nosy_PrestateProvider_AbsolutePreStateCommitment__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/types.PrestateProvider

// skipping Fuzz_Nosy_TraceAccessor_Get__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/types.TraceAccessor

// skipping Fuzz_Nosy_TraceAccessor_GetL2BlockNumberChallenge__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/types.TraceAccessor

// skipping Fuzz_Nosy_TraceAccessor_GetStepData__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/types.TraceAccessor

// skipping Fuzz_Nosy_TraceProvider_Get__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/types.TraceProvider

// skipping Fuzz_Nosy_TraceProvider_GetL2BlockNumberChallenge__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/types.TraceProvider

// skipping Fuzz_Nosy_TraceProvider_GetStepData__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/types.TraceProvider

func Fuzz_Nosy_TraceType_GameType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 TraceType
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}

		t1.GameType()
	})
}

func Fuzz_Nosy_TraceType_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 TraceType
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}

		t1.String()
	})
}

func Fuzz_Nosy_ValidTraceType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var value TraceType
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}

		ValidTraceType(value)
	})
}

func Fuzz_Nosy_boolToInt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b bool) {
		boolToInt(b)
	})
}
