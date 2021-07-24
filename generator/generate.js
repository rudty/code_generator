"use strict"
const path = require("path");
const fs = require("fs");
const ejs = require("ejs");
const F = require("../F");
const csv = require("./csv.js");
const mysql = require("./mysql.js");

const DEFAULT_TEMPLATE_OPTION = Object.freeze({
    outDir: "",
    outExt: "",
    filepath: "",
    parseDir: "",
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

const checkIfFile = (path) => new Promise((resolve, reject) => {
    fs.stat(path, (_, stat) => {
        if (stat && !stat.isDirectory()) {
            reject("file already exists");
        } else {
            resolve();
        }
    });
});

const mkdir = (path) => new Promise((resolve) => {
    fs.mkdir(path, (err) => {
        if (err) {
            resolve(false);
        } else {
            resolve(true);
        }
    });
});

const mustMkdir = async (path) => {
    const r = await mkdir(path);
    if (!r) {
        await checkIfFile(path);
    }
}

const mustRemoveFile = (path) => new Promise((resolve) => {
    fs.unlink(path, (err) => {
        if (err) {
            resolve(false);
        } else {
            resolve(true);
        }
    });
});

const writeFile = (path, data) => new Promise((resolve) => {
    fs.writeFile(path, data, (err) => {
        if (err) {
            resolve(false);
        } else {
            resolve(true);
        }
    });    
});

const buildTemplateOption = (opt) => {
    return {
        ...DEFAULT_TEMPLATE_OPTION,
        ...opt, 
    };
};

exports.get = async (moduleName, templateOption, moduleOption) => {
    templateOption = buildTemplateOption(templateOption);
    const dataset = await getModule(moduleName, moduleOption);
    const renderTexts = [];
    for (const e of dataset) {
        const renderText = await ejs.renderFile(templateOption.filepath, {
            dataset: e,
            option: templateOption,
            F,
        });
        renderTexts.push(renderText);
    }
    return renderTexts;
}

exports.renderFile = async (moduleName, templateOption, moduleOption) => {
    templateOption = buildTemplateOption(templateOption);
    const dataset = await getModule(moduleName, moduleOption);
    await mustMkdir(templateOption.outDir);
    let ext = templateOption.outExt;
    if (ext.length > 0 && ext[0] !== ".") {
        ext = "." + ext;
    }
    for (const e of dataset) {
        const renderText = await ejs.renderFile(templateOption.filepath, {
            dataset: e,
            option: templateOption,
            F,
        });

        const outputFilePath = path.join(templateOption.outDir, e.info.filename + ext);
        await mustRemoveFile(outputFilePath);
        await writeFile(outputFilePath, renderText);
    }
};
