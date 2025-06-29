# Go CMS Editor for a Basic Blog

This is a lightweight CMS backend written in Go, designed to create and manage blog posts stored as Markdown files with frontmatter metadata. 

It is currently used as the CMS for posts for [my personal website](https://marcbachan.com), and is meant to be used alongside a blog structure like [this one](https://vercel.com/templates/blog/blog-starter-kit) built in Next.js (which I also use for my site).

I wrote this both as a practice exercise with Go, as well as a means of simplifying my content management process. 

Please note that I did generate portions of this README using LLMs, mainly for tidy formatting.

---

## Features

- Create posts with rich Markdown content and YAML frontmatter
- Drag-and-drop image upload with preview
- Live Markdown preview (with cover image) using `marked.js`
- Cover and OG image URLs populated from uploads
- Images stored in `/assets/img/<slug>/<filename>`
- Posts stored in `postsDir` as `.md` files
- Temporary image staging (`/tmp-preview`) before post is saved
- Edit and delete existing posts
- Filter post list by tag
- Basic login system (session-based auth)
- HTMX-enhanced forms
- Docker + Docker Compose support

---

## Project structure

```
cms/
├── main.go                # Main app entrypoint
├── handlers/              # HTTP handlers
├── model/                 # BlogPost struct
├── storage/               # Markdown I/O helpers
├── templates/             # HTMX HTML templates
├── config/                # Settings loader
├── public/                # Static files (images, CSS)
│   └── tmp-preview/       # Temporary images before post is created
|   └── styles/            # CSS
├── Dockerfile
├── docker-compose.yml
└── README.md

````

---

## Local Development

### Requirements

- Go 1.24+
- (Optional) Docker + Docker Compose

### Using `go run`

```bash
cd cms
go run main.go
````

Visit: [http://localhost:8080](http://localhost:8080)

### With Docker Compose

```bash
docker-compose up --build
```

---

## 🔐 Authentication

* On first run, you'll need to set your user, password, and secret in  an `.env` file::

```
CMS_USER="admin"
CMS_PASS="password"
SESSION_SECRET="super-secret"
```

---

## Creating posts

1. Go to `/new`
2. Fill in title, content, tags, etc.
3. Drag and drop an image
4. Click "Create Post"
5. The Markdown file is saved in `/posts/<slug>.md`
6. The image is moved to `/assets/img/<slug>/<filename>`

---

## Editing posts

* Visit `/edit/<slug>`
* Edit frontmatter, image, or content
* Live preview will render the image and Markdown in real time

---

## Deleting posts

* On the `/posts` list, click the 🗑 button
* Uses `hx-delete` and confirms via prompt

---

## Configuration 

### `config/config.json`

Set these paths to the corresponding directories for posts and images in your project. Multiple directories are not supported - the hope is to handle any custom filtering with tags on the frontend, or just store all content in a database (i.e., Postgres).

```json
{
  "postsDir": "./posts",
  "imagesDir": "../public/assets/img",
}
```

### `.env`

```
CMS_USER="admin"
CMS_PASS="password"
SESSION_SECRET="super-secret"
```


---

## Static File Handling

* Go serves static files from `./public/`
* Images are accessed via `/assets/img/...`
* Temporary uploads are in `/tmp-preview` and cleaned up after post creation
* Add custom styles to `public/styles.css`

---

## Notes

* Markdown parsing uses [marked.js](https://marked.js.org/)
* HTMX powers dynamic form actions (`hx-post`, `hx-delete`, etc.)
* Form values are converted to JSON in JS before submission

---

## 🐳 Docker Notes

To build and run manually:

```bash
docker build -t cms .
docker run -p 8080:8080 cms
```
NOTE: 

---

## Integrate with Next.js

This CMS assumes that the final site (built in Next.js) will:

* Use the `public/` folder as its static root
* Render blog posts from the Markdown files in `/posts`
* Reference images via `/assets/img/...`

You can clone this repo into your Next.js project like:

```bash
git clone https://github.com/marcbachan/go-cms.git cms
```

Then add to your root `package.json`:

```json
"scripts": {
  "dev": "concurrently \"cd cms && go run main.go\" \"next dev\"",
  "build": "next build",
  "start": "next start"
}
```

---

## Ideas for Roadmap

One of the best things about learning Go and building this project is gradually understanding what else can be done with it. These are some ideas I'd like to implement eventually (or work with others on):

* Filter post types by tag
* OAuth or JWT-based auth
* Markdown linting or preview styles
* Integrate a database like Postgres (or at least the option to do this with an optional config for larger projects)
