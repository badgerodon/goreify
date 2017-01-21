package examples

import "github.com/badgerodon/goreify/generics"

// Merge merges all the src channels into a single channel.
// If all of the src channels are closed, the returned channel
// will also be closed.
func Merge(srcs ...<-chan generics.T1) <-chan generics.T1 {
	switch len(srcs) {
	case 0:
		return nil
	case 1:
		return srcs[0]
	case 2:
		dst := make(chan generics.T1)
		go func() {
			// merge until both channels are closed, then finally close the destination
			defer close(dst)

			c0, c1 := srcs[0], srcs[1]
			for c0 != nil || c1 != nil {
				select {
				case m, ok := <-c0:
					if !ok {
						c0 = nil
						continue
					}
					dst <- m
				case m, ok := <-c1:
					if !ok {
						c1 = nil
						continue
					}
					dst <- m
				}
			}
		}()
		return dst
	default:
		left, right := srcs[:len(srcs)/2], srcs[len(srcs)/2:]
		return Merge(Merge(left...), Merge(right...))
	}
}
