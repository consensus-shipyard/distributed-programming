package perfectlink

type Impl struct {
	Output Output
}

func New() *Impl {
	return &Impl{}
}

func (pl *Impl) Send(message []byte) {
	pl.Output.Receive(message)
}
