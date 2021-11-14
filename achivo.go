func JuegoMongo1(w http.ResponseWriter, r http.Request) {
    w.Header().Set("content-type", "application/json")

    var body map[string]interface{}
    err := json.NewDecoder(r.Body).Decode(&body)
    if err != nil {
        fmt.Println(err)
    }
    data, err := json.Marshal(body)

	//observar la info
    nuevo := string(data)
    fmt.Println(nuevo)

    var juego JuegoMongo
     = json.NewDecoder(r.Body).Decode(&juego)

    clintOptions := options.Client().ApplyURI("mongodb://34.125.189.71:27017")
    client, err := mongo.Connect(context.TODO(), clintOptions)
    if err != nil {
        fmt.Println("Mongo.connect() ERROR: ", err)
        os.Exit(1)
    }
    col := client.Database("squid-game").Collection("games")

    fmt.Println("ClientOptions Type: ", reflect.TypeOf(clintOptions), "\n")

    ctx,  := context.WithTimeout(context.Background(), 15time.Second)

    fmt.Println("Collection Type: ", reflect.TypeOf(col), "\n")

    result, insertErr := col.InsertOne(ctx, bson.D{
		{Key: "ID", Value: body["ID"]},
		{Key: "juego", Value: body["juego"]},
		{Key: "max", Value: body["max"]},
		{Key: "players", Value: body["players"]},
		{Key: "worker", Value: body["worker"]}},
	})
    if insertErr != nil {
        fmt.Println("InsertONE Error:", insertErr)
        os.Exit(1)
    } else {
        fmt.Println("InsertOne() result type: ", reflect.TypeOf(result))
        fmt.Println("InsertOne() api result type: ", result)

        newID := result.InsertedID
        fmt.Println("InsertedOne(), newID", newID)
        fmt.Println("InsertedOne(), newID type:", reflect.TypeOf(newID))

    }

    json.NewEncoder(w).Encode(result)

    //fmt.Fprintf(w, "Se Funcion")
}