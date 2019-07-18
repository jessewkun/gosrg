package ui

func (gv *GView) cursorUp() error {
	ox, oy := gv.View.Origin()
	cx, cy := gv.View.Cursor()
	if err := gv.View.SetCursor(cx, cy-1); err != nil && oy > 0 {
		if err := gv.View.SetOrigin(ox, oy-1); err != nil {
			return err
		}
		return nil
	}
	return nil
}

func (gv *GView) cursorDown() error {
	cx, cy := gv.View.Cursor()
	lineHeight := gv.View.LinesHeight()
	if cy+1 >= lineHeight {
		return nil
	}
	if err := gv.View.SetCursor(cx, cy+1); err != nil {
		ox, oy := gv.View.Origin()
		if err := gv.View.SetOrigin(ox, oy+1); err != nil {
			return err
		}
		return nil
	}
	return nil
}

func (gv *GView) cursorBegin() error {
	if err := gv.View.SetCursor(0, 0); err != nil {
		if err := gv.View.SetOrigin(0, 0); err != nil {
			return err
		}
		return nil
	}
	return nil
}

func (gv *GView) cursorLast() error {
	lineHeight := gv.View.LinesHeight()
	lineHeight--
	lastLine, _ := gv.View.Line(lineHeight)
	lastLineWidth := len(lastLine)
	opView.debug("lineHeight:", lineHeight, "lastLine:", lastLine, "lastLineWidth:", lastLineWidth)
	if err := gv.View.SetCursor(lastLineWidth, lineHeight); err != nil {
		if err := gv.View.SetOrigin(lastLineWidth, lineHeight); err != nil {
			return err
		}
		return nil
	}
	return nil
}

func (gv *GView) cursorEnd() error {
	lineHeight := gv.View.LinesHeight()
	lineHeight--
	opView.debug("lineHeight:", lineHeight)
	if err := gv.View.SetCursor(0, lineHeight); err != nil {
		if err := gv.View.SetOrigin(0, lineHeight); err != nil {
			return err
		}
		return nil
	}
	return nil
}

func (gv *GView) setCursor(x int, y int) error {
	x = 0
	if err := gv.View.SetCursor(x, y); err != nil {
		if err := gv.View.SetOrigin(x, y); err != nil {
			return err
		}
		return nil
	}
	return nil
}
