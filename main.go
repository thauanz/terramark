package main

import (
	"fmt"
	"github.com/hashicorp/terraform-config-inspect/tfconfig"
	"os"
	"sort"
)

func main() {
	arg := os.Args[1]

	module, _ := tfconfig.LoadModule(arg)
	printVariables(module)
	printOutputs(module)
}

func printVariables(m *tfconfig.Module) {
	variables := m.Variables
	keys := make([]string, 0, len(variables))

	for k, _ := range variables {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var table string

	table = "| Name | Type | Description | Default Value | Required |\n"
	table = table + "| --- | --- | --- | :---: | :---: |\n"

	for _, k := range keys {
		v := variables[k]
		varType := v.Type
		if varType == "" {
			varType = "string"
		}

		var varDefault interface{}
		varRequired := "no"

		switch v.Default.(type) {
		case nil:
			varDefault = ""
			varRequired = "yes"
		case interface{}:
			varDefault = v.Default
		}

		if varDefault == nil {
			varRequired = "yes"
		}
		row := fmt.Sprintf("| %s | %s | %s | %v | %s |\n", v.Name, varType, v.Description, varDefault, varRequired)
		table = table + row
	}

	fmt.Printf("## Inputs\n\n")
	fmt.Println(table)
}

func printOutputs(m *tfconfig.Module) {
	outputs := m.Outputs
	keys := make([]string, 0, len(outputs))

	for k, _ := range outputs {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var table string
	table = "| Name | Description |\n"
	table = table + "| --- | --- |\n"

	for _, k := range keys {
		o := outputs[k]
		row := fmt.Sprintf("| %s | %s |\n", o.Name, o.Description)
		table = table + row
	}

	fmt.Printf("## Outputs\n\n")
	fmt.Println(table)
}
