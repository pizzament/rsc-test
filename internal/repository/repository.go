package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pizzament/rsc-test/internal/model"
)

type Repository struct {
	pool *pgxpool.Pool
}

// NewRepository инициализируем репозиторий.
func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{pool: pool}
}

// AddCount метод репозитория для увеличения счётчика для bannerID.
func (r *Repository) AddCount(ctx context.Context, bannerID model.BannerID, timestamp time.Time) error {
	const query = `
	INSERT INTO logs(banner_id, time_stamp, count)
	VALUES ($1, $2, 1)
	ON CONFLICT (banner_id, time_stamp) DO UPDATE SET count = logs.count + 1;`

	_, err := r.pool.Exec(ctx, query, bannerID, timestamp)
	if err != nil {
		return err
	}

	return nil
}

// ReceiveStats метод репозитория для получения данных по bannerID.
func (r *Repository) ReceiveStats(ctx context.Context, bannerID model.BannerID, from time.Time, to time.Time) ([]model.Stat, error) {
	const query = `
	SELECT time_stamp, count
	FROM logs
	WHERE banner_id = $1 AND time_stamp >= $2 AND time_stamp < $3
	ORDER BY time_stamp;`

	rows, err := r.pool.Query(ctx, query, bannerID, from, to)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// маппим ответ бд в []model.Stat
	var stats []model.Stat
	for rows.Next() {
		var stat model.Stat
		if err := rows.Scan(&stat.Timestamp, &stat.Count); err != nil {
			return nil, err
		}
		stats = append(stats, stat)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return stats, nil
}
