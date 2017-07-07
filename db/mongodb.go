package db

import "strconv"

type MongoDB struct {
	table string
}

//set a table name to perform action
func (m *MongoDB)Table(name string) {
	m.table = name
}

//perform your raw query
func (m MongoDB)Query(q string) string {
	return "MongoDB>> Query on -> " + m.table + " " + q
}

//find a record using it's primary id
func (m MongoDB)Find(id int) string {
	return "MongoDB>> SELECT * FROM " + m.table + " WHERE id=" + strconv.Itoa(id)
}

//fetch all the records from the table
func (m MongoDB)Get() string {
	return "MongoDB>> SELECT * FROM " + m.table
}

