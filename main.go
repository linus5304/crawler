package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	html := `
			<html>
				<body>
					<a href="/path/one">
						<span>Boot.dev</span>
					</a>
					<a href="https://other.com/path/one">
						<span>Boot.dev</span>
					</a>
				</body>
			</html>
			`
	rawBaseURL := "https://blog.boot.dev"
	getURLsFromHTML(html, rawBaseURL)
}
