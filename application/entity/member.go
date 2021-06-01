package entity

type Member struct {
	Audit
	Id                   int64
	UserId               string
	Name                 string
	Password             string
	AccountNonExpired    bool
	AccountNonLocked     bool
	CredentialNonExpired bool
	Enabled              bool
	Roles                []string
	MemberInformationId  int64
}

type MemberInformation struct {
	Audit
	Id    int64
	Email string
}
