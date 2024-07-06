package api

import "time"

type GithubUser struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

type LinkRef struct {
	Href string `json:"href"`
}

type Links struct {
	Self           LinkRef `json:"self"`
	HTML           LinkRef `json:"html"`
	Issue          LinkRef `json:"issue"`
	Comments       LinkRef `json:"comments"`
	ReviewComments LinkRef `json:"review_comments"`
	ReviewComment  LinkRef `json:"review_comment"`
	Commits        LinkRef `json:"commits"`
	Statuses       LinkRef `json:"statuses"`
}

type GithubRepo struct {
	ID                       int        `json:"id"`
	NodeID                   string     `json:"node_id"`
	Name                     string     `json:"name"`
	FullName                 string     `json:"full_name"`
	Private                  bool       `json:"private"`
	Owner                    GithubUser `json:"owner"`
	HTMLURL                  string     `json:"html_url"`
	Description              string     `json:"description"`
	Fork                     bool       `json:"fork"`
	URL                      string     `json:"url"`
	ForksURL                 string     `json:"forks_url"`
	KeysURL                  string     `json:"keys_url"`
	CollaboratorsURL         string     `json:"collaborators_url"`
	TeamsURL                 string     `json:"teams_url"`
	HooksURL                 string     `json:"hooks_url"`
	IssueEventsURL           string     `json:"issue_events_url"`
	EventsURL                string     `json:"events_url"`
	AssigneesURL             string     `json:"assignees_url"`
	BranchesURL              string     `json:"branches_url"`
	TagsURL                  string     `json:"tags_url"`
	BlobsURL                 string     `json:"blobs_url"`
	GitTagsURL               string     `json:"git_tags_url"`
	GitRefsURL               string     `json:"git_refs_url"`
	TreesURL                 string     `json:"trees_url"`
	StatusesURL              string     `json:"statuses_url"`
	LanguagesURL             string     `json:"languages_url"`
	StargazersURL            string     `json:"stargazers_url"`
	ContributorsURL          string     `json:"contributors_url"`
	SubscribersURL           string     `json:"subscribers_url"`
	SubscriptionURL          string     `json:"subscription_url"`
	CommitsURL               string     `json:"commits_url"`
	GitCommitsURL            string     `json:"git_commits_url"`
	CommentsURL              string     `json:"comments_url"`
	IssueCommentURL          string     `json:"issue_comment_url"`
	ContentsURL              string     `json:"contents_url"`
	CompareURL               string     `json:"compare_url"`
	MergesURL                string     `json:"merges_url"`
	ArchiveURL               string     `json:"archive_url"`
	DownloadsURL             string     `json:"downloads_url"`
	IssuesURL                string     `json:"issues_url"`
	PullsURL                 string     `json:"pulls_url"`
	MilestonesURL            string     `json:"milestones_url"`
	NotificationsURL         string     `json:"notifications_url"`
	LabelsURL                string     `json:"labels_url"`
	ReleasesURL              string     `json:"releases_url"`
	DeploymentsURL           string     `json:"deployments_url"`
	CreatedAt                time.Time  `json:"created_at"`
	UpdatedAt                time.Time  `json:"updated_at"`
	PushedAt                 time.Time  `json:"pushed_at"`
	GitURL                   string     `json:"git_url"`
	SSHURL                   string     `json:"ssh_url"`
	CloneURL                 string     `json:"clone_url"`
	SvnURL                   string     `json:"svn_url"`
	Homepage                 any        `json:"homepage"`
	Size                     int        `json:"size"`
	StargazersCount          int        `json:"stargazers_count"`
	WatchersCount            int        `json:"watchers_count"`
	Language                 any        `json:"language"`
	HasIssues                bool       `json:"has_issues"`
	HasProjects              bool       `json:"has_projects"`
	HasDownloads             bool       `json:"has_downloads"`
	HasWiki                  bool       `json:"has_wiki"`
	HasPages                 bool       `json:"has_pages"`
	HasDiscussions           bool       `json:"has_discussions"`
	ForksCount               int        `json:"forks_count"`
	MirrorURL                any        `json:"mirror_url"`
	Archived                 bool       `json:"archived"`
	Disabled                 bool       `json:"disabled"`
	OpenIssuesCount          int        `json:"open_issues_count"`
	License                  any        `json:"license"`
	AllowForking             bool       `json:"allow_forking"`
	IsTemplate               bool       `json:"is_template"`
	WebCommitSignoffRequired bool       `json:"web_commit_signoff_required"`
	Topics                   []any      `json:"topics"`
	Visibility               string     `json:"visibility"`
	Forks                    int        `json:"forks"`
	OpenIssues               int        `json:"open_issues"`
	Watchers                 int        `json:"watchers"`
	DefaultBranch            string     `json:"default_branch"`
}

type BranchReference struct {
	Label string     `json:"label"`
	Ref   string     `json:"ref"`
	Sha   string     `json:"sha"`
	User  GithubUser `json:"user"`
	Repo  GithubRepo `json:"repo"`
}

type PullRequest struct {
	URL                string          `json:"url"`
	ID                 int             `json:"id"`
	NodeID             string          `json:"node_id"`
	HTMLURL            string          `json:"html_url"`
	DiffURL            string          `json:"diff_url"`
	PatchURL           string          `json:"patch_url"`
	IssueURL           string          `json:"issue_url"`
	Number             int             `json:"number"`
	State              string          `json:"state"`
	Locked             bool            `json:"locked"`
	Title              string          `json:"title"`
	User               GithubUser      `json:"user"`
	Body               any             `json:"body"`
	CreatedAt          time.Time       `json:"created_at"`
	UpdatedAt          time.Time       `json:"updated_at"`
	ClosedAt           any             `json:"closed_at"`
	MergedAt           any             `json:"merged_at"`
	MergeCommitSha     string          `json:"merge_commit_sha"`
	Assignee           any             `json:"assignee"`
	Assignees          []any           `json:"assignees"`
	RequestedReviewers []any           `json:"requested_reviewers"`
	RequestedTeams     []any           `json:"requested_teams"`
	Labels             []any           `json:"labels"`
	Milestone          any             `json:"milestone"`
	Draft              bool            `json:"draft"`
	CommitsURL         string          `json:"commits_url"`
	ReviewCommentsURL  string          `json:"review_comments_url"`
	ReviewCommentURL   string          `json:"review_comment_url"`
	CommentsURL        string          `json:"comments_url"`
	StatusesURL        string          `json:"statuses_url"`
	Head               BranchReference `json:"head"`
	Base               BranchReference `json:"base"`
	Links              Links           `json:"_links"`
	AuthorAssociation  string          `json:"author_association"`
	AutoMerge          any             `json:"auto_merge"`
	ActiveLockReason   any             `json:"active_lock_reason"`
}
