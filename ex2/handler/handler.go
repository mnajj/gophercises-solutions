package handler

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"net/http"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p := r.URL.Path
		if v, ok := pathsToUrls[p]; ok {
			http.Redirect(w, r, v, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var pathUrls []pathUrl
	err := yaml.Unmarshal(yml, &pathUrls)
	fmt.Println(pathUrls)
	if err != nil {
		return nil, err
	}
	ymlMap := buildMap(pathUrls)
	fmt.Println(len(ymlMap))
	return MapHandler(ymlMap, fallback), nil
}

type pathUrl struct {
	Path string
	Url  string
}

func buildMap(paths []pathUrl) map[string]string {
	pathsMap := map[string]string{}
	for _, v := range paths {
		pathsMap[v.Path] = v.Url
	}
	return pathsMap
}
