import { DATASTAR } from '../../../engine/consts';
export const DATASTAR_SSE_EVENT = `${DATASTAR}-sse`;
export const STARTED = 'started';
export const FINISHED = 'finished';
export const ERROR = 'error';
export const RETRYING = 'retrying';
export const RETRIES_FAILED = 'retries-failed';
export function datastarSSEEventWatcher(eventType, fn) {
    document.addEventListener(DATASTAR_SSE_EVENT, (event) => {
        if (event.detail.type !== eventType)
            return;
        const { argsRaw } = event.detail;
        fn(argsRaw);
    });
}
export function dispatchSSE(type, elId, argsRaw) {
    document.dispatchEvent(new CustomEvent(DATASTAR_SSE_EVENT, {
        detail: { type, elId, argsRaw },
    }));
}
