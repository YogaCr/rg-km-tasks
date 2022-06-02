package repository

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type CartItemRepository struct {
	db *sql.DB
}

func NewCartItemRepository(db *sql.DB) *CartItemRepository {
	return &CartItemRepository{db: db}
}

func (c *CartItemRepository) FetchCartItems() ([]CartItem, error) {
	var sqlStatement string
	var cartItems []CartItem

	//TODO: add sql statement here
	//HINT: join table cart_items and products

	sqlStatement = "SELECT cart_items.id as id, cart_items.quantity as quantity, product_id, products.category as category, products.product_name as product_name, products.price as price from cart_items join products on products.id = cart_items.product_id"
	res, err := c.db.Query(sqlStatement)
	if err != nil {
		return cartItems, err
	}
	defer res.Close()

	for res.Next() {
		item := &CartItem{}
		errScan := res.Scan(&item.ID, &item.Quantity, &item.ProductID, &item.Category, &item.ProductName, &item.Price)
		if errScan != nil {
			return cartItems, errScan
		}
		cartItems = append(cartItems, *item)
	}

	return cartItems, nil // TODO: replace this
}

func (c *CartItemRepository) FetchCartByProductID(productID int64) (CartItem, error) {
	//TODO : you must fetch the cart by product id
	//HINT : you can use the where statement

	sqlStatement := "SELECT cart_items.id as id, cart_items.quantity as quantity, product_id, products.category as category, products.product_name as product_name, products.price as price from cart_items join products on products.id = cart_items.product_id WHERE product_id = ?"
	row := c.db.QueryRow(sqlStatement, productID)

	cartItem := &CartItem{}
	err := row.Scan(&cartItem.ID, &cartItem.Quantity, &cartItem.ProductID, &cartItem.Category, &cartItem.ProductName, &cartItem.Price)

	return *cartItem, err // TODO: replace this
}

func (c *CartItemRepository) InsertCartItem(cartItem CartItem) error {
	// TODO: you must insert the cart item

	sqlStatement := "INSERT INTO cart_items (product_id, quantity) VALUES (?,?)"
	_, err := c.db.Exec(sqlStatement, cartItem.ProductID, cartItem.Quantity)
	return err // TODO: replace this
}

func (c *CartItemRepository) IncrementCartItemQuantity(cartItem CartItem) error {
	//TODO : you must update the quantity of the cart item
	sqlStatement := "UPDATE cart_items SET quantity = ? WHERE id = ?"
	_, err := c.db.Exec(sqlStatement, cartItem.Quantity+1, cartItem.ID)
	if err != nil {
		log.Fatal(err)
	}
	return err // TODO: replace this
}

func (c *CartItemRepository) ResetCartItems() error {
	//TODO : you must reset the cart items
	//HINT : you can use the delete statement

	sqlStatement := "DELETE FROM cart_items"
	_, err := c.db.Exec(sqlStatement)
	return err // TODO: replace this
}

func (c *CartItemRepository) TotalPrice() (int, error) {
	var sqlStatement string
	//TODO : you must calculate the total price of the cart items
	//HINT : you can use the sum statement

	total := 0

	sqlStatement = "SELECT sum(cart_items.quantity * products.price) FROM cart_items join products on products.id = cart_items.product_id"
	row := c.db.QueryRow(sqlStatement)
	row.Scan(&total)
	return total, nil // TODO: replace this
}
