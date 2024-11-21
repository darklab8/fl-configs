package configs

import (
	"github.com/darklab8/fl-configs/configs/configs_mapped"
	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/data_mapped/equipment_mapped/equip_mapped"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/inireader"
	"github.com/darklab8/fl-configs/configs/configs_settings"
	"github.com/darklab8/fl-configs/configs/configs_settings/logus"
	"github.com/darklab8/go-utils/utils/utils_logus"
)

// ExampleImportIniSection demonstrates how to add section to specific section
func Example_importIniSection() {
	// can be having imperfections related to how to handle comments. To improve some day

	freelancer_folder := configs_settings.Env.FreelancerFolder
	configs := configs_mapped.NewMappedConfigs()
	logus.Log.Debug("scanning freelancer folder", utils_logus.FilePath(freelancer_folder))

	// Reading to ini universal custom format and mapping to ORM objects
	// which have both reading and writing back capabilities
	configs.Read(freelancer_folder)

	var new_section *inireader.Section = &inireader.Section{}
	mapped_gun := &equip_mapped.Gun{}
	mapped_gun.Map(new_section)

	mapped_gun.Nickname.Set("my_gun_nickname")
	mapped_gun.IdsName.Set(3453453)
	mapped_gun.HPGunType.Set("some_hp_type")

	configs.Equip.Files[0].Sections = append(configs.Equip.Files[0].Sections, new_section)

	// Write without Dry Run for writing to files modified values back!
	configs.Write(configs_mapped.IsDruRun(true))
}
