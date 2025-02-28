package orm

import (
	"context"
	"fmt"
	"github.com/andriykusevol/aktemplategorm/internal/domain/valueobject/queryfilter"

	"gorm.io/gorm"
)

type PspDto struct {
	ID             *uint          `gorm:"primaryKey;autoIncrement"`
	PspCode        string         `gorm:"size:16"`
	PspShortName   *string        `gorm:"size:50"`
	PspCountryCode *string        `gorm:"size:3"`
	AuditRecord    AuditRecordDto `gorm:"embedded"`
}

func (PspDto) TableName() string {
	return "Psp"
}

func WrapPageQuery(ctx context.Context, db *gorm.DB, pq queryfilter.PaginationQuery, out interface{}) (*queryfilter.Pagination, error) {

	if pq.OnlyCount {
		var count int64
		err := db.Count(&count).Error
		if err != nil {
			return nil, err
		}
		return &queryfilter.Pagination{
			CurrentPage: 0,
			PageSize:    0,
			PagesCount:  0,
			HasMore:     false,
			ItemsCount:  count,
		}, nil
	}

	total, err := FindPage(ctx, db, pq, out)
	_ = total
	if err != nil {
		return nil, err
	}

	hasMore := (pq.CurrentPage * pq.PageSize) < uint(total)

	pagesCount := uint(total) / pq.GetPageSize()
	pagesCount += 1

	return &queryfilter.Pagination{
		CurrentPage: pq.GetCurrent(),
		PageSize:    pq.GetPageSize(),
		PagesCount:  pagesCount,
		HasMore:     hasMore,
		ItemsCount:  total,
	}, nil
}

func FindPage(ctx context.Context, db *gorm.DB, pq queryfilter.PaginationQuery, out interface{}) (int64, error) {
	var count int64
	err := db.Count(&count).Error
	if err != nil {
		return 0, err
	}
	if count == 0 {
		return count, nil
	}

	current, pageSize := int(pq.GetCurrent()), int(pq.GetPageSize())
	if current > 0 && pageSize > 0 {
		db = db.Offset((current - 1) * pageSize).Limit(pageSize)
	}
	if pageSize > 0 {
		db = db.Limit(pageSize)
	}
	err = db.Find(out).Error
	return count, err
}

func Paginate(pp queryfilter.PaginationQuery) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		current, pageSize := int(pp.GetCurrent()), int(pp.GetPageSize())
		fmt.Println(current, pageSize)
		if current > 0 && pageSize > 0 {
			offset := (current - 1) * pageSize
			return db.Offset(offset).Limit(pageSize)
		}
		if pageSize > 0 {
			return db.Limit(pageSize)
		}
		return db
	}
}
