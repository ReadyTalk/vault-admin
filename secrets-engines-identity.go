package main

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"path/filepath"
	"strings"
	// "github.com/davecgh/go-spew/spew"
)

type EntityConfig struct {
	Entity        map[string]interface{} `json:"entity,omitempty"`
	EntityAliases []EntityAlias          `json:"entity-aliases,omitempty"`
}

type EntityAlias struct {
	Name          string `json:"name"`
	MountPath     string `json:"mount_path"`
	MountAccessor string `json:"mount_accessor"`
}

func configureIdentitySecretsEngine(secretsEngine SecretsEngine) {

	entityList := make(map[string]string)

	files, err := ioutil.ReadDir(Spec.ConfigurationPath + "/secrets-engines/" + secretsEngine.Path + "entities")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {

		success, content := getJsonFile(Spec.ConfigurationPath + "/secrets-engines/" + secretsEngine.Path + "entities/" + file.Name())
		if success {
			var config EntityConfig

			filename := file.Name()
			entity_name := filename[0 : len(filename)-len(filepath.Ext(filename))]
			err = json.Unmarshal([]byte(content), &config)
			if err != nil {
				log.Fatal("Error parsing entity  ["+secretsEngine.Path+"entities/"+entity_name+"]", err)
			}

			// Create/Update the Entity (by Name)
			_, err := Vault.Write("/identity/entity/name/"+entity_name, config.Entity)
			if err != nil {
				log.Fatal("Error writing entity ", entity_name, err)
			}
			entityList[entity_name] = entity_name

		} else {
			log.Warn("Identity entity file has wrong extension.  Will not be processed: ", file.Name())
		}
	}

	// Cleanup Entities
	cleanupEntities(secretsEngine, entityList)
}

func cleanupEntities(secretsEngine SecretsEngine, entityList map[string]string) {

	_, existing_entities := getSecretList(secretsEngine.Path + "entity/name")
	for _, v := range existing_entities {

		if !strings.HasPrefix(v, "entity_") {
			if _, ok := entityList[v]; ok {
				log.Debug("Entity [" + v + "] exists in configuration, no cleanup necessary")
			} else {
				log.Info("Entity [" + v + "] does not exist in configuration, prompting to delete")
				if askForConfirmation("Delete entity ["+v+"] [y/n]?: ", 3) {
					_, err := Vault.Delete(secretsEngine.Path + "entity/name/" + v)
					if err != nil {
						log.Fatal("Error deleting entity ["+v+"] ", err)
					}
					log.Info("Entity [" + v + "] deleted")
				} else {
					log.Info("Leaving entity [" + v + "] even though it is not in config")
				}
			}
		}
	}
}

// LookupAuthMountAccessor returns the accessor_id of the auth mount
// configured at `path`
func LookupAuthMountAccessor(path string) string {
	authMounts, _ := VaultSys.ListAuth()
	for mountPath, mount := range authMounts {
		if mountPath == path {
			return mount.Accessor
		}
	}

	return ""
}
