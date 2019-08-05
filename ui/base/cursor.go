package base

import (
	"gosrg/utils"
)

func (gv *GView) CursorUp() error {
	ox, oy := gv.View.Origin()
	cx, cy := gv.View.Cursor()
	if err := gv.View.SetCursor(cx, cy-1); err != nil && oy > 0 {
		if err := gv.View.SetOrigin(ox, oy-1); err != nil {
			utils.Error.Println(err)
			return err
		}
		return nil
	}
	return nil
}

func (gv *GView) CursorDown() error {
	cx, cy := gv.View.Cursor()
	ox, oy := gv.View.Origin()
	lineHeight := gv.View.LinesHeight()
	lineHeight--
	if cy+oy+1 > lineHeight {
		return nil
	}
	if err := gv.View.SetCursor(cx, cy+1); err != nil {
		utils.Error.Println(err)
		if err := gv.View.SetOrigin(ox, oy+1); err != nil {
			utils.Error.Println(err)
			return err
		}
		return nil
	}
	return nil
}

func (gv *GView) CursorBegin() error {
	if err := gv.View.SetCursor(0, 0); err != nil {
		utils.Error.Println(err)
	}
	if err := gv.View.SetOrigin(0, 0); err != nil {
		utils.Error.Println(err)
		return err
	}
	return nil
}

func (gv *GView) CursorEnd(flag bool) error {
	_, row := gv.View.Size()
	row--
	lineHeight := gv.View.LinesHeight()
	lineHeight--
	cx, cy, ox, oy := 0, 0, 0, 0
	lastLine, _ := gv.View.Line(lineHeight)
	lastLineWidth := 0
	if flag == true {
		lastLineWidth = len(lastLine)
	}

	if lineHeight > row {
		cx, cy = lastLineWidth, row
		ox, oy = 0, lineHeight-row
	} else {
		cx, cy = lastLineWidth, lineHeight
		ox, oy = 0, 0
	}
	if err := gv.View.SetCursor(cx, cy); err != nil {
		utils.Error.Println(err)
	}
	if err := gv.View.SetOrigin(ox, oy); err != nil {
		utils.Error.Println(err)
		return err
	}
	return nil
}

func (gv *GView) SetCursor(x int, y int) error {
	if err := gv.View.SetCursor(x, y); err != nil {
		// if err := gv.View.SetOrigin(x, y); err != nil {
		// 	utils.Error.Println(err)
		// 	return err
		// }
		utils.Error.Println(err)
		return nil
	}
	return nil
}
