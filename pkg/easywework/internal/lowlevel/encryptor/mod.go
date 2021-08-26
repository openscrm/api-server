package encryptor

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"io"
	"openscrm/pkg/easywework/internal/lowlevel/pkcs7"
)

type WorkWXPayload struct {
	Msg       []byte
	ReceiveID []byte
}

type WorkWXEncryptor struct {
	aesKey        []byte
	entropySource io.Reader
}

type WorkWXEncryptorOption interface {
	applyTo(x *WorkWXEncryptor)
}

type customEntropySource struct {
	inner io.Reader
}

func WithEntropySource(e io.Reader) WorkWXEncryptorOption {
	return &customEntropySource{inner: e}
}

func (o *customEntropySource) applyTo(x *WorkWXEncryptor) {
	x.entropySource = o.inner
}

var errMalformedEncodingAESKey = errors.New("malformed EncodingAESKey")

func NewWorkWXEncryptor(
	encodingAESKey string,
	opts ...WorkWXEncryptorOption,
) (*WorkWXEncryptor, error) {
	aesKey, err := base64.StdEncoding.DecodeString(encodingAESKey + "=")
	if err != nil {
		return nil, err
	}

	if len(aesKey) != 32 {
		return nil, errMalformedEncodingAESKey
	}

	obj := WorkWXEncryptor{
		aesKey:        aesKey,
		entropySource: rand.Reader,
	}
	for _, o := range opts {
		o.applyTo(&obj)
	}

	return &obj, nil
}

func (e *WorkWXEncryptor) Decrypt(base64Msg []byte) (WorkWXPayload, error) {
	// base64 decode
	bufLen := base64.StdEncoding.DecodedLen(len(base64Msg))
	buf := make([]byte, bufLen)
	n, err := base64.StdEncoding.Decode(buf, base64Msg)
	if err != nil {
		return WorkWXPayload{}, err
	}
	buf = buf[:n]

	// init cipher
	block, err := aes.NewCipher(e.aesKey)
	if err != nil {
		return WorkWXPayload{}, err
	}

	iv := e.aesKey[:16]
	state := cipher.NewCBCDecrypter(block, iv)

	// decrypt in-place in the allocated temp buffer
	state.CryptBlocks(buf, buf)
	buf = pkcs7.Unpad(buf)

	// assemble decrypted payload
	// drop the 16-byte random prefix
	msgLen := binary.BigEndian.Uint32(buf[16:20])
	msg := buf[20 : 20+msgLen]
	receiveID := buf[20+msgLen:]

	return WorkWXPayload{
		Msg:       msg,
		ReceiveID: receiveID,
	}, nil
}

func (e *WorkWXEncryptor) prepareBufForEncryption(payload *WorkWXPayload) ([]byte, error) {
	resultMsgLen := 16 + 4 + len(payload.Msg) + len(payload.ReceiveID)

	// allocate buffer
	buf := make([]byte, 16, resultMsgLen)

	// add random prefix
	_, err := io.ReadFull(e.entropySource, buf) // len(buf) == 16 at this moment
	if err != nil {
		return nil, err
	}

	buf = buf[:cap(buf)] // grow to full capacity
	binary.BigEndian.PutUint32(buf[16:], uint32(len(payload.Msg)))
	copy(buf[20:], payload.Msg)
	copy(buf[20+len(payload.Msg):], payload.ReceiveID)

	return pkcs7.Pad(buf), nil
}

func (e *WorkWXEncryptor) Encrypt(payload *WorkWXPayload) (string, error) {
	buf, err := e.prepareBufForEncryption(payload)
	if err != nil {
		return "", err
	}

	// init cipher
	block, err := aes.NewCipher(e.aesKey)
	if err != nil {
		return "", err
	}

	iv := e.aesKey[:16]
	state := cipher.NewCBCEncrypter(block, iv)

	// encrypt in-place as we own the buffer
	state.CryptBlocks(buf, buf)

	return base64.StdEncoding.EncodeToString(buf), nil
}
