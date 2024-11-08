package tags

type User struct {
	Username         string `json:"username"`
	Password         string
	ShouldBeIncluded int `yaml:"should_be_included"`
}

type WillBeExcluded struct {
	Hello string
	World string
}
