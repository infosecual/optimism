package certman

import (
	"testing"

	"github.com/ethereum-optimism/optimism/op-service/tls"
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

func Fuzz_Nosy_CertMan_GetCertificate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cm *CertMan
		fill_err = tp.Fill(&cm)
		if fill_err != nil {
			return
		}
		var hello *tls.ClientHelloInfo
		fill_err = tp.Fill(&hello)
		if fill_err != nil {
			return
		}
		if cm == nil || hello == nil {
			return
		}

		cm.GetCertificate(hello)
	})
}

func Fuzz_Nosy_CertMan_GetClientCertificate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cm *CertMan
		fill_err = tp.Fill(&cm)
		if fill_err != nil {
			return
		}
		var hello *tls.CertificateRequestInfo
		fill_err = tp.Fill(&hello)
		if fill_err != nil {
			return
		}
		if cm == nil || hello == nil {
			return
		}

		cm.GetClientCertificate(hello)
	})
}

func Fuzz_Nosy_CertMan_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cm *CertMan
		fill_err = tp.Fill(&cm)
		if fill_err != nil {
			return
		}
		if cm == nil {
			return
		}

		cm.Stop()
	})
}

func Fuzz_Nosy_CertMan_Watch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cm *CertMan
		fill_err = tp.Fill(&cm)
		if fill_err != nil {
			return
		}
		if cm == nil {
			return
		}

		cm.Watch()
	})
}

func Fuzz_Nosy_CertMan_load__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cm *CertMan
		fill_err = tp.Fill(&cm)
		if fill_err != nil {
			return
		}
		if cm == nil {
			return
		}

		cm.load()
	})
}

func Fuzz_Nosy_CertMan_run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cm *CertMan
		fill_err = tp.Fill(&cm)
		if fill_err != nil {
			return
		}
		if cm == nil {
			return
		}

		cm.run()
	})
}
