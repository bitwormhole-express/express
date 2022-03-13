package impl

import (
	"bytes"
	"crypto/rand"
	"crypto/sha1"
	"strconv"
	"sync"

	"github.com/bitwomrhole-express/express/station/app/service"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/util"
)

type UUIDGeneratorImpl struct {
	markup.Component `id:"the-uuid-generator"`

	mutex sync.Mutex
	index int64
}

func (inst *UUIDGeneratorImpl) _Impl() service.UUIDGenerator {
	return inst
}

func (inst *UUIDGeneratorImpl) nextIndex() int64 {
	inst.mutex.Lock()
	defer inst.mutex.Unlock()
	inst.index++
	return inst.index
}

func (inst *UUIDGeneratorImpl) GenUUID(aTypeName, aID string) string {

	nonce := make([]byte, 32)
	now := util.Now()
	index := inst.nextIndex()

	rand.Reader.Read(nonce)

	builder := bytes.Buffer{}
	builder.Write(nonce)
	builder.WriteString(aTypeName)
	builder.WriteString(aID)
	builder.WriteString(now.String())
	builder.WriteString(strconv.FormatInt(index, 10))

	sum := sha1.Sum(builder.Bytes())
	return util.StringifyBytes(sum[:])
}
