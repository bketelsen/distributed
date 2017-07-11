package contextrpc

type TraceKey int32

const TraceID TraceKey = 0

type HelloArgs struct {
	Who string
}

type HelloReply struct {
	Message string
}
