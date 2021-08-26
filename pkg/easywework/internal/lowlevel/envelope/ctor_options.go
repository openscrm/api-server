package envelope

import (
	"io"
)

type ProcessorOption interface {
	applyTo(x *Processor)
}

type customEntropySource struct {
	inner io.Reader
}

func WithEntropySource(e io.Reader) ProcessorOption {
	return &customEntropySource{inner: e}
}

func (o *customEntropySource) applyTo(x *Processor) {
	x.entropySource = o.inner
}

type customTimeSource struct {
	inner TimeSource
}

func WithTimeSource(t TimeSource) ProcessorOption {
	return &customTimeSource{inner: t}
}

func (o *customTimeSource) applyTo(x *Processor) {
	x.timeSource = o.inner
}
