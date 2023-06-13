package application

import (
	"fmt"
	"sync"
)

type Photoshop struct {
	saved      *[]string
	inProgress *[]string
}

var once sync.Once

var photoshopInstance *Photoshop

func GetPhotoshopApp() *Photoshop {
	if photoshopInstance == nil {
		once.Do(func() {
			photoshopInstance = &Photoshop{}
		})
	}
	return photoshopInstance
}

func (p *Photoshop) Save() {
	if p.inProgress != nil {
		saved := *p.inProgress
		if p.saved == nil {
			p.saved = &saved
		} else {
			*p.saved = append(*p.saved, saved...)
		}
		p.inProgress = nil
	} else {
		fmt.Println("No work in progress to save")
	}
}

func (p *Photoshop) Undo() {
	if p.inProgress != nil {
		*p.inProgress = (*p.inProgress)[:len(*p.inProgress)-1]
		if len(*p.inProgress) == 0 {
			p.inProgress = nil
		}
	} else if p.saved != nil {
		*p.saved = (*p.saved)[:len(*p.saved)-1]
		if len(*p.saved) == 0 {
			p.saved = nil
		}
	} else {
		fmt.Println("Nothing to undo")
	}
}

func (p *Photoshop) Work(val string) {
	if p.inProgress != nil {
		*p.inProgress = append(*p.inProgress, val)
	} else {
		inProgress := []string{val}
		p.inProgress = &inProgress
	}
}

func (p *Photoshop) DisplayWork() {
	values := "==============\nSaved Work: %+v\nWork in progress: %+v\n==============\n"
	fmt.Printf(values, p.saved, p.inProgress)
}
