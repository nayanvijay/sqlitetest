package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nayanvijay/sqlitetest/pkg"
	"github.com/olekukonko/tablewriter"
)

func main() {

	run()
}

/*main application inteface*/
func run() {
	db, err := pkg.CreateDb("policy.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
main:
	for {

		var choice string
		fmt.Println("Sqlite3 CRUD")
		fmt.Println("1. Create Policy")
		fmt.Println("2. Get All Policies")
		fmt.Println("3. Update New Policy")
		fmt.Println("4. Delete Policy")
		fmt.Println("5. Exit")
		fmt.Println("ENTER YOUR CHOICE:= ")
		if _, err := fmt.Scanf("%s", &choice); err != nil {
			log.Println(err.Error())
			break main
		}

		switch choice {

		case "1":
			data, err := readPolicy()
			if err != nil {
				log.Println(err.Error())
				break
			}
			policy := pkg.NewPolicyFrom(data)
			res, err := db.InsertPolicy(policy)
			if err != nil {
				log.Println(err.Error())
			} else {
				r, _ := res.LastInsertId()
				log.Println("last Inserted id := ", r)
			}

		case "2":
			policyResults, err := db.GetAll()
			if err != nil {
				log.Println(err.Error())
			} else {
				table := tablewriter.NewWriter(os.Stdout)
				table.SetHeader([]string{"Id", "APIVersion", "Kind", "Metadata", "Spec"})
				table.SetRowLine(true)

				for _, v := range policyResults {
					table.Append([]string{v.ID, v.APIVersion, v.Kind, v.Metadata, v.Spec})
				}
				table.Render()
			}
		case "3":
			data, err := readPolicy()
			if err != nil {
				log.Println(err.Error())
				break
			}
			newPolicy := pkg.NewPolicyFrom(data)
			var id string
			fmt.Println("Enter policy id:= ")
			fmt.Scanf("%s", &id)
			if err := db.UpdateById(newPolicy, id); err != nil {
				log.Println(err.Error())
			}

		case "4":
			var id string
			fmt.Println("Enter policy id:= ")
			fmt.Scanf("%s", &id)
			if err := db.DeletePolicyByID(id); err != nil {
				log.Println(err.Error())
			}
		case "5":
			break main
		default:
			fmt.Println("Invalid option")
		}
	}

}

//Gets policy path from stdin
func readPolicy() ([]byte, error) {

	var path string
	fmt.Println("Enter new policy File Path:= ")
	fmt.Scanf("%s", &path)
	return os.ReadFile(path)

}

//NOTE: This is code to add dummy policy from certain path
/*func InsertAll(policies []pkg.Policy) {

	db, err := pkg.CreateDb("policy.db")
	if err != nil {
		log.Println(err.Error())
	}
	defer db.Close()

	for _, v := range policies {
		res, err := db.InsertPolicy(v)
		if err != nil {
			log.Println(err.Error())
			continue
		}
		id, _ := res.LastInsertId()
		log.Println("lastInserted ID:= ", id)
	}
}

func init() {


	filepath.Walk("./tmp", func(path string, info fs.FileInfo, err error) error {
		if filepath.Ext(info.Name()) == ".yaml" {
			data, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			p := pkg.NewPolicyFrom(data)
			policies = append(policies, p)
		}
		return err
	})
}*/
