"use strict";

const path = require("path");
const fs = require("fs");
const csv = require('csv-parse');
const util = require("util");

const parser = util.promisify(csv);
const readFile = util.promisify(fs.readFile);

const checkEach = (dataset, index, fn) => {
    for (let i = 1; i < dataset.length; ++i) {
        if (!fn(dataset[i][index])) {
            return false;
        }
    }
    return true;
}
const isNumberRange = (v, min, max) => {
    const n = Number(v);
    if (n % 1 !== 0) {
        return false;
    }

    if (n < min) {
        return false;
    }

    if (n > max) {
        return false;
    }

    return true; 
};

const isInt32 = (dataset, index) => checkEach(dataset, index, (v) => isNumberRange(v, -2147483648, 2147483647));
const isInt64 = (dataset, index) => checkEach(dataset, index, (v) => isNumberRange(v, -9223372036854775808, 9223372036854775807));
const isFloat = (dataset, index) => checkEach(dataset, index, Number);
const isDate = (dataset, index) => checkEach(dataset, index, (v) => !Number.isNaN(Date.parse(v)))
const isBool = (dataset, index) => checkEach(dataset, index, (v) => {
    const upper = v.toUpperCase();
    if (upper === "true" || upper === "false") {
        return true;
    }
    return false;
});

const getColumnType = (dataset, index) => {
    if (isDate(dataset, index)) {
        return "date";
    }

    if (isInt32(dataset, index)) {
        return "int32";
    }

    if (isInt64(dataset, index)) {
        return "int64";
    }

    if (isFloat(dataset, index)) {
        return "float64";
    }

    if (isBool(dataset, index)) {
        return "bool";
    }

    return "string";
};

const getDefaultType = (dataset) => {
    const types = [];
    for (let i = 0; i < dataset[0].length; ++i) {
        const type = getColumnType(dataset, i);
        types.push({
            name: dataset[0][i],
            type,
        })
    } 
    return types;
};

exports.get = async (filePath, option) => {
    const data = await readFile(filePath);
    
    if (!option) {
        option = {
            bom: true,
            from: 0,
        };
    }

    const dataset = await parser(data, option);

    option.typeInfo = getDefaultType(dataset);
    option.dataset = dataset;
    option.filepath = filePath;
    option.filename = path.basename(filePath, ".csv");
    
    return option;
};