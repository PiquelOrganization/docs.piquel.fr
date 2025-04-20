# docs.piquel.fr

What I use to serve documentation. Basically a webserver that serves Markdown as HTML.

## TODO

- Verify Github signature in webhook
- fix: webhook always fails as server restarts as soon as request is received
- Serve assets in a static dir
- Render the documentation based on requirements
- add HOME_ROUTE to make a homemage
- make reading files concurrent (every getFilesFromDir call should be goroutine)
