package domain

import (
	"errors"
	"fmt"
	"time"
)

var (
	ErrNotFound  = errors.New("not found")
	ErrConflict  = errors.New("conflict")
	ErrForbidden = errors.New("forbidden")
	ErrInvalid   = errors.New("invalid input")
)

type Role string

const (
	RoleLead      Role = "lead"
	RoleCurator   Role = "curator"
	RoleAnnotator Role = "annotator"
	RoleReviewer  Role = "reviewer"
)

type Sensitivity string

const (
	Public     Sensitivity = "public"
	Restricted Sensitivity = "restricted"
	Sensitive  Sensitivity = "sensitive"
)

type ProjectStatus string

const (
	Draft    ProjectStatus = "draft"
	Active   ProjectStatus = "active"
	Archived ProjectStatus = "archived"
)

type CorpusStatus string

const (
	Collected CorpusStatus = "collected"
	Cleansed  CorpusStatus = "cleansed"
	Licensed  CorpusStatus = "licensed"
	Released  CorpusStatus = "released"
	Withdrawn CorpusStatus = "withdrawn"
)

type BatchStatus string

const (
	BatchPending   BatchStatus = "pending"
	BatchRunning   BatchStatus = "running"
	BatchPaused    BatchStatus = "paused"
	BatchDone      BatchStatus = "done"
	BatchFailed    BatchStatus = "failed"
	BatchCancelled BatchStatus = "cancelled"
)

type TaskStatus string

const (
	TaskQueued    TaskStatus = "queued"
	TaskRunning   TaskStatus = "running"
	TaskSucceeded TaskStatus = "succeeded"
	TaskFailed    TaskStatus = "failed"
	TaskCancelled TaskStatus = "cancelled"
)

type User struct {
	ID           int64
	Email, Name  string
	Role         Role
	PasswordHash string
	CreatedAt    time.Time
}
type Session struct {
	ID, UserID           int64
	TokenHash            string
	ExpiresAt, RevokedAt *time.Time
	CreatedAt            time.Time
}
type Project struct {
	ID                   int64
	Name, Description    string
	QuotaBytes           int64
	Status               ProjectStatus
	Version              int64
	CreatedAt, UpdatedAt time.Time
}
type Corpus struct {
	ID, ProjectID            int64
	Title, Language, License string
	Sensitivity              Sensitivity
	Status                   CorpusStatus
	Bytes                    int64
	Version                  int64
	CreatedAt, UpdatedAt     time.Time
}
type Batch struct {
	ID, CorpusID          int64
	Name                  string
	Status                BatchStatus
	Priority              int
	Concurrency           int
	Version               int64
	StartedAt, FinishedAt *time.Time
	CreatedAt             time.Time
}
type AnnotationTask struct {
	ID, BatchID, AssigneeID int64
	Segment                 string
	Status                  TaskStatus
	Attempts                int
	LeaseUntil              *time.Time
	Version                 int64
	CreatedAt, UpdatedAt    time.Time
}
type Analysis struct {
	ID, ProjectID, CorpusID, AuthorID int64
	Kind, Source, AIUse               string
	Status                            string
	Result                            string
	CreatedAt                         time.Time
}
type AuditEvent struct {
	ID, ActorID                int64
	EntityType                 string
	EntityID                   int64
	Action, Outcome, RequestID string
	Details                    string
	CreatedAt                  time.Time
}
type LicenseChange struct {
	ID, CorpusID, ActorID     int64
	Previous, Current, Reason string
	CreatedAt                 time.Time
}

func (c Corpus) Transition(next CorpusStatus) error {
	ok := false
	switch c.Status {
	case Collected:
		ok = next == Cleansed
	case Cleansed:
		ok = next == Licensed
	case Licensed:
		ok = next == Released || next == Withdrawn
	case Released:
		ok = next == Withdrawn
	}
	if !ok {
		return fmt.Errorf("%w: corpus transition %s -> %s", ErrInvalid, c.Status, next)
	}
	return nil
}
func (b Batch) Transition(next BatchStatus) error {
	ok := false
	switch b.Status {
	case BatchPending:
		ok = next == BatchRunning || next == BatchCancelled
	case BatchRunning:
		ok = next == BatchPaused || next == BatchDone || next == BatchFailed || next == BatchCancelled
	case BatchPaused:
		ok = next == BatchRunning || next == BatchCancelled
	case BatchFailed:
		ok = next == BatchRunning
	}
	if !ok {
		return fmt.Errorf("%w: batch transition", ErrInvalid)
	}
	return nil
}
func (t AnnotationTask) Transition(next TaskStatus) error {
	ok := false
	switch t.Status {
	case TaskQueued:
		ok = next == TaskRunning || next == TaskCancelled
	case TaskRunning:
		ok = next == TaskSucceeded || next == TaskFailed || next == TaskCancelled
	case TaskFailed:
		ok = next == TaskQueued
	}
	if !ok {
		return fmt.Errorf("%w: task transition", ErrInvalid)
	}
	return nil
}
