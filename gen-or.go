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

// OrBool return only the first result from multiple channels.
// Use ctx to cancel the stream processing.
func OrBool(
	channels ...<-chan bool,
) <-chan bool {
	switch len(channels) {
	case 0: // when there are no channels to handle
		return nil
	case 1: // when there is only one channel in the list
		return channels[0]
	}

	orDone := make(chan bool)
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
			case <-OrBool(append(channels[3:], orDone)...):
				// here, recursively call or() to after the 3rd channel
			}
		}
	}()
	return orDone
}

// OrByte return only the first result from multiple channels.
// Use ctx to cancel the stream processing.
func OrByte(
	channels ...<-chan byte,
) <-chan byte {
	switch len(channels) {
	case 0: // when there are no channels to handle
		return nil
	case 1: // when there is only one channel in the list
		return channels[0]
	}

	orDone := make(chan byte)
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
			case <-OrByte(append(channels[3:], orDone)...):
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

// OrUint return only the first result from multiple channels.
// Use ctx to cancel the stream processing.
func OrUint(
	channels ...<-chan uint,
) <-chan uint {
	switch len(channels) {
	case 0: // when there are no channels to handle
		return nil
	case 1: // when there is only one channel in the list
		return channels[0]
	}

	orDone := make(chan uint)
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
			case <-OrUint(append(channels[3:], orDone)...):
				// here, recursively call or() to after the 3rd channel
			}
		}
	}()
	return orDone
}

// OrUint64 return only the first result from multiple channels.
// Use ctx to cancel the stream processing.
func OrUint64(
	channels ...<-chan uint64,
) <-chan uint64 {
	switch len(channels) {
	case 0: // when there are no channels to handle
		return nil
	case 1: // when there is only one channel in the list
		return channels[0]
	}

	orDone := make(chan uint64)
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
			case <-OrUint64(append(channels[3:], orDone)...):
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

// OrFloat64 return only the first result from multiple channels.
// Use ctx to cancel the stream processing.
func OrFloat64(
	channels ...<-chan float64,
) <-chan float64 {
	switch len(channels) {
	case 0: // when there are no channels to handle
		return nil
	case 1: // when there is only one channel in the list
		return channels[0]
	}

	orDone := make(chan float64)
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
			case <-OrFloat64(append(channels[3:], orDone)...):
				// here, recursively call or() to after the 3rd channel
			}
		}
	}()
	return orDone
}
