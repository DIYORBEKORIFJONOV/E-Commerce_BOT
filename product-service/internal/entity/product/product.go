package productentity

type (
	UpdateNameReq struct {
		Name    string `json:"name"`
		NewName string `json:"newname"`
	}

	Product struct {
		ProductID   string  `json:"productId"`
		Description string  `json:"description"`
		Colour      string  `json:"colour"`
		Size        int32   `json:"size"`
		Price       float32 `json:"price"`
		Quantity    int32   `json:"quantity"`
		CreatedAt   string  `json:"createdat"`
		UpdatedAt   string  `json:"updatedat"`
		PhotoURL    string  `json:"photourl"`
		MainProductId string `json:"mainproductId"`
		IsDeleted bool `json:"isDeleted"`
	}
	ProductMain struct {
		Category string `json:"category"`
		Name string `json:"name"`
		MainProductId string `json:"mainproductId"`
	}

	CreateProductReq struct {
		Name     string `json:"name"`
		Category string `json:"category"`
	}

	AddModelReq struct {
		MainProductId string `json:"mainproductId"`
		Description string  `json:"description"`
		Colour      string  `json:"colour"`
		Size        int32   `json:"size"`
		Price       float32 `json:"price"`
		Quantity    int32   `json:"quantity"`
		PhotoURL    string  `json:"photourl"`
	}

	GetProductsReq struct {
		Field string `json:"field"`
		Value string `json:"value"`
		Page  int64  `json:"page"`
		Limit int64  `json:"limit"`
	}

	GetProductsRes struct {
		Products []Product `json:"product"`
		Count    int64     `json:"count"`
	}

	UpdateProductReq struct {
		ProductID   string  `json:"productId"`
		Description string  `json:"description"`
		Colour      string  `json:"colour"`
		Size        int32   `json:"size"`
		Price       float32 `json:"price"`
		Quantity    int32   `json:"quantity"`
		PhotoURL    string  `json:"photourl"`
	}

	DeleteProductReq struct {
		ProductID string `json:"productId"`
		IsDeleted bool   `json:"is_deleted"`
	}

	GeneralResponseProduct struct {
		Status  bool   `json:"status"`
		Message string `json:"message"`
	}

	CreateCategoryReq struct {
		Category string `json:"category"`
	}

	GetCategoriesReq struct{}

	GetcategoriesRes struct {
		Categories []*CreateCategoryReq `json:"category"`
		Count      int64               `json:"count"`
	}

	UpdateCategoryReq struct {
		Category    string `json:"category"`
		NewCategory string `json:"newcategory"`
	}

	DeleteCategoryReq struct {
		Category string `json:"category"`
	}
)
