package component

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/shiwano/submarine/server/battle/server/battle/context"
)

func TestEquipmentItem(t *testing.T) {
	Convey("Equipment", t, func() {
		startTime := time.Now()
		params := &context.SubmarineParams{
			TorpedoCount:           2,
			TorpedoCooldownSeconds: 10,
			PingerCooldownSeconds:  20,
		}
		e := NewEquipment(params)

		Convey("#ToApiType", func() {
			Convey("should returns a Equipment message", func() {
				m := e.ToAPIType()
				So(m, ShouldNotBeNil)
				So(m.Coerce(), ShouldBeNil)
			})
		})

		Convey("#TryConsumeTorpedo", func() {
			Convey("should consume torpedo", func() {
				So(e.TryConsumeTorpedo(startTime), ShouldBeTrue)
				So(e.TryConsumeTorpedo(startTime), ShouldBeTrue)
				So(e.TryConsumeTorpedo(startTime), ShouldBeFalse)
			})

			Convey("should start cool-down of consumed torpedo", func() {
				So(e.TryConsumeTorpedo(startTime), ShouldBeTrue)
				So(e.TryConsumeTorpedo(startTime), ShouldBeTrue)
				So(e.TryConsumeTorpedo(startTime.Add(time.Second*10)), ShouldBeTrue)
			})
		})

		Convey("#TryConsumePinger", func() {
			Convey("should consume torpedo", func() {
				So(e.TryConsumePinger(startTime), ShouldBeTrue)
				So(e.TryConsumePinger(startTime), ShouldBeFalse)
			})

			Convey("should start cool-down of consumed torpedo", func() {
				So(e.TryConsumePinger(startTime), ShouldBeTrue)
				So(e.TryConsumePinger(startTime.Add(time.Second*10)), ShouldBeFalse)
				So(e.TryConsumePinger(startTime.Add(time.Second*20)), ShouldBeTrue)
			})
		})
	})
}
