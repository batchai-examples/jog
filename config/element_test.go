package config

import (
	"testing"
)

func TestElementT_UnmarshalYAML(t *testing.T) {
	tests := []struct {
		name    string
		input   map[string]interface{}
		wantErr bool
	}{
		{
			name: "Happy path",
			input: map[string]interface{}{
				"color":       "#FF0000",
				"print":       true,
				"print-format": "%s",
			},
			wantErr: false,
		},
		{
			name: "Negative path - invalid color",
			input: map[string]interface{}{
				"color":       "invalid_color",
				"print":       true,
				"print-format": "%s",
			},
			wantErr: true,
		},
		{
			name: "Negative path - invalid print format",
			input: map[string]interface{}{
				"color":       "#FF0000",
				"print":       true,
				"print-format": "%5.s",
			},
			wantErr: true,
		},
		{
			name: "Corner case - empty input",
			input: map[string]interface{}{},
			wantErr: false,
		},
		{
			name: "Negative path - missing print format",
			input: map[string]interface{}{
				"color": "#FF0000",
				"print": true,
			},
			wantErr: false,
		},
		{
			name: "Negative path - invalid boolean value for print",
			input: map[string]interface{}{
				"color":       "#FF0000",
				"print":       "invalid_bool",
				"print-format": "%s",
			},
			wantErr: true,
		},
		{
			name: "Negative path - invalid color format",
			input: map[string]interface{}{
				"color":       "#ZZZZZ",
				"print":       true,
				"print-format": "%s",
			},
			wantErr: true,
		},
		{
			name: "Happy path - with default values",
			input: map[string]interface{}{
				"color": "#FF0000",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &ElementT{}
			err := i.FromMap(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestElementT_ToMap(t *testing.T) {
	tests := []struct {
		name    string
		input   Element
		want    map[string]interface{}
	}{
		{
			name: "Happy path",
			input: &ElementT{
				Color:       util.Color{Value: "#FF0000"},
				Print:       true,
				PrintFormat: "%s",
			},
			want: map[string]interface{}{
				"color":       "#FF0000",
				"print":       true,
				"print-format": "%s",
			},
		},
		{
			name: "Corner case - default values",
			input: &ElementT{},
			want: map[string]interface{}{
				"color":       "#FFFFFF",
				"print":       true,
				"print-format": "%s",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.ToMap()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestElementT_Reset(t *testing.T) {
	tests := []struct {
		name string
		want Element
	}{
		{
			name: "Happy path",
			want: &ElementT{
				Color:       util.Color{Value: "#FFFFFF"},
				Print:       true,
				PrintFormat: "%s",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &ElementT{}
			i.Reset()
			if !reflect.DeepEqual(i, tt.want) {
				t.Errorf("Reset() = %v, want %v", i, tt.want)
			}
		})
	}
}

func TestElementT_PrintFormat(t *testing.T) {
	tests := []struct {
		name    string
		input   Element
		want    string
	}{
		{
			name: "Happy path",
			input: &ElementT{
				PrintFormat: "%s",
			},
			want: "%s",
		},
		{
			name: "Corner case - default value",
			input: &ElementT{},
			want: "%s",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.PrintFormat()
			if got != tt.want {
				t.Errorf("PrintFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestElementT_Color(t *testing.T) {
	tests := []struct {
		name    string
		input   Element
		want    util.Color
	}{
		{
			name: "Happy path",
			input: &ElementT{
				Color: util.Color{Value: "#FF0000"},
			},
			want: util.Color{Value: "#FF0000"},
		},
		{
			name: "Corner case - default value",
			input: &ElementT{},
			want: util.Color{Value: "#FFFFFF"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.Color()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Color() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestElementT_Print(t *testing.T) {
	tests := []struct {
		name    string
		input   Element
		want    bool
	}{
		{
			name: "Happy path",
			input: &ElementT{
				Print: true,
			},
			want: true,
		},
		{
			name: "Corner case - default value",
			input: &ElementT{},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.Print()
			if got != tt.want {
				t.Errorf("Print() = %v, want %v", got, tt.want)
			}
		})
	}
}
