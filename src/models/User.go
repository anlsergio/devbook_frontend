package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"webapp/src/config"
	"webapp/src/requests"
)

// User - represents a person signed up to DevBook
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Username  string    `json:"username,omitempty"`
	Email     string    `json:"email,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	Followers []User    `json:"followers"`
	Following []User    `json:"following"`
	Posts     []Post    `json:"posts"`
}

// GetFullUserProfile - calls multiple API endpoints by making use of concurrency, in order to fetch all data regarding a given user
func GetFullUserProfile(userID uint64, r *http.Request) (User, error) {
	userChannel := make(chan User)
	followersChannel := make(chan []User)
	followingChannel := make(chan []User)
	postsChannel := make(chan []Post)

	go FetchUserDataByUserID(userChannel, userID, r)
	go GetFollowersByUserID(followersChannel, userID, r)
	go GetFollowingByUserID(followingChannel, userID, r)
	go GetPostsByUserID(postsChannel, userID, r)

	var (
		user      User
		followers []User
		following []User
		posts     []Post
	)

	// i < 4, where 4 is the number of API calls that needs to be completed
	for i := 0; i < 4; i++ {
		select {
		case userReady := <-userChannel:
			// 0 is the default value of User.ID of type uin64
			if userReady.ID == 0 {
				return User{}, errors.New("something went wrong while trying to fetch the User data")
			}

			user = userReady

		case followersReady := <-followersChannel:
			if followersReady == nil {
				return User{}, errors.New("something went wrong while trying to fetch the dataset of followers")
			}

			followers = followersReady

		case followingReady := <-followingChannel:
			if followingReady == nil {
				return User{}, errors.New("something went wrong while trying to fetch the dataset of users being followed")
			}

			following = followingReady

		case postsReady := <-postsChannel:
			if postsReady == nil {
				return User{}, errors.New("something went wrong while trying to fetch the dataset of posts")
			}

			posts = postsReady
		}
	}

	user.Followers = followers
	user.Following = following
	user.Posts = posts

	return user, nil
}

// FetchUserDataByUserID - calls the API in order to get data from a given user
func FetchUserDataByUserID(c chan<- User, userID uint64, r *http.Request) {
	endpoint_URL := fmt.Sprintf("%s/users/%d", config.APIURL, userID)
	response, err := requests.RequestWithAuthentication(r, http.MethodGet, endpoint_URL, nil)
	if err != nil {
		c <- User{}
		return
	}
	defer response.Body.Close()

	var user User
	if err = json.NewDecoder(response.Body).Decode(&user); err != nil {
		c <- User{}
		return
	}

	c <- user
}

// GetFollowersByUserID - calls the API in order to get users that are following a given user
func GetFollowersByUserID(c chan<- []User, userID uint64, r *http.Request) {
	endpoint_URL := fmt.Sprintf("%s/users/%d/followers", config.APIURL, userID)
	response, err := requests.RequestWithAuthentication(r, http.MethodGet, endpoint_URL, nil)
	if err != nil {
		c <- nil
		return
	}
	defer response.Body.Close()

	var followers []User
	if err = json.NewDecoder(response.Body).Decode(&followers); err != nil {
		c <- nil
		return
	}

	c <- followers
}

// GetFollowingByUserID - calls the API in order to get users that a given user is currently following
func GetFollowingByUserID(c chan<- []User, userID uint64, r *http.Request) {
	endpoint_URL := fmt.Sprintf("%s/users/%d/following", config.APIURL, userID)
	response, err := requests.RequestWithAuthentication(r, http.MethodGet, endpoint_URL, nil)
	if err != nil {
		c <- nil
		return
	}
	defer response.Body.Close()

	var following []User
	if err = json.NewDecoder(response.Body).Decode(&following); err != nil {
		c <- nil
		return
	}

	c <- following
}

// GetPostsByUserID - calls the API in order to get posts from a given user
func GetPostsByUserID(c chan<- []Post, userID uint64, r *http.Request) {
	endpoint_URL := fmt.Sprintf("%s/users/%d/posts", config.APIURL, userID)
	response, err := requests.RequestWithAuthentication(r, http.MethodGet, endpoint_URL, nil)
	if err != nil {
		c <- nil
		return
	}
	defer response.Body.Close()

	var posts []Post
	if err = json.NewDecoder(response.Body).Decode(&posts); err != nil {
		c <- nil
		return
	}

	c <- posts
}
