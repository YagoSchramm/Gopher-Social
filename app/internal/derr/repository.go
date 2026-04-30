package derr

var NotFound = NewRepositoryError("NOT_FOUND", "resource not found")
var Conflict = NewClientError("CONFLICT", "resource already exists")
