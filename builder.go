/**
 * vcard generic builder that must be inherit by all builders
 * all builders should have same methods
 * should contains common functions and creators for common properties for all versions (properties must have same characteristics)
 */
package vcard
import (
	"strings"
)

type Builder struct {
	/**
	 * possible values:
	 *	 - ignore - if a property with cardinality 1 or *1 is already set and is added for the second time, the new added is ignored (the first property is kept)
	 *	 - overwrite - if a property with cardinality 1 or *1 is already set and is added for the second time, the old one is overwritten.
	 *  default: overwrite
	 */
	 addPropertyScenario string

	// vcard prperty structure
	properties []IProperty

	// vcard string
	cardString strings.Builder

	version string
}


/**
* adding property scenario
* possible values:
*	 - ignore - if a property with cardinality 1 or *1 is already set and is added for the second time, the new added is ignored (the first property is kept)
*	 - overwrite - if a property with cardinality 1 or *1 is already set and is added for the second time, the old one is overwritten.
*  default: overwrite
*/
func (b *Builder) SetAddPropertyScenario(v string) {
	if v != "ignore" {
		v = "overwrite"
	}
	b.addPropertyScenario = v;
}

func (b *Builder) GetAddPropertyScenario() string {
	v := b.addPropertyScenario
	if v!= "ignore" {
		v = "overwrite"
	}
	return v
}


func (b *Builder) GetString() string {
	return b.cardString.String()
}

func (b *Builder) GetVersion() string {
	return b.version
}

func (b *Builder) Build() string {
	// reset the string
	b.cardString.Reset()

	// write begin property
	b.cardString.WriteString(b.RenderProperty(b.NewBeginProperty()))
	b.cardString.WriteString("\r\n")

	// write version property
	b.cardString.WriteString(b.RenderProperty(b.NewVersionProperty()))
	b.cardString.WriteString("\r\n")

	for _, p := range b.properties {
		switch p.GetName() {
			case "BEGIN", "END", "VERSION":
				// these properties are manually added in the correct order
				continue
			default:
				// render a property
				b.cardString.WriteString(b.RenderProperty(p))
				b.cardString.WriteString("\r\n")
		}
	}

	// write end property
	b.cardString.WriteString(b.RenderProperty(b.NewEndProperty()))

	return b.GetString()
}

/**
 * render a property
 */

func (b *Builder) RenderProperty(p IProperty) string {
	var s strings.Builder

	s.WriteString(p.GetName())
	s.WriteString(b.RenderParameters(p.GetParameters()))
	s.WriteString(":")
	s.WriteString(b.RenderPropertyValue(p))

	return FormatLine(s.String())
}

func (b *Builder) RenderPropertyValue(p IProperty) string {
	var s strings.Builder

	for idx, v := range p.GetValue() {
		if idx > 0 {
			s.WriteString(",")
		}
		s.WriteString(v.GetString())
	}

	return s.String()
}


/**
 * render property' parameters
 */
func (b *Builder) RenderParameters(parameters map[string]IParameter) string {
	var s strings.Builder
	for _, p := range parameters {
		s.WriteString(";")
		s.WriteString(b.RenderParameter(p))
	}
	return s.String()
}

/**
 * render a parameter
 */

func (b *Builder) RenderParameter(p IParameter) string {
	var (
		s strings.Builder
	)

	// @TODO: render parameters
	s.WriteString(p.GetName())
	s.WriteString("=")

	firstWrite := true

	for _, pv := range p.GetValue() {
		if strings.Contains(pv, "\"") {
			// param values must not contains "
			continue
		}

		if !firstWrite {
			s.WriteString(",")
		}
		firstWrite = false
		if strings.ContainsAny(pv, ":;,") {
			// values than contains :;, must be double quoted
			s.WriteString("\"")
			s.WriteString(pv)
			s.WriteString("\"")
		} else {
			s.WriteString(pv)
		}
	}

	return s.String()
}


/**
 * return a list of properties
 */
 func (b *Builder) GetProperty(name string) []IProperty {
    var result []IProperty
    for _ , p := range b.properties {
		if (p.GetName() == strings.ToUpper(name)) {
			result = append(result, p)
		}
	}
	return result
}

/**
 * add a property
 */
func (b *Builder) AddProperty(p IProperty) {

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
 * delete a proprty by name
 */
func (b *Builder) DeleteProperty(name string) {
	for idx, p := range b.properties {
		if (p.GetName() == strings.ToUpper(name)) {
			b.properties = append(b.properties[:idx], b.properties[idx+1:]...)
		}
	}
}

/**
 *  create BEGIN property
 */
func (b *Builder) NewBeginProperty() *VCardProperty {
	p := NewProperty("begin")
	p.SetCardinality("1")
	p.SetAllowMultipleValues(false)
	p.AddValue(NewText("VCARD"))

	return p
}

/**
 *  create END property
 */
 func (b *Builder) NewEndProperty() *VCardProperty {
	p := NewProperty("end")
	p.SetCardinality("1")
	p.SetAllowMultipleValues(false)
	p.AddValue(NewText("VCARD"))

	return p
}

/**
 *  create VERSION property
 */
 func (b *Builder) NewVersionProperty() *VCardProperty {
	p := NewProperty("version")
	p.SetCardinality("1")
	p.SetAllowMultipleValues(false)
	p.AddValue(NewText(b.GetVersion()))

	return p
}

/**
 *  create custom property
 */
 func (b *Builder) NewCustomProperty(name string) *VCardProperty {
	p := NewProperty(name)
	return p
}

/**
 * validate vcard structure, values & parameters
 */
func (b *Builder) Validate() bool {
	return true
}

/**
 *  Content lines  SHOULD be folded to a maximum width of 75 octets, excluding the line
 *  break.  Multi-octet characters MUST remain contiguous.
 * A logical line MAY be continued on the next physical line anywhere
 *  between two characters by inserting a CRLF immediately followed by a
 *  single white space character (space (U+0020) or horizontal tab
 *  (U+0009))
 */
func FormatLine(s string) string {
	var (
		rLine strings.Builder
		line strings.Builder
	)
	line.Grow(75)

	saveLine := false
	for i, r := range s {
		if line.Len() + len(string(r)) <= 75 {
			line.WriteRune(r)
		} else {
			saveLine = true
		}

		if saveLine || (i + len(string(r)) == len(s)) {
			// we have 75 chars or is the last char to save
			if rLine.Len() > 0 {
				rLine.WriteString("\r\n ")
			}
			rLine.Grow(len(line.String()))
			rLine.WriteString(line.String())

			saveLine = false
			line.Reset()
		}
	}
	return rLine.String()
}

func (b *Builder) NewFBUrlProperty() *VCardProperty {
	// not available in 3.0
	return nil
}

func (b *Builder) NewClassProperty() *VCardProperty {
	// not available in 4.0
	return nil
}

/**
 *  create N property
 */
 func (b *Builder) NewNProperty() *VCardProperty {
	p := NewProperty("N")
	p.SetCardinality("*1")
	p.SetAllowMultipleValues(false)
	return p
 }

 /**
 *  create FN property
 */
 func (b *Builder) NewFnProperty() *VCardProperty {
	p := NewProperty("FN")
	p.SetCardinality("1*")
	p.SetAllowMultipleValues(false)
	return p
 }


/**
 *  create BDAY property
 */
 func (b *Builder) NewBDayProperty() *VCardProperty {
	p := NewProperty("BDAY")
	p.SetCardinality("*1")
	p.SetAllowMultipleValues(false)
	return p
 }

/**
 *  create ANNIVERSARY property
 */
 func (b *Builder) NewAnniversaryProperty() *VCardProperty {
	p := NewProperty("ANNIVERSARY")
	p.SetCardinality("*1")
	p.SetAllowMultipleValues(false)
	return p
 }

 /**
 *  create ADR property
 */
 func (b *Builder) NewAdrProperty() *VCardProperty {
	p := NewProperty("ADR")
	p.SetCardinality("*")
	p.SetAllowMultipleValues(false)
	return p
 }

 /**
 *  create TEL property
 */
func (b *Builder) NewTelProperty() *VCardProperty {
	p := NewProperty("TEL")
	p.SetCardinality("*")
	p.SetAllowMultipleValues(false)
	return p
}

/**
 *  create EMAIL property
 */
 func (b *Builder) NewEmailProperty() *VCardProperty {
	p := NewProperty("EMAIL")
	p.SetCardinality("*")
	p.SetAllowMultipleValues(false)
	return p
}

 /**
 *  create NICKNAME property
 */
 func (b *Builder) NewNicknameProperty() *VCardProperty {
	p := NewProperty("NICKNAME")
	p.SetCardinality("*")
	p.SetAllowMultipleValues(true)
	return p
 }

/**
 *  create PHOTO property
 */
 func (b *Builder) NewPhotoProperty() *VCardProperty {
	p := NewProperty("PHOTO")
	p.SetCardinality("*")
	p.SetAllowMultipleValues(false)
	return p
 }


/**
 *  create URL property
 */
 func (b *Builder) NewUrlProperty() *VCardProperty {
	p := NewProperty("URL")
	p.SetCardinality("*")
	p.SetAllowMultipleValues(false)
	return p
}

/**
 *  create KEY property
 */
 func (b *Builder) NewKeyProperty() *VCardProperty {
	p := NewProperty("KEY")
	p.SetCardinality("*")
	p.SetAllowMultipleValues(false)
	return p
}


/**
 *  create SOUND property
 */
 func (b *Builder) NewSoundProperty() *VCardProperty {
	p := NewProperty("SOUND")
	p.SetCardinality("*")
	p.SetAllowMultipleValues(false)
	return p
}


/**
 *  create UID property
 */
 func (b *Builder) NewUidProperty() *VCardProperty {
	p := NewProperty("UID")
	p.SetCardinality("*1")
	p.SetAllowMultipleValues(false)
	return p
}

/**
 *  create TZ property
 */
 func (b *Builder) NewTzProperty() *VCardProperty {
	p := NewProperty("TZ")
	p.SetCardinality("*")
	p.SetAllowMultipleValues(false)
	return p
}


/**
 *  create TITLE property
 */
 func (b *Builder) NewTitleProperty() *VCardProperty {
	p := NewProperty("TITLE")
	p.SetCardinality("*")
	p.SetAllowMultipleValues(false)
	return p
}

/**
 *  create ROLE property
 */
 func (b *Builder) NewRoleProperty() *VCardProperty {
	p := NewProperty("ROLE")
	p.SetCardinality("*")
	p.SetAllowMultipleValues(false)
	return p
}


/**
 *  create LOGO property
 */
 func (b *Builder) NewLogoProperty() *VCardProperty {
	p := NewProperty("LOGO")
	p.SetCardinality("*")
	p.SetAllowMultipleValues(false)
	return p
}

/**
 *  create ORG property
 */
 func (b *Builder) NewOrgProperty() *VCardProperty {
	p := NewProperty("ORG")
	p.SetCardinality("*")
	p.SetAllowMultipleValues(false)
	return p
}


/**
 *  create CATEGORIES property
 */
 func (b *Builder) NewCategoriesProperty() *VCardProperty {
	p := NewProperty("CATEGORIES")
	p.SetCardinality("*")
	p.SetAllowMultipleValues(true)
	return p
}

/**
 *  create NOTE property
 */
 func (b *Builder) NewNoteProperty() *VCardProperty {
	p := NewProperty("NOTE")
	p.SetCardinality("*")
	p.SetAllowMultipleValues(false)
	return p
}

/**
 *  create PRODID property
 */
 func (b *Builder) NewProdIdProperty() *VCardProperty {
	p := NewProperty("PRODID")
	p.SetCardinality("*1")
	p.SetAllowMultipleValues(false)
	return p
}

/**
 *  create REV property
 */
 func (b *Builder) NewRevProperty() *VCardProperty {
	p := NewProperty("REV")
	p.SetCardinality("*1")
	p.SetAllowMultipleValues(false)
	return p
}

/**
 *  create GEO property
 */
 func (b *BuilderV3) NewGeoProperty() *VCardProperty {
	p := NewProperty("GEO")
	p.SetCardinality("*")
	p.SetAllowMultipleValues(false)
	return p
}

/**
 *  create SOURCE property
 */
 func (b *Builder) NewSourceProperty() *VCardProperty {
	p := NewProperty("SOURCE")
	p.SetCardinality("*")
	p.SetAllowMultipleValues(false)
	return p
}

func (b *Builder) NewLabelProperty() *VCardProperty {
	return nil
}
func (b *Builder) NewCalAdrUriProperty() *VCardProperty {
	return nil
}

/**
 *  create CALURI property
 */
func (b *Builder) NewCalUriProperty() *VCardProperty {
	return nil
}

func (b *Builder) NewRelatedProperty() *VCardProperty {
	return nil
}

func (b *Builder) NewAgentProperty() *VCardProperty {
	return nil
}

func (b *Builder) NewSortStringProperty() *VCardProperty {
	return nil
}

func (b *Builder) NewMailerProperty() *VCardProperty {
	return nil
}

func (b *Builder) NewNameProperty() *VCardProperty {
	return nil
}
func (b *Builder) NewProfileProperty() *VCardProperty {
	return nil
}
func (b *Builder) NewXmlProperty() *VCardProperty {
	return nil
}
func (b *Builder) NewKindProperty() *VCardProperty {
	return nil
}
func (b *Builder) NewGenderProperty() *VCardProperty {
	return nil
}
func (b *Builder) NewImppProperty() *VCardProperty {
	return nil
}
func (b *Builder) NewLangProperty() *VCardProperty {
	return nil
}
func (b *Builder) NewMemberProperty() *VCardProperty {
	return nil
}
func (b *Builder) NewClientPidMappProperty() *VCardProperty {
	return nil
}