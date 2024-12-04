package util

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColorsFromLabel_HappyPath(t *testing.T) {
	// Given a valid label with multiple color names separated by commas
	label := "FgRed,BgGreen,OpBold"

	// When calling ColorsFromLabel with the given label
	style, err := ColorsFromLabel(label)

	// Then no error should be returned and the style should contain the expected colors and options
	assert.NoError(t, err)
	assert.Equal(t, color.FgRed|color.BgGreen|color.OpBold, style)
}

func TestColorsFromLabel_InvalidColorName(t *testing.T) {
	// Given a label with an invalid color name
	label := "FgRed,UnknownColor,BgGreen"

	// When calling ColorsFromLabel with the given label
	_, err := ColorsFromLabel(label)

	// Then an error should be returned indicating that the unknown color name is not allowed
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unknown color name 'UnknownColor' in")
}

func TestColorsFromLabel_EmptyLabel(t *testing.T) {
	// Given an empty label
	label := ""

	// When calling ColorsFromLabel with the given label
	style, err := ColorsFromLabel(label)

	// Then no error should be returned and the style should be reset to default
	assert.NoError(t, err)
	assert.Equal(t, color.FgDefault, style)
}

func TestColorsFromLabel_SingleColor(t *testing.T) {
	// Given a label with a single valid color name
	label := "FgBlue"

	// When calling ColorsFromLabel with the given label
	style, err := ColorsFromLabel(label)

	// Then no error should be returned and the style should contain only the specified color
	assert.NoError(t, err)
	assert.Equal(t, color.FgBlue, style)
}

func TestColorsFromLabel_MultipleOptions(t *testing.T) {
	// Given a label with multiple options separated by commas
	label := "OpBold,OpItalic"

	// When calling ColorsFromLabel with the given label
	style, err := ColorsFromLabel(label)

	// Then no error should be returned and the style should contain only the specified options
	assert.NoError(t, err)
	assert.Equal(t, color.OpBold|color.OpItalic, style)
}

func TestColorsFromLabel_MixedColorAndOption(t *testing.T) {
	// Given a label with both valid color names and options separated by commas
	label := "FgRed,OpBold"

	// When calling ColorsFromLabel with the given label
	style, err := ColorsFromLabel(label)

	// Then no error should be returned and the style should contain the specified color and option
	assert.NoError(t, err)
	assert.Equal(t, color.FgRed|color.OpBold, style)
}

func TestColorsFromLabel_LeadingTrailingSpaces(t *testing.T) {
	// Given a label with leading and trailing spaces around valid color names
	label := "  FgGreen,BgBlue  "

	// When calling ColorsFromLabel with the given label
	style, err := ColorsFromLabel(label)

	// Then no error should be returned and the style should contain only the specified colors
	assert.NoError(t, err)
	assert.Equal(t, color.FgGreen|color.BgBlue, style)
}

func TestColorsFromLabel_CommaOnly(t *testing.T) {
	// Given a label with only commas
	label := ","

	// When calling ColorsFromLabel with the given label
	style, err := ColorsFromLabel(label)

	// Then no error should be returned and the style should be reset to default
	assert.NoError(t, err)
	assert.Equal(t, color.FgDefault, style)
}

func TestColorsFromLabel_SingleInvalidColorName(t *testing.T) {
	// Given a label with a single invalid color name
	label := "UnknownColor"

	// When calling ColorsFromLabel with the given label
	_, err := ColorsFromLabel(label)

	// Then an error should be returned indicating that the unknown color name is not allowed
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unknown color name 'UnknownColor' in")
}

func TestColorsFromLabel_MultipleInvalidColorNames(t *testing.T) {
	// Given a label with multiple invalid color names separated by commas
	label := "UnknownColor1,UnknownColor2"

	// When calling ColorsFromLabel with the given label
	_, err := ColorsFromLabel(label)

	// Then an error should be returned indicating that the unknown color names are not allowed
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unknown color name 'UnknownColor1' in")
	assert.Contains(t, err.Error(), "unknown color name 'UnknownColor2' in")
}

func TestColorsFromLabel_EmptyString(t *testing.T) {
	// Given an empty string as the label
	label := ""

	// When calling ColorsFromLabel with the given label
	style, err := ColorsFromLabel(label)

	// Then no error should be returned and the style should be reset to default
	assert.NoError(t, err)
	assert.Equal(t, color.FgDefault, style)
}

func TestColorsFromLabel_WhitespaceOnly(t *testing.T) {
	// Given a string with only whitespace as the label
	label := "   "

	// When calling ColorsFromLabel with the given label
	style, err := ColorsFromLabel(label)

	// Then no error should be returned and the style should be reset to default
	assert.NoError(t, err)
	assert.Equal(t, color.FgDefault, style)
}

func TestColorsFromLabel_MixedCaseColorNames(t *testing.T) {
	// Given a label with mixed case color names
	label := "FgRed,bGgreen,OpBold"

	// When calling ColorsFromLabel with the given label
	style, err := ColorsFromLabel(label)

	// Then no error should be returned and the style should contain the expected colors and options
	assert.NoError(t, err)
	assert.Equal(t, color.FgRed|color.BgGreen|color.OpBold, style)
}

func TestColorsFromLabel_MixedCaseOptionNames(t *testing.T) {
	// Given a label with mixed case option names
	label := "Opbold,bGiTalic"

	// When calling ColorsFromLabel with the given label
	style, err := ColorsFromLabel(label)

	// Then no error should be returned and the style should contain the expected options
	assert.NoError(t, err)
	assert.Equal(t, color.OpBold|color.OpItalic, style)
}
