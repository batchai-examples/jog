package config

import (
	"testing"
)

func TestSeparatorField_UnmarshalYAML(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name: "Happy path with label",
			input: `
label: ":="
`,
			wantErr: false,
		},
		{
			name: "Happy path without label",
			input: ``,
			wantErr: false,
		},
		{
			name:    "Negative path - invalid YAML",
			input:   `invalid yaml`,
			wantErr: true,
		},
		{
			name:    "Corner case - empty input",
			input:   ``,
			wantErr: false,
		},
		{
			name: "Negative path - non-string label",
			input: `
label: 123
`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var sf SeparatorFieldT
			err := sf.UnmarshalYAML(func(v interface{}) error {
				return unmarshalYAMLString(tt.input, v)
			})
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalYAML() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSeparatorField_MarshalYAML(t *testing.T) {
	tests := []struct {
		name    string
		input   SeparatorFieldT
		want    interface{}
	}{
		{
			name: "Happy path",
			input: SeparatorFieldT{
				ElementT: ElementT{},
				Label:    ":=",
			},
			want: map[string]interface{}{
				"label": ":=",
			},
		},
		{
			name: "Corner case - empty label",
			input: SeparatorFieldT{
				ElementT: ElementT{},
				Label:    "",
			},
			want: map[string]interface{}{
				"label": "=",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.input.MarshalYAML()
			if err != nil {
				t.Errorf("MarshalYAML() error = %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalYAML() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSeparatorField_Reset(t *testing.T) {
	tests := []struct {
		name string
		want SeparatorFieldT
	}{
		{
			name: "Happy path",
			want: SeparatorFieldT{
				ElementT: ElementT{},
				Label:    "=",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sf := &SeparatorFieldT{
				ElementT: ElementT{},
				Label:    ":=",
			}
			sf.Reset()
			if !reflect.DeepEqual(*sf, tt.want) {
				t.Errorf("Reset() got = %v, want %v", sf, tt.want)
			}
		})
	}
}

func TestSeparatorField_FromMap(t *testing.T) {
	tests := []struct {
		name    string
		input   map[string]interface{}
		wantErr bool
	}{
		{
			name: "Happy path with label",
			input: map[string]interface{}{
				"label": ":=",
			},
			wantErr: false,
		},
		{
			name: "Happy path without label",
			input: map[string]interface{}{},
			wantErr: false,
		},
		{
			name:    "Negative path - invalid label type",
			input: map[string]interface{}{
				"label": 123,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var sf SeparatorFieldT
			err := sf.FromMap(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromMap() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSeparatorField_ToMap(t *testing.T) {
	tests := []struct {
		name    string
		input   SeparatorFieldT
		want    map[string]interface{}
	}{
		{
			name: "Happy path",
			input: SeparatorFieldT{
				ElementT: ElementT{},
				Label:    ":=",
			},
			want: map[string]interface{}{
				"label": ":=",
			},
		},
		{
			name: "Corner case - empty label",
			input: SeparatorFieldT{
				ElementT: ElementT{},
				Label:    "",
			},
			want: map[string]interface{}{
				"label": "=",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.ToMap()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToMap() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func unmarshalYAMLString(s string, v interface{}) error {
	return yaml.Unmarshal([]byte(s), v)
}
