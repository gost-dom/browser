# Why Gost-DOM?

Ghost-DOM was born out of the philosophy.

- TDD is the most efficient way to write _most_ code.
- TDD of UI behaviour tests should be "gray box tests" written in the back-end language

## Why does TDD make you faster?

When done right, TDD is efficient because it favours fast feedback loops. The
fast feedback loops makes it a joy to continuously work towards the goal in small
increments; small increments are easier to handle. And when a change didn't
produce the intended effect, it's significantly easier to determine why. 

TDD only works when a small snippet of code can provide _fast and relevant_
feedback. But for the vast majority of the code, a small test provides the best
means of fast feedback.

The faster the feedback cycle, the more efficient the development cycle. I've
worked with sub-100 millisecond feedback cycles, running tests so frequently
that when they broke, I could just undo the the last change and retry;
this was much faster than trying to figure out what was wrong.

> [!Note]
>
> _Some_ types of application may not lend themself well to TDD. For example,
> game development, video and audio codecs, AI training algorithms are some
> areas I imaging might not do so well under TDD, and you may need to search for
> other means of feedback. 
>
> No matter what you work with, to improve efficiency, you should optimise for
> fast feedback, e.g. when working the the visual design, running a browser with
> live-reload provides fast visual (relevant) feedback.
>
> TDD just happen to work best for _most of the code_ in _most_ applications.

### TDD and web applications

When the implementation of behaviour in a web application requires JavaScript, a
browser-like environment with a JavaScript runtime is necessary to test.
JavaScript projects have had this ability without needing a real browser through
projects like:

- [jsdom](https://github.com/jsdom/jsdom)
- [happy-dom](https://github.com/capricorn86/happy-dom/wiki/)
- [Zombie](https://github.com/assaf/zombie) (no longer maintained)
- [PhantomJS](https://github.com/ariya/phantomjs) (no longer maintained)

JSDom and Happy-dom are typically used to unit test front end components in
Single-page applications, i.e., not the web page as a whole, but testing
individual components in isolation, in a browser-like environment. Zonbie and
PhantomJS were headless browsers, i.e. JavaScript libraries simulating the
behaviour of a browser.

However, applications with server-side rendering requires a "browser" to request
an HTML page, parse the HTML, and execute the client-side script in a proper
environment where it's internal DOM has been exposed to JavaScript. 

Writing meaningful tests of individual parts is not feasible; For hypermedia
frameworks like HTMX and Datastar, the application behaviour is the choreography
of rendered HTML, the attributes on specific elements, the endpoints that the
framework calls, the request and response headers and bodies.

Generally, the only option is to use a real browser, but developers often
struggle with:

- Erratic tests cases (they sometimes fail, but pass when rerunning)
- Fragile test cases (they break when code changes, but the system works)
- Slow tests caused by the overhead, due to launching external programs,
  inter-process communication, and network traffic.
- Slow test caused by testing functionslity requiring time to pass, e.g.,
  throttled or debounced behaviour.

If we accept the premise that the effectiveness of TDD is directly affected by
the speed of the test feedback cycle, then the overhead of browser automation
negatively impacts the benefit of TDD; possibly even to the point where it's
slowing you down instead of speeding you up.

#### A headless browser to the rescue.

A _headless browser_ written in the same language as the tests cases mitigates
the problems faced by developers when using browser automation.

- By running as a library, you elliminate the overhead of launching a new
  process, as well as inter-process communication.
- Running in-process allows the embedder (the test case) control over internals,
  such as the passing of time, avoiding delays in test cases for throttled
  behaviour, as tests may fast forward simulated time.
- Test cases can much better synchronise with the browser itself, avoiding
  fragile tests.

> [!Note]
>
> Running a real browser, like chrome, in headless mode _does not_ make it a
> headless browser; it's still a browser running in headless mode, and has the
> problems inherent to browser automation.

<!--
When it comes to web user interfaces, a real browser is difficult to avoid
unless the application is void of any JavaScript code. But with a significant
amount of behaviour implemented on the client, testing needs a JavaScript
runtime in a browser-like environment. For many projects, only a real browser
provides this environment.


In addition to the overhead inherent to the technical nature, some test are also
inherently slow because the behaviour under test _depends the passing of time_.
For example, auto-complete search fields would often employ some kind of
throttling or debounce behaviour to reduce unnecessary traffic and load,
typically waiting around 2-500 milliseconds.

As a result, in these scenarios, tests are often written _after_ the system has
been manually verified to work as intended.

> [!Note]
>
> While you may not be able to solve slow tests, you _should always_ try to
> solve erratic and fragile tests. In fact, making them less fragile may make
> them slower; but if you want to avoid bugs creeping into existing
> functionaily; you need a reliable test suite.

-->

## Testing the UI should be "Gray Box Tests"

A test case will typically have 3 phases:

1. Setting up an initial state of the system
2. Exercising behaviour under test through some external stimuli
3. Verifying the end state of the system

When testing the behaviour of the user interface, the external stimuli will
exercise the user interface, simulating click or keyboard input. Part of the
verification relates to the response communicated to the user, and here the test
must interact with the user interface.

But setting up initial entities in the system, or verifying the state of
persisted entities, this should reuse the existing code you already have in 

Some projects use only black box testing when testing the user interface. That
[that is a flawed strategy](./1.1.1 Black box testing is a flawed strategy.md),
and might even not be possible.

### Mocking becomes an option

UI tests are typically full integration tests, i.e. when exercising the
UI, all layers of the system are exercised, persisting state in a database.

If you choose to isolate the test to verify the behaviour of the UI alone, a
viable approach is mocking the actual "use case".

This ties the test more closely to the responsibilities of the UI layer in
isolation. E.g., when following the "clean architecture", the responsibilities
of the user interface layer are primarily conversion of user input to use case
calls, and conversion of the outcomes of the use cases to proper user responses.
Related responsibilities include authentication and checking input validity.[^1]

The test case becomes just that, verifying that given some user input in the
browser, the use case is called with verifyable arguments, error messages are
verified for invalid inputs, as well as the feedback to the user is verified for
different pre-programmed mocked responses from the use case call.

> [!tip]
>
> I do not intend to imply that mocking is preferable to integration tests;
> neither that the "clean architecture" is superior in any way; it's just meant
> as an example. All approaches have pros and cons.
>
> But the ability to replace some dependencies is extremely powerful.
> Particularly when systems start to incorporate messages-driven aspects, or
> other delayed processing the runtime behaviour becomes unpredictable, and
> stable tests are more difficult to maintain.
>
> Replaceing components to improve predictability can significantly improve the
> stability of the test suite.

> [!caution]
>
> Not all technologies facilitate mocking equally well when testing the UI.
> Gost-DOM facilitates mocking as well as any other test in the system, i.e. -
> it only depends on how well your application architecture supports it; whereas
> .NET makes it _very_ difficult; although [Nancy](https://nancyfx.org/)
> _appears to_ support this much better.[^2]

### TDD'ing the UI

So to drive development of a new behaviour in the a test, it will typically have
these steps.

1. Set up the initial state, inserting domain objects in an empty database; or
   mocking a pre-programmed response.
2. Create a browser instance, possibly logging in if the feature requires an
   authenticated user.
3. Navigate to the page with the behaviour to verify
4. Interact with the page.
5. Verify that the proper response was communicated to the user
6. Verify data was updated correctly/mocks were called with the correct
   arguments.

When your test code can reuse production code for setting up initial state, and
verifying end state; and when the UI interaction runs in the same thread as the
test case, the test is fast and reliable, and facilitates a TDD loop.

This facilitates refactoring, as well as choosing the easy and simple solution
up front, and slowly refactoring to patterns as you identify the, extracting
common behaviour in middlewares, replacing one 3rd party CSRF library with a
different one, and have the test case verify that protection still works.

Over time, a large part of the test code will be extracted to test helpers.
E.g., I would explicitly interact with the login page when testing login
behaviour; but all other tests that merely repends on a logged in user will use
a test helper to create a browser in a logged in state.


## Conclusion

In this document, I set the premises:

- TDD is the most efficient way to write _majority of the code_ in most
  applications.
- Web user interfaces has historically been an exception to this rule.
- This is caused by the overhead and erraticness of browser automation.
- Because of this, web UI is typically developed using a more traditional
  production-code-first approach; _possibly_ adding a few tests afterwards.
- Black box testring the user interface is a flawed strategy, tests do not
  reflect business rules, and may rely on behaviour not exposed through the UI.

Gost-DOM aims to support a TDD process by elliminating the problems of browser
automation, achiving a fast feedback loop for a reliable test suite that is a
joy to work with. This is achieved by:

- elliminating the overhead involved in launching a browser and chatty
  inter-process communication.
- allow bypassing the overhead of the TCP stack when communicating with the web
  application under test, further reducing overhead.
- simulate time, allowing fast-forwarding of time, testing throttling and
  debounce behaviour without the delay.
- providing completely predictable runtime execution, elliminating erraticness
  of tests.
- writing the test code in the same language as production code, allowing back
  door manipulation.

This makes it fast and predictable, facilitating a TDD process to drive the
implementation of behaviour of a web user interface; allowing developers to work
in small increments.

Furthermore, by allowing bypassing the TCP stack, Gost-DOM facilitates mocking
of dependencies - the HTTP server is just a code dependency to the test case
like any other test case. This permits tests to run in full isolation, and as a
consequence, when mocking database layers or other shared state, the tests cases
can run in parallel, without any hassle of managing database contexts, TCP
ports, etc.

Gost is specifically written for hypermedia applications in mind. Other
architectures may have better options. Read more here: [When not to use Gost-DOM](bad-cases-for-gost-dom.md)

### TDD vs. Test Strategy

Here, I have been describing TDD as a means to increse effeciency during
development; not as the test strategy for a project. A side effect of TDD is
a test suite performing verification, and facilitating safe refactoring.

Each test that the TDD process leaves behind typically focus on an isolated
behaviour of the system; be it the business rules in the domain layer, how the
user interface present validation errors to the user, or how the database layer
detects conflicting updates.

I wouldn't trust a system that did not have tests verifying the behaviour of the
system as a whole, going through a complete "flow"; E.g., for a web shop, a flow
from login to checkout.

In my experience, a good TDD process leaves little to add, requiring possibly
only a handful of such tests written after code was developed. But they are very
worth having, preventing an issue such as users unable to log in because you
accidentally changed the login page to only be visible for authenticated
users.[^3]

And the strategy for those tests, you are free to choose. Having a suite of
black box tests that do not run during a normal TDD cycle could be a perfectly
reasonable strategy.


[^1]: Input is valid if input conforms to the specified format; e.g., a numeric
    input does not contain numbers, a date input can be parsed as a date. This
    is not to be confused with whether the input is _acceptable_; e.g. the
    return date must be later than the outbound date. It's typically in the
    domain layer you would check for acceptability. Testing error cases would
    involve scenarios that don't call the domain layer because the input is not
    valid, as well as those that do because input is valid, but not acceptable.

[^2]: This is based on my last .NET project. If this has improved, please let me
    know with an example, and I'll correct. I do not have experience with Nancy;
    my note about Nancy is based on information on the project's web site.

[^3]: Yes, this is actually an example of a bug introduced while refactoring
    authentication in the front end. It wasn't caught by any of the unit tests -
    only an integration test that had been added after code was written.
