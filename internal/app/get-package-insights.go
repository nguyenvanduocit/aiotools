package app

import "time"

type PackageInsight struct {
}

type GetPackageInsightsRequest struct {
	PackageName string
}

type GetPackageInsightsResponse struct {
	AnalyzedAt time.Time  `json:"analyzedAt,omitempty"`
	Collected  Collected  `json:"collected,omitempty"`
	Evaluation Evaluation `json:"evaluation,omitempty"`
	Score      Score      `json:"score,omitempty"`
}

type Author struct {
	Name string `json:"name,omitempty"`
}

type Publisher struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
}

type Maintainers struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
}

type Repository struct {
	Type      string `json:"type,omitempty"`
	URL       string `json:"url,omitempty"`
	Directory string `json:"directory,omitempty"`
}

type Links struct {
	Npm        string `json:"npm,omitempty"`
	Homepage   string `json:"homepage,omitempty"`
	Repository string `json:"repository,omitempty"`
	Bugs       string `json:"bugs,omitempty"`
}

type Dependencies map[string]string
type DevDependencies map[string]string
type PeerDependencies map[string]string

type Releases struct {
	From  time.Time `json:"from,omitempty"`
	To    time.Time `json:"to,omitempty"`
	Count int       `json:"count,omitempty"`
}

type Metadata struct {
	Name             string           `json:"name,omitempty"`
	Scope            string           `json:"scope,omitempty"`
	Version          string           `json:"version,omitempty"`
	Description      string           `json:"description,omitempty"`
	Keywords         []string         `json:"keywords,omitempty"`
	Date             time.Time        `json:"date,omitempty"`
	Author           Author           `json:"author,omitempty"`
	Publisher        Publisher        `json:"publisher,omitempty"`
	Maintainers      []Maintainers    `json:"maintainers,omitempty"`
	Repository       Repository       `json:"repository,omitempty"`
	Links            Links            `json:"links,omitempty"`
	License          string           `json:"license,omitempty"`
	Dependencies     Dependencies     `json:"dependencies,omitempty"`
	DevDependencies  DevDependencies  `json:"devDependencies,omitempty"`
	PeerDependencies PeerDependencies `json:"peerDependencies,omitempty"`
	Releases         []Releases       `json:"releases,omitempty"`
	Readme           string           `json:"readme,omitempty"`
}

type Downloads struct {
	From  time.Time `json:"from,omitempty"`
	To    time.Time `json:"to,omitempty"`
	Count int       `json:"count,omitempty"`
}

type Npm struct {
	Downloads  []Downloads `json:"downloads,omitempty"`
	StarsCount int         `json:"starsCount,omitempty"`
}

type Issues struct {
	Count      int  `json:"count,omitempty"`
	OpenCount  int  `json:"openCount,omitempty"`
	IsDisabled bool `json:"isDisabled,omitempty"`
}

type Contributors struct {
	Username     string `json:"username,omitempty"`
	CommitsCount int    `json:"commitsCount,omitempty"`
}

type Commits struct {
	From  time.Time `json:"from,omitempty"`
	To    time.Time `json:"to,omitempty"`
	Count int       `json:"count,omitempty"`
}

type Statuses struct {
	Context string `json:"context,omitempty"`
	State   string `json:"state,omitempty"`
}

type Github struct {
	Homepage         string         `json:"homepage,omitempty"`
	StarsCount       int            `json:"starsCount,omitempty"`
	ForksCount       int            `json:"forksCount,omitempty"`
	SubscribersCount int            `json:"subscribersCount,omitempty"`
	Issues           Issues         `json:"issues,omitempty"`
	Contributors     []Contributors `json:"contributors,omitempty"`
	Commits          []Commits      `json:"commits,omitempty"`
	Statuses         []Statuses     `json:"statuses,omitempty"`
}

type Files struct {
	ReadmeSize   int  `json:"readmeSize,omitempty"`
	TestsSize    int  `json:"testsSize,omitempty"`
	HasNpmIgnore bool `json:"hasNpmIgnore,omitempty"`
	HasChangelog bool `json:"hasChangelog,omitempty"`
}

type Urls struct {
	Original string `json:"original,omitempty"`
	Service  string `json:"service,omitempty"`
	Shields  string `json:"shields,omitempty"`
	Content  string `json:"content,omitempty"`
}

type Modifiers struct {
	Branch string `json:"branch,omitempty"`
}

type Info struct {
	Service   string    `json:"service,omitempty"`
	Type      string    `json:"type,omitempty"`
	Modifiers Modifiers `json:"modifiers,omitempty"`
}

type Badges struct {
	Urls Urls `json:"urls,omitempty"`
	Info Info `json:"info,omitempty"`
}

type Warn struct {
	Pkgid      string `json:"pkgid,omitempty"`
	StatusCode int    `json:"statusCode,omitempty"`
	Code       string `json:"code,omitempty"`
}

type OutdatedDependencies map[string]string
type Source struct {
	Files                Files                `json:"files,omitempty"`
	Badges               []Badges             `json:"badges,omitempty"`
	Linters              []string             `json:"linters,omitempty"`
	OutdatedDependencies OutdatedDependencies `json:"outdatedDependencies,omitempty"`
}

type Collected struct {
	Metadata Metadata `json:"metadata,omitempty"`
	Npm      Npm      `json:"npm,omitempty"`
	Github   Github   `json:"github,omitempty"`
	Source   Source   `json:"source,omitempty"`
}
type Quality struct {
	Carefulness float64 `json:"carefulness,omitempty"`
	Tests       float64 `json:"tests,omitempty"`
	Health      float64 `json:"health,omitempty"`
	Branding    float64 `json:"branding,omitempty"`
}
type Popularity struct {
	CommunityInterest     int     `json:"communityInterest,omitempty"`
	DownloadsCount        float64 `json:"downloadsCount,omitempty"`
	DownloadsAcceleration float64 `json:"downloadsAcceleration,omitempty"`
	DependentsCount       int     `json:"dependentsCount,omitempty"`
}
type Maintenance struct {
	ReleasesFrequency  int `json:"releasesFrequency,omitempty"`
	CommitsFrequency   int `json:"commitsFrequency,omitempty"`
	OpenIssues         int `json:"openIssues,omitempty"`
	IssuesDistribution int `json:"issuesDistribution,omitempty"`
}
type Evaluation struct {
	Quality     Quality     `json:"quality,omitempty"`
	Popularity  Popularity  `json:"popularity,omitempty"`
	Maintenance Maintenance `json:"maintenance,omitempty"`
}
type Detail struct {
	Quality     float64 `json:"quality,omitempty"`
	Popularity  float64 `json:"popularity,omitempty"`
	Maintenance float64 `json:"maintenance,omitempty"`
}
type Score struct {
	Final  float64 `json:"final,omitempty"`
	Detail Detail  `json:"detail,omitempty"`
}

func (a *PackageInsight) GetPackageInsight(request *GetPackageInsightsRequest) (*GetPackageInsightsResponse, error) {
	return nil, nil
}
