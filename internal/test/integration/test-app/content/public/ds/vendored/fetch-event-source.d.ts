/**
 * Represents a message sent in an event stream
 * https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events/Using_server-sent_events#Event_stream_format
 */
export interface EventSourceMessage {
    /** The event ID to set the EventSource object's last event ID value. */
    id: string;
    /** A string identifying the type of event described. */
    event: string;
    /** The event data */
    data: string;
    /** The reconnection interval (in milliseconds) to wait before retrying the connection */
    retry?: number;
}
/**
 * Converts a ReadableStream into a callback pattern.
 * @param stream The input ReadableStream.
 * @param onChunk A function that will be called on each new byte chunk in the stream.
 * @returns {Promise<void>} A promise that will be resolved when the stream closes.
 */
export declare function getBytes(stream: ReadableStream<Uint8Array>, onChunk: (arr: Uint8Array) => void): Promise<void>;
/**
 * Parses arbitary byte chunks into EventSource line buffers.
 * Each line should be of the format "field: value" and ends with \r, \n, or \r\n.
 * @param onLine A function that will be called on each new EventSource line.
 * @returns A function that should be called for each incoming byte chunk.
 */
export declare function getLines(onLine: (line: Uint8Array, fieldLength: number) => void): (arr: Uint8Array) => void;
/**
 * Parses line buffers into EventSourceMessages.
 * @param onId A function that will be called on each `id` field.
 * @param onRetry A function that will be called on each `retry` field.
 * @param onMessage A function that will be called on each message.
 * @returns A function that should be called for each incoming line buffer.
 */
export declare function getMessages(onId: (id: string) => void, onRetry: (retry: number) => void, onMessage?: (msg: EventSourceMessage) => void): (line: Uint8Array, fieldLength: number) => void;
export declare const EventStreamContentType = "text/event-stream";
export interface FetchEventSourceInit extends RequestInit {
    /**
     * The request headers. FetchEventSource only supports the Record<string,string> format.
     */
    headers?: Record<string, string>;
    /**
     * Called when a response is received. Use this to validate that the response
     * actually matches what you expect (and throw if it doesn't.) If not provided,
     * will default to a basic validation to ensure the content-type is text/event-stream.
     */
    onopen?: (response: Response) => Promise<void>;
    /**
     * Called when a message is received. NOTE: Unlike the default browser
     * EventSource.onmessage, this callback is called for _all_ events,
     * even ones with a custom `event` field.
     */
    onmessage?: (ev: EventSourceMessage) => void;
    /**
     * Called when a response finishes. If you don't expect the server to kill
     * the connection, you can throw an exception here and retry using onerror.
     */
    onclose?: () => void;
    /**
     * Called when there is any error making the request / processing messages /
     * handling callbacks etc. Use this to control the retry strategy: if the
     * error is fatal, rethrow the error inside the callback to stop the entire
     * operation. Otherwise, you can return an interval (in milliseconds) after
     * which the request will automatically retry (with the last-event-id).
     * If this callback is not specified, or it returns undefined, fetchEventSource
     * will treat every error as retriable and will try again after 1 second.
     */
    onerror?: (err: any) => number | null | undefined | void;
    /**
     * If true, will keep the request open even if the document is hidden.
     * By default, fetchEventSource will close the request and reopen it
     * automatically when the document becomes visible again.
     */
    openWhenHidden?: boolean;
    /** The Fetch function to use. Defaults to window.fetch */
    fetch?: typeof fetch;
    /** The retry interval in milliseconds. Defaults to 1_000 */
    retryInterval?: number;
    /** The scaler for the retry interval. Defaults to 2 */
    retryScaler?: number;
    /** The maximum retry interval in milliseconds. Defaults to 30_000 */
    retryMaxWaitMs?: number;
    /** The maximum number of retries before giving up. Defaults to 10 */
    retryMaxCount?: number;
}
export declare function fetchEventSource(input: RequestInfo, elId: string, { signal: inputSignal, headers: inputHeaders, onopen: inputOnOpen, onmessage, onclose, onerror, openWhenHidden, fetch: inputFetch, retryInterval, retryScaler, retryMaxWaitMs, retryMaxCount, ...rest }: FetchEventSourceInit): Promise<void>;
