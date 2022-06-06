package repository

import "database/sql"

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (p *ProductRepository) FetchProductByID(id int64) (Product, error) {
	//TODO: You must implement this function fot fetch product by id
	sqlStatement := `SELECT * FROM products WHERE id = ?`
	var product Product
	row := p.db.QueryRow(sqlStatement, id)
	err := row.Scan(&product.ID, &product.ProductName, &product.Category, &product.Price, &product.Quantity)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (p *ProductRepository) FetchProductByName(productName string) (Product, error) {
	// TODO: You must implement this function for fetch product by name
	sqlStatement := `SELECT * FROM products WHERE product_name = ?`
	var product Product
	row := p.db.QueryRow(sqlStatement, productName)
	err := row.Scan(&product.ID, &product.ProductName, &product.Category, &product.Price, &product.Quantity)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (p *ProductRepository) FetchProducts() ([]Product, error) {
	// TODO: You must implement this function for fetch all products
	sqlStatement := `SELECT * FROM products`
	var products []Product
	rows, err := p.db.Query(sqlStatement)
	if err != nil {
		return products, err
	}
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.ProductName, &product.Category, &product.Price, &product.Quantity)
		if err != nil {
			return products, err
		}
		products = append(products, product)
	}
	return products, nil
}
