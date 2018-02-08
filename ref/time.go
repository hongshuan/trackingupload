Constants
func After(d Duration) <-chan Time
func Tick(d Duration) <-chan Time
func Sleep(d Duration)
type Duration
    func ParseDuration(s string) (Duration, error)
    func Since(t Time) Duration
    func Until(t Time) Duration
    func (d Duration) Hours() float64
    func (d Duration) Minutes() float64
    func (d Duration) Nanoseconds() int64
    func (d Duration) Round(m Duration) Duration
    func (d Duration) Seconds() float64
    func (d Duration) String() string
    func (d Duration) Truncate(m Duration) Duration
type Location
    func FixedZone(name string, offset int) *Location
    func LoadLocation(name string) (*Location, error)
    func (l *Location) String() string
type ParseError
    func (e *ParseError) Error() string
type Ticker
    func NewTicker(d Duration) *Ticker
    func (t *Ticker) Stop()
type Time
    func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time
    func Now() Time
    func Parse(layout, value string) (Time, error)
    func ParseInLocation(layout, value string, loc *Location) (Time, error)
    func Unix(sec int64, nsec int64) Time

    func (t Time) Add(d Duration) Time
    func (t Time) Sub(u Time) Duration

    func (t Time) AddDate(years int, months int, days int) Time
    func (t Time) AppendFormat(b []byte, layout string) []byte
    func (t Time) Clock() (hour, min, sec int)
    func (t Time) Date() (year int, month Month, day int)

    func (t Time) Before(u Time) bool
    func (t Time) Equal(u Time) bool
    func (t Time) After(u Time) bool

    func (t Time) Format(layout string) string
    func (t *Time) GobDecode(data []byte) error
    func (t Time) GobEncode() ([]byte, error)
    func (t Time) ISOWeek() (year, week int)
    func (t Time) In(loc *Location) Time
    func (t Time) IsZero() bool
    func (t Time) Local() Time
    func (t Time) Location() *Location
    func (t Time) Round(d Duration) Time
    func (t Time) String() string
    func (t Time) Truncate(d Duration) Time
    func (t Time) UTC() Time

    func (t Time) Unix() int64
    func (t Time) UnixNano() int64

    func (t Time) MarshalBinary() ([]byte, error)
    func (t Time) MarshalJSON() ([]byte, error)
    func (t Time) MarshalText() ([]byte, error)

    func (t *Time) UnmarshalBinary(data []byte) error
    func (t *Time) UnmarshalJSON(data []byte) error
    func (t *Time) UnmarshalText(data []byte) error

    func (t Time) Year() int
    func (t Time) Month() Month
    func (t Time) Day() int
    func (t Time) Hour() int
    func (t Time) Minute() int
    func (t Time) Second() int
    func (t Time) Nanosecond() int
    func (t Time) YearDay() int
    func (t Time) Weekday() Weekday

    func (t Time) Zone() (name string, offset int)
type Timer
    func AfterFunc(d Duration, f func()) *Timer
    func NewTimer(d Duration) *Timer
    func (t *Timer) Reset(d Duration) bool
    func (t *Timer) Stop() bool
type Weekday
    func (d Weekday) String() string
type Month
    func (m Month) String() string
