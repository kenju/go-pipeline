// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package pipeline

// OrInterface return only the first result from multiple channels.
// Use ctx to cancel the stream processing.
func OrInterface(
	channels ...<-chan interface{},
) <-chan interface{} {
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
			case <-OrInterface(append(channels[3:], orDone)...):
				// here, recursively call or() to after the 3rd channel
			}
		}
	}()
	return orDone
}

// OrString return only the first result from multiple channels.
// Use ctx to cancel the stream processing.
func OrString(
	channels ...<-chan string,
) <-chan string {
	switch len(channels) {
	case 0: // when there are no channels to handle
		return nil
	case 1: // when there is only one channel in the list
		return channels[0]
	}

	orDone := make(chan string)
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
			case <-OrString(append(channels[3:], orDone)...):
				// here, recursively call or() to after the 3rd channel
			}
		}
	}()
	return orDone
}

// OrInt return only the first result from multiple channels.
// Use ctx to cancel the stream processing.
func OrInt(
	channels ...<-chan int,
) <-chan int {
	switch len(channels) {
	case 0: // when there are no channels to handle
		return nil
	case 1: // when there is only one channel in the list
		return channels[0]
	}

	orDone := make(chan int)
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
			case <-OrInt(append(channels[3:], orDone)...):
				// here, recursively call or() to after the 3rd channel
			}
		}
	}()
	return orDone
}

// OrFloat32 return only the first result from multiple channels.
// Use ctx to cancel the stream processing.
func OrFloat32(
	channels ...<-chan float32,
) <-chan float32 {
	switch len(channels) {
	case 0: // when there are no channels to handle
		return nil
	case 1: // when there is only one channel in the list
		return channels[0]
	}

	orDone := make(chan float32)
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
			case <-OrFloat32(append(channels[3:], orDone)...):
				// here, recursively call or() to after the 3rd channel
			}
		}
	}()
	return orDone
}