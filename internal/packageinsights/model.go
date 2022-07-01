package packageinsights

import "time"

type Response struct {
	AnalyzedAt time.Time  `json:"analyzedAt"`
	Collected  Collected  `json:"collected"`
	Evaluation Evaluation `json:"evaluation"`
	Score      Score      `json:"score"`
}
type Publisher struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
type Maintainers struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
type Repository struct {
	Type      string `json:"type"`
	URL       string `json:"url"`
	Directory string `json:"directory"`
}
type Links struct {
	Npm        string `json:"npm"`
	Homepage   string `json:"homepage"`
	Repository string `json:"repository"`
	Bugs       string `json:"bugs"`
}

type Metadata struct {
	Name              string        `json:"name"`
	Scope             string        `json:"scope"`
	Version           string        `json:"version"`
	Description       string        `json:"description"`
	Keywords          []string      `json:"keywords"`
	Date              time.Time     `json:"date"`
	Publisher         Publisher     `json:"publisher"`
	Maintainers       []Maintainers `json:"maintainers"`
	Repository        Repository    `json:"repository"`
	Links             Links         `json:"links"`
	License           string        `json:"license"`
	HasSelectiveFiles bool          `json:"hasSelectiveFiles"`
}
type Downloads struct {
	From  time.Time `json:"from"`
	To    time.Time `json:"to"`
	Count int       `json:"count"`
}
type Npm struct {
	Downloads  []Downloads `json:"downloads"`
	StarsCount int         `json:"starsCount"`
}
type Distribution map[string]int
type Issues struct {
	Count        int          `json:"count"`
	OpenCount    int          `json:"openCount"`
	Distribution Distribution `json:"distribution"`
	IsDisabled   bool         `json:"isDisabled"`
}
type Contributors struct {
	Username     string `json:"username"`
	CommitsCount int    `json:"commitsCount"`
}
type Commits struct {
	From  time.Time `json:"from"`
	To    time.Time `json:"to"`
	Count int       `json:"count"`
}
type Statuses struct {
	Context string `json:"context"`
	State   string `json:"state"`
}
type Github struct {
	Homepage         string         `json:"homepage"`
	StarsCount       int            `json:"starsCount"`
	ForksCount       int            `json:"forksCount"`
	SubscribersCount int            `json:"subscribersCount"`
	Issues           Issues         `json:"issues"`
	Contributors     []Contributors `json:"contributors"`
	Commits          []Commits      `json:"commits"`
	Statuses         []Statuses     `json:"statuses"`
}
type Files struct {
	ReadmeSize   float64 `json:"readmeSize"`
	TestsSize    float64 `json:"testsSize"`
	HasChangelog bool    `json:"hasChangelog"`
}
type Urls struct {
	Original string `json:"original"`
	Shields  string `json:"shields"`
	Content  string `json:"content"`
}
type Modifiers struct {
	Type string `json:"type"`
}
type Info struct {
	Service   string    `json:"service"`
	Type      string    `json:"type"`
	Modifiers Modifiers `json:"modifiers"`
}
type Urls0 struct {
	Original string `json:"original"`
	Service  string `json:"service"`
	Shields  string `json:"shields"`
	Content  string `json:"content"`
}
type Info0 struct {
	Service string `json:"service"`
	Type    string `json:"type"`
}
type Badges struct {
	Urls Urls `json:"urls,omitempty"`
	Info Info `json:"info,omitempty"`
}
type Source struct {
	Files   Files    `json:"files"`
	Badges  []Badges `json:"badges"`
	Linters []string `json:"linters"`
}
type Collected struct {
	Metadata Metadata `json:"metadata"`
	Npm      Npm      `json:"npm"`
	Github   Github   `json:"github"`
	Source   Source   `json:"source"`
}
type Quality struct {
	Carefulness float64 `json:"carefulness"`
	Tests       float64 `json:"tests"`
	Health      float64 `json:"health"`
	Branding    float64 `json:"branding"`
}
type Popularity struct {
	CommunityInterest     float64 `json:"communityInterest"`
	DownloadsCount        float64 `json:"downloadsCount"`
	DownloadsAcceleration float64 `json:"downloadsAcceleration"`
	DependentsCount       float64 `json:"dependentsCount"`
}
type Maintenance struct {
	ReleasesFrequency  float64 `json:"releasesFrequency"`
	CommitsFrequency   float64 `json:"commitsFrequency"`
	OpenIssues         float64 `json:"openIssues"`
	IssuesDistribution float64 `json:"issuesDistribution"`
}
type Evaluation struct {
	Quality     Quality     `json:"quality"`
	Popularity  Popularity  `json:"popularity"`
	Maintenance Maintenance `json:"maintenance"`
}
type Detail struct {
	Quality     float64 `json:"quality"`
	Popularity  float64 `json:"popularity"`
	Maintenance float64 `json:"maintenance"`
}
type Score struct {
	Final  float64 `json:"final"`
	Detail Detail  `json:"detail"`
}
