package data

type MainResponse struct {
	Data childrenResponse `json:"data"`
}

type childrenResponse struct {
	Children []struct {
		Data mainDataResponse
	} `json:"children"`
}

type mainDataResponse struct {
	Subreddit       string `json:"subreddit"`
	Title           string `json:"title"`
	SubredditPrefix string `json:"subreddit_name_prefixed"`
	Author          string `json:"author"`
	IsVideo         bool   `json:"is_video"`
	UrlDestination  string `json:"url_overridden_by_dest"`
}
