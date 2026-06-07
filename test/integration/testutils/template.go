package testutils

import (
	"strings"
	"text/template"

	yaml "gopkg.in/yaml.v2"
)

// GoldenTemplateFuncs are the template helpers used to render the golden output
// .tpl files. They are shared by the Prometheus CLI integration tests and the
// library use-case tests, both of which render the same templates and must
// therefore agree on how values are rendered.
var GoldenTemplateFuncs = template.FuncMap{
	// yamlValue renders a string the way Sloth's YAML serializers do when it is
	// emitted as a mapping value. The version comes from `git describe` and can
	// be an all-digit short git SHA (e.g. "6950793"), which the serializers quote
	// so it is not parsed back as a number. Using this for the sloth_version label
	// keeps the golden files correct regardless of whether the version happens to
	// look numeric.
	"yamlValue": func(v string) (string, error) {
		b, err := yaml.Marshal(v)
		if err != nil {
			return "", err
		}
		return strings.TrimRight(string(b), "\n"), nil
	},
}
