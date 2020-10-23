/**
 * vcard builder V4
 */
package vcard

type BuilderV4 struct {
	*Builder
}



/**
 *  create KIND property
 */
 func (b *BuilderV4) NewKindProperty() *VCardProperty {
	p := NewProperty("KIND")
	p.SetCardinality("*1")
	p.SetAllowMultipleValues(false)
	return p
}

/**
 *  create XML property
 */
 func (b *BuilderV4) NewXmlProperty() *VCardProperty {
	p := NewProperty("XML")
	p.SetCardinality("*")
	p.SetAllowMultipleValues(false)
	return p
}


/**
 *  create GENDER property
 */
 func (b *BuilderV4) NewGenderProperty() *VCardProperty {
	p := NewProperty("GENDER")
	p.SetCardinality("*1")
	p.SetAllowMultipleValues(false)
	return p
 }


/**
 *  create IMPP property
 */
func (b *BuilderV4) NewImppProperty() *VCardProperty {
	p := NewProperty("IMPP")
	p.SetCardinality("*")
	p.SetAllowMultipleValues(false)
	return p
}

/**
 *  create LANG property
 */
 func (b *BuilderV4) NewLangProperty() *VCardProperty {
	p := NewProperty("LANG")
	p.SetCardinality("*")
	p.SetAllowMultipleValues(false)
	return p
}



/**
 *  create MEMBER property
 */
 func (b *BuilderV4) NewMemberProperty() *VCardProperty {
	p := NewProperty("MEMBER")
	p.SetCardinality("*")
	p.SetAllowMultipleValues(false)
	return p
}

/**
 *  create RELATED property
 */
 func (b *BuilderV4) NewRelatedProperty() *VCardProperty {
	p := NewProperty("RELATED")
	p.SetCardinality("*")
	p.SetAllowMultipleValues(false)
	return p
}


/**
 *  create CLIENTPIDMAP property
 */
 func (b *BuilderV4) NewClientPidMappProperty() *VCardProperty {
	p := NewProperty("CLIENTPIDMAP")
	p.SetCardinality("*")
	p.SetAllowMultipleValues(false)
	return p
}


/**
 *  create FBURL property
 */
 func (b *BuilderV4) NewFBUrlProperty() *VCardProperty {
	p := NewProperty("FBURL")
	p.SetCardinality("*")
	p.SetAllowMultipleValues(false)
	return p
}

/**
 *  create CALADRURI property
 */
 func (b *BuilderV4) NewCalAdrUriProperty() *VCardProperty {
	p := NewProperty("CALADRURI")
	p.SetCardinality("*")
	p.SetAllowMultipleValues(false)
	return p
}

/**
 *  create CALURI property
 */
func (b *BuilderV4) NewCalUriProperty() *VCardProperty {
	p := NewProperty("CALURI")
	p.SetCardinality("*")
	p.SetAllowMultipleValues(false)
	return p
}

/**
 *  validate properties/values/parameters
 *  @TODO
 */

 func (b *BuilderV4) Validate() bool {
	 return true
 }

/**
* create a new v4.0 card builder
*/
func NewBuilderV4() BuilderV4 {
	b := BuilderV4{
		&Builder{
			version: "4.0",
		},
	}
	return b
}
