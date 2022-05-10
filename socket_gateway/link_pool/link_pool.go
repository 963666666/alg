package link_pool

type LinkPool struct {
}

type LinkOpts struct {
	ServerPort       int32
	ListCount        int
	ListSize         int
	ProtoLenByteSize int
}

var instance *LinkPool

func init() {
	instance = new(LinkPool)
}

func GetInstance() *LinkPool {
	return instance
}



