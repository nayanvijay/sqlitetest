package pkg

import (
	"database/sql"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

/*Schema for policy Table*/
var initSchema = `
CREATE TABLE IF NOT EXISTS policy(
id INTEGER PRIMARY KEY AUTOINCREMENT, 
apiversion varchar(20), 
kind varchar(20), 
metadata  varchar(200),
spec varchar(2000)
);
`

/*struct to store query result for policy*/
type PolicyResult struct {
	ID         string
	APIVersion string
	Kind       string
	Metadata   string
	Spec       string
}

/*Creates the the Table with Given Schema*/
func (db *DBconn) CreateTables(schema string) error {
	_, err := db.conn.Exec(schema)
	return err
}

/*Struct wrapper around *sql.Db*/
type DBconn struct {
	conn *sql.DB
}

/*Creates the db if Not found in given location*/
func CreateDb(dsn string) (DBconn, error) {
	_, err := os.Stat(dsn)
	if os.IsNotExist(err) {
		log.Println("Creating sqlite database ", dsn)
		_, err := os.OpenFile(dsn, os.O_CREATE|os.O_WRONLY, 0660)
		if err != nil {
			return DBconn{}, err
		}
	}
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return DBconn{}, err
	}
	conc := DBconn{conn: db}
	err = conc.CreateTables(initSchema)
	return conc, err
}

/*insert given policy in table*/
func (db *DBconn) InsertPolicy(p Policy) (sql.Result, error) {

	stm, err := db.conn.Prepare("INSERT INTO policy(apiversion,kind,metadata,spec) values(?,?,?,?)")
	if err != nil {
		return nil, err
	}

	return stm.Exec(
		p.APIVersion,
		p.Kind,
		p.GetMetadata(),
		p.GetSpec(),
	)

}

/*gets all policy form policy table*/
func (db *DBconn) GetAll() ([]PolicyResult, error) {
	pResult := []PolicyResult{}

	stm, err := db.conn.Prepare("SELECT * FROM policy")
	if err != nil {
		return []PolicyResult{}, err
	}

	rows, err := stm.Query()
	if err != nil {
		return []PolicyResult{}, err

	}
	defer rows.Close()

	for rows.Next() {
		temp := PolicyResult{}
		if err := rows.Scan(&temp.ID,
			&temp.APIVersion,
			&temp.Kind,
			&temp.Metadata,
			&temp.Spec); err != nil {

			log.Println(err.Error())
			continue
		}
		pResult = append(pResult, temp)

	}
	return pResult, nil
}

/*delete the policy with given id*/
func (db *DBconn) DeletePolicyByID(id string) error {
	stm, err := db.conn.Prepare("DELETE FROM policy WHERE id = ?")
	if err != nil {
		return err
	}
	res, err := stm.Exec(id)
	ra, _ := res.RowsAffected()
	log.Println("Rows Affected:= ", ra)
	return err
}

/*delete the policy with given id*/
func (db *DBconn) UpdateById(np Policy, id string) error {
	stm, err := db.conn.Prepare("UPDATE policy SET apiversion=?,kind=?,metadata=?,spec=? WHERE id = ?")
	if err != nil {
		return err
	}
	re, err := stm.Exec(np.APIVersion, np.Kind, np.GetMetadata(), np.GetSpec(), id)
	i, _ := re.RowsAffected()
	log.Println("RowAffected := ", i)
	return err
}

/*close the db connection*/
func (db *DBconn) Close() {
	db.conn.Close()
}
