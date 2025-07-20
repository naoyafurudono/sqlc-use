package plugin

import (
	"encoding/json"
	"testing"
)

func TestOptions_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		json    string
		want    Options
		wantErr bool
	}{
		{
			name: "valid options with package",
			json: `{"package": "myapp.db", "format": "json"}`,
			want: Options{
				Package: "myapp.db",
				Format:  "json",
			},
			wantErr: false,
		},
		{
			name: "only package",
			json: `{"package": "db"}`,
			want: Options{
				Package: "db",
				Format:  "",
			},
			wantErr: false,
		},
		{
			name: "empty options",
			json: `{}`,
			want: Options{
				Package: "",
				Format:  "",
			},
			wantErr: false,
		},
		{
			name:    "invalid json",
			json:    `{invalid}`,
			want:    Options{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got Options
			err := json.Unmarshal([]byte(tt.json), &got)

			if (err != nil) != tt.wantErr {
				t.Errorf("json.Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && got != tt.want {
				t.Errorf("json.Unmarshal() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestDefaultOptions(t *testing.T) {
	got := DefaultOptions()
	want := Options{
		Format: "json",
	}

	if got != want {
		t.Errorf("DefaultOptions() = %+v, want %+v", got, want)
	}
}
