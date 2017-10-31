package mantisdb

import (
	"database/sql"

	"fmt"
	//"log"

	_ "github.com/denisenkom/go-mssqldb"
)

func TestConn() {
	fmt.Printf("mantisdb.TestConn()\n")
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d",
		"10.8.145.104",    // server
		"db_mantis_app",   // user
		"Hud723KueTbs!hf", // pw
		1433)              // port

	conn, err := sql.Open("mssql", connString)
	if err != nil {
		fmt.Printf("open connection failed: %s\n", err.Error())
	}

	defer conn.Close()

	stmt, err := conn.Prepare("select 1, 'abc'")
	if err != nil {
		//log.Fatal("Prepare failed:", err.Error())
		fmt.Printf("Prepare failed:: %s\n", err.Error())
		return
	}

	defer stmt.Close()

	row := stmt.QueryRow()
	var somenumber int64
	var somechars string
	err = row.Scan(&somenumber, &somechars)
	if err != nil {
		//log.Fatal("Scan failed:", err.Error())
		fmt.Printf("Scan failed: %s\n", err.Error())
		return
	}
	fmt.Printf("somenumber:%d\n", somenumber)
	fmt.Printf("somechars:%s\n", somechars)
}

/*
<add name="NXA.API.Data.Mantis.Properties.Settings.db_mantisConnectionString"
connectionString="Data Source=10.8.145.104;Initial Catalog=db_mantis;
Persist Security Info=True;
User ID=db_mantis_app;
Password=Hud723KueTbs!hf"
providerName="System.Data.SqlClient" />


[global::System.Data.Linq.Mapping.FunctionAttribute(Name="dbo.mts_product_by_product_id")]
public ISingleResult<mts_product_by_product_idResult> mts_product_by_product_id([global::System.Data.Linq.Mapping.ParameterAttribute(DbType="VarChar(50)")] string product_id)
{
IExecuteResult result = this.ExecuteMethodCall(this, ((MethodInfo)(MethodInfo.GetCurrentMethod())), product_id);
return ((ISingleResult<mts_product_by_product_idResult>)(result.ReturnValue));
}

public partial class mts_product_by_product_idResult
{
private int _product_no;

private string _product_name;

private string _product_id;

private string _service_code;

private string _display_name;

private string _product_details;

private System.DateTime _registered_datetime;

private System.DateTime _last_updated_datetime;

private string _inserted_user;

private string _updated_user;

private string _redirect_url;

private string _scope;

private string _secret_key;

private string _nisms_service_code;

private string _nisms_game_type;

private string _corebilling_service_code;

*/
