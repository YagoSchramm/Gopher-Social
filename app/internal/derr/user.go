package derr

var InvalidUserName = NewClientError("INVALID_USER_NAME", "a user with that username already exists")
var InvalidUserEmail = NewClientError("INVALID_USER_EMAIL", "a user with that email already exists")
