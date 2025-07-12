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
> No matter what you work with, to improve efficiency, you should optimise the
> feedback loop, e.g. when working the the visual design, running a browser with
> live-reload provides fast visual (relevant) feedback.
>
> TDD just happen to work really well for _most of the code_ in _most_ applications.

### TDD and web applications

When it comes to web user interfaces, a real browser is difficult to avoid
unless the application is void of any JavaScript code. But with a significant
amount of behaviour implemented on the client, testing needs a JavaScript
runtime in a browser-like environment. For many projects, only a real browser
provides this environment.

Using a real browser is not without problems, and developers often struggle
with:

- Erratic tests cases (they sometimes fail, but pass when rerunning)
- Fragile test cases (they break when code changes, but the system works)
- Slow tests, and initial startup overhead, due to launching external programs,
  inter-process communication, and network traffic.

If we accept the premise that the effectiveness of TDD is directly affected by
the speed of the test feedback cycle, then the overhead of browser automation
negatively impacts the benefit of TDD; possibly even to the point where it's
slowing you down instead of speeding you up.

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

### The "headless browser"

A headless browser is a piece of software that in many ways acts as a browser,
but doesn't render a user interface. Headless browsers are often intended for
testing. The headless browser offers APIs allowing test code to simulate user
interaction with the application.

As the browser needs a JavaScript runtime to be of any use, the libraries that
simulate a browser environment are generally written in JavaScript.

- [happy-dom](https://github.com/capricorn86/happy-dom/wiki/)
- [jsdom](https://github.com/jsdom/jsdom)
- [zombie](https://github.com/assaf/zombie) (no longer maintained)
- [PhantomJS](https://github.com/ariya/phantomjs) (no longer maintained)

Happy-dom and jsdom are tools used extensively to "unit test" front-end code for
single-page applications. I.e., they don't act as full browsers, but provide
enough implementation of the DOM, to allow the developer to test their
application. These types of test don't "open a web page", but renders and tests
a single component in isolation.

As these are all JavaScript libraries, these tools also allows you to "hack the
runtime", for example, using a tool like [lolex] to simulate time, allowing
fast-forwarding of time when verifying throttled or debounced behaviour
permitting instantaneous feedback.

[lolex]: https://www.npmjs.com/package/lolex

> [!Note]
>
> Running a real browser, like chrome, in headless mode _does not_ make it a
> headless browser; it's still a browser running in headless mode, and has the
> problems inherent to browser automation.
>
> This is still an improvement compared to early browser automation methods
> where browsers did not include a headless mode, requiring the desktop
> environment on the build server (though unix-like systems could create a
> virtual desktop environment using a tool like [xvfb])

[xvfb]: https://www.x.org/archive/X11R7.7/doc/man/man1/Xvfb.1.xhtml

## TDD of UI behaviour tests should be "gray box tests" written in the back-end language

Many projects employ a black box testing strategy when it comes to testing the
UI. This is a flawed approach, but before explaining why, let's get the
terminology clear.

### Black box vs. white box testing

A _black box test_ is a test exercising the system from the outside with no
knowledge of the internals of the system (we cannot see inside, it's black).
Testing the system entirely through the UI is an example of a black-box test.

Obviously, a white box test is the opposite, you test the system with intimate
knowledge of its internals. Unit testing is an example of white box testing. TDD
typically generates white box tests.

Gray box testing uses knowledge with the internals, but stimulates the system
from the outside. A gray box test might use the internals of the syste setup the
system in the initial state and/or verify the end state (also known as back-door
manipulation), but exercise the user interface relating to the behaviour under test.

### Black box testing is a flawed strategy

A test case will typically have 3 cases

1. Setting up an initial state
2. Exercising the system, providing some external stimuli
3. Verifying the end state of the system

When testing the behaviour of the user interface, the external stimuli will
exercise the user interface, and for the part that the verification that relates
to the response communicated to the user, verification also need to interact
with the user interface. 

But setting up initial entities in the system, or verifying the state of
persisted entities, this should just use back-door manipulation. Performing
these task through the user interface is not just a bad idea, it can be outright
impossible.

#### Black box testing leads to fragile tests

User interfaces may change. When the user interface for feature X is changed,
only the tests for feature X should be updated.

Black box testing the user interface will lead to fragile tests, as changes to
a part of the user interface affect the outcome of tests for other features.

> [!INFO]
>
> A _fragile test_, as a test anti-pattern. It is a test the often fails when
> you make code changes despite the feature still works as intended. Fragile
> tests are usually caused by test cases too closely coupled to the
> implementation.

#### Black box testing the behaviour of the UI is not possible

As the application grows, so does the number of use cases, but not all use cases
are initiated from the user interface.

Some use cases can be triggered by the passing of time, e.g., when an invoice is
past its due date, send a reminder to the customer. Some may be triggered by
events from other systems; When a user account was deleted in the identity
provider, mark all comments created by the user as deleted. And some may be
triggered by a system administrator using a CLI, like creating a new tenant in a
multi-tenant SaaS start-up, that doesn't have an automated sign-up process yet.

It's simply impossible to black box test all behaviours of the system by
exercising the user interface alone.

#### Black box tests operate at the wrong level of abstration

Imagine a user story, "Close account", and there's a rule, "A member cannot
close an account that was created by an administrator". The behaviour in the
user interface could be that for such an account, the normal button to delete an
account is replaced by a message saying "please contact ...".

To setup the initial state, you need an account that was created by an
administrator, and this should be expressed as concisely as possible in the test
case. For example, it could be:

```Go
creator := User{ Kind: UserKindAdministrator }
account := Account{ CreatedBy: creator }
userRepo.Save(creator)
accountRepo.Save(account)
currentUser := User{ Kind: UserKindMember }
```

A black box does not clearly expresses the business rules: an account created by
an administrator. The black box test describes the process to create such an
account. That process is not the focus of _this_ test; that process is tested
elsewhere.

#### Mocking becomes an option

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
only a handful of such tests written after code was developed.

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
