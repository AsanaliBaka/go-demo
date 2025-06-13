package link

import (
	"fmt"
	"go/adv-demo/pkg/db"
)

type LinkRepo struct {
	DataBase *db.Db
}

func NewLinkRepo(db *db.Db) *LinkRepo {
	return &LinkRepo{
		DataBase: db,
	}
}

func (repo *LinkRepo) Create(link *Link) (*Link, error) {
	result := repo.DataBase.DB.Create(link)

	if result.Error != nil {
		return nil, result.Error
	}

	return link, nil
}
func (repo *LinkRepo) Delete(id uint) error {
	result := repo.DataBase.Delete(&Link{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil

}
func (repo *LinkRepo) Update(link *Link) (*Link, error) {

	result := repo.DataBase.Updates(link)

	if result.Error != nil {
		return nil, result.Error
	}

	return link, nil
}
func (repo *LinkRepo) Get(hash string) (*Link, error) {
	var link Link
	result := repo.DataBase.First(&link, "hash = ?", hash)

	if result.Error != nil {
		return nil, result.Error
	}

	return &link, nil
}

func (repo *LinkRepo) GetById(id uint) (*Link, error) {
	var newLink Link

	resultLink := repo.DataBase.First(&newLink, "ID = ?", id)

	if resultLink == nil {
		return nil, fmt.Errorf("not faund")
	}

	return &newLink, nil
}

func (repo *LinkRepo) GetLinks(limit, offset uint) []Link {
	var links []Link
	repo.DataBase.
		Table("links").
		Where("deleted_at is null").
		Order("id asc").
		Limit(int(limit)).
		Offset(int(offset)).Scan(&links)

	return links
}

func (repo *LinkRepo) Counter() int64 {

	var count int64

	repo.DataBase.
		Table("links").
		Where("deleted_at is null").
		Count(&count)

	return count

}
