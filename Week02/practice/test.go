package sync

func Go(x func()) {
	if err := recover(); err != nil {
		// xxxx
	}
	go x()
}
