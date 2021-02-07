package shared

// DomainError is an error that represent a violation of certain bussines rules
type DomainError interface {
	Message() string
}
