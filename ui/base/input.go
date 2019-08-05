package base

import (
	"math"
	"strings"
)

type Input struct {
	// input label
	Label string

	// input name
	Name string

	// input value
	Value string
}

// padLabel used to fill the label to the maximum length
func (i *Input) padLabel(f *Form) {
	if f.Done {
		return
	}
	padLen := f.MaxLabel - len(i.Label)
	if padLen > 0 {
		switch f.LabelAlign {
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
	if f.MarginLeft > 0 {
		i.Label = strings.Repeat(PAD_STR, f.MarginLeft) + i.Label
	}
}
