package release

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"runtime"
	"strings"
	"time"
)

type GithubApi struct {
	URL       string `json:"url"`
	AssetsURL string `json:"assets_url"`
	UploadURL string `json:"upload_url"`
	HTMLURL   string `json:"html_url"`
	ID        int    `json:"id"`
	Author    struct {
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
	} `json:"author"`
	NodeID          string    `json:"node_id"`
	TagName         string    `json:"tag_name"`
	TargetCommitish string    `json:"target_commitish"`
	Name            string    `json:"name"`
	Draft           bool      `json:"draft"`
	Prerelease      bool      `json:"prerelease"`
	CreatedAt       time.Time `json:"created_at"`
	PublishedAt     time.Time `json:"published_at"`
	Assets          []struct {
		URL      string `json:"url"`
		ID       int    `json:"id"`
		NodeID   string `json:"node_id"`
		Name     string `json:"name"`
		Label    string `json:"label"`
		Uploader struct {
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
		} `json:"uploader"`
		ContentType        string    `json:"content_type"`
		State              string    `json:"state"`
		Size               int       `json:"size"`
		DownloadCount      int       `json:"download_count"`
		CreatedAt          time.Time `json:"created_at"`
		UpdatedAt          time.Time `json:"updated_at"`
		BrowserDownloadURL string    `json:"browser_download_url"`
	} `json:"assets"`
	TarballURL string `json:"tarball_url"`
	ZipballURL string `json:"zipball_url"`
	Body       string `json:"body"`
	Reactions  struct {
		URL        string `json:"url"`
		TotalCount int    `json:"total_count"`
		Num1       int    `json:"+1"`
		Num10      int    `json:"-1"`
		Laugh      int    `json:"laugh"`
		Hooray     int    `json:"hooray"`
		Confused   int    `json:"confused"`
		Heart      int    `json:"heart"`
		Rocket     int    `json:"rocket"`
		Eyes       int    `json:"eyes"`
	} `json:"reactions"`
}

var ErrImproperRepo = errors.New("improper repo name, should be in format owner/repository")
var ErrFailedFetch = errors.New("failed to fetch release information")
var ErrFailedReadResp = errors.New("failed to read response body")
var ErrJsonParse = errors.New("there was an error parsing json")
var ErrNoAssets = errors.New("there are no assets for this repo")
var ErrCheckPkg = errors.New("error checking distro package manager")
var ErrNoRepoUrl = errors.New("no repo url found")

func Github(repo string, match string) (string, error) {
	if !strings.Contains(repo, "/") {
		return "", ErrImproperRepo
	}

	// Send a GET request to the GitHub API
	url := fmt.Sprintf("https://api.github.com/repos/%s/releases/latest", repo)
	response, err := http.Get(url)
	if err != nil {
		return "", ErrFailedFetch
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", ErrFailedReadResp
	}

	// Convert the response body to a string
	responseStr := string(body)

	var data GithubApi
	if err := json.Unmarshal([]byte(responseStr), &data); err != nil {
		return "", ErrJsonParse
	}
	assets := data.Assets
	if len(assets) < 1 {
		return "", ErrNoAssets
	}

	var searchStr string
	if match == "" {
		var pkgName string
		if runtime.GOOS == "linux" {
			pkgName, err = CheckPkg()
			if err != nil {
				return "", ErrCheckPkg
			}
		}
		archName := runtime.GOARCH
		searchStr = fmt.Sprintf("%s.%s", archName, pkgName)
	} else {
		searchStr = match
	}

	var repoUrl string
	for _, v := range assets {
		if strings.Contains(v.BrowserDownloadURL, searchStr) {
			repoUrl = v.BrowserDownloadURL
			break
		}

	}
	if len(repoUrl) < 10 {
		return "", ErrNoRepoUrl
	}

	return repoUrl, nil
}
