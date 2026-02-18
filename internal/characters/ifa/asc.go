package ifa

import (
	"github.com/genshinsim/gcsim/pkg/core/attributes"
	"github.com/genshinsim/gcsim/pkg/core/event"
	"github.com/genshinsim/gcsim/pkg/core/player/character"
	"github.com/genshinsim/gcsim/pkg/modifier"
)

const (
	rescueEssentialsKey = "rescue-essentials"
	rescueEssentialsMax = 150
	a4Key               = "ifa-a4-buff"
)

// When Ifa is in the Nightsoul's Blessing state, every 1 Nightsoul Point
// out of the total in his entire party will grant him 1 Rescue Essentials
// point. Rescue Essentials will increase the Swirl, Electro-Charged, and
// Lunar-Charged DMG dealt by all nearby party members.

func (c *char) a1() {
	if c.Base.Ascension < 1 {
		return
	}
}

// When nearby party members trigger Nightsoul Bursts,
// Ifa's Elemental Mastery increases by 80 for 10s.

func (c *char) a4() {
	if c.Base.Ascension < 4 {
		return
	}
	c.Core.Events.Subscribe(event.OnNightsoulBurst, func(args ...any) bool {
		c.AddStatMod(character.StatMod{
			Base:         modifier.NewBaseWithHitlag(a4Key, 10*60),
			AffectedStat: attributes.EM,
			Amount: func() ([]float64, bool) {
				return c.a4Buff, true
			},
		})
		return false
	}, a4Key)
}
