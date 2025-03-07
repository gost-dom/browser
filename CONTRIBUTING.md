# Contributor's Guide

Aweseome. Apart from actually help writing new features, here are a few other
ways to contribute:

- **Provide feedback**. This is super important. The feedback from users help
  set the direction, and prioritise features.
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
