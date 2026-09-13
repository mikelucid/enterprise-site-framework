package domain

type Recommendation struct {
	ID          string
	Type        string
	Action      string
	Description string
	Confidence  float32
	SiteID      string
	Metadata    map[string]string
}
