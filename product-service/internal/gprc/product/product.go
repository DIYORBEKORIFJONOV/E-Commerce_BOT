package productgrpc

import (
	"context"
	productentity "productservice/internal/entity/product"
	productusecase "productservice/internal/usecase/product/productservice"
	productpb "productservice/pkg/protos/gen/product"

	"google.golang.org/grpc"
)

type productGrpcService struct {
	productpb.UnimplementedProductServiceServer
	product *productusecase.ProductUseCaseIml
}

func RegisterProductGrpcService(grpcServer *grpc.Server,product *productusecase.ProductUseCaseIml) {
	productpb.RegisterProductServiceServer(grpcServer, &productGrpcService{
		product: product,
	})
}

func (p *productGrpcService)CreateProduct(ctx context.Context, req  *productpb.CreateProductReq) (*productpb.MainProduct, error) {
    reqEntity := &productentity.CreateProductReq{
        Name:     req.Name,
        Category: req.Category,
    }

    res, err := p.product.CreateProduct(ctx, reqEntity)
    if err != nil {
        return nil, err
    }

    return &productpb.MainProduct{
		Category: res.Category,
		MainProductId: res.MainProductId,
		Name: res.Name,
    }, nil
}


func (p *productGrpcService) AddModel(ctx context.Context, req *productpb.AddModelReq) (*productpb.Product, error) {
	addModelReq := &productentity.AddModelReq{
		Description: req.Description,
		Colour:      req.Colour,
		Size:        req.Size,
		Price:       req.Price,
		Quantity:    req.Quantity,
		PhotoURL:    req.Photourl,
		MainProductId: req.MainProductId,
	}

	product, err := p.product.AddModel(ctx, addModelReq)
	if err != nil {
		return nil, err
	}

	return &productpb.Product{
		ProductId:   product.ProductID,
		Description: product.Description,
		Colour:      product.Colour,
		Size:        product.Size,
		Price:       product.Price,
		Quantity:    product.Quantity,
		Createdat:   product.CreatedAt,
		Updatedat:   product.UpdatedAt,
		Photourl:    product.PhotoURL,
		MainProductId: product.MainProductId,
	}, nil
}


func (p *productGrpcService) GetAllProduct(ctx context.Context, req *productpb.GetProductsReq) (*productpb.GetProductsRes, error) {
	getReq := &productentity.GetProductsReq{
		Field: req.Field,
		Value: req.Value,
		Page:  req.Page,
		Limit: req.Limit,
	}

	productsEntity, count, err := p.product.GetAllProduct(ctx, getReq)
	if err != nil {
		return nil, err
	}

	var products []*productpb.Product
	for _, prod := range productsEntity {
		products = append(products, &productpb.Product{
			ProductId:   prod.ProductID,
			Description: prod.Description,
			Colour:      prod.Colour,
			Size:        prod.Size,
			Price:       prod.Price,
			Quantity:    prod.Quantity,
			Createdat:   prod.CreatedAt,
			Updatedat:   prod.UpdatedAt,
			Photourl:    prod.PhotoURL,
			MainProductId: prod.MainProductId,
		})
	}

	return &productpb.GetProductsRes{
		Product: products,
		Count:   count,
	}, nil
}



func (p *productGrpcService) UpdateProduct(ctx context.Context, req *productpb.UpdateProductReq) (*productpb.Product, error) {
	updateReq := &productentity.UpdateProductReq{
		ProductID:   req.ProductId,
		Description: req.Description,
		Colour:      req.Colour,
		Size:        req.Size,
		Price:       req.Price,
		Quantity:    req.Quantity,
		PhotoURL: req.PhotoUrl,
	}

	updatedProduct, err := p.product.UpdateProduct(ctx, updateReq)
	if err != nil {
		return nil, err
	}

	return &productpb.Product{
		ProductId:   updatedProduct.ProductID,
		Description: updatedProduct.Description,
		Colour:      updatedProduct.Colour,
		Size:        updatedProduct.Size,
		Price:       updatedProduct.Price,
		Quantity:    updatedProduct.Quantity,
		Createdat:   updatedProduct.CreatedAt,
		Updatedat:   updatedProduct.UpdatedAt,
		Photourl:    updatedProduct.PhotoURL,
		MainProductId: updatedProduct.MainProductId,
	}, nil
}

func (p *productGrpcService) DeleteProduct(ctx context.Context, req *productpb.DeleteProductReq) (*productpb.GeneralResponseProduct, error) {
	deleteReq := &productentity.DeleteProductReq{
		ProductID: req.ProductId,
		IsDeleted: req.IsDeleted,
	}
	res, err := p.product.DeleteProduct(ctx, deleteReq)
	if err != nil {
		return nil, err
	}
	return &productpb.GeneralResponseProduct{
		Status:  res.Status,
		Message: res.Message,
	}, nil
}

func (p *productGrpcService) CreateCategory(ctx context.Context, req *productpb.CreateCategoryReq) (*productpb.GeneralResponseProduct, error) {
	res, err := p.product.CreateCategory(ctx, &productentity.CreateCategoryReq{Category: req.Category})
	if err != nil {
		return nil, err
	}
	return &productpb.GeneralResponseProduct{
		Status:  res.Status,
		Message: res.Message,
	}, nil
}

func (p *productGrpcService) GetAllCategory(ctx context.Context, req *productpb.GetCategoriesReq) (*productpb.GetcategoriesRes, error) {
	res, err := p.product.GetAllCategory(ctx, &productentity.GetCategoriesReq{})
	if err != nil {
		return nil, err
	}
	var categories []*productpb.CreateCategoryReq
	for _, cat := range res.Categories {
		categories = append(categories, &productpb.CreateCategoryReq{
			Category: cat.Category,
		})
	}
	return &productpb.GetcategoriesRes{
		Category: categories,
		Count:    res.Count,
	}, nil
}

func (p *productGrpcService) UpdateCategory(ctx context.Context, req *productpb.UpdateCategoryReq) (*productpb.GeneralResponseProduct, error) {
	updateReq := &productentity.UpdateCategoryReq{
		Category:    req.Category,
		NewCategory: req.Newcategory,
	}
	res, err := p.product.UpdateCategory(ctx, updateReq)
	if err != nil {
		return nil, err
	}
	return &productpb.GeneralResponseProduct{
		Status:  res.Status,
		Message: res.Message,
	}, nil
}

func (p *productGrpcService) DeleteCategory(ctx context.Context, req *productpb.DeleteCategoryReq) (*productpb.GeneralResponseProduct, error) {
	res, err := p.product.DeleteCategory(ctx, &productentity.DeleteCategoryReq{
		Category: req.Category,
	})
	if err != nil {
		return nil, err
	}
	return &productpb.GeneralResponseProduct{
		Status:  res.Status,
		Message: res.Message,
	}, nil
}



func (p *productGrpcService) UpdateName(ctx context.Context, req *productpb.UpdateNameReq) (*productpb.GeneralResponseProduct, error) {
	res, err := p.product.UpdateName(ctx, &productentity.UpdateNameReq{
		Name:    req.Name,
		NewName: req.Newname,
	})
	if err != nil {
		return nil, err
	}
	return &productpb.GeneralResponseProduct{
		Status:  res.Status,
		Message: res.Message,
	}, nil
}



func (p *productGrpcService)GetMainProduct(ctx context.Context, req *productpb.GetMainProductReq) (*productpb.GetMainProductRes, error) {
	mainProducts,err := p.product.GetMainProduct(ctx,req.Field,req.Value)
	if err != nil {
		return nil,err
	}
	response := productpb.GetMainProductRes{}
	for _,m := range mainProducts {
		response.Products = append(response.Products,&productpb.MainProduct{
			Category: m.Category,
			Name: m.Name,
			MainProductId:m.MainProductId ,
		})
		response.Count++
	}

	return &response,err
}