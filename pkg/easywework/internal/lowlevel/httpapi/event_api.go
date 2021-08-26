package httpapi

import (
	"io/ioutil"
	"net/http"
	"openscrm/pkg/easywework/internal/lowlevel/envelope"
)

type EnvelopeHandler interface {
	OnIncomingEnvelope(rx envelope.Envelope) error
}

func (h *LowLevelHandler) eventHandler(rw http.ResponseWriter, r *http.Request) {
	// app bodies are assumed small
	// we can't do streaming parse/decrypt/verification anyway
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	// signature verification is inside EnvelopeProcessor
	ev, err := h.ep.HandleIncomingMsg(r.URL, body)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.eh.OnIncomingEnvelope(ev)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	// currently we always return empty 200 responses
	// any reply is to be sent asynchronously
	// this might change in the future (maybe save a couple of RTT or so)
	rw.WriteHeader(http.StatusOK)
}
