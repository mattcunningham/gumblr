package tumblr

import "encoding/json"

// The following type corresponds to the response message
type Response struct {
	Meta     Meta            `json:"meta"`     // HTTP response message
	Response json.RawMessage `json:"response"` // API-specific results
}

type Meta struct {
	Status int    `json:"status"` // the 3-digit HTTP Status-Code (e.g., 200)
	Msg    string `json:"msg"`    // the HTTP Reason-Phrase (e.g., OK)
}

// /info — Retrieve Blog Info
type BlogInfo struct {
	Blog struct {
		Title                string `json:"title"`                   // The display title of the blog
		PostCount            int    `json:"posts"`                   // The total number of posts to this blog
		Name                 string `json:"name"`                    // The short blog name that appears before tumblr.com in a standard blog hostname
		Updated              int    `json:"updated"`                 // The time of the most recent post, in seconds since the epoch
		Description          string `json:"description"`             // The blog's description
		Ask                  bool   `json:"ask"`                     // Indicates whether the blog allows questions
		AskAnon              bool   `json:"ask_anon"`                // Indicates whether the blog allows anonymous questions
		Likes                int    `json:"likes"`                   // Number of likes for this user
		IsBlockedFromPrimary bool   `json:"is_blocked_from_primary"` // Indicates whether this blog has been blocked by the calling user's primary blog
	} `json:"blog"`
}

// /avatar — Retrieve a Blog Avatar
type BlogAvatar struct {
	AvatarURL string `json:"avatar_url"` // The URL of the avatar image.
}

// /likes - Retrieve Blog's Likes
type Likes struct {
	LikedPost  []Post `json:"liked_posts"` // An array of post objects (posts liked by the user)
	LikedCount int    `json:"liked_count"` // Total number of liked posts
}

// /followers — Retrieve a Blog's Followers
type BlogFollowers struct {
	TotalUsers int `json:"total_users"` // The number of users currently following the blog
	Users      []struct {
		Name      string `json:"name"`      // The user's name on tumblr
		Following bool   `json:"following"` // Whether the caller is following the user
		URL       string `json:"url"`       // The URL of the user's primary blog
		Updated   int    `json:"updated"`   // The time of the user's most recent post, in seconds since the epoch
	} `json:"users"`
}

type BlogList struct {
	Posts []Post `json:"posts"`
}

// /posts – Retrieve Published Posts
type BlogPosts struct {
	BlogInfo          // Each response includes a blog object that is the equivalent of an /info response.
	Posts      []Post `json:"posts"`
	TotalPosts int    `json:"total_posts"` // The total number of post available for this request, useful for paginating through results
}

type Post struct {
	BlogName    string   `json:"blog_name"`    // The short name used to uniquely identify a blog
	ID          int      `json:"id"`           // The post's unique ID
	PostURL     string   `json:"post_url"`     // The location of the post
	Type        string   `json:"type"`         // The type of post
	Timestamp   int      `json:"timestamp"`    // The time of the post, in seconds since the epoch
	Date        string   `json:"date"`         // The GMT date and time of the post, as a string
	Format      string   `json:"format"`       // The post format: html or markdown
	ReblogKey   string   `json:"reblog_key"`   // The key used to reblog this post
	Tags        []string `json:"tags"`         // Tags applied to the post
	Bookmarklet bool     `json:"bookmarklet"`  // Indicates whether the post was created via the Tumblr bookmarklet
	Mobile      bool     `json:"mobile"`       // Indicates whether the post was created via mobile/email publishing
	SourceURL   string   `json:"source_url"`   // The URL for the source of the content (for quotes, reblogs, etc.)
	SourceTitle string   `json:"source_title"` // The title of the source site
	Liked       bool     `json:"liked"`        // Indicates if a user has already liked a post or not
	State       string   `json:"state"`        // Indicates the current state of the post
	// Text posts
	Title string `json:"title,omitempty"` // The optional title of the post
	Body  string `json:"body,omitempty"`  // The full post body
	// Photo posts
	Caption string `json:"caption,omitempty"` // The user-supplied caption
	Photos  []struct {
		Caption      string `json:"caption,omitempty"` // user supplied caption for the individual photo
		OriginalSize struct {
			Height int    `json:"height,omitempty"` // height of the image
			Width  int    `json:"width,omitempty"`  // width of the image
			URL    string `json:"url,omitempty"`    // location of the photo file
		} `json:"original_size,omitempty"`
		AlternateSizes []struct {
			Height int    `json:"height,omitempty"` // height of the photo
			Width  int    `json:"width,omitempty"`  // width of the photo
			URL    string `json:"url,omitempty"`    // Location of the photo file
		} `json:"alt_sizes,omitempty"` // alternate photo sizes
	} `json:"photos,omitempty"`
	// Quote posts
	Text   string `json:"text,omitempty"`   // The text of the quote
	Source string `json:"source,omitempty"` // Full HTML for the source of the quote
	// Link posts
	URL         string `json:"url,omitempty"`         // The link
	Author      string `json:"author,omitempty"`      // The author of the article the link points to
	Excerpt     string `json:"excerpt,omitempty"`     // An excerpt from the article the link points to
	Publisher   string `json:"publisher,omitempty"`   // The publisher of the article the link points to
	Description string `json:"description,omitempty"` // A user-supplied description
	// Chat posts
	Dialogue []struct {
		Name   string `json:"name,omitempty"`   // name of the speaker
		Label  string `json:"label,omitempty"`  // label of the speaker
		Phrase string `json:"phrase,omitempty"` // text
	} `json:"dialogue,omitempty"`
	// Audio posts
	AudioPlayer string `json:"player,omitempty"`       // HTML for embedding the audio player
	PlayCount   int    `json:"plays,omitempty"`        // Number of times the audio post has been played
	AlbumArt    string `json:"album_art,omitempty"`    // Location of the audio file's ID3 album art image
	Artist      string `json:"artist,omitempty"`       // The audio file's ID3 artist value
	Album       string `json:"album,omitempty"`        // The audio file's ID3 album value
	TrackName   string `json:"track_name,omitempty"`   // The audio file's ID3 title value
	TrackNumber int    `json:"track_number,omitempty"` // The audio file's ID3 track value
	Year        int    `json:"year,omitempty"`         // The audio file's ID3 year value
	// Video posts
	Player []struct {
		Width     int    `json:"width,omitempty"`      // the width of the video player
		EmbedCode string `json:"embed_code,omitempty"` // HTML for embedding the video player
	} `json:"player,omitempty"`
	// Answer posts
	AskingName string `json:"asking_name,omitempty"` // The blog name of the user asking the question
	AskingURL  string `json:"asking_url,omitempty"`  // The blog URL of the user asking the question
	Question   string `json:"question,omitempty"`    // The question being asked
	Answer     string `json:"answer,omitempty"`      // The answer given
}

// /user/info – Get a User's Information
type UserInfo struct {
	User struct {
		Following         int    `json:"following"`           // The number of blogs the user is following
		DefaultPostFormat string `json:"default_post_format"` // The default posting format - html, markdown or raw
		Name              string `json:"name"`                // The user's tumblr short name
		Likes             int    `json:"likes"`               // The total count of the user's likes
		Blogs             []struct {
			Name      string `json:"name"`      // the short name of the blog
			URL       string `json:"url"`       // the URL of the blog
			Title     string `json:"title"`     // the title of the blog
			Primary   bool   `json:"primary"`   // indicates if this is the user's primary blog
			Followers int    `json:"followers"` // total count of followers for this blog
			Tweet     string `json:"tweet"`     // indicate if posts are tweeted auto, Y, N
			Facebook  string `json:"facebook"`  // indicate if posts are sent to facebook Y, N
			Type      string `json:"type"`      // indicates whether a blog is public or private
		} `json:"blogs"` // Each item is a blog the user has permissions to post to
	} `json:"user"`
}

// /user/following
type UserFollowing struct {
	TotalBlogs int `json:"total_blogs"` // The number of blogs the user is following
	Blogs      []struct {
		Name        string `json:"name"`        // the user name attached to the blog that's being followed
		URL         string `json:"url"`         // the URL of the blog that's being followed
		Updated     int    `json:"updated"`     // the time of the most recent post, in seconds since the epoch
		Title       string `json:"title"`       // the title of the blog
		Description string `json:"description"` // the description of the blog
	} `json:"blogs"`
}
