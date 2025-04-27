package product_service

import (
	"context"

	productentity "github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/entity/product"
	clientgrpcserver "github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/infastructure/client_grpc_server"
	adjustrequestproduct "github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/service/product/adjustRequest"
	adjustresponseproduct "github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/service/product/adjustResponse"
	productpb "github.com/diyorbek/E-Commerce_BOT/api_gateway/pkg/protos/gen/product"
)

type ProductService struct {
	O   clientgrpcserver.ServiceClient
	Req *adjustrequestproduct.AdjustRequest
	Res *adjustresponseproduct.AdjustResponse
}

func NewProductReqService(O *clientgrpcserver.ServiceClient)*ProductService {
	res := adjustresponseproduct.NewAddJsutResponse()
	req := adjustrequestproduct.NewProductReq()
	
	return &ProductService{
		O: *O,
		Req: req,
		Res: res,
	}
}


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

func (p *ProductService) CreateCategory(ctx context.Context,
	req *productentity.CreateCategoryReq) (*productentity.GeneralResponseProduct, error) {

	res, err := p.O.ProductService().CreateCategory(ctx, p.Req.CreateCategoryReq(req))
	if err != nil {
		return &productentity.GeneralResponseProduct{}, err
	}
	res1 :=p.Res.GeneralResponseProduct(res)
	return &res1, nil
}

func (p *ProductService) GetAllCategory(ctx context.Context) (categories *productentity.GetcategoriesRes, err error) {
	res, err := p.O.ProductService().GetAllCategory(ctx, &productpb.GetCategoriesReq{})
	if err != nil {
		return nil, err
	}

	return p.Res.Categories(res.Category, res.Count), nil
}

func (p *ProductService) UpdateCategory(ctx context.Context, req *productentity.UpdateCategoryReq) (res *productentity.GeneralResponseProduct, err error) {
	response, err := p.O.ProductService().UpdateCategory(ctx, p.Req.UpdateCategoryReq(req))
	if err != nil {
		return nil, err
	}

	res1 := p.Res.GeneralResponseProduct(response)
	return &res1, nil
}

func (p *ProductService) DeleteCategory(ctx context.Context, req *productentity.DeleteCategoryReq) (res *productentity.GeneralResponseProduct, err error) {
	response, err := p.O.ProductService().DeleteCategory(ctx, p.Req.DeleteCategoryReq(req))
	if err != nil {
		return nil, err
	}
	req1 := p.Res.GeneralResponseProduct(response)

	return &req1, nil
}


func (p *ProductService) UpdateName(ctx context.Context, req *productentity.UpdateNameReq) (res *productentity.GeneralResponseProduct, err error) {
	response, err := p.O.ProductService().UpdateName(ctx, p.Req.UpdateNameReq(req))
	if err != nil {
	  return nil, err
	}
	req1 := p.Res.GeneralResponseProduct(response)
  
	return &req1, nil
  }
  
  func (p *ProductService) GetMainProduct(ctx context.Context, field, value string) (res []*productentity.ProductMain, err error) {
	res1, err := p.O.ProductService().GetMainProduct(ctx, p.Req.GetMainProductReq(field, value))
	if err != nil {
	  return nil, err
	}
	r := p.Res.MainProducts(res1, res1.Count)
	return r, nil
  }
  