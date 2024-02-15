package cpix

import "encoding/xml"

type ContentKeyUsageRule struct {
	XMLName xml.Name `xml:"ContentKeyUsageRule"`

	// Attr
	Id                xml.Attr `xml:"id,attr,omitempty"`
	KId               xml.Attr `xml:"kid,attr"` // Required
	IntendedTrackType xml.Attr `xml:"intendedTrackType,attr,omitempty"`

	// Child
	KeyPeriodFilter struct {
		XMLName  xml.Name `xml:"KeyPeriodFilter,omitempty"`
		PeriodId xml.Attr `xml:"periodId,attr"` // Required
	} `xml:"KeyPeriodFilter,omitempty"` // 0...n

	// * not supported in speke

	LabelFilter struct {
		XMLName xml.Name `xml:"LabelFilter,omitempty"`
		Label   xml.Attr `xml:"label,attr"` // Required
	} `xml:"LabelFilter,omitempty"` // 0...n

	VideoFilter struct {
		XMLName xml.Name `xml:"VideoFilter,omitempty"`

		MinPixels xml.Attr `xml:"minPixels,attr,omitempty"`
		MaxPixels xml.Attr `xml:"maxPixels,attr,omitempty"`
		HDR       xml.Attr `xml:"hdr,attr,omitempty"`
		WCG       xml.Attr `xml:"wcg,attr,omitempty"`
		MinFps    xml.Attr `xml:"minFps,attr,omitempty"`
		MaxFps    xml.Attr `xml:"maxFps,attr,omitempty"`
	} `xml:"VideoFilter,omitempty"` // 0...n

	AudioFilter struct {
		XMLName xml.Name `xml:"AudioFilter,omitempty"`

		MinChannels xml.Attr `xml:"minChannels,attr,omitempty"`
		MaxChannels xml.Attr `xml:"maxChannels,attr,omitempty"`
	} `xml:"AudioFilter,omitempty"`
	// 0...n

	BitrateFilter struct {
		XMLName xml.Name `xml:"BitrateFilter,omitempty"`

		MinBitrate xml.Attr `xml:"minBitrate,attr,omitempty"`
		MaxBitrate xml.Attr `xml:"maxBitrate,attr,omitempty"`
	} `xml:"BitrateFilter,omitempty"` // 0...n

}

type ContentKeyUsageRuleList struct {
	XMLName xml.Name `xml:"ContentKeyUsageRuleList,omitempty"`

	// Attr
	Id            xml.Attr `xml:"id,attr,omitempty"`
	UpdateVersion xml.Attr `xml:"updateVersion,attr,omitempty"`

	// Child
	List []*ContentKeyUsageRule `xml:"ContentKeyUsageRule"` // 0...n
}
