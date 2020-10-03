package constdef

const IPFSHost = "localhost:5001"
const IPFSMyId = "Qma6wFrmDQ48u7GrZvvzrj5RkMAFBCNepNx62fVe5134GL"

type UserStatus int32

const (
	EdgeStatus    UserStatus = 1
	GatewayStatus UserStatus = 2
	CloudStatus   UserStatus = 3
)

const SaveListLesslength = 2

const MyCloudNode = "/ip4/127.0.0.1/tcp/4001/ipfs/Qma6wFrmDQ48u7GrZvvzrj5RkMAFBCNepNx62fVe5134GL"
