"use strict";

const selectInternal = (name, iter, first) => {
    if (iter.name) {
        return iter.name;
    }

    if (iter.get && (typeof m.get) === "function") {
        return iter.get(name);
    }

    if (first) {
        const r = [];
        for (const e of iter) {
            r.push(selectInternal(name, e, false));
        }
        return r;
    }

    throw new Error("?");
};

/**
 * iterable 에서 인자로 받은 이름을 가진 원소를 추출합니다.
 * @param {String} name 
 * @param {Iterable} iter 
 * @returns 추출한 값의 배열
 */
exports.select = (name, iter) => selectInternal(name, iter, true);
