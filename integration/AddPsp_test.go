package integration

import (
	"fmt"
	json "github.com/goccy/go-json"
	"testing"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func Test_AddPsp(t *testing.T) {

	jsonStr := `{"id":2,"name":"Alice","email":"alice@example.com"}`
	var user User
	err := json.Unmarshal([]byte(jsonStr), &user)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Decoded User: %+v\n", user)

	fmt.Println("!!! TEST AddPsp")

}
