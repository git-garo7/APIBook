package controllers

import "net/http"

//CreateUser = insert user into database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create user"))
}

//SearchUsers = search for all users saved in the database
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("looking for all users"))
}

//SearchUser = search for a user in the database
func SearchUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("looking one user"))
}

//UpdateUser = update user data
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("updating user"))
}

//DeleteUser = delete database user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("deleting user"))
}
