package observer

func (n *Notifier) register(o observer) {
	n.observerList = append(n.observerList, o)
}

func (n *Notifier) deregister(o observer) {
	n.observerList = removeFromslice(n.observerList, o)
}

func (n *Notifier) notifyAll() {
	for _, observer := range n.observerList {
		observer.update(i.name)
	}
}

func removeFromslice(observerList []observer, observerToRemove observer) []observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.getID() == observer.getID() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}