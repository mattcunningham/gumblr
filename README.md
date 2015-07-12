# Tumblr
[![License MIT](https://img.shields.io/npm/l/express.svg)](http://opensource.org/licenses/MIT)
[![GoDoc](http://godoc.org/github.com/mattcunningham/gumblr?status.png)](http://godoc.org/github.com/mattcunningham/gumblr)

## Installing
    go get github.com/mattcunningham/gumblr

## Creating a client
All Tumblr API calls will be made through the `Tumblr` type.  To create a Tumblr client:

    client := tumblr.New(
        "<Insert Consumer Key>",
        "<Insert Consumer Secret",
        "<Insert Oauth Key>",
        "<Insert Oauth Secret>"
    )
A simple way to receive the necessary credentials is by accessing the Tumblr API console at https://api.tumblr.com/console.

## Supported Methods
### Blog Requests
    client.BlogInfo("staff.tumblr.com")
    client.BlogAvatar("staff.tumblr.com")
    client.BlogAvatarAndSize("staff.tumblr.com", 24)
    client.BlogLikes("staff.tumblr.com", make(map[string]string))
    client.BlogFollowers("staff.tumblr.com", make(map[string]string))
    client.BlogQueuedPosts("staff.tumblr.com", make(map[string]string))
    client.BlogLikes("staff.tumblr.com", make(map[string]string))

### Blog Actions
    client.Post("staff.tumblr.com", make(map[string]string))
    client.PostEdit("staff.tumblr.com", 12345, make(map[string]string))
    client.PostReblog("staff.tumblr.com", 12344321, "r3bl0gk3y", make(map[string]string))
    client.PostDelete("staff.tumblr.com", 4321234)

### User Requests
    client.UserInfo()
    client.UserDashboard(make(map[string]string))
    client.UserLikes(make(map[string]string))
    client.UserFollowing(make(map[string]string))

### User Actions
    client.UserFollow("staff.tumblr.com")
    client.UserUnfollow("staff.tumblr.com")
    client.UserLike(1234431, "r3b10gk3y")
    client.UserUnlike(4321234, "r3b10gk3y")

## Tagged Posts
    client.TaggedPosts("gifs", make(map[string]string))
