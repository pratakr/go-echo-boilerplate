package main

import (
	"fmt"
	"os"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "user:password@tcp(localhost:3306)/database_name?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Get table information from the database
	rows, err := db.Raw("DESCRIBE table_name").Rows()
	if err != nil {
		panic("failed to get table information")
	}
	defer rows.Close()

	// Create struct fields based on the column names and types
	var fields []string
	for rows.Next() {
		var field, fieldType, _, _, _, _ string
		if err := rows.Scan(&field, &fieldType, _, _, _, _); err != nil {
			panic("failed to scan row")
		}
		goType := sqlTypeToGoType(fieldType)
		fields = append(fields, fmt.Sprintf("%s %s", field, goType))
	}

	// Generate struct based on the table name and fields
	structName := "TableName"
	structDefinition := fmt.Sprintf("type %s struct {\n\t%s\n}", structName, strings.Join(fields, "\n\t"))

	// Write struct to file
	filename := "/path/to/struct.go"
	file, err := os.Create(filename)
	if err != nil {
		panic("failed to create file")
	}
	defer file.Close()

	_, err = file.WriteString(structDefinition)
	if err != nil {
		panic("failed to write to file")
	}

	fmt.Printf("Struct definition written to %s\n", filename)
}

// Helper function to convert SQL types to Golang types
func sqlTypeToGoType(sqlType string) string {
	switch sqlType {
	case "bigint":
		return "int64"
	case "int":
		return "int32"
	case "varchar", "text":
		return "string"
	case "datetime":
		return "time.Time"
	default:
		return "interface{}"
	}
}
