# design

the original static simply converted amber template files to html and stored this state to a json file.  this forces a one to one correlation between amber and html and also means that content can't be done in markdown

## features

* content is written in markdown
* layout and look and feel are done in amber, this reduces the number of boilerplate code necessary for each article since most of a page's look and feel doesn't change with content updates.
* source template files can go in its own repository or a secondary repository, compiled sources get placed in a github pages directory
* compilation happens with github actions, no need for a cli tool, all you do you is commit code (major difference between hugo)
* a general page which lists all content and shows list form of the latest content which should be automatically generated, this is like your homepage with links to all content
* a single page of the article should also be generated where you can view a non-summarized, original view of the content, use custom routing to access the page.  in other words, all content is loaded.
* single page architecture, this generates less html, in fact, just a single html page.
* images should be stored in other locations referenced through links, not in the repository itself.
* i guess this means there's no need for a static cli so this repository is useless.  maybe a web app with oauth is all that's needed.  no install necessary of a cli.


## outline

* version
* compile
  * output
  * source
  * force
* test
* clean
* init
* config

## workflow

1. STATIC_SOURCE and STATIC_DEST env variables should be defined
1. list files in STATIC_SOURCE, nested files need to be considered
1. since this is a single page article, all article conversions will be stored in the same html file.  each article will be referenced by hashtag, there will be an overall list of articles that can be navigated, as soon as they click a particular article, the raw contents will be displayed, this is a bit of js magic.
1. in order to create the overall index, all markdown files need to be iterated first, then the amber files, amber files provide the structure and layout of the page

## html layout

* content needs to be embedded in the html, but not displayed

## next steps

* investigate github actions


