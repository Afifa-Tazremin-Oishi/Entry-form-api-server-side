package model

// Country struct for db table - country
type Hrdept struct {
	Code  int    `json:"code"`
	Dept  string `json:"dept"`
	// Examlevel   int     `json:"examlevel" validate:"required"`
	// Subid       int     `json:"subid"`
	// Monthid     int     `json:"monthid"`
	// Sessionyear int     `json:"sessionyear"`
	// Amount      float32 `json:"amount"`
}
// type Country struct{
// 	Id int `json:"id"`
// 	Names string `json:"names"`
// }