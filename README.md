# static
Opinionated static site compiler written in golang

## Description

static is a command tool to aide in the building of web sites where content is generally unchanged, a blog with a commenting system and API backend that keeps a lot of state server-side is not a great use case for static.  However, if you'd like a simple way of creating web content using an efficient templating language then static is for you.  You don't need to write any javascript code or learn any heavyweight frameworks.  All content is compiled ahead of time and can be served from various web servers like apache or reverse proxies like nginx.  Content and layout should be the focus of your site and static takes care of most of the other meanial tasks.

## Features

1.  Compile .amber, .markdown, .md files into .html
1.  Allow for testing of static content locally before staging to production
1.  Facilitate staging of sites

## Requirements

1.  golang 1.9+
1.  `go get github.com/fatih/color`
1.  `go get gopkg.in/alecthomas/kingpin.v2`
1.  `go get github.com/russross/blackfriday`
1.  `go get github.com/eknkc/amber`

## Usage

### `static init`

Creates the following hierarchy and files in the current working directory:

`
assets/css/site.css
assets/js/config.js
assets/js/site.js
index.amber
indexbar.amber
layout.amber
menubar.amber
.static.clean
.static.conf
`

This command clones a repository from github locally.

### `static config`

Lists configuration stored in .static.conf

`static config`

```javascript
{"exclude":"indexbar.amber,footer.amber, layout.amber"}
```

#### `static config --clear`

Removes all configuration parameters.

`static config --list`

Lists all the configuration parameters.

`static config --exclude="indexbar.amber, footer.amber, layout.amber"`

#### .static.conf

Contains a json structure of all configuration parameters.

Parameter | Description
---|---
exclude | Templates to exclude during compilation

### `static compile`

Converts all .amber files to .html

Parameter | Description
---|---
src | Location of source .amber files
out | Location to store compiled .html files, directory must exist and user must have write permissions.

`static compile --src=test --out=production`

### `static clean`

Removes all compiled (html) files along with .static.clean.

#### `static clean --all`

Removes all files including .static.conf

#### .static.clean

Contains a json structure of all compiled files by static.

### `static version`

Shows version of static.

### `static test`

Runs a mini-http server with content

## FAQ

### Why not just use Hugo?

static is a tool that uses a very specific workflow to generate static sites, all steps are very prescriptive and purposely simplified to keep focus on content.  Users already have a large choice of frameworks like bootstrap, material design lite, jquery, etc, static is not trying to re-invent solutions to those problems, but rather to simplify the workflow for creating and maintaining static sites.  Hugo is a great tool for generating static sites with a wealth of plugins and probably rivals Wordpress in many respects.  

## TODO

- [x] generate .static file that lists all the compiled output files
- [x] configuration file to avoid certain file types like indexbar.amber
- [x] source and destination
- [ ] clone skeleton from github
- [ ] upload to github pages or git repository
- [ ] upload to s3
- [ ] documentation for deploying to web servers
- [ ] automatically generate web server configuration, e.g. nginx.conf
- [ ] letsencrypt integration
- [ ] docker file
- [ ] automatically download assets like material-ui, bootstrap, etc
- [ ] markdown support
