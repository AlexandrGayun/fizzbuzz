package linear

import "fizzbuzztest/processor"

type Core struct {
	Conveyor []processor.Handler
}

func CreateLinearCore() processor.Core {
	return &Core{}
}

func (lc *Core) WriteHandlers(handlers []processor.Handler) error {
	lc.Conveyor = handlers
	return nil
}

func (lc *Core) Process(input any) (any, error) {
	var (
		handlerRes any
		err        error
	)
	for i, handler := range lc.Conveyor {
		if i == 0 {
			handlerRes, err = handler(input)
		} else {
			handlerRes, err = handler(handlerRes)
		}
		if err != nil {
			return nil, err
		}
	}
	return handlerRes, nil
}
