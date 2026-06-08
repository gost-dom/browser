package fetch

import "time"

// SimulatedDelay defines the simulated time that must pass from making a
// request until the response is processed in the event loop.
//
// Deprecated: This is HIGHLY experimental, helping carve out the mechanism to
// control time. Once that is in place, controlling delay will be
// context-dependent.
var SimulatedDelay time.Duration
