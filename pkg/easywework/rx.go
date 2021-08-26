package workwx

import (
	"net/http"
	"openscrm/pkg/easywework/internal/lowlevel/envelope"
	"openscrm/pkg/easywework/internal/lowlevel/httpapi"
)

// RxMessageHandler 用来接收消息的接口。
type RxMessageHandler interface {
	// OnIncomingMessage 一条消息到来时的回调。
	OnIncomingMessage(msg *RxMessage) error
}

type LowLevelEnvelopeHandler struct {
	highLevelHandler RxMessageHandler
}

var _ httpapi.EnvelopeHandler = (*LowLevelEnvelopeHandler)(nil)

func (h *LowLevelEnvelopeHandler) OnIncomingEnvelope(rx envelope.Envelope) error {
	msg, err := fromEnvelope(rx.Msg)
	if err != nil {
		return err
	}

	return h.highLevelHandler.OnIncomingMessage(msg)
}

type HTTPHandler struct {
	inner *httpapi.LowLevelHandler
}

var _ http.Handler = (*HTTPHandler)(nil)

func (h *HTTPHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.inner.ServeHTTP(rw, r)
}
