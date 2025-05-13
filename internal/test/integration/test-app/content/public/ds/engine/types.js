import { DATASTAR } from './consts';
export var PluginType;
(function (PluginType) {
    PluginType[PluginType["Attribute"] = 1] = "Attribute";
    PluginType[PluginType["Watcher"] = 2] = "Watcher";
    PluginType[PluginType["Action"] = 3] = "Action";
})(PluginType || (PluginType = {}));
export var Requirement;
(function (Requirement) {
    Requirement[Requirement["Allowed"] = 0] = "Allowed";
    Requirement[Requirement["Must"] = 1] = "Must";
    Requirement[Requirement["Denied"] = 2] = "Denied";
    Requirement[Requirement["Exclusive"] = 3] = "Exclusive";
})(Requirement || (Requirement = {}));
export const DATASTAR_SIGNAL_EVENT = `${DATASTAR}-signals`;
