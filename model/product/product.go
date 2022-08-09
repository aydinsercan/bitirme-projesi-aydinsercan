package ProductModel

import (
	db "github.com/aydinsercan/basketapi/model"
)

type (
	Product     struct{ db.DB }
	ProductData struct {
		Id            int64  `db:"id" json:"id"`
		ProductName   string `db:"productName" json:"product_name"`
		Price         int    `db:"price" json:"price"`
		VatPercentage int    `db:"vatpercentage" json:"vatpercentage"`
	}
)

func (p *Product) List() ([]ProductData, error) {
	strSQL := `SELECT id, productName, price, vatpercentage FROM products`
	data := []ProductData{}
	pool := p.GetPool()
	err := pool.Select(&data, strSQL)
	return data, err
}

func (p *Product) Get(productId int64) (ProductData, error) {
	strSQL := `SELECT id, productName, price, vatpercentage FROM products WHERE id=?`
	data := ProductData{}
	pool := p.GetPool()
	err := pool.Get(&data, strSQL)
	return data, err
}
