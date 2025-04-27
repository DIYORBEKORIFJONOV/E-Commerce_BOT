package product_service

import (
	clientgrpcserver "api_gateway/internal/infastructure/client_grpc_server"
	"api_gateway/internal/service/order/adjust/adjustresponse"
	adjustrequestproduct "api_gateway/internal/service/product/adjustRequest"
)


type OrderServiceforbot struct {
	O   clientgrpcserver.ServiceClient
	Req *adjustrequestproduct.AdjustRequest
	Res *adjustresponse.AdjustResponse
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