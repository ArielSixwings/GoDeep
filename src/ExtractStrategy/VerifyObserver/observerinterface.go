package observer


type observer interface {
	getID() string
	verify() error
}