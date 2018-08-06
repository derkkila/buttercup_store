package model
type Product struct {
        Id string `json:"id"`
        Name string  `json:"name"`
        Description string `json:"description"`
        ProdType string `json:"prodtype"`
        Category string `json:"category"`
        Price string `json:"price"`
        Qty string `json:"qty"`
}
