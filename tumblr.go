package tumblr

import (
	"github.com/kurrik/oauth1a"
	"net/url"
	"strconv"
)

const (
	apiBlogUrl      = "http://api.tumblr.com/v2/blog/"            // api url to get blog data or write to a blog
	apiUserUrl      = "http://api.tumblr.com/v2/user/"            // api url to get user data or perform actions
	apiTaggedUrl    = "http://api.tumblr.com/v2/tagged?"          // api url to get tagged posts
	requestTokenUrl = "http://www.tumblr.com/oauth/request_token" // oauth request-token URL
	authorizeUrl    = "https://www.tumblr.com/oauth/authorize"    // oauth authorize URL
	accessTokenUrl  = "http://www.tumblr.com/oauth/access_token"  // oauth access-token URL
)

type Tumblr struct {
	oauthService oauth1a.Service    // oauth service used to sign HTTP requests
	config       oauth1a.UserConfig // used within the oauth HTTP signing
	apiKey       string             // consumer key used for certain API requests
}

// This is the initialization method.
// An easy way to get the credentials is to access the interactive console:
// https://api.tumblr.com/console
func New(consumerKey, consumerSecret, oauthKey, oauthSecret string) *Tumblr {
	service := &oauth1a.Service{
		RequestURL:   requestTokenUrl,
		AuthorizeURL: authorizeUrl,
		AccessURL:    accessTokenUrl,
		ClientConfig: &oauth1a.ClientConfig{
			ConsumerKey:    consumerKey,
			ConsumerSecret: consumerSecret,
		},
		Signer: new(oauth1a.HmacSha1Signer),
	}
	config := oauth1a.NewAuthorizedConfig(oauthKey, oauthSecret)
	return &Tumblr{
		oauthService: *service,
		config:       *config,
		apiKey:       consumerKey,
	}
}

// This method returns general information about the blog, such as the title,
// number of posts, and other high-level data.
// blogHostname - The standard or custom blog hostname (e.g., example.tumblr.com, example.com)
func (api Tumblr) BlogInfo(blogHostname string) BlogInfo {
	var blogInfo BlogInfo
	requestURL := apiBlogUrl + blogHostname + "/info"
	api.info(requestURL, &blogInfo)
	return blogInfo
}

// This method returns a URL to a blog's avatar with default size 64
// blogHostname - The standard or custom blog hostname (e.g., example.tumblr.com, example.com)
func (api Tumblr) BlogAvatar(blogHostname string) []byte {
	return api.BlogAvatarAndSize(blogHostname, 64)
}

// This method returns a URL to a blog's avatar with a custom size
// blogHostname - The standard or custom blog hostname (e.g., example.tumblr.com, example.com)
// size - The size of the avatar (square, one value for both length and width).
//        Must be one of the values: 16, 24, 30, 40, 48, 64, 96, 128, 512
func (api Tumblr) BlogAvatarAndSize(blogHostname string, size int) []byte {
	requestURL := apiBlogUrl + blogHostname + "/avatar/" + strconv.Itoa(size)
	return api.rawGet(requestURL)
}

// This method can be used to retrieve the publicly exposed likes from a blog.
// blogHostname - The standard or custom blog hostname (e.g., example.tumblr.com, example.com)
// params - A map of the params that are included in this request. Possible parameters:
//          * limit - The number of results to return.  Default: 20 (1–20, inclusive)
//          * offset - Liked post number to start at.  Default: 0 (First post)
//          * before - Retrieve posts liked before the specified timestamp. Default: None
//          * after - Retrieve posts liked after the specified timestamp. Default: None
func (api Tumblr) BlogLikes(blogHostname string, params map[string]string) Likes {
	var blogLikes Likes
	requestURL := apiBlogUrl + blogHostname + "/likes?"
	urlParams := url.Values{}
	urlParams.Set("api_key", api.apiKey)
	for key, value := range params {
		urlParams.Set(key, value)
	}
	requestURL = requestURL + urlParams.Encode()
	api.info(requestURL, &blogLikes)
	return blogLikes
}

// This method retrieves a blog's followers
// blogHostname - The standard or custom blog hostname (e.g., example.tumblr.com, example.com)
// params - A map of the params that are included in this request. Possible parameters:
//          * limit - The number of results to return.  Default: 20 (1–20, inclusive)
//          * offset - Liked post number to start at.  Default: 0 (First follower)
func (api Tumblr) BlogFollowers(blogHostname string, params map[string]string) BlogFollowers {
	var blogFollowers BlogFollowers
	requestURL := apiBlogUrl + blogHostname + "/followers?"
	urlParams := url.Values{}
	for key, value := range params {
		urlParams.Set(key, value)
	}
	requestURL = requestURL + urlParams.Encode()
	api.info(requestURL, &blogFollowers)
	return blogFollowers
}

// This method retrieves a list of a blog's published posts
// blogHostname - The standard or custom blog hostname (e.g., example.tumblr.com, example.com)
// params - A map of the params that are included in this request. Possible parameters:
//          * type - The type of post to return. text, quote, link, answer, video, audio, photo, chat
//          * id - A specific post ID
//          * tag - Limits the response to posts with the specified tag
//          * limit - The number of posts to return: 1–20, inclusive
//          * offset - Post number to start at
//          * reblog_info - Indicates whether to return reblog information (specify true or false)
//          * notes_info - Indicates whether to return notes information (specify true or false).
//          * filter - Specifies the post format to return, other than HTML (text or raw)
func (api Tumblr) BlogPosts(blogHostname string, params map[string]string) BlogPosts {
	var blogPosts BlogPosts
	requestURL := apiBlogUrl + blogHostname + "/posts?"
	urlParams := url.Values{}
	urlParams.Set("api_key", api.apiKey)
	for key, value := range params {
		urlParams.Set(key, value)
	}
	requestURL = requestURL + urlParams.Encode()
	api.info(requestURL, &blogPosts)
	return blogPosts
}

// This method retrieves a list of a blog's queued posts.
// blogHostname - The standard or custom blog hostname (e.g., example.tumblr.com, example.com)
// params - A map of the params that are included in this request. Possible parameters:
//          * offset - Post number to start at (Default: 0)
//          * limit - The number of results to return: 1–20, inclusive.
//          * filter - Specifies the post format to return, other than HTML (text or raw)
func (api Tumblr) BlogQueuedPosts(blogHostname string, params map[string]string) BlogList {
	var queuedPosts BlogList
	requestURL := apiBlogUrl + blogHostname + "/posts/queue?"
	urlParams := url.Values{}
	for key, value := range params {
		urlParams.Set(key, value)
	}
	requestURL = requestURL + urlParams.Encode()
	api.info(requestURL, &queuedPosts)
	return queuedPosts
}

// This method is used to post a blog post to a blog
// blogHostname - The standard or custom blog hostname (e.g., example.tumblr.com, example.com)
// params - A map of the params that are included in this request. Possible parameters:
//          * type - The type of post to create. Specify one of the following:
//                   text, photo, quote, link, chat, audio, video
//          * state - The state of the post. Specify one of the following:  published, draft, queue, private
//          * tags - Comma-separated tags for this post
//          * tweet - Manages the autotweet (if enabled) for this post: set to off for no tweet,
//                    or enter text to override the default tweet
//          * date - The GMT date and time of the post, as a string
//          * format - Sets the format type of post. Supported formats are: html & markdown
//          * slug - Add a short text summary to the end of the post URL
//           TEXT POSTS
//          * title - The optional title of the post, HTML entities must be escaped
//          * body - The full post body, HTML allowed
//           PHOTO POSTS
//          * caption - The user-supplied caption, HTML allowed
//          * link - The "click-through URL" for the photo
//          * source - The photo source URL
//          * data - One or more image files (submit multiple times to create a slide show)
//           QUOTE POSTS
//          * quote - The full text of the quote, HTML entities must be escpaed
//          * source - Cited source, HTML allowed
//           LINK POSTS
//          * title - The title of the page the link points to, HTML entities should be escaped.
//          * url - The link
//          * description - A user-supplied description, HTML allowed.
//           CHAT POSTS
//          * title - The title of the chat
//          * conversation - The text of the conversation/chat, with dialogue labels (no HTML).
//           AUDIO POSTS
//          * caption - The user-supplied caption
//          * external_url - The URL of the site that hosts the audio file (not tumblr)
//          * data - An audio file
//           VIDEO POSTS
//          * caption - The user-supplied caption
//          * embed - HTML embed code for the video
//          * data - A video file
func (api Tumblr) Post(blogHostname string, params map[string]string) Meta {
	requestURL := apiBlogUrl + blogHostname + "/post"
	urlParams := url.Values{}
	for key, value := range params {
		urlParams.Set(key, value)
	}
	response := api.post(requestURL, urlParams.Encode())
	return response.Meta
}

// This method is used to edit a blog post to a blog
// blogHostname - The standard or custom blog hostname (e.g., example.tumblr.com, example.com)
// id - The id of the blog post
// params - The list of possible parameters are listed above the Post method
func (api Tumblr) PostEdit(blogHostname string, id int, params map[string]string) Meta {
	requestURL := apiBlogUrl + blogHostname + "/post/edit"
	urlParams := url.Values{}
	urlParams.Set("id", strconv.Itoa(id))
	for key, value := range params {
		urlParams.Set(key, value)
	}
	response := api.post(requestURL, urlParams.Encode())
	return response.Meta
}

// This method is used to reblog a blog post to a blog
// blogHostname - The standard or custom blog hostname (e.g., example.tumblr.com, example.com)
// id - The ID of the reblogged post
// reblogKey - The reblog key for the reblogged post – get the reblog key with a BlogPosts request
// params - The list of possible parameters are listed above the Post method, along with:
//          * comment - A comment added to the reblogged post
func (api Tumblr) PostReblog(blogHostname string, id int, reblogKey string, params map[string]string) Meta {
	requestURL := apiBlogUrl + blogHostname + "/post/reblog"
	urlParams := url.Values{}
	urlParams.Set("id", strconv.Itoa(id))
	urlParams.Set("reblog_key", reblogKey)
	for key, value := range params {
		urlParams.Set(key, value)
	}
	response := api.post(requestURL, urlParams.Encode())
	return response.Meta
}

// This method is used to delete a blog post from a blog
// blogHostname - The standard or custom blog hostname (e.g., example.tumblr.com, example.com)
// id - The ID of the post to delete
func (api Tumblr) PostDelete(blogHostname string, id int) Meta {
	requestURL := apiBlogUrl + blogHostname + "/post/delete"
	urlParams := url.Values{}
	urlParams.Set("id", strconv.Itoa(id))
	response := api.post(requestURL, urlParams.Encode())
	return response.Meta
}

// This method is used to retrieve the user's account information that matches
// the OAuth credentials submitted with the request.
func (api Tumblr) UserInfo() UserInfo {
	var userInfo UserInfo
	requestURL := apiUserUrl + "info"
	api.info(requestURL, &userInfo)
	return userInfo
}

// This method is used to retrieve the dashboard that matches the OAuth credentials
// submitted with the request.
// params - A map of the params that are included in this request. Possible parameters:
//          * limit - The number of results to return: 1–20, inclusive
//          * offset - Post number to start at
//          * type - The type of post to return. text, photo, quote, link, chat, audio, video, answer
//          * since_id - Return posts that have appeared after this ID
//          * reblog_info - Indicates whether to return reblog information (specify true or false).
//          * notes_info - Indicates whether to return notes information (specify true or false).
func (api Tumblr) UserDashboard(params map[string]string) BlogList {
	var userDashboard BlogList
	requestURL := apiUserUrl + "dashboard?"
	urlParams := url.Values{}
	for key, value := range params {
		urlParams.Set(key, value)
	}
	requestURL = requestURL + urlParams.Encode()
	api.info(requestURL, &userDashboard)
	return userDashboard
}

// This method can be used to retrieve the publicly exposed likes from a blog.
// params - A map of the params that are included in this request. Possible parameters:
//          * limit - The number of results to return.  Default: 20 (1–20, inclusive)
//          * offset - Liked post number to start at.  Default: 0 (First post)
//          * before - Retrieve posts liked before the specified timestamp. Default: None
//          * after - Retrieve posts liked after the specified timestamp. Default: None
func (api Tumblr) UserLikes(params map[string]string) Likes {
	var userLikes Likes
	requestURL := apiUserUrl + "likes?"
	urlParams := url.Values{}
	for key, value := range params {
		urlParams.Set(key, value)
	}
	requestURL = requestURL + urlParams.Encode()
	api.info(requestURL, &userLikes)
	return userLikes
}

// This method is used to retrieve the blogs followed by the user whose OAuth credentials
// are submitted with the request.
// params - A map of the params that are included in this request. Possible parameters:
//          * limit - The number of results to return.  Default: 20 (1–20, inclusive)
//          * offset - Liked post number to start at.  Default: 0 (First post)
func (api Tumblr) UserFollowing(params map[string]string) UserFollowing {
	var userFollowing UserFollowing
	requestURL := apiUserUrl + "following?"
	urlParams := url.Values{}
	for key, value := range params {
		urlParams.Set(key, value)
	}
	requestURL = requestURL + urlParams.Encode()
	api.info(requestURL, &userFollowing)
	return userFollowing
}

// This method is used to follow a specific URL
// followURL - The url to follow, formatted (blogname.tumblr.com, blogname.com)
func (api Tumblr) UserFollow(followURL string) Meta {
	requestURL := apiUserUrl + "follow"
	urlParams := url.Values{}
	urlParams.Set("url", followURL)
	response := api.post(requestURL, urlParams.Encode())
	return response.Meta
}

// This method is used to unfollow a specific URL
// unfollowURL - The url to unfollow, formatted (blogname.tumblr.com, blogname.com)
func (api Tumblr) UserUnfollow(unfollowURL string) Meta {
	requestURL := apiUserUrl + "unfollow"
	urlParams := url.Values{}
	urlParams.Set("url", unfollowURL)
	response := api.post(requestURL, urlParams.Encode())
	return response.Meta
}

// This method is used to like a specific blog post
// id - The ID of the blog post to be liked
// reblogKey - The reblog key string
func (api Tumblr) UserLike(id int, reblogKey string) Meta {
	requestURL := apiUserUrl + "like"
	urlParams := url.Values{}
	urlParams.Set("id", strconv.Itoa(id))
	urlParams.Set("reblog_key", reblogKey)
	response := api.post(requestURL, urlParams.Encode())
	return response.Meta
}

// This method is used to unlike a specific blog post
// id - The ID of the blog post to be unliked
// reblogKey - The reblog key string
func (api Tumblr) UserUnlike(id int, reblogKey string) Meta {
	requestURL := apiUserUrl + "unlike"
	urlParams := url.Values{}
	urlParams.Set("id", strconv.Itoa(id))
	urlParams.Set("reblog_key", reblogKey)
	response := api.post(requestURL, urlParams.Encode())
	return response.Meta
}

// This method is used to retrieve posts that are tagged with a specified tag.
// tag - The tag on the posts you'd like to retrieve.
// params - A map of the params that are included in this request. Possible parameters:
//          * before - The timestamp of when you'd like to see posts before.
//                     If the Tag is a "featured" tag, use the "featured_timestamp"
//                     on the post object for pagination.
//          * limit - The number of results to return: 1–20, inclusive
//          * filter - Specifies the post format to return, other than HTML (text or raw)
func (api Tumblr) TaggedPosts(tag string, params map[string]string) []Post {
	var taggedPosts []Post
	requestURL := apiTaggedUrl
	urlParams := url.Values{}
	urlParams.Set("tag", tag)
	urlParams.Set("api_key", api.apiKey)
	for key, value := range params {
		urlParams.Set(key, value)
	}
	requestURL = requestURL + urlParams.Encode()
	api.info(requestURL, &taggedPosts)
	return taggedPosts
}
