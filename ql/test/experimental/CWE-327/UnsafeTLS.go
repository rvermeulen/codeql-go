package main

import (
	"crypto/tls"
)

func main() {}

func minMaxTlsVersion() {
	{
		config := &tls.Config{}
		config.MinVersion = 0 // BAD
	}
	{
		config := &tls.Config{}
		config.MaxVersion = 0 // GOOD
	}
	///
	{
		config := &tls.Config{
			MinVersion: 0, // BAD
		}
		_ = config
	}
	{
		config := &tls.Config{
			MaxVersion: 0, // GOOD
		}
		_ = config
	}
	///
	{
		config := &tls.Config{}
		config.MinVersion = tls.VersionSSL30 // BAD
	}
	{
		config := &tls.Config{}
		config.MaxVersion = tls.VersionSSL30 // BAD
	}
	///
	{
		config := &tls.Config{}
		config.MinVersion = tls.VersionTLS10 // BAD
	}
	{
		config := &tls.Config{}
		config.MaxVersion = tls.VersionTLS10 // BAD
	}
	///
	{
		config := &tls.Config{}
		config.MinVersion = tls.VersionTLS11 // BAD
	}
	{
		config := &tls.Config{}
		config.MaxVersion = tls.VersionTLS11 // BAD
	}
	///
	{
		config := &tls.Config{
			MinVersion: tls.VersionTLS11, // BAD
		}
		_ = config
	}
	{
		config := &tls.Config{
			MaxVersion: tls.VersionTLS11, // BAD
		}
		_ = config
	}
	{
		config := &tls.Config{
			MinVersion: tls.VersionTLS12, // GOOD
		}
		_ = config
	}
	{
		config := &tls.Config{
			MaxVersion: tls.VersionTLS13, // GOOD
		}
		_ = config
	}
	///
	{
		config := &tls.Config{
			MinVersion: 0x0300, // BAD
		}
		_ = config
	}
	{
		config := &tls.Config{
			MaxVersion: 0x0301, // BAD
		}
		_ = config
	}
}

func cipherSuites() {
	{
		config := &tls.Config{
			CipherSuites: []uint16{
				tls.TLS_RSA_WITH_RC4_128_SHA,                // BAD
				tls.TLS_RSA_WITH_AES_128_CBC_SHA256,         // BAD
				tls.TLS_ECDHE_ECDSA_WITH_RC4_128_SHA,        // BAD
				tls.TLS_ECDHE_RSA_WITH_RC4_128_SHA,          // BAD
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256, // BAD
				tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256,   // BAD
			},
		}
		_ = config
	}
	{
		config := &tls.Config{
			CipherSuites: []uint16{
				tls.TLS_RSA_WITH_RC4_128_SHA, // BAD
			},
		}
		_ = config
	}
	{
		config := &tls.Config{
			CipherSuites: []uint16{
				tls.TLS_RSA_WITH_AES_128_CBC_SHA256, // BAD
			},
		}
		_ = config
	}
	{
		config := &tls.Config{
			CipherSuites: []uint16{
				tls.TLS_ECDHE_ECDSA_WITH_RC4_128_SHA, // BAD
			},
		}
		_ = config
	}
	{
		config := &tls.Config{
			CipherSuites: []uint16{
				tls.TLS_ECDHE_RSA_WITH_RC4_128_SHA, // BAD
			},
		}
		_ = config
	}
	{
		config := &tls.Config{
			CipherSuites: []uint16{
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256, // BAD
			},
		}
		_ = config
	}
	{
		config := &tls.Config{
			CipherSuites: []uint16{
				tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256, // BAD
			},
		}
		_ = config
	}
	{
		config := &tls.Config{
			CipherSuites: []uint16{
				tls.TLS_CHACHA20_POLY1305_SHA256, // GOOD
			},
		}
		_ = config
	}
	{
		config := &tls.Config{}
		config.CipherSuites = make([]uint16, 0)
		config.CipherSuites = append(config.CipherSuites, tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256) // BAD
	}
	{
		config := &tls.Config{}
		config.CipherSuites = make([]uint16, 0)
		insecureSuites := tls.InsecureCipherSuites()
		for _, v := range insecureSuites {
			config.CipherSuites = append(config.CipherSuites, v.ID) // TODO: should be flagged as BAD.
		}
	}
	{
		config := &tls.Config{}
		cipherSuites := make([]uint16, 0)
		insecureSuites := tls.InsecureCipherSuites()
		for _, v := range insecureSuites {
			cipherSuites = append(cipherSuites, v.ID)
		}
		config.CipherSuites = cipherSuites // TODO: should be flagged as BAD.
	}
	{
		config := &tls.Config{}
		cipherSuites := make([]uint16, 0)
		insecureSuites := tls.InsecureCipherSuites()
		for i := range insecureSuites {
			cipherSuites = append(cipherSuites, insecureSuites[i].ID)
		}
		config.CipherSuites = cipherSuites // BAD
	}
}
