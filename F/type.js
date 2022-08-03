"use strict";

const toJavaType = (t) => {
    if (t) {
        t = t.toLowerCase();
        if (t.includes("date")) {
            return "LocalDateTime";
        }
        if (t.includes("byte") || t.includes("int8") || t.includes("tinyint")) {
            return "byte";
        }
        if (t.includes("smallint") || t.includes("int16") || t.includes("smallint")) {
            return "short";
        }
        if (t.includes("int64") || t.includes("long") || t.includes("bigint")) {
            return "long";
        }
        if (t.includes("int32") || t.includes("int") || t.includes("enum")) {
            return "int"
        }
        if (t.includes("float64") || t.includes("double")) {
            return "double";
        }
        if (t.includes("float32") || t.includes("float")) {
            return "float";
        }
        if (t.includes("string") || t.includes("char") || t.includes("text")) {
            return "String";
        }
    }
    return t;
};

exports.toJavaType = toJavaType;

const toSharpType = (t) => {
    if (t) {
        t = t.toLowerCase();
        if (t.includes("date")) {
            return "DateTime";
        }
    }
    return toJavaType(t);
};

exports.toSharpType = toSharpType;

exports.javaParseString = (t, value) => {
    const javaType = toJavaType(t);
    if (javaType.includes("int")) {
        return "Integer.parseInt(" + value + ")";
    }
    if (javaType.includes("long")) {
        return "Long.parseLong(" + value + ")";
    }
    if (javaType.includes("float")) {
        return "Float.parseFloat(" + value + ")";
    }
    if (javaType.includes("double")) {
        return "Double.parseDouble(" + value + ")";
    }
    if (javaType.includes("LocalDateTime")) {
        return "LocalDateTime.parse(" + value + ")";
    }
    return value;
};

exports.toSharpReaderMethod = (t) => {
    if (t) {
        t = t.toLowerCase();
        if (t.includes("date")) {
            return "GetDateTime";
        }
        if (t.includes("byte") || t.includes("int8") || t.includes("tinyint")) {
            return "GetByte";
        }
        if (t.includes("smallint") || t.includes("int16") || t.includes("smallint")) {
            return "GetInt16";
        }
        if (t.includes("int64") || t.includes("long") || t.includes("bigint")) {
            return "GetInt64";
        }
        if (t.includes("int32") || t.includes("int") || t.includes("enum")) {
            return "GetInt32"
        }
        if (t.includes("float64") || t.includes("double")) {
            return "GetDouble";
        }
        if (t.includes("float32") || t.includes("float")) {
            return "GetFloat";
        }
        if (t.includes("string") || t.includes("char") || t.includes("text")) {
            return "GetString";
        }
    }
    return t;
};