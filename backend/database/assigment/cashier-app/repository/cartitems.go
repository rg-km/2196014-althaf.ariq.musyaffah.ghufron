package repository

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type CartItemRepository struct {
	db *sql.DB
}

func NewCartItemRepository(db *sql.DB) *CartItemRepository {
	return &CartItemRepository{db: db}
}

func (c *CartItemRepository) FetchCartItems() ([]CartItem, error) {
	//var sqlStatement string
	var cartItems []CartItem

	//TODO: add sql statement here
	//HINT: join table cart_items and products
	sqlStatement := `SELECT 
	cart_items.id,
	products.category, 
	cart_items.product_id,   
	products.product_name, 
	products.price,
	cart_items.quantity
	FROM cart_items 
	INNER JOIN products 
	ON cart_items.product_id = products.id`

	rows, err := c.db.Query(sqlStatement)
	if err != nil {
		return cartItems, err
	}

	for rows.Next() {
		var cartItem CartItem
		err := rows.Scan(
			&cartItem.ID,
			&cartItem.Category,
			&cartItem.ProductID,
			&cartItem.ProductName,
			&cartItem.Price,
			&cartItem.Quantity,
		)
		if err != nil {
			return cartItems, err
		}
		cartItems = append(cartItems, cartItem)
	}

	return cartItems, nil

}

func (c *CartItemRepository) FetchCartByProductID(productID int64) (CartItem, error) {
	var cartItem CartItem
	//var sqlStatement string
	//TODO : you must fetch the cart by product id
	//HINT : you can use the where statement
	sqlStatement := `SELECT 
	cart_items.id,
	products.category,
	cart_items.product_id,  
	products.product_name, 
	products.price,
	cart_items.quantity
	FROM cart_items 
	INNER JOIN products 
	ON cart_items.product_id = products.id
	WHERE cart_items.product_id = ?`

	row := c.db.QueryRow(sqlStatement, productID)
	err := row.Scan(
		&cartItem.ID,
		&cartItem.Category,
		&cartItem.ProductID,
		&cartItem.ProductName,
		&cartItem.Price,
		&cartItem.Quantity,
	)
	if err != nil {
		return cartItem, err
	}

	return cartItem, nil

}

func (c *CartItemRepository) InsertCartItem(cartItem CartItem) error {
	// TODO: you must insert the cart item

	sqlStatement := `INSERT INTO cart_items (product_id, quantity) 
	VALUES (?, ?)`
	_, err := c.db.Exec(sqlStatement, cartItem.ProductID, cartItem.Quantity)
	if err != nil {
		return err
	}
	return nil

}

func (c *CartItemRepository) IncrementCartItemQuantity(cartItem CartItem) error {
	//TODO : you must update the quantity of the cart item
	//HINT : you can use the where statement
	sqlStatement := `UPDATE cart_items SET quantity = quantity + 1 
	WHERE product_id = ?`
	_, err := c.db.Exec(sqlStatement, cartItem.ProductID, cartItem.Quantity)
	if err != nil {
		return err
	}

	return nil
}

func (c *CartItemRepository) ResetCartItems() error {
	//TODO : you must reset the cart items
	//HINT : you can use the delete statement
	sqlStatement := `DELETE FROM cart_items`
	_, err := c.db.Exec(sqlStatement)
	if err != nil {
		return err
	}

	return nil
}

func (c *CartItemRepository) TotalPrice() (int, error) {
	//var sqlStatement string
	//TODO : you must calculate the total price of the cart items
	//HINT : you can use the sum statement

	var totalPrice int
	sqlStatement := `SELECT SUM(products.price * cart_items.quantity) 
	FROM cart_items 
	INNER JOIN products ON cart_items.product_id = products.id`
	err := c.db.QueryRow(sqlStatement).Scan(&totalPrice)
	if err != nil {
		return totalPrice, err
	}

	return totalPrice, nil
}
