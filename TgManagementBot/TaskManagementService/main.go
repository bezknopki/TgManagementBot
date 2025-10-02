package main

import (
	"context"
	"fmt"
	"log"

	"github.com/bezknopki/TgManagementBot/general/task"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TaskManagementService struct {
	db *pgxpool.Pool
}

func NewTaskService(db *pgxpool.Pool) *TaskManagementService {
	return &TaskManagementService{db: db}
}

func ConnectDB(dsn string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("connect db: %w", err)
	}

	return pool, nil
}

func CreateTask(ctx context.Context, db *pgxpool.Pool, task task.Task) (int, error) {
	var id int
	err := db.QueryRow(ctx,
		`INSERT INTO tasks (deadline, responsible_user, assigned_user, created_by, location_id, priority, title, description)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`,
		task.Deadline, task.ResponsibleUser, task.AssignedUser, task.CreatedBy, task.LocationId, task.Priority, task.Title, task.Description).Scan(&id)
	return id, err
}

func GetTask(ctx context.Context, db *pgxpool.Pool, id int) (*task.Task, error) {
	task := &task.Task{}
	err := db.QueryRow(ctx, `SELECT id, date_created, date_completed, deadline, responsible_user, assigned_user, created_by,
	location_id, stage, priority, cancellation_reason, title, message, description FROM tasks WHERE id = $1`, id).
		Scan(&task.Id, &task.DateCreated, &task.DateCompleted, &task.Deadline, &task.ResponsibleUser, &task.AssignedUser,
			&task.CreatedBy, &task.LocationId, &task.Stage, &task.Priority, &task.CancellationReason, &task.Title, &task.Message, &task.Description)

	if err != nil {
		return nil, err
	}
	return task, nil
}

func main() {
	dsn := "postgres://postgres:6969@localhost:5433/tasks_db?sslmode=disable"
	pool, err := ConnectDB(dsn)
	if err != nil {
		log.Fatalf("Cannot connect to DB: %v", err)
	}
	defer pool.Close()

	//tmService := NewTaskService(pool)
}
