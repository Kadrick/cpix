package cpix

import (
	"encoding/xml"
)

type CPIX struct {
	XMLName xml.Name `xml:"CPIX"`

	// namespace / namespace : value : definition
	NSCpix  xml.Attr `xml:"cpix,attr"`  // cpix : urn:dashif:org:cpix : https://dashif.org/docs/DASH-IF-CPIX-v2-0.pdf
	NSEnc   xml.Attr `xml:"enc,attr"`   // xml enc : http://www.w3.org/2001/04/xmlenc# : http://www.w3.org/2001/04/xmlenc#
	NSPskc  xml.Attr `xml:"pskc,attr"`  // key contanier : urn:ietf:params:xml:ns:keyprov:pskc : https://datatracker.ietf.org/doc/html/rfc6030
	NSSpeke xml.Attr `xml:"speke,attr"` // speke (aws) : urn:aws:amazon:com:speke : https://docs.aws.amazon.com/speke/latest/documentation/what-is-speke.html

	// Attr
	Id        xml.Attr `xml:"id,attr,omitempty"`
	Name      xml.Attr `xml:"name,attr,omitempty"`
	ContentId xml.Attr `xml:"contentId,attr,omitempty"`
	Version   xml.Attr `xml:"version,attr,omitempty"`

	// Child
	DeliveryDataList        *DeliveryDataList        `xml:"DeliveryDataList,omitempty"`        // 0...1
	ContentKeyList          *ContentKeyList          `xml:"ContentKeyList,omitempty"`          // 0...1
	DRMSystemList           *DRMSystemList           `xml:"DRMSystemList,omitempty"`           // 0...1
	ContentKeyPeriodList    *ContentKeyPeriodList    `xml:"ContentKeyPeriodList,omitempty"`    // 0...1
	ContentKeyUsageRuleList *ContentKeyUsageRuleList `xml:"ContentKeyUsageRuleList,omitempty"` // 0...1
	UpdateHistoryItemList   *UpdateHistoryItemList   `xml:"UpdateHistoryItemList,omitempty"`   // 0...1

	// TODO: Signature
}
