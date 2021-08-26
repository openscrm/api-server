package envelope

import (
	"crypto/rand"
	"encoding/xml"
	"errors"
	"io"
	"math/big"
	"net/url"
	"openscrm/pkg/easywework/internal/lowlevel/encryptor"
	"openscrm/pkg/easywework/internal/lowlevel/signature"
	"strconv"
)

type Processor struct {
	token         string
	encryptor     *encryptor.WorkWXEncryptor
	entropySource io.Reader
	timeSource    TimeSource
}

func NewProcessor(token string, encodingAESKey string, opts ...ProcessorOption) (*Processor, error) {
	obj := Processor{
		token:         token,
		encryptor:     nil, // XXX init later
		entropySource: rand.Reader,
		timeSource:    DefaultTimeSource{},
	}
	for _, o := range opts {
		o.applyTo(&obj)
	}

	enc, err := encryptor.NewWorkWXEncryptor(
		encodingAESKey,
		encryptor.WithEntropySource(obj.entropySource),
	)
	if err != nil {
		return nil, err
	}
	obj.encryptor = enc

	return &obj, nil
}

var errInvalidSignature = errors.New("invalid signature")

func (p *Processor) HandleIncomingMsg(url *url.URL, body []byte) (Envelope, error) {
	// xml unmarshal
	var x xmlRxEnvelope
	err := xml.Unmarshal(body, &x)
	if err != nil {
		return Envelope{}, err
	}

	// check signature
	if !signature.VerifyHTTPRequestSignature(p.token, url, x.Encrypt) {
		return Envelope{}, errInvalidSignature
	}

	// decrypt message
	msg, err := p.encryptor.Decrypt([]byte(x.Encrypt))
	if err != nil {
		return Envelope{}, err
	}

	// assemble envelope to return
	return Envelope{
		ToUserName: x.ToUserName,
		AgentID:    x.AgentID,
		Msg:        msg.Msg,
		ReceiveID:  msg.ReceiveID,
	}, nil
}

func (p *Processor) MakeOutgoingEnvelope(msg []byte) ([]byte, error) {
	workwxPayload := encryptor.WorkWXPayload{
		Msg:       msg,
		ReceiveID: nil,
	}
	encryptedMsg, err := p.encryptor.Encrypt(&workwxPayload)
	if err != nil {
		return nil, err
	}

	ts := p.timeSource.GetCurrentTimestamp().Unix()
	nonce, err := makeNonce(p.entropySource)
	if err != nil {
		return nil, err
	}

	msgSignature := signature.MakeDevMsgSignature(
		p.token,
		strconv.FormatInt(ts, 10),
		nonce,
		encryptedMsg,
	)

	envelope := xmlTxEnvelope{
		XMLName: xml.Name{},
		Encrypt: cdataNode{
			CData: encryptedMsg,
		},
		MsgSignature: cdataNode{
			CData: msgSignature,
		},
		Timestamp: ts,
		Nonce: cdataNode{
			CData: nonce,
		},
	}

	result, err := xml.Marshal(envelope)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func makeNonce(entropySource io.Reader) (string, error) {
	limit := big.NewInt(1)
	limit = limit.Lsh(limit, 64)
	n, err := rand.Int(entropySource, limit)
	if err != nil {
		return "", err
	}
	return n.String(), nil
}
