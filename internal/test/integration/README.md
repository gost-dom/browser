# Test app

This folder contains a simple test app, including HTMX, that is being used in
the test suite for gost-dom.

```sh
test-app/ # The actual HTML server implementation as an http.Handler
test-app-main/ # An app with a real server for inspecting in a browser.
```

## Simple dev mode

The `dev` script will start the test server, and restart when a source file is
saved. The script uses `nodemon` from the NPM ecosystem. You can install it if
you have node.js and NPM installed using:

```
> npm i -g nodemon
```
