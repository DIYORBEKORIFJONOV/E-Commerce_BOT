package adjustresponse

import (
	productentity "api_gateway/internal/entity/product"
	productpb "api_gateway/pkg/protos/gen/product"
)

// --- Single conversions ---

func Product(pb *productpb.Product) *productentity.Product {
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

func MainProduct(pb *productpb.MainProduct) *productentity.ProductMain {
	if pb == nil {
		return nil
	}
	return &productentity.ProductMain{
		MainProductId: pb.MainProductId,
		Category:      pb.Category,
		Name:          pb.Name,
	}
}

func GeneralResponseProduct(pb *productpb.GeneralResponseProduct) *productentity.GeneralResponseProduct {
	if pb == nil {
		return nil
	}
	return &productentity.GeneralResponseProduct{
		Status:  pb.Status,
		Message: pb.Message,
	}
}

// --- List conversions ---

func Products(pbProducts []*productpb.Product, count int64) *productentity.GetProductsRes {
	products := make([]productentity.Product, 0, len(pbProducts))
	for _, pb := range pbProducts {
		products = append(products, *Product(pb))
	}
	return &productentity.GetProductsRes{
		Products: products,
		Count:    count,
	}
}

func MainProducts(pbProducts []*productpb.MainProduct, count int64) []productentity.ProductMain {
	products := make([]productentity.ProductMain, 0, len(pbProducts))
	for _, pb := range pbProducts {
		if pb != nil {
			products = append(products, *MainProduct(pb))
		}
	}
	return products
}

func Categories(pbCategories []*productpb.CreateCategoryReq, count int64) *productentity.GetcategoriesRes {
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
