// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"game-custom-com/internal/dao/internal"
)

// internalFeedbackDao is internal type for wrapping internal DAO implements.
type internalFeedbackDao = *internal.FeedbackDao

// feedbackDao is the data access object for table feedback.
// You can define custom methods on it to extend its functionality as you wish.
type feedbackDao struct {
	internalFeedbackDao
}

var (
	// Feedback is globally public accessible object for table feedback operations.
	Feedback = feedbackDao{
		internal.NewFeedbackDao(),
	}
)

// Fill with you ideas below.
