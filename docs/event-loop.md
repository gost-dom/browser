# Timeouts and the Event Loop

> [!INFO]
> 
> In Gost-DOM, the event loop is not "running" in the background. Instead,
> client code must explicitly tell Gost-DOM to run enqueued tasks.

In a browser, the event loop is executed in the main thread of the browser; the
thread that also handles user interaction, and updates the DOM accordingly.
JavaScript code will not run concurrently with DOM updates.

When using Gost-DOM, you might manipulate the DOM. Either directly, or using the
packages to [simulate user input]. As a consequence, delayed functions, e.g.,
`setTimeout` or `setInterval` callbacks don't execute as time pass. You have to
explicitly advance time.

> [!CAUTION]
>
> All interaction with a window must happen in the same goroutine. 

## Gost-DOM runs in simulated time.

Gost-DOM uses simulated time, allowing tests to fast-forward time. So verifying
throttling behaviour with a 200ms delay does not need to wait 200 milliseconds.
Client code advance time by calling these two functions.

- `Window.Clock().Advance(time.Duration)` advances time for a certain amount of
  time, running all timeout and interval callbacks that should run in that
  period.
- `Window.Clock().RunAll()` will run until all `setTimeout` callbacks are
  called. If a callback adds a new callback, they will be executed too.

> [!WARNING]
>
> Both `Advance()` and `RunAll()` will try to prevent an infinite loop caused by
> new callbacks continuously being registered. It expects that the number of
> registered callbacks decrease over time; and will panic if they don't. This
> will be triggered if a `setTimeout` callback continuously registers new
> timeouts. Likewise, `RunAll()` will effectively panic if there are any
> intervals registered (that are not cleaned up).^[1] ^[2]

> [!NOTE]
>
> Simulated time only affects execution of `setTimeout`/`setInterval` callbacks.
> The return value of `new Date()` will return the system clock.

[simulate user input]: ./simulating-user-input.md

## Waiting for async operations

If Gost-DOM client code calls, `Click()` on a button that triggers a `fetch`
call, the data will be fetched in a new goroutine.

To wait for the code to complete, use `Window.Clock().ProcessEvents()`. This
will wait for all promises to have settled. Alternately,
`Window.Clock().ProcessEventsWhile()` will wait as long as a specific condition
is true.

> [!WARNING]
>
> `ProcessEvents()`/`ProcessEventsWhile()` are new functions added as the
> behaviour was necessary to write predictable test code when using `fetch`. The
> entire scenario exposed a problem not properly anticipated originally, and as
> a consequence, the entire clock/event loop may require some changes.


[^1]: There are tuning parameters for this behaviour; but they are not exposed.
    Please submit an issue if this you experience issues with this.

[^2]: A recursive `setInterval` call with zero delay will trigger this
    behaviour, but shouldn't for the `Advance()` case: 
    https://github.com/gost-dom/browser/issues/45)

