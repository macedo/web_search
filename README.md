# web_search

this is a command line for searching with Google, DuckDuckGo, Github and StackOverflow services.

## How to Install

You can download web_search binaries from the [release page](https://github.com/macedo/web_search/releases) on this repository.

Move the relevant binary for your OS to `/usr/local/bin` (or other directory that is on your $PATH)

(might require to make the downloaded file executable with `chmod +x`)
 
## Usage

You can use the `web_search` as the following:
 
`$ web_search -engine=<engine> <term>`

For example:

```
web_search -engine=duckduckgo golang command line applications
```

For convenience, is recommended add some alias to your shell.

```
alias google='web_search -engine=google'
alias ddg='web_search -engine=duckduckgo'
alias github='web_search -engine=github'
alias stackoverflow='web_search -engine=stackoverflow'
```

## Motivation

* Want to build a command line (that have use) using golang

If you like this project, please consider giving me a ⭐.

## Discussion

### Issues Page

You can discuss or ask for help in [issues](https://github.com/macedo/web_search/issues).

## Contribute

### Bug Reports / Feature Requests
If you want to report a bug or request a new feature, feel free to open a [new issue](https://github.com/macedo/web_search/issues).
