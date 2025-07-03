# Contributor's Guide

Aweseome. Apart from actually help writing new features, here are a few other
ways to contribute:

- **Provide feedback**. This is super important. The feedback from users help
  set the direction, and prioritise features.
- Providing test projects to check for compatibility with existing frameworks
  and libraries, with the potential to serve as a foundation to extend the
  existing test suite.
- Help with documentation. Maybe you had a problem that could have been avoided
  by better documentation.
- Help maintain a website. Currently online at https://gostdom.net - but it's a
  very minimal Jekyll site. I hope this will over time be a good source of
  information, and documentation about using the project

No matter how you would like to contribute, I encourage you to [join the
discussions](https://github.com/orgs/gost-dom/discussions)

## Contributing to the code

There are a few issues marked as "good first issue", as they seem isolated
enough to be easy to start on. But they aren't the highest priority.

Otherwise, here are some guidelines:

- Let yourself known before starting working on a feature to avoid conflicts
- There is not a "correct" prioritization. Without user feedback, prioritization
  is based on what I experience, and what I believe would make most sense
- Browse through the [Developer docs](./docs/Developer-docs.md), but don't be
  afraid to ask questions about the implementation.
- If there's something you find to be weird or unidiomatic, chances are that I'd agree. Any suggestions 
- **Read the docs**. This project implements specific standards. Both MDN and the
  Living Standard are sources of information.
    - [MDN Developer](https://developer.mozilla.org/en-US/docs/Web)
    - [WHATWG Living Standard](https://html.spec.whatwg.org/multipage)

Regarding conflicts. The code base is being heavily refactored, as [new
insights](https://github.com/orgs/gost-dom/discussions/50) leads to better
implementations. If this leads to breaking changes (from a client's point of
view), I'd normally announce this ahead of publishing; but refactors that don't
break the API aren't.

## Provide examples for other frameworks.

In order to check how Gost-DOM works with different frameworks, example projects
are necessary. You can help by creating new projects, or extend the existing
ones with more use cases, to help find compatibility issues.

https://github.com/gost-dom/browser/issues/124

### Test project "requirement"

In order to easily consume the project from Gost-DOM, the project must:

- Be `git clone`able. E.g., a public github project is great.
- The root `http.Handler` should be exported, either as a variable, or a
  function to construct an instance.
- The root handler must serve all JavaScript code, I recommend go's
  [embed](https://pkg.go.dev/embed) package.
- JavaScript code **must be bundled** if you use ESM (preferably committed to git,
  but I can manage an npm build)
- Optional: Have a `main` module, to make it easy to test in a browser for
  comparison. I can manage creating it, but it'll save me the work.
- Optional: Create a `readme` describing the project, particularly if different
  actions depend on different browser APIs.

Regarding ESM bundles, Gost-DOM will ignore `type="module` on `<script>` tags,
and will treat it as a class script source. So as long as the script doesn't
`import` other modules, it should be fine.

#### Things to consider:

- Can the framework be grouped into layers of functionality?
- Can you create separate routes to check core behaviour separately from
  extended/advanced behaviour.
  - If the framework supports plugins, can we test the core functionality without
    any plugins, and have different pages for the different plugin behaviour.

### Example

As an example, the project's own _extremely limited_ [HTMX
test](https://github.com/gost-dom/browser/tree/main/internal/test/integration)

- `test-app` exports `CreateServer()` that returns an http server.
- The server uses `embed` to serve `htmx.js` from `contents/` directory.
- `test-app-main` has a `main` to launch the app for comparison in a real browser.

<!--
### Optional: Create a PR to the community-examples

You can optionally 

1. Fork the [community-examples](https://github.com/gost-dom/community-examples)
   project.
2. Create the project in a subdirectory of the project. The directory must be
   self-contained following the requirements listed above
3. Create a PR back to the community-examples

### Optional: Include a permissive license

If you provide a permissive license, preferable one without _any_
restrictions, even just the requirement to keep the license file, I could easily
integrate your example into the main project test suite, if it would make sense.

I would always make sure contributors are mentioned (unless you prefer not to).
-->

