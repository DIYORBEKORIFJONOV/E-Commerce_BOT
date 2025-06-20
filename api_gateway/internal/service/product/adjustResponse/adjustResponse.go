package adjustresponseproduct

import (
	productentity "github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/entity/product"
	productpb "github.com/diyorbek/E-Commerce_BOT/api_gateway/pkg/protos/gen/product"
)

type AdjustResponse struct{}

func NewAddJsutResponse()*AdjustResponse {
	return &AdjustResponse{}
}
func (p *AdjustResponse) Product(pb *productpb.Product) *productentity.Product {
	if pb == nil {
		return nil
	}
	return &productentity.Product{
		ProductID:     pb.ProductId,
		Description:   pb.Description,
		Colour:        pb.Colour,
		Size:          pb.Size,
		Price:         pb.Price,
		Quantity:      pb.Quantity,
		CreatedAt:     pb.Createdat,
		UpdatedAt:     pb.Updatedat,
		PhotoURL:      pb.Photourl,
		MainProductId: pb.MainProductId,
	}
}

func (p *AdjustResponse) MainProduct(pb *productpb.MainProduct) *productentity.ProductMain {
	if pb == nil {
		return nil
	}
	return &productentity.ProductMain{
		MainProductId: pb.MainProductId,
		Category:      pb.Category,
		Name:          pb.Name,
	}
}

func (p *AdjustResponse) GeneralResponseProduct(pb *productpb.GeneralResponseProduct) productentity.GeneralResponseProduct {
	if pb == nil {
		return productentity.GeneralResponseProduct{}
	}
	return productentity.GeneralResponseProduct{
		Status:  pb.Status,
		Message: pb.Message,
	}
}

// --- List conversions ---

func (p *AdjustResponse) Products(pbProducts *productpb.GetProductsRes, count int64) *productentity.GetProductsRes {
	products := make([]productentity.Product, 0, len(pbProducts.Product))
	for _, pb := range pbProducts.Product {
		if pb != nil {
			products = append(products, productentity.Product{
				ProductID:     pb.ProductId,
				MainProductId: pb.MainProductId,
				PhotoURL:      pb.Photourl,
				Colour:        pb.Colour,
				Size:          pb.Size,
				Price:         pb.Price,
				Quantity:      pb.Quantity,
				Description:   pb.Description,
				CreatedAt:     pb.Createdat,
				UpdatedAt:     pb.Updatedat,
			})
		}
	}

	return &productentity.GetProductsRes{
		Products: products,
		Count:    count,
	}
}

func (p *AdjustResponse) MainProducts(pbProducts *productpb.GetMainProductRes, count int64) []*productentity.ProductMain {
	products := make([]*productentity.ProductMain, 0, len(pbProducts.Products)) 
  
	for _, pb := range pbProducts.Products {
	  if pb != nil {
		products = append(products, &productentity.ProductMain{
		  Category:      pb.Category,
		  Name:          pb.Name,
		  MainProductId: pb.MainProductId,
		})
	  }
	}
  
	return products
  }
  

func (p *AdjustResponse) Categories(pbCategories []*productpb.CreateCategoryReq, count int64) *productentity.GetcategoriesRes {
	categories := make([]*productentity.CreateCategoryReq, 0, len(pbCategories))
	for _, pb := range pbCategories {
		categories = append(categories, &productentity.CreateCategoryReq{
			Category: pb.Category,
		})
	}
	return &productentity.GetcategoriesRes{
		Categories: categories,
		Count:      count,
	}
}
