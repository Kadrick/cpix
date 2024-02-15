package cpix

import "encoding/xml"

type ContentKeyPeriod struct {
	XMLName xml.Name `xml:"ContentKeyPeriod"`

	// Attr
	Id xml.Attr `xml:"id,attr,omitempty"`

	Index xml.Attr `xml:"index,attr,omitempty"`
	Start xml.Attr `xml:"start,attr,omitempty"`
	End   xml.Attr `xml:"end,attr,omitempty"`
}

type ContentKeyPeriodList struct {
	XMLName xml.Name `xml:"ContentKeyPeriodList,omitempty"`

	// Attr
	Id            xml.Attr `xml:"id,attr,omitempty"`
	UpdateVersion xml.Attr `xml:"updateVersion,attr,omitempty"`

	// Child
	List []*ContentKeyPeriod `xml:"ContentKeyPeriod"` // 0...n
}
