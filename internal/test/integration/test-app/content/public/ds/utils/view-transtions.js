export const docWithViewTransitionAPI = document;
export const supportsViewTransitions = !!docWithViewTransitionAPI.startViewTransition;
export function modifyViewTransition(callback, mods) {
    if (mods.has('viewtransition') && supportsViewTransitions) {
        const cb = callback; // I hate javascript
        callback = (...args) => document.startViewTransition(() => cb(...args));
    }
    return callback;
}
