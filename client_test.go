package darktrace

import "testing"

func TestMakeURL(t *testing.T) {
	data := [][3]string{
		[3]string{"https://darktrace.example.com", "/devices", "https://darktrace.example.com/devices"},
		[3]string{"https://darktrace.example.com", "autocomplete", "https://darktrace.example.com/autocomplete"},
		[3]string{"https://darktrace.example.com/", "/devices", "https://darktrace.example.com/devices"},
		[3]string{"https://darktrace.example.com", "/autocomplete/", "https://darktrace.example.com/autocomplete"},
	}
	for _, in := range data {
		c := Client{baseUrl: in[0]}
		if u := c.makeURL(in[1]); u != in[2] {
			t.Errorf("Client.makeURL('%s'): found '%s', expected '%s'", in[1], u, in[2])
		}
	}
}
