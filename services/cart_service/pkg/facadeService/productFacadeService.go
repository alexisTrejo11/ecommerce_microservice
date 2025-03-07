package facadeService

import "github.com/google/uuid"

type Product struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Price       float64   `json:"price"`
	IsAvalaible bool      `json:"isAvalaible"`
	Disccount   float64   `json:"disccount"`
}

type ProductFacadeService interface {
	GetProductById(id uuid.UUID) (*Product, error)
	GetProductsByIdIn(id []uuid.UUID) (*[]Product, error)
}

type ProductFacadeServiceImpl struct {
}

func NewProductFacadeService() ProductFacadeService {
	return &ProductFacadeServiceImpl{}
}

// Return Generic until Prodcut Serive is implemented
func (p *ProductFacadeServiceImpl) GetProductById(id uuid.UUID) (*Product, error) {
	return p.getDummyProduct(), nil
}

func (p *ProductFacadeServiceImpl) GetProductsByIdIn(id []uuid.UUID) (*[]Product, error) {
	prodcuts := make([]Product, len(id))
	for range id {
		prodcuts = append(prodcuts, *p.getDummyProduct())
	}
	return &prodcuts, nil
}

func (p *ProductFacadeServiceImpl) getDummyProduct() *Product {
	return &Product{
		Id:          uuid.New(),
		Name:        "Product Name",
		Price:       100.00,
		IsAvalaible: true,
		Disccount:   10.00,
	}
}
