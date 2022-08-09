package CartModel

import (
	"database/sql"

	db "github.com/aydinsercan/basketapi/model"
)

type (
	Cart     struct{ db.DB }
	CartData struct {
		Id            int64         `db:"id" json:"id"`
		ProductId     int64         `db:"productId" json:"product_id" validate:"required"`
		ProductName   string        `db:"productName" json:"product_name"`
		Qty           int           `db:"qty" json:"qty" validate:"required"`
		AID           string        `db:"AID" json:"-"`
		UserId        sql.NullInt64 `db:"userId" json:"-"`
		IsOrdered     int           `db:"isOrdered"`
		TotalPrice    int           `db:"totalprice" json:"totalprice"`
		VatPercentage int           `db:"vatpercentage" json:"vatpercentage"`
	}
)

func (c *Cart) List(param CartData) ([]CartData, error) {
	return c.listByAID(param.AID)
}

func (c *Cart) listByAID(AID string) ([]CartData, error) {

	strSQL := `SELECT a.id, a.productId, b.productName, a.qty, 
				CASE
					WHEN a.qty <=3 THEN b.price*a.qty
					WHEN a.qty > 3 THEN b.price*a.qty - TRUNCATE((b.price*8)/100,0)* (a.qty-3)
				END as totalprice, 
				b.vatpercentage as vatpercentage FROM carts 
				a INNER JOIN products b ON a.productId=b.id WHERE AID=? AND isOrdered=0`
	data := []CartData{}
	pool := c.GetPool()
	err := pool.Select(&data, strSQL, AID)

	return data, err
}

func (c *Cart) Insert(param *CartData) (int64, error) {
	args := []interface{}{}
	strSQL := `SELECT id, qty FROM carts WHERE `

	strSQL += `AID=?`
	args = append(args, param.AID)

	strSQL += ` AND productId=?`
	args = append(args, param.ProductId)

	pool := c.GetPool()

	cart := CartData{}
	if err := pool.Get(&cart, strSQL, args...); err != nil && err != sql.ErrNoRows {
		return 0, err
	}

	if !((CartData{}) == cart) {
		param.Id = cart.Id
		param.Qty += cart.Qty
		return param.Id, c.Update(param)
	}

	strSQL = `INSERT INTO carts (productId, qty, AID, userId) VALUES (?, ?, ?, ?)`
	args = []interface{}{param.ProductId, param.Qty, param.AID, param.UserId}
	res, err := pool.Exec(strSQL, args...)
	if err != nil {
		return 0, err
	}
	insertId, _ := res.LastInsertId()

	return insertId, err
}

func (c *Cart) Update(param *CartData) error {
	pool := c.GetPool()
	strSQL := `UPDATE carts SET qty=? WHERE id=?`
	_, err := pool.Exec(strSQL, param.Qty, param.Id)
	return err
}

func (c *Cart) Delete(id int64) error {
	pool := c.GetPool()
	strSQL := `DELETE FROM carts WHERE id=?`
	_, err := pool.Exec(strSQL, id)
	return err
}

//Flush all data placeOrder
func (c *Cart) Order() error {
	pool := c.GetPool()
	strSQL := `DELETE FROM carts`
	_, err := pool.Exec(strSQL)
	return err
}
