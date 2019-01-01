package main

func Or(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0: // when there are no channels to handle
		return nil
	case 1: // when there is only one channel in the list
		return channels[0]
	}

	orDone := make(chan interface{})
	go func() {
		defer close(orDone)

		switch len(channels) {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			case <-Or(append(channels[3:], orDone)...):
				// here, recursively call or() to after the 3rd channel
			}
		}
	}()
	return orDone
}
