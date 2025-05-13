// Authors: Delaney Gillilan
// Icon: material-symbols:clock-loader-60-sharp
// Slug: Sets the indicator signal used when fetching data via SSE
// Description: must be a valid signal name
import { PluginType, Requirement, } from '../../../../engine/types';
import { modifyCasing, trimDollarSignPrefix } from '../../../../utils/text';
import { DATASTAR_SSE_EVENT, FINISHED, STARTED, } from '../shared';
export const Indicator = {
    type: PluginType.Attribute,
    name: 'indicator',
    keyReq: Requirement.Exclusive,
    valReq: Requirement.Exclusive,
    onLoad: ({ el, key, mods, signals, value }) => {
        const signalName = key
            ? modifyCasing(key, mods)
            : trimDollarSignPrefix(value);
        const { signal } = signals.upsertIfMissing(signalName, false);
        const watcher = ((event) => {
            const { type, elId, } = event.detail;
            if (elId !== el.id)
                return;
            switch (type) {
                case STARTED:
                    signal.value = true;
                    break;
                case FINISHED:
                    signal.value = false;
                    // Remove the event listener only when finished, in case the element is removed while the request is still in progress
                    document.removeEventListener(DATASTAR_SSE_EVENT, watcher);
                    break;
            }
        });
        document.addEventListener(DATASTAR_SSE_EVENT, watcher);
    },
};
