package protobuf

import (
	"github.com/davidkhala/goutils"
	"google.golang.org/protobuf/proto"
)

// MarshalOrPanic serializes a protobuf message and panics if this operation fails
func MarshalOrPanic(pb proto.Message) []byte {
	data, err := proto.Marshal(pb)
	goutils.PanicError(err)
	return data
}
