package observer


type observer interface {
	update(string)
	getID() string
	verify() error
}

type subject interface {
	register(Observer observer)
	deregister(Observer observer)
	notifyAll()
}