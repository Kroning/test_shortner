package storage

import (
	"context"
	"fmt"
	cfg "github.com/Kroning/test_shortner/internal/config"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PgDb struct {
	Pool *pgxpool.Pool
}

// Obtains a pgx pool and returns it inside of PgDb struct
func PgConnect(ctx context.Context, cfg cfg.Config) (*PgDb, error) {
	pool, err := GetPool(ctx, cfg)
	if err != nil {
		return nil, err
	}
	return &PgDb{Pool: pool}, nil
}

// Connects to DB and tests connection
func GetPool(ctx context.Context, cfg cfg.Config) (*pgxpool.Pool, error) {
	db := cfg.Db
	dburl := "postgres://" + db.Username + ":" + db.Password + "@" + db.Host + ":" + db.Port + "/" + db.Dbname
	dbpool, err := pgxpool.New(ctx, dburl)
	if err != nil {
		return nil, err
	}
	//defer dbpool.Close() - No need actually

	// In container DB can start a few seconds.
	// Docker with "depends_on" wait for container, but not DB.
	// This is workaround for start up.
	cnt := 0
	for true {
		_, err = dbpool.Acquire(ctx)
		if err != nil {
			cnt++
			if cnt > 5 {
				return nil, err
			}
			fmt.Println("No connect to database, attempt ", cnt)
			time.Sleep(2 * time.Second)
			continue
		}
		fmt.Println("DB connection succesfull")
		break
	}

	return dbpool, err
}

// Returns all links that have alias LIKE '%alias%'
func (db *PgDb) LoadLinksByAlias(ctx context.Context, alias string) ([][]string, error) {
	result := [][]string{}
	query := "SELECT id, alias, url, TO_CHAR(created_at, 'yyyy-mm-dd hh:mm:ss') FROM links WHERE deleted_at IS NULL and alias LIKE $1;"
	rows, err := db.Pool.Query(ctx, query, "%"+alias+"%")
	defer rows.Close()
	if err != nil {
		log.Println("Error ", err, " while executing query ", query)
		return nil, fmt.Errorf("Error %s while executing query %s", err, query)
	}
	for rows.Next() {
		var id, alias, url, created string
		err = rows.Scan(&id, &alias, &url, &created)
		if err != nil {
			log.Println("Error ", err, " while Scan query ", query)
			return nil, fmt.Errorf("Error %s while executing query %s", err, query)
		}
		slice := []string{id, alias, url, created}
		result = append(result, slice)
	}
	if err = rows.Err(); err != nil {
		log.Println("Error ", err, " while Scan query ", query)
		return nil, fmt.Errorf("Error %s while calling rows.Next() in query %s", err, query)
	}
	return result, nil
}

func (db *PgDb) CheckLinkExistance(ctx context.Context, alias string) (string, error) {
	query := "SELECT url FROM links WHERE alias = $1 and deleted_at IS NULL;"
	row := db.Pool.QueryRow(ctx, query, alias)
	var url string
	err := row.Scan(&url)
	if err == pgx.ErrNoRows {
		return "", LinkNotExists
	} else if err != nil {
		log.Println("Error ", err, " while Scan query in CheckLinkExistance")
		return "", err
	}
	return url, nil
}

func (db *PgDb) InsertLink(ctx context.Context, alias string, url string) error {
	query := "INSERT INTO links VALUES(default,$1,$2,NOW(),NULL);"
	_, err := db.Pool.Exec(ctx, query, alias, url)
	if err != nil {
		log.Println("Link innsertion error: ", err)
		return err
	}
	return nil
}
