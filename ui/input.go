package ui

import (
	"math"
	"strings"
)

type Input struct {
	Label string
	Name  string
	Value string
}

func (i *Input) padLabel(f *Form) {
	if f.done {
		return
	}
	padLen := f.MaxLabel - len(i.Label)
	if padLen > 0 {
		switch f.labelAlign {
		case ALIGN_LEFT:
			i.Label += strings.Repeat(PAD_STR, padLen)
		case ALIGN_MIDDLE:
			if padLen%2 != 0 {
				t := int(math.Floor(float64(padLen / 2)))
				left := strings.Repeat(PAD_STR, t)
				right := strings.Repeat(PAD_STR, t+1)
				i.Label = left + i.Label + right
			} else {
				t := strings.Repeat(PAD_STR, padLen/2)
				i.Label = t + i.Label + t
			}
		case ALIGN_RIGHT:
			i.Label = strings.Repeat(PAD_STR, padLen) + i.Label
		}
	}
	if f.marginLeft > 0 {
		i.Label = strings.Repeat(PAD_STR, f.marginLeft) + i.Label
	}
}
