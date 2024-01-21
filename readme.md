# Alfred Web Searches

> [Alfred](https://www.alfredapp.com/) workflow to search through any website (easily extendable list)

<img src="https://i.imgur.com/IxhwjuN.png" width="500" alt="img">

This workflow similar to [DuckDuckGo Bangs](https://duckduckgo.com/bang?) allows you to fuzzy search through a selection of websites that you can [add and contribute to](workflow/websites.csv).

You simply search for the website you want to scope your search too, press enter and make your search.

Some results have a prefix to signify the scope of the search. For example results prefixed with `r:` will search through Reddit subreddits. Here is the full list of prefixes and their descriptions.

| Prefix | Description                   |
| ------ | ----------------------------- |
| r:     | Search subreddits on Reddit   |
| d:     | Search documentation websites |
| g:     | Search GitHub                 |
| s:     | Search stack exchange sites   |
| f:     | Search forums                 |
| t:     | Search Google Translate       |

Each prefix also has a corresponding icon for visual feedback of prefix meaning.

And since the workflow allows for fuzzy searching, you can even make searches like that:

<img src="https://i.imgur.com/eekNFrr.png" width="500" alt="img">

And they will match.

You can also attach a hotkey that will take the selected text and then will let you search for that text on any of the websites.

## Install

Download workflow from [GitHub releases](../../releases/latest).

See [here](https://github.com/deanishe/awgo/wiki/Catalina) for instructions on fixing permissions in macOS refusing to run Go binary.

## Setup

The workflow is written in [Go](https://golang.org/) and uses [AwGo](https://github.com/deanishe/awgo) library for all Alfred related things.

It uses [modd](https://github.com/cortesi/modd) and [Alfred command](https://godoc.org/github.com/jason0x43/go-alfred/alfred) to ease its development.

## Run

1. Run `alfred link` (makes symbolic link of [`workflow`](workflow) directory)
2. Run `modd` (starts a process that automatically builds the workflow with `alfred build` on any changes you make to `.go` files, this builds and places a binary inside [`workflow`](workflow) directory.)
3. Make changes to code or modify Alfred objects to do what you want! Open debugger in Alfred or run the workflow with `workflow:log` passed in as argument to see the logs Alfred produces.

![](https://i.imgur.com/FFYOecx.png)

## Contribute

Always open to useful ideas or fixes in form of issues or PRs.

Can [open new issue](../../issues/new/choose) (search [existing issues](../../issues) first) or [start discussion](../../discussions).

It's okay to submit draft PR as you can get help along the way to make it merge ready.

Join [Discord](https://discord.com/invite/TVafwaD23d) for more indepth discussions on this repo and [others](https://github.com/nikitavoloboev#src).

### 🖤

[Support on GitHub](https://github.com/sponsors/nikitavoloboev) or look into [other projects](https://nikiv.dev/projects).

[![Discord](https://img.shields.io/badge/Discord-100000?style=flat&logo=discord&logoColor=white&labelColor=black&color=black)](https://discord.com/invite/TVafwaD23d) [![X](https://img.shields.io/badge/nikitavoloboev-100000?logo=X&color=black)](https://twitter.com/nikitavoloboev) [![nikiv.dev](https://img.shields.io/badge/nikiv.dev-black)](https://nikiv.dev)
