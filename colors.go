package main

import (
	"fmt"

	"github.com/gookit/color"
	"github.com/gookit/goutil/strutil"
)

// Colors ...
var Colors = map[string]color.Color{
	"FgBlack":         color.FgBlack,
	"FgRed":           color.FgRed,
	"FgGreen":         color.FgGreen,
	"FgYellow":        color.FgYellow,
	"FgBlue":          color.FgBlue,
	"FgMagenta":       color.FgMagenta,
	"FgCyan":          color.FgCyan,
	"FgWhite":         color.FgWhite,
	"FgDefault":       color.FgDefault,
	"FgDarkGray":      color.FgDarkGray,
	"FgLightRed":      color.FgLightRed,
	"FgLightGreen":    color.FgLightGreen,
	"FgLightYellow":   color.FgLightYellow,
	"FgLightBlue":     color.FgLightBlue,
	"FgLightMagenta":  color.FgLightMagenta,
	"FgLightCyan":     color.FgLightCyan,
	"FgLightWhite":    color.FgLightWhite,
	"FgGray":          color.FgGray,
	"BgBlack":         color.BgBlack,
	"BgRed":           color.BgRed,
	"BgGreen":         color.BgGreen,
	"BgYellow":        color.BgYellow,
	"BgBlue":          color.BgBlue,
	"BgMagenta":       color.BgMagenta,
	"BgCyan":          color.BgCyan,
	"BgWhite":         color.BgWhite,
	"BgDefault":       color.BgDefault,
	"BgDarkGray":      color.BgDarkGray,
	"BgLightRed":      color.BgLightRed,
	"BgLightGreen":    color.BgLightGreen,
	"BgLightYellow":   color.BgLightYellow,
	"BgLightBlue":     color.BgLightBlue,
	"BgLightMagenta":  color.BgLightMagenta,
	"BgLightCyan":     color.BgLightCyan,
	"BgLightWhite":    color.BgLightWhite,
	"BgGray":          color.BgGray,
	"OpReset":         color.OpReset,
	"OpBold":          color.OpBold,
	"OpFuzzy":         color.OpFuzzy,
	"OpItalic":        color.OpItalic,
	"OpUnderscore":    color.OpUnderscore,
	"OpBlink":         color.OpBlink,
	"OpFastBlink":     color.OpFastBlink,
	"OpReverse":       color.OpReverse,
	"OpConcealed":     color.OpConcealed,
	"OpStrikethrough": color.OpStrikethrough,
	"Red":             color.Red,
	"Cyan":            color.Cyan,
	"Gray":            color.Gray,
	"Blue":            color.Blue,
	"Black":           color.Black,
	"Green":           color.Green,
	"White":           color.White,
	"Yellow":          color.Yellow,
	"Magenta":         color.Magenta,
	"Bold":            color.Bold,
	"Normal":          color.Normal,
	"LightRed":        color.LightRed,
	"LightCyan":       color.LightCyan,
	"LightBlue":       color.LightBlue,
	"LightGreen":      color.LightGreen,
	"LightWhite":      color.LightWhite,
	"LightYellow":     color.LightYellow,
	"LightMagenta":    color.LightMagenta,
}

// ColorsFromLabel ...
func ColorsFromLabel(label string) (color.Style, error) {
	colorNames := strutil.Split(label, ",")
	r := make([]color.Color, 0, len(colorNames))

	for _, colorName := range colorNames {
		c, has := Colors[colorName]
		if !has {
			return nil, fmt.Errorf("unknown color name '%s' in '%s'. allowed: %v", colorName, label, Colors)
		}
		r = append(r, c)
	}

	return color.New(r...), nil
}