package common

type Module interface {
	GetName() string
	GetDescription() string
	GetSlug() string
	GetWaybarOutput() (*WaybarOutput, error)
	SaveLastRun()
	GetLastRun() string
	WriteOutput() error
	Run() error
	GetRunInterval() int64
	GetRunIntervalOnBattery() int64
	RunCommand(name string, args []string) error
	IsEnabled() bool
	SuspendOnBattery() bool
}
