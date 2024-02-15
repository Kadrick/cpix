package cpix

import "encoding/xml"

type DeliveryData struct {
	XMLName xml.Name `xml:"DeliveryData"`

	// Attr
	Id            xml.Attr `xml:"id,attr,omitempty"`
	Name          xml.Attr `xml:"name,attr,omitempty"`
	UpdateVersion xml.Attr `xml:"updateVersion,attr,omitempty"`

	// Child

	// TODO
	DeliveryKey struct {
		XMLName xml.Name `xml:"DeliveryKey"`
	} // Required

	// TODO
	DocumentKey struct {
		XMLName xml.Name `xml:"DocumentKey"`
	} // Required

	// TODO
	MACMethod struct {
		XMLName xml.Name `xml:"MACMethod"`
	} // 0...1

	Description struct {
		XMLName xml.Name `xml:"Description"`
		Value   string   `xml:",chardata"`
	} // 0...1

	SendingEntity struct {
		XMLName xml.Name `xml:"SendingEntity"`
		Value   string   `xml:",chardata"`
	} // 0...1

	SenderPointOfContact struct {
		XMLName xml.Name `xml:"SenderPointOfContact"`
		Value   string   `xml:",chardata"`
	} // 0...1

	ReceivingEntity struct {
		XMLName xml.Name `xml:"ReceivingEntity"`
		Value   string   `xml:",chardata"`
	} // 0...1
}

type DeliveryDataList struct {
	XMLName xml.Name `xml:"DeliveryDataList,omitempty"`

	// Attr
	Id            xml.Attr `xml:"id,attr,omitempty"`
	UpdateVersion xml.Attr `xml:"updateVersion,attr,omitempty"`

	// Child
	List []*DeliveryData `xml:"DeliveryData"` // 0...n
}
