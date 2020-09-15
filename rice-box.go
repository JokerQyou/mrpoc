package main

import (
	"time"

	"github.com/GeertJohan/go.rice/embedded"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    "20200915113656_test1.down.sql",
		FileModTime: time.Unix(1600141083, 0),

		Content: string("DROP TABLE IF EXISTS test;\n"),
	}
	file3 := &embedded.EmbeddedFile{
		Filename:    "20200915113656_test1.up.sql",
		FileModTime: time.Unix(1600141072, 0),

		Content: string("CREATE TABLE IF NOT EXISTS test (\n    id INTEGER PRIMARY KEY,\n    name TEXT NOT NULL UNIQUE\n);\n"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1600141083, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // "20200915113656_test1.down.sql"
			file3, // "20200915113656_test1.up.sql"

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`./sql`, &embedded.EmbeddedBox{
		Name: `./sql`,
		Time: time.Unix(1600141083, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"20200915113656_test1.down.sql": file2,
			"20200915113656_test1.up.sql":   file3,
		},
	})
}
