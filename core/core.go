package core

import (
	"regexp"
	"os"
	"strings"
)

func environmentMap() map[string]string {
	envs := os.Environ()
	envMap := make(map[string]string)
	for _, envString := range envs  {
		kv := strings.SplitN(envString, "=", 2)
		envMap[kv[0]] = kv[1]
		}
	return envMap
}

func Resolvables(template string) []string {
	reg := regexp.MustCompile("%%[^%]+%%")
	resolvAll := reg.FindAllString(template, -1)
	//why golang, why no sets ?!! why ????
	resolvMap := make(map[string]string)
	for _, resolv := range resolvAll {
		if _, present := resolvMap[resolv]; !present {
			resolvMap[resolv] = resolv
		}
	}

	keys := make([]string, len(resolvMap))
	i := 0
	for k, _ := range resolvMap {
		keys[i] = k
		i++
	}

	return keys
}

func OutfileName(s string) (replaced string) {
	reg := regexp.MustCompile("\\.tmpl$")
	replaced = reg.ReplaceAllLiteralString(s, "")
	if s == replaced {
		replaced += ".out"
	}
	return replaced
}

/*
func ResolveFileTemplate(envMap map[string]string, templateFile, outfile string) {
	file, fileErr := ioutil.ReadFile(templateFile)
	if fileErr != nil {
		panic(fileErr)
	}


}*/
