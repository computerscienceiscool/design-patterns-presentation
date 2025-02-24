package observer

import "fmt"

type Observer interface {
	Update(videoTitle string)
}

type Subscriber struct {
	name string
}

func (s *Subscriber) Update(videoTitle string) {
	fmt.Printf("%s received notification for new video: %s\n", s.name, videoTitle)
}

type YouTubeChannel struct {
	observers []Observer
}

func (yc *YouTubeChannel) Subscribe(observer Observer) {
	yc.observers = append(yc.observers, observer)
}

func (yc *YouTubeChannel) Notify(videoTitle string) {
	for _, observer := range yc.observers {
		observer.Update(videoTitle)
	}
}

func DemoObserver() {
	channel := &YouTubeChannel{}
	sub1 := &Subscriber{name: "Alice"}
	channel.Subscribe(sub1)

	channel.Notify("New Go Tutorial")
}
