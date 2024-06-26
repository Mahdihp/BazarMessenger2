package postgres

import (
	"BazarMessenger/common"
	"BazarMessenger/config"
	"BazarMessenger/ent"
	"context"
	"entgo.io/ent/dialect"
	"fmt"
	"log"
)

type PostgresqlDB struct {
	config config.DbConfig
	client *ent.Client
}

func (m *PostgresqlDB) Conn() *ent.Client {
	return m.client
}

func New(appConfig config.DbConfig) *PostgresqlDB {
	ConnetionString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		appConfig.Host, appConfig.DbPort, appConfig.Username, appConfig.DbName, appConfig.Password)

	client, err := ent.Open(dialect.Postgres, ConnetionString)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	DataSeed(client)

	return &PostgresqlDB{config: appConfig, client: client}
}

func DataSeed(client *ent.Client) {
	ctx := context.Background()
	//if err := client.Schema.Create(ctx); err != nil {
	//	log.Fatalf("failed creating schema resources: %v", err)
	//}

	count, err := client.Role.Query().Count(ctx)
	if err != nil {
		log.Fatalf("failed count query: %v", err)
	}
	if count == 0 {
		rols, _ := client.Role.CreateBulk(
			client.Role.Create().SetName(common.ADMIN).SetDescription("ادمین"),
			client.Role.Create().SetName(common.PREMIUM).SetDescription("کاربر پولی"),
			client.Role.Create().SetName(common.USER).SetDescription("کاربر معمولی"),
		).Save(ctx)

		client.User.Create().
			SetMobile("09339466051").
			SetUsername("mahdi").
			SetActive(true).
			SetDeleted(false).
			AddRoles(rols...).
			Save(ctx)
	}
}
