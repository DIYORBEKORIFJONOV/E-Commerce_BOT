package productusecase

import (
	"context"
	productentity "productservice/internal/entity/product"
)

type productUseCase interface {
	CreateProduct(ctx context.Context, req *productentity.CreateProductReq) (res *productentity.ProductMain, err error)
	AddModel(ctx context.Context, req *productentity.AddModelReq) (product *productentity.Product, err error)
	GetAllProduct(ctx context.Context, req *productentity.GetProductsReq) (products []*productentity.Product,count int64, err error)
	UpdateProduct(ctx context.Context, req *productentity.UpdateProductReq) (product *productentity.Product, err error)
	DeleteProduct(ctx context.Context, req *productentity.DeleteProductReq) (res *productentity.GeneralResponseProduct, err error)
	CreateCategory(ctx context.Context, req *productentity.CreateCategoryReq) (res *productentity.GeneralResponseProduct, err error)
	GetAllCategory(ctx context.Context, req *productentity.GetCategoriesReq) (categories *productentity.GetcategoriesRes, err error)
	UpdateCategory(ctx context.Context, req *productentity.UpdateCategoryReq) (res *productentity.GeneralResponseProduct, err error)
	DeleteCategory(ctx context.Context, req *productentity.DeleteCategoryReq) (res *productentity.GeneralResponseProduct, err error)
	UpdateName(ctx context.Context, req *productentity.UpdateNameReq) (res *productentity.GeneralResponseProduct, err error)
	GetMainProduct(ctx context.Context,field,value string)(products []*productentity.ProductMain, err error)
}

type ProductUseCaseIml struct {
	product productUseCase
}

func NewProductUsage(product productUseCase) *ProductUseCaseIml {
	return &ProductUseCaseIml{
		product: product,
	}
}

func (p *ProductUseCaseIml) CreateProduct(ctx context.Context, req *productentity.CreateProductReq) (res *productentity.ProductMain, err error) {
	return p.product.CreateProduct(ctx, req)
}

func (p *ProductUseCaseIml) AddModel(ctx context.Context, req *productentity.AddModelReq) (product *productentity.Product, err error) {
	return p.product.AddModel(ctx, req)
}

func (p *ProductUseCaseIml) GetAllProduct(ctx context.Context, req *productentity.GetProductsReq) (products []*productentity.Product,count int64, err error) {
	return p.product.GetAllProduct(ctx, req)
}

func (p *ProductUseCaseIml) UpdateProduct(ctx context.Context, req *productentity.UpdateProductReq) (product *productentity.Product, err error) {
	return p.product.UpdateProduct(ctx, req)
}

func (p *ProductUseCaseIml) DeleteProduct(ctx context.Context, req *productentity.DeleteProductReq) (res *productentity.GeneralResponseProduct, err error) {
	return p.product.DeleteProduct(ctx, req)
}

func (p *ProductUseCaseIml) CreateCategory(ctx context.Context, req *productentity.CreateCategoryReq) (res *productentity.GeneralResponseProduct, err error) {
	return p.product.CreateCategory(ctx, req)
}

func (p *ProductUseCaseIml) GetAllCategory(ctx context.Context, req *productentity.GetCategoriesReq) (categories *productentity.GetcategoriesRes, err error) {
	return p.product.GetAllCategory(ctx, req)
}

func (p *ProductUseCaseIml) UpdateCategory(ctx context.Context, req *productentity.UpdateCategoryReq) (res *productentity.GeneralResponseProduct, err error) {
	return p.product.UpdateCategory(ctx, req)
}

func (p *ProductUseCaseIml) DeleteCategory(ctx context.Context, req *productentity.DeleteCategoryReq) (res *productentity.GeneralResponseProduct, err error) {
	return p.product.DeleteCategory(ctx, req)
}

func (p *ProductUseCaseIml) UpdateName(ctx context.Context, req *productentity.UpdateNameReq) (res *productentity.GeneralResponseProduct, err error) {
	return p.product.UpdateName(ctx, req)
}

func (p *ProductUseCaseIml)GetMainProduct(ctx context.Context,field,value string)(products []*productentity.ProductMain, err error) {
	return p.product.GetMainProduct(ctx,field,value)
}