package models

type SaveURLReq struct{
	OriginalURL 	string				`json:"url"`
	CustomKey 		string				`json:"custom_key,omitempty"`
}
type SaveURLResp struct {
	KeyID			int					`json:"-"`
	Key				string				`json:"key"`
}

type RedirectReq struct {
	Key				string
}
type RedirectResp struct {
	OriginalURL 	string
}

type URL struct {
	Id				int					`json:"-" db:"url_id"`
	URL          	string              `json:"url" db:"url"`
	Key				string				`json:"key" db:"key"`
}

