// Mgmt
// Copyright (C) 2013-2018+ James Shubin and the project contributors
// Written by James Shubin <james@shubin.ca> and the project contributors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"io/ioutil"
	"log"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

func parsePkg(path, filename, templates string) error {
	var c config
	filePath := filepath.Join(path, filename)
	log.Printf("Data: %s", filePath)
	cfgFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	err = yaml.UnmarshalStrict(cfgFile, &c)
	if err != nil {
		return err
	}
	err = parseFuncs(c, path, templates)
	if err != nil {
		return err
	}
	return nil
}
