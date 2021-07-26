//Models/UserModel.go
package Models
type Student struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	LastName   string `json:"lastname"`
	DOB   string `json:"dob"`
	Address string `json:"address"`
	Subject string `json:"subject"`
	Mark string `json:"mark"`
}
func (b *Student) TableName() string {
	return "student"
}