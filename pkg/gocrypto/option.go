package gocrypto

import "crypto"

const (
	modeECB = "ECB"
	modeCBC = "CBC"
	modeCFB = "CFB"
	modeCTR = "CTR"
)

var (
	defaultAesKey = []byte("mKoF_pL,NjI9=I;w") // aes key
	defaultDesKey = []byte("VgY7*uHb")         // des key
	defaultMode   = "ECB"

	defaultRsaFormat   = "PKCS#1"
	defaultRsaHashType = crypto.SHA1
)

type aesOptions struct {
	// the length of the key must be one of 16,24,32, corresponding to
	// AES-128,AES-192,AES-256 respectively.
	aesKey []byte
	// there are four operating modes in total, ECB CBC CFB CTR
	mode string
}

// AesOption set the aes options.
type AesOption func(*aesOptions)

func (o *aesOptions) apply(opts ...AesOption) { _ = "STUB: not implemented"; return }

func defaultAesOptions() *aesOptions { _ = "STUB: not implemented"; return nil }

// WithAesKey set aes key
func WithAesKey(key []byte) AesOption { _ = "STUB: not implemented"; return *new(AesOption) }

// WithAesModeCBC set mode to CBC
func WithAesModeCBC() AesOption { _ = "STUB: not implemented"; return *new(AesOption) }

// WithAesModeECB set mode to ECB
func WithAesModeECB() AesOption { _ = "STUB: not implemented"; return *new(AesOption) }

// WithAesModeCFB set mode to CFB
func WithAesModeCFB() AesOption { _ = "STUB: not implemented"; return *new(AesOption) }

// WithAesModeCTR set mode to CTR
func WithAesModeCTR() AesOption { _ = "STUB: not implemented"; return *new(AesOption) }

// ------------------------------------------------------------------------------------------

type desOptions struct {
	desKey []byte // the length of the key must be 8
	mode   string // there are four operating modes in total, ECB CBC CFB CTR
}

// DesOption set the des options.
type DesOption func(*desOptions)

func (o *desOptions) apply(opts ...DesOption) { _ = "STUB: not implemented"; return }

func defaultDesOptions() *desOptions { _ = "STUB: not implemented"; return nil }

// WithDesKey set des key
func WithDesKey(key []byte) DesOption { _ = "STUB: not implemented"; return *new(DesOption) }

// WithDesModeCBC set mode to CBC
func WithDesModeCBC() DesOption { _ = "STUB: not implemented"; return *new(DesOption) }

// WithDesModeECB set mode to ECB
func WithDesModeECB() DesOption { _ = "STUB: not implemented"; return *new(DesOption) }

// WithDesModeCFB set mode to CFB
func WithDesModeCFB() DesOption { _ = "STUB: not implemented"; return *new(DesOption) }

// WithDesModeCTR set mode to CTR
func WithDesModeCTR() DesOption { _ = "STUB: not implemented"; return *new(DesOption) }

// ------------------------------------------------------------------------------------------

type rsaOptions struct {
	// rsa key pair format
	format string
	// hash types for signatures and signature verification
	hashType crypto.Hash
}

// RsaOption set the rsa options.
type RsaOption func(*rsaOptions)

func (o *rsaOptions) apply(opts ...RsaOption) { _ = "STUB: not implemented"; return }

func defaultRsaOptions() *rsaOptions { _ = "STUB: not implemented"; return nil }

// WithRsaFormatPKCS1 set format
func WithRsaFormatPKCS1() RsaOption { _ = "STUB: not implemented"; return *new(RsaOption) }

// WithRsaFormatPKCS8 set format
func WithRsaFormatPKCS8() RsaOption { _ = "STUB: not implemented"; return *new(RsaOption) }

// WithRsaHashTypeMd5 set hash type
func WithRsaHashTypeMd5() RsaOption { _ = "STUB: not implemented"; return *new(RsaOption) }

// WithRsaHashTypeSha1 set hash type
func WithRsaHashTypeSha1() RsaOption { _ = "STUB: not implemented"; return *new(RsaOption) }

// WithRsaHashTypeSha256 set hash type
func WithRsaHashTypeSha256() RsaOption { _ = "STUB: not implemented"; return *new(RsaOption) }

// WithRsaHashTypeSha512 set hash type
func WithRsaHashTypeSha512() RsaOption { _ = "STUB: not implemented"; return *new(RsaOption) }

// WithRsaHashType set hash type
func WithRsaHashType(hash crypto.Hash) RsaOption { _ = "STUB: not implemented"; return *new(RsaOption) }
