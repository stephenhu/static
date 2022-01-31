# static
command line tool for building static sites written in golang.

## Description

static is a command tool to aide in the building of web sites where content is generally unchanged, a blog with a commenting system and API backend that keeps a lot of state server-side is not a great use case for static.  A fixed blog or website where all access is read is a great fit for static.

If you'd like a simple way of creating web content using an efficient templating language then static is for you.  You don't need to write any javascript code or learn any heavyweight frameworks though there is flexibility to use these as well.  All content is compiled ahead of time and can be served from various web servers like apache or reverse proxies like nginx.  Content and layout should be the focus of your site and static takes care of most of the other menial tasks.

static also provides a github action that can be used to automate many steps, see: [github.com/stephenhu/static-action]() for more details.

## Features

1.  Compile .amber and .md files into index.html (single page)
1.  static will create an index of all articles and list these like a home page, by clicking on an individual page, static also provides a single view by leveraging css and js magic.
1.  Allow for testing of web content locally before staging to production

## Dependencies

1.  golang 1.11+
1.  `go get github.com/fatih/color`
1.  `go get github.com/spf13/cobra`
1.  `go get github.com/russross/blackfriday`
1.  `go get github.com/eknkc/amber`
1.  `go get github.com/PuerkitoBio/goquery`

## Usage

### .staticignore

this file contains a list of files to ignore when doing a build, by default README.md is the only file ignored.  you can ignore processing of additional files by adding a single name per line.

### `static build`

the build command expects the following files minimally: `layout.amber` and `index.amber`.  you can `import` other .amber files into `index.amber`, but these are the required files.  `layout.amber` and `index.amber` provide the essential layout of your pages and uses amber template markup.  see [github.com/eknkc/amber]() for more details about the markup language.

Parameter | Description
---|---
src | Location of source .amber files
out | Location to store compiled .html files, directory must exist and user must have write permissions.

`static build --src=test --out=production`

### `static test`

Runs your site in a local web server for testing.  after starting the test server, please open up your browser to [http://127.0.0.1:8888]() and you can check all links and the layout, make modifications and see this updated in real time without restarting the browser.


### `static version`

Shows version of static.

## FAQ

### Why not just use Hugo?

static is a tool that uses a very specific workflow to generate static sites, all steps are very prescriptive and purposely simplified to keep focus on content.  Users already have a large choice of frameworks like bootstrap, material design lite, jquery, angularjs, reactjs, vuejs, etc, static is not trying to re-invent solutions to those problems, but rather to simplify the workflow for creating and maintaining static sites.  Hugo is a great tool for generating static sites with a wealth of plugins and probably rivals Wordpress in many respects.
