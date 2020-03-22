package counters2

type alertCounter int

func New(value int) alertCounter {
	return alertCounter(value)
}
