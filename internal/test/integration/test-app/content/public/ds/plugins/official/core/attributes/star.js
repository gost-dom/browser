import { PluginType, Requirement, } from '../../../../engine/types';
export const Star = {
    type: PluginType.Attribute,
    name: 'star',
    keyReq: Requirement.Denied,
    valReq: Requirement.Denied,
    onLoad: () => {
        alert('YOU ARE PROBABLY OVERCOMPLICATING IT');
    },
};
