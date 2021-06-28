package playerdb

import (
	"github.com/df-mc/dragonfly/server/player"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/google/uuid"
	"time"
)

func fromJson(d jsonData) player.Data {
	return player.Data{
		UUID:            uuid.MustParse(d.UUID),
		Username:        d.Username,
		Position:        d.Position,
		Velocity:        d.Velocity,
		Yaw:             d.Yaw,
		Pitch:           d.Pitch,
		Health:          d.Health,
		MaxHealth:       d.MaxHealth,
		Hunger:          d.Hunger,
		FoodTick:        d.FoodTick,
		ExhaustionLevel: d.ExhaustionLevel,
		SaturationLevel: d.SaturationLevel,
		XpLevel:         d.XpLevel,
		XpTotal:         d.XpTotal,
		XpPercentage:    d.XpPercentage,
		XpSeed:          d.XpSeed,
		GameMode:        dataToGameMode(d.GameMode),
		Effects:         dataToEffects(d.Effects),
		FireTicks:       d.FireTicks,
		FallDistance:    d.FallDistance,
		Inventory:       dataToInv(d.Inventory),
	}
}

func toJson(d player.Data) jsonData {
	return jsonData{
		UUID:            d.UUID.String(),
		Username:        d.Username,
		Position:        d.Position,
		Velocity:        d.Velocity,
		Yaw:             d.Yaw,
		Pitch:           d.Pitch,
		Health:          d.Health,
		MaxHealth:       d.MaxHealth,
		Hunger:          d.Hunger,
		FoodTick:        d.FoodTick,
		ExhaustionLevel: d.ExhaustionLevel,
		SaturationLevel: d.SaturationLevel,
		XpLevel:         d.XpLevel,
		XpTotal:         d.XpTotal,
		XpPercentage:    d.XpPercentage,
		XpSeed:          d.XpSeed,
		GameMode:        gameModeToData(d.GameMode),
		Effects:         effectsToData(d.Effects),
		FireTicks:       d.FireTicks,
		FallDistance:    d.FallDistance,
		Inventory:       invToData(d.Inventory),
	}
}

type jsonData struct {
	UUID                             string
	Username                         string
	Position, Velocity               mgl64.Vec3
	Yaw, Pitch                       float64
	Health, MaxHealth                float64
	Hunger                           int
	FoodTick                         int
	ExhaustionLevel, SaturationLevel float64
	XpLevel, XpTotal                 int
	XpPercentage                     float64
	XpSeed                           int
	GameMode                         uint8
	Inventory                        jsonInventoryData
	Effects                          []jsonEffect
	FireTicks                        int64
	FallDistance                     float64
}

type jsonInventoryData struct {
	Items      []jsonSlot
	Boots      []byte
	Leggings   []byte
	Chestplate []byte
	Helmet     []byte
	OffHand    []byte
	MainHand   uint32
}

type jsonSlot struct {
	Item []byte
	Slot int
}

type jsonEffect struct {
	ID       int
	Level    int
	Duration time.Duration
	Ambient  bool
}
