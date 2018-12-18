package main

// func dBAddObject(w HDWallet) bool {
// 	session, err := mgo.Dial("localhost:27017")
// 	if err != nil {
// 		log.Println("Could not connect to mongo: ", err.Error())
// 		return false
// 	}
// 	defer session.Close()
//
// 	// Optional. Switch the session to a monotonic behavior.
// 	session.SetMode(mgo.Monotonic, true)
//
// 	c := session.DB("hdwallet").C("wallet")
// 	_, err = c.UpsertId(w.Username, w)
// 	if err != nil {
// 		log.Println("Error creating Profile: ", err.Error())
// 		return false
// 	}
// 	return true
// }
//
// func dbGetObject(name string) HDWallet {
// 	session, err := mgo.Dial("localhost:27017")
// 	if err != nil {
// 		log.Println("Could not connect to mongo: ", err.Error())
// 		return HDWallet{}
// 	}
// 	defer session.Close()
//
// 	// Optional. Switch the session to a monotonic behavior.
// 	session.SetMode(mgo.Monotonic, true)
// 	c := session.DB("hdwallet").C("wallet")
// 	var wallet HDWallet
// 	//err = c.Find(bson.M{}).All(&wallet)
// 	err = c.Find(bson.M{"username": name}).One(&wallet)
//
// 	return wallet
//
// }
