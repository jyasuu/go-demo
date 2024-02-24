package main

import (
	"fmt"
	"math"
	"sync"
	"net/http"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"

    "context"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
	
	"github.com/redis/go-redis/v9"
	
)

// Struct definition
type Point struct {
	X, Y int
}

// Function with parameters and return value
func add(a, b int) int {
	return a + b
}

// Function with multiple return values
func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return dividend / divisor, nil
}

// Interface definition
type Shape interface {
	Area() float64
}

// Struct implementing an interface
type Circle struct {
	Radius float64
}

// Method for the Circle struct to implement the Shape interface
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func connectToPostgreSQL() (*gorm.DB, error) {
    dsn := "user=postgres password=postgres dbname=postgres host=localhost port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    return db, nil
}


type User struct {
    ID       uint   `gorm:"primaryKey"`
    Username string `gorm:"unique"`
    Email    string
}

func createUser(db *gorm.DB, user *User) error {
    result := db.Create(user)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

func getUserByID(db *gorm.DB, userID uint) (*User, error) {
    var user User
    result := db.First(&user, userID)
    if result.Error != nil {
        return nil, result.Error
    }
    return &user, nil
}

func updateUser(db *gorm.DB, user *User) error {
    result := db.Save(user)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

func deleteUser(db *gorm.DB, user *User) error {
    result := db.Delete(user)
    if result.Error != nil {
        return result.Error
    }
    return nil
}



// You will be using this Trainer type later in the program
type Trainer struct {
    Name string
    Age  int
    City string
}


var ctx = context.Background()

func ExampleClient() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    err := rdb.Set(ctx, "key", "value", 0).Err()
    if err != nil {
        panic(err)
    }

    val, err := rdb.Get(ctx, "key").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("key", val)

    val2, err := rdb.Get(ctx, "key2").Result()
    if err == redis.Nil {
        fmt.Println("key2 does not exist")
    } else if err != nil {
        panic(err)
    } else {
        fmt.Println("key2", val2)
    }
    // Output: key value
    // key2 does not exist
}


func raceExample1() {
	var w sync.WaitGroup
	var sum = 0
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go func() {
			defer w.Done()
			sum++
		}()
	}
	w.Wait()
	fmt.Println("final sum is", sum)
}

func raceExample2() {
	var w sync.WaitGroup
	var m sync.Mutex
	var sum = 0
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go func() {
			defer w.Done()
			m.Lock()
			sum++
			m.Unlock()
		}()
	}
	w.Wait()
	fmt.Println("final sum is", sum)
}

func raceExample3() {
	var w sync.WaitGroup
	ch := make(chan bool, 1)
	var sum = 0
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go func() {
			defer w.Done()
			ch <- true
			sum++
			<-ch
		}()
	}
	w.Wait()
	fmt.Println("final sum is", sum)
}



func main() {
	// Variable declaration and assignment
	// Type inference
	// Conditional statement
	// Looping construct
	// Array declaration and initialization
	// Slice creation
	// Map declaration and initialization
	// Accessing values
	// Printing values
	// Modifying values
	// Printing updated values
	// Calling a function
	// Error handling
	// Using a struct
	// Printing values
	// Looping through a slice
	// Using an interface
	basicExample()

	raceExample1()
	raceExample1()
	raceExample1()
	raceExample2()
	raceExample2()
	raceExample2()
	raceExample3()
	raceExample3()
	raceExample3()

	
    // Perform database migration
    // Create a user
    // Query user by ID
    // Update user
    // Delete user
    postgresExample()

	// Set client options
	// Connect to MongoDB
	// Check the connection
	// create a value into which the result can be decoded
	// Pass these options to the Find method
	// Here's an array in which you can store the decoded documents
	// Passing bson.D{{}} as the filter matches all documents in the collection
	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	// create a value into which the single document can be decoded
	// Close the cursor once finished
	mongoExample()


	ExampleClient()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
	})
	
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"Daniel": "123456",
		"Sam":    "abc123",
	}))

	authorized.GET("/hello/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")

		firstname := c.DefaultQuery("firstname", "None")
		lastname := c.Query("lastname")

		c.JSON(http.StatusOK, gin.H{
			"name":      name,
			"action":    action,
			"firstname": firstname,
			"lastname":  lastname,
		})
	})
	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func mongoExample() {
	credential := options.Credential{
		Username: "root",
		Password: "123456",
	}

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(credential)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("test").Collection("trainers")

	fmt.Println("Connected to MongoDB!")
	{
		ash := Trainer{"Ash", 10, "Pallet Town"}
		misty := Trainer{"Misty", 10, "Cerulean City"}
		brock := Trainer{"Brock", 15, "Pewter City"}

		insertResult, err := collection.InsertOne(context.TODO(), ash)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Inserted a single document: ", insertResult.InsertedID)

		trainers := []interface{}{misty, brock}

		insertManyResult, err := collection.InsertMany(context.TODO(), trainers)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)

		filter := bson.D{{"name", "Ash"}}

		update := bson.D{
			{"$inc", bson.D{
				{"age", 1},
			}},
		}

		updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

		var result2 Trainer

		err = collection.FindOne(context.TODO(), filter).Decode(&result2)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Found a single document: %+v\n", result2)

	}
	{

		findOptions := options.Find()
		findOptions.SetLimit(2)

		var results []*Trainer

		cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
		if err != nil {
			log.Fatal(err)
		}

		for cur.Next(context.TODO()) {

			var elem Trainer
			err := cur.Decode(&elem)
			if err != nil {
				log.Fatal(err)
			}

			results = append(results, &elem)
		}

		if err := cur.Err(); err != nil {
			log.Fatal(err)
		}

		cur.Close(context.TODO())

		fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
	}

	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}

func postgresExample() {
	db, err := connectToPostgreSQL()
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal(err)
	}

	newUser := &User{Username: "john_doe", Email: "john.doe@example.com"}
	err = createUser(db, newUser)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Created User:", newUser)

	userID := newUser.ID
	user, err := getUserByID(db, userID)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("User by ID:", user)

	user.Email = "updated_email@example.com"
	err = updateUser(db, user)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Updated User:", user)

	err = deleteUser(db, user)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Deleted User:", user)

	// defer db.Close()
}

func basicExample() {
	var x int = 10
	y := 5

	if x > y {
		fmt.Println("x is greater than y")
	} else {
		fmt.Println("y is greater than or equal to x")
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	numbers := [3]int{1, 2, 3}

	slice := numbers[1:3]

	person := map[string]string{
		"name":  "John",
		"age":   "30",
		"city":  "New York",
		"email": "john@example.com",
	}

	name := person["name"]
	age := person["age"]
	city := person["city"]
	email := person["email"]

	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %s\n", age)
	fmt.Printf("City: %s\n", city)
	fmt.Printf("Email: %s\n", email)

	person["age"] = "31"
	person["city"] = "San Francisco"

	fmt.Printf("Updated Age: %s\n", person["age"])
	fmt.Printf("Updated City: %s\n", person["city"])

	sum := add(x, y)

	result, err := divide(10.0, 2.0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	p := Point{X: 1, Y: 2}

	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Point: %+v\n", p)

	for index, value := range slice {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	var shape Shape
	shape = Circle{Radius: 5.0}
	area := shape.Area()
	fmt.Printf("Circle Area: %.2f\n", area)
}