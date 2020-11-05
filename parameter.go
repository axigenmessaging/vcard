/**
* define the parameter of the vcard's property
*/
package vcard

import (
	"strings"
)


type Parameter struct {
	name string
	value []string
	mayHaveMultipleValues bool
}

func (p *Parameter) SetName(n string) {
	p.name = strings.ToUpper(n)
}

func (p *Parameter) GetName() string {
	return p.name
}

func (p *Parameter) SetAllowMultipleValues(b bool) {
	p.mayHaveMultipleValues = b
}

func (p *Parameter) AllowMultipleValues() bool {
	return p.mayHaveMultipleValues
}


func (p *Parameter) Validate() (result bool, err error) {
	result = true
	return
}

/**
* append a new value to a parameter
*/
func (p *Parameter) AddValue(value string) {
	for _, elValue := range p.value {
		if (elValue == value) {
			// the parameter already has this value
			return
		}
	}
	p.value = append(p.value, value)
}

/**
* set parameter values, overwriting all values
*/
func (p *Parameter) SetValue(value []string) {
	p.value = value
}

/**
* return parameter values
*/
func (p *Parameter) GetValue() []string {
	return p.value
}

/*
 * param        = param-name "=" param-value *("," param-value)
 */
func (p *Parameter) String() string {
	var strBuilder strings.Builder
	if len(p.value) > 0 {
		strBuilder.WriteString(";")
		strBuilder.WriteString(p.name)
		strBuilder.WriteString("=")
		for i, elValue := range p.value {
			if (i != 0) {
				strBuilder.WriteString(",")
			}
			// be sure it has no " chars
			elValue = strings.ReplaceAll(elValue, "\"", "")

			// is has special chars, enclose quote the value
			if (strings.ContainsAny(elValue, ":;,")) {
				elValue = "\""  +elValue + "\""
			}
			strBuilder.WriteString(elValue)
		}
	}
	return strBuilder.String()
}

/**
 * has at least one value that is not empty
 */
func (p *Parameter) IsEmpty() bool {
	if len(p.value) > 0 {
		for _, v := range p.value {
			if v != "" {
				return false
			}
		}
	}
	return true
}

func NewParameter(name string) *Parameter {
	p := &Parameter {
		name: "",
		value: []string{},
		mayHaveMultipleValues: false,
	}
	p.SetName(name)
	return p
}