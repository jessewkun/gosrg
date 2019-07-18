package ui

import (
	"fmt"
	"gosrg/utils"
)

func (gv *GView) cursorUp() error {
	ox, oy := gv.View.Origin()
	cx, cy := gv.View.Cursor()
	if err := gv.View.SetCursor(cx, cy-1); err != nil && oy > 0 {
		if err := gv.View.SetOrigin(ox, oy-1); err != nil {
			utils.Logger.Println(err)
			return err
		}
		return nil
	}
	return nil
}

func (gv *GView) cursorDown() error {
	cx, cy := gv.View.Cursor()
	_, oy := gv.View.Origin()
	lineHeight := gv.View.LinesHeight()
	lineHeight--
	if cy+oy+1 >= lineHeight {
		return nil
	}
	utils.Logger.Println(cy, oy, lineHeight)
	if err := gv.View.SetCursor(cx, cy+1); err != nil {
		utils.Logger.Println(err)
		ox, oy := gv.View.Origin()
		if err := gv.View.SetOrigin(ox, oy+1); err != nil {
			utils.Logger.Println(err)
			return err
		}
		return nil
	}
	return nil
}

func (gv *GView) cursorBegin() error {
	if err := gv.View.SetCursor(0, 0); err != nil {
		utils.Logger.Println(err)
	}
	if err := gv.View.SetOrigin(0, 0); err != nil {
		utils.Logger.Println(err)
		return err
	}
	kView.cursorDebug()
	return nil
}

func (gv *GView) cursorEnd(flag bool) error {
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
		utils.Logger.Println(err)
	}
	if err := gv.View.SetOrigin(ox, oy); err != nil {
		utils.Logger.Println(err)
		return err
	}
	kView.cursorDebug()
	return nil
}

func (gv *GView) setCursor(x int, y int) error {
	if err := gv.View.SetCursor(x, y); err != nil {
		// if err := gv.View.SetOrigin(x, y); err != nil {
		// 	utils.Logger.Println(err)
		// 	return err
		// }
		utils.Logger.Println(err)
		return nil
	}
	return nil
}

func (gv *GView) cursorDebug() {
	x, y := gv.View.Size()
	ox, oy := gv.View.Origin()
	cx, cy := gv.View.Cursor()
	rx, ry := gv.View.ReadPos()
	wx, wy := gv.View.WritePos()

	str := fmt.Sprintf("size: x: %d y: %d  orign x: %d y: %d  cursor x: %d y: %d  read x: %d y: %d  write x: %d y: %d", x, y, ox, oy, cx, cy, rx, ry, wx, wy)
	opView.debug(str)
}
