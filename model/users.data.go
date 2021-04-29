package model

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Id        int    `json:"id" `
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form:"last_name"`
	Gender    string `json:"gender" form:"gender"`
	IpAddress string `json:"ip_address" form:"ip_address"`
}

func GetallUsers() ([]User, error) {
	users := make([]User, 0)
	file, errfile := os.Open("MOCK_DATA.json")
	if errfile != nil {
		return nil, fmt.Errorf("%s", errfile)
	}
	defer file.Close()
	//fileByte, _ := io.ReadAll(file)

	err := json.NewDecoder(file).Decode(&users)
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}
	//fmt.Println(users)
	return users, nil

}

func check(err error) {
	if err != nil {
		fmt.Errorf("%w", err)
	}
}

func NextID(u *User, us []User) []User {
	id := len(us) + 1
	u.Id = id
	us = append(us, *u)
	return us

}
func FindUserById(id int, users []User) (*User, int, error) {
	//users, err := GetallUsers()
	for _, user := range users {
		if user.Id == id {
			return &user, id, nil
		}

	}
	return nil, 0, fmt.Errorf("User With id %d not Found", id)
}
