package httpapi

import (
	"github.com/pkg/errors"
	"net/http"
	"net/url"
	"openscrm/pkg/easywework/internal/lowlevel/encryptor"
	"openscrm/pkg/easywework/internal/lowlevel/envelope"
	"openscrm/pkg/easywework/internal/lowlevel/signature"
)

type LowLevelHandler struct {
	token     string
	encryptor *encryptor.WorkWXEncryptor
	ep        *envelope.Processor
	eh        EnvelopeHandler
}

var _ http.Handler = (*LowLevelHandler)(nil)

func (h *LowLevelHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// 测试回调模式请求
		h.echoTestHandler(rw, r)

	case http.MethodPost:
		// 回调事件
		h.eventHandler(rw, r)

	default:
		// unhandled app method
		rw.WriteHeader(http.StatusNotImplemented)
	}
}

func (h *LowLevelHandler) EchoTest(url *url.URL) (echoMsg []byte, err error) {
	if !signature.VerifyHTTPRequestSignature(h.token, url, "") {
		return nil, errors.New("verify signature failed")
	}

	adapter := URLValuesForEchoTestAPI(url.Query())
	args, err := adapter.ToEchoTestAPIArgs()
	if err != nil {
		return nil, err
	}

	payload, err := h.encryptor.Decrypt([]byte(args.EchoStr))
	if err != nil {
		return nil, err
	}
	return payload.Msg, nil
}
