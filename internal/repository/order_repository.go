//talks directly to the database and perform crud operation related to orders.
package repository

import (
	"database/sql"
	"mini-order-api/internal/model"
)
type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {  //function returns the pointer address
	return &OrderRepository{db: db}     //var repo OrderRepository repo.db = db
}

func (r *OrderRepository) Create(order *model.Order) error {        /*Go struct → SQL query → Database
Database → Returns ID → Stored in struct */
	query := `INSERT INTO orders (customer_name, item, quantity) 
	VALUES ($1, $2) RETURNING id`
	    return r.db.QueryRow(query,
        order.CustomerName,
        order.Amount,
    ).Scan(&order.ID)
}

func (r *OrderRepository) GetALL() ([]*model.Order,error){
	rows,err := r.db.Query("SELECT id, customer_name,amount FROM orders")
	if err != nil {
		return nil, err
	}
defer rows.Close()
//returns the result of the query as a set of rows, which can be iterated over to access the data. It is used to execute SQL SELECT statements and retrieve data from the database. The returned *sql.Rows object provides methods to iterate through the result set and scan the values into Go variables or structs.
var orders []*model.Order

for rows.Next() {
	var o model.Order
	if err := rows.Scan(&o.ID, &o.CustomerName, &o.Amount); err != nil {
		return nil,err
}
orders =append(orders, &o)
}
return orders, nil
}