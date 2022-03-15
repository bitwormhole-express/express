package impl

import (
	"context"
	"errors"

	"github.com/bitwomrhole-express/express/community-server/app/data/dao"
	"github.com/bitwomrhole-express/express/community-server/app/data/dxo"
	"github.com/bitwomrhole-express/express/community-server/app/data/entity"
	"github.com/bitwomrhole-express/express/community-server/app/service"
	"github.com/bitwomrhole-express/express/community-server/app/web/dto"
	"github.com/bitwormhole/starter-object-bucket/buckets"
	"github.com/bitwormhole/starter-security/keeper"
	"github.com/bitwormhole/starter/markup"
)

type UploadServiceImpl struct {
	markup.Component `id:"express-UploadService"`

	Subjects      keeper.SubjectManager `inject:"#keeper-subject-manager"`
	BucketDrivers buckets.Manager       `inject:"#buckets.Manager"`
	BucketDAO     dao.BucketDAO         `inject:"#express-data-bucket-dao"`
}

func (inst *UploadServiceImpl) _Impl() service.UploadService {
	return inst
}

func (inst *UploadServiceImpl) BeginUpload(c context.Context, up *dto.Upload) (*dto.Upload, error) {

	sub, err := inst.Subjects.GetSubject(c)
	if err != nil {
		return nil, err
	}

	list, err := inst.listBucketsBySubject(sub)
	if err != nil {
		return nil, err
	}

	selector := up.BucketSelector
	bucket, err := inst.selectBucket(selector, list)
	if err != nil {
		return nil, err
	}

	driver, err := inst.BucketDrivers.FindDriver(bucket.Provider)
	if err != nil {
		return nil, err
	}

	conn, err := driver.GetConnector().Open(bucket)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	up2, err := inst.convertUpload1(up)
	if err != nil {
		return nil, err
	}

	o1 := conn.GetObject(up.ObjectName)
	up3, err := o1.UploadByAPI(up2)
	if err != nil {
		return nil, err
	}

	return inst.convertUpload2(up3)
}

func (inst *UploadServiceImpl) EndUpload(c context.Context, id string, up *dto.Upload) (*dto.Upload, error) {

	return nil, errors.New("no impl")
}

func (inst *UploadServiceImpl) listBucketsBySubject(subject keeper.Subject) ([]*buckets.Bucket, error) {

	session, err := subject.GetSession(true)
	if err != nil {
		return nil, err
	}

	ident := session.GetIdentity()
	uid := ident.UserID() // int64
	accountID := dxo.AccountID(uid)

	src, err := inst.BucketDAO.ListBucketsByOwner(accountID)
	if err != nil {
		return nil, err
	}

	dst := make([]*buckets.Bucket, 0)
	for _, item1 := range src {
		item2 := inst.convertBucket(item1)
		dst = append(dst, item2)
	}

	return dst, nil
}

func (inst *UploadServiceImpl) selectBucket(selector string, list []*buckets.Bucket) (*buckets.Bucket, error) {
	// todo ...
	if list == nil {
		return nil, errors.New("no bucket")
	}
	if len(list) < 1 {
		return nil, errors.New("no bucket")
	}
	b := list[0]
	if b == nil {
		return nil, errors.New("no bucket")
	}
	return b, nil
}

func (inst *UploadServiceImpl) convertUpload1(up1 *dto.Upload) (*buckets.HTTPUploading, error) {
	up2 := &buckets.HTTPUploading{}
	up2.Method = up1.Method
	up2.URL = up1.URL
	up2.ContentLength = up1.ContentLength
	up2.ContentMD5 = up1.ContentMD5
	up2.ContentType = up1.ContentType
	up2.UseHTTPS = true
	up2.RequestHeaders = up1.Headers
	return up2, nil
}

func (inst *UploadServiceImpl) convertUpload2(up2 *buckets.HTTPUploading) (*dto.Upload, error) {
	up1 := &dto.Upload{}
	up1.Method = up2.Method
	up1.URL = up2.URL
	up1.ContentLength = up2.ContentLength
	up1.ContentMD5 = up2.ContentMD5
	up1.ContentType = up2.ContentType
	up1.Headers = up2.RequestHeaders
	return up1, nil
}

func (inst *UploadServiceImpl) convertBucket(b1 *entity.Bucket) *buckets.Bucket {
	b2 := &buckets.Bucket{}
	b2.Provider = b1.Provider
	b2.Credential = "todo..."
	return b2
}
