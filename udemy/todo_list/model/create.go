package model

import "fmt"

func CreateTodo(name, todo string) error {
	insertQ, err := con.Query("INSERT INTO TODO VALUES(?, ?)", name, todo)

	defer insertQ.Close()
	if err != nil {
		fmt.Println("Error inserting data", err)
		return err
	}
	return nil
}

func DeleteTodo(name string) error {
	insertQ, err := con.Query("DELETE FROM TODO WHERE name=?", name)

	defer insertQ.Close()
	if err != nil {
		fmt.Println("Error deleting data", err)
		return err
	}
	return nil
}
