package tests

import "go_fiber_template/app/database"

func GetTestQuery() []TestModel {
	result := []TestModel{}

	database.DBConn.Unscoped().
		Table("customer").
		Find(&result)

	return result
}
