package vcard

import (
	"strings"
	"net/url"
	"strconv"
	"regexp"
	"encoding/base64"
)



/**
 * dataFormat:
 * URI, EMAIL, DATE, TIME, DATE-TIME, DATE-AND-OR-TIME, TIMESTAMP, BOOLEAN  , INTEGER , FLOAT , UTC-OFFSET, LANGUAGE-TAG, ADDRESS, BOOLEAN, INTEGER, FLOAT
 *
 */

func ValidateData(d IData, dataFormat string) bool {
	switch (strings.ToUpper(d.GetType())) {
		case "TEXT":
			switch (strings.ToUpper(dataFormat)) {
				case "URI":
					return IsUri(d.GetValue())
				case "EMAIL":
					return IsEmail(d.GetValue())
				case "BOOLEAN":
					return IsBoolean(d.GetValue())
				case "INTEGER":
					return IsInteger(d.GetValue())
				case "FLOAT":
					return IsFloat(d.GetValue())
				case "DATE":
					return IsDate(d.GetValue())
				case "TIME":
					return IsTime(d.GetValue())
				case "DATE-TIME":
					return IsDatetime(d.GetValue())
				case "DATE-AND-OR-TIME":
					return IsDatetime(d.GetValue()) || IsDate(d.GetValue()) || (strings.HasPrefix(d.GetValue(), "T") && IsTime(strings.TrimPrefix(d.GetValue(), "T")))
				case "UTC-OFFSET":
					return IsUtcOffset(d.GetValue())
				case "TIMESTAMP":
					return IsTimestamp(d.GetValue())
			}
			return true
		default:
			if strings.ToUpper(d.GetType()) == strings.ToUpper(dataFormat) {
				return true
			}
			return false
	}
}

func IsUri(s string) bool {
	_, err := url.ParseRequestURI(s)
	if err != nil {
		return false
	}
	return true
}

func IsEmail(s string) bool {
	if len(s) < 3 && len(s) > 254 {
		return false
	}
	matched, _ := regexp.MatchString("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$", s)

	return matched
}


func IsBoolean(v string) bool {
	val := strings.ToUpper(v)
	if val == "FALSE" || val == "TRUE" {
		return true
	}
	return false
}

func IsInteger(s string) bool {
	_, err := strconv.ParseInt(s, 10, 64)
	if (err == nil) {
		return true
	}
	return false
}

func IsFloat(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	if (err == nil) {
		return true
	}
	return false
}

func IsTimestamp(s string) bool {
	matched, _ := regexp.MatchString(`^[0-9]{4}[0-9]{2}[0-9]{2}T[0-9]{6}(Z|[\-+][0-9]{2,4})?$`, s)
	return matched
}

/**
 *	date          = year    [month  day]
 *                  / year "-" month
 *                  / "--"     month [day]
 *                  / "--"      "-"   day
 *    date-noreduc  = year     month  day
 *                  / "--"     month  day
 *                  / "--"      "-"   day
 *	 date-complete = year     month  day
*/
func IsDate(s string) bool {
	matched, _ := regexp.MatchString(`^([0-9]{4}|\-\-)([0-9]{2}|\-){0,2}$`, s)
	return matched
}


/**
 time          = hour [minute [second]] [zone]
                   /  "-"  minute [second]  [zone]
                   /  "-"   "-"    second   [zone]
     time-notrunc  = hour [minute [second]] [zone]
     time-complete = hour  minute  second   [zone]
     time-designator = %x54  ; uppercase "T"
*/
func IsTime(s string) bool {
	matched, _ := regexp.MatchString(`^([0-9]{2}|\-){1,3}Z?([\-+][0-9]{2,4}){0,1}$`, s)
	return matched
}


func IsDatetime(s string) bool {
	matched, _ := regexp.MatchString(`^([0-9]{4}|\-\-)([0-9]{2}|\-){2}T[0-9]{2}([0-9]{2}){0,2}$`, s)
	return matched
}

func IsUtcOffset(s string) bool {
	matched, _ := regexp.MatchString(`^[+\-][0-9]{2}([0-9]{2})?$`, s)
	return matched
}

func IsBase64Encoded(s string) bool {
	_, err := base64.StdEncoding.DecodeString(s)
	return err == nil
}