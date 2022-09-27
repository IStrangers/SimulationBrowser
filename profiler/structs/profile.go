package structs

import "time"

type Profile struct {
	name      string
	startTime time.Time
	endTime   time.Time
}

func (profile *Profile) GetElapsedTime() time.Duration {
	return profile.endTime.Sub(profile.startTime)
}

func (profile *Profile) GetName() string {
	return profile.name
}
