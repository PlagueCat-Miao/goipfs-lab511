package constdef

const IPFSHost = "localhost:5001"

type UserStatus int32

const (
	EdgeStatus    UserStatus = 1
	GatewayStatus UserStatus = 2
	CloudStatus   UserStatus = 3
)

const SaveListLesslength = 2

const IPFSNodeUrlFormat = "/ip4/%s/tcp/4001/ipfs/%s"

const LocalTestNode = "/ip4/127.0.0.1/tcp/4001/ipfs/Qma6wFrmDQ48u7GrZvvzrj5RkMAFBCNepNx62fVe5134GL"

type OpTypeEvent int32

const (
	Add    OpTypeEvent = 1
	Update OpTypeEvent = 2
)
