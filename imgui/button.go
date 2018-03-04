package imgui

import (
	"image/color"

	"github.com/qeedquan/go-media/math/f64"
)

type ButtonFlags uint

const (
	ButtonFlagsRepeat ButtonFlags = 1 << iota
	ButtonFlagsPressedOnClickRelease
	ButtonFlagsPressedOnClick
	ButtonFlagsPressedOnRelease
	ButtonFlagsPressedOnDoubleClick
	ButtonFlagsFlattenChildren
	ButtonFlagsAllowItemOverlap
	ButtonFlagsDontClosePopups
	ButtonFlagsDisabled
	ButtonFlagsAlignTextBaseLine
	ButtonFlagsNoKeyModifiers
	ButtonFlagsNoHoldingActiveID
	ButtonFlagsPressedOnDragDropHold
	ButtonFlagsNoNavFocus
)

func (c *Context) Button(label string, size f64.Vec2) bool {
	return c.ButtonEx(label, size, 0)
}

func (c *Context) SmallButton(label string) bool {
	style := c.GetStyle()
	y := style.FramePadding.Y
	style.FramePadding.Y = 0
	pressed := c.ButtonEx(label, f64.Vec2{}, ButtonFlagsAlignTextBaseLine)
	style.FramePadding.Y = y
	return pressed
}

func (c *Context) ArrowButton(strId string, dir Dir) bool {
	window := c.GetCurrentWindow()
	if window.SkipItems {
		return false
	}

	style := c.GetStyle()
	dc := &window.DC
	id := window.GetID(strId)
	sz := c.GetFrameHeight()
	bb := f64.Rectangle{dc.CursorPos, dc.CursorPos.Add(f64.Vec2{sz, sz})}
	c.ItemSize(bb.Size(), 0)
	if !c.ItemAdd(bb, id, nil) {
		return false
	}

	hovered, held, pressed := c.ButtonBehavior(bb, id, 0)

	var col color.RGBA
	if hovered && held {
	} else if hovered {
	} else {
	}

	c.RenderNavHighlight(bb, id)
	c.RenderFrame(bb.Min, bb.Max, col, true, style.FrameRounding)
	c.RenderArrow(bb.Min.Add(style.FramePadding), dir)

	return pressed
}

func (c *Context) ButtonEx(label string, size f64.Vec2, flags ButtonFlags) bool {
	window := c.GetCurrentWindow()
	if window.SkipItems {
		return false
	}

	dc := &window.DC
	style := c.GetStyle()
	pos := dc.CursorPos
	// try to vertically align buttons that are smaller/have no padding so that text baseline matches (bit hacky, since it shouldn't be a flag)
	if flags&ButtonFlagsAlignTextBaseLine != 0 && style.FramePadding.Y < dc.CurrentLineTextBaseOffset {
		pos.Y += dc.CurrentLineTextBaseOffset - style.FramePadding.Y
	}

	var itemSize f64.Vec2
	bb := f64.Rectangle{pos, pos.Add(itemSize)}
	id := ID(0)
	_, _, pressed := c.ButtonBehavior(bb, id, flags)

	// render
	//c.RenderFrame(bb.Min, bb.Max, col, true, style.FrameRounding)

	return pressed
}

func (c *Context) ButtonBehavior(bb f64.Rectangle, id ID, flags ButtonFlags) (outHovered, outHeld, pressed bool) {
	if flags&ButtonFlagsDisabled != 0 {
		if c.ActiveId == id {
			c.ClearActiveID()
		}
		return
	}

	// default behavior requires click+release on same spot
	mask := ButtonFlagsPressedOnClickRelease |
		ButtonFlagsPressedOnClick |
		ButtonFlagsPressedOnRelease |
		ButtonFlagsPressedOnDoubleClick
	if flags&mask == 0 {
		flags |= ButtonFlagsPressedOnClickRelease
	}

	window := c.GetCurrentWindow()
	if flags&ButtonFlagsFlattenChildren != 0 && c.HoveredRootWindow == window {
		c.HoveredWindow = window
	}

	return
}
