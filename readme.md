# Alfred Web Searches [![Workflows](https://img.shields.io/badge/More%20Workflows-🎩-purple.svg)](https://github.com/learn-anything/alfred-workflows) [![Thanks](https://img.shields.io/badge/Say%20Thanks-💗-ff69b4.svg)](https://www.patreon.com/nikitavoloboev)
> [Alfred](https://www.alfredapp.com/) workflow to search through any website on the web (easily extendable list)

<img src="https://i.imgur.com/FELOBBm.png" width="500" alt="img">

This workflow similar to [DuckDuckGo Bangs](https://duckduckgo.com/bang?) allows you to fuzzy search through a selection of websites that you can add and contribute to in [here](https://github.com/nikitavoloboev/alfred-web-searches/edit/master/workflow/websites.csv).

You simply search for the website you want to scope your search too, press enter and make your search.

Each search has a prefix to signify the theme or scope of the search. For example results prefixed with `r: ` will search through subreddits. Here is the full list of prefixes and their descriptions.

|  Prefix |  Decription |
|---|---|
|  r: | Search subreddits on Reddit  |
|  d: | Search documentation websites |
|  g: | Search GitHub |
|  s: | Search stack exchange sites |
| f:  | Search forums |
| t:  | Search Google Translate |
| w:  | Search ordinary websites (none of the above) |

Each prefix also has a corresponding icon for visual feedback of prefix meaning.

And since the workflow allows for fuzzy searching, you can even make searches like that:

<img src="https://i.imgur.com/Rf4N6jK.png" width="500" alt="img">

And they will match.

You can also attach a hotkey that will take the selected text and then will let you search for that text on any of the websites.

## Install
Download the workflow from [GitHub releases](../../releases/latest).

## Contributing
[Suggestions](../../issues/) and pull requests are highly encouraged!

## Developing
If you want to add features and things to the workflow. It is best to use [this Alfred CLI tool](https://godoc.org/github.com/jason0x43/go-alfred/alfred) which you can install by running:

`go install github.com/jason0x43/go-alfred/alfred`

You can then clone this repo and run `alfred link` inside it. This will make a symbolic link of the [`workflow`](workflow) directory.

You can then make changes to the code and after run `alfred build` to build the go binary to `workflow` directory. Which you can then use from inside Alfred [script filters](https://www.alfredapp.com/help/workflows/inputs/script-filter/).

## Credits
The workflow uses [AwGo](https://github.com/deanishe/awgo) library for all the Alfred related things.

## Thank you 💜
You can support what I do on [Patreon](https://www.patreon.com/nikitavoloboev) or look into [other projects](https://nikitavoloboev.xyz/projects) I shared.

## License
MIT © [Nikita Voloboev](https://www.nikitavoloboev.xyz)