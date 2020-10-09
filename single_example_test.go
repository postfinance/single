package single

import "log"

// ExampleNew is an example how to use single
func ExampleNew() {
	// create a new lockfile in /var/lock/filename
	one, err := New("filename", WithLockPath("/tmp"))
	if err != nil {
		log.Fatal(err)
	}

	// lock and defer unlocking
	if err := one.Lock(); err != nil {
		log.Fatal(err)
	}

	// run

	if err := one.Unlock(); err != nil {
		log.Println(err)
	}
}
