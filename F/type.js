"use strict";

const toJavaType = (t) => {
    if (t) {
        switch (t) {
            case "date":
                return "LocalDateTime";
            case "int32":
                return "int"
            case "int64":
                return "long";
            case "float32":
                return "float";
            case "float64":
                return "double";
            case "string":
                return "String";
        }
    }
    return t;
};

exports.toJavaType = toJavaType;

exports.javaParseString = (t, value) => {
    const javaType = toJavaType(t);
    switch (javaType) {
    case "int":
        return "Integer.parseInt(" + value + ")";
    case "long":
        return "Long.parseLong(" + value + ")";
    case "float":
        return "Float.parseFloat(" + value + ")";
    case "double":
        return "Double.parseDouble(" + value + ")";
    case "LocalDateTime":
        return "LocalDateTime.parse(" + value + ")";
    }
    return value;
};