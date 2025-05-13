import { tagHas, tagToMs } from './tags';
export function debounce(callback, wait, leading = false, trailing = true) {
    let timer = -1;
    const resetTimer = () => timer && clearTimeout(timer);
    return (...args) => {
        resetTimer();
        if (leading && !timer) {
            callback(...args);
        }
        timer = setTimeout(() => {
            if (trailing) {
                callback(...args);
            }
            resetTimer();
        }, wait);
    };
}
export function throttle(callback, wait, leading = true, trailing = false) {
    let waiting = false;
    return (...args) => {
        if (waiting)
            return;
        if (leading) {
            callback(...args);
        }
        waiting = true;
        setTimeout(() => {
            waiting = false;
            if (trailing) {
                callback(...args);
            }
        }, wait);
    };
}
export function modifyTiming(callback, mods) {
    const debounceArgs = mods.get('debounce');
    if (debounceArgs) {
        const wait = tagToMs(debounceArgs);
        const leading = tagHas(debounceArgs, 'leading', false);
        const trailing = !tagHas(debounceArgs, 'notrail', false);
        callback = debounce(callback, wait, leading, trailing);
    }
    const throttleArgs = mods.get('throttle');
    if (throttleArgs) {
        const wait = tagToMs(throttleArgs);
        const leading = !tagHas(throttleArgs, 'noleading', false);
        const trailing = tagHas(throttleArgs, 'trail', false);
        callback = throttle(callback, wait, leading, trailing);
    }
    return callback;
}
