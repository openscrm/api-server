package httpapi

import (
	"errors"
	"net/http"
	"net/url"
	"openscrm/pkg/easywework/internal/lowlevel/signature"
	"strconv"
)

type ToEchoTestAPIArgs interface {
	ToEchoTestAPIArgs() (EchoTestAPIArgs, error)
}

type EchoTestAPIArgs struct {
	MsgSignature string
	Timestamp    int64
	Nonce        string
	EchoStr      string
}

type URLValuesForEchoTestAPI url.Values

var _ ToEchoTestAPIArgs = URLValuesForEchoTestAPI{}

var errMalformedArgs = errors.New("malformed arguments for echo test API")

func (x URLValuesForEchoTestAPI) ToEchoTestAPIArgs() (EchoTestAPIArgs, error) {
	var msgSignature string
	{
		l := x["msg_signature"]
		if len(l) != 1 {
			return EchoTestAPIArgs{}, errMalformedArgs
		}
		msgSignature = l[0]
	}

	var timestamp int64
	{
		l := x["timestamp"]
		if len(l) != 1 {
			return EchoTestAPIArgs{}, errMalformedArgs
		}
		timestampStr := l[0]

		timestampInt, err := strconv.ParseInt(timestampStr, 10, 64)
		if err != nil {
			return EchoTestAPIArgs{}, errMalformedArgs
		}

		timestamp = timestampInt
	}

	var nonce string
	{
		l := x["nonce"]
		if len(l) != 1 {
			return EchoTestAPIArgs{}, errMalformedArgs
		}
		nonce = l[0]
	}

	var echoStr string
	{
		l := x["echostr"]
		if len(l) != 1 {
			return EchoTestAPIArgs{}, errMalformedArgs
		}
		echoStr = l[0]
	}

	return EchoTestAPIArgs{
		MsgSignature: msgSignature,
		Timestamp:    timestamp,
		Nonce:        nonce,
		EchoStr:      echoStr,
	}, nil
}

func (h *LowLevelHandler) echoTestHandler(rw http.ResponseWriter, r *http.Request) {
	url := r.URL

	if !signature.VerifyHTTPRequestSignature(h.token, url, "") {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	adapter := URLValuesForEchoTestAPI(url.Query())
	args, err := adapter.ToEchoTestAPIArgs()
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	payload, err := h.encryptor.Decrypt([]byte(args.EchoStr))
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusOK)
	// No way to signal failure with the typical HTTP handler method signature
	_, _ = rw.Write(payload.Msg)
}

func (h *LowLevelHandler) Decrypt(echoStr string) ([]byte, error) {
	payload, err := h.encryptor.Decrypt([]byte(echoStr))
	if err != nil {
		return nil, err
	}
	return payload.Msg, err
}
