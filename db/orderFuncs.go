package db

import "EJawTest/models"

func GetOrderFromDB(orderID uint) (*models.OrderStruct, error) {
	order := &models.OrderStruct{}
	err := db.QueryRow("SELECT id, date FROM orders WHERE id = $1", orderID).Scan(&order.Order.ID, &order.Order.Date)
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT op.id, op.order_id, op.product_id, op.amount, p.name, p.description, p.serial_number, p.seller_id FROM order_product op JOIN product p ON op.product_id = p.id WHERE op.order_id = $1", orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orderProducts []models.OrderProduct
	for rows.Next() {
		var op models.OrderProduct
		var product models.Product
		err = rows.Scan(&op.ID, &op.OrderID, &op.ProductID, &op.Amount, &product.Name, &product.Description, &product.SerialNumber, &product.SellerID)
		if err != nil {
			return nil, err
		}
		op.Product = product
		orderProducts = append(orderProducts, op)
	}

	order.Items = orderProducts

	return order, nil
}

func DeleteOrderFromDB(orderID uint) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("DELETE FROM order_product WHERE order_id = $1", orderID)
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM orders WHERE id = $1", orderID)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func CreateOrderInDB(order models.Order, items []models.OrderProduct) (*models.OrderStruct, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	err = tx.QueryRow("INSERT INTO orders (date) VALUES ($1) RETURNING id, date", order.Date).Scan(&order.ID, &order.Date)
	if err != nil {
		return nil, err
	}

	for _, item := range items {
		err = tx.QueryRow("INSERT INTO order_product (order_id, product_id, amount) VALUES ($1, $2, $3) RETURNING id", order.ID, item.ProductID, item.Amount).Scan(&item.ID)
		if err != nil {
			return nil, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	orderStruct := &models.OrderStruct{
		Order: order,
		Items: items,
	}

	return orderStruct, nil
}

func UpdateOrderInDB(orderStruct models.OrderStruct) (*models.OrderStruct, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	err = tx.QueryRow("UPDATE orders SET date = $1 WHERE id = $2 RETURNING id, date", orderStruct.Order.Date, orderStruct.Order.ID).Scan(&orderStruct.Order.ID, &orderStruct.Order.Date)
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec("DELETE FROM order_product WHERE order_id = $1", orderStruct.Order.ID)
	if err != nil {
		return nil, err
	}

	for _, item := range orderStruct.Items {
		err = tx.QueryRow("INSERT INTO order_product (order_id, product_id, amount) VALUES ($1, $2, $3) RETURNING id", orderStruct.Order.ID, item.ProductID, item.Amount).Scan(&item.ID)
		if err != nil {
			return nil, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &orderStruct, nil
}

func GetAllOrdersFromDB(page, pageSize int) ([]models.OrderStruct, error) {
	var orders []models.OrderStruct
	offset := (page - 1) * pageSize

	rows, err := db.Query("SELECT id, date FROM orders LIMIT $1 OFFSET $2", pageSize, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		order := models.OrderStruct{}
		err := rows.Scan(&order.Order.ID, &order.Order.Date)
		if err != nil {
			return nil, err
		}

		productRows, err := db.Query("SELECT op.id, op.order_id, op.product_id, op.amount, p.name, p.description, p.serial_number, p.seller_id FROM order_product op JOIN product p ON op.product_id = p.id WHERE op.order_id = $1", order.Order.ID)
		if err != nil {
			return nil, err
		}
		defer productRows.Close()

		var orderProducts []models.OrderProduct
		for productRows.Next() {
			var op models.OrderProduct
			var product models.Product
			err = productRows.Scan(&op.ID, &op.OrderID, &op.ProductID, &op.Amount, &product.Name, &product.Description, &product.SerialNumber, &product.SellerID)
			if err != nil {
				return nil, err
			}
			op.Product = product
			orderProducts = append(orderProducts, op)
		}

		order.Items = orderProducts
		orders = append(orders, order)
	}

	return orders, nil
}
