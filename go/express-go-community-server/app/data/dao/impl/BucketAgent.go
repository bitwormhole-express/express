package impl

import (
	"github.com/bitwomrhole-express/express/community-server/app/data/dao"
	"github.com/bitwormhole/starter-object-bucket/buckets"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/vlog"
)

///////////////////////////////////////////////////////////////////////////////

// BucketAgentImpl ...
type BucketAgentImpl struct {
	markup.Component `id:"express-data-bucket-agent" initMethod:"Init"`

	BucketDriver string `inject:"${bucket.default.driver}"`

	BM      buckets.Manager     `inject:"#buckets.Manager"`
	Context application.Context `inject:"context"`

	conn buckets.Connection
}

func (inst *BucketAgentImpl) _Impl() dao.BucketAgent {
	return inst
}

func (inst *BucketAgentImpl) Init() error {
	// return inst.loadBucket()
	return nil
}

func (inst *BucketAgentImpl) loadBucket() error {

	driver, err := inst.BM.FindDriver(inst.BucketDriver)
	if err != nil {
		return err
	}

	props := inst.Context.GetProperties()
	cfg, err := driver.GetBucket("bucket", "default", props)
	if err != nil {
		return err
	}

	conn, err := driver.GetConnector().Open(cfg)
	if err != nil {
		return err
	}

	o := conn.GetObject(".test")
	ok, err := o.Exists()
	if err != nil {
		return err
	}
	if ok {
		vlog.Debug("ok")
	}

	inst.conn = conn
	return nil
}

func (inst *BucketAgentImpl) GetBucket() buckets.Connection {
	return inst.conn
}
