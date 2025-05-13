export declare const DATASTAR_SSE_EVENT = "datastar-sse";
export declare const STARTED = "started";
export declare const FINISHED = "finished";
export declare const ERROR = "error";
export declare const RETRYING = "retrying";
export declare const RETRIES_FAILED = "retries-failed";
export interface DatastarSSEEvent {
    type: string;
    elId: string;
    argsRaw: Record<string, string>;
}
export interface CustomEventMap {
    [DATASTAR_SSE_EVENT]: CustomEvent<DatastarSSEEvent>;
}
export type WatcherFn<K extends keyof CustomEventMap> = (this: Document, ev: CustomEventMap[K]) => void;
declare global {
    interface Document {
        addEventListener<K extends keyof CustomEventMap>(type: K, listener: WatcherFn<K>): void;
        removeEventListener<K extends keyof CustomEventMap>(type: K, listener: WatcherFn<K>): void;
        dispatchEvent<K extends keyof CustomEventMap>(ev: CustomEventMap[K]): void;
    }
}
export declare function datastarSSEEventWatcher(eventType: string, fn: (argsRaw: Record<string, string>) => void): void;
export declare function dispatchSSE(type: string, elId: string, argsRaw: Record<string, string>): void;
