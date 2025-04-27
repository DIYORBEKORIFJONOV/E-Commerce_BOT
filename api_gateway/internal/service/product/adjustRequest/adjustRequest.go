package adjustrequestproduct

import (
	productentity "github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/entity/product"
	productpb "github.com/diyorbek/E-Commerce_BOT/api_gateway/pkg/protos/gen/product"
)

type AdjustRequest struct {
}

func NewProductReq() *AdjustRequest {
	return &AdjustRequest{}
}

func (p *AdjustRequest)CreateProductReq(req *productentity.CreateProductReq) *productpb.CreateProductReq {
	return &productpb.CreateProductReq{
		Name:     req.Name,
		Category: req.Category,
	}
}

func (p *AdjustRequest)AddModelReq(req *productentity.AddModelReq) *productpb.AddModelReq {
	return &productpb.AddModelReq{
		MainProductId: req.MainProductId,
		Description:   req.Description,
		Colour:        req.Colour,
		Size:          req.Size,
		Price:         req.Price,
		Quantity:      req.Quantity,
		Photourl:      req.PhotoURL,
	}
}

func (p *AdjustRequest)GetProductsReq(req *productentity.GetProductsReq) *productpb.GetProductsReq {
	return &productpb.GetProductsReq{
		Field: req.Field,
		Value: req.Value,
		Page:  req.Page,
		Limit: req.Limit,
	}
}

func (p *AdjustRequest)UpdateProductReq(req *productentity.UpdateProductReq) *productpb.UpdateProductReq {
	return &productpb.UpdateProductReq{
		ProductId:   req.ProductID,
		Description: req.Description,
		Colour:      req.Colour,
		Size:        req.Size,
		Price:       req.Price,
		Quantity:    req.Quantity,
		PhotoUrl:    req.PhotoURL,
	}
}

func (p *AdjustRequest)DeleteProductReq(req *productentity.DeleteProductReq) *productpb.DeleteProductReq {
	return &productpb.DeleteProductReq{
		ProductId: req.ProductID,
		IsDeleted: req.IsDeleted,
	}
}

func (p *AdjustRequest)CreateCategoryReq(req *productentity.CreateCategoryReq) *productpb.CreateCategoryReq {
	return &productpb.CreateCategoryReq{
		Category: req.Category,
	}
}

func (p *AdjustRequest)UpdateCategoryReq(req *productentity.UpdateCategoryReq) *productpb.UpdateCategoryReq {
	return &productpb.UpdateCategoryReq{
		Category:    req.Category,
		Newcategory: req.NewCategory,
	}
}

func (p *AdjustRequest)DeleteCategoryReq(req *productentity.DeleteCategoryReq) *productpb.DeleteCategoryReq {
	return &productpb.DeleteCategoryReq{
		Category: req.Category,
	}
}

func (p *AdjustRequest)UpdateNameReq(req *productentity.UpdateNameReq) *productpb.UpdateNameReq {
	return &productpb.UpdateNameReq{
		Name:    req.Name,
		Newname: req.NewName,
	}
}

func (p *AdjustRequest)GetMainProductReq(field, value string) *productpb.GetMainProductReq {
	return &productpb.GetMainProductReq{
		Field: field,
		Value: value,
	}
}
