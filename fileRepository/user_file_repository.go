package fileRepository

import (
	"awesomeProject/helper"
	"awesomeProject/model"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
	"time"
)


type UserRepositoryI interface {
	Create (u *model.User) (*model.User, error)
	Get (email *string, id *int32) *model.User
	GetAll() []*model.User
	FakeDelete (id int32) error
	Delete (id int32) error
	Edit (u model.User) *model.User
}

type UserFileRepository struct {
	idMutex *sync.Mutex
}

func NewUserFileRepository () *UserFileRepository{
	return &UserFileRepository{
		idMutex: &sync.Mutex{},
	}
}

func (ufr *UserFileRepository) Create(user *model.User) error {
	users, err := helper.GetAllUsers()
	if err !=nil {
		return err
	}
	user.ID, _ = ufr.GetNextID()
	users1:=append(*users, *user)
	file, err:=os.OpenFile("./datastore/user.txt", os.O_WRONLY|os.O_CREATE, 0600)
	if err !=nil {
		return  err
	}
	bytes, err:=json.MarshalIndent(&users1, "", "")
		if err !=nil {
			return  err
		}
	length, err := file.Write(bytes)
	if err !=nil {
		return err
	}
	fmt.Println("New user ", user, " was saved and ", length, "bytes was written")
	defer file.Close()
	return nil
	}

func (ufr *UserFileRepository) Get(email *string, id *int32) (*model.User, error) {
	users, err := helper.GetAllUsers()
	if err !=nil {
		return nil, err
	}
	for _, value:=range *users{
		if value.Email==*email||value.ID==*id{
			return &value, nil
		}
	}
	return nil ,errors.New("user with this ID/e-mail not found")
}

func (ufr *UserFileRepository) GetAll() (*[]model.User, error) {
	users, err := helper.GetAllUsers()
	if err !=nil {
		return nil, err
	}
	return users, nil
}

func (ufr *UserFileRepository) Delete(id int32) error {
	users, err := helper.GetAllUsers()
	if err !=nil {
		return err
	}
	users1:=*users
	for i, value:=range users1{

		if value.ID==id{
			value.Deleted=time.Now().String()
			users1[i].Deleted=time.Now().String()

			file, err:=os.OpenFile("./datastore/user.txt", os.O_WRONLY|os.O_CREATE, 0600)
			if err !=nil {
				return  err
			}
			bytes, err:=json.MarshalIndent(&users1, "", "")
			if err !=nil {
				return  err
			}
			length, err := file.Write(bytes)
			if err !=nil {
				return err
			}
			fmt.Println("User  ", users1[i].ID, " was deleted ", length, "bytes was written")
			defer file.Close()
			return nil

		}
		i++
	}
	return errors.New("User not found")
}


func (ufr *UserFileRepository) Edit(user model.User) *model.User {
	panic("implement me")
}

func (ufr *UserFileRepository) GetNextID() (int32, error) {
	ufr.idMutex.Lock()
	b1, err:=ioutil.ReadFile("./datastore/user_ID.txt")
	if err !=nil {
		panic(err)
	}
	fID:=FileID{}
	err =json.Unmarshal(b1, &fID)
	if err !=nil {
		panic(err)
	}
	fID.ID+=1
	f, err:=os.OpenFile("./datastore/user_ID.txt", os.O_WRONLY, 0600)
	b1, err=json.Marshal(fID)
	n, err := f.Write(b1)
	fmt.Println("new ID ", fID, " was generated and",n, " bytes was written")
	err=f.Close()
	ufr.idMutex.Unlock()
	return fID.ID,nil
}

type FileID struct {
	ID int32 `json:"id"`
}