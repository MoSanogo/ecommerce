package repository

import (
	"database/sql"
	"ecommerce-grpc-api/internal/models"
)

type ProductRepository interface {
	GetProductByID(id string) (*models.Product, error)
	CreateProduct(product *models.Product) error
	UpdateProduct(product *models.Product) error
	DeleteProduct(id string) error
	ListProducts() ([]*models.Product, error)
}

type ProductRepositoryImpl struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &ProductRepositoryImpl{db: db}
}

func (r *ProductRepositoryImpl) GetProductByID(id string) (*models.Product, error) {
	var product models.Product
	query := "SELECT id, tenant_id, name, description, price, stock_quantity, category, image_urls, created_at, updated_at FROM products WHERE id = ?"
	err := r.db.QueryRow(query, id).Scan(&product.ID, &product.TenantID, &product.Name, &product.Description, &product.Price,
		&product.StockQuantity, &product.Category, &product.ImageURLs, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepositoryImpl) CreateProduct(product *models.Product) error {
	query := "INSERT INTO products (id, tenant_id, name, description, price, stock_quantity, category, image_urls) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	_, err := r.db.Exec(query, product.ID, product.TenantID, product.Name, product.Description,
		product.Price, product.StockQuantity, product.Category, product.ImageURLs)
	if err != nil {
		return err
	}
	return nil
}

func (r *ProductRepositoryImpl) UpdateProduct(product *models.Product) error {
	query := "UPDATE products SET tenant_id = ?, name = ?, description = ?, price = ?, stock_quantity = ?, category = ?, image_urls = ? WHERE id = ?"
	_, err := r.db.Exec(query, product.TenantID, product.Name, product.Description,
		product.Price, product.StockQuantity, product.Category, product.ImageURLs, product.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *ProductRepositoryImpl) DeleteProduct(id string) error {
	query := "DELETE FROM products WHERE id = ?"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *ProductRepositoryImpl) ListProducts() ([]*models.Product, error) {
	query := "SELECT id, tenant_id, name, description, price, stock_quantity, category, image_urls, created_at, updated_at FROM products"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*models.Product
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.TenantID, &product.Name, &product.Description,
			&product.Price, &product.StockQuantity, &product.Category, &product.ImageURLs,
			&product.CreatedAt, &product.UpdatedAt); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
