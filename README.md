# Go URL Shortener with Background Processing

This is a Go URL Shortener project. This is a fun, small-scale project built to get a better understanding of GO, especially its concurrency features, and connect it with a simple JavaScript frontend.

## What’s the Project About?

The idea is simple:

- **Shorten URLs:** You enter a long URL, and the backend gives you a shorter version.
- **Redirect:** When you visit the short URL, it takes you to the original page.
- **Background Processing:** Every time someone uses a short URL, the server updates click analytics in the background using Go routines.

## How Does It Work?

1. **The Go Backend:**
   - Acts as an HTTP server.
   - Offers a few endpoints: one for shortening URLs, one for redirecting, and one for fetching analytics.
   - Uses Go routines to handle click counting in the background, keeping your main redirect flow super fast.

2. **The JavaScript Frontend:**
   - A simple HTML page with an input field.
   - When you submit a URL, JavaScript makes an HTTP request to the Go server.
   - Displays the resulting short URL so you can use it immediately.

3. **Local Connection:**
   - The Go server runs locally (e.g on `http://localhost:8080`).
   - Your browser loads the HTML page (either served by Go or opened directly), and the JavaScript communicates with your Go API through standard HTTP requests.

## Getting Started

- **Set Up:** Make sure you have Go installed. You can run the server locally on your machine.
- **Run the Server:** Run the Go app to start the backend.
- **Open the Frontend:** Either serve your HTML from the Go server or open it directly in your browser.
- **Test It Out:** Enter a URL in the form, get your short URL.
