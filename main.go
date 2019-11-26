package main

import (
	"encoding/json"
	"fmt"
	"github.com/otiai10/copy"
	"io/ioutil"
	"os"
	"time"
	"regexp"
)

type configuration struct {
	basePath    string `json:"basePath"`
	subtitle    string `json:"subtitle"`
	author      string `json:"author"`
	siteURL     string `json:"siteURL"`
	currentYear int    `json:"currentYear"`
}

func check(e error) {
	// Log message with specified arguments.

	if e != nil {
		panic(e)
	}
}

func readFile(fileName string) []byte {
	// Read file and close the file.

	file, err := ioutil.ReadFile(fileName)
	check(err)
	fmt.Println(string(file))
	return file
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

	return " ".join(regexp.sub("(?s)<.*?>", " ", text).split()[:words])
}

func readHeaders(text string) {
	// Parse headers in text and yield (key, value, end-index) tuples.

	for match in regexp.finditer(r'\s*<!--\s*(.+?)\s*:\s*(.+?)\s*-->\s*|.+', text):
        if not match.group(1):
            break
        yield match.group(1), match.group(2), match.end()
}

func formatDate(dateString int) {
	// Convert yyyy-mm-dd date string to RFC 2822 format date string.

	// return dateString.Format(time.RFC2822)
	return time.Now().Format(time.RFC2822)
}

func readContent(filename string) {
	// Read content and metadata from file into a dictionary.

	// Read file content.
    text = fread(filename)

    // Read metadata and save it in a dictionary.
    date_slug = os.path.basename(filename).split('.')[0]
    match = regexp.search(r'^(?:(\d\d\d\d-\d\d-\d\d)-)?(.+)$', date_slug)
    content = {
        'date': match.group(1) or '1970-01-01',
        'slug': match.group(2),
    }

    // Read headers.
    end = 0
    for key, val, end in read_headers(text):
        content[key] = val

    // Separate content from headers.
    text = text[end:]

    // Convert Markdown content to HTML.
    if filename.endswith(('.md', '.mkd', '.mkdn', '.mdown', '.markdown')):
        try:
            if _test == 'ImportError':
                raise ImportError('Error forced by test')
            import commonmark
            text = commonmark.commonmark(text)
        except ImportError as e:
            log('WARNING: Cannot render Markdown in {}: {}', filename, str(e))

    // Update the dictionary with content and RFC 2822 date.
    content.update({
        'content': text,
        'rfc_2822_date': rfc_2822_format(content['date'])
    })

    return content
}

func renderHTML(template string, content []byte) {
	// Replace placeholders in template with values froconfiguration.

	return regexp.sub(r'{{\s*([^}\s]+)\s*}}', lambda match: str(params.get(match.group(1), match.group(0))), template)
}

func makePages(source string, dist string, layout string, content []byte) {
	// Generate pages from page content.

	items = []

    for (srcPath in glob.glob(src)) {
        content = read_content(srcPath)

		page_configuration = dicconfiguration, **content)

        // Populate placeholders in content if content-rendering is enabled.
        ipage_configuration.get('render') == 'yes':
		rendered_content = rendepage_configuration['content'], page_configuration)
		page_configuration['content'] = rendered_content
		content['content'] = rendered_content

        items.append(content)

        dstPath = render(dst, page_configuration)
        output = render(layout, page_configuration)

        log('Rendering {} => {} ...', srcPath, dstPath)
        fwrite(dstPath, output)
	}

    return sorted(items, key=lambda x: x['date'], reverse=True)
}

func makeList(posts string, dist string, listLayout string, itemLayout string) {
	// Generate list page for a blog.

	items = []

    for (post in posts) {
		item_configuration = dicconfiguration, **post)
		item_configuration['summary'] = truncate(post['content'])
		item = render(itemLayout, item_configuration)
		items.append(item)
		configuration['content'] = ''.join(items)
		dstPath = render(dst, configuration)
		output = render(listLayout, configuration)

		log('Rendering list => {} ...', dstPath)
		fwrite(dstPath, output)
	}
}

func main() {
	// Create a new _site directory from scratch.

	if _, err := os.Stat("_site"); os.IsNotExist(err) {
		err := os.RemoveAll("_site")
		check(err)
	}

	err := copy.Copy("static", "_site")
	check(err)

	// Default configuration.
	config := []configuration{
		{
			basePath:    "",
			subtitle:    "Lorem Ipsum",
			author:      "Admin",
			siteURL:     "http://localhost:8000",
			currentYear: time.Now().Year(),
		},
	}

	// If config.json exists, load it.
	if _, err := os.Stat("config.json"); err == nil {
		byteValue := readFile("config.json")

		json.Unmarshal(byteValue, &config)
	}

	check(err)

	// Load layouts.
	pageLayout := readFile("layout/page.html")
	postLayout := readFile("layout/post.html")
	listLayout := readFile("layout/list.html")
	itemLayout := readFile("layout/item.html")
	feedXML := readFile("layout/feed.xml")
	itemXML := readFile("layout/item.xml")

	// Combine layouts to form final layouts.
	postLayout := renderHTML(pageLayout, pageLayout)
	listLayout := renderHTML(pageLayout, listLayout)

	// Create site pages.
	makePages("content/_index.html", "_site/index.html", pageLayout, "home", config)
	makePages("content/[!_]*.html", "_site/{{ slug }}/index.html", pageLayout, "home", config)

	// Create blogs.
	blogPosts := makePages("content/blog/*.md", "_site/blog/{{ slug }}/index.html", postLayout, "blog", config)
	newsPosts := makePages("content/news/*.html", "_site/news/{{ slug }}/index.html", postLayout, "news", config)

	// Create blog list pages.
	makeList(blogPosts, "_site/blog/index.html", listLayout, itemLayout, "blog", "Blog", config)
	makeList(newsPosts, "_site/news/index.html", listLayout, itemLayout, "news", "News", config)

	// Create RSS feeds.
	makeList(blogPosts, "_site/blog/rss.xml", feed_xml, item_xml, "blog", "Blog", config)
	makeList(newsPosts, "_site/news/rss.xml", feed_xml, item_xml, "news", "News", config)
}
