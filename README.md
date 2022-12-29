
# Gorm Migrator  
![Build Status](https://travis-ci.org/nock/nock.svg)
![Coverage Status](http://img.shields.io/badge/coverage-100%25-brightgreen.svg)
![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)

With this package you can manage migrations and seeds in gorm easily

## Features 🔥 
- Migrate data with history meta table in database
- Migrate specific migration via name of the struct
- Seed data with history meta table in database
- Seed specific data via name of the struct
- Down all migrated data 
- Down the last number of migrated data
- Down all seed data 
- Down the last number of seed data
- Completely struct base
- Cli commands
- Simple usage
- 100% test coverage

## Installation 
Run the following comand in your project

~~~bash  
  go get github.com/hoitek-go/gorm-migrator
~~~

## Usage/Examples  
**Sample migration files:**

1-create-nurse.go
~~~go  
  package migrations

  import (
  	"myproject/db"
  	"myproject/internal/nurse/domain"
  )
  
  type CreateNurse struct{}
  
  func (s CreateNurse) Up() {
  	db.Default.Migrator().CreateTable(&domain.Nurse{})
  }
  
  func (s CreateNurse) Down() {
  	db.Default.Migrator().DropTable(&domain.Nurse{})
  }
~~~ 

2-create-user.go
~~~go  
  package migrations

  import (
  	"myproject/db"
  	"myproject/internal/user/domain"
  )
  
  type CreateUser struct{}
  
  func (s CreateUser) Up() {
  	db.Default.Migrator().CreateTable(&domain.User{})
  }
  
  func (s CreateUser) Down() {
  	db.Default.Migrator().DropTable(&domain.User{})
  }
~~~ 

**Sample seed files:**

1-seed-nurses.go
~~~go  
  package seeds

  import (
  	"myproject/db"
  	"myproject/internal/nurse/domain"
  )
  
  type SeedNurse struct{}

  func (s SeedNurse) Up() {
  	db.Default.Create([]domain.Nurse{
            {Name: "Nurse1"},
            {Name: "Nurse2"},
  	})
  }
  
  func (s SeedNurse) Down() {
  	db.Default.Unscoped().Where("name in ?", []string{"Nurse1", "Nurse2"}).Delete(&domain.Nurse{})
  }

~~~ 

**Write the following codes in your main module:**

~~~go  
  package main

  import (
    "os"

    migrator "github.com/hoitek-go/gorm-migrator"
    "myproject/db"
    "myproject/db/migrations"
    "myproject/db/seeds"
  )

  func main() {
    // Set gorm instance
    migrator.SetGorm(db.Init())

    // Set migrations
    migrator.SetMigrations(
      migrations.CreateNurse{},
      migrations.CreateUser{},
    )
    
    // Set seeds
    migrator.SetSeeds(
      seeds.SeedNurse{},
    )

    // Set arguments
    migrator.SetArgs(os.Args)
  }
~~~  

## Migrate Up

Migrate up all migrations into database

~~~bash  
  go run main.go migrate up
~~~

## Migrate Specific Struct

Migrate specific migration via name of the struct

For example assume that User is a struct:

~~~bash  
  go run main.go migrate up User
~~~

## Migrate Down Last n Number

Migrate down the last number of n migrations from database

~~~bash  
  go run main.go migrate down n
~~~

## Migrate Down All

Migrate down all migrations from database

~~~bash  
  go run main.go migrate down
~~~

## Seed Up

Seed up all seeds into databse

~~~bash  
  go run main.go seed up
~~~

## Seed Specific Struct

Seed specific struct via name of the struct

For example assume that UserSeed is a struct:

~~~bash  
  go run main.go seed up UserSeed
~~~

## Seed Down Last n Number

Seed down the last number of n seeds from database

~~~bash  
  go run main.go seed down n
~~~

## Seed Down All

Seed down all seeds from database

~~~bash  
  go run main.go seed down
~~~

## Run Tests

~~~bash  
  make test
~~~

## Export Test Coverage

~~~bash  
  make testcov
~~~

## Tech Stack  
**Server:** Golang, Gorm
 
## Licence  
[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)  
[![GPLv3 License](https://img.shields.io/badge/License-GPL%20v3-yellow.svg)](https://choosealicense.com/licenses/gpl-3.0/)  
[![AGPL License](https://img.shields.io/badge/license-AGPL-blue.svg)](https://choosealicense.com/licenses/gpl-3.0/)  
 