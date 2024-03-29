package service

import (
	"Fp-TokoBelanja/model/entity"
	"Fp-TokoBelanja/model/input"
	"Fp-TokoBelanja/repository"
	"errors"
)

type ProductService interface {
	CreateProduct(role_user string, inputBody input.ProductCreateInput) (entity.Product, error)
	GetAllProducts() ([]entity.Product, error)
	UpdateProduct(role_user string, id_product int, input input.ProductUpdateInput) (entity.Product, error)
	DeleteProduct(role_user string, id_product int) error
}

type productService struct {
	productRepository  repository.ProductRepository
	categoryRepository repository.CategoryRepository
}

func NewProductService(productRepository repository.ProductRepository, categoryRepository repository.CategoryRepository) *productService {
	return &productService{productRepository, categoryRepository}
}

func (s *productService) CreateProduct(role_user string, inputBody input.ProductCreateInput) (entity.Product, error) {
	if role_user != "admin" {
		return entity.Product{}, errors.New("you are not admin")
	}

	categoryData, err := s.categoryRepository.FindById(inputBody.CategoryID)
	if err != nil {
		return entity.Product{}, err
	}
	if categoryData.ID == 0 {
		return entity.Product{}, errors.New("category not found")
	}

	product := entity.Product{
		Title:      inputBody.Title,
		Price:      inputBody.Price,
		Stock:      inputBody.Stock,
		CategoryID: inputBody.CategoryID,
	}

	return s.productRepository.Save(product)
}

func (s *productService) GetAllProducts() ([]entity.Product, error) {
	return s.productRepository.FindAll()
}

func (s *productService) UpdateProduct(role_user string, id_product int, input input.ProductUpdateInput) (entity.Product, error) {
	if role_user != "admin" {
		return entity.Product{}, errors.New("you are not admin")
	}

	productData, err := s.productRepository.FindById(id_product)
	if err != nil {
		return entity.Product{}, err
	}
	if productData.ID == 0 {
		return entity.Product{}, errors.New("product not found")
	}

	categoryData, err := s.categoryRepository.FindById(input.CategoryID)
	if err != nil {
		return entity.Product{}, err
	}
	if categoryData.ID == 0 {
		return entity.Product{}, errors.New("category not found")
	}

	product := entity.Product{
		Title:      input.Title,
		Price:      input.Price,
		Stock:      input.Stock,
		CategoryID: input.CategoryID,
	}

	_, err = s.productRepository.Update(id_product, product)
	if err != nil {
		return entity.Product{}, err
	}

	return s.productRepository.FindById(id_product)
}

func (s *productService) DeleteProduct(role_user string, id_product int) error {
	if role_user != "admin" {
		return errors.New("you are not admin")
	}

	productData, err := s.productRepository.FindById(id_product)
	if err != nil {
		return err
	}
	if productData.ID == 0 {
		return errors.New("product not found")
	}

	return s.productRepository.Delete(id_product)
}
