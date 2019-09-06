package network

type ProcessorIf interface {
	// must goroutine safe
	Route(interface{}) error
	// must goroutine safe
	Unmarshal(data []byte) (interface{}, error)
	// must goroutine safe
	Marshal(id uint32, msg interface{}) ([][]byte, error)
}
