package models

import (
	"errors"
	"time"
)

type Estudiante struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func Crear(e Estudiante) error {
	q := `INSERT INTO
			estudiantes (name, age, active)
			VALUES (?, ?, ?)`

	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(q)

	if err != nil {
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(e.Name, e.Age, e.Active)

	if err != nil {
		return err
	}

	i, _ := r.RowsAffected()

	if i != 1 {
		return errors.New("Se esperava uma linha afetada")
	}

	return nil
}

func Consultar() (estudiantes []Estudiante, err error) {
	q := `SELECT id, name, age, active, created_at, updated_at
							FROM estudiantes`
	db := getConnection()
	defer db.Close()

	rows, err := db.Query(q)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		e := Estudiante{}

		err = rows.Scan(&e.ID,
			&e.Name,
			&e.Age,
			&e.Active,
			&e.CreatedAt,
			&e.UpdatedAt,
		)
		if err != nil {
			return
		}
		estudiantes = append(estudiantes, e)
	}
	return estudiantes, nil
}

func Actualizar(e Estudiante) error {
	q := `UPDATE estudiantes
			SET name = ?, age = ?, active = ?, updated_at = now()
			WHERE id = ?`

	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(q)

	if err != nil {
		return err
	}

	defer stmt.Close()

	r, err := stmt.Exec(e.Name, e.Age, e.Active, e.ID)

	if err != nil {
		return err
	}

	i, _ := r.RowsAffected()

	if i != 1 {
		return errors.New("Se esperava uma linha afetada")
	}

	return nil

}

func Deletar(ID int) error {
	q := `DELETE FROM estudiantes WHERE id = ?`

	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(q)

	if err != nil {
		return err
	}

	defer stmt.Close()

	r, err := stmt.Exec(ID)

	if err != nil {
		return err
	}

	i, _ := r.RowsAffected()

	if i != 1 {
		return errors.New("Se esperava uma linha afetada")
	}

	return nil
}

func Pesquisa(termo string) (estudiantes []Estudiante, err error) {
	q := `SELECT * FROM estudiantes WHERE name like ?;`

	db := getConnection()
	defer db.Close()

	rows, err := db.Query(q, "%"+termo+"%")
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		e := Estudiante{}

		err = rows.Scan(&e.ID,
			&e.Name,
			&e.Age,
			&e.Active,
			&e.CreatedAt,
			&e.UpdatedAt,
		)
		if err != nil {
			return
		}
		estudiantes = append(estudiantes, e)
	}
	return estudiantes, nil
}

func ConsultaParaUm(id int) (estudiante Estudiante, err error) {
	q := `SELECT * FROM estudiantes WHERE id = ?;`

	db := getConnection()
	defer db.Close()

	e := Estudiante{}

	err = db.QueryRow(q, id).Scan(&e.ID,
		&e.Name,
		&e.Age,
		&e.Active,
		&e.CreatedAt,
		&e.UpdatedAt)
	if err != nil {
		return
	}

	return e, nil
}
