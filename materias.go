package main

import (
	"database/sql"
	"errors"
	"fmt"
)

type Materia struct {
	IdMateria       int
	Descrip         string
	CantidadAlumnos int
}

func CrearMateria(m Materia) error {
	if m.Descrip == "" {
		return errors.New("error: descrip no puede ser nulo")
	}
	q := `INSERT INTO materias(descrip, cantidad_alumnos)
			VALUES ($1,$2)`
	db := GetConnection()
	defer db.Close()

	//Preparando datos nulos
	cantidadNull := sql.NullInt32{}

	if m.CantidadAlumnos == 0 {
		cantidadNull.Valid = false
	} else {
		cantidadNull.Valid = true
		cantidadNull.Int32 = int32(m.CantidadAlumnos)
	}

	//Preparando sentencia
	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	r, err := stmt.Exec(&m.Descrip, &cantidadNull)
	if err != nil {
		return err
	}
	i, _ := r.RowsAffected()
	if i != 1 {
		return errors.New("error: mas de una fila afectada")
	}
	fmt.Println("Materia creada correctamente!")
	return nil
}

func LeerMaterias() ([]Materia, error) {
	q := `SELECT * FROM materias`
	db := GetConnection()
	defer db.Close()

	cantidadNull := sql.NullInt32{}

	stmt, err := db.Query(q)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var materias []Materia
	for stmt.Next() {
		m := Materia{}
		err := stmt.Scan(&m.IdMateria, &m.Descrip, &cantidadNull)
		if err != nil {
			return nil, err
		}

		m.CantidadAlumnos = int(cantidadNull.Int32)

		materias = append(materias, m)
	}
	return materias, nil
}

func BorrarMateria(id int) error {
	q := `DELETE FROM materias WHERE id_materia = $1`
	db := GetConnection()
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	rows, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	i, _ := rows.RowsAffected()
	if i != 1 {
		return errors.New("error: más de una fila modificada")
	}
	fmt.Println("Borrado correctamente")
	return nil
}

func ActualizarMateria(e Materia) error {
	var q string
	q = `UPDATE materias
	SET descrip = $1, cantidad_alumnos = $3
	WHERE id_materia = $2`

	db := GetConnection()
	defer db.Close()

	if e.CantidadAlumnos == 0 {
		q = `UPDATE materias
		SET descrip = $1
		WHERE id_materia = $2`
	}

	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	if e.CantidadAlumnos == 0 {
		rows, err := stmt.Exec(e.Descrip, e.IdMateria)
		if err != nil {
			return err
		}
		i, _ := rows.RowsAffected()
		if i != 1 {
			fmt.Println(i)
			return errors.New("error: más de una fila modificada")
		}

	} else {
		rows, err := stmt.Exec(e.Descrip, e.IdMateria, e.CantidadAlumnos)
		if err != nil {
			return err
		}
		i, _ := rows.RowsAffected()
		if i != 1 {
			fmt.Println(i)
			return errors.New("error: más de una fila modificada")
		}

	}
	fmt.Println("Actualizado correctamente")
	return nil
}
