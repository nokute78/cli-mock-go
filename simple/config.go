/*
   Copyright 2020 Takahiro Yamashita

   Licensed under the Apache License, Version 2.0 (the "License");

   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package main

import (
	"errors"
	"flag"
)

// ConfigArgsMissing represents no Args error
var ConfigNoArgs error = errors.New("No Args")

type Config struct {
	showVersion bool
}

// Pass os.Args[1:]
func Configure(args []string) (*Config, error) {
	ret := &Config{}
	if len(args) < 1 {
		return nil, ConfigNoArgs
	}

	opt := flag.NewFlagSet("simple", flag.ContinueOnError)
	opt.BoolVar(&ret.showVersion, "V", false, "show Version")

	err := opt.Parse(args)

	return ret, err
}
