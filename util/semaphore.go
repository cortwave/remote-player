package util

type empty struct{}
type semaphore chan empty

//Acquire n resources
func (s semaphore) Acquire(n int) {
	e := empty{}
	for i := 0; i < n; i++ {
		s <- e
	}
}

//Release n resources
func (s semaphore) Release(n int) {
	for i := 0; i < n; i++ {
		<-s
	}
}

//Lock semaphore
func (s semaphore) Lock() {
	s.Release(1)
}

//Unlock semaphore
func (s semaphore) Unlock() {
	s.Acquire(1)
}
