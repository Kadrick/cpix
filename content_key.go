// pskc ref: https://datatracker.ietf.org/doc/html/rfc6030#section-11
package cpix

import "encoding/xml"

type (
	PlainValue struct {
		XMLName xml.Name `xml:"PlainValue"`
		Value   string   `xml:",chardata"`
	}

	EncryptedValue struct {
		XMLName xml.Name `xml:"EncryptedValue"`
	}

	Secret struct {
		XMLName xml.Name `xml:"Secret"`

		PlainValue     *PlainValue     `xml:"PlainValue,omitempty"`
		EncryptedValue *EncryptedValue `xml:"EncryptedValue,omitempty"`
	}

	Data struct {
		XMLName xml.Name `xml:"Data"`
		// pskc
		Secret *Secret `xml:"Secret,omitempty"`
	}
)

// cpix
type ContentKey struct {
	XMLName xml.Name `xml:"ContentKey"`

	// Attr
	Id        xml.Attr `xml:"id,attr,omitempty"`
	Algorithm xml.Attr `xml:"Algorithm,attr,omitempty"`

	KId          xml.Attr `xml:"kid,attr"` // required
	ExplicitIV   xml.Attr `xml:"explicitIV,attr,omitempty"`
	DependsOnKey xml.Attr `xml:"dependsOnKey,attr,omitempty"`

	// Child
	Issuer *struct {
		XMLName xml.Name `xml:"Issuer"`
		Value   string   `xml:",chardata"`
	} `xml:"Issuer,omitempty"`

	KeyProfileId *struct {
		XMLName xml.Name `xml:"KeyProfileId"`
		Value   string   `xml:",chardata"`
	} `xml:"KeyProfileId,omitempty"`

	KeyReference *struct {
		XMLName xml.Name `xml:"KeyReference"`
		Value   string   `xml:",chardata"`
	} `xml:"KeyReference,omitempty"`

	FriendlyName *struct {
		XMLName xml.Name `xml:"FriendlyName"`
		Value   string   `xml:",chardata"`
	} `xml:"FriendlyName,omitempty"`

	UserId *struct {
		XMLName xml.Name `xml:"UserId"`
		Value   string   `xml:",chardata"`
	} `xml:"UserId,omitempty"`

	Data *Data `xml:"Data,omitempty"`

	// TODO
	AlgorithmParameters *struct {
	} `xml:"AlgorithmParameters,omitempty"`

	// TODO
	Policy *struct {
	} `xml:"Policy,omitempty"`

	// TODO
	Extensions []*struct {
	} `xml:"Extensions,omitempty"`
}

type ContentKeyList struct {
	XMLName xml.Name `xml:"ContentKeyList,omitempty"`

	// Attr
	Id            xml.Attr `xml:"id,attr,omitempty"`
	UpdateVersion xml.Attr `xml:"updateVersion,attr,omitempty"`

	// Child
	List []*ContentKey `xml:"ContentKey"` // 0...n
}
