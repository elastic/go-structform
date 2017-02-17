package main

import (
	"bufio"
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
	dataFile := flag.String("d", "", "input data file for use to fill out")

	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		return errors.New("Missing input file")
	}

	userData, err := loadData(*dataFile)
	if err != nil {
		return fmt.Errorf("Failed to read data file: %v", err)
	}

	gen := struct {
		Import    []string
		Templates map[string]string
		Main      string
	}{}
	if err = loadConfigInto(args[0], &gen); err != nil {
		return err
	}

	dat := struct {
		Data interface{}
	}{}
	if err = loadConfigInto(args[0], &dat, ucfg.VarExp); err != nil {
		return err
	}

	var T *template.Template
	var defaultFuncs = template.FuncMap{
		"toLower":    strings.ToLower,
		"toUpper":    strings.ToUpper,
		"capitalize": strings.Title,
		"isnil": func(v interface{}) bool {
			return v == nil
		},
		"dict":   makeDict,
		"invoke": makeInvokeCommand(&T), // invoke another template with named parameters
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

	var data interface{}
	switch {
	case userData == nil && dat.Data != nil:
		data = dat.Data
	case userData != nil && dat.Data == nil:
		data = userData
	case userData != nil && dat.Data != nil:
		data = map[string]interface{}{
			"config":   userData,
			"template": dat.Data,
		}
	}

	if err := T.Execute(&buf, data); err != nil {
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

		err := loadConfigInto(file, &gen)
		if err != nil {
			return nil, err
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

func loadConfig(file string, extraOpts ...ucfg.Option) (cfg *ucfg.Config, err error) {
	opts := append(append([]ucfg.Option{}, extraOpts...), cfgOpts...)
	cfg, err = yaml.NewConfigWithFile(file, opts...)
	if err != nil {
		err = fmt.Errorf("Failed to read file %v with: %v", file, err)
	}
	return
}

func loadConfigInto(file string, to interface{}, extraOpts ...ucfg.Option) error {
	cfg, err := loadConfig(file, extraOpts...)
	if err == nil {
		err = readConfig(cfg, to, extraOpts...)
	}
	return err
}

func readConfig(cfg *ucfg.Config, to interface{}, extraOpts ...ucfg.Option) error {
	opts := append(append([]ucfg.Option{}, extraOpts...), cfgOpts...)
	if err := cfg.Unpack(to, opts...); err != nil {
		return fmt.Errorf("Parsing template file failed with %v", err)
	}
	return nil
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

func makeInvokeCommand(T **template.Template) func(string, ...interface{}) (string, error) {
	return func(name string, values ...interface{}) (string, error) {
		params, err := makeDict(values...)
		if err != nil {
			return "", err
		}

		var buf bytes.Buffer
		err = (*T).ExecuteTemplate(&buf, name, params)
		return buf.String(), err

	}
}

func loadData(file string) (map[string]string, error) {
	if file == "" {
		return nil, nil
	}

	meta := struct {
		Entries map[string]struct {
			Default     string
			Description string
		} `config:",inline"`
	}{}

	err := loadConfigInto(file, &meta, ucfg.VarExp)
	if err != nil {
		return nil, err
	}

	reader := bufio.NewReader(os.Stdin)

	state := map[string]string{}
	for name, entry := range meta.Entries {
		// parse default value
		T, err := template.New("").Parse(entry.Default)
		if err != nil {
			return nil, fmt.Errorf("Failed to parse data entry %v: %v", name, err)
		}

		var buf bytes.Buffer
		if err := T.Execute(&buf, state); err != nil {
			return nil, fmt.Errorf("Failed to evaluate data entry %v: %v", name, err)
		}

		// ask user for input
		defaultValue := buf.String()
		fmt.Printf("%v\n%v [%v]: ", entry.Description, name, defaultValue)
		value, err := reader.ReadString('\n')
		if err != nil {
			return nil, fmt.Errorf("Error waiting for user input: %v", err)
		}

		value = strings.TrimSpace(value)
		if value == "" {
			value = defaultValue
		}

		state[name] = value
	}

	return state, nil
}
