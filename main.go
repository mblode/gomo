package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

func check(e error) {
	// Log message with specified arguments.

	if e != nil {
		panic(e)
	}
}

func readFile(fileName string) {
	// Read file and close the file.

	dat, err := ioutil.ReadFile(fileName)
	check(err)
	fmt.Print(string(dat))
}

func writeFile(fileName string, text string) {
	// Write content to file and close the file.

	textAsBytes := []byte(text)
	err := ioutil.WriteFile(fileName, textAsBytes, 0644)
	check(err)
}

func truncateText(text string, numberOfWords int) {
	// Remove tags and truncate text to the specified number of words.

	if numberOfWords == 0 {
		numberOfWords = 25
	}

	// return ' '.join(re.sub('(?s)<.*?>', ' ', text).split()[:words])
}

func readHeaders(text string) {
	// Parse headers in text and yield (key, value, end-index) tuples.
}

func formatDate(dateString string) {
	// Convert yyyy-mm-dd date string to RFC 2822 format date string.
}

func readContent(filename string) {
	// Read content and metadata from file into a dictionary.
}

func renderHTML(template string) {
	// Replace placeholders in template with values from params.
}

func makePages(source string, dist string, layout string) {
	// Generate pages from page content.
}

func makeList(posts string, dist string, listLayout string, itemLayout string) {
	// Generate list page for a blog.

}

func main() {
	// Create a new _site directory from scratch.

    if os.path.isdir('_site'):
        shutil.rmtree('_site')
    shutil.copytree('static', '_site')

    // Default parameters.
    params = {
        'base_path': '',
        'subtitle': 'Lorem Ipsum',
        'author': 'Admin',
        'site_url': 'http://localhost:8000',
        'current_year': datetime.datetime.now().year
    }

    // If params.json exists, load it.
    if os.path.isfile('params.json'):
        params.update(json.loads(fread('params.json')))

    // Load layouts.
    page_layout = fread('layout/page.html')
    post_layout = fread('layout/post.html')
    list_layout = fread('layout/list.html')
    item_layout = fread('layout/item.html')
    feed_xml = fread('layout/feed.xml')
    item_xml = fread('layout/item.xml')

    // Combine layouts to form final layouts.
    post_layout = render(page_layout, content=post_layout)
    list_layout = render(page_layout, content=list_layout)

    // Create site pages.
    make_pages('content/_index.html', '_site/index.html',
               page_layout, **params)
    make_pages('content/[!_]*.html', '_site/{{ slug }}/index.html',
               page_layout, **params)

    // Create blogs.
    blog_posts = make_pages('content/blog/*.md',
                            '_site/blog/{{ slug }}/index.html',
                            post_layout, blog='blog', **params)
    news_posts = make_pages('content/news/*.html',
                            '_site/news/{{ slug }}/index.html',
                            post_layout, blog='news', **params)

    // Create blog list pages.
    make_list(blog_posts, '_site/blog/index.html',
              list_layout, item_layout, blog='blog', title='Blog', **params)
    make_list(news_posts, '_site/news/index.html',
              list_layout, item_layout, blog='news', title='News', **params)

    // Create RSS feeds.
    make_list(blog_posts, '_site/blog/rss.xml',
              feed_xml, item_xml, blog='blog', title='Blog', **params)
    make_list(news_posts, '_site/news/rss.xml',
              feed_xml, item_xml, blog='news', title='News', **params)
}
