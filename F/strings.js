"use strict";

const isString = (a) => a.constructor === String;

const toString = (v) => {
    if (isString(v)) {
        return v;
    }

    if (Array.isArray(v)) {
        return v.join(",");
    }
    
    return JSON.stringify(v);
};

const toCamelPascal = (name) => {
    let toUpper = false;
    const r = [];
    for (let i = 1; i < name.length; i++) {
        switch(name[i]) {
        case "-":
        case "_":
            toUpper = true;
            break;
        default:
            if (toUpper) {
                r.push(name[i].toUpperCase());
                toUpper = false;
            } else {
                r.push(name[i]);
            }
        }
    }
    return r.join("");
};

exports.toPascal = (v) => {
    const name = toString(v);
    if (name.length === 0) {
        return "";
    }
    const firstChar = name[0].toUpperCase();
    return firstChar + toCamelPascal(name);
};

exports.toCamel = (v) => {
    const name = toString(v);
    if (name.length === 0) {
        return "";
    }
    const firstChar = name[0].toLowerCase();
    return firstChar + toCamelPascal(name); 
};


exports.toSnake = (v) => {
    const name = toString(v);
    if (name.length === 0) {
        return "";
    } 

    if (name.length === 1) {
        return name.toLowerCase();
    }

    const r = [name[0].toLowerCase()];
    for (let i = 1; i < name.length; i++) { 
        const e = name[i];
        const l = e.toLowerCase();

        if (e !== l && name[i - 1] !== "_") {
            r.push("_"); 
        } 

        r.push(l);
    }

    return r.join("");
};
