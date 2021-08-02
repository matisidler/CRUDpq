package main

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/lib/pq"
)

//Creando estructura estudiantes
type Estudiante struct {
	Id                 uint
	Name               string
	Age                uint
	Active             bool
	CreatedAt          time.Time
	UpdatedAt          time.Time
	MateriasInscriptas uint
}

//Creando estudiantes en la BD

func CrearEstudiante(e Estudiante) error {
	q := `INSERT INTO estudiantes(name, age, active, materias_inscriptas)
			VALUES($1, $2, $3, $4)`
	db := GetConnection()
	defer db.Close()

	//Admitiendo valores nulos
	nameNull := sql.NullString{}
	ageNull := sql.NullInt32{}

	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	//Aca le digo que si e.Name es nulo, me active nameNull. En caso de que no sea nulo, lo desactivo y su valor va a ser el de e.Name
	if e.Name == "" {
		nameNull.Valid = false
	} else {
		nameNull.Valid = true
		nameNull.String = e.Name
	}

	if e.Age == 0 {
		ageNull.Valid = false
	} else {
		ageNull.Valid = true
		ageNull.Int32 = int32(e.Age)
	}

	//Acá le paso el valor de nameNull, que si name es nulo va a ser nulo, si name tiene valor, va a mostrar ese valor.
	rows, err := stmt.Exec(nameNull, ageNull, e.Active, e.MateriasInscriptas)
	if err != nil {
		return err
	}
	i, _ := rows.RowsAffected()
	if i != 1 {
		return errors.New("error: más de una fila modificada")
	}
	fmt.Println("Creado correctamente")
	return nil
}

func BorrarEstudiante(id int) error {
	q := `DELETE FROM estudiantes WHERE id = $1`
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

func ActualizarEstudiante(e Estudiante) error {
	q := `UPDATE estudiantes
			SET name = $1, age = $2, active = $3, updated_at = NOW(), materias_inscriptas = $4
			WHERE id = $5`
	db := GetConnection()
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	rows, err := stmt.Exec(e.Name, e.Age, e.Active, e.MateriasInscriptas, e.Id)
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

func LeerEstudiantes() ([]Estudiante, error) {
	q := `SELECT * FROM estudiantes`
	db := GetConnection()
	defer db.Close()

	nameNull := sql.NullString{}
	ageNull := sql.NullInt32{}
	materiasNull := sql.NullInt32{}
	activeNull := sql.NullBool{}
	fechaNull := pq.NullTime{}

	stmt, err := db.Query(q)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var estudiantes []Estudiante
	for stmt.Next() {
		e := Estudiante{}
		err := stmt.Scan(&e.Id, &nameNull, &ageNull, &activeNull, &e.CreatedAt, &fechaNull, &materiasNull)
		if err != nil {
			return nil, err
		}
		e.Name = nameNull.String
		e.Age = uint(ageNull.Int32)
		e.MateriasInscriptas = uint(materiasNull.Int32)
		e.Active = activeNull.Bool
		e.UpdatedAt = fechaNull.Time

		estudiantes = append(estudiantes, e)
	}
	return estudiantes, nil
}
