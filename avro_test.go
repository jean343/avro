package avro_test

import "github.com/jean343/avro"

func ConfigTeardown() {
	// Reset the caches
	avro.DefaultConfig = avro.Config{}.Freeze()
}
