package stat

import (
	"go/adv-demo/pkg/db"
	"time"

	"gorm.io/datatypes"
)

type StatRepo struct {
	*db.Db
}

func NewStatRepo(db *db.Db) *StatRepo {
	return &StatRepo{
		Db: db,
	}
}

func (repo *StatRepo) AddClic(linkId uint) {
	var stat Stat
	currentDate := datatypes.Date(time.Now())
	repo.DB.Find(&stat, "link_id = ? and date = ?", linkId, currentDate)

	if stat.ID == 0 {
		repo.DB.Create(&Stat{
			LinkId: linkId,
			Clicks: 1,
			Date:   currentDate,
		})
	} else {
		stat.Clicks += 1
		repo.DB.Save(&stat)
	}

}

func (repo *StatRepo) GetStat(by string, from, to time.Time) []GetStatResponse {
	var stats []GetStatResponse
	var selectQuert string

	switch by {
	case FilterByMonth:
		selectQuert = "to_char(date , 'YYYY-MM') as period, sum(clicks)"

	case FilterByDay:
		selectQuert = "to_char(date , 'YYYY-MM') as period, sum(clicks)"
	}

	repo.DB.Table("stats").
		Select(selectQuert).
		Where(
			"date BETWEEN ? AND ?", from, to).
		Group("period").
		Order("period").Scan(&stats)

	return stats
}
