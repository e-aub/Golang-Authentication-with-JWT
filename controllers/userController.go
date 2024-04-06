package controllers

import (
	databse "Golang-Authentication-with-JWT/database"
	helper "Golang-Authentication-with-JWT/helpers"
	"Golang-Authentication-with-JWT/models"
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var useCollection *mongo.Collection = databse.OpenCollection(databse.Client, "user")
var validate = validator.New()

func HashPassword() {

}

func VerifyPassword() {}

func Signup() {}

func Login() {}

func GetUsers() {}

func GetUser() {}
