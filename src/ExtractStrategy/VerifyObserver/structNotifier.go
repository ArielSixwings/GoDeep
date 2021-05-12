package observer

type Notifier struct {
	verificators 		[]observer
	structHadChanged 	[]bool
	structWasAllocated 	[]bool
	name 				string
}