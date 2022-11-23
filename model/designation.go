package model

// Country struct for db table - country
type Desig struct {
	Code         int    `json:"code"`
	Designation  string `json:"designation"`
	Sdesignation string `json:"sdesignation"`
	// Examlevel   int     `json:"examlevel" validate:"required"`
	// Subid       int     `json:"subid"`
	// Monthid     int     `json:"monthid"`
	// Sessionyear int     `json:"sessionyear"`
	// Amount      float32 `json:"amount"`
}
type Country struct {
	Id    int    `json:"id"`
	Names string `json:"names"`
}
