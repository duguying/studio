// Package utils 包注释
package utils

import (
	"fmt"
	"time"

	"github.com/arran4/golang-ical"
)

// GenerateICS 生成日历事件
func GenerateICS(id string, start, end, stamp time.Time, summary, address, description, link, attendee string) {
	cal := ics.NewCalendar()
	cal.SetMethod(ics.MethodRequest)
	event := cal.AddEvent(fmt.Sprintf("%s@ics.duguying.net", id))
	event.SetCreatedTime(time.Now())
	event.SetDtStampTime(stamp)
	event.SetModifiedAt(time.Now())
	event.SetStartAt(start)
	event.SetEndAt(end)
	event.SetSummary(summary)
	event.SetLocation(address)
	event.SetDescription(description)
	event.SetURL(link)
	event.SetOrganizer("studio@ics.duguying.net", ics.WithCN("This Machine"))
	event.AddAttendee(attendee,
		ics.CalendarUserTypeIndividual, ics.ParticipationStatusNeedsAction,
		ics.ParticipationRoleReqParticipant, ics.WithRSVP(true),
	)
	content := cal.Serialize()
	fmt.Println(content)
}
