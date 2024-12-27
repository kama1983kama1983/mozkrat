package middleware

import (
    "fmt"
    "log"
    "strconv"
    _ "github.com/mattn/go-sqlite3" // Import your database driver
)

// Dbinsert inserts a new record into the specified table
func Dbinsert(table string, fields []string, values map[string]interface{}) bool {
    query := "INSERT INTO " + table + " (" + fields[0]
    for _, field := range fields[1:] {
        query += ", " + field
    }
    query += ") VALUES (?"
    for range fields[1:] {
        query += ", ?"
    }
    query += ")"
	db := InitDb()
	defer db.Close()
    stmt, err := db.Prepare(query)
    if err != nil {
        log.Fatal(err)
        return false
    }
    defer stmt.Close()

    args := make([]interface{}, len(fields))
    for i, field := range fields {
        args[i] = values[field]
    }

    _, err = stmt.Exec(args...)
    return err == nil
}

// DbDelete deletes a record from the specified table
func DbDelete(table string, id string) bool {
	db := InitDb()
	defer db.Close()
    if id != "" {
        sql := fmt.Sprintf(`DELETE FROM %s WHERE id = ?`, table)
        _, err := db.Exec(sql, id)
        if err != nil {
            log.Fatalf("لم يتم حذف البيانات: %v", err)
            return false
        }
        return true
    }
    return false
}

// DbUpdate updates a record in the specified table
func DbUpdate(table string, id string, fields []string, values map[string]interface{}) bool {
    ids, err := strconv.Atoi(id)
    if err != nil {
        log.Fatalf("الرقم التعريفي خطأ %v", err)
        return false
    }

    query := "UPDATE " + table + " SET "
    for i, field := range fields {
        query += field + " = ?"
        if i < len(fields)-1 {
            query += ", "
        }
    }
    query += " WHERE id = ?"
    db := InitDb()
    defer db.Close()
    stmt, err := db.Prepare(query)
    if err != nil {
        log.Fatal(err)
        return false
    }
    defer stmt.Close()

    args := make([]interface{}, len(fields)+1)
    for i, field := range fields {
        args[i] = values[field]
    }
    args[len(fields)] = ids

    _, err = stmt.Exec(args...)
    return err == nil
}

// DbGetOne retrieves a single record by ID
func DbGetOne(table string, id string) (map[string]interface{}, error) {
    if id != "" {
        ids, err := strconv.Atoi(id)
        if err != nil {
            return nil, fmt.Errorf("خطأ في الرقم التعريفي %v", err)
        }

		sql := fmt.Sprintf(`SELECT * FROM %s WHERE id = %s`, table,id)
		db := InitDb()
		defer db.Close()
		row,err := db.Query(sql)
		if err != nil{
			return nil ,err
		}
        result := make(map[string]interface{})
        columns, err := row.Columns()
        if err != nil {
            return nil, err
        }

        values := make([]interface{}, len(columns))
        for i := range values {
            values[i] = new(interface{})
        }

        err = row.Scan(values...)
        if err != nil {
            return nil, fmt.Errorf("لا يوجد الرقم المسلسل %d", ids)
        }

        for i, col := range columns {
            result[col] = *(values[i].(*interface{}))
        }
        return result, nil
    }
    return nil, fmt.Errorf("خطأ في الرقم التعريفي %s", id)
}

// DbGetAll retrieves all records from the specified table
func DbGetAll(table string) ([]map[string]interface{}, error) {
    sql := fmt.Sprintf("SELECT * FROM %s", table)
	db := InitDb()
	defer db.Close()
    rows, err := db.Query(sql)
    if err != nil {
        return nil, fmt.Errorf("يوجد خطأ في ارسال البيانات %v", err)
    }
    defer rows.Close()

    var results []map[string]interface{}
    columns, err := rows.Columns()
    if err != nil {
        return nil, err
    }

    for rows.Next() {
        record := make(map[string]interface{})
        values := make([]interface{}, len(columns))
        for i := range values {
            values[i] = new(interface{})
        }

        err := rows.Scan(values...)
        if err != nil {
            return nil, err
        }

        for i, col := range columns {
            record[col] = *(values[i].(*interface{}))
        }

        results = append(results, record)
    }

    if err = rows.Err(); err != nil {
        return nil, err
    }

    return results, nil
}