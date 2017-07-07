package db

import "strconv"

type Mysql struct {
	table string
}

//set a table name to perform action
func (m *Mysql)Table(name string) {
	m.table = name
}

//perform your raw query
func (m Mysql)Query(q string) string {
	return "Mysql>> Query on -> " + m.table + " " + q
}

//find a record using it's primary id
func (m Mysql)Find(id int) string {
	return "Mysql>> SELECT * FROM " + m.table + " WHERE id=" + strconv.Itoa(id)
}

//fetch all the records from the table
func (m Mysql)Get() string {
	return "Mysql>> SELECT * FROM " + m.table
}

