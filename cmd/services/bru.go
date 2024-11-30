package services

import (
	"context"
	"time"
)

// defining struct for the service
type Coffee struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Roast     string    `json:"roast"`
	Image     string    `json:"image"`
	Region    string    `json:"region"`
	Price     float32   `json:"price"`
	GrindUnit int16     `json:"grind_unit"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// to define the query and the data going in the query
func (c *Coffee) GetAllCoffees() ([]*Coffee, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)

	defer cancel()

	query := `SELECT id, name , image , roast , region , price , grind_unit, created_at , updated_at FROM coffees`
	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	var coffees []*Coffee

	// PREPARES THE NEXT ROW TO READ ANOTHER AGAIN
	for rows.Next() {
		var coffee Coffee
		err := rows.Scan(
			&coffee.ID, // this order should be the same as the query order
			&coffee.Name,
			&coffee.Image,
			&coffee.Roast,
			&coffee.Region,
			&coffee.Price,
			&coffee.GrindUnit,
			&coffee.CreatedAt,
			&coffee.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		coffees = append(coffees, &coffee)

	}

	return coffees, nil

}

// for creating the service with json params
func (c *Coffee) CreateCoffee(coffee Coffee) (*Coffee, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)

	defer cancel()

	query := `
		    INSERT INTO coffees (name, image, region, roast, price, grind_unit, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING *`
	// check
	_, err := db.ExecContext(
		ctx,
		query,
		// this order should be the same as the query order
		coffee.Name,
		coffee.Image,
		coffee.Roast,
		coffee.Region,
		coffee.Price,
		coffee.GrindUnit,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return nil, err
	}

	return &coffee, nil
}
