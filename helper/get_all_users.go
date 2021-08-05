package helper

import (
	"awesomeProject/model"
	"encoding/json"
	"io/ioutil"
)

func GetAllUsers() (*[]model.User, error) {
	dataFromFile, err:=ioutil.ReadFile("./datastore/user.txt")
	if err !=nil {
		return nil, err
	}
	users := new([] model.User)
	err =json.Unmarshal(dataFromFile, &users)
	return users, nil
}
