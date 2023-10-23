package helpers

import "testing"

func TestGetUrl(t *testing.T) {
	tests := []struct {
		name   string
		apiUrl string
		apiKey string
		params map[string]string
		want   string
	}{
		{
			name:   "test with no params",
			apiUrl: "https://api.etherscan.io/api",
			apiKey: "myapikey",
			params: map[string]string{},
			want:   "https://api.etherscan.io/api/?apikey=myapikey",
		},
		{
			name:   "test with one param",
			apiUrl: "https://api.etherscan.io/api",
			apiKey: "myapikey",
			params: map[string]string{"param1": "value1"},
			want:   "https://api.etherscan.io/api/?apikey=myapikey&param1=value1",
		},
		{
			name:   "test with multiple params",
			apiUrl: "https://api.etherscan.io/api",
			apiKey: "myapikey",
			params: map[string]string{"param1": "value1", "param2": "value2"},
			want:   "https://api.etherscan.io/api/?apikey=myapikey&param1=value1&param2=value2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetUrl(tt.apiUrl, tt.apiKey, tt.params)
			if got != tt.want {
				t.Errorf("GetUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
