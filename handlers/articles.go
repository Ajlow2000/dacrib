package handlers

import (
	"net/http"
	"strings"
	"time"

	"github.com/Ajlow2000/dacrib/components"
)

// In-Memory Representation of an article
type Article struct {
	title         string
	datePublished time.Time
	body          string
}

// A structured key
type articleKey string

// Get a structured key derived from the article title.
// Enforces a simple pattern for consistent urls
func (a Article) getKey() articleKey {
	key := strings.ToLower(a.title)
	key = strings.ReplaceAll(key, " ", "-")
	return articleKey(key)
}

// A collection of articles to be served
var MyArticles = map[articleKey]Article{}

// Handle the base page for articles
func Articles(w http.ResponseWriter, r *http.Request) {
	components.Articles().Render(r.Context(), w)
}

// Handle article specific pages
func ShowArticle(w http.ResponseWriter, r *http.Request) {
	/////////// temp /////////////
	MyArticles = make(map[articleKey]Article, 50)

	article_foo := Article{
		title:         "Foo",
		datePublished: time.Now(),
		body:          "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus fermentum, arcu ut euismod iaculis, dui nulla pulvinar neque, ut ultrices sem nisi vel ante. Nulla rutrum libero at lectus varius dictum. Donec a dolor nibh. Aenean mollis nulla eget semper placerat. Ut finibus felis eget urna finibus, vel congue ipsum vulputate. Aliquam gravida dolor eget neque efficitur tincidunt. Cras ut felis ut magna interdum euismod eu nec dui. Phasellus nec porttitor purus. Vivamus id nisl nec nisl facilisis feugiat sit amet vitae orci. Cras aliquam convallis nisi et eleifend. Ut sed gravida tortor. Suspendisse ante tellus, euismod id placerat sit amet, varius ac dolor. Morbi eu fringilla sem. Proin in massa sagittis, imperdiet nunc accumsan, placerat diam",
	}

	article_bar := Article{
		title:         "Bar",
		datePublished: time.Now(),
		body:          "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus fermentum, arcu ut euismod iaculis, dui nulla pulvinar neque, ut ultrices sem nisi vel ante. Nulla rutrum libero at lectus varius dictum. Donec a dolor nibh. Aenean mollis nulla eget semper placerat. Ut finibus felis eget urna finibus, vel congue ipsum vulputate. Aliquam gravida dolor eget neque efficitur tincidunt. Cras ut felis ut magna interdum euismod eu nec dui. Phasellus nec porttitor purus. Vivamus id nisl nec nisl facilisis feugiat sit amet vitae orci. Cras aliquam convallis nisi et eleifend. Ut sed gravida tortor. Suspendisse ante tellus, euismod id placerat sit amet, varius ac dolor. Morbi eu fringilla sem. Proin in massa sagittis, imperdiet nunc accumsan, placerat diam",
	}

	article_foobar := Article{
		title:         "Foo Bar",
		datePublished: time.Now(),
		body:          "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus fermentum, arcu ut euismod iaculis, dui nulla pulvinar neque, ut ultrices sem nisi vel ante. Nulla rutrum libero at lectus varius dictum. Donec a dolor nibh. Aenean mollis nulla eget semper placerat. Ut finibus felis eget urna finibus, vel congue ipsum vulputate. Aliquam gravida dolor eget neque efficitur tincidunt. Cras ut felis ut magna interdum euismod eu nec dui. Phasellus nec porttitor purus. Vivamus id nisl nec nisl facilisis feugiat sit amet vitae orci. Cras aliquam convallis nisi et eleifend. Ut sed gravida tortor. Suspendisse ante tellus, euismod id placerat sit amet, varius ac dolor. Morbi eu fringilla sem. Proin in massa sagittis, imperdiet nunc accumsan, placerat diam",
	}

	MyArticles[article_foo.getKey()] = article_foo
	MyArticles[article_bar.getKey()] = article_bar
	MyArticles[article_foobar.getKey()] = article_foobar
	//////////////////////////////

	url := strings.Split(r.URL.Path, "/")
	requestedArticle := url[len(url)-1]

	article, ok := MyArticles[articleKey(requestedArticle)]
	if !ok {
		NotFoundGoArticles(w, r)

	} else {
		components.ShowArticle(article.body).Render(r.Context(), w)
	}

}
