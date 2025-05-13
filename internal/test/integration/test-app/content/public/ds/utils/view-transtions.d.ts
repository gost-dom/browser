import type { EventCallbackHandler, Modifiers } from "../engine/types";
export interface DocumentSupportingViewTransitionAPI {
    startViewTransition(updateCallback: () => Promise<void> | void): IViewTransition;
}
export interface IViewTransition {
    finished: Promise<void>;
    ready: Promise<void>;
    updateCallbackDone: Promise<void>;
    skipTransition(): void;
}
export declare const docWithViewTransitionAPI: DocumentSupportingViewTransitionAPI;
export declare const supportsViewTransitions = true;
export declare function modifyViewTransition(callback: EventCallbackHandler, mods: Modifiers): EventCallbackHandler;
