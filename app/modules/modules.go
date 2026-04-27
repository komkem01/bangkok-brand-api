package modules

import (
	"log/slog"
	"sync"

	"bangkok-brand/app/modules/district"
	"bangkok-brand/app/modules/entities"
	"bangkok-brand/app/modules/example"
	"bangkok-brand/app/modules/gender"
	"bangkok-brand/app/modules/prefix"
	"bangkok-brand/app/modules/province"
	"bangkok-brand/app/modules/sentry"
	"bangkok-brand/app/modules/specs"
	"bangkok-brand/app/modules/subdistrict"
	"bangkok-brand/app/modules/zipcode"
	"bangkok-brand/internal/config"
	"bangkok-brand/internal/database"
	"bangkok-brand/internal/log"
	"bangkok-brand/internal/otel/collector"

	exampletwo "bangkok-brand/app/modules/example-two"

	appConf "bangkok-brand/config"
	// "bangkok-brand/app/modules/kafka"
)

type Modules struct {
	Conf   *config.Module[appConf.Config]
	Specs  *specs.Module
	Log    *log.Module
	OTEL   *collector.Module
	Sentry *sentry.Module
	DB     *database.DatabaseModule
	ENT    *entities.Module
	// Kafka *kafka.Module
	Example     *example.Module
	Example2    *exampletwo.Module
	Gender      *gender.Module
	Prefix      *prefix.Module
	Province    *province.Module
	District    *district.Module
	Subdistrict *subdistrict.Module
	Zipcode     *zipcode.Module
}

func modulesInit() {
	confMod := config.New(&appConf.App)
	specsMod := specs.New(config.Conf[specs.Config](confMod.Svc))
	conf := confMod.Svc.Config()

	logMod := log.New(config.Conf[log.Option](confMod.Svc))
	otel := collector.New(config.Conf[collector.Config](confMod.Svc))
	log := log.With(slog.String("module", "modules"))

	sentryMod := sentry.New(config.Conf[sentry.Config](confMod.Svc))

	db := database.New(conf.Database.Sql)
	entitiesMod := entities.New(db.Svc.DB())
	exampleMod := example.New(config.Conf[example.Config](confMod.Svc), entitiesMod.Svc)
	exampleMod2 := exampletwo.New(config.Conf[exampletwo.Config](confMod.Svc), entitiesMod.Svc)
	genderMod := gender.New(config.Conf[gender.Config](confMod.Svc), entitiesMod.Svc)
	prefixMod := prefix.New(config.Conf[prefix.Config](confMod.Svc), entitiesMod.Svc)
	provinceMod := province.New(config.Conf[province.Config](confMod.Svc), entitiesMod.Svc)
	districtMod := district.New(config.Conf[district.Config](confMod.Svc), entitiesMod.Svc)
	subdistrictMod := subdistrict.New(config.Conf[subdistrict.Config](confMod.Svc), entitiesMod.Svc)
	zipcodeMod := zipcode.New(config.Conf[zipcode.Config](confMod.Svc), entitiesMod.Svc)
	// kafka := kafka.New(&conf.Kafka)
	mod = &Modules{
		Conf:        confMod,
		Specs:       specsMod,
		Log:         logMod,
		OTEL:        otel,
		Sentry:      sentryMod,
		DB:          db,
		ENT:         entitiesMod,
		Example:     exampleMod,
		Example2:    exampleMod2,
		Gender:      genderMod,
		Prefix:      prefixMod,
		Province:    provinceMod,
		District:    districtMod,
		Subdistrict: subdistrictMod,
		Zipcode:     zipcodeMod,
		// Kafka: kafka,
	}

	log.Infof("all modules initialized")
}

var (
	once sync.Once
	mod  *Modules
)

func Get() *Modules {
	once.Do(modulesInit)

	return mod
}
