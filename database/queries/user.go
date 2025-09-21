package queries

import "fmt"

func GetUserByEmailQuery(email string) string {
	query := `select id, email, createdAt from users where email = '%s'`
	return fmt.Sprintf(query, email)
}

func InsertUser(email string, password string) string {
	query := "insert into users (email, password) values ('%s', '%s')"
	return fmt.Sprintf(query, email, password)
}
