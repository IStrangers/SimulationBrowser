package structs

import "time"

type Profiler struct {
	profiles []*Profile
}

func CreateProfiler() *Profiler {
	var profiles []*Profile
	return &Profiler{
		profiles: profiles,
	}
}

func (profiler *Profiler) Start(profileName string) {
	for _, profile := range profiler.profiles {
		if profile.name == profileName {
			profile.startTime = time.Now()
			return
		}
	}

	profiler.profiles = append(profiler.profiles, &Profile{
		name:      profileName,
		startTime: time.Now(),
	})
}

func (profiler *Profiler) Stop(profileName string) {
	for _, profile := range profiler.profiles {
		if profile.name == profileName {
			profile.endTime = time.Now()
		}
	}
}

func (profiler *Profiler) GetProfile(profileName string) *Profile {
	var rProfile *Profile

	for _, profile := range profiler.profiles {
		if profile.name == profileName {
			rProfile = profile
		}
	}

	return rProfile
}
