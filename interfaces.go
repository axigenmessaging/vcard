package vcard

/**
 * property interface
 */
 type IProperty interface {
	/**
	 * set the name of the property (the function must convert property uppercase)
	 */
	SetName(n string)

	/**
	 * get the name of the property
	 */
	GetName() string

	/**
	 * set cardinality
	 */
	SetCardinality(v string)

	/**
	 * get cardinality
	 */
	 GetCardinality() string

	/**
	 * allow multiple values
	 */
	 SetAllowMultipleValues(v bool)

	/**
	 * add a value to the property
	 * if property is single value => the old value will be rewritten
	 * if property is multivalue => the value will be appended
	 */
	AddValue(v IData)

	/**
	 * the old value will be rewritten
	 * if the property is single value, and v has multiple values, only the first one will be kept
	 */
	SetValue(v []IData)

	/**
	 * return the values
	 */
	GetValue() []IData

	/**
	 * return first value - usefull when you know the attribut has a single value
	 */
    GetFirstValue() IData

	/**
	 * add property parameter
	 */
	AddParameter(param IParameter)

	/**
	 * set parameters removing all existing
	 */
	SetParameters(param map[string]IParameter)

	/**
	 * return parameters list
	 */
	GetParameters() map[string]IParameter
}


/**
 * data interface for property values
 */
 type IData interface {
	Validate() bool
	GetType() string
	SetValue(s string)
	GetValue() string
	IsEmpty() bool
	GetString() string
}

/**
 * property parameter interface
 */
type IParameter interface {
	GetName() string
	SetName(n string)
	AddValue(v string)
	SetValue(v []string)
	GetValue() []string

	/**
	 * set if the partemerter is allowed to have multiple values
	 */
	SetAllowMultipleValues(b bool)

	/**
	 * check if the paramerter allow multiple values
	 */
	AllowMultipleValues() bool

	/**
	 * validate parameter values
	 */
	Validate() (bool, error)

	IsEmpty() bool
}

type IBuilder interface {
	/**
	* generate string from card stucture
	*/
	GetString() string

	/**
	 * return vcard version (the builder is specific for one version)
	 */
	GetVersion() string


	/**
	 * validate the properties and parameters
	 */
	Validate() bool

	/**
	 * add property
	 */

	AddProperty(p IProperty)

	/**
	* return a list of properties
	*/
	GetProperty(name string) []IProperty

	DeleteProperty(name string)


	/**
	 * scenario for rewriting / use properties with cardinality 1 or *1
	 *  ignore - fist added will be used -> if the property exists if you try to add it again, it will be ignored
	 *  overwrite - last added will be used -> last added property will overwrite the existing one
	 */
	SetAddPropertyScenario(v string)

	GetAddPropertyScenario() string

	Build() string


	NewBeginProperty() *VCardProperty
	NewEndProperty() *VCardProperty
	NewVersionProperty() *VCardProperty
	NewCustomProperty(name string) *VCardProperty
	NewFBUrlProperty() *VCardProperty
	NewClassProperty() *VCardProperty
	NewNProperty() *VCardProperty
	NewFnProperty() *VCardProperty
	NewBDayProperty() *VCardProperty
	NewAnniversaryProperty() *VCardProperty
	NewAdrProperty() *VCardProperty
	NewTelProperty() *VCardProperty
	NewEmailProperty() *VCardProperty
	NewNicknameProperty() *VCardProperty
	NewPhotoProperty() *VCardProperty
	NewUrlProperty() *VCardProperty
	NewKeyProperty() *VCardProperty
	NewSoundProperty() *VCardProperty
	NewUidProperty() *VCardProperty
	NewTzProperty() *VCardProperty
	NewTitleProperty() *VCardProperty
	NewRoleProperty() *VCardProperty
	NewLogoProperty() *VCardProperty
	NewOrgProperty() *VCardProperty
	NewCategoriesProperty() *VCardProperty
	NewNoteProperty() *VCardProperty
	NewProdIdProperty() *VCardProperty
	NewRevProperty() *VCardProperty
	NewGeoProperty() *VCardProperty
	NewSourceProperty() *VCardProperty
	NewLabelProperty() *VCardProperty
	NewCalAdrUriProperty() *VCardProperty
	NewCalUriProperty() *VCardProperty
	NewRelatedProperty() *VCardProperty
	NewAgentProperty() *VCardProperty
	NewSortStringProperty() *VCardProperty
	NewMailerProperty() *VCardProperty
	NewNameProperty() *VCardProperty
	NewProfileProperty() *VCardProperty
	NewXmlProperty() *VCardProperty
	NewKindProperty() *VCardProperty
	NewGenderProperty() *VCardProperty
	NewImppProperty() *VCardProperty
	NewLangProperty() *VCardProperty
	NewMemberProperty() *VCardProperty
	NewClientPidMappProperty() *VCardProperty
}
