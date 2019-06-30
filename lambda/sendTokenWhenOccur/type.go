package main

type EewData struct {
	Version Version `json:"version"`
	Common  Common  `json:"common"`
	Details Details `json:"details"`
}

type Version struct {
	CommonVersion  string `json:"common_version"`
	DetailsVersion string `json:"details_version"`
}

type Common struct {
	Datatype     string `json:"datatype"`
	Msgid        string `json:"msgid"`
	Sendid       string `json:"sendid"`
	Senddatetime string `json:"senddatetime"`
}

type EewInfo struct {
	Eewid           string  `json:"eewid"`
	Sequence        int     `json:"sequence"`
	IsFinal         int     `json:"is_final"`
	OccuredDatetime int64   `json:"occured_datetime"`
	ReportDatetime  int64   `json:"report_datetime"`
	Hypocode        int     `json:"hypocode"`
	HypocenterIsSea int     `json:"hypocenter_is_sea"`
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	Depth           int     `json:"depth"`
	Magnitude       float64 `json:"magnitude"`
}

type AreaInfo struct {
	Alert     int     `json:"alert"`
	Intensity float64 `json:"intensity"`
	STime     int64   `json:"s_time"`
}

type Details struct {
	Type     int                 `json:"type"`
	Office   int                 `json:"office"`
	Alert    int                 `json:"alert"`
	Cancel   int                 `json:"cancel"`
	Eewinfo  EewInfo             `json:"eewinfo"`
	Areainfo map[string]AreaInfo `json:"areainfo"`
}
