"use strict";
const generator = require("./generator");
const opt = require("./config.json");
/**
 * csv 
 * csvDir: "."
 * 
 * database
 * "database" : "dbname",
 * "user": "username",
 * "password" : "password",
 */
(async () => {
    await generator.renderFile(
        opt.run, 
        opt.template_option,
        opt.module_option);
})();
