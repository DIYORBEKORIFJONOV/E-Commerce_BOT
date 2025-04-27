package product_service

import (
	productentity "api_gateway/internal/entity/product"
	clientgrpcserver "api_gateway/internal/infastructure/client_grpc_server"
	adjustrequestproduct "api_gateway/internal/service/product/adjustRequest"
	adjustresponseproduct "api_gateway/internal/service/product/adjustResponse"
	"context"
)

type ProductService struct {
	O   clientgrpcserver.ServiceClient
	Req *adjustrequestproduct.AdjustRequest
	Res *adjustresponseproduct.AdjustResponse
}

// type productUseCase interface {
// 	CreateProduct(ctx context.Context, req *productentity.CreateProductReq) (res *productentity.ProductMain, err error)
// 	AddModel(ctx context.Context, req *productentity.AddModelReq) (product *productentity.Product, err error)
// 	GetAllProduct(ctx context.Context, req *productentity.GetProductsReq) (products *productentity.GetProductsRes, err error)
// 	UpdateProduct(ctx context.Context, req *productentity.UpdateProductReq) (product *productentity.Product, err error)
// 	DeleteProduct(ctx context.Context, req *productentity.DeleteProductReq) (res *productentity.GeneralResponseProduct, err error)
// 	CreateCategory(ctx context.Context, req *productentity.CreateCategoryReq) (res *productentity.GeneralResponseProduct, err error)
// 	GetAllCategory(ctx context.Context) (categories *productentity.GetcategoriesRes, err error)
// 	UpdateCategory(ctx context.Context, req *productentity.UpdateCategoryReq) (res *productentity.GeneralResponseProduct, err error)
// 	DeleteCategory(ctx context.Context, req *productentity.DeleteCategoryReq) (res *productentity.GeneralResponseProduct, err error)
// 	UpdateName(ctx context.Context, req *productentity.UpdateNameReq) (res *productentity.GeneralResponseProduct, err error)
// 	GetMainProduct(ctx context.Context,field,value string)(products []*productentity.ProductMain, err error)
// }

func (p *ProductService) CreateProduct(ctx context.Context, req *productentity.CreateProductReq) (res *productentity.ProductMain, err error) {
	res1, err := p.O.ProductService().CreateProduct(ctx, p.Req.CreateProductReq(req))
	if err != nil {
		return nil, err
	}
	return p.Res.MainProduct(res1), nil
}

func (p *ProductService) AddModel(ctx context.Context, req *productentity.AddModelReq) (res *productentity.Product, err error) {
	res1, err := p.O.ProductService().AddModel(ctx, p.Req.AddModelReq(req))
	if err != nil {
		return nil, err
	}
	return p.Res.Product(res1), nil
}

func (p *ProductService) GetAllProduct(ctx context.Context, req *productentity.GetProductsReq) (res *productentity.GetProductsRes, err error) {
	res1, err := p.O.ProductService().GetAllProduct(ctx, p.Req.GetProductsReq(req))
	if err != nil {
		return nil, err
	}
	return p.Res.Products(res1, res1.Count), nil
}

func (p *ProductService) UpdateProduct(ctx context.Context, req *productentity.UpdateProductReq) (res *productentity.Product, err error) {
	res1, err := p.O.ProductService().UpdateProduct(ctx, p.Req.UpdateProductReq(req))
	if err != nil {
		return nil, err
	}
	return p.Res.Product(res1), nil
}

func (p *ProductService) DeleteProduct(ctx context.Context, req *productentity.DeleteProductReq) (res *productentity.GeneralResponseProduct, err error) {
	res1, err := p.O.ProductService().DeleteProduct(ctx, p.Req.DeleteProductReq(req))
	if err != nil {
		return nil, err
	}
	res2 := p.Res.GeneralResponseProduct(res1)
	return &res2, nil
}
