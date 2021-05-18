package pubsub

// subject interface
type subject interface {
	Register(Observer Observer)
	Deregister(Observer Observer)
	notifyAll()
}

// BaseSubject struct
type BaseSubject struct {
	subject
	ObserverList []Observer
	Name         string
}

// Register add Observer to subject
func (b *BaseSubject) Register(o Observer) {
	b.ObserverList = append(b.ObserverList, o)
}

// Deregister remove Observer from subject
func (b *BaseSubject) Deregister(o Observer) {
	b.ObserverList = removeFromSlice(b.ObserverList, o)
}

// notifyAll notify all observers
func (b *BaseSubject) notifyAll() {
	for _, observer := range b.ObserverList {
		observer.InformObserver(b.Name)
	}
}

// removeFromSlice help func
func removeFromSlice(observerList []Observer, observerToRemove Observer) []Observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.GetObserverID() == observer.GetObserverID() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}