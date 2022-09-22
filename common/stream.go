package common

type Stream struct {
	source <-chan any
	walks  []WalkFunc
}
type FilterFunc func(item any) bool
type WalkFunc func(item any, pipe chan<- any)

func Just(items ...any) Stream {
	source := make(chan any, len(items))
	for _, item := range items {
		source <- item
	}
	close(source)
	return Range(source)
}

func Range(source <-chan any) Stream {
	return Stream{
		source: source,
	}
}

func (s Stream) Filter(fn FilterFunc) Stream {
	return s.Walk(func(item any, pipe chan<- any) {
		if fn(item) {
			pipe <- item
		}
	})
}

func (s Stream) Walk(fn WalkFunc) Stream {
	s.walks = append(s.walks, fn)
	return s
}

func (s Stream) startWalkWork() Stream {
	pipe := make(chan any)
	go func() {
		for item := range s.source {
			for _, fn := range s.walks {
				fn(item, pipe)
			}
		}
		close(pipe)
	}()
	return Range(pipe)
}

func (s Stream) Count() (count int) {
	newStream := s.startWalkWork()
	for range newStream.source {
		count++
	}
	return
}

func (s Stream) First() any {
	newStream := s.startWalkWork()
	for item := range newStream.source {
		return item
	}
	return nil
}
