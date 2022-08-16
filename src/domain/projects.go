package domain

type ProjectList struct {
	Projects []Project
}

type Project struct {
	Name, Id string
}
