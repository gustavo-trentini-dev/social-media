package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"frontend/src/config"
	"frontend/src/requests"
	"net/http"
	"time"
)

type User struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Nick      string    `json:"nick"`
	CreatedAt time.Time `json:"createdAt"`
	Followers []User    `json:"followers"`
	Following []User    `json:"following"`
	Posts     []Post    `json:"posts"`
}

func GetFullUser(userID uint64, r *http.Request) (User, error) {
	userChannel := make(chan User)
	followerChanel := make(chan []User)
	followingChanel := make(chan []User)
	postsChanel := make(chan []Post)

	go GetUser(userChannel, userID, r)
	go GetFollowers(followerChanel, userID, r)
	go GetFollowing(followingChanel, userID, r)
	go GetPosts(postsChanel, userID, r)

	var (
		user      User
		followers []User
		following []User
		posts     []Post
	)

	for i := 0; i < 4; i++ {
		select {
		case userLoaded := <-userChannel:
			if userLoaded.ID == 0 {
				return User{}, errors.New("Error searching user")
			}

			user = userLoaded
		case followersLoaded := <-followerChanel:
			if followersLoaded == nil {
				return User{}, errors.New("Error searching followers")
			}

			followers = followersLoaded

		case followingLoaded := <-followingChanel:
			if followingLoaded == nil {
				return User{}, errors.New("Error searching following")
			}

			following = followingLoaded

		case postsLoaded := <-postsChanel:
			if postsLoaded == nil {
				return User{}, errors.New("Error searching posts")
			}

			posts = postsLoaded
		}
	}

	user.Followers = followers
	user.Following = following
	user.Posts = posts

	return user, nil
}

func GetUser(channel chan<- User, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d", config.API_URL, userID)
	resp, err := requests.DoRequestWithAuth(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- User{}
		return
	}
	defer resp.Body.Close()

	var user User
	if err = json.NewDecoder(resp.Body).Decode(&user); err != nil {
		channel <- User{}
		return
	}

	channel <- user
}

func GetFollowers(channel chan<- []User, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/followers", config.API_URL, userID)
	resp, err := requests.DoRequestWithAuth(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}
	defer resp.Body.Close()

	var followers []User
	if err = json.NewDecoder(resp.Body).Decode(&followers); err != nil {
		channel <- nil
		return
	}

	if followers == nil {
		channel <- make([]User, 0)
		return
	}

	channel <- followers
}

func GetFollowing(channel chan<- []User, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/following", config.API_URL, userID)
	resp, err := requests.DoRequestWithAuth(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}
	defer resp.Body.Close()

	var following []User
	if err = json.NewDecoder(resp.Body).Decode(&following); err != nil {
		channel <- nil
		return
	}

	if following == nil {
		channel <- make([]User, 0)
		return
	}

	channel <- following
}

func GetPosts(channel chan<- []Post, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/posts", config.API_URL, userID)
	resp, err := requests.DoRequestWithAuth(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}
	defer resp.Body.Close()

	var posts []Post
	if err = json.NewDecoder(resp.Body).Decode(&posts); err != nil {
		channel <- nil
		return
	}

	if posts == nil {
		channel <- make([]Post, 0)
		return
	}

	channel <- posts
}
