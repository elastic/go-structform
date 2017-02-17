package gotype

// define mktmpl alias
//go:generate -command mktmpl go run ../internal/gen/gen_yaml.go

//go:generate mktmpl -f -o stacks_generated.go stacks_tmpl.yml
//go:generate mktmpl -f -o unfold_primitive_generated.go unfold_primitive.yml
//go:generate mktmpl -f -o unfold_arr_generated.go unfold_arr.yml
//go:generate mktmpl -f -o unfold_map_generated.go unfold_map.yml
//go:generate mktmpl -f -o unfold_sel_generated.go unfold_sel.yml
