package structs

import "net/url"

type History struct {
	previousPages []*url.URL
	nextPages []*url.URL
}

func CreateHistory() *History {
	return &History{

	}
}

func (history *History) NextPages() []*url.URL {
	return history.nextPages
}

func (history *History) AllPages() []*url.URL {
	return history.previousPages
}

func (history *History) PageCount() int {
	return len(history.previousPages)
}

func (history *History) Push(URL *url.URL) {
	history.nextPages = nil
	history.previousPages = append(history.previousPages, URL)
}

func (history *History) Last() *url.URL {
	return history.previousPages[len(history.previousPages)-1]
}

func (history *History) PopNext() {
	if len(history.nextPages) > 0 {
		history.previousPages = append(history.previousPages, history.nextPages[len(history.nextPages)-1])
		history.nextPages = nil
	}
}

func (history *History) Pop() {
	if len(history.previousPages) > 0 {
		history.nextPages = append(history.nextPages, history.previousPages[len(history.previousPages)-1])
		history.previousPages = history.previousPages[:len(history.previousPages)-1]
	}
}

