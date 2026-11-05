package repositories

import (
	"Kaho_BaaS/internal/apps/account/models"
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	gormpostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var testDB *gorm.DB

func TestMain(m *testing.M) {
	ctx := context.Background()

	dbContainer, err := postgres.Run(
		ctx,
		"postgres:17-alpine",
		postgres.WithDatabase("database"),
		postgres.WithUsername("user"),
		postgres.WithPassword("password"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(30*time.Second)),
	)
	if err != nil {
		log.Fatalf("could not start postgres container: %v", err)
	}

	connStr, err := dbContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		log.Fatalf("could not get connection string: %v", err)
	}

	testDB, err = gorm.Open(gormpostgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	if err := testDB.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("could not migrate database: %v", err)
	}

	code := m.Run()

	if err := dbContainer.Terminate(ctx); err != nil {
		log.Fatalf("could not teardown postgres container: %v", err)
	}

	os.Exit(code)
}

func TestFindUserByEmail_EmptyTable(t *testing.T) {
	repository := NewAccountRepository(testDB)

	user, err := repository.FindUserByEmail(context.TODO(), "nobody@demo.com")

	assert.ErrorIs(t, err, gorm.ErrRecordNotFound)
	assert.Nil(t, user)
}

func TestFindUserByEmail_ReturnsMatchingUser(t *testing.T) {
	repository := NewAccountRepository(testDB)

	_, err := repository.Create(context.TODO(), &models.Register{
		Email:    "first@demo.com",
		Password: "hashed-first",
		Name:     "First",
	})
	assert.Nil(t, err)

	_, err = repository.Create(context.TODO(), &models.Register{
		Email:    "second@demo.com",
		Password: "hashed-second",
		Name:     "Second",
	})
	assert.Nil(t, err)

	user, err := repository.FindUserByEmail(context.TODO(), "second@demo.com")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "second@demo.com", user.Email)
	assert.Equal(t, "Second", user.Name)

	user, err = repository.FindUserByEmail(context.TODO(), "unknown@demo.com")

	assert.ErrorIs(t, err, gorm.ErrRecordNotFound)
	assert.Nil(t, user)
}
