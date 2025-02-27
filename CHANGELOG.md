# Changelog - Gost-DOM


## [0.3.0](https://github.com/gost-dom/browser/compare/v0.2.0...v0.3.0) (2025-02-27)

This is a breaking change because of when callbacks will be executed. See the
commit messages for the event loop for details

### Features

* "Event loop" with delay and interval ([a132662](https://github.com/gost-dom/browser/commit/a132662bdf05067beb7bb41a795ab158dd8392f3))
* Change v8go dependency ([df1ef32](https://github.com/gost-dom/browser/commit/df1ef32fdb5d0499bb38f249956190ce7cbb533d))


### Bug Fixes

* ownerDocument when moving to new document ([590a5f0](https://github.com/gost-dom/browser/commit/590a5f0feadd49ad846d5b600024b962a11984e4))

## [0.2.0](https://github.com/gost-dom/browser/compare/v0.1.2...v0.2.0) (2025-02-19)

### Features

    * New url package (4ff3300 (https://github.com/gost-dom/browser/commit/4ff3300cbf5291e1f18dd9faea422291f49a13e6))
    * Support HTMLAnchorElement.String() (3e2efc7 (https://github.com/gost-dom/browser/commit/3e2efc7384c204db8aadc7805a9e6777b1b5eaec))

## [0.1.2](https://github.com/gost-dom/browser/compare/v0.1.1...v0.1.2) (2025-02-06)

### Bug Fixes

    * Fix ParseURLBase issues (a4ed7a5 (https://github.com/gost-dom/browser/commit/a4ed7a5c28c9a947a1a2dc50827c3b21dc9b46b2))

## [0.1.1](https://github.com/gost-dom/browser/compare/v0.1.0...v0.1.1) (2025-02-06)

### Bug Fixes

    * Passing an int to V8.NewValue (aec0653 (https://github.com/gost-dom/browser/commit/aec0653b949d40692cec9402cdf6bf995a6ba365))
