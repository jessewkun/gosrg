package base

import (
	"gosrg/utils"
)

// CursorUp used to move the cursor up
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

// CursorDown used to move the cursor down
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

// CursorBegin used to move the cursor to the begining of the view
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

// CursorEnd used to move the cursor to the end of the view
// if flag is true, the cursor will moved to the end of last line,
// if flag is false, the cursor will moved to the beginning of last line
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

// SetCursor used to set cursor's position
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
