package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/elastic/go-ucfg"
	"github.com/elastic/go-ucfg/yaml"

	"golang.org/x/tools/imports"
)

var cfgOpts = []ucfg.Option{
	ucfg.PathSep("."),
	ucfg.ResolveEnv,
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	to := flag.String("o", "", "write to")
	format := flag.Bool("f", false, "format output using goimports")

	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		return errors.New("Missing input file")
	}

	cfg, err := yaml.NewConfigWithFile(args[0], cfgOpts...)
	if err != nil {
		return fmt.Errorf("Failed to read file with: %v", err)
	}

	gen := struct {
		Import    []string
		Templates map[string]string
		Data      interface{}
		Main      string
	}{}

	if err := cfg.Unpack(&gen, cfgOpts...); err != nil {
		return fmt.Errorf("Parsing template file failed with %v", err)
	}

	var T *template.Template
	var defaultFuncs = template.FuncMap{
		"toLower":    strings.ToLower,
		"toUpper":    strings.ToUpper,
		"capitalize": strings.Title,
		"isnil": func(v interface{}) bool {
			return v == nil
		},
		"dict": makeDict,
		"invoke": func(name string, values ...interface{}) (string, error) {
			params, err := makeDict(values...)
			if err != nil {
				return "", err
			}

			var buf bytes.Buffer
			err = T.ExecuteTemplate(&buf, name, params)
			return buf.String(), err
		},
	}

	T, err = loadTemplates(template.New("").Funcs(defaultFuncs), gen.Import)
	if err != nil {
		return err
	}

	T, err = addTemplates(T, gen.Templates)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	header := fmt.Sprintf("// This file has been generated from '%v', do not edit\n", args[0])
	buf.WriteString(header)
	T = T.New("master")
	T, err = T.Parse(gen.Main)
	if err != nil {
		return fmt.Errorf("Parsing 'template' fields failed with %v", err)
	}

	if err := T.Execute(&buf, gen.Data); err != nil {
		return fmt.Errorf("executing template failed with %v", err)
	}

	content := buf.Bytes()
	if *format {
		content, err = imports.Process(*to, content, nil)
		if err != nil {
			return fmt.Errorf("Applying goimports failed with: %v", err)
		}
	}

	if *to != "" {
		return ioutil.WriteFile(*to, content, 0644)
	}

	_, err = os.Stdout.Write(content)
	return err
}

func loadTemplates(T *template.Template, files []string) (*template.Template, error) {
	for _, file := range files {
		gen := struct {
			Import    []string
			Templates map[string]string
		}{}

		cfg, err := yaml.NewConfigWithFile(file, cfgOpts...)
		if err != nil {
			return nil, fmt.Errorf("Failed to read template file %v: %v", file, err)
		}

		if err := cfg.Unpack(&gen, cfgOpts...); err != nil {
			return nil, fmt.Errorf("Parsing template file %v failed with %v", file, err)
		}

		T, err = loadTemplates(T, gen.Import)
		if err != nil {
			return nil, err
		}

		T, err = addTemplates(T, gen.Templates)
		if err != nil {
			return nil, err
		}
	}

	return T, nil
}

func addTemplates(T *template.Template, templates map[string]string) (*template.Template, error) {
	for name, content := range templates {
		var err error

		T = T.New(name)
		T, err = T.Parse(content)
		if err != nil {
			return nil, fmt.Errorf("failed to parse template %v: %v", name, err)
		}
	}

	return T, nil
}

func makeDict(values ...interface{}) (map[string]interface{}, error) {
	if len(values)%2 != 0 {
		return nil, errors.New("invalid dict call")
	}

	dict := make(map[string]interface{}, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, errors.New("dict keys must be strings")
		}
		dict[key] = values[i+1]
	}
	return dict, nil
}
