"use strict"
const fs = require("fs");
const ejs = require("ejs");
const F = require("../F");
const csv = require("./csv.js");
const mysql = require("./mysql.js");


const dirExists = (path) => new Promise((resolve) => {
    fs.access(path, (err) => {
        if (err) {
            resolve(false);
        } else {
            resolve(true);
        }
    });
});

const getModule = (moduleName, ...option) => {
    switch (moduleName) {
        case "csv":
            return csv.get(...option);
        case "mysql":
            return mysql.get(...option);
    }
    throw new Error("not support");
};

exports.get = async (moduleName, templateOption, moduleOption) => {
    const dataset = await getModule(moduleName, moduleOption);
    for (const e of dataset) {
        const render = await ejs.renderFile(templateOption.filepath, {
            dataset: e,
            option: templateOption,
            F,
        });
        console.log(render);
    }
}