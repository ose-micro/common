package claims

import "fmt"

type Permission struct {
	Resource string `json:"resource"`
	Action   string `json:"action"`
}

func (p *Permission) String() string {
	return fmt.Sprintf("%s:%s", p.Resource, p.Action)
}
