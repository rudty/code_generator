"use strict";
const F = require("../F");
const assert = require("assert");

describe('toPascal', () => {
    const pascalString = "HelloWorld";

    it('==', () => {
        const a = "HelloWorld";
        const b = F.toPascal(a);
        assert.deepStrictEqual(a, pascalString);
        assert.deepStrictEqual(b, pascalString);
    });

    it('h', () => {
        const a = "helloWorld";
        const b = F.toPascal(a);
        assert.deepStrictEqual(b, pascalString);
    });

    it('_', () => {
        const a = "hello_world";
        const b = F.toPascal(a);
        assert.deepStrictEqual(b, pascalString);
    });

    it('_W', () => {
        const a = "hello_World";
        const b = F.toPascal(a);
        assert.deepStrictEqual(b, pascalString);
    });
});

describe('toCamel', () => {
    const camelString = "helloWorld";
    it('==', () => {
        const a = "helloWorld";
        const b = F.toCamel(a);
        assert.deepStrictEqual(a, camelString);
    });

    it('h', () => {
        const a = "HelloWorld";
        const b = F.toCamel(a);
        assert.deepStrictEqual(b, camelString);
    });

    it('_', () => {
        const a = "hello_world";
        const b = F.toCamel(a);
        assert.deepStrictEqual(b, camelString);
    });

    it('_W', () => {
        const a = "hello_World";
        const b = F.toCamel(a);
        assert.deepStrictEqual(b, camelString);
    });
});

describe('toSnake', () => {
    const snakeString = "hello_world";
    it('==', () => {
        const a = "hello_world"
        const b = F.toSnake(a);
        assert.deepStrictEqual(b, snakeString);
    });

    it('h', () => {
        const a = "HelloWorld";
        const b = F.toSnake(a);
        assert.deepStrictEqual(b, snakeString);
    });

    it('_W', () => {
        const a = "hello_World";
        const b = F.toSnake(a);
        assert.deepStrictEqual(b, snakeString);
    });
});
