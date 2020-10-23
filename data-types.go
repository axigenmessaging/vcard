package vcard

import (
	"strings"
	"regexp"
	"net/http"
)

/*
"text": The "text" value type should be used to identify values that
contain human-readable text.  As for the language, it is controlled
by the LANGUAGE property parameter defined in Section 5.1.

Examples for "text":

	this is a text value
	this is one value,this is another
	this is a single value\, with a comma encoded

A formatted text line break in a text value type MUST be represented
as the character sequence backslash (U+005C) followed by a Latin
small letter n (U+006E) or a Latin capital letter N (U+004E), that
is, "\n" or "\N".


GetString returns the string of the value structure. If the is compound, it creates the string joing components by ";" and escape  ";" char from components' value
*/

func EscapeValue(v string) string {
	v = strings.ReplaceAll(v, "\\", "\\\\")
	v = strings.ReplaceAll(v, ",", "\\,")
	v = strings.ReplaceAll(v, "\\n", "\\\\n")

	// this is optional for value without componets
	v = strings.ReplaceAll(v, ";", "\\;")
	return v
}


type TextValue struct {
	value string
}

func (v *TextValue) Validate() bool {
	return true
}
func (v *TextValue) GetType() string {
	return "TEXT"
}
func (v *TextValue) SetValue(val string) {
	v.value = val
}
func (v *TextValue) GetValue() string {
	return v.value;
}
func (v *TextValue) IsEmpty() bool {
	return v.value == ""
}

func (v *TextValue) GetString() string {
	return EscapeValue(v.value)
}


func NewText(s string) *TextValue{
	return &TextValue {
		value: s,
	}
}


/**
 * The components correspond, in sequence, to the sex
 *     (biological), and gender identity.  Each component is optional.
 *
 *     Sex component:  A single letter.  M stands for "male", F stands
 *        for "female", O stands for "other", N stands for "none or not
 *        applicable", U stands for "unknown".
 *
 *     Gender identity component:  Free-form text.
 */
 type GenderValue struct {
	*TextValue
	Sex  string
	Identity string
}

func (v *GenderValue) GetType() string {
	return "GENDER"
}

func (v *GenderValue) Validate() bool {
	var result = false
	switch (v.Sex) {
		case "A", "M", "F", "O", "N", "U", "":
			result = true
	}
	return result
}

func (v *GenderValue) GetString() string {
	var s strings.Builder
	s.WriteString(EscapeValue(v.Sex))
	if len(v.Identity) > 0 {
		s.WriteString(";")
		s.WriteString(EscapeValue(v.Identity))
	}
	return s.String()
}

func NewGender(s string, i string) *GenderValue {
	return &GenderValue {
		Sex: s,
		Identity: i,
	}
}


/**
 * GEO
 */

 /**

			  geo-URI       = geo-scheme ":" geo-path
			  geo-scheme    = "geo"
			  geo-path      = coordinates p
			  coordinates   = coord-a "," coord-b [ "," coord-c ]

			  coord-a       = num
			  coord-b       = num
			  coord-c       = num

			  p             = [ crsp ] [ uncp ] *parameter
			  crsp          = ";crs=" crslabel
			  crslabel      = "wgs84" / labeltext
			  uncp          = ";u=" uval
			  uval          = pnum
			  parameter     = ";" pname [ "=" pvalue ]
			  pname         = labeltext
			  pvalue        = 1*paramchar
			  paramchar     = p-unreserved / unreserved / pct-encoded

			  labeltext     = 1*( alphanum / "-" )
			  pnum          = 1*DIGIT [ "." 1*DIGIT ]
			  num           = [ "-" ] pnum
			  unreserved    = alphanum / mark
			  mark          = "-" / "_" / "." / "!" / "~" / "*" /
							  "'" / "(" / ")"
			  pct-encoded   = "%" HEXDIG HEXDIG
			  p-unreserved  = "[" / "]" / ":" / "&" / "+" / "$"
			  alphanum      = ALPHA / DIGIT
  */
 type GeoValue struct {
	 *TextValue
	 Lat  string //;
	 Lon string
	 Alt string
	 P string
 }

 func (v *GeoValue) GetType() string {
	 return "GEO"
 }

 func (v *GeoValue) Validate() bool {
	 var result = false

	 return result
 }

 func (v *GeoValue) GetString() string {
	 var s strings.Builder

	 s.WriteString("geo:")
	 s.WriteString(EscapeValue(v.Lat))
	 s.WriteString(",")
	 s.WriteString(EscapeValue(v.Lon))
	 if len(v.Alt) > 0 {
		s.WriteString(",")
		s.WriteString(EscapeValue(v.Alt))
	 }

	 return s.String()
 }

 func NewGeo(lat string, lon string, alt string) *GeoValue {
	 return &GeoValue {
		 Lat: lat,
		 Lon: lon,
		 Alt: alt,
	 }
 }


 /**
 * ADDRESS
 */


type AddressValue struct {
	*TextValue
	Pobox  string //the post office box;
	Ext string // the extended address (e.g., apartment or suite number);
	Street string //the street address;
	Locality string //the locality (e.g., city);
	Region string //the region (e.g., state or province);
	PostalCode string //the postal code;
	Country string //the country name
}

func (v *AddressValue) GetType() string {
	return "ADDRESS"
}

func (v *AddressValue) Validate() bool {
	return true
}

func (v *AddressValue) IsEmpty() bool {
	return v.Pobox == "" && v.Ext == ""  && v.Street == "" && v.Locality == "" && v.Region == "" && v.PostalCode == "" && v.Country == ""
}

func (v *AddressValue) GetString() string {
	var s strings.Builder
	s.WriteString(EscapeValue(v.Pobox))
	s.WriteString(";")
	s.WriteString(EscapeValue(v.Ext))
	s.WriteString(";")
	s.WriteString(EscapeValue(v.Street))
	s.WriteString(";")
	s.WriteString(EscapeValue(v.Locality))
	s.WriteString(";")
	s.WriteString(EscapeValue(v.Region))
	s.WriteString(";")
	s.WriteString(EscapeValue(v.PostalCode))
	s.WriteString(";")
	s.WriteString(EscapeValue(v.Country))

	return s.String()
}

func NewAddress() *AddressValue {
	return &AddressValue {
		Pobox: "",
		Ext: "",
		Street: "",
		Locality: "",
		Region: "",
		PostalCode: "",
		Country: "",
	}
}


/**
 * NAME
 * <Family Name>; <Given Name>; <Middle Name>; <Honorific Prefixes>;
 * <Honorific Postfixes>
 **/

type NameValue struct {
	*TextValue
	FamilyName  []string //family names/surname;
	GivenName []string // given names
	MiddleName []string //Additional Names
	HonorificPrefixes []string
	HonorificSuffixes []string
}

func (v *NameValue) GetType() string {
	return "NAME"
}

func (v *NameValue) Validate() bool {
	return true
}

func (v *NameValue) AddFamilyName(n string) {
	if n != "" {
		v.FamilyName = append(v.FamilyName, n)
	}
}

func (v *NameValue) AddGivenName(n string) {
	if n != "" {
		v.GivenName = append(v.GivenName, n)
	}
}

func (v *NameValue) AddMiddleName(n string) {
	if n != "" {
		v.MiddleName = append(v.MiddleName, n)
	}
}

func (v *NameValue) AddHonorificPrefix(n string) {
	if n != "" {
		v.HonorificPrefixes = append(v.HonorificPrefixes, n)
	}
}

func (v *NameValue) AddHonorificSuffix(n string) {
	if n != "" {
		v.HonorificSuffixes = append(v.HonorificSuffixes, n)
	}
}

func (v *NameValue) GetString() string {
	var (
		s strings.Builder
		idx int
		vs string
	)

	for idx, vs = range v.FamilyName {
		if idx != 0 {
			s.WriteString(",")
		}
		s.WriteString(EscapeValue(vs))
	}
	s.WriteString(";")

	for idx, vs = range v.GivenName {
		if idx != 0 {
			s.WriteString(",")
		}
		s.WriteString(EscapeValue(vs))
	}
	s.WriteString(";")

	for idx, vs = range v.MiddleName {
		if idx != 0 {
			s.WriteString(",")
		}
		s.WriteString(EscapeValue(vs))
	}
	s.WriteString(";")

	for idx, vs = range v.HonorificPrefixes {
		if idx != 0 {
			s.WriteString(",")
		}
		s.WriteString(EscapeValue(vs))
	}
	s.WriteString(";")

	for idx, vs = range v.HonorificSuffixes {
		if idx != 0 {
			s.WriteString(",")
		}
		s.WriteString(EscapeValue(vs))
	}

	return s.String()
}

func (v *NameValue) IsEmpty() bool {
	return len(v.FamilyName) == 0 && len(v.GivenName) == 0 && len(v.MiddleName) == 0 && len(v.HonorificPrefixes) == 0 && len(v.HonorificSuffixes) == 0;
}

func NewName() *NameValue {
	v := &NameValue {}
	return v
}


/**
 * ORGANIZATION
 */


 /**
 * A single structured text value consisting of components
 *      separated by the SEMICOLON character (U+003B).
 *	   component *(";" component)
 *     [Organization];[SubUnit1];[SubUnit2]; <repeats>
 */
 type OrganizationValue struct {
	 *TextValue
	 Company  string //the post office box;
	 Departments []string // the extended address (e.g., apartment or suite number);
 }

 func (v *OrganizationValue) GetType() string {
	 return "ORG"
 }

 func (v *OrganizationValue) Validate() bool {
	 return true
 }

 func (v *OrganizationValue) IsEmpty() bool {
	 return v.Company == "" && len(v.Departments) == 0
 }

 func (v *OrganizationValue) GetString() string {
	var s strings.Builder

	s.WriteString(EscapeValue(v.Company))
	if len(v.Departments) > 0 {
		s.WriteString(";")
		for idx, vs := range v.Departments {
			if idx != 0 {
				s.WriteString(";")
			}
			s.WriteString(EscapeValue(vs))
		}
	}

	return s.String()
 }

 func NewOrganization(c string, d []string) *OrganizationValue {
	 return &OrganizationValue {
		 Company: c,
		 Departments: d,
	 }
}

 /**
 *  photo value
 */
 type PhotoValue struct {
	*TextValue
	IsUrl  bool // if the value is an url;
	IsB64Encoded bool // if the value is a base64 encoded string
	MediaType string // mime type of the value
}

func (v *PhotoValue) IsEmpty() bool {
	return v.value == ""
}

func (v *PhotoValue) GetType() string {
	return "PHOTO"
}

func (v *PhotoValue) Validate() bool {
	return true
}

/**
 *  detect if a value is a Data URI to decompose the uri,  an URI, or a simple file string
 */
func (v *PhotoValue) AutodetectValue(s string) {
	if (IsUri(s)) {
		// the value it's an url or something like data:image/jpeg;base64, xxxx
		re := regexp.MustCompile(`^data\:([^\:;]+)?(;base64)?,(.*)`)
		tmp := re.FindStringSubmatch(s)

		if len(tmp) == 0 {
			// it not a data URI
			v.IsUrl = true
			v.IsB64Encoded = false //IsBase64Encoded(s)
			v.SetValue(s)
			v.MediaType = "" //http.DetectContentType([]byte(s))
		} else {
			/*
				 tmp -> 0-> full match
						 1 -> media type
						 2 -> base64
						 3 -> file content
				it's possible that 1 and/or 2 to be missing
			*/
			// it's a data schema
			fileContent := tmp[len(tmp)-1]
			if len(tmp) > 2 {
				// we have at least media type or base64
				if strings.ToLower(tmp[1]) == ";base64" || (len(tmp)>3 && strings.ToLower(tmp[2]) == ";base64") {
					// media type is missing or all elements are present
					v.IsB64Encoded = true
				}
				if strings.ToLower(tmp[1]) != ";base64" {
					v.MediaType = tmp[1]
				}
			}
			if v.MediaType == "" {
				v.MediaType = http.DetectContentType([]byte(fileContent))
			}

			v.SetValue(fileContent)
		}
	} else {
		v.IsUrl = false
		v.IsB64Encoded = IsBase64Encoded(s)
		v.SetValue(s)
		v.MediaType = http.DetectContentType([]byte(s))
	}
}

func NewPhoto(s string) *PhotoValue {
	p := new(PhotoValue)
	if len(s)>0 {
		p.AutodetectValue(s)
	}
	return p
}