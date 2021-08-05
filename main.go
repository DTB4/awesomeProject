package main

import (
	"awesomeProject/fileRepository"
	"awesomeProject/model"
)

//func handler (w http.ResponseWriter, r *http.Request)  {
//	if r.URL.Path=="/api/v1"{
//		fmt.Fprintf(w, "This must be nginx server default \n My request url is %s", r.URL.Path)
//	}
//	if r.URL.Path=="/api/v2" {
//		fmt.Fprintf(w, "This must be nginx server mysite \n My request url is %s", r.URL.Path)
//	}
//}
func main() {
	userRepository := fileRepository.NewUserFileRepository()

	//u:=model.User{
	//	ID: 9,
	//	Email: "petr9@gm.com",
	//}
	//findedUser, err:= userRepository.Get(&u.Email, &u.ID)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(findedUser, "we found this")


	u := model.User{
		ID: 9,
	}
	userRepository.Delete(u.ID)



	//allUsers, err := userRepository.GetAll()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println("Start of Users list")
	//for _, value:=range *allUsers{
	//	fmt.Println(value)
	//}
	//fmt.Println("End of Users list")




	//u := model.User{
	//	ID:          0,
	//	Name:        "Petya",
	//	Email:       "petr@gm.com",
	//	Password:    "pasword123",
	//	Location:    "home19283012x120830131",
	//	PhoneNumber: "05099922213",
	//	Deleted: "",
	//}
	//err := userRepository.Create(&u)
	//if err !=nil {
	//	fmt.Println(err.Error())
	//	return
	//}

//	fmt.Println("Server is listening for http//")
//	http.HandleFunc("/api/", handler)
//log.Fatal(http.ListenAndServe(":8080",nil))
}
