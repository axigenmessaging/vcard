package vcard

import (
	"strings"
)

type VCardV3 struct {
	properties []IProperty
	/**
	 * possible values:
	 *	 - ignore - if a property with cardinality 1 or *1 is already set and is added for the second time, the new added is ignored (the first property is kept)
	 *	 - overwrite - if a property with cardinality 1 or *1 is already set and is added for the second time, the old one is overwritten.
	 *  default: overwrite
	 */
	 addPropertyScenario string
}

func (b *VCardV3) SetAddPropertyScenario(v string) {
	if v != "ignore" {
		v = "overwrite"
	}
	b.addPropertyScenario = v;
}

func (b *VCardV3) GetAddPropertyScenario() string {
	v := b.addPropertyScenario
	if v!= "ignore" {
		v = "overwrite"
	}
	return v
}

func (vc *VCardV3) GetProperties() []IProperty {
	return vc.properties
}

/**
 * create property
 */
func (vc *VCardV3) CreateProperty(name string) IProperty {
	p := NewProperty(name)

	switch (p.GetName()) {
		case "BEGIN", "END":
			p.SetCardinality("1")
			p.SetAllowMultipleValues(false)
			p.AddValue(NewText("VCARD"))
		case "PROFILE":
			p.SetCardinality("*1")
			p.SetAllowMultipleValues(false)
			p.AddValue(NewText("VCARD"))
		case "VERSION":
			p.SetCardinality("1")
			p.SetAllowMultipleValues(false)
			p.AddValue(NewText("3.0"))
		case "FN", "N":
			p.SetCardinality("1")
			p.SetAllowMultipleValues(false)
		case "NICKNAME":
			p.SetCardinality("*")
			p.SetAllowMultipleValues(true)
		case "PHOTO","ADR","LABEL","TEL","EMAIL","GEO", "TITLE", "ROLE", "LOGO", "AGENT", "ORG", "CATEGORIES", "NOTE", "PRODID", "SORT-STRING", "SOUND", "KEY":
			p.SetCardinality("*")
			p.SetAllowMultipleValues(false)
		case "BDAY", "MAILER","TZ", "REV", "SOURCE", "UID", "CLASS":
			p.SetCardinality("*1")
			p.SetAllowMultipleValues(false)
		default:
			p.SetCardinality("*")
			p.SetAllowMultipleValues(false)
	}

	return p
}

/**
 * add a property
 */
 func (b *VCardV3) AddProperty(p IProperty) {

	if p.GetCardinality() == "1" || p.GetCardinality() == "*1" {
		// only one property should exists
		switch (b.GetAddPropertyScenario()) {
			case "ignore":
				propExists := b.GetProperty(p.GetName())
				if len(propExists) > 0 {
					// ignore item
					return
				}
			case "overwrite":
				// remove all existing properties of the same type (name)
				b.DeleteProperty(p.GetName())
		}
	}

   b.properties = append(b.properties, p)
}

/**
 * return a list of properties
 */
 func (b *VCardV3) GetProperty(name string) []IProperty {
    var result []IProperty
    for _ , p := range b.properties {
		if (p.GetName() == strings.ToUpper(name)) {
			result = append(result, p)
		}
	}
	return result
}

/**
 * delete a proprty by name
 */
 func (b *VCardV3) DeleteProperty(name string) {
	for idx, p := range b.properties {
		if (p.GetName() == strings.ToUpper(name)) {
			b.properties = append(b.properties[:idx], b.properties[idx+1:]...)
		}
	}
}



/**
 * create a parameter and attachit to a property
 * parameters & parameters values are specific to properties
 */
func (vc *VCardV3) AddPropertyParameter(p IProperty, name string, value []string) {
	param := NewParameter(name)

	switch (p.GetName()) {
		case "ENCODING":
			if len(value) > 0 {
				if (strings.ToLower(value[0]) == "base64") {
					value[0] = "b"
				}
			}
		case "VALUE":
		case "CHARSET":
		case "LANGUAGE":
		case "CONTEXT":
	}

	param.SetValue(value)

	p.AddParameter(param)
}

func (vc *VCardV3) Build() string {
	builder := NewBuilder(vc)
	return builder.Build()
}


func NewVCardV3() *VCardV3{
	v := VCardV3{}
	v.SetAddPropertyScenario("overwrite")
	return &v
}

