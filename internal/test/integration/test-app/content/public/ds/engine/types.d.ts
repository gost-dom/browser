import type { EffectFn, Signal } from '../vendored/preact-core';
import type { SignalsRoot } from './signals';
export type OnRemovalFn = () => void;
export declare enum PluginType {
    Attribute = 1,
    Watcher = 2,
    Action = 3
}
export interface DatastarPlugin {
    type: PluginType;
    name: string;
}
export declare enum Requirement {
    Allowed = 0,
    Must = 1,
    Denied = 2,
    Exclusive = 3
}
export interface DatastarSignalEvent {
    added: Array<string>;
    removed: Array<string>;
    updated: Array<string>;
}
export declare const DATASTAR_SIGNAL_EVENT = "datastar-signals";
export interface CustomEventMap {
    [DATASTAR_SIGNAL_EVENT]: CustomEvent<DatastarSignalEvent>;
}
export type WatcherFn<K extends keyof CustomEventMap> = (this: Document, ev: CustomEventMap[K]) => void;
declare global {
    interface Document {
        dispatchEvent<K extends keyof CustomEventMap>(ev: CustomEventMap[K]): void;
        addEventListener<K extends keyof CustomEventMap>(type: K, listener: WatcherFn<K>): void;
        removeEventListener<K extends keyof CustomEventMap>(type: K, listener: WatcherFn<K>): void;
    }
}
export interface AttributePlugin extends DatastarPlugin {
    type: PluginType.Attribute;
    onGlobalInit?: (ctx: InitContext) => void;
    onLoad: (ctx: RuntimeContext) => OnRemovalFn | void;
    keyReq?: Requirement;
    valReq?: Requirement;
    argNames?: string[];
}
export interface WatcherPlugin extends DatastarPlugin {
    type: PluginType.Watcher;
    onGlobalInit?: (ctx: InitContext) => void;
}
export type ActionPlugins = Record<string, ActionPlugin>;
export type ActionMethod = (ctx: RuntimeContext, ...args: any[]) => any;
export interface ActionPlugin extends DatastarPlugin {
    type: PluginType.Action;
    fn: ActionMethod;
}
export type GlobalInitializer = (ctx: InitContext) => void;
export type InitContext = {
    plugin: DatastarPlugin;
    signals: SignalsRoot;
    effect: (fn: EffectFn) => OnRemovalFn;
    actions: Readonly<ActionPlugins>;
    removals: Map<string, Map<number, OnRemovalFn>>;
    applyToElement: (el: HTMLorSVGElement) => void;
};
export type HTMLorSVGElement = Element & (HTMLElement | SVGElement);
export type Modifiers = Map<string, Set<string>>;
export type RuntimeContext = InitContext & {
    plugin: DatastarPlugin;
    el: HTMLorSVGElement;
    rawKey: Readonly<string>;
    key: Readonly<string>;
    value: Readonly<string>;
    mods: Modifiers;
    genRX: () => <T>(...args: any[]) => T;
    fnContent?: string;
};
export type NestedValues = {
    [key: string]: NestedValues | any;
};
export type NestedSignal = {
    [key: string]: NestedSignal | Signal<any>;
};
export type RuntimeExpressionFunction = (ctx: RuntimeContext, ...args: any[]) => any;
export type EventCallbackHandler = (...args: any[]) => void;
