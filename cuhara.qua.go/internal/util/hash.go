package util

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"math"
	"strings"

	"golang.org/x/crypto/argon2"
)

const (
	Argon2HashID = "argon2id"
)

var (
	ErrInvalidArgon2Hash         = errors.New("invalid argon2 hash")
	ErrIncompatibleArgon2Version = errors.New("incompatible argon2 version")
)

var (
	DefaultArgon2Params = &Argon2Params{
		Time:       1,
		Memory:     64 * 1024,
		Threads:    4,
		KeyLength:  32,
		SaltLength: 16,
	}
)

type Argon2Params struct {
	Time       uint32
	Memory     uint32
	Threads    uint8
	KeyLength  uint32
	SaltLength uint32
}

func DefaultArgon2ParamsFromEnv() *Argon2Params {
	return &Argon2Params{
		Time:       GetEnvAsUint32("AUTH_HASHING_ARGON2_TIME", DefaultArgon2Params.Time),
		Memory:     GetEnvAsUint32("AUTH_HASHING_ARGON2_MEMORY", DefaultArgon2Params.Memory),
		Threads:    GetEnvAsUint8("AUTH_HASHING_ARGON2_THREADS", DefaultArgon2Params.Threads),
		KeyLength:  GetEnvAsUint32("AUTH_HASHING_ARGON2_KEY_LENGTH", DefaultArgon2Params.KeyLength),
		SaltLength: GetEnvAsUint32("AUTH_HASHING_ARGON2_SALT_LENGTH", DefaultArgon2Params.SaltLength),
	}
}

func HashPassword(password string, params *Argon2Params) (hash string, err error) {
	salt, err := generateSalt(params.SaltLength)
	if err != nil {
		return "", err
	}

	key := argon2.IDKey([]byte(password), salt, params.Time, params.Memory, params.Threads, params.KeyLength)

	b64salt := base64.RawStdEncoding.EncodeToString(salt)
	b64key := base64.RawStdEncoding.EncodeToString(key)

	return fmt.Sprintf("$%s$v=%d$m=%d,t=%d,p=%d$%s$%s", Argon2HashID, argon2.Version, params.Memory, params.Time, params.Threads, b64salt, b64key), nil
}

func ComparePasswordAndHash(password string, hash string) (matches bool, err error) {
	params, salt, key, err := decodeArgon2Hash(hash)
	if err != nil {
		return false, err
	}

	pKey := argon2.IDKey([]byte(password), salt, params.Time, params.Memory, params.Threads, params.KeyLength)

	keyLen := len(key)
	pKeyLen := len(pKey)

	if keyLen > math.MaxInt32 || pKeyLen > math.MaxInt32 {
		return false, ErrInvalidArgon2Hash
	}

	if subtle.ConstantTimeEq(int32(keyLen), int32(pKeyLen)) == 0 {
		return false, nil
	}

	if subtle.ConstantTimeCompare(key, pKey) == 0 {
		return false, nil
	}

	return true, nil
}

func decodeArgon2Hash(hash string) (params *Argon2Params, salt []byte, key []byte, err error) {
	vals := strings.Split(hash, "$")
	if len(vals) != 6 {
		return nil, nil, nil, ErrInvalidArgon2Hash
	}
	if vals[1] != Argon2HashID {
		return nil, nil, nil, ErrInvalidArgon2Hash
	}

	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, ErrIncompatibleArgon2Version
	}
	if version != argon2.Version {
		return nil, nil, nil, ErrIncompatibleArgon2Version
	}

	params = &Argon2Params{}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &params.Memory, &params.Time, &params.Threads)
	if err != nil {
		return nil, nil, nil, ErrInvalidArgon2Hash
	}

	salt, err = base64.RawStdEncoding.DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, ErrInvalidArgon2Hash
	}

	key, err = base64.RawStdEncoding.DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, ErrInvalidArgon2Hash
	}

	keyLength := len(key)
	if keyLength > math.MaxInt32 {
		return nil, nil, nil, ErrInvalidArgon2Hash
	}

	params.KeyLength = uint32(keyLength)

	return params, salt, key, nil
}

func generateSalt(n uint32) ([]byte, error) {
	b := make([]byte, n)

	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}
