export type ConfigHead = {
    style?: "morph" | "append" | "none" | "merge" | undefined;
    block?: boolean | undefined;
    ignore?: boolean | undefined;
    shouldPreserve?: ((arg0: Element) => boolean) | undefined;
    shouldReAppend?: ((arg0: Element) => boolean) | undefined;
    shouldRemove?: ((arg0: Element) => boolean) | undefined;
    afterHeadMorphed?: ((arg0: Element, arg1: {
        added: Node[];
        kept: Element[];
        removed: Element[];
    }) => void) | undefined;
};
export type ConfigCallbacks = {
    beforeNodeAdded?: ((arg0: Node) => boolean) | undefined;
    afterNodeAdded?: ((arg0: Node) => void) | undefined;
    beforeNodeMorphed?: ((arg0: Element, arg1: Node) => boolean) | undefined;
    afterNodeMorphed?: ((arg0: Element, arg1: Node) => void) | undefined;
    beforeNodeRemoved?: ((arg0: Element) => boolean) | undefined;
    afterNodeRemoved?: ((arg0: Element) => void) | undefined;
    beforeAttributeUpdated?: ((arg0: string, arg1: Element, arg2: "update" | "remove") => boolean) | undefined;
};
export type Config = {
    morphStyle?: "outerHTML" | "innerHTML" | undefined;
    ignoreActive?: boolean | undefined;
    ignoreActiveValue?: boolean | undefined;
    restoreFocus?: boolean | undefined;
    callbacks?: ConfigCallbacks | undefined;
    head?: ConfigHead | undefined;
};
export type NoOp = Function;
export type ConfigHeadInternal = {
    style: "merge" | "append" | "morph" | "none";
    block?: boolean | undefined;
    ignore?: boolean | undefined;
    shouldPreserve: ((arg0: Element) => boolean) | NoOp;
    shouldReAppend: ((arg0: Element) => boolean) | NoOp;
    shouldRemove: ((arg0: Element) => boolean) | NoOp;
    afterHeadMorphed: ((arg0: Element, arg1: {
        added: Node[];
        kept: Element[];
        removed: Element[];
    }) => void) | NoOp;
};
export type ConfigCallbacksInternal = {
    beforeNodeAdded: ((arg0: Node) => boolean) | NoOp;
    afterNodeAdded: ((arg0: Node) => void) | NoOp;
    beforeNodeMorphed: ((arg0: Node, arg1: Node) => boolean) | NoOp;
    afterNodeMorphed: ((arg0: Node, arg1: Node) => void) | NoOp;
    beforeNodeRemoved: ((arg0: Node) => boolean) | NoOp;
    afterNodeRemoved: ((arg0: Node) => void) | NoOp;
    beforeAttributeUpdated: ((arg0: string, arg1: Element, arg2: "update" | "remove") => boolean) | NoOp;
};
export type ConfigInternal = {
    morphStyle: "outerHTML" | "innerHTML";
    ignoreActive?: boolean | undefined;
    ignoreActiveValue?: boolean | undefined;
    restoreFocus?: boolean | undefined;
    callbacks: ConfigCallbacksInternal;
    head: ConfigHeadInternal;
};
export type IdSets = {
    persistentIds: Set<string>;
    idMap: Map<Node, Set<string>>;
};
export type Morph = Function;
/**
 * @typedef {object} ConfigHead
 *
 * @property {'merge' | 'append' | 'morph' | 'none'} [style]
 * @property {boolean} [block]
 * @property {boolean} [ignore]
 * @property {function(Element): boolean} [shouldPreserve]
 * @property {function(Element): boolean} [shouldReAppend]
 * @property {function(Element): boolean} [shouldRemove]
 * @property {function(Element, {added: Node[], kept: Element[], removed: Element[]}): void} [afterHeadMorphed]
 */
/**
 * @typedef {object} ConfigCallbacks
 *
 * @property {function(Node): boolean} [beforeNodeAdded]
 * @property {function(Node): void} [afterNodeAdded]
 * @property {function(Element, Node): boolean} [beforeNodeMorphed]
 * @property {function(Element, Node): void} [afterNodeMorphed]
 * @property {function(Element): boolean} [beforeNodeRemoved]
 * @property {function(Element): void} [afterNodeRemoved]
 * @property {function(string, Element, "update" | "remove"): boolean} [beforeAttributeUpdated]
 */
/**
 * @typedef {object} Config
 *
 * @property {'outerHTML' | 'innerHTML'} [morphStyle]
 * @property {boolean} [ignoreActive]
 * @property {boolean} [ignoreActiveValue]
 * @property {boolean} [restoreFocus]
 * @property {ConfigCallbacks} [callbacks]
 * @property {ConfigHead} [head]
 */
/**
 * @typedef {function} NoOp
 *
 * @returns {void}
 */
/**
 * @typedef {object} ConfigHeadInternal
 *
 * @property {'merge' | 'append' | 'morph' | 'none'} style
 * @property {boolean} [block]
 * @property {boolean} [ignore]
 * @property {(function(Element): boolean) | NoOp} shouldPreserve
 * @property {(function(Element): boolean) | NoOp} shouldReAppend
 * @property {(function(Element): boolean) | NoOp} shouldRemove
 * @property {(function(Element, {added: Node[], kept: Element[], removed: Element[]}): void) | NoOp} afterHeadMorphed
 */
/**
 * @typedef {object} ConfigCallbacksInternal
 *
 * @property {(function(Node): boolean) | NoOp} beforeNodeAdded
 * @property {(function(Node): void) | NoOp} afterNodeAdded
 * @property {(function(Node, Node): boolean) | NoOp} beforeNodeMorphed
 * @property {(function(Node, Node): void) | NoOp} afterNodeMorphed
 * @property {(function(Node): boolean) | NoOp} beforeNodeRemoved
 * @property {(function(Node): void) | NoOp} afterNodeRemoved
 * @property {(function(string, Element, "update" | "remove"): boolean) | NoOp} beforeAttributeUpdated
 */
/**
 * @typedef {object} ConfigInternal
 *
 * @property {'outerHTML' | 'innerHTML'} morphStyle
 * @property {boolean} [ignoreActive]
 * @property {boolean} [ignoreActiveValue]
 * @property {boolean} [restoreFocus]
 * @property {ConfigCallbacksInternal} callbacks
 * @property {ConfigHeadInternal} head
 */
/**
 * @typedef {Object} IdSets
 * @property {Set<string>} persistentIds
 * @property {Map<Node, Set<string>>} idMap
 */
/**
 * @typedef {Function} Morph
 *
 * @param {Element | Document} oldNode
 * @param {Element | Node | HTMLCollection | Node[] | string | null} newContent
 * @param {Config} [config]
 * @returns {undefined | Node[]}
 */
/**
 *
 * @type {{defaults: ConfigInternal, morph: Morph}}
 */
export var Idiomorph: {
    defaults: ConfigInternal;
    morph: Morph;
};
