# Go CMS Editor for a Basic Blog

This is a lightweight CMS backend written in Go, designed to create and manage blog posts stored as Markdown files with frontmatter. It's meant to be run alongside a Next.js frontend that statically renders these posts.

It is currently used as the CMS for posts for [my personal website](https://marcbachan.com), and is meant to be used alongside a Next.js blog model like [this one](https://vercel.com/templates/blog/blog-starter-kit).

I wrote this both as a practice exercise with Go, as well as a means of consolidating my content management process. 

---

## Structure

```
cms/
├── config.json          # Global settings for post/image directories
├── main.go              # Entry point
├── handlers/            # HTTP handlers (API + UI)
├── storage/             # Markdown file writer
├── model/               # Shared struct definitions
├── utils/               # Helpers (e.g., slugify)
├── templates/           # HTMX-powered HTML UI
├── posts/               # Generated markdown posts
├── images/              # Uploaded image assets
└── README.md            # You're here!
```

---

## Usage

### Configuration

Edit `cms/config.json` to define global output paths:

```json
{
  "postsDir": "posts",
  "imagesDir": "images"
}
```

---

### Local Development

#### Run only the CMS:

```bash
cd cms
go run main.go
```

#### Run CMS alongside Next.js

From the root:

```bash
yarn dev
```

> This runs both:
>
> * Next.js dev server
> * Go CMS server (`cms/main.go`)
>   via `concurrently`

---

## Features

### Create posts

* POST `/api/posts` (via form or JSON)
* Stores `.md` files with frontmatter like:

```yaml
---
title: "My first post"
excerpt: "a byline or summary text"
coverImage: "/images/my_first_post/photo.jpg"
date: "2025-06-23"
ogImage:
  url: "/images/my_first_post/photo.jpg"
tags: [go, blog, photo]
---
```

* Markdown body text follows frontmatter

---

### Drag-and-Drop Image Upload

* Drag image into the form
* Automatically uploads to:

  ```
  images/<slugified_title>/<filename>
  ```
* Returns web path like:

  ```
  /assets/img/<slugified_title>/<filename>
  ```
* Auto-fills `coverImage` and `ogImage.url` fields in form

---

## UI Access

Visit [`http://localhost:8080/new`](http://localhost:8080/new) to use the HTMX-based editor form.

---

## 🐳 Docker Support

You can also build and run the CMS as a Docker container:

```bash
cd cms
docker build -t cms-editor .
docker run -p 8080:8080 cms-editor
```

---

## Output Paths

By default:

* Markdown posts go to `cms/posts/*.md`
* Uploaded images go to `cms/images/<slug>/*`

Make sure your Next.js project reads from these paths (or symlink them if needed).
---

## Requirements

* Go 1.21+
* (Optional) `concurrently` in your Next.js project: `yarn add -D concurrently`