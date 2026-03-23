package app

import "github.com/adm87/msgen/palette"

type TaskStatus uint8

const (
	TaskStatusPending TaskStatus = iota
	TaskStatusInProgress
	TaskStatusCompleted
	TaskStatusFailed
)

const (
	taskStatusPendingRunes    = '◌'
	taskStatusInProgressRunes = '●'
	taskStatusCompletedRunes  = '✔'
	taskStatusFailedRunes     = '✘'
)

func GetStatusIcon(status TaskStatus) string {
	switch status {
	case TaskStatusPending:
		return palette.PendingStatus.Render(string(taskStatusPendingRunes))
	case TaskStatusInProgress:
		return palette.InProgressStatus.Render(string(taskStatusInProgressRunes))
	case TaskStatusCompleted:
		return palette.SuccessStatus.Render(string(taskStatusCompletedRunes))
	case TaskStatusFailed:
		return palette.ErrorStatus.Render(string(taskStatusFailedRunes))
	default:
		return "?"
	}
}
