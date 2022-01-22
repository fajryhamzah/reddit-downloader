package data

type MainResponse struct {
	Data childrenResponse `json:"data"`
}

type childrenResponse struct {
	Children []struct {
		Data DataResponse
	} `json:"children"`
}

type DataResponse struct {
	Subreddit       string `json:"subreddit"`
	Title           string `json:"title"`
	SubredditPrefix string `json:"subreddit_name_prefixed"`
	Author          string `json:"author"`
	IsVideo         bool   `json:"is_video"`
	UrlDestination  string `json:"url_overridden_by_dest"`
	Permalink       string `json:"permalink"`
	SecureMedia     struct {
		RedditVideo redditVideo `json:"reddit_video"`
	} `json:"secure_media"`
}

type redditVideo struct {
	FallbackUrl string `json:"fallback_url"`
	HlsUrl      string `json:"hls_url"`
	IsGIF       bool   `json:"is_gif"`
	Bitrate     int32  `json:"bitrate_kbps"`
}
