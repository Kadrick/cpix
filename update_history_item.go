package cpix

import "encoding/xml"

type UpdateHistoryItem struct {
	XMLName xml.Name `xml:"UpdateHistoryItem"`

	// Attr
	Id            xml.Attr `xml:"id,attr,omitempty"`
	UpdateVersion xml.Attr `xml:"updateVersion,attr"` // Required
	Index         xml.Attr `xml:"index,attr"`         // Required
	Source        xml.Attr `xml:"source,attr"`        // Required
	Date          xml.Attr `xml:"date,attr"`          // Required
}

type UpdateHistoryItemList struct {
	XMLName xml.Name `xml:"UpdateHistoryItemList,omitempty"`

	// Attr
	Id            xml.Attr `xml:"id,attr,omitempty"`
	UpdateVersion xml.Attr `xml:"updateVersion,attr,omitempty"`

	// Child
	List []*UpdateHistoryItem `xml:"UpdateHistoryItem"` // 0...n
}
