package release

import (
	"fmt"
	"testing"
)

const (
	rcloneVer = "1.63.0"
	ytVer     = "2023.06.22"
)

func TestGithub(t *testing.T) {
	tests{
		{
			InputRepo: "rclone/rclone", InputMatch: "", F: Github,
			Expected:    fmt.Sprintf("https://github.com/rclone/rclone/releases/download/v%s/rclone-v%s-linux-amd64.deb", rcloneVer, rcloneVer),
			ExpectedErr: nil,
		},
		{
			InputRepo: "myztillx-test-repo", InputMatch: "", F: Github,
			Expected:    "",
			ExpectedErr: ErrImproperRepo,
		},
		{
			InputRepo: "myztillx/test-repo", InputMatch: "", F: Github,
			Expected:    "",
			ExpectedErr: ErrNoAssets,
		},
		{
			InputRepo: "yt-dlp/yt-dlp", InputMatch: "_linux", F: Github,
			Expected:    fmt.Sprintf("https://github.com/yt-dlp/yt-dlp/releases/download/%s/yt-dlp_linux", ytVer),
			ExpectedErr: nil,
		},
	}.validate(t)
}
