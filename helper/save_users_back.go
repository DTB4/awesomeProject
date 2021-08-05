package helper

import (
	"awesomeProject/model"
	"encoding/json"
	"fmt"
	"os"
)

func SaveUsersBack(users *[]model.User) error {
	file, err := os.OpenFile("./datastore/user.txt", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	bytes, err := json.MarshalIndent(&users, "", "")
	if err != nil {
		return err
	}
	length, err := file.Write(bytes)
	if err != nil {
		return err
	}
	fmt.Println(length, "bytes was written")
	defer file.Close()
	return nil
}
