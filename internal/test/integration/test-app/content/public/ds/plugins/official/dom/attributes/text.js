// Authors: Delaney Gillilan
// Icon: tabler:typography
// Slug: Set the text content of an element
// Description: This attribute sets the text content of an element to the result of the expression.
import { runtimeErr } from '../../../../engine/errors';
import { PluginType, Requirement, } from '../../../../engine/types';
export const Text = {
    type: PluginType.Attribute,
    name: 'text',
    keyReq: Requirement.Denied,
    valReq: Requirement.Must,
    onLoad: (ctx) => {
        const { el, effect, genRX } = ctx;
        const rx = genRX();
        if (!(el instanceof HTMLElement)) {
            runtimeErr('TextInvalidElement', ctx);
        }
        return effect(() => {
            const res = rx(ctx);
            el.textContent = `${res}`;
        });
    },
};
