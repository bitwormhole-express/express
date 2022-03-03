package impl

import (
	"errors"

	"bitwomrhole.com/djaf/express-go-server/server/data/dao"
	"bitwomrhole.com/djaf/express-go-server/server/data/entity"
	"bitwomrhole.com/djaf/express-go-server/server/service"
	"bitwomrhole.com/djaf/express-go-server/server/web/dto"
	"github.com/bitwormhole/starter/markup"
)

type PackageServiceImpl struct {
	markup.Component `id:"express-PackageService"`

	PackageDAO dao.PackageDAO `inject:"#express-data-package-dao"`
}

func (inst *PackageServiceImpl) _Impl() service.PackageService {
	return inst
}

func (inst *PackageServiceImpl) Find(o *dto.Package) ([]*dto.Package, error) {

	return nil, errors.New("no impl")
}

func (inst *PackageServiceImpl) FindOne(id dto.PackageID) (*dto.Package, error) {

	return nil, errors.New("no impl")
}

func (inst *PackageServiceImpl) Insert(o1 *dto.Package) (*dto.Package, error) {

	ent, err := inst.dto2entity(o1)
	if err != nil {
		return nil, err
	}

	ent, err = inst.PackageDAO.Insert(ent)
	if err != nil {
		return nil, err
	}

	o2, err := inst.entity2dto(ent)
	if err != nil {
		return nil, err
	}
	return o2, nil
}

func (inst *PackageServiceImpl) Remove(id dto.PackageID) error {

	return errors.New("no impl")
}

func (inst *PackageServiceImpl) Update(id dto.PackageID, o *dto.Package) (*dto.Package, error) {

	return nil, errors.New("no impl")
}

func (inst *PackageServiceImpl) entity2dto(o1 *entity.Package) (*dto.Package, error) {
	o2 := &dto.Package{
		ID:          dto.PackageID(o1.ID),
		FromAddress: o1.FromAddress,
		ToAddress:   o1.ToAddress,
		Comment:     o1.Comment,
		Status:      o1.Status,
		Transport:   o1.Transport,
	}
	o2.UUID = o1.UUID
	return o2, nil
}

func (inst *PackageServiceImpl) dto2entity(o1 *dto.Package) (*entity.Package, error) {
	o2 := &entity.Package{
		ID:          entity.PackageID(o1.ID),
		FromAddress: o1.FromAddress,
		ToAddress:   o1.ToAddress,
		Comment:     o1.Comment,
		Status:      o1.Status,
		Transport:   o1.Transport,
	}
	o2.UUID = o1.UUID
	return o2, nil
}
