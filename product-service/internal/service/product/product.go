package productservice

import (
	"context"
	"fmt"
	productentity "productservice/internal/entity/product"
	productuserepository "productservice/internal/usecase/product/productrepository"
	logger "productservice/log"
	"time"

	"github.com/google/uuid"
)

type ProductService struct {
	logger * logger.Logger
	product *productuserepository.ProductUseRepositoryIml
}

func NewProductService(
	logger *logger.Logger,
	product *productuserepository.ProductUseRepositoryIml,
)(*ProductService) {
	
	return	&ProductService{
		logger: logger,
		product: product,
	}
}




const op  = "operation"

func (p *ProductService)CreateProduct(ctx context.Context, req *productentity.CreateProductReq) (res *productentity.ProductMain, err error) {
	logCtx := p.logger.WithContext(map[string]string{
		"operation": "CreateProduct",
		"product":  req.Name,
		"categtory":req.Category,
	})

	cotigories,err := p.product.GetAllCategory(ctx)
	if err != nil {
		logCtx.Error(err,err.Error())
		return nil,err
	}

	var checkCategory bool = false

	for _,category := range cotigories {
		if category.Category == req.Category {
			checkCategory = true
			break
		}
	}
	if !checkCategory {
		logCtx.Error(fmt.Errorf("category does not exists"),"category does not exists")
		return nil,fmt.Errorf("category does not exists")
	}

	r,err := p.product.GetMainProduct(ctx,"name",req.Name)
	if err != nil {
		if err.Error() != "products not found" {
			logCtx.Error(err,err.Error())
			return nil,err
		}
	}
	if r != nil {
		logCtx.Error(fmt.Errorf("name already exists"),"name already exists")
		return nil,fmt.Errorf("name already exists")
	}


	productMain := &productentity.ProductMain{
		Category: req.Category,
		Name: req.Name,
		MainProductId: uuid.NewString(),

	}
	err = p.product.AddProduct(ctx,productMain)
	if err != nil {
		logCtx.Error(err,err.Error())
		return nil,err
	}

	

	logCtx.Info("Product succsessfuly created")
	return productMain,nil
}


func (p *ProductService)AddModel(ctx context.Context, req *productentity.AddModelReq) (*productentity.Product, error) {
	logCtx := p.logger.WithContext(map[string]string{
		"operation": "AddModel",
		"productId":req.MainProductId,
	})

	mainProduct,err := p.product.GetMainProduct(ctx,"mainproductid",req.MainProductId)
	if err != nil {
		logCtx.Error(err,err.Error())
		return nil,err
	}
	
	product := &productentity.Product{
		ProductID:uuid.NewString(),
		Description: req.Description,
		Colour: req.Colour,
		Size: req.Size,
		Price: req.Price,
		Quantity: req.Quantity,
		PhotoURL: req.PhotoURL,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:  time.Now().Format("2006-01-02 15:04:05"),
		MainProductId: mainProduct[0].MainProductId,
	}

	err = p.product.AddModelToProduct(ctx,product)
	if err != nil {
		logCtx.Error(err,err.Error())
		return nil,err
	}

	logCtx.Info("model successfulty added")
	return product,err
}


func (p *ProductService)GetAllProduct(ctx context.Context, req *productentity.GetProductsReq) ([]*productentity.Product, int64, error) {
	logCtx := p.logger.WithContext(map[string]string{
		"operation":"GetAllProduct",
	})	

	products,err := p.product.GetAllProduct(ctx,req)
	if err != nil {
		logCtx.Error(err,err.Error())
		return nil,0,err
	}

	logCtx.Info("getting successfuly")
	return products,int64(len(products)),nil
}

func (p *ProductService)UpdateProduct(ctx context.Context, req *productentity.UpdateProductReq) (product *productentity.Product, err error) {
	logCtx := p.logger.WithContext(map[string]string{
		"operation":"UpdateProduct",
	})
	
	product,err = p.product.UpdateProduct(ctx,req)
	if err != nil {
		logCtx.Error(err,err.Error())
		return nil,err
	}

	logCtx.Info("product successfuly updated")

	return product,nil
}

func (p *ProductService)DeleteProduct(ctx context.Context, req *productentity.DeleteProductReq) (res *productentity.GeneralResponseProduct, err error) {
	logCtx := p.logger.WithContext(map[string]string{
		"operation":"DeleteProduct",
	})

	err = p.product.DeleteProduct(ctx,req.ProductID,req.IsDeleted)
	if err != nil {
		logCtx.Error(err,err.Error())
		return nil,err
	}
	
	logCtx.Info("successfuly deleted")
	return	 &productentity.GeneralResponseProduct{
		Status: true,
		Message: "successfuly deleted",
	},nil

}


func (p *ProductService)CreateCategory(ctx context.Context, req *productentity.CreateCategoryReq) (res *productentity.GeneralResponseProduct, err error){
	logCtx := p.logger.WithContext(map[string]string{
		"operation":"CreateCategory",
	})
	categories,err := p.product.GetAllCategory(ctx)
	if err != nil {
		logCtx.Error(err,err.Error())
		return nil,err
	}

	for _,c:= range categories {
		if c.Category == req.Category {
			err := fmt.Errorf("category already exists")
			logCtx.Error(err,err.Error())
			return nil,err
		}
	}

	err= p.product.CreateCategory(ctx,req.Category)
	if err != nil {
		logCtx.Error(err,err.Error())
		return nil,err
	}

	message := "category successfuly created"

	logCtx.Info(message)
	return &productentity.GeneralResponseProduct{
		Status: true,
		Message: message,
	},nil
}

func (c *ProductService)GetAllCategory(ctx context.Context, req *productentity.GetCategoriesReq) (*productentity.GetcategoriesRes, error) {
	logCtx := c.logger.WithContext(map[string]string{
		op:"GetAllCategory",
	})

	categories,err := c.product.GetAllCategory(ctx)
	if err != nil {
		logCtx.Error(err,err.Error())
		return nil,err
	}
	

	logCtx.Info("successfuly getting")
	return &productentity.GetcategoriesRes{
		Categories: categories,
		Count:int64(len(categories)),
	},nil
}

func (c *ProductService)UpdateCategory(ctx context.Context, req *productentity.UpdateCategoryReq) (res *productentity.GeneralResponseProduct, err error) {
	logCtx := c.logger.WithContext(map[string]string{
		op:"UpdateCategory",
	})

	err = c.product.UpdateCategory(ctx,req.Category,req.NewCategory)
	if err != nil {
		logCtx.Error(err,err.Error())
		return nil,err
	}

	message := "category successfuly updated"
	
	logCtx.Info(message)

	return &productentity.GeneralResponseProduct{
		Message: message,
		Status: true,
	},nil
}

func (p *ProductService)DeleteCategory(ctx context.Context, req *productentity.DeleteCategoryReq) (res *productentity.GeneralResponseProduct, err error) {
	logCtx := p.logger.WithContext(map[string]string{
		op:"DeleteCategory",
	})

	err = p.product.DeleteCategory(ctx,req.Category)
	if err != nil {
		logCtx.Error(err,err.Error())
		return nil,err
	}

	message := "category successfuly deleted"

	logCtx.Info(message)

	return &productentity.GeneralResponseProduct{
		Message: message,
		Status: true,
	},nil
}

func (c *ProductService)UpdateName(ctx context.Context, req *productentity.UpdateNameReq) (res *productentity.GeneralResponseProduct, err error) {
	logCtx := c.logger.WithContext(map[string]string{
		op:"UpdateName",
	})

	err = c.product.Update(ctx,req.Name,req.NewName)
	if err != nil {
		logCtx.Error(err,err.Error())
		return nil,err
	}

	message := "name successfuly updated"

	logCtx.Info(message)

	return &productentity.GeneralResponseProduct{
		Message: message,
		Status: true,
	},nil
}


func (c *ProductService)GetMainProduct(ctx context.Context,field,value string)(products []*productentity.ProductMain, err error) {
	logCtx := c.logger.WithContext(map[string]string{
		op:"GetMainProduct",
	})

	mainProduct,err := c.product.GetMainProduct(ctx,field,value)
	if err != nil {
		logCtx.Error(err,err.Error())
		return nil,err
	}

	return mainProduct,nil
	
}