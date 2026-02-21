package employee

import (
	"context"
	"database/sql"
	"log"
)

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) Repository {
	return &postgresRepository{db: db}
}

func (r postgresRepository) GetById(ctx context.Context, id int) (*Employee, error) {
	query := "SELECT name, email, salary, active FROM employee WHERE id = $1"

	row := r.db.QueryRow(query, id)

	var name string
	var email string
	var salary float64
	var active bool

	err := row.Scan(&name, &email, &salary, &active)
	if err != nil {
		return &Employee{}, err
	}

	return &Employee{
		ID: id,
		Name: name,
		Email: email,
		Salary: salary,
		Active: active,
	}, nil
}


func (r postgresRepository) GetAll(ctx context.Context) ([]Employee, error) {
	query := "SELECT id, name, email, salary, active FROM employee"

	emps, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer emps.Close()

	data := []Employee{}

	var name string
	var email string
	var salary float64
	var active bool
	var id int

	for emps.Next() {
		err := emps.Scan(&id, &name, &email, &salary, &active)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, Employee{ID: id,Name: name, Email: email, Salary: salary, Active: active})
	}

	return data, nil
}

func (r *postgresRepository) RegisterEmployee(ctx context.Context, emp *Employee) (int, error){
	query :=`INSERT INTO employee (name, email, salary, active) 
			VALUES ($1, $2, $3, $4) RETURNING id;`

	var pk int

	err := r.db.QueryRow(query, emp.Name, emp.Email, emp.Salary, emp.Active).Scan(&pk)
	if err != nil {
		return 0, err
	}

	return pk, nil
}

func (r *postgresRepository) Fire(ctx context.Context, id int) (*Employee, error) {
	query := "UPDATE employee SET active = false WHERE id = $1"

	_, err := r.db.Exec(query, id)
	if err != nil {
		return &Employee{}, err
	}

	emp, err := r.GetById(ctx, id)
	if err != nil {
		return &Employee{}, err
	}
	
	return emp, nil 
}

func (r *postgresRepository) Employ(ctx context.Context, id int) (*Employee, error) {
	query := "UPDATE employee SET active = true WHERE id = $1"

	_, err := r.db.Exec(query, id)
	if err != nil {
		return &Employee{}, err
	}

	emp, err := r.GetById(ctx, id)
	if err != nil {
		return &Employee{}, err
	}
	
	return emp, nil 
}