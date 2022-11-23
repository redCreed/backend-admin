package dto

type Swag struct {
	Swagger string                 `json:"swagger"`
	Info    interface{}            `json:"info"`
	Paths   map[string]interface{} `json:"paths"`
}

type Api struct {
	Handler     string `json:"handler"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Uri         string `json:"uri"`
}

type MenuApiList struct {
	MenuId int    `json:"menu_id"`
	ApiId  int    `json:"api_id"`
	Handle string `json:"handle" `
	Title  string `json:"title" `
	Path   string `json:"path" `
	Method string `json:"method" `
}
