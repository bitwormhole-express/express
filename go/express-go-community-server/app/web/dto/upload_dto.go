package dto

import "github.com/bitwormhole/starter/util"

type Upload struct {
	Base

	ID string

	Method         string
	URL            string
	ObjectName     string // aka. Path
	BucketSelector string
	ContentType    string
	ContentMD5     util.Hex
	ContentLength  int64
	Headers        map[string]string
}
