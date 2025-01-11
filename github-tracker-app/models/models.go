package models

type Commit struct {
	ID      string     `json:"id"`
	Message string     `json:"message"`
	Author  CommitUser `json:"author"`
}

type Repository struct {
	FullName string `json:"full_name"`
}

type PushEvent struct {
	Repository Repository `json:"repository"`
	HeadCommit Commit     `json:"head_commit"`
}

type CommitUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
