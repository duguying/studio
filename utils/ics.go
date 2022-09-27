// Package utils 包注释
package utils

import (
	"fmt"
	"time"

	ics "github.com/arran4/golang-ical"
)

// GenerateICS 生成日历事件
func GenerateICS(id string, date, end time.Time, period time.Duration,
	summary, address, description, link, attendee string) string {
	cal := ics.NewCalendar()
	cal.SetMethod(ics.MethodRequest)
	event := cal.AddEvent(fmt.Sprintf("%s@ics.duguying.net", id))
	event.SetCreatedTime(time.Now())
	event.SetDtStampTime(date)
	event.SetModifiedAt(time.Now())
	event.SetStartAt(date)
	event.SetEndAt(date.Add(period))
	event.SetSummary(summary)
	event.SetLocation(address)
	event.SetDescription(description)
	if link != "" {
		event.SetURL(link)
	}
	event.SetOrganizer("studio@ics.duguying.net", ics.WithCN("This Machine"))
	event.AddAttendee(attendee,
		ics.CalendarUserTypeIndividual, ics.ParticipationStatusNeedsAction,
		ics.ParticipationRoleReqParticipant, ics.WithRSVP(true),
	)
	return cal.Serialize()
}
