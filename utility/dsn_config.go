package utility

func getDsn() string {

	host := fetchProperty("host")

	user := fetchProperty("user")

	password := fetchProperty("password")

	dbname := fetchProperty("dbname")

	port := fetchProperty("port")

	sslmode := fetchProperty("sslmode")

	return "host=" + host + " " + "user=" + user + " " + "password=" + password + " " + "dbname=" + dbname + " " + "port=" + port + " " + "sslmode=" + sslmode
}

func fetchProperty(key string) string {
	val, err := resourceManager.GetProperty(key)

	if err != nil {
		Log.Error(err.Error())
		errorMessage, _ := resourceManager.GetProperty("database.connection.error")
		Log.Error(errorMessage)
		panic(errorMessage)
	}

	return val
}
