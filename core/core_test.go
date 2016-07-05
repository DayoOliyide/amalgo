package core

import (
	"testing"
	"os"
	"reflect"
)


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

func TestEnvironmentMap(t *testing.T)  {
	cases := []struct {
		envKey, envVal, expectedVal string
	}{
		{"AMALGO_TEST_ENV_1", "value1", "value1"},
		{"AMALGO_TEST_ENV_2", "value2=value2",  "value2=value2"},
		{"AMALGO_TEST_ENV_3", "/some/file/path",  "/some/file/path"},
		{"AMALGO_TEST_ENV_4", "space seperated value",  "space seperated value"},
		{"AMALGO_TEST_ENV_5", "comma,seperated,value",  "comma,seperated,value"},
		{"AMALGO_TEST_ENV_6", "semicolon;seperated;value",  "semicolon;seperated;value"},
	}
	for _, c := range cases {
		os.Setenv(c.envKey, c.envVal)
		envMap := environmentMap()
		os.Unsetenv(c.envKey)
		actualVal, present := envMap[c.envKey]
		if !present || actualVal != c.expectedVal {
			t.Errorf("Environment value is %q, expected value is %q. Got %q",
			c.envVal, c.expectedVal, actualVal)
		}
	}
}

func TestResolvables(t *testing.T) {
	cases := []struct {
		input string; expectedVal []string
	}{
		{"%%TestCase1%%", []string{"%%TestCase1%%"}},
		{"%TestCase2%", []string{}},
		{"%%TestCase3%%blahblah%NONO%", []string{"%%TestCase3%%"}},
		{"%%TestCase4A%%\n\tblah%%TestCase4B%%blahblah\n  %%TestCase4C%%", []string{"%%TestCase4A%%",
			                                                                    "%%TestCase4B%%",
		                                                                            "%%TestCase4C%%"}},
	}
	for _, c := range cases {
		actualVal := Resolvables(c.input)
		if !reflect.DeepEqual(c.expectedVal, actualVal) {
			t.Errorf("Provided Template -> %q Expected value -> %q Actual value -> %q", c.input,
				c.expectedVal, actualVal)
		}
	}
}
