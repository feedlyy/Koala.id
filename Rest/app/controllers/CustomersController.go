package controllers

import (
	"Koala/Rest/app/models"
	"Koala/Rest/db"
	"encoding/json"
	"github.com/dchest/uniuri"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	en_translations "gopkg.in/go-playground/validator.v10/translations/en"
	"log"
	"net/http"
	"time"
)

func getPwd (s string) []byte {
	return []byte(s)
}

func hashAndSalt (pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
	}

	return string(hash)
}

func Register (w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(405)
		_ = json.NewEncoder(w).Encode("Method Not Allowed")
	} else {
		var customers models.Customers
		var arr_customers []models.Customers //for show lastInserted Data

		//form value input
		id := uniuri.NewLen(10)
		name := r.FormValue("customer_name")
		phone := r.FormValue("phone_number")
		email := r.FormValue("email")
		dob := r.FormValue("dob")
		sex := r.FormValue("sex")
		salt := uniuri.NewLen(10)
		password := r.FormValue("password") + salt
		createdAt := time.Now()

		layout := "2006-01-02"
		dobb, _ := time.Parse(layout, dob) // for inserting later

		bytePwd := getPwd(password) // convert input string pwd + salt into byte
		hashPwd := hashAndSalt(bytePwd) // convert bytePwd into hashed + salt string

		//translator for validator input
		translator := en.New()
		uni := ut.New(translator, translator)

		// this is usually known or extracted from http 'Accept-Language' header
		// also see uni.FindTranslator(...)
		trans, found := uni.GetTranslator("en")
		if !found {
			log.Fatal("translator not found")
		}

		v := validator.New()

		if err := en_translations.RegisterDefaultTranslations(v, trans); err != nil {
			log.Fatal(err)
		}

		_ = v.RegisterTranslation("required", trans, func(ut ut.Translator) error {
			return ut.Add("required", "{0} is a required field", true) // see universal-translator for details
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("required", fe.Field())
			return t
		})

		check := models.Customers{
			CustomerId:       id,
			CustomerName:     name,
			Email: 			  email,
			PhoneNumber: 	  phone,
			Dob: 			  dobb,
			Sex: 			  sex,
			Salt:             salt,
			Password:         hashPwd,
			CreatedDate:      createdAt,
		}

		err := v.Struct(check)

		dbConnect := db.Connect()

		//validation for unique phone numbers
		uniquePhone, errorr := dbConnect.Query("select phone_number from customers where phone_number = $1", phone)
		if errorr != nil {
			log.Fatal(errorr)
		}
		phones := []string{}
		for uniquePhone.Next() {
			if err := uniquePhone.Scan(&customers.PhoneNumber); err != nil {
				log.Fatal(err.Error())
			} else {
				phones = append(phones, customers.PhoneNumber)
			}
		}
		//validation for unique emails
		uniqueEmails, errorrs := dbConnect.Query("select email from customers where email = $1", email)
		if errorrs != nil {
			log.Fatal(errorrs)
		}
		emails := []string{}
		for uniqueEmails.Next() {
			if err := uniqueEmails.Scan(&customers.Email); err != nil {
				log.Fatal(err.Error())
			} else {
				emails = append(emails, customers.Email)
			}
		}

		if err != nil {
			for _, errors := range err.(validator.ValidationErrors) {
				w.WriteHeader(400)
				_ = json.NewEncoder(w).Encode(errors.Translate(trans))
			}
		} else if len(phones) != 0 {
			w.WriteHeader(400)
			_ = json.NewEncoder(w).Encode("Phone Number already registered, please use another one")
		} else if len(emails) != 0 {
			w.WriteHeader(400)
			_ = json.NewEncoder(w).Encode("Email already registered, please use another one")
		} else {
			sqlStatement := `insert into customers (customer_id, customer_name, email, 
                       phone_number, dob, sex, salt, password, create_date)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING customer_id`
			customer_id := ""
			err = dbConnect.QueryRow(sqlStatement, id, name, email, phone, dobb, sex, salt,
				hashPwd, createdAt).Scan(&customer_id)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(customer_id)

			getLast, err := dbConnect.Query("select * from customers where customer_id = $1", customer_id)
			if err != nil {
				log.Fatal(err)
			}

			// get the data from query and add it to customers struct
			for getLast.Next() {
				if err := getLast.Scan(&customers.CustomerId, &customers.CustomerName, &customers.PhoneNumber,
					&customers.Email, &customers.Dob,
					&customers.Sex, &customers.Salt, &customers.Password, &customers.CreatedDate); err != nil {
					log.Fatal(err.Error())
				} else {
					arr_customers = append(arr_customers, customers)
				}
			}
			defer dbConnect.Close()

			type Response struct { // to show the response after register
				Message string
				Data []models.Customers
			}

			var response Response

			response.Message = "Success"
			response.Data = arr_customers

			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(response)

		}
	}


}

func GetToken () {

}

func RefreshToken () {

}
