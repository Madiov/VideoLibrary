package repository

import (
	"fmt"
	"gorm.io/gorm"
)

type BaseRepository struct {
	db *gorm.DB
}

func NewBaseRepository(db *gorm.DB) *BaseRepository {
	return &BaseRepository{db: db}
}

func (b *BaseRepository) ReadAll(Response interface{}) error {
	db := b.db.Find(&Response)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (b *BaseRepository) ReadAllBy(Response interface{}, column string, value interface{}) error {
	q := fmt.Sprintf("%s = ? ", column)
	db := b.db.Find(&Response).Where(q, value)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (b *BaseRepository) ReadBy(model interface{}, column string, value interface{}) error {
	q := fmt.Sprintf("%s = ? ", column)
	db := b.db.First(&model).Where(q, value)
	if db.Error == nil {
		return db.Error
	}
	return nil
}

func (b *BaseRepository) Create(model interface{}) error {
	db := b.db.Create(&model)
	if db.RowsAffected == 0 {
		return db.Error
	}
	return nil
}

func (b *BaseRepository) UpdateBy(model interface{}, UpdateObject interface{}, column string, value interface{}) error {
	q := fmt.Sprintf("%s = ? ", column)
	db := b.db.Model(&model).Where(q, value).Updates(&UpdateObject)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (b *BaseRepository) DeleteBy(model interface{}, column string, value interface{}) error {
	q := fmt.Sprintf("%s = ? ", column)
	db := b.db.Delete(&model).Where(q, value)
	if db.Error != nil {
		return db.Error
	}
	return nil
}
