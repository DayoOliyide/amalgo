package core

import "testing"


func TestOutfileName(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"TestFile.tmpl", "TestFile"},
		{"TestFile.tmpl.tmpl", "TestFile.tmpl"},
		{"TestFile.conf.tmpl", "TestFile.conf"},
		{"TestFile", "TestFile.out"},
		{"", ".out"},
	}
	for _, c := range cases {
		got := OutfileName(c.in)
		if got != c.want {
			t.Errorf("OutfileName(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
