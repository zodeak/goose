package dialectquery

import (
	"fmt"
)

type Quest struct{}

var _ Querier = (*Quest)(nil)

func (qr *Quest) CreateTable(tableName string) string {
	q := `CREATE TABLE %s (
		version_id long NOT NULL,
		is_applied boolean NOT NULL,
		tstamp timestamp NOT NULL,
		is_deleted boolean NOT NULL
	)`
	return fmt.Sprintf(q, tableName)
}

func (qr *Quest) InsertVersion(tableName string) string {
	q := `INSERT INTO %s (version_id, is_applied, tstamp, is_deleted) VALUES ($1, $2, NOW(), FALSE)`
	return fmt.Sprintf(q, tableName)
}

func (qr *Quest) DeleteVersion(tableName string) string {
	q := `UPDATE %s SET is_deleted=TRUE WHERE version_id=$1 AND is_deleted=FALSE`
	return fmt.Sprintf(q, tableName)
}

func (qr *Quest) GetMigrationByVersion(tableName string) string {
	q := `SELECT tstamp, is_applied FROM %s WHERE version_id=$1 AND is_deleted=FALSE ORDER BY tstamp DESC LIMIT 1`
	return fmt.Sprintf(q, tableName)
}

func (qr *Quest) ListMigrations(tableName string) string {
	q := `SELECT version_id, is_applied FROM %s WHERE is_deleted=FALSE ORDER BY version_id DESC`
	return fmt.Sprintf(q, tableName)
}

func (qr *Quest) GetLatestVersion(tableName string) string {
	q := `SELECT max(version_id) FROM %s WHERE is_deleted=FALSE`
	return fmt.Sprintf(q, tableName)
}

func (qr *Quest) TableExists(tableName string) string {
	q := `SELECT * FROM %s LIMIT 1`
	return fmt.Sprintf(q, tableName)
}
