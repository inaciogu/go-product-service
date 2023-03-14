package usecases

import entity "github.com/inaciogu/go-product-service/internal/entities"

type ListAllProductsOutputDto struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type ListAllProductsUseCase struct {
	ProductRepository entity.ProductRepository
}

func NewListAllProductsUseCase(productRepository entity.ProductRepository) *ListAllProductsUseCase {
	return &ListAllProductsUseCase{
		ProductRepository: productRepository,
	}
}

func (u *ListAllProductsUseCase) Execute() ([]*ListAllProductsOutputDto, error) {
	products, err := u.ProductRepository.FindAll()

	if err != nil {
		return nil, err
	}

	var productsOutput []*ListAllProductsOutputDto

	for _, product := range products {
		productsOutput = append(productsOutput, &ListAllProductsOutputDto{
			ID:    product.ID,
			Name:  product.Name,
			Price: product.Price,
		})
	}
	return productsOutput, nil
}
