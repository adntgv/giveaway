package giveaway

import (
	"fmt"
	"strings"

	ig "github.com/ahmdrz/goinsta"
)

type InstagramApp interface {
	GetLikers() ([]string, error)
	GetCommenters() ([]string, error)
	String() string
}

type Session struct {
	Shortcode     string    `json:"shotrcode"`
	PostID        string    `json:"post_id"`
	CommentsCount int       `json:"comments_count"`
	LikesCount    int       `json:"likes_count"`
	Commenters    []ig.User `json:"commenters"`
	Likers        []ig.User `json:"likers"`
}

// App struct holding all info regarding a give away
type App struct {
	//ex: instagram.com/p/{shortcode}/...
	insta   *ig.Instagram
	session Session `json:"session"`
}

func (app *App) GetCommenters() ([]string, error) {
	return getUniqueUsernames(app.session.Commenters), nil
}

func (app *App) GetLikers() ([]string, error) {
	return getUniqueUsernames(app.session.Likers), nil
}

func getUniqueUsernames(users []ig.User) []string {
	var usernames []string
	unique := make(map[string]bool, len(users))
	for _, c := range users {
		username := c.Username
		if _, exists := unique[username]; !exists {
			unique[username] = true
			usernames = append(usernames, username)
		}
	}
	return usernames
}

// DefaultApp returns an instance of an app
func DefaultApp(login, password, shortcode string) (InstagramApp, error) {
	insta := ig.New(login, password)
	if err := insta.Login(); err != nil {
		return nil, err
	}
	app := &App{
		insta: insta,
		session: Session{
			Shortcode: shortcode,
			PostID:    ShortcodeToInstaID(shortcode),
		},
	}

	if err := app.FillLikers(); err != nil {
		return nil, err
	}
	if err := app.FillCommenters(); err != nil {
		return nil, err
	}

	return app, nil
}

// FillCommenters of a session
func (app *App) FillCommenters() error {
	fm, err := app.insta.GetMedia(app.session.PostID)
	if err != nil {
		return err
	}
	fm.Sync()
	var commenters []ig.User
	for _, item := range fm.Items {
		app.session.CommentsCount = item.CommentCount
		item.Comments.Sync()
		for item.Comments.Next() {
			comments := item.Comments.Items
			for _, c := range comments {
				item.Comments.Sync()
				commenters = append(commenters, c.User)
			}

		}
	}
	app.session.Commenters = commenters
	return nil
}

// FillLikers of a session
func (app *App) FillLikers() error {
	fm, err := app.insta.GetMedia(app.session.PostID)
	if err != nil {
		return err
	}
	fm.Sync()
	var likers []ig.User
	for _, item := range fm.Items {
		item.SyncLikers()
		app.session.LikesCount = item.Likes
		likers = item.Likers
	}
	app.session.Likers = likers
	return nil
}

func (app *App) String() string {
	str := ""
	return str
}

// ShortcodeToInstaID transforms shortcode to id
// ex: instagram.com/p/{shortcode}/...
//
func ShortcodeToInstaID(Shortcode string) string {
	id := 0
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"
	for _, c := range Shortcode {
		id = (id * 64) + strings.Index(alphabet, string(c))

	}
	return fmt.Sprint(id)
}
