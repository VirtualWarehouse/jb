# jb
🏦 CLI for jobcan slack integration

## Instruction
### What you need

- An access token named `d`
- A slack workspace token
- A channel ID

1. Run `go run main.go login` or `./jb` after building a binary.
2. Open your slack workspace in your browser.
3. Open Devtools.
4. Go to `Application` -> `Cookies` -> `d`.
5. Copy the contents of `d` and paste it to your terminal.
6. Enter `JSON.parse(localStorage.localConfig_v2).teams` in the browser console.
7. Then you get a list of the team you are in. Copy the token of the team you want to use.
8. Paste the token to your terminal.
9. Copy the channel ID and paste it.

Original:
https://github.com/anoriqq/jb
