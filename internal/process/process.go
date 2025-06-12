package process

type Processor struct{}

func NewProcessor(storageDir string) *Processor {
	return &Processor{}
}
