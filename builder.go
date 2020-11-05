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
	vcard IVCard
	/**
	 * possible values:
	 *	 - ignore - if a property with cardinality 1 or *1 is already set and is added for the second time, the new added is ignored (the first property is kept)
	 *	 - overwrite - if a property with cardinality 1 or *1 is already set and is added for the second time, the old one is overwritten.
	 *  default: overwrite
	 */
	 addPropertyScenario string

	// vcard string
	cardString strings.Builder
}



func (b *Builder) GetString() string {
	return b.cardString.String()
}


func (b *Builder) Build() string {
	// reset the string
	b.cardString.Reset() // b.cardString = strings.Builder{}

	// write begin property
	b.cardString.WriteString(b.RenderProperty(b.vcard.CreateProperty("begin")))
	b.cardString.WriteString("\r\n")

	// write version property
	b.cardString.WriteString(b.RenderProperty(b.vcard.CreateProperty("version")))
	b.cardString.WriteString("\r\n")

	for _, p := range b.vcard.GetProperties() {
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
	b.cardString.WriteString(b.RenderProperty(b.vcard.CreateProperty("end")))

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


func NewBuilder(vc IVCard) *Builder {
	b := Builder{
		vcard: vc,
	}

	return &b
}
