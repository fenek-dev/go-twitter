package tweets

import "github.com/jackc/pgx/v5"

type Repository struct {
	conn *pgx.Conn
}

func NewRepository(conn *pgx.Conn) *Repository {
	return &Repository{
		conn: conn,
	}
}

func (r *Repository) GetById() {

}

func (r *Repository) Create() {

}

func (r *Repository) Delete() {

}
