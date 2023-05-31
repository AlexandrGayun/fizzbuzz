package processor

type Core interface {
	WriteHandlers(handlers []Handler) error
	Process(input any) (any, error)
}

type Handler func(args ...any) (any, error)

type Processor struct {
	Core Core
}

func CreateNewProcessor(core Core, handlers []Handler) (Processor, error) {
	p := Processor{Core: core}
	err := p.Core.WriteHandlers(handlers)
	if err != nil {
		return p, err
	}
	return p, nil
}

func (p Processor) Run(input any) (any, error) {
	res, err := p.Core.Process(input)
	if err != nil {
		return nil, err
	}
	return res, nil
}
