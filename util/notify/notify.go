package notify

type Notify struct {
	c        chan bool
	size     int
	notified bool
}

func NewNotify() *Notify {
	size := 32
	c := make(chan bool, size)
	notify := Notify{c, size, false}
	return &notify
}

func NewNotifyWithSize(size int) *Notify {
	c := make(chan bool, size)
	notify := Notify{c, size, false}
	return &notify
}

func (n Notify) Check() bool {
	select {
	case <-n.c:
		return true
	default:
		return false
	}
}

func (n Notify) Notify() {
	if n.notified {
		return
	}
	n.notified = true
	i := 0
	for i < n.size {
		n.c <- true
		i++
	}
}

func (n Notify) Wait() {
	<-n.c
}
