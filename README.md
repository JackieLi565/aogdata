# Advent Of Go Data

The purpose of this module is to allow users to request data from [Advent of Code](https://adventofcode.com/) using their session token.

## Getting Started

Make sure that you have your session token ready! If you do not know how to access your session token see [How Do I Get My Session Token?](#session-token)

Add as a dependency

```
go get github.com/JackieLi565/aogdata
```

Make a `.env` file with the following key

```env
AOC_SESSION=YOUR_SESSION_TOKEN
```

If you are posting your code publicly I highly encourage creating a `.gitignore` file.

## Session Token

Advent of Code gives each individual user a unique set of data. To access this data you will need your session token!

- Visit [Advent of Code](https://adventofcode.com/)
- Make sure that you are logged in
- Open devtools and goto the network section
- Refresh the page and open the `adventofcode.com` document
- Find the `cookie` key in the [Request Headers](https://developer.mozilla.org/en-US/docs/Glossary/Request_header)
- You now should be able to see your session token under `session`

## More to Come

[ ] - Support for pre-formate data.
[ ] - Option to use test cases.
