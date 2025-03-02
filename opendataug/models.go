package opendataug

// District represents a district in Uganda
type District struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	TownStatus bool   `json:"town_status"`
	RegionID   string `json:"region_id"`
	RegionName string `json:"region_name"`
}

// County represents a county in Uganda
type County struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Code       string `json:"code"`
	DistrictID string `json:"district_id"`
	CreatedAt  string `json:"created_at,omitempty"`
	UpdatedAt  string `json:"updated_at,omitempty"`
}

// Subcounty represents a subcounty in Uganda
type Subcounty struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Code      string `json:"code"`
	CountyID  string `json:"county_id"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Parish represents a parish in Uganda
type Parish struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	SubcountyID string `json:"subcounty_id"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
}

// Village represents a village in Uganda
type Village struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Code      string `json:"code"`
	ParishID  string `json:"parish_id"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Response is a generic response wrapper
type Response struct {
	Data  interface{} `json:"data"`
	Meta  Meta        `json:"meta,omitempty"`
	Error string      `json:"error,omitempty"`
}

// Meta contains pagination information
type Meta struct {
	CurrentPage int `json:"current_page"`
	LastPage    int `json:"last_page"`
	PerPage     int `json:"per_page"`
	Total       int `json:"total"`
}
