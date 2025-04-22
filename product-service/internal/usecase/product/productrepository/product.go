package productuserepository

import (
	"context"
	productentity "productservice/internal/entity/product"
)

type productUseRepository interface {
	AddProduct(ctx context.Context, product *productentity.ProductMain) (err error)
	AddModelToProduct(ctx context.Context, product *productentity.Product) (err error)
	GetAllProduct(ctx context.Context, req *productentity.GetProductsReq) (product []*productentity.Product,err error)
	UpdateProduct(ctx context.Context, req *productentity.UpdateProductReq) (product *productentity.Product, err error)
	DeleteProduct(ctx context.Context, productId string, hardDelete bool) (err error)
	CreateCategory(ctx context.Context, nameCategory string) (err error)
	GetAllCategory(ctx context.Context) (categories []*productentity.CreateCategoryReq,err error)
	UpdateCategory(ctx context.Context, category, newCategory string) (err error)
	DeleteCategory(ctx context.Context, nameCategory string) (err error)
	Update(ctx context.Context, name, newName string) (err error)
	GetMainProduct(ctx context.Context, field, value string) ([]*productentity.ProductMain, error)
}

type ProductUseRepositoryIml struct {
	product productUseRepository
}

func NewProductUsageRepository(product productUseRepository) *ProductUseRepositoryIml {
	return &ProductUseRepositoryIml{
		product: product,
	}
}

func (p *ProductUseRepositoryIml) AddProduct(ctx context.Context, product *productentity.ProductMain) (err error) {
	return p.product.AddProduct(ctx, product)
}

func (p *ProductUseRepositoryIml) AddModelToProduct(ctx context.Context, product *productentity.Product) (err error) {
	return p.product.AddModelToProduct(ctx, product)
}

func (p *ProductUseRepositoryIml) GetAllProduct(ctx context.Context, req *productentity.GetProductsReq) (product []*productentity.Product,err error) {
	return p.product.GetAllProduct(ctx, req)
}

func (p *ProductUseRepositoryIml) UpdateProduct(ctx context.Context, req *productentity.UpdateProductReq) (product *productentity.Product, err error) {
	return p.product.UpdateProduct(ctx, req)
}

func (p *ProductUseRepositoryIml) DeleteProduct(ctx context.Context, productId string, hardDelete bool) (err error) {
	return p.product.DeleteProduct(ctx, productId, hardDelete)
}

func (p *ProductUseRepositoryIml) CreateCategory(ctx context.Context, nameCategory string) (err error) {
	return p.product.CreateCategory(ctx, nameCategory)
}

func (p *ProductUseRepositoryIml) GetAllCategory(ctx context.Context) (categories []*productentity.CreateCategoryReq,err error) {
	return p.product.GetAllCategory(ctx)
}

func (p *ProductUseRepositoryIml) UpdateCategory(ctx context.Context, category, newCategory string) (err error) {
	return p.product.UpdateCategory(ctx, category, newCategory)
}

func (p *ProductUseRepositoryIml) DeleteCategory(ctx context.Context, nameCategory string) (err error) {
	return p.product.DeleteCategory(ctx, nameCategory)
}

func (p *ProductUseRepositoryIml) Update(ctx context.Context, name, newName string) (err error) {
	return p.product.Update(ctx, name, newName)
}


func (p *ProductUseRepositoryIml)GetMainProduct(ctx context.Context, field, value string) ([]*productentity.ProductMain, error){
	return p.product.GetMainProduct(ctx,field,value)
}