package product_service

import (
	productentity "api_gateway/internal/entity/product"
	clientgrpcserver "api_gateway/internal/infastructure/client_grpc_server"
	adjustrequestproduct "api_gateway/internal/service/product/adjustRequest"
	adjustresponseproduct "api_gateway/internal/service/product/adjustResponse"
	productpb "api_gateway/pkg/protos/gen/product"
	"context"

)


type ProductService struct {
	O   clientgrpcserver.ServiceClient
	Req *adjustrequestproduct.AdjustRequest
	Res *adjustresponseproduct.AdjustResponse
}


// type productUseCase interface {
// 	DeleteCategory(ctx context.Context, req *productentity.DeleteCategoryReq) (res *productentity.GeneralResponseProduct, err error)
// 	UpdateName(ctx context.Context, req *productentity.UpdateNameReq) (res *productentity.GeneralResponseProduct, err error)
// 	GetMainProduct(ctx context.Context,field,value string)(products []*productentity.ProductMain, err error)
// }

func (p *ProductService)CreateCategory(ctx context.Context, 
	req *productentity.CreateCategoryReq) (productentity.GeneralResponseProduct, error) {
	
	res,err := p.O.ProductService().CreateCategory(ctx,p.Req.CreateCategoryReq(req))
	if err != nil {
		return productentity.GeneralResponseProduct{},err
	}

	return p.Res.GeneralResponseProduct(res),nil
}


func (p *ProductService)GetAllCategory(ctx context.Context) (categories *productentity.GetcategoriesRes, err error) {
	res,err := p.O.ProductService().GetAllCategory(ctx,&productpb.GetCategoriesReq{})
	if err != nil {
		return nil,err
	}

	return p.Res.Categories(res.Category,res.Count),nil
}

func (p *ProductService)UpdateCategory(ctx context.Context, req *productentity.UpdateCategoryReq) (res *productentity.GeneralResponseProduct, err error) {
	response,err := p.O.ProductService().UpdateCategory(ctx,p.Req.UpdateCategoryReq(req))
	if err != nil {
		return nil,err
	}

	res1 := p.Res.GeneralResponseProduct(response)
	return &res1,nil
}


func (p *ProductService)DeleteCategory(ctx context.Context, req *productentity.DeleteCategoryReq) (res *productentity.GeneralResponseProduct, err error) {
	response,err := p.O.ProductService().DeleteCategory(ctx,p.Req.DeleteCategoryReq(req))
	if err != nil {
		return nil,err
	}
	req1 :=  p.Res.GeneralResponseProduct(response)

	return &req1,nil
}