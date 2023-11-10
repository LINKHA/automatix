// Code generated by goctl. DO NOT EDIT.
package types

type Server struct {
	ServerId      string `json:"serverId"`
	Name          string `json:"name"`
	ServerType    int64  `json:"serverType"`
	Switch        int64  `json:"switch"`
	StartTime     int64  `json:"startTime"`
	Area          int64  `json:"area"`
	Tags          string `json:"tags"`
	MaxOnline     int64  `json:"maxOnline"`
	MaxQueue      int64  `json:"maxQueue"`
	MaxSign       int64  `json:"maxSign"`
	TemplateValue string `json:"templateValue"`
}

type GetServerReq struct {
	ServerId string `json:"serverId"`
}

type GetServerResp struct {
	ReturnCode    int64  `json:"returnCode"`
	ServerId      string `json:"serverId"`
	Name          string `json:"name"`
	ServerType    int64  `json:"serverType"`
	Switch        int64  `json:"switch"`
	StartTime     int64  `json:"startTime"`
	Area          int64  `json:"area"`
	Tags          string `json:"tags"`
	MaxOnline     int64  `json:"maxOnline"`
	MaxQueue      int64  `json:"maxQueue"`
	MaxSign       int64  `json:"maxSign"`
	TemplateValue string `json:"templateValue"`
}

type SetServerReq struct {
	ServerId      string `json:"serverId"`
	Name          string `json:"name"`
	ServerType    int64  `json:"serverType"`
	Switch        int64  `json:"switch"`
	Area          int64  `json:"area"`
	Tags          string `json:"tags"`
	MaxOnline     int64  `json:"maxOnline"`
	MaxQueue      int64  `json:"maxQueue"`
	MaxSign       int64  `json:"maxSign"`
	TemplateValue string `json:"templateValue"`
}

type SetServerResp struct {
	ReturnCode int64 `json:"returnCode"`
}

type CreateServerReq struct {
	Name          string `json:"name"`
	ServerType    int64  `json:"serverType"`
	Switch        int64  `json:"switch"`
	StartTime     int64  `json:"startTime"`
	Area          int64  `json:"area"`
	Tags          string `json:"tags"`
	MaxOnline     int64  `json:"maxOnline"`
	MaxQueue      int64  `json:"maxQueue"`
	MaxSign       int64  `json:"maxSign"`
	TemplateValue string `json:"templateValue"`
}

type CreateServerResp struct {
	ServerId   string `json:"serverId"`
	ReturnCode int64  `json:"returnCode"`
}

type LoginServerReq struct {
	PlayerId string `json:"playerId"`
	ServerId string `json:"serverId"`
}

type LoginServerResp struct {
	ReturnCode int64 `json:"returnCode"`
}
