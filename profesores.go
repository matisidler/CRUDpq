package main

import (
	"database/sql"
	"errors"
	"fmt"
)

//Creando profesores
type Profesor struct {
	Id        uint
	Nombre    string
	IdMateria uint
	Salario   uint
}

func CrearProfesor(p Profesor) error {
	if p.IdMateria == 0 {
		return errors.New("error: IdMateria no puede ser 0.")
	}
	q := `INSERT INTO profesores(nombre, id_materia, salario)
			VALUES ($1,$2,$3)`
	db := GetConnection()
	defer db.Close()

	//Preparando datos nulos
	nombreNull := sql.NullString{}
	salarioNull := sql.NullInt32{}

	if p.Nombre == "" {
		nombreNull.Valid = false
	} else {
		nombreNull.Valid = true
		nombreNull.String = p.Nombre
	}

	if p.Salario == 0 {
		salarioNull.Valid = false
	} else {
		salarioNull.Valid = true
		salarioNull.Int32 = int32(p.Salario)
	}

	//Preparando sentencia
	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	r, err := stmt.Exec(&nombreNull, &p.IdMateria, &salarioNull)
	if err != nil {
		return err
	}
	i, _ := r.RowsAffected()
	if i != 1 {
		return errors.New("error: mas de una fila afectada")
	}
	fmt.Println("Profesor creado correctamente!")
	return nil
}

func LeerProfesores() ([]Profesor, error) {
	q := `SELECT * FROM profesores`
	db := GetConnection()
	defer db.Close()

	nombreNull := sql.NullString{}
	salarioNull := sql.NullInt32{}

	stmt, err := db.Query(q)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var profesores []Profesor
	for stmt.Next() {
		p := Profesor{}
		err := stmt.Scan(&p.Id, &nombreNull, &p.IdMateria, &salarioNull)
		if err != nil {
			return nil, err
		}
		p.Nombre = nombreNull.String
		p.Salario = uint(salarioNull.Int32)

		profesores = append(profesores, p)
	}
	return profesores, nil
}

func BorrarProfesor(id int) error {
	q := `DELETE FROM profesores WHERE id_profesor = $1`
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

func ActualizarProfesor(e Profesor) error {
	q := `UPDATE profesores
			SET nombre = $1, id_materia = $2, salario = $3
			WHERE id_profesor = $4`
	db := GetConnection()
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	rows, err := stmt.Exec(e.Nombre, e.IdMateria, e.Salario, e.Id)
	if err != nil {
		return err
	}
	i, _ := rows.RowsAffected()
	if i != 1 {
		fmt.Println(i)
		return errors.New("error: más de una fila modificada")
	}
	fmt.Println("Actualizado correctamente")
	return nil
}
