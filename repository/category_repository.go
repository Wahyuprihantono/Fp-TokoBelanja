package repository

import (
	"Fp-TokoBelanja/model/entity"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Save(category entity.Category) (entity.Category, error)
	FindAll() ([]entity.Category, error)
	FindById(id_category int) (entity.Category, error)
	Update(id_category int, category entity.Category) (entity.Category, error)
	Delete(id_category int) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) Save(category entity.Category) (entity.Category, error) {
	err := r.db.Create(&category).Error
	return category, err
}

func (r *categoryRepository) FindAll() ([]entity.Category, error) {
	var categories []entity.Category
	err := r.db.Find(&categories).Error
	return categories, err
}

func (r *categoryRepository) FindById(id_category int) (entity.Category, error) {
	var category entity.Category
	err := r.db.Where("id = ?", id_category).Find(&category).Error
	return category, err
}

func (r *categoryRepository) Update(id_category int, category entity.Category) (entity.Category, error) {
	err := r.db.Where("id = ?", id_category).Updates(&category).Error
	return category, err
}

func (r *categoryRepository) Delete(id_category int) error {
	err := r.db.Where("id = ?", id_category).Delete(&entity.Category{}).Error
	return err
}
