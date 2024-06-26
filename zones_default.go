package cycling

var DEFAULT_ZONES = Zones{
	{Level: 1,
		Name:             "Active Recovery",
		Description:      "Very low-level exercise, too low to induce significant physiological adaptations. Minimal sensation of leg effort/fatigue. Requires no concentration to maintain pace, and continuous conversation is possible. Typically used for active recovery after strenuous days or between interval efforts.",
		MinDuration:      10800,
		MaxDuration:      36000,
		RecoveryInterval: 0.0,
		MaxHeartRate:     0.68,
		MaxPower:         0.55},
	{Level: 2,
		Name:             "Endurance",
		Description:      "All day pace. A sensation of leg effort/fatigue is generally low but may rise periodically. Concentration to maintain pace is only required at the highest or longest end of the range. Continuous conversation is still possible. Frequent (daily) training sessions of moderate duration (e.g., two hours) at Level 2 are possible, but complete recovery from very long workouts may take more than 24 hrs.",
		MinDuration:      3600,
		MaxDuration:      14400,
		RecoveryInterval: 0.0,
		MaxHeartRate:     0.83,
		MaxPower:         0.75},
	{Level: 3,
		Name:             "Tempo",
		Description:      "Typical intensity of fartlek workout. Greater sensation of leg effort/fatigue than at Level 2. Requires concentration to maintain alone. Conversation must be somewhat halting. Consecutive days of Level 3 training are still possible if the duration is not excessive and dietary carbohydrate intake is adequate.",
		MinDuration:      1200,
		MaxDuration:      3600,
		RecoveryInterval: 0.25,
		MaxHeartRate:     0.94,
		MaxPower:         0.90},
	{Level: 4,
		Name:             "Lactate Threshold",
		Description:      "Just below to just above TT effort. Continuous sensation of moderate-high leg effort/fatigue. Continuous conversation is difficult at best. Effort sufficiently high that sustained exercise at this level is mentally very taxing therefore typically performed in training as multiple blocks of short duration. Consecutive days of training at Level 4 are possible, but such workouts are generally only performed when sufficiently rested so as to be able to maintain intensity.",
		MinDuration:      600,
		MaxDuration:      1800,
		RecoveryInterval: 0.5,
		MaxHeartRate:     1.05,
		MaxPower:         1.05},
	{Level: 5,
		Name:             "VO2 MAX",
		Description:      "Intense longer intervals intended to increase VO2max. Strong to severe sensations of leg effort/fatigue. Conversation is not possible due to often ‘ragged’ breathing. Should generally be attempted only when adequately recovered, consecutive days of Level 5 work are not necessarily desirable, even if possible.",
		MinDuration:      180,
		MaxDuration:      480,
		RecoveryInterval: 1.0,
		MaxHeartRate:     1.15,
		MaxPower:         1.20},
	{Level: 6,
		Name:             "Anaerobic Capacity",
		Description:      "Short high-intensity intervals designed to increase anaerobic capacity. Severe sensation of leg effort/fatigue, and conversation impossible. Consecutive days of extended Level 6 training are usually not attempted.",
		MinDuration:      30,
		MaxDuration:      180,
		RecoveryInterval: 1.5,
		MaxHeartRate:     1.20,
		MaxPower:         1.40},
	{Level: 7,
		Name:             "Neuromuscular Power",
		Description:      "Very short, very high-intensity efforts (e.g., jumps, standing starts, short sprints) that generally place greater stress on musculoskeletal rather than metabolic systems. Power is useful as a guide, but only in reference to prior similar efforts, not TT pace.",
		MinDuration:      1,
		MaxDuration:      30,
		RecoveryInterval: 10.0,
		MaxHeartRate:     1.30,
		MaxPower:         3.00}}
