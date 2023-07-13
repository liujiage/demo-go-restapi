package common

import(
	"github.com/golobby/container/v3"
	"github.com/jmoiron/sqlx"
 
)

var (
	Instance = container.New()
	//todo define project root path
)

/****
Inject object into container 
****/
func init(){
	//init and inject it as a singleton 
	Instance.NamedSingleton("db", func() DBHelper {
		//inital connect to database
		dbHelper := DBModelHelper{}
		dbHelper.GetConnect()
		return &dbHelper
	})
}

/****
Get common object from container
***/
func GetDBConnect() *sqlx.DB{
	//todo got the db ....
	var helper DBHelper
    Instance.NamedResolve(&helper, "db")
	return helper.GetConnect()
}

