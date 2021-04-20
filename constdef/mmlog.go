package constdef

const ParamNil = "nil"

type ErrLogType int32

const (
	Undefine ErrLogType = 0
	Leaf     ErrLogType = 1
	Node     ErrLogType = 2
	Show     ErrLogType = 3
)
