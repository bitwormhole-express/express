package impl

import (
	"errors"

	"github.com/bitwomrhole-express/express/station/app/data/dao"
	"github.com/bitwomrhole-express/express/station/app/data/dxo"
	"github.com/bitwomrhole-express/express/station/app/data/entity"
	"github.com/bitwomrhole-express/express/station/app/service"
	"github.com/bitwomrhole-express/express/station/app/web/dto"
	"github.com/bitwormhole/starter/markup"
)

// PackageServiceImpl ...
type PackageServiceImpl struct {
	markup.Component `id:"express-PackageService"`

	PackageDAO dao.PackageDAO `inject:"#express-data-package-dao"`
}

func (inst *PackageServiceImpl) _Impl() service.PackageService {
	return inst
}

// Find ...
func (inst *PackageServiceImpl) Find(o *dto.Package) ([]*dto.Package, error) {

	return nil, errors.New("no impl")
}

// FindOne ...
func (inst *PackageServiceImpl) FindOne(id dxo.PackageID) (*dto.Package, error) {

	return nil, errors.New("no impl")
}

// Insert ...
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

// Remove ...
func (inst *PackageServiceImpl) Remove(id dxo.PackageID) error {

	return errors.New("no impl")
}

// Update ...
func (inst *PackageServiceImpl) Update(id dxo.PackageID, o *dto.Package) (*dto.Package, error) {

	return nil, errors.New("no impl")
}

func (inst *PackageServiceImpl) entity2dto(o1 *entity.Package) (*dto.Package, error) {
	o2 := &dto.Package{
		ID:          o1.ID,
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
		ID:          o1.ID,
		FromAddress: o1.FromAddress,
		ToAddress:   o1.ToAddress,
		Comment:     o1.Comment,
		Status:      o1.Status,
		Transport:   o1.Transport,
	}
	o2.UUID = o1.UUID
	return o2, nil
}
