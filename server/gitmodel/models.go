package gitmodel

type Data struct {
	TotalCount int64   `json:"total_count"`
	Items      []Items `json:"items"`
}

type Items struct {
	Name  string `json:"name"`
	Stars int64  `json:"stargazers_count"`
}
