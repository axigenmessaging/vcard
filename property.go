package vcard

import (
	"strings"
)



/**
 * generic vcard property
 */

type VCardProperty struct {
	// property name - should be uppercase
	name string

	// property values
	values []IData

	// property parameters
	parameters map[string]IParameter

	cardinality string

	allowMultipleValues bool

	acceptedValueTypes []string

}

/**
 * set accepted value types: see data-validators for accepted types;
 *
 */

func (p *VCardProperty) SetAcceptedValueTypes(acceptedTypes []string) {
	p.acceptedValueTypes = acceptedTypes
}

/**
*  set tag name upper case
*/
func (p *VCardProperty) SetName(n string) {
	p.name = strings.ToUpper(n);
}

func (p *VCardProperty) GetName() string {
	return p.name
}

func (p *VCardProperty) SetCardinality(v string) {
	switch (v) {
		case "1", "*1", "1*", "*":
			// do nothing - the value is correct
		default:
			v = "*"
	}
	p.cardinality = v
}

func (p *VCardProperty) GetCardinality() string {
	return p.cardinality
}

func (p *VCardProperty) SetAllowMultipleValues(v bool) {
	p.allowMultipleValues = v
}

func (p *VCardProperty) GetAllowMultipleValues() bool {
	return p.allowMultipleValues
}

/**
* add a value to property
*/
func (p *VCardProperty) AddValue(v IData) {
	p.values = append(p.values, v)
}

/**
 * get the list of values of a property
 */
func (p *VCardProperty) GetValue() []IData {
	return p.values
}

/**
 * get first value from a property
 */
func (p *VCardProperty) GetFirstValue() IData {
	if (len(p.values) > 0) {
		return p.values[0]
	}
	return nil
}

/**
 * set property list values, reseting existing ones
 */

func (p *VCardProperty) SetValue(values []IData)  {
	 p.values = values
}

/**
 * add a parameter to a property
 */
func (p *VCardProperty) AddParameter(param IParameter) {
   existingParam, ok := p.parameters[param.GetName()];
   if ok {
	 for _, nv := range param.GetValue() {
		 existingParam.AddValue(nv)
	 }
   } else {
	   if p.parameters == nil {
			p.parameters = map[string]IParameter{}
	   }
	   p.parameters[param.GetName()] = param
   }
}

/**
 * set the parameters list, reseting any existing parameter
 */
func (p *VCardProperty) SetParameters(paramList map[string]IParameter) {
   p.parameters = paramList
}

/**
 * get the property parameters list
 */
func (p *VCardProperty) GetParameters() map[string]IParameter {
	return p.parameters
}

/**
 *	create a generic property
 */

 func NewProperty(name string) *VCardProperty {
	if name == "" {
		return nil
	}

	p := &VCardProperty{}
	p.SetName(strings.ToUpper(name))

	return p
}