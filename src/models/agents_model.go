package models

type Agent struct {
	Name string `json:"name"`
}

func GetAgentsList() []Agent {
	return []Agent{{Name: "Agent000"}, {Name: "Agent001"}, {Name: "Agent007"}}
}
