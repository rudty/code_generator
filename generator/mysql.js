"use strict";

const mysql = require("mysql2/promise");
const F = require("../F");

let db;

/**
 * 테이블 이름을 가져옵니다
 * @returns 대상 데이터베이스의 테이블 이름들
 */
const getTableNames = async () => {
    const [rows, _] = await db.query(`
        select 
            table_name as name
        from information_schema.tables 
        where table_schema = ?;`, db.config.database);
    
    return F.select("name", rows);
};

/**
 * 대상 테이블의 컬럼 정보를 가져옵니다.
 * @param {String} tableName 
 * @returns 컬럼 정보
 */
const getColumns = async (tableName) => {
    const [rows, _] = await db.query(`
    select
        column_name as 'name',
        column_type as 'type',
        column_key as 'key',
        case is_nullable when 'YES' then 1 else 0 end as 'isnull'
    from information_schema.columns
    where table_schema = ? and table_name = ?;`,
    [db.config.database, tableName]);
    return rows;
};

exports.get = async (connection) => {
    try {
        db = await mysql.createConnection(connection);

        const tables = [];
        const tableNames = await getTableNames(db.config.database);
        for (const tableName of tableNames) {
            const columns = await getColumns(tableName);
            
            const table = {
                name: tableName,
                columns: columns,
            };

            const maybePkColumn = columns.filter(c => c.key.includes("PR"));

            if (maybePkColumn.length > 0) {
                table.pkColumn = maybePkColumn[0];
                table.columnsWithoutPk = columns.filter(e => e !== table.pkColumn);
            }
            
            tables.push(table);

            return tables;
        }
    } finally {
        if (db) {
            db.destroy();
        }
    }
};