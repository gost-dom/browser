export declare const DSP: string;
export declare const DSS: string;
export declare const DATASTAR = "datastar";
export declare const DATASTAR_REQUEST = "Datastar-Request";
export declare const DefaultSseRetryDurationMs = 1000;
export declare const DefaultExecuteScriptAttributes = "type module";
export declare const DefaultFragmentsUseViewTransitions = false;
export declare const DefaultMergeSignalsOnlyIfMissing = false;
export declare const DefaultExecuteScriptAutoRemove = true;
export declare const FragmentMergeModes: {
    readonly Morph: "morph";
    readonly Inner: "inner";
    readonly Outer: "outer";
    readonly Prepend: "prepend";
    readonly Append: "append";
    readonly Before: "before";
    readonly After: "after";
    readonly UpsertAttributes: "upsertAttributes";
};
export declare const DefaultFragmentMergeMode: "morph";
export declare const EventTypes: {
    readonly MergeFragments: "datastar-merge-fragments";
    readonly MergeSignals: "datastar-merge-signals";
    readonly RemoveFragments: "datastar-remove-fragments";
    readonly RemoveSignals: "datastar-remove-signals";
    readonly ExecuteScript: "datastar-execute-script";
};
