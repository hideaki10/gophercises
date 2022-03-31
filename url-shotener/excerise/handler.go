package urlshotener

import (
	"net/http"

	yaml "gopkg.in/yaml.v2"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {

	// return handler func
	return func(w http.ResponseWriter, r *http.Request) {
		// 判定 request path 是否在pathesToUrls 中
		// 如果在，则重定向到对应的 url
		// 没有就返回 fallback的handler

		url := r.URL.Path

		if dest, ok := pathsToUrls[url]; ok {
			http.Redirect(w, r, dest, 200)
			return
		}

		fallback.ServeHTTP(w, r)

	}

}

func YAMLHandler(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// parse yaml
	var pathurls []pathUrl
	err := yaml.Unmarshal(yamlBytes, &pathurls)
	if err != nil {
		return nil, err
	}
	// build map
	var buildMap map[string]string

	for _, p := range pathurls {
		buildMap[p.Path] = p.URL
	}
	return MapHandler(buildMap, fallback), nil
}

type pathUrl struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}
