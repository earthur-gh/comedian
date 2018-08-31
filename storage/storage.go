package storage

import (
	"time"

	"github.com/maddevsio/comedian/model"
)

// Storage is interface for all supported storages(e.g. MySQL, Postgresql)
type Storage interface {
	// CreateStandup creates standup entry in database
	CreateStandup(model.Standup) (model.Standup, error)

	// UpdateStandup updates standup entry in database
	UpdateStandup(model.Standup) (model.Standup, error)

	// SelectStandup selects standup entry from database
	SelectStandup(int64) (model.Standup, error)

	// SelectStandupByMessageTS selects standup entry by messageTS from database
	SelectStandupByMessageTS(string) (model.Standup, error)

	// SelectStandupsByChannelID selects standup entry by channel ID from database
	SelectStandupsByChannelID(string) ([]model.Standup, error)

	// SelectStandupsByChannelID selects standup entry by channel ID and time period from database
	SelectStandupsByChannelIDForPeriod(string, time.Time, time.Time) ([]model.Standup, error)

	SelectStandupsFiltered(string, string, time.Time, time.Time) ([]model.Standup, error)
	// SelectStandupsForPeriod selects standup entrys for time period from database

	SelectStandupsForPeriod(dateStart, dateEnd time.Time) ([]model.Standup, error)

	// DeleteStandup deletes standup entry from database
	DeleteStandup(int64) error

	// ListStandups returns array of standup entries from database
	ListStandups() ([]model.Standup, error)

	// CreateStandupUser creates standupUser entry in database
	CreateStandupUser(model.StandupUser) (model.StandupUser, error)

	// Checks if user has admin role
	IsAdmin(string, string) bool

	//FindStandupUser finds standup user
	FindStandupUser(username string) (model.StandupUser, error)

	//FindStandupUsers finds standup users
	FindStandupUsers(username string) ([]model.StandupUser, error)

	//FindStandupUserInChannel finds standup user in channel
	FindStandupUserInChannel(string, string) (model.StandupUser, error)

	//FindStandupUserInChannelByUserID finds standup user in channel by Slack member ID
	FindStandupUserInChannelByUserID(string, string) (model.StandupUser, error)

	//GetNonReporters returns a list of non reporters
	GetNonReporters(string, time.Time, time.Time) ([]model.StandupUser, error)

	//GetNonReporter returns a non reporter
	GetNonReporter(string, string, time.Time, time.Time) ([]model.StandupUser, error)

	//CheckNonReporter checks if a user is non reporter
	CheckNonReporter(model.StandupUser, time.Time, time.Time) (bool, error)

	// DeleteStandupUser deletes standup_users entry from database
	DeleteStandupUserByUsername(string, string) error

	// ListStandupUsersByChannelID returns array of standupUser entries from database
	ListStandupUsersByChannelID(string) ([]model.StandupUser, error)

	// ListAllStandupUsers returns array of standupUser entries from database
	ListAllStandupUsers() ([]model.StandupUser, error)

	// CreateStandupTime creates standup time entry in database
	CreateStandupTime(model.StandupTime) (model.StandupTime, error)

	// DeleteStandupTime deletes time entry from database
	DeleteStandupTime(string) error

	// ListStandupTime returns standup time entry from database
	ListStandupTime(string) (model.StandupTime, error)

	// ListAllStandupTime returns standup time entry for all channels from database
	ListAllStandupTime() ([]model.StandupTime, error)

	//GetAllChannels returns a list of all channels
	GetAllChannels() ([]string, error)
}
