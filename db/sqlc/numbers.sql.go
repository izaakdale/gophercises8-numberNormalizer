// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: numbers.sql

package db

import (
	"context"
)

const addNumber = `-- name: AddNumber :one
INSERT INTO numbers (
    number
) VALUES (
    $1
)
RETURNING id, number
`

func (q *Queries) AddNumber(ctx context.Context, number string) (Number, error) {
	row := q.db.QueryRowContext(ctx, addNumber, number)
	var i Number
	err := row.Scan(&i.ID, &i.Number)
	return i, err
}

const getNumbers = `-- name: GetNumbers :many
SELECT id, number FROM numbers
`

func (q *Queries) GetNumbers(ctx context.Context) ([]Number, error) {
	rows, err := q.db.QueryContext(ctx, getNumbers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Number{}
	for rows.Next() {
		var i Number
		if err := rows.Scan(&i.ID, &i.Number); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
