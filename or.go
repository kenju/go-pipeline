package pipeline

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "GenType=interface{},byte,string,int,uint64,float32,float64"

// OrGenType return only the first result from multiple channels.
// Use ctx to cancel the stream processing.
func OrGenType(
	channels ...<-chan GenType,
) <-chan GenType {
	switch len(channels) {
	case 0: // when there are no channels to handle
		return nil
	case 1: // when there is only one channel in the list
		return channels[0]
	}

	orDone := make(chan GenType)
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
			case <-OrGenType(append(channels[3:], orDone)...):
				// here, recursively call or() to after the 3rd channel
			}
		}
	}()
	return orDone
}
