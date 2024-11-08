package tags

type User struct {
	Username         string `json:"username"`
	Password         string
	ShouldBeIncluded int `yaml:"should_be_included"`
}
