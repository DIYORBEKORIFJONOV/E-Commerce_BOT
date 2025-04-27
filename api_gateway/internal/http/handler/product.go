package handler

import (
	productentity "api_gateway/internal/entity/product"
	"api_gateway/internal/infastructure/minao"
	productusecase "api_gateway/internal/usecase/product"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	product *productusecase.ProductUseCaseIml
    minioPhoto *minao.Client
}

func NewProductHandler(
    product *productusecase.ProductUseCaseIml,
    minioClient *minao.Client,
) *ProductHandler {
    return &ProductHandler{
        product:    product,
        minioPhoto: minioClient,
    }
}



// CreateProduct godoc
// @Summary      Create a new product
// @Description  Создаёт новый продукт на основе JSON-запроса
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        product  body      productentity.CreateProductReq  true  "Product payload"
// @Success      201      {object}  productentity.Product          "Created"
// @Failure      400      {object}  productentity.ErrorResponse    "Bad Request"
// @Failure      500      {object}  productentity.ErrorResponse    "Internal Server Error"
// @Router       /products [post]
func (p *ProductHandler) CreateProduct(c *gin.Context) {
    var req productentity.CreateProductReq

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, productentity.ErrorResponse{
            Error: err.Error(),
        })
        return
    }

    mainProduct, err := p.product.CreateProduct(c.Request.Context(), &req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, productentity.ErrorResponse{
            Error: err.Error(),
        })
        return
    }

    c.JSON(http.StatusCreated, mainProduct)
}



// AddModel godoc
// @Summary      Добавляет модель товара с фото
// @Description  Принимает form-data: все поля AddModelReq + файл photo
// @Tags         products
// @Accept       multipart/form-data
// @Produce      json
// @Param        mainproductId formData string true  "MainProductId"
// @Param        description    formData string false "Description"
// @Param        colour         formData string false "Colour"
// @Param        size           formData int    false "Size"
// @Param        price          formData number false "Price"
// @Param        quantity       formData int    false "Quantity"
// @Param        photo          formData file   true  "Photo file"
// @Success      201 {object} productentity.Product  "Созданный товар"
// @Failure      400 {object} productentity.ErrorResponse          "Некорректный запрос"
// @Failure      500 {object} productentity.ErrorResponse          "Ошибка сервера"
// @Router       /products/addmodel [post]
func (p *ProductHandler) AddModel(c *gin.Context) {
    var req productentity.AddModelReq
    if err := c.ShouldBind(&req); err != nil {
        c.JSON(http.StatusBadRequest,productentity.ErrorResponse{Error: err.Error()})
        return
    }
    fileHeader, err := c.FormFile("photo")
    if err != nil {
        c.JSON(http.StatusBadRequest,productentity.ErrorResponse{Error: "photo is required"})
        return
    }
    objectName, err := p.minioPhoto.AddPhoto(c.Request.Context(), fileHeader)
    if err != nil {
        c.JSON(http.StatusInternalServerError,productentity.ErrorResponse{Error: err.Error()})
        return
    }
    req.PhotoURL = objectName
    product, err := p.product.AddModel(c.Request.Context(), &req)
    if err != nil {
        c.JSON(http.StatusBadRequest, productentity.ErrorResponse{Error: err.Error()})
        return
    }
    c.JSON(http.StatusCreated, product)
}



// UpdateProductName godoc
// @Summary      Update product name
// @Description  Обновляет название продукта по ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        payload  body      productentity.UpdateNameReq          true  "Payload для обновления названия продукта"
// @Success      200      {object}  productentity.GeneralResponseProduct  "Успешный ответ с обновлённым продуктом"
// @Failure      400      {object}  productentity.ErrorResponse          "Неверный запрос"
// @Failure      500      {object}  productentity.ErrorResponse          "Внутренняя ошибка сервера"
// @Router       /products/name [put]
func (p *ProductHandler) UpdateProductName(c *gin.Context) {
    var req productentity.UpdateNameReq

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, productentity.ErrorResponse{Error: err.Error()})
        return
    }

    generalRes, err := p.product.UpdateName(c.Request.Context(), &req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, productentity.ErrorResponse{Error: err.Error()})
        return
    }

    c.JSON(http.StatusOK, generalRes)
}


func (p *ProductHandler)GetMainProduct(c *gin.Context) {
    
    field := c.Query("field")
    page := c.Query("page")

   mainProducts,err :=  p.product.GetMainProduct(c.Request.Context(),field,page)
   if err != nil {
    c.JSON(http.StatusInternalServerError,productentity.ErrorResponse{Error: err.Error()})
    return
   }

   c.JSON(http.StatusOK,mainProducts)
}




// GetAllProduct godoc
// @Summary      List products
// @Description  Получить список продуктов с фильтрацией и пагинацией
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        field   query     string  false  "Поле для фильтрации (например, name, category)"
// @Param        value   query     string  false  "Значение для фильтрации"
// @Param        page    query     int     false  "Номер страницы"       default(1)
// @Param        limit   query     int     false  "Размер страницы"      default(10)
// @Success      200     {array}   productentity.GetProductsRes        "Успешный ответ — список продуктов"
// @Failure      400     {object}  productentity.ErrorResponse "Ошибка в параметрах запроса"
// @Failure      500     {object}  productentity.ErrorResponse "Внутренняя ошибка сервера"
// @Router       /products [get]
func (p *ProductHandler) GetAllProduct(c *gin.Context) {
    var req productentity.GetProductsReq

    if err := c.ShouldBindQuery(&req); err != nil {
        c.JSON(http.StatusBadRequest, productentity.ErrorResponse{Error: err.Error()})
        return
    }

    products, err := p.product.GetAllProduct(c.Request.Context(), &req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, productentity.ErrorResponse{Error: err.Error()})
        return
    }

    c.JSON(http.StatusOK, products)
}



// UpdateProduct godoc
// @Summary      Partially update an existing product
// @Description  Обновляет продукт по данным из тела запроса (multipart/form-data). Фото обновляется только если предоставлено.
// @Tags         products
// @Accept       multipart/form-data
// @Produce      json
// @Param        id           formData string  true  "Product ID"
// @Param        name        formData string  false "Product name"
// @Param        category    formData string  false "Category"
// @Param        description formData string  false "Description"
// @Param        photo       formData file    false "Photo file"
// @Success      200         {object} productentity.Product           "Успешный ответ — обновлённый продукт"
// @Failure      400         {object} productentity.ErrorResponse     "Ошибка в данных запроса"
// @Failure      500         {object} productentity.ErrorResponse     "Внутренняя ошибка сервера"
// @Router       /products [patch]
func (p *ProductHandler) UpdateProduct(c *gin.Context) {
    var req productentity.UpdateProductReq

    if err := c.ShouldBind(&req); err != nil {
        c.JSON(http.StatusBadRequest, productentity.ErrorResponse{Error: err.Error()})
        return
    }

    fileHeader, err := c.FormFile("photo")

    if err == nil {
        objectName, err := p.minioPhoto.AddPhoto(c.Request.Context(), fileHeader)
        if err != nil {
            c.JSON(http.StatusInternalServerError, productentity.ErrorResponse{Error: err.Error()})
            return
        }
        req.PhotoURL = objectName
    } else {
        req.PhotoURL = ""
        if err != http.ErrMissingFile {
     
            c.JSON(http.StatusBadRequest, productentity.ErrorResponse{Error: "Error with photo upload"})
            return
        }
    }

    product, err := p.product.UpdateProduct(c.Request.Context(), &req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, productentity.ErrorResponse{Error: err.Error()})
        return
    }

    c.JSON(http.StatusOK, product)
}


// DeleteProduct godoc
// @Summary      Delete an existing product
// @Description  Удаляет продукт по переданному идентификатору.
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        payload  body      productentity.DeleteProductReq      true  "Payload для удаления продукта"
// @Success      200      {object}  productentity.GeneralResponseProduct  "Результат удаления"
// @Failure      400      {object}  productentity.ErrorResponse          "Неверный запрос"
// @Failure      500      {object}  productentity.ErrorResponse          "Внутренняя ошибка сервера"
// @Router       /products [delete]
func (p *ProductHandler) DeleteProduct(c *gin.Context) {
    var req productentity.DeleteProductReq

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, productentity.ErrorResponse{Error: err.Error()})
        return
    }

    res, err := p.product.DeleteProduct(c.Request.Context(), &req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, productentity.ErrorResponse{Error: err.Error()})
        return
    }

    c.JSON(http.StatusOK, res)
}


// CreateCategory godoc
// @Summary      Create a new category
// @Description  Создаёт категорию
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        payload  body      productentity.CreateCategoryReq   true  "Category payload"
// @Success      200      {object}  productentity.GeneralResponseCategory
// @Failure      400      {object}  productentity.ErrorResponse
// @Failure      500      {object}  productentity.ErrorResponse
// @Router       products/categories [post]
func (p *ProductHandler) CreateCategory(c *gin.Context) {
    var req productentity.CreateCategoryReq

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, productentity.ErrorResponse{Error: err.Error()})
        return
    }

    generalRes, err := p.product.CreateCategory(c.Request.Context(), &req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, productentity.ErrorResponse{Error: err.Error()})
        return
    }

    c.JSON(http.StatusOK, generalRes)
}



// GetAllCategory godoc
// @Summary      Retrieve all categories
// @Description  Возвращает список всех категорий продуктов
// @Tags         categories
// @Accept       json
// @Produce      json
// @Success      200  {object}  productentity.GetcategoriesRes
// @Failure      500  {object}  productentity.ErrorResponse
// @Router       products/categories [get]
func (p *ProductHandler) GetAllCategory(c *gin.Context) {
    categories, err := p.product.GetAllCategory(c.Request.Context())
    if err != nil {
        c.JSON(http.StatusInternalServerError, productentity.ErrorResponse{Error: err.Error()})
        return
    }
    c.JSON(http.StatusOK, categories)
}



// UpdateCategory godoc
// @Summary      Update an existing category
// @Description  Обновляет название или другие параметры категории
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        payload  body      productentity.UpdateCategoryReq   true  "Category update payload"
// @Success      200      {object}  productentity.GeneralResponseCategory
// @Failure      400      {object}  productentity.ErrorResponse
// @Failure      500      {object}  productentity.ErrorResponse
// @Router       products/categories [put]
func (p *ProductHandler) UpdateCategory(c *gin.Context) {
    var req productentity.UpdateCategoryReq

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, productentity.ErrorResponse{Error: err.Error()})
        return
    }

    generalResponse, err := p.product.UpdateCategory(c.Request.Context(), &req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, productentity.ErrorResponse{Error: err.Error()})
        return
    }

    c.JSON(http.StatusOK, generalResponse)
}


// DeleteCategory godoc
// @Summary      Delete a category
// @Description  Удаляет категорию по переданному идентификатору
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        payload  body      productentity.DeleteCategoryReq   true  "Category deletion payload"
// @Success      200      {object}  productentity.GeneralResponseCategory
// @Failure      400      {object}  productentity.ErrorResponse
// @Failure      500      {object}  productentity.ErrorResponse
// @Router       products/categories [delete]
func (p *ProductHandler) DeleteCategory(c *gin.Context) {
    var req productentity.DeleteCategoryReq

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, productentity.ErrorResponse{Error: err.Error()})
        return
    }

    generalRes, err := p.product.DeleteCategory(c.Request.Context(), &req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, productentity.ErrorResponse{Error: err.Error()})
        return
    }

    c.JSON(http.StatusOK, generalRes)
}

