package enum

type ResStatus int

const (
	ServerError ResStatus =500
	UnAuthorized ResStatus=403
	NotFound ResStatus=404
	Success ResStatus =200
	Accepted ResStatus=201
	DeleteSuccess ResStatus=204
)

func (receiver ResStatus) String() string {
	switch receiver {
	case ServerError:
		return "server error"
	case UnAuthorized:
		return "unauthorized access"
	case NotFound:
		return "resource not found"
	case Success:
		return "get resource successful"
	case Accepted:
		return "resource action accepted"
	case DeleteSuccess:
		return "remove resource successful"
	default:
		return "unknown error"
	}
}
