/**
 * vcard builder V4
 */
package vcard

type BuilderV3 struct {
	*Builder
}

/**
 *  create NAME property
 */
 func (b *BuilderV3) NewNameProperty() *VCardProperty {
	p := NewProperty("NAME")
	p.SetCardinality("*1")
	p.SetAllowMultipleValues(false)
	return p
}

/**
 *  create PROFILE property
 */
 func (b *BuilderV3) NewProfileProperty() *VCardProperty {
	p := NewProperty("PROFILE")
	p.SetCardinality("*1")
	p.SetAllowMultipleValues(false)
	return p
}



 /**
 *  create LABEL property
 */
 func (b *BuilderV3) NewLabelProperty() *VCardProperty {
	p := NewProperty("LABEL")
	p.SetCardinality("*")
	p.SetAllowMultipleValues(false)
	return p
 }



/**
 *  create MAILER property
 */
func (b *BuilderV3) NewMailerProperty() *VCardProperty {
	p := NewProperty("MAILER")
	p.SetCardinality("*1")
	p.SetAllowMultipleValues(false)
	return p
}


/**
 *  create AGENT property
 */
 func (b *BuilderV3) NewAgentProperty() *VCardProperty {
	p := NewProperty("AGENT")
	p.SetCardinality("*")
	p.SetAllowMultipleValues(false)
	return p
}


/**
 *  create SORT-STRING property
 */
 func (b *BuilderV3) NewSortStringProperty() *VCardProperty {
	p := NewProperty("SORT-STRING")
	p.SetCardinality("*")
	p.SetAllowMultipleValues(false)
	return p
}



/**
 *  create CLASS property
 */
 func (b *BuilderV3) NewClassProperty() *VCardProperty {
	p := NewProperty("CLASS")
	p.SetCardinality("*")
	p.SetAllowMultipleValues(false)
	return p
}


/**
 *  validate properties/values/parameters
 *  @TODO
 */

 func (b *BuilderV3) Validate() bool {
	 return true
 }

/**
* create a new v4.0 card builder
*/
func NewBuilderV3() BuilderV3 {
	b := BuilderV3 {
			&Builder{
				version: "3.0",
			},
	}
	return b
}

