package cpix

import "encoding/xml"

type (
	PSSH struct {
		XMLName xml.Name `xml:"PSSH"`
		Value   string   `xml:",chardata"` // base64
	}

	ContentProtectionData struct {
		XMLName xml.Name `xml:"ContentProtectionData"`
		Value   string   `xml:",chardata"` // base64
	}

	URIExtXKey struct {
		XMLName xml.Name `xml:"URIExtXKey"`
		Value   string   `xml:",chardata"` // base64
	}

	HLSSignalingData struct {
		XMLName      xml.Name `xml:"HLSSignalingData"`
		PlaylistType xml.Attr `xml:"playlist,attr"` // "master" or "media"
	}

	SmoothStreamingProtectionHeaderData struct {
		XMLName xml.Name `xml:"SmoothStreamingProtectionHeaderData"`
		Value   string   `xml:",chardata"` // string
	}

	HDSSignalingData struct {
		XMLName xml.Name `xml:"HDSSignalingData"`
		Value   string   `xml:",chardata"` // base64
	}

	/*=============================================*/

	KeyFormat struct {
		XMLName xml.Name `xml:"KeyFormat"`
		Value   string   `xml:",chardata"` // base64
	}

	KeyFormatVersions struct {
		XMLName xml.Name `xml:"KeyFormatVersions"`
		Value   string   `xml:",chardata"` // base64
	}
)

type DRMSystem struct {
	XMLName xml.Name `xml:"DRMSystem"`

	// Attr
	Id            xml.Attr `xml:"id,attr,omitempty"`
	Name          xml.Attr `xml:"name,attr,omitempty"`
	UpdateVersion xml.Attr `xml:"updateVersion,attr,omitempty"`

	SystemId xml.Attr `xml:"systemId,attr"` // Required, https://dashif.org/identifiers/content_protection/
	KId      xml.Attr `xml:"kid,attr"`      // Required

	// Child
	PSSH                                *PSSH                                `xml:"PSSH,omitempty"`                                // 0...1
	ContentProtectionData               *ContentProtectionData               `xml:"ContentProtectionData,omitempty"`               // 0...1
	URIExtXKey                          *URIExtXKey                          `xml:"URIExtXKey,omitempty"`                          // 0...1
	HLSSignalingData                    *HLSSignalingData                    `xml:"HLSSignalingData,omitempty"`                    // 0...2
	SmoothStreamingProtectionHeaderData *SmoothStreamingProtectionHeaderData `xml:"SmoothStreamingProtectionHeaderData,omitempty"` // 0...1
	HDSSignalingData                    *HDSSignalingData                    `xml:"HDSSignalingData,omitempty"`                    // 0...1

	// Speke
	KeyFormat         *KeyFormat         `xml:"KeyFormat,omitempty"`         // 0...1
	KeyFormatVersions *KeyFormatVersions `xml:"KeyFormatVersions,omitempty"` // 0...1
}

type DRMSystemList struct {
	XMLName xml.Name `xml:"DRMSystemList,omitempty"`

	// Attr
	Id            xml.Attr `xml:"id,attr,omitempty"`
	UpdateVersion xml.Attr `xml:"updateVersion,attr,omitempty"`

	// Child
	List []*DRMSystem `xml:"DRMSystem"` // 0...n
}
