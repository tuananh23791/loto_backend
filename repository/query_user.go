package repository

import (
	"context"
	"fmt"
	"time"
	ErrorCode "travel/config"
	"travel/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
	"github.com/fatih/structs"
	"go.mongodb.org/mongo-driver/mongo"
)

var collectionUser *mongo.Collection

func InsertUser(user model.User) (int, string, model.User) {
	getCollectionUser()
	errorCode, errorMessage := checkUserIsExist(user)

	if errorCode != ErrorCode.SUCCESS {
		return errorCode, errorMessage, user
	}

	if errorCode == ErrorCode.USER_EXITS {
		return ErrorCode.USER_EXITS, ErrorCode.USER_EXITS_MESSAGE, user
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 5)
	user.Password = string(hash)

	errorCode, errorMessage, tokenString := createJwtToken(user)
	if errorCode != ErrorCode.SUCCESS {
		return errorCode, errorMessage, user
	}
	user.Token = tokenString

	result, err := collectionUser.InsertOne(context.TODO(), structs.Map(user))

	if err != nil {
		fmt.Println("InsertOne error" + err.Error())
		return ErrorCode.SOME_THING_WENT_WRONG, "Insert User error - " + err.Error(), user
	}
	// user.ID = res.InsertedID.(string)
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		user.ID = oid.Hex()
	} else {
		// Not objectid.ObjectID, do what you want
	}
	return ErrorCode.SUCCESS, "", user
}

func checkUserIsExist(user model.User) (int, string) {
	filter := bson.D{{"phone_number", user.PhoneNumber}}
	var result model.User
	err := collectionUser.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return ErrorCode.SUCCESS, ""
		}
		// fmt.Println("checkUserIsExist error" + err.Err().Error())
		return ErrorCode.SOME_THING_WENT_WRONG, "checkUserIsExist - " + err.Error()
	}
	fmt.Println("checkUserIsExist error", result)
	return ErrorCode.USER_EXITS, "User is exist"
}

func Login(user model.User) (int, string, model.User) {
	getCollectionUser()
	var result model.User
	err := collectionUser.FindOne(context.TODO(), bson.D{{"phone_number", user.PhoneNumber}}).Decode(&result)
	if err != nil {
		fmt.Println("sai phone_number", user)
		return ErrorCode.INVALID_PHONE_PASSWORD, "Find User error - " + err.Error(), user
	}
	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password))

	if err != nil {
		fmt.Println("sai password")
		return ErrorCode.INVALID_PHONE_PASSWORD, "Find User error - " + err.Error(), result
	}

	errorCode, errorMessage, tokenString := createJwtToken(user)
	if errorCode != ErrorCode.SUCCESS {
		return errorCode, errorMessage, result
	}

	filter := bson.D{{"phone_number", user.PhoneNumber}}
	update := bson.M{"$set": bson.M{"token": tokenString}}
	_, err = collectionUser.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return ErrorCode.SOME_THING_WENT_WRONG, "Update User error - " + err.Error(), user
	}
	result.Token = tokenString

	return ErrorCode.SUCCESS, "", result
}

func UpdateUser(user model.User, tokenString string) (int, string, model.User) {
	getCollectionUser()
	var result model.User
	errorCode, errorMessage := validateToken(tokenString)

	if errorCode != ErrorCode.SUCCESS {
		return errorCode, errorMessage, result
	}

	// var mapUser map[string]interface{}
	user.PhoneNumber = ""
	user.ID = ""
	user.Token = ""
	user.Password = ""
	user.Role = ""
	mapUser := structs.Map(user)
	fmt.Println("Map user", mapUser)
	delete(mapUser, "ID")

	if len(mapUser) == 0 {
		return ErrorCode.MISSING_FIELD, "Update User - Field Invalid", user
	}

	filter := bson.M{"token": tokenString}
	update := bson.M{"$set": mapUser}
	_, err := collectionUser.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return ErrorCode.SOME_THING_WENT_WRONG, "Update User error - " + err.Error(), user
	}

	err = collectionUser.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return ErrorCode.SOME_THING_WENT_WRONG, "Update User, find user error - " + err.Error(), result
	}

	return ErrorCode.SUCCESS, "", result
}

func createJwtToken(user model.User) (int, string, string) {
	secret := []byte(ErrorCode.Secretkey)

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["fullname"] = user.FullName
	claims["phone_number"] = user.PhoneNumber
	claims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix() //Token hết hạn sau 30 ngay
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	"FullName":    user.FullName,
	// 	"PhoneNumber": user.PhoneNumber,
	// })
	tokenString, err := token.SignedString(secret)
	if err != nil {
		fmt.Println("createJwtToken error", err.Error())
		return ErrorCode.SOME_THING_WENT_WRONG, "createJwtToken error - " + err.Error(), ""
	}

	return ErrorCode.SUCCESS, "", tokenString
}

func getCollectionUser() {
	if collectionUser == nil {
		Connect()
		collectionUser = db.Collection("user")
	}
	// ctx, _ = context.WithTimeout(context.Background(), config.CONNECTION_TIME_OUT*time.Second)
}

func validateToken(tokenString string) (int, string) {
	getCollectionUser()
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(ErrorCode.Secretkey), nil
	})
	if err != nil {
		return ErrorCode.TOKEN_INVALID, "parse token error - " + err.Error()
	}
	// Check expiration times
	now := time.Now().Unix()
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
		if !claims.VerifyExpiresAt(now, false) {
			return ErrorCode.TOKEN_EXPIRED, "Token Expired"
		}
		err := claims.Valid()
		if err != nil {
			return ErrorCode.TOKEN_EXPIRED, "Token Expired"
		}
	} else {
		return ErrorCode.SOME_THING_WENT_WRONG, "claims token fail"
	}

	filter := bson.D{{"token", tokenString}}
	var result model.User
	err = collectionUser.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return ErrorCode.TOKEN_INVALID, "TOKEN INVALID"
		}
		// fmt.Println("checkUserIsExist error" + err.Err().Error())
		return ErrorCode.SOME_THING_WENT_WRONG, "Find Token error - " + err.Error()
	}

	return ErrorCode.SUCCESS, ""
}
