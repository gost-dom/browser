// Icon: ic:baseline-get-app
// Slug: Use a GET request to fetch data from a server using Server-Sent Events matching the Datastar SDK interface
// Description: Remember, SSE is just a regular SSE request but with the ability to send 0-inf messages to the client.
import {
  DATASTAR,
  DATASTAR_REQUEST,
  DefaultSseRetryDurationMs,
} from "../../../../engine/consts";
import { runtimeErr } from "../../../../engine/errors";
import { fetchEventSource } from "../../../../vendored/fetch-event-source";
import { dispatchSSE, ERROR, FINISHED, RETRYING, STARTED } from "../shared";
const isWrongContent = (err) => `${err}`.includes("text/event-stream");
export const sse = async (ctx, method, url, args) => {
  const { el, signals } = ctx;
  const elId = el.id;
  const {
    headers: userHeaders,
    contentType,
    includeLocal,
    selector,
    openWhenHidden,
    retryInterval,
    retryScaler,
    retryMaxWaitMs,
    retryMaxCount,
    abort,
  } = Object.assign(
    {
      headers: {},
      contentType: "json",
      includeLocal: false,
      selector: null,
      openWhenHidden: false, // will keep the request open even if the document is hidden.
      retryInterval: DefaultSseRetryDurationMs, // the retry interval in milliseconds
      retryScaler: 2, // the amount to multiply the retry interval by each time
      retryMaxWaitMs: 30_000, // the maximum retry interval in milliseconds
      retryMaxCount: 10, // the maximum number of retries before giving up
      abort: undefined,
    },
    args
  );
  const action = method.toLowerCase();
  let cleanupFn = () => {};
  try {
    dispatchSSE(STARTED, elId, {});
    if (!url?.length) {
      throw runtimeErr("SseNoUrlProvided", ctx, { action });
    }
    const initialHeaders = {};
    initialHeaders[DATASTAR_REQUEST] = true;
    // We ignore the content-type header if using form data
    // if missing the boundary will be set automatically
    if (contentType === "json") {
      initialHeaders["Content-Type"] = "application/json";
    }
    const headers = Object.assign({}, initialHeaders, userHeaders);
    const req = {
      method,
      headers,
      openWhenHidden,
      retryInterval,
      retryScaler,
      retryMaxWaitMs,
      retryMaxCount,
      signal: abort,
      onopen: async (response) => {
        if (response.status >= 400) {
          const status = response.status.toString();
          dispatchSSE(ERROR, elId, { status });
        }
      },
      onmessage: (evt) => {
        if (!evt.event.startsWith(DATASTAR)) {
          console.log("NOT STARTS WITH");
          return;
        }
        const type = evt.event;
        const argsRawLines = {};
        const lines = evt.data.split("\n");
        for (const line of lines) {
          const colonIndex = line.indexOf(" ");
          const key = line.slice(0, colonIndex);
          let argLines = argsRawLines[key];
          if (!argLines) {
            argLines = [];
            argsRawLines[key] = argLines;
          }
          const value = line.slice(colonIndex + 1);
          argLines.push(value);
        }
        const argsRaw = {};
        for (const [key, lines] of Object.entries(argsRawLines)) {
          argsRaw[key] = lines.join("\n");
        }
        dispatchSSE(type, elId, argsRaw);
      },
      onerror: (error) => {
        if (isWrongContent(error)) {
          // don't retry if the content-type is wrong
          throw runtimeErr("InvalidContentType", ctx, { url });
        }
        // do nothing and it will retry
        if (error) {
          console.error(error.message);
          dispatchSSE(RETRYING, elId, { message: error.message });
        }
      },
    };
    const urlInstance = new URL(url, window.location.origin);
    const queryParams = new URLSearchParams(urlInstance.search);
    if (contentType === "json") {
      console.log("JSON");
      const json = signals.JSON(false, !includeLocal);
      if (method === "GET") {
        queryParams.set(DATASTAR, json);
      } else {
        req.body = json;
      }
      console.log("JSON DONE");
    } else if (contentType === "form") {
      console.log("FORM");
      const formEl = selector
        ? document.querySelector(selector)
        : el.closest("form");
      if (formEl === null) {
        if (selector) {
          throw runtimeErr("SseFormNotFound", ctx, { action, selector });
        }
        throw runtimeErr("SseClosestFormNotFound", ctx, { action });
      }
      if (el !== formEl) {
        const preventDefault = (evt) => evt.preventDefault();
        formEl.addEventListener("submit", preventDefault);
        cleanupFn = () => formEl.removeEventListener("submit", preventDefault);
      }
      if (!formEl.checkValidity()) {
        formEl.reportValidity();
        cleanupFn();
        console.log("Not valid");
        return;
      }
      const formData = new FormData(formEl);
      console.log("PING");
      if (method === "GET") {
        const formParams = new URLSearchParams(formData);
        for (const [key, value] of formParams) {
          queryParams.set(key, value);
        }
      } else {
        req.body = formData;
      }
    } else {
      throw runtimeErr("SseInvalidContentType", ctx, { action, contentType });
    }
    console.log("PING: " + queryParams.toString());
    try {
      urlInstance.search = queryParams.toString();
    } catch (err) {
      console.error("err", err);
      throw err;
    }
    console.log("PING2");
    try {
      await fetchEventSource(urlInstance.toString(), elId, req);
    } catch (error) {
      if (!isWrongContent(error)) {
        throw runtimeErr("SseFetchFailed", ctx, { method, url, error });
      }
      // exit gracefully and do nothing if the content-type is wrong
      // this can happen if the client is sending a request
      // where no response is expected, and they haven't
      // set the content-type to text/event-stream
    }
  } finally {
    console.log("FINALLY");
    dispatchSSE(FINISHED, elId, {});
    cleanupFn();
  }
};
