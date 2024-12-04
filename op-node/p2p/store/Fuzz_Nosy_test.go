package store

import (
	"testing"
	"github.com/ethereum-optimism/optimism/op-node/p2p"
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
	
	func Fuzz_Nosy_DecayApplicationScores_Apply__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var d *DecayApplicationScores
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var rec *scoreRecord
		fill_err = tp.Fill(&rec)
		if fill_err != nil {
			return
		}
	if d == nil|| rec == nil {
		return
	}

	d.Apply(rec)
	})
}

func Fuzz_Nosy_extendedStore_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *extendedStore
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.Close()
	})
}

func Fuzz_Nosy_ipBanBook_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var d *ipBanBook
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
	if d == nil {
		return
	}

	d.Close()
	})
}

func Fuzz_Nosy_ipBanBook_GetIPBanExpiration__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var d *ipBanBook
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var ip net.IP
		fill_err = tp.Fill(&ip)
		if fill_err != nil {
			return
		}
	if d == nil {
		return
	}

	d.GetIPBanExpiration(ip)
	})
}

func Fuzz_Nosy_ipBanBook_SetIPBanExpiration__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var d *ipBanBook
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var ip net.IP
		fill_err = tp.Fill(&ip)
		if fill_err != nil {
			return
		}
		var expirationTime time.Time
		fill_err = tp.Fill(&expirationTime)
		if fill_err != nil {
			return
		}
	if d == nil {
		return
	}

	d.SetIPBanExpiration(ip, expirationTime)
	})
}

func Fuzz_Nosy_ipBanBook_startGC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var d *ipBanBook
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
	if d == nil {
		return
	}

	d.startGC()
	})
}

func Fuzz_Nosy_ipBanRecord_LastUpdated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *ipBanRecord
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.LastUpdated()
	})
}

func Fuzz_Nosy_ipBanRecord_MarshalBinary__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *ipBanRecord
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.MarshalBinary()
	})
}

func Fuzz_Nosy_ipBanRecord_SetLastUpdated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *ipBanRecord
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var t2 time.Time
		fill_err = tp.Fill(&t2)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.SetLastUpdated(t2)
	})
}

func Fuzz_Nosy_ipBanRecord_UnmarshalBinary__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *ipBanRecord
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var d2 []byte
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.UnmarshalBinary(d2)
	})
}

func Fuzz_Nosy_metadataBook_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var m *metadataBook
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
	if m == nil {
		return
	}

	m.Close()
	})
}

func Fuzz_Nosy_metadataBook_GetPeerMetadata__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var m *metadataBook
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var id peer.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
	if m == nil {
		return
	}

	m.GetPeerMetadata(id)
	})
}

func Fuzz_Nosy_metadataBook_SetPeerMetadata__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var m *metadataBook
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var id peer.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var md PeerMetadata
		fill_err = tp.Fill(&md)
		if fill_err != nil {
			return
		}
	if m == nil {
		return
	}

	m.SetPeerMetadata(id, md)
	})
}

func Fuzz_Nosy_metadataBook_startGC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var m *metadataBook
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
	if m == nil {
		return
	}

	m.startGC()
	})
}

func Fuzz_Nosy_metadataRecord_Apply__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var rec *metadataRecord
		fill_err = tp.Fill(&rec)
		if fill_err != nil {
			return
		}
	if rec == nil {
		return
	}

	md := newMetadataRecord()
	md.Apply(rec)
	})
}

func Fuzz_Nosy_metadataRecord_SetLastUpdated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var t1 time.Time
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}

	m := newMetadataRecord()
	m.SetLastUpdated(t1)
	})
}

func Fuzz_Nosy_metadataRecord_UnmarshalBinary__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte) {
	m := newMetadataRecord()
	m.UnmarshalBinary(d1)
	})
}

func Fuzz_Nosy_peerBanBook_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var d *peerBanBook
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
	if d == nil {
		return
	}

	d.Close()
	})
}

func Fuzz_Nosy_peerBanBook_GetPeerBanExpiration__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var d *peerBanBook
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var id peer.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
	if d == nil {
		return
	}

	d.GetPeerBanExpiration(id)
	})
}

func Fuzz_Nosy_peerBanBook_SetPeerBanExpiration__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var d *peerBanBook
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var id peer.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var expirationTime time.Time
		fill_err = tp.Fill(&expirationTime)
		if fill_err != nil {
			return
		}
	if d == nil {
		return
	}

	d.SetPeerBanExpiration(id, expirationTime)
	})
}

func Fuzz_Nosy_peerBanBook_startGC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var d *peerBanBook
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
	if d == nil {
		return
	}

	d.startGC()
	})
}

func Fuzz_Nosy_peerBanRecord_LastUpdated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *peerBanRecord
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.LastUpdated()
	})
}

func Fuzz_Nosy_peerBanRecord_MarshalBinary__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *peerBanRecord
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.MarshalBinary()
	})
}

func Fuzz_Nosy_peerBanRecord_SetLastUpdated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *peerBanRecord
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var t2 time.Time
		fill_err = tp.Fill(&t2)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.SetLastUpdated(t2)
	})
}

func Fuzz_Nosy_peerBanRecord_UnmarshalBinary__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *peerBanRecord
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var d2 []byte
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.UnmarshalBinary(d2)
	})
}

func Fuzz_Nosy_recordsBook[K ~string, V record]_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var d *recordsBook[K, V]
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
	if d == nil {
		return
	}

	d.Close()
	})
}

// skipping Fuzz_Nosy_recordsBook[K ~string, V record]_deleteRecord__ because parameters include func, chan, or unsupported interface: K

// skipping Fuzz_Nosy_recordsBook[K ~string, V record]_dsKey__ because parameters include func, chan, or unsupported interface: K

// skipping Fuzz_Nosy_recordsBook[K ~string, V record]_getRecord__ because parameters include func, chan, or unsupported interface: K

// skipping Fuzz_Nosy_recordsBook[K ~string, V record]_hasExpired__ because parameters include func, chan, or unsupported interface: V

func Fuzz_Nosy_recordsBook[K ~string, V record]_prune__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var d *recordsBook[K, V]
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
	if d == nil {
		return
	}

	d.prune()
	})
}

// skipping Fuzz_Nosy_recordsBook[K ~string, V record]_setRecord__ because parameters include func, chan, or unsupported interface: K

func Fuzz_Nosy_recordsBook[K ~string, V record]_startGC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var d *recordsBook[K, V]
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
	if d == nil {
		return
	}

	d.startGC()
	})
}

func Fuzz_Nosy_scoreBook_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var d *scoreBook
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
	if d == nil {
		return
	}

	d.Close()
	})
}

func Fuzz_Nosy_scoreBook_GetPeerScore__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var d *scoreBook
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var id peer.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
	if d == nil {
		return
	}

	d.GetPeerScore(id)
	})
}

func Fuzz_Nosy_scoreBook_GetPeerScores__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var d *scoreBook
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var id peer.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
	if d == nil {
		return
	}

	d.GetPeerScores(id)
	})
}

// skipping Fuzz_Nosy_scoreBook_SetScore__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p/store.ScoreDiff

func Fuzz_Nosy_scoreBook_startGC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var d *scoreBook
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
	if d == nil {
		return
	}

	d.startGC()
	})
}

func Fuzz_Nosy_scoreRecord_LastUpdated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte) {
	s, err := deserializeScoresV0(d1)
	if err != nil {
		return
	}
	s.LastUpdated()
	})
}

func Fuzz_Nosy_scoreRecord_MarshalBinary__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte) {
	s, err := deserializeScoresV0(d1)
	if err != nil {
		return
	}
	s.MarshalBinary()
	})
}

func Fuzz_Nosy_scoreRecord_SetLastUpdated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var d1 []byte
		fill_err = tp.Fill(&d1)
		if fill_err != nil {
			return
		}
		var t2 time.Time
		fill_err = tp.Fill(&t2)
		if fill_err != nil {
			return
		}

	s, err := deserializeScoresV0(d1)
	if err != nil {
		return
	}
	s.SetLastUpdated(t2)
	})
}

func Fuzz_Nosy_scoreRecord_UnmarshalBinary__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte, d2 []byte) {
	s, err := deserializeScoresV0(d1)
	if err != nil {
		return
	}
	s.UnmarshalBinary(d2)
	})
}

func Fuzz_Nosy_GossipScores_Apply__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var g GossipScores
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var rec *scoreRecord
		fill_err = tp.Fill(&rec)
		if fill_err != nil {
			return
		}
	if rec == nil {
		return
	}

	g.Apply(rec)
	})
}

// skipping Fuzz_Nosy_IPBanStore_GetIPBanExpiration__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p/store.IPBanStore

// skipping Fuzz_Nosy_IPBanStore_SetIPBanExpiration__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p/store.IPBanStore

func Fuzz_Nosy_IncrementErrorResponses_Apply__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var i IncrementErrorResponses
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		var rec *scoreRecord
		fill_err = tp.Fill(&rec)
		if fill_err != nil {
			return
		}
	if rec == nil {
		return
	}

	i.Apply(rec)
	})
}

func Fuzz_Nosy_IncrementRejectedPayloads_Apply__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var i IncrementRejectedPayloads
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		var rec *scoreRecord
		fill_err = tp.Fill(&rec)
		if fill_err != nil {
			return
		}
	if rec == nil {
		return
	}

	i.Apply(rec)
	})
}

func Fuzz_Nosy_IncrementValidResponses_Apply__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var i IncrementValidResponses
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		var rec *scoreRecord
		fill_err = tp.Fill(&rec)
		if fill_err != nil {
			return
		}
	if rec == nil {
		return
	}

	i.Apply(rec)
	})
}

// skipping Fuzz_Nosy_MetadataStore_GetPeerMetadata__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p/store.MetadataStore

// skipping Fuzz_Nosy_MetadataStore_SetPeerMetadata__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p/store.MetadataStore

// skipping Fuzz_Nosy_PeerBanStore_GetPeerBanExpiration__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p/store.PeerBanStore

// skipping Fuzz_Nosy_PeerBanStore_SetPeerBanExpiration__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p/store.PeerBanStore

// skipping Fuzz_Nosy_ScoreDatastore_GetPeerScore__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p/store.ScoreDatastore

// skipping Fuzz_Nosy_ScoreDatastore_GetPeerScores__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p/store.ScoreDatastore

// skipping Fuzz_Nosy_ScoreDatastore_SetScore__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p/store.ScoreDatastore

// skipping Fuzz_Nosy_ScoreDiff_Apply__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p/store.ScoreDiff

func Fuzz_Nosy_ipBanUpdate_Apply__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var p ipBanUpdate
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var rec *ipBanRecord
		fill_err = tp.Fill(&rec)
		if fill_err != nil {
			return
		}
	if rec == nil {
		return
	}

	p.Apply(rec)
	})
}

func Fuzz_Nosy_peerBanUpdate_Apply__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var p peerBanUpdate
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var rec *peerBanRecord
		fill_err = tp.Fill(&rec)
		if fill_err != nil {
			return
		}
	if rec == nil {
		return
	}

	p.Apply(rec)
	})
}

// skipping Fuzz_Nosy_record_LastUpdated__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p/store.record

// skipping Fuzz_Nosy_record_SetLastUpdated__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p/store.record

// skipping Fuzz_Nosy_recordDiff[V record]_Apply__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/p2p/store.recordDiff[V github.com/ethereum-optimism/optimism/op-node/p2p/store.record]

func Fuzz_Nosy_serializeScoresV0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var scores scoreRecord
		fill_err = tp.Fill(&scores)
		if fill_err != nil {
			return
		}

	serializeScoresV0(scores)
	})
}

// skipping Fuzz_Nosy_startGc__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

