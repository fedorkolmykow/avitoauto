package models


type SaveURLReq struct{
	OriginalURL 	string				`json:"url"`
	CustomKey 		string				`json:"custom_key"`
}
type SaveURLResp struct {
	Key				string				`json:"key"`
}

type RedirectReq struct {
	Key				string
}
type RedirectResp struct {
	OriginalURL 	string
}

type URL struct {
	URLId			int					`json:"-" db:"url_id"`
	URL          	string              `json:"url" db:"url"`
	Key				string				`json:"key" db:"key"`
}

