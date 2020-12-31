package errors

//used to return business service error messages
type ServiceError struct {
	Message string `json:"message"`
}
